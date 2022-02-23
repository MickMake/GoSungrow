package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/api"
	"errors"
	"fmt"
	"strings"
)

func (sg *SunGrow) Get(endpoint string, request string) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		if sg.Error != nil {
			break
		}
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		if request != "" {
			ret = ret.SetRequestByJson(api.Json(request))
			if ret.IsError() {
				fmt.Println(ret.Help())
				sg.Error = ret.GetError()
				break
			}
		}

		ret = ret.Call()
		if ret.IsError() {
			fmt.Println(ret.Help())
			sg.Error = ret.GetError()
			break
		}

		switch {
		case sg.OutputType.IsNone():

		case sg.OutputType.IsFile():
			sg.Error = ret.WriteFile()

		case sg.OutputType.IsRaw():
			fmt.Println(ret.GetData(true))

		case sg.OutputType.IsJson():
			fmt.Println(ret.GetData(false))

		default:
		}
	}
	return ret
}

func (sg *SunGrow) RefGet(endpoint string, request interface{}) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		if sg.Error != nil {
			break
		}
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		if request != nil {
			ret = ret.SetRequest(request)
			if ret.IsError() {
				sg.Error = ret.GetError()
				break
			}
		}

		ret = ret.Call()
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}
	}
	return ret
}

func (sg *SunGrow) GetHighLevel(name string, args ...string) error {
	for range Only.Once {
		name = strings.ToLower(name)
		switch name {
		case "stats":
			sg.Error = sg.GetCurrentStats()
		case "template":
			args = fillArray(2, args)
			sg.Error = sg.GetData(args[0], args[1])
		default:
			sg.Error = errors.New("unknown high-level command")
		}
	}
	return sg.Error
}

func (sg *SunGrow) ListHighLevel() {
	fmt.Println("stats - Get current inverter stats, (last 5 minutes).")
	fmt.Println("template [date] [template_id] - Get data from template.")
}

func (sg *SunGrow) AllCritical() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.Get("AppService.powerDevicePointList", "")
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsList", "")
		if ep.IsError() {
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		psId := _getPsList.GetPsId()

		ep = sg.Get("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("WebAppService.showPSView", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		// ep = sg.Get("AppService.findPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		ep = sg.Get("AppService.getPowerStatistics", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsDetail", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsDetailWithPsType", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsHealthState", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsListStaticData", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		ep = sg.Get("AppService.getPsWeatherList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		// ep = sg.Get("AppService.queryAllPsIdAndName", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		// ep = sg.Get("AppService.queryDeviceListByUserId", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }

		ep = sg.Get("AppService.queryDeviceListForApp", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

	}

	sg.Error = ep.GetError()
	return sg.Error
}

func (sg *SunGrow) GetCurrentStats() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.Get("AppService.getPsList", "")
		if ep.IsError() {
			break
		}
		_getPsList := getPsList.AssertResultData(ep)
		psId := _getPsList.GetPsId()

		if sg.OutputType.IsHuman() {
			_queryDeviceList := getPsList.AssertResultData(ep)
			points := _queryDeviceList.GetData()
			for _, r := range points {
				for _, c := range r {
					fmt.Printf("\"%s\",", c)
				}
				fmt.Println()
			}
		}

		// ep = sg.Get("AppService.getPsDetail", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		// if ep.IsError() {
		// 	break
		// }
		//
		// if sg.OutputType.IsHuman() {
		// 	_queryDeviceList := getPsDetail.AssertResultData(ep)
		// 	points := _queryDeviceList.GetDataByName("SH10RT")
		// 	for _, p := range points {
		// 		t := api.NewDateTime(p.TimeStamp)
		// 		fmt.Printf("\"%v\",\"%s\",\"%s\",\"%s\",\"%s\"\n",
		// 			t,
		// 			p.PointGroupName,
		// 			p.PointName,
		// 			p.Value,
		// 			p.Unit,
		// 		)
		// 	}
		// }

		ep = sg.Get("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		if ep.IsError() {
			break
		}

		if sg.OutputType.IsHuman() {
			_queryDeviceList := queryDeviceList.AssertResultData(ep)
			points := _queryDeviceList.GetCsvByName("SH10RT")
			fmt.Printf("%v", points)
		}
	}

	sg.Error = ep.GetError()
	return sg.Error
}

func (sg *SunGrow) GetData(date string, template string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		ep := sg.RefGet("WebAppService.queryUserCurveTemplateData", queryUserCurveTemplateData.RequestData{
			TemplateID: template,
		})

		data := queryUserCurveTemplateData.AssertResultData(ep)
		var pskeys string
		var points string
		pointNames := make(map[string]string)
		pointUnits := make(map[string]string)
		for dn, dr := range data.PointsData.Devices {
			for pn, pr := range dr.Points {
				pointNames["p"+pr.PointID] = pr.PointName
				pointUnits["p"+pr.PointID] = pr.Unit
				pskeys += "," + dn
				points += "," + pn
			}
		}
		pskeys = strings.TrimPrefix(pskeys, ",")
		points = strings.TrimPrefix(points, ",")

		//
		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)
		psId := sg.GetPsId()

		ep2 := sg.RefGet("AppService.queryMutiPointDataList", queryMutiPointDataList.RequestData{
			PsID:           psId,
			PsKey:          pskeys,
			Points:         points,
			MinuteInterval: "5",
			StartTimeStamp: when.GetDayStartTimestamp(),
			EndTimeStamp:   when.GetDayEndTimestamp(),
		})

		//
		csv := api.NewCsv()
		csv = csv.SetHeader([]string{
			"Date/Time",
			"Device Name",
			"Point Name",
			"Value",
			"Units",
		})

		data2 := queryMutiPointDataList.AssertResultData(ep2)
		for deviceName, deviceRef := range data2.Devices {
			for pointId, pointRef := range deviceRef.Points {
				for _, tim := range pointRef.Times {
					csv = csv.AddRow([]string{
						tim.Key.String(),
						deviceName,
						fmt.Sprintf("%s (%s)", pointNames[pointId], pointId),
						tim.Value,
						pointUnits[pointId],
					})
				}
			}
		}

		fmt.Println(csv.String())
	}

	return sg.Error
}

func (sg *SunGrow) GetPsId() int64 {
	var ret int64

	sOut := sg.OutputType
	for range Only.Once {
		sg.OutputType.SetNone()

		ep := sg.Get("AppService.getPsList", "")
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsId()
	}
	sg.OutputType = sOut

	return ret
}

func fillArray(count int, args []string) []string {
	var ret []string
	for range Only.Once {
		//
		// if len(args) == 0 {
		//	break
		// }
		ret = make([]string, count)
		for i, e := range args {
			ret[i] = e
		}
	}
	return ret
}
