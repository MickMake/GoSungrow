package iSolarCloud


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
