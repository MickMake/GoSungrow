package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getTemplateList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"strconv"
	"time"
)


func (sg *SunGrow) GetByJson(endpoint string, request string) api.EndPoint {
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
			ret = ret.SetRequestByJson(output.Json(request))
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
				sg.Error = ret.WriteDataFile()

			case sg.OutputType.IsRaw():
				fmt.Println(ret.GetJsonData(true))

			case sg.OutputType.IsJson():
				fmt.Println(ret.GetJsonData(false))

			default:
		}
	}
	return ret
}

func (sg *SunGrow) GetByStruct(endpoint string, request interface{}, cache time.Duration) api.EndPoint {
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

		ret = ret.SetCacheTimeout(cache)
		// if ret.CheckCache() {
		// 	ret = ret.ReadCache()
		// 	if !ret.IsError() {
		// 		break
		// 	}
		// }

		ret = ret.Call()
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		// sg.Error = ret.WriteCache()
		// if sg.Error != nil {
		// 	break
		// }
	}

	return ret
}

// func (sg *SunGrow) GetHighLevel(name string, args ...string) error {
// 	for range Only.Once {
// 		name = strings.ToLower(name)
// 		// if name == "stats" {
// 		// 	sg.Error = sg.GetCurrentStats()
// 		// 	break
// 		// }
// 		//
// 		// if name == "template" {
// 		// 	args = fillArray(3, args)
// 		// 	if args[0] == "" {
// 		// 		sg.Error = errors.New("need a date")
// 		// 		break
// 		// 	}
// 		// 	sg.Error = sg.GetTemplateData(args[0], args[1], args[2])
// 		// 	break
// 		// }
// 		//
// 		// if name == "template-points" {
// 		// 	args = fillArray(1, args)
// 		// 	sg.Error = sg.GetTemplatePoints(args[0])
// 		// 	break
// 		// }
//
// 		// if name == "points" {
// 		// 	args = fillArray(2, args)
// 		// 	if args[0] == "" {
// 		// 		sg.Error = errors.New("need a date")
// 		// 		break
// 		// 	}
// 		// 	points := api.CreatePoints(args)
// 		// 	sg.Error = sg.GetPointData(args[0], points)
// 		// 	break
// 		// }
//
// 		sg.Error = errors.New("unknown high-level command")
// 	}
// 	return sg.Error
// }
//
// func (sg *SunGrow) ListHighLevel() {
// 	fmt.Println("stats - Get current inverter stats, (last 5 minutes).")
// 	fmt.Println("\tdata get stats")
// 	fmt.Println("")
//
// 	fmt.Println("template [date] [template_id] - Get data from template.")
// 	fmt.Println("\tdata get template - Get data using default template 8042 for today.")
// 	fmt.Println("\tdata get template 2022 8040 - Get year data for template 8040 for the year 2022.")
// 	fmt.Println("\tdata get template 202202 8040 - Get month data for template 8040 for the month 202202.")
// 	fmt.Println("\tdata get template 20220202 8040 - Get day data for template 8040 for the day 20220202.")
// 	fmt.Println("\tdata get template 2022 - Get year data for default template 8042 for the year 2022.")
// 	fmt.Println("")
//
// 	fmt.Println("points <date> <device_id.point_id> ... - Get data from points list.")
// 	fmt.Println("")
// }

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

func (sg *SunGrow) GetCurrentStats() error {
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

		sg.Error = sg.Output(_getPsList, table, "")
		if sg.Error != nil {
			break
		}


		// ep = sg.GetByJson("AppService.queryDeviceList", fmt.Sprintf(`{"ps_id":"%d"}`, psId))
		ep = sg.GetByStruct(
			"AppService.queryDeviceList",
			queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout,
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

		sg.Error = sg.Output(ep2, table, "")
		if sg.Error != nil {
			break
		}

		// if sg.OutputType.IsHuman() {
		// 	_queryDeviceList := queryDeviceList.AssertResultData(ep)
		// 	points := _queryDeviceList.GetDataByName("SH10RT")
		// 	// @TODO - Improve output of this!
		// 	fmt.Printf("%v", points)
		// }
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
		psId := sg.GetPsId()

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

		sg.Error = sg.Output(ep2, table, "")
		if sg.Error != nil {
			break
		}

		//
		// csv := api.NewCsv()
		// sg.Error = csv.SetHeader([]string{
		// 	"Date/Time",
		// 	"PointId Name",
		// 	"Point Name",
		// 	"Value",
		// 	"Units",
		// })
		//
		// data := queryMutiPointDataList.AssertResultData(ep)
		// for deviceName, deviceRef := range data.Devices {
		// 	for pointId, pointRef := range deviceRef.Points {
		// 		for _, tim := range pointRef.Times {
		// 			gp := pointNames.GetPoint(deviceName, pointId)
		// 			_ = csv.AddRow([]string {
		// 				tim.Key.PrintFull(),
		// 				deviceName,
		// 				fmt.Sprintf("%s (%s)", gp.Description, pointId),
		// 				tim.Value,
		// 				gp.Unit,
		// 			})
		// 		}
		// 	}
		// }
		//
		// switch {
		// 	case sg.OutputType.IsNone():
		//
		// 	case sg.OutputType.IsHuman():
		// 		table.Print()
		//
		// 	case sg.OutputType.IsFile():
		// 		// a := queryMutiPointDataList.Assert(ep)
		// 		// suffix := fmt.Sprintf("%s-%s", when, "data")
		// 		// fn := a.GetCsvFilename()
		// 		sg.Error = table.WriteCsvFile()
		//
		// 	case sg.OutputType.IsRaw():
		// 		fmt.Println(ep.GetJsonData(true))
		//
		// 	case sg.OutputType.IsJson():
		// 		fmt.Println(ep.GetJsonData(false))
		//
		// 	default:
		// }
	}

	return sg.Error
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

			sg.Error = sg.Output(ep2, table, "")
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

		sg.Error = sg.Output(ep2, table, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) Output(endpoint api.EndPoint, table output.Table, graphFilter string) error {
	for range Only.Once {
		switch {
			case sg.OutputType.IsNone():

			case sg.OutputType.IsHuman():
				table.Print()

			case sg.OutputType.IsFile():
				// a := queryMutiPointDataList.Assert(endpoint)
				// suffix := fmt.Sprintf("%s-%s", when, "data")
				// fn := a.GetCsvFilename()
				sg.Error = table.WriteCsvFile()

			case sg.OutputType.IsRaw():
				fmt.Println(endpoint.GetJsonData(true))

			case sg.OutputType.IsJson():
				fmt.Println(endpoint.GetJsonData(false))

			case sg.OutputType.IsGraph():
				sg.Error = table.SetGraphFromJson(output.Json(graphFilter))
				if sg.Error != nil {
					break
				}
				sg.Error = table.CreateGraph()
				if sg.Error != nil {
					break
				}
				//

			default:
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPsId() int64 {
	var ret int64

	// sOut := sg.OutputType
	for range Only.Once {
		// sg.OutputType.SetNone()

		// ep := sg.GetByJson("AppService.getPsList", "")
		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsId()
	}
	// sg.OutputType = sOut

	return ret
}

func fillArray(count int, args []string) []string {
	var ret []string
	for range Only.Once {
		if count < len(args) {
			count = len(args)
		}
		ret = make([]string, count)
		for i, e := range args {
			ret[i] = e
		}
	}
	return ret
}

// func (sg *SunGrow) Graph(csv api.Table, timeCol int, ) {
// 	foo := New("Testing 1. 2. 3.")
//
// 	err := foo.SetFilename("HelloWorld.png")
// 	fmt.Println(err)
//
// 	now := time.Now()
// 	var times []time.Time
// 	for i := 0; i < 16; i++ {
// 		now = now.Add(time.Minute * 5)
// 		times = append(times, now)
// 	}
// 	err = foo.SetX("Date", times...)
// 	fmt.Println(err)
//
// 	var values []float64
// 	for i := 0; i < 16; i++ {
// 		then := (float64(i) * Randy(-200, 500)) +  Randy(-5000, 10000)
// 		values = append(values, then)
// 	}
//
// 	err = foo.SetY("Power", values...)
// 	fmt.Println(err)
//
// 	foo.SetRangeY(-6000, 12000)
//
// 	err = foo.Generate()
// 	fmt.Println(err)
// }
