package iSolarCloud

// package printTable
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/olekukonko/tablewriter"
// 	"os"
// )
//
//
// // writeTableTabular outputs tabular data to STDOUT
// func writeTableTabular(data [][]string, cols ...string) {
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetHeader(cols)
// 	table.SetBorder(false)
// 	table.AppendBulk(data)
//
// 	table.Render()
// }
//
// // writeTableJSON outputs JSON data to STDOUT
// func writeTableJSON(data [][]string, cols ...string) {
// 	mappedData := make([]map[string]string, 0)
// 	for i := range data {
// 		rowData := make(map[string]string)
// 		for j := range data[i] {
// 			rowData[cols[j]] = data[i][j]
// 		}
// 		mappedData = append(mappedData, rowData)
// 	}
// 	jsonData, err := json.Marshal(mappedData)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))
// }
//
// // writeTable outputs JSON or tabular data to STDOUT
// func writeTable(outType OutputType, data [][]string, cols ...string) {
// 	writeTableTabular(data, cols...)
// }
//
// //func EntityFilename(entity string) string {
// //	return strings.TrimSuffix(entity, ".json") + ".json"
// //}

// getPowerStationInfo - Contains state decoding.

// func (sgd *SunGrowData) getPsList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPsList",
// 			getPsList.RequestData{ },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getPsList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) queryDeviceList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.queryDeviceList",
// 			queryDeviceList.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := queryDeviceList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryDeviceList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - api get queryDeviceInfo '{"device_sn":"B2281302388","uuid":1179879}'
// func (sgd *SunGrowData) queryDeviceInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.queryDeviceInfo",
// 			queryDeviceInfo.RequestData{ Uuid: "1179877", DeviceSn: "B2281302388" },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := queryDeviceInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryDeviceInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) queryDeviceListForApp(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.queryDeviceListForApp",
// 			queryDeviceListForApp.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := queryDeviceListForApp.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryDeviceListForApp-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPsDetailWithPsType(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPsDetailWithPsType",
// 			getPsDetailWithPsType.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getPsDetailWithPsType.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsDetailWithPsType-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPsDetail(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPsDetail",
// 			getPsDetail.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getPsDetail.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsDetail-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) findPsType(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.findPsType",
// 			findPsType.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := findPsType.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("findPsType-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO -
// func (sgd *SunGrowData) getAllDeviceByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getAllDeviceByPsId",
// 			getAllDeviceByPsId.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getAllDeviceByPsId.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getAllDeviceByPsId-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getDeviceList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getDeviceList",
// 			getDeviceList.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getDeviceList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getDeviceList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getIncomeSettingInfos(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getIncomeSettingInfos",
// 			getIncomeSettingInfos.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getIncomeSettingInfos.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getIncomeSettingInfos-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getKpiInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getKpiInfo",
// 			getKpiInfo.RequestData{ },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getKpiInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getKpiInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPowerChargeSettingInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerChargeSettingInfo",
// 			getPowerChargeSettingInfo.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getPowerChargeSettingInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerChargeSettingInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getHouseholdStoragePsReport(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		// fmt.Println(request.Date.Original())
// 		// {"date_id":"20221001","date_type":"1","ps_id":"1129147"}
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getHouseholdStoragePsReport",
// 			getHouseholdStoragePsReport.RequestData{ PsId: request.PsId, DateType: request.Date.DateType, DateID: request.Date.Original() },
// 			api.DefaultTimeout,
// 		)
// 		if ep.IsError() {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		data := getHouseholdStoragePsReport.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getHouseholdStoragePsReport-%d-%s", request.PsId, request.Date.Original())
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - Not working
// func (sgd *SunGrowData) getPowerStationBasicInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerStationBasicInfo",
// 			getPowerStationBasicInfo.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPowerStationBasicInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = data.Error
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerStationBasicInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - Add in the ability to increment values by 5 minutes.
// Return from this function is an array of 288 values.
// func (sgd *SunGrowData) getPowerStationData(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerStationData",
// 			getPowerStationData.RequestData{ PsId: request.PsId, DateType: request.Date.DateType, DateID: request.Date.Original() },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPowerStationData.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerStationData-%d-%s", request.PsId, request.Date.Original())
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPowerStationForHousehold(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerStationForHousehold",
// 			getPowerStationForHousehold.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPowerStationForHousehold.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerStationForHousehold-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - Figure out how to properly flatten some of these "two field" arrays.
// func (sgd *SunGrowData) getPowerStationInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerStationInfo",
// 			getPowerStationInfo.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPowerStationInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerStationInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPowerStatistics(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPowerStatistics",
// 			getPowerStatistics.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPowerStatistics.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPowerStatistics-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPsHealthState(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPsHealthState",
// 			getPsHealthState.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsHealthState.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsHealthState-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - Figure out how to properly flatten some of these "two field" arrays.
// func (sgd *SunGrowData) powerDevicePointList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.powerDevicePointList",
// 			powerDevicePointList.RequestData{ },
// 			api.DefaultTimeout,
// 		)
//
// 		data := powerDevicePointList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("powerDevicePointList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getPsWeatherList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getPsWeatherList",
// 			getPsWeatherList.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getPsWeatherList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getPsWeatherList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// @TODO - Not working
// func (sgd *SunGrowData) getRemoteUpgradeTaskList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getRemoteUpgradeTaskList",
// 			getRemoteUpgradeTaskList.RequestData{ PsIdList: "1171348,1121412"},	// PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getRemoteUpgradeTaskList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getRemoteUpgradeTaskList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) reportList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.reportList",
// 			reportList.RequestData{ PsId: request.PsId, ReportType: request.ReportType },
// 			api.DefaultTimeout,
// 		)
//
// 		data := reportList.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("reportList-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) getReportData(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.getReportData",
// 			getReportData.RequestData{ PsId: request.PsId, ReportType: request.ReportType },
// 			api.DefaultTimeout,
// 		)
//
// 		data := getReportData.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("getReportData-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) psForcastInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.psForcastInfo",
// 			psForcastInfo.RequestData{ PsId: request.PsId },
// 			api.DefaultTimeout,
// 		)
//
// 		data := psForcastInfo.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("psForcastInfo-%d", request.PsId)
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) queryPowerStationInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
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

// func (sgd *SunGrowData) getPsIdState(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"WebAppService.getPsIdState",
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

// func (sgd *SunGrowData) queryPsProfit(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.queryPsProfit",
// 			queryPsProfit.RequestData{ PsId: request.PsId, DateType: request.Date.DateType, DateID: request.Date.Original() },
// 			api.DefaultTimeout,
// 		)
//
// 		data := queryPsProfit.Assert(ep)
// 		if data.Error != nil {
// 			response.Error = ep.GetError()
// 			break
// 		}
//
// 		response.Filename = data.SetFilenamePrefix("queryPsProfit-%d-%s", request.PsId, request.Date.Original())
// 		response.Data = data.GetEndPointData()
// 		response.Table = data.GetEndPointDataTable()
// 	}
// 	return response
// }

// func (sgd *SunGrowData) queryAllPsIdAndName(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
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

// func (sgd *SunGrowData) queryPsIdList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
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

// @TODO - Not working
// func (sgd *SunGrowData) queryPsNameByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"AppService.queryPsNameByPsId",
// 			queryPsNameByPsId.RequestData{ PsId: request.PsId },
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

// func (sgd *SunGrowData) showPSView(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"WebAppService.showPSView",
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

// func (sgd *SunGrowData) getMaxDeviceIdByPsId(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.SunGrow.GetByStruct(
// 			"WebIscmAppService.getMaxDeviceIdByPsId",
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
// 		response.Data.Table.SetSaveFile(sgd.sunGrow.SaveAsFile)
// 		response.Data.Table.OutputType = sgd.sunGrow.OutputType
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
// 		ok = sgd.sunGrow.RequestRequiresArgs(endpoint)
// 	}
// 	return yes
// }
//
//
// func (sgd *SunGrowData) energyTrend(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sgd.sunGrow.GetByStruct(
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
// 		ep := sgd.sunGrow.GetByStruct(
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
// 		ep := sgd.sunGrow.GetByStruct(
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
// 		ep := sgd.sunGrow.GetByStruct(
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
// 		ep := sgd.sunGrow.GetByStruct(
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
// 		ep := sgd.sunGrow.GetByStruct(
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
