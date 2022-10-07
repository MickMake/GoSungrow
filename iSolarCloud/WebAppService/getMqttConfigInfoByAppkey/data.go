package getMqttConfigInfoByAppkey

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
)


const Url = "/v1/commonService/getMqttConfigInfoByAppkey"
const Disabled = false

const (
	WebAppKey = "B0455FBE7AA0328DB57B59AA729F05D8"
	LoginAppKey = "93D72E60331ABDCDC7B39ADC2D1F32B3"
)

type RequestData struct {
	AppKey string `json:"app_key" required:"true"`
}

// IsValid Checks for validity of results data.
func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

// Help provides more info to the user on request JSON fields.
func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

// ResultData holds data returned from the API.
type ResultData struct {
	Code             string   `json:"code"`
	MqttPassword     string   `json:"mqtt_password"`
	MqttRsaPublicKey string   `json:"mqtt_rsa_public_key"`
	MqttType         string   `json:"mqtt_type"`
	MqttURLList      []string `json:"mqtt_url_list"`
	MqttURLListLan   []string `json:"mqtt_url_list_lan"`
	MqttUsername     string   `json:"mqtt_username"`
}

// IsValid Checks for validity of results data.
func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// 	case e.Dummy == "":
	// 		break
	// 	default:
	// 		err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
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
// }

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		e.Error = table.SetHeader(
			"AppKey",
			"Name",
			"Value",
		)
		if e.Error != nil {
			break
		}

		// @TODO - Think about providing an apiReflect function that does this automatically.
		_ = table.AddRow(e.Request.AppKey, "Code", e.Response.ResultData.Code)
		_ = table.AddRow(e.Request.AppKey, "Mqtt Username", e.Response.ResultData.MqttUsername)
		_ = table.AddRow(e.Request.AppKey, "Mqtt Password", e.Response.ResultData.MqttPassword)
		_ = table.AddRow(e.Request.AppKey, "Mqtt Rsa Public Key", e.Response.ResultData.MqttRsaPublicKey)
		_ = table.AddRow(e.Request.AppKey, "Mqtt Type", e.Response.ResultData.MqttType)
		_ = table.AddRow(e.Request.AppKey, "Mqtt URL List", e.Response.ResultData.MqttURLList)
		_ = table.AddRow(e.Request.AppKey, "Mqtt URL List Lan", e.Response.ResultData.MqttURLListLan)

	}

	return table
}
