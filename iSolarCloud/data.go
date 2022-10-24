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

func (sg *SunGrow) GetEndpoints(endpoints []string, psIds []valueTypes.Integer, date valueTypes.DateTime, reportType string, faultTypeCode string) error {
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

		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}

		if date.IsZero() {
			date = valueTypes.NewDateTime(valueTypes.Now)
		}

		for _, endpoint := range endpoints {
			args := sg.RequestArgs(endpoint)
			fmt.Printf("args:%s\n", args)
			if len(args) == 0 {
				// We don't have any request args.
				req := SunGrowDataRequest {}
				response := data.GetByApi(endpoint, req)
				if response.Error != nil {
					sg.Error = response.Error
					break
				}
				if data.Error != nil {
					sg.Error = data.Error
					break
				}
				sg.Error = response.Table.Output()
				if sg.Error != nil {
					break
				}

				continue
			}

			if _, ok := args["PsId"]; !ok {
				// If we don't need a PsId
				req := SunGrowDataRequest {
					Date:          date,
					ReportType:    reportType,
					FaultTypeCode: faultTypeCode,
				}
				response := data.GetByApi(endpoint, req)
				if response.Error != nil {
					sg.Error = response.Error
					break
				}
				if data.Error != nil {
					sg.Error = data.Error
					break
				}
				sg.Error = response.Table.Output()
				if sg.Error != nil {
					break
				}

				continue
			}

			if _, ok := args["PsId"]; ok {
				var finalData api.DataMap
				finalFilename := endpoint + "_"

				for _, psId := range psIds {
					finalFilename += psId.String() + "-"

					req := SunGrowDataRequest {
						PsId:          psId,
						Date:          date,
						ReportType:    reportType,
						FaultTypeCode: faultTypeCode,
					}

					// if data.HasArgs(endpoint) {
					// }
					// _, ok := data.FuncExists(endpoint)
					// if ok {
					// 	response = data.GetByFunc(endpoint, req)
					// } else {
					// 	data.Error = nil
					// 	response = data.GetByApi(endpoint, req)
					// }
					response := data.GetByApi(endpoint, req)
					if response.Error != nil {
						sg.Error = response.Error
						break
					}
					if data.Error != nil {
						sg.Error = data.Error
						break
					}

					finalData.AppendMap(response.Data)
				}

				finalTable := finalData.CreateTable()
				// response.Table.SetTitle(response.Title)
				finalFilename = strings.TrimSuffix(finalFilename, "-")
				finalTable.SetFilePrefix(finalFilename)
				finalTable.SetGraphFilter("")
				finalTable.SetSaveFile(sg.SaveAsFile)
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
	PsId          valueTypes.Integer  `json:"ps_id,omitempty"`
	ReportType    string              `json:"report_type,omitempty"`
	Date          valueTypes.DateTime `json:"date,omitempty"`
	FaultTypeCode string              `json:"fault_type_code,omitempty"`
}
type SunGrowDataResponse struct {
	Data     api.DataMap
	Table    output.Table
	Filename string
	Title    string
	Error    error
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

func (sgd *SunGrowData) GetByApi(endpoint string, request SunGrowDataRequest) SunGrowDataResponse {
	var response SunGrowDataResponse
	response.Table = output.NewTable()
	for range Only.Once {
		req, err := json.Marshal(request)
		if err != nil {
			fmt.Printf("GetByApi - ERR: %s\n", err)
		}

		data := sgd.SunGrow.GetByJson(endpoint, string(req))
		if data.IsError() {
			sgd.Error = data.GetError()
			break
		}

		fn := ""
		if request.PsId.String() != "" {
			fn += "-" + request.PsId.String()
		}
		if !request.Date.IsZero() {
			fn += "-" + request.Date.Original()
		}
		response.Filename = data.SetFilenamePrefix(fn)
		if response.Title == "" {
			response.Title = fmt.Sprintf("Data Request %s", endpoint)
		}

		response.Data = data.GetEndPointData()
		response.Table = data.GetEndPointDataTable()
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
			findCodeValueList.RequestData{ CodeType: "1" },
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
