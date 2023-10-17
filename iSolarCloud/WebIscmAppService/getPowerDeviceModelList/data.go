package getPowerDeviceModelList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/getPowerDeviceModelList"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPowerDeviceModelList"
)

// ./gojson.sh WebIscmAppService.getPowerDeviceModelList '{"device_type":"44"}'

type RequestData struct {
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"ModelId"`

	ModelId           valueTypes.Integer `json:"model_id"`
	InverterModelType valueTypes.Integer `json:"inverter_model_type"`
	DeviceModelCode   valueTypes.String  `json:"device_model_code"`
	ModelName         valueTypes.String  `json:"model_name"`
	IsCountryCheck    valueTypes.Bool    `json:"is_country_check"`
	IsReadSet         valueTypes.Bool    `json:"is_read_set"`
	IsReset           valueTypes.Bool    `json:"is_reset"`
	IsThirdParty      valueTypes.Bool    `json:"is_third_party"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
