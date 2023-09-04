package getMqttConfigInfoByAppkey

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)


const Url = "/v1/commonService/getMqttConfigInfoByAppkey"
const Disabled = false
const EndPointName = "WebAppService.getMqttConfigInfoByAppkey"

const (
	AppKey      = "93D72E60331ABDCDC7B39ADC2D1F32B3"
	WebAppKey   = "B0455FBE7AA0328DB57B59AA729F05D8"
	LoginAppKey = "93D72E60331ABDCDC7B39ADC2D1F32B3"
)

type RequestData struct {
	AppKey valueTypes.String `json:"app_key" required:"true"`
}

// IsValid Checks for validity of results data.
func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
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
	MqttURLList      []valueTypes.String `json:"mqtt_url_list"`
	MqttURLListLan   []valueTypes.String `json:"mqtt_url_list_lan"`
	MqttUsername     string   `json:"mqtt_username"`
}

// IsValid Checks for validity of results data.
func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
