package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/findPsType"
	"GoSungrow/iSolarCloud/AppService/getAllDeviceByPsId"
	"GoSungrow/iSolarCloud/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/AppService/getDeviceModelInfoList"
	"GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getTemplateList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)


func (sg *SunGrow) GetPointNamesFromTemplate(template string) api.TemplatePoints {
	var ret api.TemplatePoints

	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		ep := sg.GetByStruct(
			"WebAppService.queryUserCurveTemplateData",
			queryUserCurveTemplateData.RequestData{TemplateID: template},
			time.Hour,
		)
		if sg.Error != nil {
			break
		}

		data := queryUserCurveTemplateData.AssertResultData(ep)
		for dn, dr := range data.PointsData.Devices {
			for _, pr := range dr.Points {
				if pr.Unit == "null" {
					pr.Unit = ""
				}
				ret = append(ret, api.TemplatePoint {
					PsKey:       dn,
					PointId:     "p" + pr.PointID,
					Description: pr.PointName,
					Unit:        pr.Unit,
				})
			}
		}
	}

	return ret
}

func (sg *SunGrow) GetTemplateData(template string, date string, filter string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)

		var psIds []api.Integer
		psIds, sg.Error = sg.StringToPids()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			pointNames := sg.GetPointNamesFromTemplate(template)
			if sg.Error != nil {
				break
			}

			ep := sg.GetByStruct(
				"AppService.queryMutiPointDataList",
				queryMutiPointDataList.RequestData {
					// PsID:           api.SetIntegerValue(psId),
					PsID:           psId,
					PsKey:          api.SetPsKeyValue(pointNames.PrintKeys()),
					Points:         pointNames.PrintPoints(),
					MinuteInterval: "5",
					StartTimeStamp: when.GetDayStartTimestamp(),
					EndTimeStamp:   when.GetDayEndTimestamp(),
				},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			// data := queryMutiPointDataList.AssertResultData(ep)
			data := queryMutiPointDataList.Assert(ep)
			table := data.GetPointDataTable(pointNames)
			if table.Error != nil {
				sg.Error = table.Error
				break
			}

			table.SetTitle("Template %s on %s for %d", template, when.String(), psId)
			table.SetFilePrefix(data.SetFilenamePrefix("%s-%s-%d", when.String(), template, psId))
			table.SetGraphFilter(filter)
			table.SetSaveFile(sg.SaveAsFile)
			table.OutputType = sg.OutputType
			sg.Error = table.Output()
			if sg.Error != nil {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetTemplatePoints(template string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		table := output.NewTable()
		sg.Error = table.SetHeader(
			"PointStruct Id",
			"Description",
			"Unit",
			)
		if sg.Error != nil {
			break
		}

		ss := sg.GetPointNamesFromTemplate(template)
		for _, s := range ss {
			sg.Error = table.AddRow(
				api.NameDevicePoint(s.PsKey, s.PointId),
				s.Description,
				s.Unit,
			)
			if sg.Error != nil {
				break
			}
		}
		if sg.Error != nil {
			break
		}

		table.SetTitle("Template %s", template)
		table.SetFilePrefix(template)
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) AllCritical() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.GetByJson("AppService.powerDevicePointList", "")
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep = sg.GetByJson("AppService.getPsList", "")
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)

		for _, psId := range _getPsList.GetPsIds() {
			ep = sg.GetByJson("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("WebAppService.showPSView", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			// ep = sg.GetByJson("AppService.findPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if ep.IsError() {
			// sg.Error = ep.GetError()
			// break
			// }

			ep = sg.GetByJson("AppService.getPowerStatistics", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.getPsDetail", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.getPsDetailWithPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.getPsHealthState", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.getPsListStaticData", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			ep = sg.GetByJson("AppService.getPsWeatherList", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			// ep = sg.GetByJson("AppService.queryAllPsIdAndName", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if ep.IsError() {
			// sg.Error = ep.GetError()
			// break
			// }

			// ep = sg.GetByJson("AppService.queryDeviceListByUserId", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if ep.IsError() {
			// sg.Error = ep.GetError()
			// break
			// }

			ep = sg.GetByJson("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}
		}
	}

	sg.Error = ep.GetError()
	return sg.Error
}

func (sg *SunGrow) PrintCurrentStats() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.Assert(ep)
		table := _getPsList.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("getPsList")
		table.SetFilePrefix(_getPsList.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}

		for _, psId := range _getPsList.GetPsIds() {
			ep = sg.GetByStruct(
				"AppService.queryDeviceList",
				// queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
				queryDeviceList.RequestData{PsId: psId},
				time.Second*60,
			)
			if sg.Error != nil {
				break
			}

			data := queryDeviceList.Assert(ep)

			table = data.GetDataTable()
			if table.Error != nil {
				sg.Error = table.Error
				break
			}

			table.SetTitle("queryDeviceList %s", psId)
			table.SetFilePrefix(data.SetFilenamePrefix("%s", psId))
			table.SetGraphFilter("")
			table.SetSaveFile(sg.SaveAsFile)
			table.OutputType = sg.OutputType
			sg.Error = table.Output()
			if sg.Error != nil {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPointNames(devices ...string) error {
	for range Only.Once {
		if len(devices) == 0 {
			devices = getPowerDevicePointNames.DeviceTypes
		}
		for _, dt := range devices {
			ep := sg.GetByStruct(
				"AppService.getPowerDevicePointNames",
				getPowerDevicePointNames.RequestData{DeviceType: dt},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			data := getPowerDevicePointNames.Assert(ep)
			table := data.GetPointDataTable()
			if table.Error != nil {
				sg.Error = table.Error
				break
			}

			table.SetTitle("getPowerDevicePointNames %s", dt)
			table.SetFilePrefix(data.SetFilenamePrefix(""))
			table.SetGraphFilter("")
			table.SetSaveFile(sg.SaveAsFile)
			table.OutputType = sg.OutputType
			sg.Error = table.Output()
			if sg.Error != nil {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetTemplates() error {
	for range Only.Once {
		ep := sg.GetByStruct(
			"AppService.getTemplateList",
			getTemplateList.RequestData{},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		data := getTemplateList.Assert(ep)
		table := data.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("getTemplateList")
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetIsolarcloudMqtt(appKey string) error {
	for range Only.Once {
		if appKey == "" {
			appKey = sg.GetAppKey()
		}

		ep := sg.GetByStruct(
			"WebAppService.getMqttConfigInfoByAppkey",
			getMqttConfigInfoByAppkey.RequestData{AppKey: appKey},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		data := getMqttConfigInfoByAppkey.Assert(ep)
		table := data.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("MQTT info")
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetRealTimeData(psKey string) error {
	for range Only.Once {
		if psKey == "" {
			var psKeys []string
			psKeys, sg.Error = sg.GetPsKeys()
			if sg.Error != nil {
				break
			}
			fmt.Printf("psKeys: %v\n", psKeys)
			psKey = strings.Join(psKeys, ",")
		}

		ep := sg.GetByStruct(
			"AppService.queryDeviceRealTimeDataByPsKeys",
			queryDeviceRealTimeDataByPsKeys.RequestData{PsKeyList: psKey},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		data := queryDeviceRealTimeDataByPsKeys.Assert(ep)
		table := data.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("Real Time Data %s", psKey)
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) CmdDataPsDetail(psIds ...api.Integer) error {
	for range Only.Once {
		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}

		var data SunGrowData
		data.New(sg)

		for _, psId := range psIds {
			response := data.Get("getPsDetail", SunGrowDataRequest{PsId: psId})
			if response.Error != nil {
				break
			}
			sg.Error = response.Table.Output()
			if sg.Error != nil {
				break
			}

			response = data.Get("getPsDetailWithPsType", SunGrowDataRequest{PsId: psId})
			if response.Error != nil {
				break
			}
			sg.Error = response.Table.Output()
			if sg.Error != nil {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPointData(date string, pointNames api.TemplatePoints, psIds ...api.Integer) error {
	for range Only.Once {
		if len(pointNames) == 0 {
			sg.Error = errors.New("no points defined")
			break
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)

		if len(psIds) == 0 {
			psIds, sg.Error = sg.StringToPids()
			if sg.Error != nil {
				break
			}
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(
				"AppService.queryMutiPointDataList",
				queryMutiPointDataList.RequestData{
					PsID:           psId,
					PsKey:          api.SetPsKeyValue(pointNames.PrintKeys()),
					Points:         pointNames.PrintPoints(),
					MinuteInterval: "5",
					StartTimeStamp: when.GetDayStartTimestamp(),
					EndTimeStamp:   when.GetDayEndTimestamp(),
				},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			data := queryMutiPointDataList.Assert(ep)
			table := data.GetPointDataTable(pointNames)
			if table.Error != nil {
				sg.Error = table.Error
				break
			}

			table.SetTitle("Point Data %s", psId)
			table.SetFilePrefix(data.SetFilenamePrefix("%d", psId))
			table.SetGraphFilter("")
			table.SetSaveFile(sg.SaveAsFile)
			table.OutputType = sg.OutputType
			sg.Error = table.Output()
			if sg.Error != nil {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) SearchPointNames(pns ...string) error {
	for range Only.Once {
		table := output.NewTable()

		_ = table.SetHeader(
			"DeviceType",
			"Id",
			"Period",
			"Point Id",
			"Point Name",
			"Show Point Name",
			"Translation Id",
		)

		if len(pns) == 0 {
			fmt.Println("Searching up to id 1000 within getPowerDevicePointInfo")
			for pni := 0; pni < 1000; pni++ {
				PrintPause(pni, 20)

				ep := sg.GetByStruct(
					"AppService.getPowerDevicePointInfo",
					// getPowerDevicePointInfo.RequestData{Id: strconv.FormatInt(int64(pni), 10)},
					getPowerDevicePointInfo.RequestData{Id: api.SetIntegerValue(int64(pni)) },
					DefaultCacheTimeout,
				)
				if sg.Error != nil {
					break
				}

				data := getPowerDevicePointInfo.Assert(ep)
				table = data.AddDataTable(table)
				if table.Error != nil {
					sg.Error = table.Error
					break
				}
			}
			fmt.Println("")
		} else {
			fmt.Printf("Searching for %v within getPowerDevicePointInfo\n", pns)
			for _, pn := range pns {
				ep := sg.GetByStruct(
					"AppService.getPowerDevicePointInfo",
					getPowerDevicePointInfo.RequestData{Id: api.SetIntegerString(pn)},
					DefaultCacheTimeout,
				)
				if sg.Error != nil {
					break
				}

				data := getPowerDevicePointInfo.Assert(ep)
				table = data.GetPointDataTable()
				if table.Error != nil {
					sg.Error = table.Error
					break
				}
			}
			fmt.Println("")
		}

		table.SetTitle("Point Names")
		table.SetFilePrefix("AppService.getPowerDevicePointInfo")
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetDeviceList(psIds ...api.Integer) error {
	for range Only.Once {
		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}

		var ret getDeviceList.Devices
		for _, psId := range psIds {
			ep := sg.GetByStruct(
				"AppService.getDeviceList",
				// getDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getDeviceList.RequestData{PsId: psId},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			data := getDeviceList.Assert(ep)
			ret = append(ret, data.GetDevices()...)
		}

		table := getDeviceList.GetDevicesTable(ret)
		table.SetTitle("All Devices")
		table.SetFilePrefix("")
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetDeviceModelInfoList() error {
	for range Only.Once {
		ep := sg.GetByStruct(
			"AppService.getDeviceModelInfoList",
			getDeviceModelInfoList.RequestData{},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		data := getDeviceModelInfoList.Assert(ep)
		table := data.GetPointDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("Models")
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetDevicePoints(psIds ...api.Integer) error {
	for range Only.Once {
		var points api.DataMap

		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}

		// getPsList
		ep := sg.GetByStruct("AppService.getPsList", getPsList.RequestData{}, DefaultCacheTimeout)
		if sg.Error != nil {
			break
		}
		PsList := getPsList.Assert(ep)
		data := PsList.GetData()
		if PsList.Error != nil {
			sg.Error = PsList.Error
			break
		}
		points.AppendMap(data)

		for _, psId := range psIds {
			// pid := strconv.Itoa(int(psId))

			// GoSungrow api raw getIncomeSettingInfos '{"ps_id":"1171348"}'
			// // getPowerStationForHousehold
			// // GoSungrow api raw getPowerStationForHousehold '{"ps_id":"1171348"}'

			// // getAreaList
			// getAreaList.Disabled
			//
			// // getPsReport
			// getPsReport.Disabled

			// // getPowerStatistics
			// // api raw getPowerStatistics '{"ps_id":"1171348"}'
			// ep = sg.GetByStruct("AppService.getPowerStatistics", getPowerStatistics.RequestData{ PsId: pid }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// PowerStatistics := getPowerStatistics.Assert(ep)
			// data = PowerStatistics.GetData()
			// if PowerStatistics.Error != nil {
			// 	sg.Error = PowerStatistics.Error
			// 	break
			// }
			// points.AppendMap(data)

			// getPowerStationData
			// api raw getPowerStationData '{"date_id":"20221007","date_type":"1","ps_id":"1171348"}'
			ep = sg.GetByStruct("AppService.getPowerStationData", getPowerStationData.RequestData{ PsId: psId, DateType: "1", DateID: "20221007"}, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			PowerStationData := getPowerStationData.Assert(ep)
			data = PowerStationData.GetData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)
			// api raw getPowerStationData '{"date_id":"202210","date_type":"2","ps_id":"1171348"}'
			ep = sg.GetByStruct("AppService.getPowerStationData", getPowerStationData.RequestData{ PsId: psId, DateType: "2", DateID: "202210"}, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			PowerStationData = getPowerStationData.Assert(ep)
			data = PowerStationData.GetData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)
			// api raw getPowerStationData '{"date_id":"2022","date_type":"3","ps_id":"1171348"}'
			ep = sg.GetByStruct("AppService.getPowerStationData", getPowerStationData.RequestData{ PsId: psId, DateType: "3", DateID: "2022"}, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			PowerStationData = getPowerStationData.Assert(ep)
			data = PowerStationData.GetData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)


			// queryDeviceList
			// api get AppService.queryDeviceList '{"ps_id":"1171348"}'
			ep = sg.GetByStruct("AppService.queryDeviceList", queryDeviceList.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			DeviceList := queryDeviceList.Assert(ep)
			data = DeviceList.GetData()
			if DeviceList.Error != nil {
				sg.Error = DeviceList.Error
				break
			}
			points.AppendMap(data)


			// getHouseholdStoragePsReport
			// api get getHouseholdStoragePsReport '{"date_id":"20221001","date_type":"1","ps_id":"1129147"}'
			ep = sg.GetByStruct(
				"AppService.getHouseholdStoragePsReport",
				getHouseholdStoragePsReport.RequestData{ DateID: "20221001", DateType: "1", PsId: psId },
				DefaultCacheTimeout,
				)
			if sg.Error != nil {
				break
			}
			HouseholdStoragePsReport := getHouseholdStoragePsReport.Assert(ep)
			data = HouseholdStoragePsReport.GetData()
			if HouseholdStoragePsReport.Error != nil {
				sg.Error = HouseholdStoragePsReport.Error
				break
			}
			points.AppendMap(data)


			// getPsDetailWithPsType
			ep = sg.GetByStruct("AppService.getPsDetailWithPsType", getPsDetailWithPsType.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			PsDetailWithPsType := getPsDetailWithPsType.Assert(ep)
			data = PsDetailWithPsType.GetData()
			if PsDetailWithPsType.Error != nil {
				sg.Error = PsDetailWithPsType.Error
				break
			}
			points.AppendMap(data)


			// getPsDetail
			ep = sg.GetByStruct("AppService.getPsDetail", getPsDetail.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			PsDetail := getPsDetail.Assert(ep)
			data = PsDetail.GetData()
			if PsDetail.Error != nil {
				sg.Error = PsDetailWithPsType.Error
				break
			}
			points.AppendMap(data)


			// getKpiInfo
			ep = sg.GetByStruct("AppService.getKpiInfo", getKpiInfo.RequestData{ }, DefaultCacheTimeout)
			if sg.Error != nil {
				break
			}
			KpiInfo := getKpiInfo.Assert(ep)
			data = KpiInfo.GetData()
			if KpiInfo.Error != nil {
				sg.Error = KpiInfo.Error
				break
			}
			points.AppendMap(data)


			// // getPowerDevicePointInfo
			// ep = sg.GetByStruct("AppService.getPowerDevicePointInfo", getPowerDevicePointInfo.RequestData{ Id: psId }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// ep4 := getPowerDevicePointInfo.Assert(ep)
			// data = ep4.GetData()
			// if ep4.Error != nil {
			// 	sg.Error = ep4.Error
			// 	break
			// }
			// points.AppendMap(data)


			// // queryDeviceRealTimeDataByPsKeys
			// ep = sg.GetByStruct("AppService.queryDeviceRealTimeDataByPsKeys", queryDeviceRealTimeDataByPsKeys.RequestData{ PsKeyList: psId }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// DeviceRealTimeDataByPsKeys := queryDeviceRealTimeDataByPsKeys.Assert(ep)
			// data = DeviceRealTimeDataByPsKeys.GetData()
			// if DeviceRealTimeDataByPsKeys.Error != nil {
			// 	sg.Error = DeviceRealTimeDataByPsKeys.Error
			// 	break
			// }
			// points.AppendMap(data)

		}

		points.Print()
	}

	return sg.Error
}


func PrintPause(index int, max int) {
	for range Only.Once {
		if index == 0 {
			fmt.Printf("\n%.3d ", index)
			break
		}

		m := math.Mod(float64(index), float64(max))
		if m == 0 {
			fmt.Printf("PAUSE")
			time.Sleep(time.Millisecond * 500)
			// fmt.Printf("\r%s%.3d ", strings.Repeat(" ", 4), pni)
			fmt.Printf("\r%.3d ", index)
		} else {
			time.Sleep(time.Millisecond * 100)
			fmt.Printf(".")
		}
	}
}

func (sg *SunGrow) GetPointName(psIds api.Integer) error {
	for range Only.Once {
		ep := sg.GetByStruct(
			"AppService.getPowerDevicePointInfo",
			getPowerDevicePointInfo.RequestData{Id: psIds},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		data := getPowerDevicePointInfo.Assert(ep)
		table := data.GetPointDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		// table2 := data.GetData()
		// fmt.Printf("%v\n", table2)

		table.SetTitle("Point Name %s", psIds)
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}


func (sg *SunGrow) GetPsIds() ([]api.Integer, error) {
	var ret []api.Integer

	for range Only.Once {
		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsIds()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsId() (api.Integer, error) {
	var ret api.Integer

	for range Only.Once {

		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		psIds := _getPsList.GetPsIds()
		if len(psIds) == 0 {
			break
		}

		ret = psIds[0]
	}

	return ret, sg.Error
}

func (sg *SunGrow) StringToPids(pids ...string) ([]api.Integer, error) {
	var psIds []api.Integer
	for range Only.Once {
		for _, pid := range pids {
			// p, err := strconv.ParseInt(pid, 10, 64)
			// if err != nil {
			// 	continue
			// }
			psIds = append(psIds, api.SetIntegerString(pid))
		}
		if len(psIds) == 0 {
			psIds, sg.Error = sg.GetPsIds()
			if sg.Error != nil {
				break
			}
		}
	}

	return psIds, sg.Error
}

func (sg *SunGrow) GetPsNames() ([]string, error) {
	var ret []string

	for range Only.Once {
		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsName()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetDevices(print bool) (getDeviceList.Devices, error) {
	var ret getDeviceList.Devices

	for range Only.Once {
		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		psIds := _getPsList.GetPsDevices()

		for _, psId := range psIds {
			ret = append(ret, getDeviceList.Device{
				Vendor:        api.SetStringValue(""),
				PsId:          psId.PsID,
				PsKey:         api.SetPsKeyValue(psId.PsID.String()),
				DeviceName:    psId.PsName,
				DeviceProSn:   psId.PsShortName,
				DeviceModel:   api.SetStringValue(""),
				DeviceType:    psId.PsType,
				DeviceCode:    api.SetIntegerValue(0),
				ChannelId:     api.SetIntegerValue(0),
				DeviceModelID: api.SetIntegerValue(0),
				TypeName:      api.SetStringValue("Ps Id"),
				DeviceState:   psId.PsHealthStatus,
				DevStatus:     psId.PsStatus.String(),
				Uuid:          api.SetIntegerValue(0),

				// 			PsFaultStatus:  d.PsFaultStatus,
				//			PsHealthStatus: d.PsHealthStatus,
				//			PsHolder:       d.PsHolder,
				//			PsID:           d.PsID,
				//			PsName:         d.PsName,
				//			PsShortName:    d.PsShortName,
				//			PsStatus:       d.PsStatus,
				//			PsType:         d.PsType,
			})

			ep = sg.GetByStruct(
				"AppService.getDeviceList",
				// getDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getDeviceList.RequestData{PsId: psId.PsID},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			data := getDeviceList.Assert(ep)
			ret = append(ret, data.GetDevices()...)
		}

		if !print {
			break
		}

		table := getDeviceList.GetDevicesTable(ret)
		table.SetTitle("All Devices")
		table.SetFilePrefix("")
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.Error != nil {
			break
		}
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsModels() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psIds []api.Integer
		psIds, sg.Error = sg.StringToPids()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(
				"AppService.getPsDetailWithPsType",
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			data := getPsDetailWithPsType.Assert(ep)
			ret = append(ret, data.GetDeviceName())
		}
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsSerials() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psIds []api.Integer
		psIds, sg.Error = sg.StringToPids()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(
				"AppService.getPsDetailWithPsType",
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			data := getPsDetailWithPsType.Assert(ep)
			ret = append(ret, data.GetDeviceSerial())
		}
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsKeys() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psIds []api.Integer
		psIds, sg.Error = sg.StringToPids()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(
				"AppService.getPsDetailWithPsType",
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if ep.IsError() {
				sg.Error = ep.GetError()
				break
			}

			data := getPsDetailWithPsType.Assert(ep)
			ret = append(ret, data.GetPsKeys()...)
		}
	}

	return ret, sg.Error
}


// ****************************************************** //

func (sg *SunGrow) GetEndpoints(endpoints []string, psIds []api.Integer, date api.DateTime) error {
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
			date = api.NewDateTime("now")
		}
		// fmt.Printf("FilePrefix: %s\n", date.Original())
		// fmt.Printf("String: %s\n", date.String())


		for _, endpoint := range endpoints {
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
				response := data.Get(endpoint, SunGrowDataRequest{ PsId: psId, Date: date })
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
	PsId       api.Integer  `json:"ps_id"`
	ReportType string       `json:"report_type"`
	Date       api.DateTime `json:"date"`
}
type SunGrowDataResponse struct {
	Data     api.DataMap
	Table    output.Table
	Filename string
	Title    string
	Error    error
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
	}
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
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
		response.Data = data.GetData()
		response.Table = data.GetDataTable()
	}
	return response
}

// func (sg *SunGrowData) getPsWeatherList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getRemoteUpgradeTaskList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.getRemoteUpgradeTaskList",
// 			getRemoteUpgradeTaskList.RequestData{ PsId: request.PsId },
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) getReportData(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) psForcastInfo(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
//
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
//
// func (sg *SunGrowData) reportList(request SunGrowDataRequest) SunGrowDataResponse {
// 	var response SunGrowDataResponse
// 	for range Only.Once {
// 		ep := sg.SunGrow.GetByStruct(
// 			"AppService.reportList",
// 			reportList.RequestData{ PsId: request.PsId },
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
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
// 		response.Data = data.GetData()
// 		response.Table = data.GetDataTable()
// 	}
// 	return response
// }
