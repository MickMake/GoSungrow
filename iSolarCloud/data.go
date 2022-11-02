package iSolarCloud

import (
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

func (sg *SunGrow) GetEndpoints(endpoints string, args ...string) error {
	for range Only.Once {
		var data SunGrowData
		data.New(sg)

		var request SunGrowDataRequest
		request = request.Set(args...)

		eps := SplitArg(endpoints)
		// if endpoints == "all" {
		// 	eps = data.GetAllEndPoints()
		// }

		if endpoints == "" {
			// fmt.Println("Additional commands available, (on top of endpoints):")
			// for _, ep := range data.GetAllEndPoints() {
			// 	fmt.Printf("\t%s\n", ep)
			// }
			sg.Error = errors.New("need an endpoint")
			break
		}

		for _, endpoint := range eps {
			ep := sg.GetEndpoint(endpoint)
			sg.Error = ep.GetError()
			if sg.Error != nil {
				break
			}

			rargs := ep.GetRequestArgNames()
			// fmt.Printf("args:%s\n", rargs)

			if _, ok := rargs[NamePsId]; !ok {
				response := data.GetByApi(ep, request)
				if response.Error != nil {
					sg.Error = response.Error
					break
				}
				if data.Error != nil {
					sg.Error = data.Error
					break
				}

				if sg.OutputType.IsTable() {
					if len(response.Data.DataTables) == 0 {
						fmt.Println("Now data tables defined - try a get instead.")
						break
					}
					for _, t := range response.Data.DataTables {
						t.OutputType = sg.OutputType
						t.SetSaveFile(sg.SaveAsFile)
						sg.Error = t.Output()
						if sg.Error != nil {
							break
						}
					}
					break
				}

				response.Data.Table.SetFilePrefix(request.GetFilename(endpoint))
				sg.Error = response.Data.Table.Output()
				break
			}

			if len(request.aPsId) == 0 {
				var pids valueTypes.PsIds
				pids, sg.Error = sg.GetPsIds()
				if sg.Error != nil {
					break
				}
				request.SetPsIds(pids.Strings())
			}
			for _, psId := range request.aPsId {
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

				if sg.OutputType.IsTable() {
					for _, t := range response.Data.DataTables {
						t.OutputType = sg.OutputType
						t.SetSaveFile(sg.SaveAsFile)
						sg.Error = t.Output()
						if sg.Error != nil {
							break
						}
					}
					break
				}

				response.Data.Table.SetFilePrefix(request.GetFilename(endpoint))
				sg.Error = response.Data.Table.Output()
			}
			break
		}
	}

	return sg.Error
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
	Filename string
	Title    string
	Error    error
}


type SunGrowData struct {
	// EndPoints EndPoints
	SunGrow   *SunGrow
	Error     error
}

func (sgd *SunGrowData) New(ref *SunGrow) {
	for range Only.Once {
		sgd.SunGrow = ref
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

func (sgd *SunGrowData) GetByEndPointName(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	response.Data.Table = output.NewTable()
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
	// response.Table = output.NewTable()
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
			response.Title = fmt.Sprintf("Data Request %s", endpoint.GetName())
		}

		response.Data = endpoint.GetEndPointData()
		// endpoint.GetEndPointDataTables()
		// response.Table = endpoint.GetEndPointResultTable()
		// response.Table.SetTitle(response.Title)
		// response.Table.SetFilePrefix(response.Filename)
		// response.Data.Table.SetGraphFilter("")
		response.Data.Table.SetSaveFile(sgd.SunGrow.SaveAsFile)
		response.Data.Table.OutputType = sgd.SunGrow.OutputType
	}
	return response
}

// func (sgd *SunGrowData) GetAllEndPoints() []string {
// 	var ret []string
// 	for ep := range sgd.EndPoints {
// 		ret = append(ret, ep)
// 	}
// 	return ret
// }
//
// func (sgd *SunGrowData) GetByFunc(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		dataEndPoint, ok := sgd.FuncExists(endpoint)
// 		if !ok {
// 			break
// 		}
//
// 		response = dataEndPoint.Func(request)
// 		if response.Error != nil {
// 			break
// 		}
// 		if sgd.Error != nil {
// 			break
// 		}
//
// 		if response.Filename == "" {
// 			response.Filename = endpoint
// 		}
// 		if response.Title == "" {
// 			response.Title = fmt.Sprintf("Data Request %s", endpoint)
// 		}
// 		response.Data.Table.SetTitle(response.Title)
// 		response.Data.Table.SetFilePrefix(response.Filename)
// 		response.Data.Table.SetGraphFilter("")
// 		response.Data.Table.SetSaveFile(sgd.SunGrow.SaveAsFile)
// 		response.Data.Table.OutputType = sgd.SunGrow.OutputType
// 	}
// 	return response
// }
//
// func (sgd *SunGrowData) FuncExists(endpoint string) (EndPoint, bool) {
// 	var dataFunc EndPoint
// 	var yes bool
// 	for range Only.Once {
// 		if dataFunc, yes = sgd.EndPoints[endpoint]; yes {
// 			yes = true
// 			break
// 		}
// 		sgd.Error = errors.New(fmt.Sprintf("unknown endpoint function '%s'", endpoint))
// 	}
// 	return dataFunc, yes
// }
//
// func (sgd *SunGrowData) HasArgs(endpoint string) bool {
// 	var yes bool
// 	for range Only.Once {
// 		dataEndPoint, ok := sgd.FuncExists(endpoint)
// 		if ok {
// 			yes = dataEndPoint.HasArgs
// 			break
// 		}
//
// 		ok = sgd.SunGrow.RequestRequiresArgs(endpoint)
// 	}
// 	return yes
// }
//
//
// func (sgd *SunGrowData) energyTrend(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.energyTrend",
// 			// energyTrend.RequestData{ PsId: request.PsId },
// 			energyTrend.RequestData{  },
// 			api.DefaultTimeout,
// 		)
//
// 		data := energyTrend.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("energyTrend-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
// func (sgd *SunGrowData) getAreaList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getAreaList",
// 			// energyTrend.RequestData{ PsId: request.PsId },
// 			getAreaList.RequestData{  },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getAreaList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getAreaList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
// // @TODO - Need to add OrgIds
// // @TODO - Need to support []string in ResultData
// func (sgd *SunGrowData) getAllPsIdByOrgIds(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getAllPsIdByOrgIds",
// 			// getAllPsIdByOrgIds.RequestData{ PsId: request.PsId },
// 			getAllPsIdByOrgIds.RequestData{ OrgIds: valueTypes.SetStringValue("362245") },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getAllPsIdByOrgIds.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getAllPsIdByOrgIds-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
// // @TODO - No data.
// func (sgd *SunGrowData) findCodeValueList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.findCodeValueList",
// 			// findCodeValueList.RequestData{ PsId: request.PsId },
// 			findCodeValueList.RequestData{ CodeType: *request.CodeType },
// 			api.DefaultTimeout,
// 		)
//
// 		data := findCodeValueList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("findCodeValueList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
// // @TODO - No data.
// func (sgd *SunGrowData) queryFaultCodes(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"WebAppService.queryFaultCodes",
// 			// queryFaultCodes.RequestData{ PsId: request.PsId },
// 			queryFaultCodes.RequestData{ FaultName: valueTypes.SetStringValue("417") },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryFaultCodes.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryFaultCodes-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
// func (sgd *SunGrowData) queryNounList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"WebAppService.queryNounList",
// 			queryNounList.RequestData{ FaultTypeCode: valueTypes.SetStringValue("718") },
// 			// queryNounList.RequestData{ },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryNounList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryNounList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Data.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
//
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
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
// 		response.Table = data.GetEndPointResultTable()
// 	}
// 	return response
// }
