package getPowerDeviceModelList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPowerDeviceModelList"
const Disabled = false
// ./gojson.sh WebIscmAppService.getPowerDeviceModelList '{"device_type":"44"}'

type RequestData struct {
	DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	DeviceModelCode   valueTypes.String  `json:"device_model_code"`
	InverterModelType valueTypes.Integer `json:"inverter_model_type"`
	IsCountryCheck    valueTypes.Integer `json:"is_country_check"`
	IsReadSet         valueTypes.Integer `json:"is_read_set"`
	IsReset           valueTypes.Integer `json:"is_reset"`
	IsThirdParty      valueTypes.Integer `json:"is_third_party"`
	ModelID           valueTypes.Integer `json:"model_id"`
	ModelName         valueTypes.String  `json:"model_name"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
