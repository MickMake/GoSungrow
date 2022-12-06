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


type SunGrowData struct {
	Args      []string
	endPoints []string
	Request   SunGrowDataRequest

	Results SunGrowDataResults

	sunGrow    *SunGrow
	outputType output.OutputType
	saveAsFile bool
	Error      error
}

func (sgd *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sgd.sunGrow = ref
		sgd.Results = make(SunGrowDataResults)

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
	sgd.Args = args
}

func (sgd *SunGrowData) SetPsIds(psids ...string) {
	for range Only.Once {
		// if len(psids) == 0 {
		// 	var pids valueTypes.PsIds
		// 	pids, sgd.Error = sgd.sunGrow.GetPsIds()
		// 	if sgd.Error != nil {
		// 		break
		// 	}
		// 	sgd.Request.SetPSIDs(pids.Strings())
		// 	break
		// }

		var pids valueTypes.PsIds
		pids = sgd.sunGrow.SetPsIds(psids...)
		if sgd.Error != nil {
			break
		}

		sgd.Request.SetPSIDs(pids.Strings())
	}
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
		args := request.GetArgs(response.Data.EndPoint)
		name := endpoint.GetArea().String() + "." + endpoint.GetName().String()

		response.Options = OutputOptions {
			Name:        endpoint.GetName().String(),
			OutputType:  sgd.sunGrow.OutputType,
			FileSuffix:  request.GetPrimaryArg(), // + " - " + request.RequestAsFilePrefix(),
			SaveAsFile:  sgd.sunGrow.SaveAsFile,
			TitleSuffix: args,
			GraphRequest: output.GraphRequest {
				Title:       name + " - " + args,
				TimeColumn:  nil,
				DataColumn:  nil,
				UnitsColumn: nil,
				NameColumn:  nil,
				// DataColumn:   nil,
				DataMin: nil,
				DataMax: nil,
				Width:   nil,
				Height:  nil,
				Error:   nil,
			},
		}
	}

	return response
}

func (sgd *SunGrowData) GetData() error {
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
			sgd.Request.SetRequired(ep.GetRequestArgNames())
			sgd.Request.SetArgs(sgd.Args...)

			// PsId not required.
			if sgd.Request.IsPsIdNotRequired() {
				var result SunGrowDataResult

				result.EndPointArea = ep.GetArea()
				result.EndPointName = ep.GetName()
				result.EndPoint = ep
				result.Request = sgd.Request
				result.Response = sgd.CallEndpoint(ep, result.Request)
				sgd.Results[result.EndPointName.String()] = result
				break
			}

			// PsId required and not set.
			if len(sgd.Request.aPsId) == 0 {
				sgd.SetPsIds()
				if sgd.Error != nil {
					break
				}
			}

			// PsId required.
			for _, psId := range sgd.Request.aPsId {
				var result SunGrowDataResult
				result.Request = sgd.Request
				result.Request.SetPsId(psId.String())

				result.EndPointArea = ep.GetArea()
				result.EndPointName = ep.GetName()
				result.EndPoint = ep
				result.Response = sgd.CallEndpoint(ep, result.Request)
				if sgd.Error != nil {
					break
				}
				sgd.Results[result.EndPointName.String() + "/" + psId.String()] = result
			}

			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) Process() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			result.Response.Data.ProcessMap()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) Output() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			sgd.Error = result.Response.Output()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowData) OutputDataTables() error {
	for range Only.Once {
		if len(sgd.Results) == 0 {
			fmt.Println("No results found.")
			break
		}

		for _, result := range sgd.Results {
			sgd.Error = result.Response.OutputDataTables()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}


type SunGrowDataResults map[string]SunGrowDataResult
type SunGrowDataResult struct {
	EndPointArea api.AreaName
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

func (sgd *SunGrowDataResult) CreateResultTable(full bool) output.Table {
	ret := sgd.Response.CreateResultTable(full)
	sgd.Error = sgd.Response.Data.Error
	return ret
}

func (sgd *SunGrowDataResult) CreateDataTables() api.Tables {
	tables := sgd.Response.CreateDataTables()
	sgd.Error = sgd.Response.Data.Error
	return tables
}

func (sgd *SunGrowDataResult) Sort() []string {
	return sgd.Response.Data.Sort()
}

func (sgd *SunGrowDataResult) Print() {
	sgd.Response.Data.Print()
}


type OutputOptions struct {
	Name        string
	TitleSuffix string
	OutputType output.OutputType
	FileSuffix string
	SaveAsFile bool
	GraphRequest output.GraphRequest

	// table.InitGraph(output.GraphRequest {
	// 	Title:        "",
	// 	TimeColumn:   output.SetString("Date/Time"),
	// 	SearchColumn: output.SetString("Point Id"),
	// 	NameColumn:   output.SetString("Point Name"),
	// 	ValueColumn:  output.SetString("Value"),
	// 	UnitsColumn:  output.SetString("Units"),
	// 	SearchString: output.SetString(""),
	// 	MinLeftAxis:  output.SetFloat(0),
	// 	MaxLeftAxis:  output.SetFloat(0),
	// })
}

type SunGrowDataResponses map[string]SunGrowDataResponse
type SunGrowDataFunction func(request SunGrowDataRequest) SunGrowDataResponse
type SunGrowDataResponse struct {
	Data     api.DataMap
	Options  OutputOptions
	Error    error
}

func (sgd *SunGrowDataResponse) CreateResultTable(full bool) output.Table {
	ret := sgd.Data.CreateResultTable(full)
	sgd.Error = sgd.Data.Error
	return ret
}

func (sgd *SunGrowDataResponse) CreateDataTables() api.Tables {
	tables := sgd.Data.CreateDataTables()
	sgd.Error = sgd.Data.Error
	return tables
}

func (sgd *SunGrowDataResponse) Output() error {
	for range Only.Once {
		// Outputs that don't drop through.
		if sgd.Options.OutputType.IsStruct() || sgd.Options.OutputType.IsList() || sgd.Options.OutputType.IsRaw() || sgd.Options.OutputType.IsJson() {
			table := sgd.Data.CreateResultTable(true)
			table.OutputType = sgd.Options.OutputType
			table.SetSaveFile(sgd.Options.SaveAsFile)
			table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			table.AppendFilePrefix(sgd.Options.FileSuffix)
			sgd.Error = table.Output()
			break
		}

		// Outputs that can drop through to DataTables.
		if sgd.Options.OutputType.IsTable() || sgd.Options.OutputType.IsXLSX() || sgd.Options.OutputType.IsCsv() || sgd.Options.OutputType.IsXML() {
			table := sgd.Data.CreateResultTable(false)
			table.OutputType = sgd.Options.OutputType
			table.SetSaveFile(sgd.Options.SaveAsFile)
			table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			table.AppendFilePrefix(sgd.Options.FileSuffix)
			sgd.Error = table.Output()
			if sgd.Error != nil {
				break
			}
			// break
		}

		sgd.Error = sgd.OutputDataTables()
	}

	return sgd.Error
}

func (sgd *SunGrowDataResponse) OutputDataTables() error {
	for range Only.Once {
		tables := sgd.Data.CreateDataTables()
		if len(tables) == 0 {
			break
		}

		// @iSolarCloud/api/struct_data.go:420
		if sgd.Options.OutputType.IsGraph() {
			for _, data := range tables {
				if !data.IsValid {
					fmt.Printf("# %s.%s - has no data.\n", data.Area, data.Name)
					continue
				}
				// table := tables[name]
				// table.OutputType = sgd.Options.OutputType
				// table.SetSaveFile(sgd.Options.SaveAsFile)
				// table.SetFilePrefix("%s", sgd.Options.Filename)

				// value := values[name]
				// fmt.Printf("Writing (%d)\n", len(value))
				// fmt.Printf("%v\n", table.GetHeaders())

				if sgd.Options.TitleSuffix == "" {
					sgd.Options.TitleSuffix = data.Table.GetTitle()
				}
				data.Table.OutputType = sgd.Options.OutputType
				data.Table.SetSaveFile(true)
				data.Table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
				data.Table.AppendFilePrefix(sgd.Options.FileSuffix)

				if sgd.Options.GraphRequest.TimeColumn == nil {
					for _, col := range data.Table.GetHeaders() {
						val := data.Values.GetCell(0, col)
						if val.IsTypeDateTime() {
							sgd.Options.GraphRequest.TimeColumn = &col
							break
						}
					}
				}

				// if sgd.Options.GraphRequest.DataColumn == nil {
				// 	for _, col := range data.Table.GetHeaders() {
				// 		val := data.Values.GetCell(0, col)
				// 		if val.IsNumber() {
				// 			sgd.Options.GraphRequest.DataColumn = &col
				// 			break
				// 		}
				// 	}
				// }

				if sgd.Options.GraphRequest.DataColumn == nil {
					for _, col := range data.Table.GetHeaders() {
						val := data.Values.GetCell(0, col)
						if val.IsNumber() {
							sgd.Options.GraphRequest.DataColumn = &col
							break
						}
					}
				}

				if sgd.Options.GraphRequest.UnitsColumn != nil {
					for _, col := range data.Table.GetHeaders() {
						if *sgd.Options.GraphRequest.UnitsColumn != col {
							continue
						}
						val := data.Values.GetCell(0, col)
						unit := val.Unit()
						if unit != "" {
							continue
						}
						sgd.Options.GraphRequest.UnitsColumn = &col
						sgd.Options.GraphRequest.DataUnit = &unit
						break
					}
				}

				if sgd.Options.GraphRequest.NameColumn == nil {
				}

				// if sgd.Options.GraphRequest.DataMin == nil {
				// 	zero := 0.0
				// 	sgd.Options.GraphRequest.DataMin = &zero
				// }

				// if sgd.Options.GraphRequest.DataMax == nil {
				// 	zero := 0.0
				// 	sgd.Options.GraphRequest.DataMax = &zero
				// }

				// if sgd.Options.GraphRequest.Width == nil {
				// }

				// if sgd.Options.GraphRequest.Height == nil {
				// }

				sgd.Error = data.Table.SetGraph(sgd.Options.GraphRequest)
				if sgd.Error != nil {
					break
				}

				sgd.Error = data.Table.Output()
				if sgd.Error != nil {
					break
				}
			}
			break
		}

		for _, data := range tables {
			fmt.Println()
			data.Table.OutputType = sgd.Options.OutputType
			data.Table.SetSaveFile(sgd.Options.SaveAsFile)
			data.Table.AppendTitle(" - %s", sgd.Options.TitleSuffix)
			data.Table.AppendFilePrefix(sgd.Options.FileSuffix)

			sgd.Error = data.Table.Output()
			if sgd.Error != nil {
				break
			}
		}
	}

	return sgd.Error
}

func (sgd *SunGrowDataResponse) Print() {
	sgd.Data.Print()
}
