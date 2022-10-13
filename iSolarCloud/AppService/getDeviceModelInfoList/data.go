package getDeviceModelInfoList

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceModelInfoList"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	ComType           string `json:"com_type"`
	DeviceFactoryID   string `json:"device_factory_id"`
	DeviceFactoryName string `json:"device_factory_name"`
	DeviceModel       string `json:"device_model"`
	DeviceModelCode   string `json:"device_model_code"`
	DeviceModelID     api.Integer  `json:"device_model_id"`
	DeviceType        api.Integer  `json:"device_type"`
	IsRemoteUpgrade   api.Bool     `json:"is_remote_upgrade"`
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Model Id",
			"Type",
			"Com Type",
			"Factory Id",
			"Factory Name",
			"Model",
			"Model Code",
			"Is Remote Upgrade",
		)
		for _, d := range e.Response.ResultData {
			if d.DeviceModel == d.DeviceModelCode {
				d.DeviceModelCode = ""
			}
			_ = table.AddRow(
				d.DeviceModelID,
				d.DeviceType,
				d.ComType,
				d.DeviceFactoryID,
				d.DeviceFactoryName,
				d.DeviceModel,
				d.DeviceModelCode,
				d.IsRemoteUpgrade,
			)
		}
	}
	return table
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}
