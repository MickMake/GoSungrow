package modelPointsPage

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/modelPointsPage"
const Disabled = false

type RequestData struct {
	DeviceModelId valueTypes.String `json:"device_model_id" required:"true"`
	DeviceType    valueTypes.String `json:"device_type"     required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ModelList []struct {
		DeviceModelId   valueTypes.Integer `json:"device_model_id"`
		DeviceModel     valueTypes.String  `json:"device_model"`
		DeviceModelCode valueTypes.String  `json:"device_model_code"`
	} `json:"modelList" PointId:"model_list" DataTable:"true" DataTableName:"Model List"`
	PointList []struct {
		PointId             valueTypes.Integer `json:"point_id" DataTableSort:"true"`
		PointName           valueTypes.String  `json:"point_name"`
		CodeId              valueTypes.Integer `json:"code_id"`
		DeviceModelId       valueTypes.Integer `json:"device_model_id"`
		IsShow              valueTypes.Bool    `json:"is_show"`
		IsSupportSecondData valueTypes.Bool    `json:"is_support_second_data"`
		OrderNum            valueTypes.Integer `json:"order_num"`
	} `json:"pointList" PointId:"point_list" DataTable:"true"`
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
