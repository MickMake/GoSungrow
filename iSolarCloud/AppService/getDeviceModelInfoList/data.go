package getDeviceModelInfoList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceModelInfoList"
const Disabled = false

type RequestData struct {
	}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	ComType           valueTypes.String `json:"com_type"`
	DeviceFactoryId   valueTypes.String `json:"device_factory_id"`
	DeviceFactoryName valueTypes.String `json:"device_factory_name"`
	DeviceModel       valueTypes.String `json:"device_model"`
	DeviceModelCode   valueTypes.String `json:"device_model_code"`
	DeviceModelId     valueTypes.Integer  `json:"device_model_id"`
	DeviceType        valueTypes.Integer  `json:"device_type"`
	IsRemoteUpgrade   valueTypes.Bool     `json:"is_remote_upgrade"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetPointDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable(
			"Model Id",
			"Type",
			"Com Type",
			"Factory Id",
			"Factory Name",
			"Model",
			"Model Code",
			"Is Remote Upgrade",
		)
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		// _ = table.SetHeader(
		// 	"Model Id",
		// 	"Type",
		// 	"Com Type",
		// 	"Factory Id",
		// 	"Factory Name",
		// 	"Model",
		// 	"Model Code",
		// 	"Is Remote Upgrade",
		// )
		for _, d := range e.Response.ResultData {
			if d.DeviceModel == d.DeviceModelCode {
				d.DeviceModelCode.SetString("")
			}
			_ = table.AddRow(d.DeviceModelId.String(),
				d.DeviceType.String(),
				d.ComType.String(),
				d.DeviceFactoryId.String(),
				d.DeviceFactoryName.String(),
				d.DeviceModel.String(),
				d.DeviceModelCode.String(),
				d.IsRemoteUpgrade.String(),
			)
		}
	}
	return table
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
