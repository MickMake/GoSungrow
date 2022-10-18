package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/findPsType"
	"GoSungrow/iSolarCloud/AppService/getAllDeviceByPsId"
	"GoSungrow/iSolarCloud/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"GoSungrow/iSolarCloud/AppService/getRemoteUpgradeTaskList"
	"GoSungrow/iSolarCloud/AppService/getReportData"
	"GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"GoSungrow/iSolarCloud/AppService/psForcastInfo"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/AppService/reportList"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

// ****************************************************** //

func (sg *SunGrow) GetEndpoints(endpoints []string, psIds []valueTypes.Integer, date valueTypes.DateTime, reportType string) error {
	for range Only.Once {
		var data SunGrowData
		data.New(sg)

		if len(endpoints) == 0 {
			fmt.Println("Available endpoints with this command:")
			for _, ep := range data.GetAllEndPoints() {
				fmt.Printf("\t%s\n", ep)
			}
			sg.Error = errors.New("need an endpoint, (or 'all')")
			break
		}

		if endpoints[0] == "all" {
			endpoints = data.GetAllEndPoints()
		}

		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}

		if date.IsZero() {
			date = valueTypes.NewDateTime(valueTypes.Now)
		}
		// fmt.Printf("FilePrefix: %s\n", date.Original())
		// fmt.Printf("String: %s\n", date.String())


		for _, endpoint := range endpoints {
			_, ok := data.Exists(endpoint)
			if !ok {
				sg.Error = errors.New(fmt.Sprintf("EndPoint '%s' does not exist.", endpoint))
				break
			}

			if !data.HasArgs(endpoint) {
				response := data.Get(endpoint, SunGrowDataRequest{ })
				if response.Error != nil {
					break
				}
				// response.Table.SetGraphFilter("")
				sg.Error = response.Table.Output()
				if sg.Error != nil {
					break
				}
				continue
			}

			for _, psId := range psIds {
				response := data.Get(endpoint, SunGrowDataRequest{ PsId: psId, Date: date, ReportType: reportType })
				if response.Error != nil {
					break
				}
				// response.Table.SetGraphFilter("")
				sg.Error = response.Table.Output()
				if sg.Error != nil {
					break
				}
			}
		}
	}

	return sg.Error
}

type SunGrowData struct {
	EndPoint  string
	EndPoints EndPoints
	SunGrow   *SunGrow
	Error     error
}

type EndPoints map[string]EndPoint
type EndPoint struct {
	Func SunGrowDataFunction
	HasArgs bool
}

type SunGrowDataFunction func(request SunGrowDataRequest) SunGrowDataResponse
type SunGrowDataRequest struct {
	PsId       valueTypes.Integer  `json:"ps_id"`
	ReportType string       `json:"report_type"`
	Date       valueTypes.DateTime `json:"date"`
}
type SunGrowDataResponse struct {
	Data     api.DataMap
	Table    output.Table
	Filename string
	Title    string
	Error    error
}


func (sg *SunGrowData) GetAllEndPoints() []string {
	var ret []string
	for ep := range sg.EndPoints {
		ret = append(ret, ep)
	}
	return ret
}

func (sg *SunGrowData) Get(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		dataEndPoint, ok := sg.Exists(endpoint)
		if !ok {
			break
		}

		response = dataEndPoint.Func(request)
		if response.Error != nil {
			break
		}
		if sg.Error != nil {
			break
		}

		if response.Filename == "" {
			response.Filename = endpoint
		}
		if response.Title == "" {
			response.Title = fmt.Sprintf("Data Request %s", endpoint)
		}
		response.Table.SetTitle(response.Title)
		response.Table.SetFilePrefix(response.Filename)
		response.Table.SetGraphFilter("")
		response.Table.SetSaveFile(sg.SunGrow.SaveAsFile)
		response.Table.OutputType = sg.SunGrow.OutputType
	}
	return response
}

func (sg *SunGrowData) Exists(endpoint string) (EndPoint, bool) {
	var dataFunc EndPoint
	var yes bool
	for range Only.Once {
		if dataFunc, yes = sg.EndPoints[endpoint]; yes {
			yes = true
			break
		}
		sg.Error = errors.New(fmt.Sprintf("unknown endpoint function '%s'", endpoint))
	}
	return dataFunc, yes
}

func (sg *SunGrowData) HasArgs(endpoint string) bool {
	var yes bool
	for range Only.Once {
		dataEndPoint, ok := sg.Exists(endpoint)
		if !ok {
			break
		}
		yes = dataEndPoint.HasArgs
	}
	return yes
}


func (sg *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sg.SunGrow = ref
		sg.EndPoints = make(EndPoints)
		sg.EndPoints["getPsList"] = EndPoint { Func: sg.getPsList, HasArgs: false }
		sg.EndPoints["queryDeviceList"] = EndPoint { Func: sg.queryDeviceList, HasArgs: true }
		sg.EndPoints["queryDeviceListForApp"] = EndPoint { Func: sg.queryDeviceListForApp, HasArgs: true }
		sg.EndPoints["getPsDetailWithPsType"] = EndPoint { Func: sg.getPsDetailWithPsType, HasArgs: true }
		sg.EndPoints["getPsDetail"] = EndPoint { Func: sg.getPsDetail, HasArgs: true }
		sg.EndPoints["findPsType"] = EndPoint { Func: sg.findPsType, HasArgs: true }
		// sg.EndPoints["getAllDeviceByPsId"] = EndPoint { Func: sg.getAllDeviceByPsId, HasArgs: false }	// Not working
		sg.EndPoints["getDeviceList"] = EndPoint { Func: sg.getDeviceList, HasArgs: true }
		sg.EndPoints["getIncomeSettingInfos"] = EndPoint { Func: sg.getIncomeSettingInfos, HasArgs: true }
		sg.EndPoints["getKpiInfo"] = EndPoint { Func: sg.getKpiInfo, HasArgs: false }
		sg.EndPoints["getPowerChargeSettingInfo"] = EndPoint { Func: sg.getPowerChargeSettingInfo, HasArgs: true }
		sg.EndPoints["getHouseholdStoragePsReport"] = EndPoint { Func: sg.getHouseholdStoragePsReport, HasArgs: true }
		// sg.EndPoints["getPowerStationBasicInfo"] = EndPoint { Func: sg.getPowerStationBasicInfo, HasArgs: true }	// Not working
		sg.EndPoints["getPowerStationData"] = EndPoint { Func: sg.getPowerStationData, HasArgs: true }
		sg.EndPoints["getPowerStationForHousehold"] = EndPoint { Func: sg.getPowerStationForHousehold, HasArgs: true }
		sg.EndPoints["getPowerStationInfo"] = EndPoint { Func: sg.getPowerStationInfo, HasArgs: true }
		sg.EndPoints["getPowerStatistics"] = EndPoint { Func: sg.getPowerStatistics, HasArgs: true }
		sg.EndPoints["getPsHealthState"] = EndPoint { Func: sg.getPsHealthState, HasArgs: true }
		sg.EndPoints["powerDevicePointList"] = EndPoint { Func: sg.powerDevicePointList, HasArgs: false }

		sg.EndPoints["getPsWeatherList"] = EndPoint { Func: sg.getPsWeatherList, HasArgs: true }
		sg.EndPoints["getRemoteUpgradeTaskList"] = EndPoint { Func: sg.getRemoteUpgradeTaskList, HasArgs: false }	// Not working
		sg.EndPoints["reportList"] = EndPoint { Func: sg.reportList, HasArgs: true }
		sg.EndPoints["getReportData"] = EndPoint { Func: sg.getReportData, HasArgs: true }
		sg.EndPoints["psForcastInfo"] = EndPoint { Func: sg.psForcastInfo, HasArgs: true }
	}
}


func (sg *SunGrowData) getPsList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPsList",
			getPsList.RequestData{ },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getPsList.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getPsList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) queryDeviceList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.queryDeviceList",
			queryDeviceList.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := queryDeviceList.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("queryDeviceList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) queryDeviceListForApp(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.queryDeviceListForApp",
			queryDeviceListForApp.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := queryDeviceListForApp.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("queryDeviceListForApp-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPsDetailWithPsType(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getPsDetailWithPsType.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getPsDetailWithPsType-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPsDetail(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPsDetail",
			getPsDetail.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getPsDetail.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getPsDetail-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) findPsType(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.findPsType",
			findPsType.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := findPsType.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("findPsType-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getAllDeviceByPsId(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getAllDeviceByPsId",
			getAllDeviceByPsId.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getAllDeviceByPsId.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getAllDeviceByPsId-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getDeviceList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getDeviceList",
			getDeviceList.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getDeviceList.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getDeviceList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getIncomeSettingInfos(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getIncomeSettingInfos",
			getIncomeSettingInfos.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getIncomeSettingInfos.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getIncomeSettingInfos-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getKpiInfo(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getKpiInfo",
			getKpiInfo.RequestData{ },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getKpiInfo.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getKpiInfo-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPowerChargeSettingInfo(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerChargeSettingInfo",
			getPowerChargeSettingInfo.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getPowerChargeSettingInfo.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerChargeSettingInfo-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getHouseholdStoragePsReport(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		// fmt.Println(request.Date.Original())
		// {"date_id":"20221001","date_type":"1","ps_id":"1129147"}
		ep := sg.SunGrow.GetByStruct(
			"AppService.getHouseholdStoragePsReport",
			getHouseholdStoragePsReport.RequestData{ PsId: request.PsId, DateType: request.Date.DateType, DateID: request.Date.Original() },
			api.DefaultTimeout,
		)
		if ep.IsError() {
			response.Error = ep.GetError()
			break
		}

		data := getHouseholdStoragePsReport.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getHouseholdStoragePsReport-%d-%s", request.PsId, request.Date.Original())
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPowerStationBasicInfo(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerStationBasicInfo",
			getPowerStationBasicInfo.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPowerStationBasicInfo.Assert(ep)
		if data.Error != nil {
			response.Error = data.Error
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerStationBasicInfo-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// @TODO - Add in the ability to increment values by 5 minutes.
// Return from this function is an array of 288 values.
func (sg *SunGrowData) getPowerStationData(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerStationData",
			getPowerStationData.RequestData{ PsId: request.PsId, DateType: request.Date.DateType, DateID: request.Date.Original() },
			api.DefaultTimeout,
		)

		data := getPowerStationData.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerStationData-%d-%s", request.PsId, request.Date.Original())
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPowerStationForHousehold(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerStationForHousehold",
			getPowerStationForHousehold.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPowerStationForHousehold.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerStationForHousehold-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPowerStationInfo(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerStationInfo",
			getPowerStationInfo.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPowerStationInfo.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerStationInfo-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPowerStatistics(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPowerStatistics",
			getPowerStatistics.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPowerStatistics.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPowerStatistics-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getPsHealthState(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPsHealthState",
			getPsHealthState.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPsHealthState.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPsHealthState-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) powerDevicePointList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.powerDevicePointList",
			powerDevicePointList.RequestData{ },
			api.DefaultTimeout,
		)

		data := powerDevicePointList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("powerDevicePointList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}


func (sg *SunGrowData) getPsWeatherList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getPsWeatherList",
			getPsWeatherList.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getPsWeatherList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getPsWeatherList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getRemoteUpgradeTaskList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getRemoteUpgradeTaskList",
			getRemoteUpgradeTaskList.RequestData{ },	// PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := getRemoteUpgradeTaskList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getRemoteUpgradeTaskList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) reportList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.reportList",
			reportList.RequestData{ PsId: request.PsId, ReportType: request.ReportType },
			api.DefaultTimeout,
		)

		data := reportList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("reportList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) getReportData(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.getReportData",
			getReportData.RequestData{ PsId: request.PsId, ReportType: request.ReportType },
			api.DefaultTimeout,
		)

		data := getReportData.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getReportData-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sg *SunGrowData) psForcastInfo(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sg.SunGrow.GetByStruct(
			"AppService.psForcastInfo",
			psForcastInfo.RequestData{ PsId: request.PsId },
			api.DefaultTimeout,
		)

		data := psForcastInfo.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("psForcastInfo-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// func (sg *SunGrowData) queryAllPsIdAndName(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.queryAllPsIdAndName",
// 			queryAllPsIdAndName.RequestData{  },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryAllPsIdAndName.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryAllPsIdAndName-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) queryPsIdList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.queryPsIdList",
// 			queryPsIdList.RequestData{  },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPsIdList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPsIdList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) queryPsNameByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.queryPsNameByPsId",
// 			queryPsNameByPsId.RequestData{  },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPsNameByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPsNameByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) queryPowerStationInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.queryPowerStationInfo",
// 			queryPowerStationInfo.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPowerStationInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPowerStationInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) queryPsProfit(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.queryPsProfit",
// 			queryPsProfit.RequestData{ PsId: request.PsId, DateID: request.DateID, DateType: request.DateType },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPsProfit.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPsProfit-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsIdState(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.getPsIdState",
// 			getPsIdState.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsIdState.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsIdState-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) showPSView(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.showPSView",
// 			showPSView.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := showPSView.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("showPSView-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getMaxDeviceIdByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.getMaxDeviceIdByPsId",
// 			getMaxDeviceIdByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getMaxDeviceIdByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getMaxDeviceIdByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
