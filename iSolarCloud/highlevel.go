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
	"GoSungrow/iSolarCloud/WebAppService/showPSView"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
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
