package iSolarCloud

import (
	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/queryUnitList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

func (sg *SunGrow) MetaUnitList() error {
	for range Only.Once {
		data := sg.NewSunGrowData()
		data.SetArgs()
		data.SetEndpoints(queryUnitList.EndPointName)

		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
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

		data := sg.NewSunGrowData()
		data.SetArgs("AppKey:" + appKey)
		data.SetEndpoints(getMqttConfigInfoByAppkey.EndPointName)

		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.Output()
		if sg.Error != nil {
			break
		}

		// ep := sg.GetByStruct(getMqttConfigInfoByAppkey.EndPointName,
		// 	getMqttConfigInfoByAppkey.RequestData{AppKey: valueTypes.SetStringValue(appKey)},
		// 	DefaultCacheTimeout,
		// )
		// if sg.IsError() {
		// 	break
		// }
		//
		// data := getMqttConfigInfoByAppkey.Assert(ep)
		// table := data.GetEndPointResultTable()
		// if table.Error != nil {
		// 	sg.Error = table.Error
		// 	break
		// }
		//
		// table.SetTitle("MQTT info")
		// table.SetFilePrefix(data.SetFilenamePrefix(""))
		// table.SetGraphFilter("")
		// table.SetSaveFile(sg.SaveAsFile)
		// table.OutputType = sg.OutputType
		// sg.Error = table.Output()
		// if sg.IsError() {
		// 	break
		// }
	}

	return sg.Error
}

func (sg *SunGrow) GetRealTimeData(psKey string) error {
	for range Only.Once {
		if psKey == "" {
			// var psKeys []string
			// psKeys, sg.Error = sg.GetPsKeys()
			// if sg.IsError() {
			// 	break
			// }
			// fmt.Printf("psKeys: %v\n", psKeys)
			// psKey = strings.Join(psKeys, ",")
		}
		// fmt.Println("TO FIX")
		// break

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
