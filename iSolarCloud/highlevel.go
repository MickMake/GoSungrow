package iSolarCloud

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsDetail"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsListStaticData"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebAppService/showPSView"
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
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

// func (sg *SunGrow) PrintCurrentStats() error {
// 	var ep api.EndPoint
// 	for range Only.Once {
// 		ep = sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		_getPsList := getPsList.Assert(ep)
// 		table := _getPsList.GetEndPointResultTable()
// 		if table.Error != nil {
// 			sg.Error = table.Error
// 			break
// 		}
//
// 		table.SetTitle("getPsList")
// 		table.SetFilePrefix(_getPsList.SetFilenamePrefix(""))
// 		table.SetGraphFilter("")
// 		table.SetSaveFile(sg.SaveAsFile)
// 		table.OutputType = sg.OutputType
// 		sg.Error = table.Output()
// 		if sg.IsError() {
// 			break
// 		}
//
// 		for _, psId := range _getPsList.GetPsIds() {
// 			ep = sg.GetByStruct(queryDeviceList.EndPointName,
// 				// queryDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
// 				queryDeviceList.RequestData{PsId: psId},
// 				time.Second*60,
// 			)
// 			if sg.IsError() {
// 				break
// 			}
//
// 			data := queryDeviceList.Assert(ep)
//
// 			table = data.GetEndPointResultTable()
// 			if table.Error != nil {
// 				sg.Error = table.Error
// 				break
// 			}
//
// 			table.SetTitle("queryDeviceList %s", psId)
// 			table.SetFilePrefix(data.SetFilenamePrefix("%s", psId))
// 			table.SetGraphFilter("")
// 			table.SetSaveFile(sg.SaveAsFile)
// 			table.OutputType = sg.OutputType
// 			sg.Error = table.Output()
// 			if sg.IsError() {
// 				break
// 			}
// 		}
// 	}
//
// 	return sg.Error
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
		psIds, sg.Error = sg.GetPsIds()
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
		psIds, sg.Error = sg.GetPsIds()
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
