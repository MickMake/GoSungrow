package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getTemplateList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"strconv"
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

		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		pointNames := sg.GetPointNamesFromTemplate(template)
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData {
				PsID:           psId,
				PsKey:          pointNames.PrintKeys(),
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
		table := data.GetDataTable(pointNames)
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		fn := data.SetFilenamePrefix("%s-%s", when.String(), template)
		sg.Error = table.SetFilePrefix(fn)
		if sg.Error != nil {
			break
		}

		sg.Error = sg.Output(ep, &table, filter)
		if sg.Error != nil {
			break
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
			"Point Id",
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

		table.Print()
	}

	return sg.Error
}

func (sg *SunGrow) AllCritical() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.GetByJson("AppService.powerDevicePointList", "")
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsList", "")
		if ep.IsError() {
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		psId := _getPsList.GetPsId()

		ep = sg.GetByJson("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("WebAppService.showPSView", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		// ep = sg.GetByJson("AppService.findPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		ep = sg.GetByJson("AppService.getPowerStatistics", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsDetail", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsDetailWithPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsHealthState", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsListStaticData", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.GetByJson("AppService.getPsWeatherList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		// ep = sg.GetByJson("AppService.queryAllPsIdAndName", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		// ep = sg.GetByJson("AppService.queryDeviceListByUserId", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		ep = sg.GetByJson("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
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
			break
		}
		_getPsList := getPsList.Assert(ep)
		psId := _getPsList.GetPsId()
		table := _getPsList.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(_getPsList, &table, "")
		if sg.Error != nil {
			break
		}

		ep = sg.GetByStruct(
			"AppService.queryDeviceList",
			queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
			time.Second * 60,
		)
		if sg.Error != nil {
			break
		}

		ep2 := queryDeviceList.Assert(ep)
		table = ep2.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) QueryDevice(psId int64) queryDeviceList.EndPoint {
	var ret queryDeviceList.EndPoint
	for range Only.Once {
		if psId == 0 {
			psId, sg.Error = sg.GetPsId()
			if sg.Error != nil {
				break
			}
		}

		// ep = sg.GetByJson("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		ep := sg.GetByStruct(
			"AppService.queryDeviceList",
			queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
			time.Second * 60,
		)
		// if sg.Error != nil {
		// 	break
		// }

		ret = queryDeviceList.Assert(ep)
	}

	return ret
}

func (sg *SunGrow) QueryPs(psId int64) getPsList.EndPoint {
	var ret getPsList.EndPoint
	for range Only.Once {
		if psId == 0 {
			psId, sg.Error = sg.GetPsId()
			if sg.Error != nil {
				break
			}
		}

		// ep = sg.GetByJson("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		ep := sg.GetByStruct(
			"AppService.getPsList",
			getPsList.RequestData{},
			time.Second * 60,
		)
		// if sg.Error != nil {
		// 	break
		// }

		ret = getPsList.Assert(ep)
	}

	return ret
}

func (sg *SunGrow) GetPointNames() error {
	for range Only.Once {
		for _, dt := range getPowerDevicePointNames.DeviceTypes {
			ep := sg.GetByStruct(
				"AppService.getPowerDevicePointNames",
				getPowerDevicePointNames.RequestData{DeviceType: dt},
				DefaultCacheTimeout,
			)
			if sg.Error != nil {
				break
			}

			ep2 := getPowerDevicePointNames.Assert(ep)
			table := ep2.GetDataTable()
			if table.Error != nil {
				sg.Error = table.Error
				break
			}

			sg.Error = sg.Output(ep2, &table, "")
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

		ep2 := getTemplateList.Assert(ep)
		table := ep2.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
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

		ep2 := getMqttConfigInfoByAppkey.Assert(ep)
		table := ep2.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
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
			fmt.Printf("%v\n", psKeys)
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

		ep2 := queryDeviceRealTimeDataByPsKeys.Assert(ep)
		table := ep2.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, nil, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPsDetails(psid string) error {
	for range Only.Once {
		var psId int64
		if psid == "" {
			psId, sg.Error = sg.GetPsId()
		} else {
			psId, sg.Error = strconv.ParseInt(psid, 10, 64)
		}
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep2 := getPsDetailWithPsType.Assert(ep)
		table := ep2.GetDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPointData(date string, pointNames api.TemplatePoints) error {
	for range Only.Once {
		if len(pointNames) == 0 {
			sg.Error = errors.New("no points defined")
			break
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)

		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData {
				PsID:           psId,
				PsKey:          pointNames.PrintKeys(),
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

		ep2 := queryMutiPointDataList.Assert(ep)
		table := ep2.GetDataTable(pointNames)
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPsId() (int64, error) {
	var ret int64

	for range Only.Once {

		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsId()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsName() (string, error) {
	var ret string

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

func (sg *SunGrow) GetPsModel() (string, error) {
	var ret string

	for range Only.Once {
		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep2 := getPsDetailWithPsType.Assert(ep)
		ret = ep2.GetDeviceName()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsSerial() (string, error) {
	var ret string

	for range Only.Once {
		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep2 := getPsDetailWithPsType.Assert(ep)
		ret = ep2.GetDeviceSerial()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsKeys() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep2 := getPsDetailWithPsType.Assert(ep)
		ret = ep2.GetPsKeys()
	}

	return ret, sg.Error
}
