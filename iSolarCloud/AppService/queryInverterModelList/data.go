package queryInverterModelList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryInverterModelList"
const Disabled = false
const EndPointName = "AppService.queryInverterModelList"

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
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"DeviceModelId"`

	DeviceModelId valueTypes.Integer `json:"device_model_id"`
	DeviceModel   valueTypes.String  `json:"device_model"`
	ModelCode     valueTypes.String  `json:"model_code"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
