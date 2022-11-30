package iSolarCloud

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


// ****************************************************** //

func (sg *SunGrow) NewSunGrowData() SunGrowData {
	var data SunGrowData
	for range Only.Once {
		data.New(sg)

		data.SetOutputType(sg.OutputType)
		data.SetSaveAsFile(sg.SaveAsFile)
	}
	return data
}


func SplitArg(arg string) []string {
	var ret []string
	for range Only.Once {
		ret = []string{arg}
		for _, s := range []string{ ",", "/", " "} {
			if strings.Contains(arg, s) {
				ret = strings.Split(arg, s)
				break
			}
		}
	}
	return ret
}


type EndPoints map[string]EndPoint
type EndPoint struct {
	Func SunGrowDataFunction
	HasArgs bool
}

type SunGrowDataFunction func(request SunGrowDataRequest) SunGrowDataResponse


type SunGrowDataResponse struct {
	Data     api.DataMap
	Name     string
	Filename string
	Title    string
	Error    error
}
type SunGrowDataResponses map[string]SunGrowDataResponse

func (sgd *SunGrowDataResponse) GetOutput(outputType output.OutputType, saveAsFile bool) error {
	for range Only.Once {
		if outputType.IsStruct() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			break
		}

		if outputType.IsList() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			break
		}

		if outputType.IsRaw() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			break
		}

		if outputType.IsJson() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			break
		}

		if outputType.IsXLSX() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			break
		}


		if outputType.IsCsv() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			if sgd.Error != nil {
				break
			}
			// break
		}

		if outputType.IsXML() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			if sgd.Error != nil {
				break
			}
			// break
		}

		if outputType.IsTable() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = outputType
			table.SetSaveFile(saveAsFile)
			table.AppendFilePrefix(sgd.Filename)
			table.SetTitle(table.GetName() + " - " + sgd.Title)
			sgd.Error = table.Output()
			if sgd.Error != nil {
				break
			}
			// break
		}


		tables := sgd.Data.CreateDataTables()
		if len(tables) == 0 {
			fmt.Printf("No data table results for '%s'\n", sgd.Title)
			break
		}

		for _, table2 := range tables {
			fmt.Println()
			table2.OutputType = outputType
			table2.SetSaveFile(saveAsFile)
			table2.AppendFilePrefix(sgd.Filename)
			table2.SetTitle(table2.GetName() + " - " + sgd.Title)
			sgd.Error = table2.Output()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}


type SunGrowData struct {
	endPoints  []string
	request    SunGrowDataRequest

	results    SunGrowDataResults

	sunGrow    *SunGrow
	outputType output.OutputType
	saveAsFile bool
	Error      error
}

func (sgd *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sgd.sunGrow = ref
		sgd.results = make(SunGrowDataResults)

		// sgd.EndPoints = make(EndPoints)
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
		//
		// sgd.EndPoints["energyTrend"] = EndPoint { Func: sgd.energyTrend, HasArgs: false }
		// sgd.EndPoints["getAreaList"] = EndPoint { Func: sgd.getAreaList, HasArgs: false }
		// sgd.EndPoints["getAllPsIdByOrgIds"] = EndPoint { Func: sgd.getAllPsIdByOrgIds, HasArgs: false }
		// sgd.EndPoints["findCodeValueList"] = EndPoint { Func: sgd.findCodeValueList, HasArgs: false }
		// sgd.EndPoints["queryFaultCodes"] = EndPoint { Func: sgd.queryFaultCodes, HasArgs: false }
		// sgd.EndPoints["queryNounList"] = EndPoint { Func: sgd.queryNounList, HasArgs: false }
	}
}

func (sgd *SunGrowData) SetOutput(t string) {
	sgd.outputType.Set(t)
}

func (sgd *SunGrowData) SetOutputType(t output.OutputType) {
	sgd.outputType = t
}

func (sgd *SunGrowData) SetSaveAsFile(yes bool) {
	sgd.saveAsFile = yes
}

func (sgd *SunGrowData) SetEndpoints(endpoints ...string) {
	sgd.endPoints = endpoints
}

func (sgd *SunGrowData) SetArgs(args ...string) {
	sgd.request.SetArgs(args...)
}

func (sgd *SunGrowData) SetPsIds(psids ...string) {
	for range Only.Once {
		if len(psids) == 0 {
			var pids valueTypes.PsIds
			pids, sgd.Error = sgd.sunGrow.GetPsIds()
			if sgd.Error != nil {
				break
			}
			sgd.request.SetPSIDs(pids.Strings())
			break
		}

		sgd.request.SetPSIDs(psids)
	}
}

func (sgd *SunGrowData) GetData(args ...string) error {
	for range Only.Once {
		if len(sgd.endPoints) == 0 {
			sgd.Error = errors.New("need an endpoint")
			break
		}

		for _, endpoint := range sgd.endPoints {
			// Lookup endpoint interface from string.
			ep := sgd.sunGrow.GetEndpoint(endpoint)
			if sgd.sunGrow.IsError() {
				sgd.Error = sgd.sunGrow.Error
				break
			}
			sgd.request.SetRequired(ep.GetRequestArgNames())

			sgd.SetArgs(args...)

			// PsId not required.
			if sgd.request.IsPsIdNotRequired() {
				var result SunGrowDataResult

				result.EndPointName = ep.GetName()
				result.EndPoint = ep
				result.Request = sgd.request
				result.Response = sgd.CallEndpoint(ep, result.Request)
				sgd.results[result.EndPointName.String()] = result
				break
			}

			// PsId required and not set.
			if len(sgd.request.aPsId) == 0 {
				sgd.SetPsIds()
				if sgd.Error != nil {
					break
				}
			}

			// PsId required.
			for _, psId := range sgd.request.aPsId {
				var result SunGrowDataResult

				result.Request = sgd.request
				result.Request.SetPsId(psId.String())

				result.EndPointName = ep.GetName()
				result.EndPoint = ep
				result.Response = sgd.CallEndpoint(ep, result.Request)
				if sgd.Error != nil {
					break
				}
				sgd.results[result.EndPointName.String() + "/" + psId.String()] = result
			}

			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) GetOutput() error {
	for range Only.Once {
		if len(sgd.results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.results {
			result.Response.Data.ProcessMap()
			if sgd.Error != nil {
				break
			}

			args := result.Request.GetArgs(result.EndPoint)
			result.Response.Filename = result.Request.RequestAsFilePrefix()
			result.Response.Title = result.EndPoint.GetArea().String() + "." + result.EndPoint.GetName().String() + " - " + args
			sgd.Error = result.Response.GetOutput(sgd.sunGrow.OutputType, sgd.sunGrow.SaveAsFile)
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) GetResults() SunGrowDataResults {
	return sgd.results
}

func (sgd *SunGrowData) CallEndpoint(endpoint api.EndPoint, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	for range Only.Once {
		if !request.Validate(endpoint) {
			request.Help(endpoint)
			sgd.Error = errors.New("missing argument")
			break
		}

		var req []byte
		req, sgd.Error = json.Marshal(request)
		if sgd.Error != nil {
			fmt.Printf("GetEndpoint - ERR: %s\n", sgd.Error)
			break
		}
		// fmt.Printf("%s\n", req)

		if string(req) != "" {
			endpoint = endpoint.SetRequestByJson(output.Json(req))
			sgd.Error = endpoint.GetError()
			if sgd.Error != nil {
				fmt.Println(endpoint.Help())
				break
			}
		}
		// fmt.Printf("%s\n", endpoint.GetRequestJson())

		endpoint = endpoint.Call()
		sgd.Error = endpoint.GetError()
		if sgd.Error != nil {
			if strings.Contains(sgd.Error.Error(), "er_token_login_invalid") {
				sgd.sunGrow.Logout()
				break
			}
			fmt.Println(endpoint.Help())
			break
		}

		response.Data = endpoint.GetEndPointData()
		response.Name = endpoint.GetName().String()
		// response.Data.Table.AppendFilePrefix(request.RequestAsFilePrefix())		// request.GetFilename(endpoint.GetName().String()))
		// response.Data.Table.SetSaveFile(sgd.saveAsFile)
		// response.Data.Table.OutputType = sgd.outputType
		// response.Title = response.Data.Table.GetTitle()
		// response.Filename = response.Data.Table.GetFilePrefix()
	}

	return response
}


type SunGrowDataResults map[string]SunGrowDataResult
type SunGrowDataResult struct {
	EndPointName api.EndPointName
	EndPoint     api.EndPoint
	Request      SunGrowDataRequest
	Response     SunGrowDataResponse

	Error        error
}

func (sgd *SunGrowDataResult) ProcessMap() error {
	sgd.Response.Data.ProcessMap()
	sgd.Error = sgd.Response.Data.Error
	return sgd.Error
}

func (sgd *SunGrowDataResult) ProcessMapForMqtt() error {
	sgd.Response.Data.ProcessMapForMqtt()
	sgd.Error = sgd.Response.Data.Error
	return sgd.Error
}

func (sgd *SunGrowDataResult) GetOutput(outputType output.OutputType, saveAsFile bool, filePrefix string) error {
	sgd.Response.Filename = filePrefix
	sgd.Error = sgd.Response.GetOutput(outputType, saveAsFile)
	return sgd.Error
}

func (sgd *SunGrowDataResult) ResultTable(full bool) output.Table {
	ret := sgd.Response.Data.CreateResultTable(full)
	sgd.Error = sgd.Response.Data.Error
	return ret
}

func (sgd *SunGrowDataResult) DataTables() output.Tables {
	ret := sgd.Response.Data.CreateDataTables()
	sgd.Error = sgd.Response.Data.Error
	return ret
}

func (sgd *SunGrowDataResult) Sort() []string {
	return sgd.Response.Data.Sort()
}

func (sgd *SunGrowDataResult) Print() {
	sgd.Response.Data.Print()
}
