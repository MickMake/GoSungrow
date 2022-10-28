package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/energyTrend"
	"GoSungrow/iSolarCloud/AppService/findCodeValueList"
	"GoSungrow/iSolarCloud/AppService/getAllPsIdByOrgIds"
	"GoSungrow/iSolarCloud/AppService/getAreaList"
	"GoSungrow/iSolarCloud/WebAppService/queryFaultCodes"
	"GoSungrow/iSolarCloud/WebAppService/queryNounList"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


// ****************************************************** //

// func (sg *SunGrow) GetEndpoints(endpoints []string, psIds []string, date string, reportType string, faultTypeCode string) error {

func (sg *SunGrow) GetEndpoints(endpoints []string, args ...string) error {
	for range Only.Once {
		var data SunGrowData
		data.New(sg)

		if len(endpoints) == 0 {
			fmt.Println("Additional commands available, (on top of endpoints):")
			for _, ep := range data.GetAllEndPoints() {
				fmt.Printf("\t%s\n", ep)
			}
			sg.Error = errors.New("need an endpoint, (or command alias)")
			break
		}

		if endpoints[0] == "all" {
			endpoints = data.GetAllEndPoints()
		}

		var request SunGrowDataRequest
		request = request.Set(args...)

		for _, endpoint := range endpoints {
			ep := sg.GetEndpoint(endpoint)
			sg.Error = ep.GetError()
			if sg.Error != nil {
				break
			}

			rargs := ep.GetRequestArgNames()
			// fmt.Printf("args:%s\n", rargs)

			if _, ok := rargs[NamePsId]; !ok {
				// We don't have any request PsId args.
				response := data.GetByApi(ep, request)
				if response.Error != nil {
					sg.Error = response.Error
					break
				}
				if data.Error != nil {
					sg.Error = data.Error
					break
				}
				response.Table.SetFilePrefix(request.GetFilename(endpoint))
				sg.Error = response.Table.Output()
				if sg.Error != nil {
					break
				}
			}

			// This method merges all ps_ids
			if _, ok := rargs[NamePsId]; ok {
				var finalData api.DataMap
				var finalRaw []byte

				var pids valueTypes.PsIds
				pids, sg.Error = sg.GetPsIds()
				if sg.Error != nil {
					break
				}
				request.SetPsIds(pids.Strings())

				for _, psId := range pids {
					request.SetPsId(psId.String())

					response := data.GetByApi(ep, request)
					if response.Error != nil {
						sg.Error = response.Error
						break
					}
					if data.Error != nil {
						sg.Error = data.Error
						break
					}

					finalData.AppendMap(response.Data)
					finalRaw = append(finalRaw, response.Table.GetRawBytes()...)
				}

				if sg.Error != nil {
					fmt.Printf("Error: %s\n", sg.Error)
					break
				}

				request.SetPsId(strings.Join(request.aPsId.Strings(), "_"))	// Hackety hack.
				finalTable := finalData.CreateTable()
				// response.Table.SetTitle(response.Title)
				finalTable.SetFilePrefix(request.GetFilename(endpoint))
				finalTable.SetGraphFilter("")
				finalTable.SetSaveFile(sg.SaveAsFile)
				finalTable.SetRaw(finalRaw)
				finalTable.OutputType = sg.OutputType

				sg.Error = finalTable.Output()
				if sg.Error != nil {
					break
				}
			}

		}
	}

	return sg.Error
}


type EndPoints map[string]EndPoint
type EndPoint struct {
	Func SunGrowDataFunction
	HasArgs bool
}

type SunGrowDataFunction func(request SunGrowDataRequest) SunGrowDataResponse

// SunGrowDataRequest - Collection of all possible request args.
type SunGrowDataRequest struct {
	PsId          *valueTypes.PsId     `json:"ps_id,omitempty"`
	ReportType    *string              `json:"report_type,omitempty"`
	DateId        *valueTypes.DateTime `json:"date_id,omitempty"`
	DateType      *string              `json:"date_type,omitempty"`
	FaultTypeCode *string              `json:"fault_type_code,omitempty"`
	Size          *valueTypes.Integer  `json:"page_size,omitempty"`
	CurPage       *valueTypes.Integer  `json:"cur_page,omitempty"`
	DeviceType    *valueTypes.String   `json:"device_type,omitempty"`
	ReportId      *valueTypes.String   `json:"report_id,omitempty"`
	CodeType      *valueTypes.String   `json:"code_type,omitempty"`
	OrgIds        *valueTypes.String   `json:"orgIds,omitempty"`
	PsIdList      *valueTypes.String   `json:"ps_id_list,omitempty"`

	aPsId         valueTypes.PsIds
}

const (
	NamePsId          = "PsId"
	NameReportType    = "ReportType"
	NameDateId        = "DateId"
	NameDateType      = "DateType"
	NameFaultTypeCode = "FaultTypeCode"
	NameSize          = "Size"
	NameCurPage       = "CurPage"
	NameDeviceType    = "DeviceType"
	NameReportId      = "ReportId"
	NameCodeType      = "CodeType"
	NameOrgIds        = "OrgIds"
	NamePsIdList      = "PsIdList"
)

// MarshalJSON - Convert value to JSON
func (sgd SunGrowDataRequest) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		var dt *string
		if sgd.DateId != nil {
			dt = &sgd.DateId.DateType
		}

		type Parse SunGrowDataRequest
		// Store result from string
		data, err = json.Marshal(Parse {
			PsId:          sgd.PsId,
			ReportType:    sgd.ReportType,
			DateId:        sgd.DateId,
			DateType:      dt,
			FaultTypeCode: sgd.FaultTypeCode,
			Size:          sgd.Size,
			CurPage:       sgd.CurPage,
			DeviceType:    sgd.DeviceType,
			ReportId:      sgd.ReportId,
			CodeType:      sgd.CodeType,
			OrgIds:        sgd.OrgIds,
			PsIdList:      sgd.PsIdList,
		})
		if err == nil {
			break
		}
	}

	return data, err
}

func (sgd *SunGrowDataRequest) Set(args ...string) SunGrowDataRequest {
	var request SunGrowDataRequest
	for range Only.Once {
		for _, arg := range args {
			a := strings.Split(arg, ":")
			if len(a) == 0 {
				continue
			}

			if len(a) == 1 {
				a = append(a, "")
			}

			switch a[0] {
				case NamePsId:
					request.aPsId = valueTypes.SetPsIdStrings(strings.Split(a[1], ","))

				case NameReportType:
					request.ReportType = &a[1]

				case NameDateId:
					val := valueTypes.SetDateTimeString(a[1])
					request.DateId = &val

				case NameFaultTypeCode:
					request.FaultTypeCode = &a[1]

				case NameSize:
					val := valueTypes.SetIntegerString(a[1])
					request.Size = &val

				case NameCurPage:
					val := valueTypes.SetIntegerString(a[1])
					request.CurPage = &val

				case NameDeviceType:
					val := valueTypes.SetStringValue(a[1])
					request.DeviceType = &val

				case NameReportId:
					val := valueTypes.SetStringValue(a[1])
					request.ReportId = &val

				case NameCodeType:
					val := valueTypes.SetStringValue(a[1])
					request.CodeType = &val

				case NameOrgIds:
					val := valueTypes.SetStringValue(a[1])
					request.OrgIds = &val

				case NamePsIdList:
					val := valueTypes.SetStringValue(a[1])
					request.PsIdList = &val
			}
		}
	}
	return request
}

func (sgd *SunGrowDataRequest) Validate(endpoint api.EndPoint) bool {
	ok := true
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			if value != "true" {
				continue
			}
			switch key {
				case NamePsId:
					// if sgd.PsId == nil {
					// 	fmt.Printf("%s is required\n", key)
					// 	ok = false
					// }

				case NameReportType:
					if sgd.ReportType == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameDateId:
					if sgd.DateId == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameFaultTypeCode:
					if sgd.FaultTypeCode == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameSize:
					if sgd.Size == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameCurPage:
					if sgd.CurPage == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameDeviceType:
					if sgd.DeviceType == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameReportId:
					if sgd.ReportId == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameCodeType:
					if sgd.CodeType == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NameOrgIds:
					if sgd.OrgIds == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}

				case NamePsIdList:
					if sgd.PsIdList == nil {
						fmt.Printf("%s is required\n", key)
						ok = false
					}
			}
		}
	}
	return ok
}

// func (sgd *SunGrowDataRequest) Create(endpoint api.EndPoint) SunGrowDataRequest {
// 	var request SunGrowDataRequest
// 	for range Only.Once {
// 		args := endpoint.GetRequestArgNames()
// 		for key, value := range args {
// 			if value != "true" {
// 				continue
// 			}
// 			switch key {
// 				case NamePsId:
// 						request.PsId = sgd.PsId
//
// 				case NameReportType:
// 					request.ReportType = sgd.ReportType
// 					if *request.ReportType == "" {
// 						*request.ReportType = "1"
// 					}
//
// 				case NameDateId:
// 					request.DateId = sgd.DateId
// 					if request.DateId.IsZero() {
// 						did := valueTypes.SetDateTimeString(time.Now().Format(valueTypes.DateTimeLayoutDay))
// 						request.DateId = &did
// 					}
//
// 				case NameFaultTypeCode:
// 					request.FaultTypeCode = sgd.FaultTypeCode
//
// 				case NameSize:
// 					request.Size = sgd.Size
// 					if request.Size.Value() == 0 {
// 						request.Size.SetValue(100)
// 					}
//
// 				case NameCurPage:
// 					request.CurPage = sgd.CurPage
// 					if request.CurPage.Value() == 0 {
// 						request.CurPage.SetValue(1)
// 					}
//
// 				case NameDeviceType:
// 					request.DeviceType = sgd.DeviceType
// 					// if request.DeviceType.String() == "" {
// 					// 	request.DeviceType.SetValue("14")	// @TODO - Need to lookup the first device_type.
// 					// }
//
// 				case NameReportId:
// 					request.ReportId = sgd.ReportId
// 					// if request.ReportId.String() == "" {
// 					// 	request.ReportId.SetValue("8042")	// @TODO - Need to lookup the first device_type.
// 					// }
//
// 				case NameCodeType:
// 					request.CodeType = sgd.CodeType
// 			}
// 		}
// 	}
// 	return request
// }

func (sgd *SunGrowDataRequest) Help(endpoint api.EndPoint) {
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			if key == NamePsId {
				continue
			}

			if key == NameDateType {
				continue
			}

			if value != "true" {
				fmt.Printf("optional - %s:value\n", key)
				continue
			}

			fmt.Printf("required - %s:value\n", key)
		}
	}
}

func (sgd *SunGrowDataRequest) GetFilename(prefix string) string {
	var ret string
	for range Only.Once {
		var aret []string

		aret = append(aret, prefix)
		if sgd.PsId != nil {
			if sgd.PsId.String() != "" {
				aret = append(aret, sgd.PsId.String())
			}
		}

		if sgd.DateId != nil {
			if sgd.DateId.Original() != "" {
				aret = append(aret, sgd.DateId.Original())
			}
		}

		if sgd.ReportType != nil {
			if *sgd.ReportType != "" {
				aret = append(aret, *sgd.ReportType)
			}
		}

		if sgd.FaultTypeCode != nil {
			if *sgd.FaultTypeCode != "" {
				aret = append(aret, *sgd.FaultTypeCode)
			}
		}

		ret = strings.Join(aret, "-")
	}
	return ret
}

func (sgd *SunGrowDataRequest) SetDate(date string) {
	did := valueTypes.SetDateTimeString(date)
	sgd.DateId = &did
	if sgd.DateId.IsZero() {
		did = valueTypes.NewDateTime(valueTypes.Now)
		sgd.DateId = &did
	}
}

func (sgd *SunGrowDataRequest) SetFaultTypeCode(ftc string) {
	sgd.FaultTypeCode = &ftc
}

func (sgd *SunGrowDataRequest) SetReportType(rt string) {
	sgd.ReportType = &rt
}

func (sgd *SunGrowDataRequest) SetPsIds(psIds []string) {
	sgd.aPsId = valueTypes.SetPsIdStrings(psIds)
}

func (sgd *SunGrowDataRequest) GetPsIds() valueTypes.PsIds {
	return sgd.aPsId
}

func (sgd *SunGrowDataRequest) SetPsId(psId string) {
	pid := valueTypes.SetPsIdString(psId)
	sgd.PsId = &pid
}


type SunGrowDataResponse struct {
	Data     api.DataMap
	Table    output.Table
	Filename string
	Title    string
	Error    error
}


type SunGrowData struct {
	// EndPoint  string
	EndPoints EndPoints
	SunGrow   *SunGrow
	Error     error
}

func (sgd *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sgd.SunGrow = ref
		sgd.EndPoints = make(EndPoints)
		// sgd.EndPoints["getPsList"] = EndPoint { Func: sgd.getPsList, HasArgs: false }
		// sgd.EndPoints["queryDeviceList"] = EndPoint { Func: sgd.queryDeviceList, HasArgs: true }
		// sgd.EndPoints["queryDeviceInfo"] = EndPoint { Func: sgd.queryDeviceInfo, HasArgs: true }
		// sgd.EndPoints["queryDeviceListForApp"] = EndPoint { Func: sgd.queryDeviceListForApp, HasArgs: true }
		// sgd.EndPoints["getPsDetailWithPsType"] = EndPoint { Func: sgd.getPsDetailWithPsType, HasArgs: true }
		// sgd.EndPoints["getPsDetail"] = EndPoint { Func: sgd.getPsDetail, HasArgs: true }
		// sgd.EndPoints["findPsType"] = EndPoint { Func: sgd.findPsType, HasArgs: true }
		// // sg.EndPoints["getAllDeviceByPsId"] = EndPoint { Func: sg.getAllDeviceByPsId, HasArgs: false }	// Not working
		// sgd.EndPoints["getDeviceList"] = EndPoint { Func: sgd.getDeviceList, HasArgs: true }
		// sgd.EndPoints["getIncomeSettingInfos"] = EndPoint { Func: sgd.getIncomeSettingInfos, HasArgs: true }
		// sgd.EndPoints["getKpiInfo"] = EndPoint { Func: sgd.getKpiInfo, HasArgs: false }
		// sgd.EndPoints["getPowerChargeSettingInfo"] = EndPoint { Func: sgd.getPowerChargeSettingInfo, HasArgs: true }
		// sgd.EndPoints["getHouseholdStoragePsReport"] = EndPoint { Func: sgd.getHouseholdStoragePsReport, HasArgs: true }
		// // sg.EndPoints["getPowerStationBasicInfo"] = EndPoint { Func: sg.getPowerStationBasicInfo, HasArgs: true }	// Not working
		// sgd.EndPoints["getPowerStationData"] = EndPoint { Func: sgd.getPowerStationData, HasArgs: true }
		// sgd.EndPoints["getPowerStationForHousehold"] = EndPoint { Func: sgd.getPowerStationForHousehold, HasArgs: true }
		// sgd.EndPoints["getPowerStationInfo"] = EndPoint { Func: sgd.getPowerStationInfo, HasArgs: true }
		// sgd.EndPoints["getPowerStatistics"] = EndPoint { Func: sgd.getPowerStatistics, HasArgs: true }
		// sgd.EndPoints["getPsHealthState"] = EndPoint { Func: sgd.getPsHealthState, HasArgs: true }
		// sgd.EndPoints["powerDevicePointList"] = EndPoint { Func: sgd.powerDevicePointList, HasArgs: false }
		// sgd.EndPoints["getPsWeatherList"] = EndPoint { Func: sgd.getPsWeatherList, HasArgs: true }
		// sgd.EndPoints["getRemoteUpgradeTaskList"] = EndPoint { Func: sgd.getRemoteUpgradeTaskList, HasArgs: false } // Not working
		// sgd.EndPoints["reportList"] = EndPoint { Func: sgd.reportList, HasArgs: true }
		// sgd.EndPoints["getReportData"] = EndPoint { Func: sgd.getReportData, HasArgs: true }
		// sgd.EndPoints["psForcastInfo"] = EndPoint { Func: sgd.psForcastInfo, HasArgs: true }
		// sgd.EndPoints["queryPowerStationInfo"] = EndPoint { Func: sgd.queryPowerStationInfo, HasArgs: true }
		// sgd.EndPoints["getPsIdState"] = EndPoint { Func: sgd.getPsIdState, HasArgs: true }
		// sgd.EndPoints["queryPsProfit"] = EndPoint { Func: sgd.queryPsProfit, HasArgs: true }
		// sgd.EndPoints["queryAllPsIdAndName"] = EndPoint { Func: sgd.queryAllPsIdAndName, HasArgs: false }
		// sgd.EndPoints["queryPsIdList"] = EndPoint { Func: sgd.queryPsIdList, HasArgs: false }
		// sgd.EndPoints["queryPsNameByPsId"] = EndPoint { Func: sgd.queryPsNameByPsId, HasArgs: true }
		// sgd.EndPoints["showPSView"] = EndPoint { Func: sgd.showPSView, HasArgs: true }
		// sgd.EndPoints["getMaxDeviceIdByPsId"] = EndPoint { Func: sgd.getMaxDeviceIdByPsId, HasArgs: true }

		sgd.EndPoints["energyTrend"] = EndPoint { Func: sgd.energyTrend, HasArgs: false }
		sgd.EndPoints["getAreaList"] = EndPoint { Func: sgd.getAreaList, HasArgs: false }
		sgd.EndPoints["getAllPsIdByOrgIds"] = EndPoint { Func: sgd.getAllPsIdByOrgIds, HasArgs: false }
		sgd.EndPoints["findCodeValueList"] = EndPoint { Func: sgd.findCodeValueList, HasArgs: false }
		sgd.EndPoints["queryFaultCodes"] = EndPoint { Func: sgd.queryFaultCodes, HasArgs: false }
		sgd.EndPoints["queryNounList"] = EndPoint { Func: sgd.queryNounList, HasArgs: false }
	}
}

func (sgd *SunGrowData) GetAllEndPoints() []string {
	var ret []string
	for ep := range sgd.EndPoints {
		ret = append(ret, ep)
	}
	return ret
}

func (sgd *SunGrowData) GetByFunc(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		dataEndPoint, ok := sgd.FuncExists(endpoint)
		if !ok {
			break
		}

		response = dataEndPoint.Func(request)
		if response.Error != nil {
			break
		}
		if sgd.Error != nil {
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
		response.Table.SetSaveFile(sgd.SunGrow.SaveAsFile)
		response.Table.OutputType = sgd.SunGrow.OutputType
	}
	return response
}

func (sgd *SunGrowData) GetByEndPointName(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	response.Table = output.NewTable()
	for range Only.Once {
		ep := sgd.SunGrow.GetEndpoint(endpoint)
		sgd.SunGrow.Error = ep.GetError()
		if sgd.SunGrow.Error != nil {
			break
		}
		response = sgd.GetByApi(ep, request)
	}
	return response
}

func (sgd *SunGrowData) GetByApi(endpoint api.EndPoint, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	response.Table = output.NewTable()
	for range Only.Once {

		if !request.Validate(endpoint) {
			request.Help(endpoint)
			break
		}

		var req []byte
		req, sgd.Error = json.Marshal(request)
		if sgd.Error != nil {
			fmt.Printf("GetByApi - ERR: %s\n", sgd.Error)
			break
		}
		fmt.Printf("%s\n", req)

		// newRequest := request.Create(endpoint)

		if string(req) != "" {
			endpoint = endpoint.SetRequestByJson(output.Json(req))
			sgd.Error = endpoint.GetError()
			if sgd.Error != nil {
				fmt.Println(endpoint.Help())
				break
			}
		}
		fmt.Printf("%s\n", endpoint.GetRequestJson())

		endpoint = endpoint.Call()
		sgd.Error = endpoint.GetError()
		if sgd.Error != nil {
			if strings.Contains(sgd.Error.Error(), "er_token_login_invalid") {
				sgd.SunGrow.Logout()
				break
			}
			fmt.Println(endpoint.Help())
			break
		}

		// switch {
		// 	case sgd.SunGrow.OutputType.IsNone():
		//
		// 	case sgd.SunGrow.OutputType.IsRaw():
		// 		if sgd.SunGrow.SaveAsFile {
		// 			sgd.SunGrow.Error = ep.WriteDataFile()
		// 			break
		// 		}
		// 		fmt.Println(ep.GetJsonData(true))
		//
		// 	case sgd.SunGrow.OutputType.IsJson():
		// 		if sgd.SunGrow.SaveAsFile {
		// 			sgd.SunGrow.Error = ep.WriteDataFile()
		// 			break
		// 		}
		// 		fmt.Println(ep.GetJsonData(false))
		//
		// 	default:
		// }

		// var fn []string
		// if request.PsId.String() != "" {
		// 	fn = append(fn, request.PsId.String())
		// }
		// if !request.DateId.IsZero() {
		// 	fn = append(fn, request.DateId.Original())
		// }
		response.Filename = request.GetFilename("")
		response.Filename = endpoint.SetFilenamePrefix(response.Filename)
		// response.Filename = request.CreateFilename(ep)
		if response.Title == "" {
			response.Title = fmt.Sprintf("Data Request %s", endpoint)
		}

		response.Data = endpoint.GetEndPointData()
		response.Table = endpoint.GetEndPointDataTable()
		// response.Table.SetTitle(response.Title)
		response.Table.SetFilePrefix(response.Filename)
		response.Table.SetGraphFilter("")
		response.Table.SetSaveFile(sgd.SunGrow.SaveAsFile)
		response.Table.OutputType = sgd.SunGrow.OutputType
	}
	return response
}

func (sgd *SunGrowData) FuncExists(endpoint string) (EndPoint, bool) {
	var dataFunc EndPoint
	var yes bool
	for range Only.Once {
		if dataFunc, yes = sgd.EndPoints[endpoint]; yes {
			yes = true
			break
		}
		sgd.Error = errors.New(fmt.Sprintf("unknown endpoint function '%s'", endpoint))
	}
	return dataFunc, yes
}

func (sgd *SunGrowData) HasArgs(endpoint string) bool {
	var yes bool
	for range Only.Once {
		dataEndPoint, ok := sgd.FuncExists(endpoint)
		if ok {
			yes = dataEndPoint.HasArgs
			break
		}

		ok = sgd.SunGrow.RequestRequiresArgs(endpoint)
	}
	return yes
}


func (sgd *SunGrowData) energyTrend(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"AppService.energyTrend",
			// energyTrend.RequestData{ PsId: request.PsId },
			energyTrend.RequestData{  },
			api.DefaultTimeout,
		)

		data := energyTrend.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("energyTrend-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sgd *SunGrowData) getAreaList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"AppService.getAreaList",
			// energyTrend.RequestData{ PsId: request.PsId },
			getAreaList.RequestData{  },
			api.DefaultTimeout,
		)

		data := getAreaList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getAreaList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// @TODO - Need to add OrgIds
// @TODO - Need to support []string in ResultData
func (sgd *SunGrowData) getAllPsIdByOrgIds(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"AppService.getAllPsIdByOrgIds",
			// getAllPsIdByOrgIds.RequestData{ PsId: request.PsId },
			getAllPsIdByOrgIds.RequestData{ OrgIds: valueTypes.SetStringValue("362245") },
			api.DefaultTimeout,
		)

		data := getAllPsIdByOrgIds.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("getAllPsIdByOrgIds-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// @TODO - No data.
func (sgd *SunGrowData) findCodeValueList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"AppService.findCodeValueList",
			// findCodeValueList.RequestData{ PsId: request.PsId },
			findCodeValueList.RequestData{ CodeType: *request.CodeType },
			api.DefaultTimeout,
		)

		data := findCodeValueList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("findCodeValueList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// @TODO - No data.
func (sgd *SunGrowData) queryFaultCodes(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"WebAppService.queryFaultCodes",
			// queryFaultCodes.RequestData{ PsId: request.PsId },
			queryFaultCodes.RequestData{ FaultName: "417" },
			api.DefaultTimeout,
		)

		data := queryFaultCodes.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("queryFaultCodes-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

func (sgd *SunGrowData) queryNounList(request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		ep := sgd.SunGrow.GetByStruct(
			"WebAppService.queryNounList",
			queryNounList.RequestData{ FaultTypeCode: "718" },
			// queryNounList.RequestData{ },
			api.DefaultTimeout,
		)

		data := queryNounList.Assert(ep)
		if data.Error != nil {
			response.Error = ep.GetError()
			break
		}

		response.Filename = data.SetFilenamePrefix("queryNounList-%d", request.PsId)
		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
	}
	return response
}

// func (sg *SunGrowData) getChnnlListByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getChnnlListByPsId",
// 			getChnnlListByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getChnnlListByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getChnnlListByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getDevInstalledPowerByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getDevInstalledPowerByPsId",
// 			getDevInstalledPowerByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getDevInstalledPowerByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getDevInstalledPowerByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsIdByUserId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsIdByUserId",
// 			getPsIdByUserId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsIdByUserId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsIdByUserId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsInfoWithJoinGridByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsInfoWithJoinGridByPsId",
// 			getPsInfoWithJoinGridByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsInfoWithJoinGridByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsInfoWithJoinGridByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsInstallerByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsInstallerByPsId",
// 			getPsInstallerByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsInstallerByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsInstallerByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsInstallerOrgInfoByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsInstallerOrgInfoByPsId",
// 			getPsInstallerOrgInfoByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsInstallerOrgInfoByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsInstallerOrgInfoByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsKpiForHoursByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsKpiForHoursByPsId",
// 			getPsKpiForHoursByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsKpiForHoursByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsKpiForHoursByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getPsListForPsDataByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.getPsListForPsDataByPsId",
// 			getPsListForPsDataByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsListForPsDataByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsListForPsDataByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) queryPsTypeByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.queryPsTypeByPsId",
// 			queryPsTypeByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPsTypeByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPsTypeByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) selectDeviceTypeByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"WebIscmAppService.selectDeviceTypeByPsId",
// 			selectDeviceTypeByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := selectDeviceTypeByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("selectDeviceTypeByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }
