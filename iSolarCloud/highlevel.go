package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getPsListStaticData"
	"GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"GoSungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"GoSungrow/iSolarCloud/WebAppService/showPSView"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
	"time"
)


func (sg *SunGrow) AllCritical() error {
	var ep api.EndPoint
	for range Only.Once {
		ep = sg.GetByJson(powerDevicePointList.EndPointName, "")
		if sg.IsError() {
			break
		}

		ep = sg.GetByJson(getPsList.EndPointName, "")
		if sg.IsError() {
			break
		}

		_getPsList := getPsList.AssertResultData(ep)

		for _, psId := range _getPsList.GetPsIds() {
			ep = sg.GetByJson(queryDeviceList.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(queryDeviceListForApp.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(showPSView.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			// ep = sg.GetByJson(findPsType.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if sg.IsError() {
			// break
			// }

			ep = sg.GetByJson(getPowerStatistics.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(getPsDetail.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(getPsDetailWithPsType.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(getPsHealthState.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(getPsListStaticData.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			ep = sg.GetByJson(getPsWeatherList.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
				break
			}

			// ep = sg.GetByJson(queryAllPsIdAndName.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if sg.IsError() {
			// break
			// }

			// ep = sg.GetByJson(queryDeviceListByUserId.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			// if sg.IsError() {
			// break
			// }

			ep = sg.GetByJson(queryDeviceListForApp.EndPointName, fmt.Sprintf(`{"ps_id":"%d"}`, psId.Value()))
			if sg.IsError() {
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
		ep = sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		_getPsList := getPsList.Assert(ep)
		table := _getPsList.GetEndPointResultTable()
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
		if sg.IsError() {
			break
		}

		for _, psId := range _getPsList.GetPsIds() {
			ep = sg.GetByStruct(queryDeviceList.EndPointName,
				// queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
				queryDeviceList.RequestData{PsId: psId},
				time.Second*60,
			)
			if sg.IsError() {
				break
			}

			data := queryDeviceList.Assert(ep)

			table = data.GetEndPointResultTable()
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
			if sg.IsError() {
				break
			}
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetIsolarcloudMqtt(appKey string) error {
	for range Only.Once {
		if appKey == "" {
			appKey = sg.GetAppKey()
		}

		ep := sg.GetByStruct(getMqttConfigInfoByAppkey.EndPointName,
			getMqttConfigInfoByAppkey.RequestData{AppKey: valueTypes.SetStringValue(appKey)},
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}

		data := getMqttConfigInfoByAppkey.Assert(ep)
		table := data.GetEndPointResultTable()
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
		if sg.IsError() {
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
			if sg.IsError() {
				break
			}
			fmt.Printf("psKeys: %v\n", psKeys)
			psKey = strings.Join(psKeys, ",")
		}

		ep := sg.GetByStruct(queryDeviceRealTimeDataByPsKeys.EndPointName,
			queryDeviceRealTimeDataByPsKeys.RequestData{PsKeyList: valueTypes.SetStringValue(psKey)},
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}

		data := queryDeviceRealTimeDataByPsKeys.Assert(ep)
		table := data.GetEndPointResultTable()
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
		if sg.IsError() {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) CmdDataPsDetail(psIds ...string) error {
	for range Only.Once {
		data := sg.NewSunGrowData()
		data.SetPsIds(psIds...)
		data.SetEndpoints(getPsDetail.EndPointName, getPsDetailWithPsType.EndPointName)

		sg.Error = data.GetData()
		if sg.IsError() {
			break
		}

		sg.Error = data.Output()
		if sg.IsError() {
			break
		}
	}

	return sg.Error
}

// func (sg *SunGrow) GetPointData(date string, pointNames api.TemplatePoints, psIds ...valueTypes.PsId) error {
// 	for range Only.Once {
// 		if len(pointNames) == 0 {
// 			sg.Error = errors.New("no points defined")
// 			break
// 		}
//
// 		if date == "" {
// 			date = valueTypes.NewDateTime("").String()
// 		}
// 		when := valueTypes.NewDateTime(date)
//
// 		if len(psIds) == 0 {
// 			psIds, sg.Error = sg.StringToPids()
// 			if sg.IsError() {
// 				break
// 			}
// 		}
//
// 		for _, psId := range psIds {
// 			ep := sg.GetByStruct(
// 				"AppService.queryMutiPointDataList",
// 				queryMutiPointDataList.RequestData{
// 					PsId:           psId,
// 					PsKeys:         valueTypes.SetPsKeyString(pointNames.PrintKeys()),      // @TODO - Fixup!
// 					Points:         valueTypes.SetPointIdsString(pointNames.PrintPoints()), // @TODO - Fixup!
// 					MinuteInterval: valueTypes.SetIntegerValue(5),                          // @TODO - Fixup!
// 					StartTimeStamp: valueTypes.SetDateTimeString(when.GetDayStartTimestamp()), // @TODO - Fixup!
// 					EndTimeStamp:   valueTypes.SetDateTimeString(when.GetDayEndTimestamp()),   // @TODO - Fixup!
// 				},
// 				DefaultCacheTimeout,
// 			)
// 			if sg.IsError() {
// 				break
// 			}
// 			if ep == nil {
// 				break
// 			}
//
// 			// data := queryMutiPointDataList.Assert(ep)
// 			// table := data.GetPointDataTable(pointNames)
// 			// if table.Error != nil {
// 			// 	sg.Error = table.Error
// 			// 	break
// 			// }
// 			//
// 			// table.SetTitle("Point Data %s", psId)
// 			// table.SetFilePrefix(data.SetFilenamePrefix("%d", psId))
// 			// table.SetGraphFilter("")
// 			// table.SetSaveFile(sg.SaveAsFile)
// 			// table.OutputType = sg.OutputType
// 			// sg.Error = table.Output()
// 			// if sg.IsError() {
// 			// 	break
// 			// }
// 		}
// 	}
//
// 	return sg.Error
// }

// func (sg *SunGrow) SearchPointNames(pns ...string) error {
// 	for range Only.Once {
// 		table := output.NewTable(
// 			"DeviceType",
// 			"Id",
// 			"Period",
// 			"Point Id",
// 			"Point Name",
// 			"Show Point Name",
// 			"Translation Id",
// 		)
//
// 		// _ = table.SetHeader(
// 		// 	"DeviceType",
// 		// 	"Id",
// 		// 	"Period",
// 		// 	"Point Id",
// 		// 	"Point Name",
// 		// 	"Show Point Name",
// 		// 	"Translation Id",
// 		// )
//
// 		if len(pns) == 0 {
// 			fmt.Println("Searching up to id 1000 within getPowerDevicePointInfo")
// 			for pni := 0; pni < 1000; pni++ {
// 				PrintPause(pni, 20)
//
// 				ep := sg.GetByStruct(
// 					"AppService.getPowerDevicePointInfo",
// 					// getPowerDevicePointInfo.RequestData{Id: strconv.FormatInt(int64(pni), 10)},
// 					getPowerDevicePointInfo.RequestData{Id: valueTypes.SetIntegerValue(int64(pni)) },
// 					DefaultCacheTimeout,
// 				)
// 				if sg.IsError() {
// 					break
// 				}
//
// 				data := getPowerDevicePointInfo.Assert(ep)
// 				table = data.AddDataTable(table)
// 				if table.Error != nil {
// 					sg.Error = table.Error
// 					break
// 				}
// 			}
// 			fmt.Println("")
// 		} else {
// 			fmt.Printf("Searching for %v within getPowerDevicePointInfo\n", pns)
// 			for _, pn := range pns {
// 				ep := sg.GetByStruct(
// 					"AppService.getPowerDevicePointInfo",
// 					getPowerDevicePointInfo.RequestData{Id: valueTypes.SetIntegerString(pn)},
// 					DefaultCacheTimeout,
// 				)
// 				if sg.IsError() {
// 					break
// 				}
//
// 				data := getPowerDevicePointInfo.Assert(ep)
// 				table = data.GetPointDataTable()
// 				if table.Error != nil {
// 					sg.Error = table.Error
// 					break
// 				}
// 			}
// 			fmt.Println("")
// 		}
//
// 		table.SetTitle("Point Names")
// 		table.SetFilePrefix("AppService.getPowerDevicePointInfo")
// 		table.SetGraphFilter("")
// 		table.SetSaveFile(sg.SaveAsFile)
// 		table.OutputType = sg.OutputType
// 		sg.Error = table.Output()
// 		if sg.IsError() {
// 			break
// 		}
// 	}
//
// 	return sg.Error
// }

// func (sg *SunGrow) GetPointName(id valueTypes.Integer) error {
// 	for range Only.Once {
// 		ep := sg.GetByStruct(
// 			"AppService.getPowerDevicePointInfo",
// 			getPowerDevicePointInfo.RequestData{Id: id},
// 			DefaultCacheTimeout,
// 		)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		data := getPowerDevicePointInfo.Assert(ep)
// 		table := data.GetPointDataTable()
// 		if table.Error != nil {
// 			sg.Error = table.Error
// 			break
// 		}
//
// 		// table2 := data.GetEndPointData()
// 		// fmt.Printf("%v\n", table2)
//
// 		table.SetTitle("Point Id %s", id)
// 		table.SetFilePrefix(data.SetFilenamePrefix(""))
// 		table.SetGraphFilter("")
// 		table.SetSaveFile(sg.SaveAsFile)
// 		table.OutputType = sg.OutputType
// 		sg.Error = table.Output()
// 		if sg.IsError() {
// 			break
// 		}
// 	}
//
// 	return sg.Error
// }

// func (sg *SunGrow) StringToPids(pids ...string) ([]valueTypes.PsId, error) {
// 	var psIds []valueTypes.PsId
// 	for range Only.Once {
// 		var psids valueTypes.PsIds
// 		psids = sg.SetPsIds(pids...)
// 		if sg.Error != nil {
// 			break
// 		}
//
// 		// for _, pid := range pids {
// 		// 	// p, err := strconv.ParseInt(pid, 10, 64)
// 		// 	// if err != nil {
// 		// 	// 	continue
// 		// 	// }
// 		// 	psIds = append(psIds, valueTypes.SetPsIdString(pid))
// 		// }
// 		// if len(psIds) == 0 {
// 		// 	psIds, sg.Error = sg.PsIds()
// 		// 	if sg.IsError() {
// 		// 		break
// 		// 	}
// 		// }
// 	}
//
// 	return psIds, sg.Error
// }

func (sg *SunGrow) GetPsNames() ([]string, error) {
	var ret []string

	for range Only.Once {
		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsName()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsModels() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psIds valueTypes.PsIds
		psIds, sg.Error = sg.PsIds()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if sg.IsError() {
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
		var psIds valueTypes.PsIds
		psIds, sg.Error = sg.PsIds()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if sg.IsError() {
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
		var psIds valueTypes.PsIds
		psIds, sg.Error = sg.PsIds()
		if sg.Error != nil {
			break
		}

		for _, psId := range psIds {
			ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
				// getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
				getPsDetailWithPsType.RequestData{PsId: psId},
				DefaultCacheTimeout)
			if sg.IsError() {
				break
			}

			data := getPsDetailWithPsType.Assert(ep)
			ret = append(ret, data.GetPsKeys()...)
		}
	}

	return ret, sg.Error
}
