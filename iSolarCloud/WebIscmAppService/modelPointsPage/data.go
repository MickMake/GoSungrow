package modelPointsPage

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
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
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PointList []struct {
		GoStructParent      GoStruct.GoStructParent  `json:"-" PointIdFromChild:"PointId" PointIdReplace:"false"`

		PointId             valueTypes.Integer `json:"point_id"`
		PointName           valueTypes.String  `json:"point_name"`
		CodeId              valueTypes.Integer `json:"code_id"`
		DeviceModelId       valueTypes.Integer `json:"device_model_id"`
		IsShow              valueTypes.Bool    `json:"is_show"`
		IsSupportSecondData valueTypes.Bool    `json:"is_support_second_data"`
		OrderNum            valueTypes.Integer `json:"order_num"`
	} `json:"pointList" PointId:"point_list" DataTable:"true" DataTableSortOn:"PointId"`
	ModelList []struct {
		GoStruct        GoStruct.GoStructParent   `json:"GoStruct" PointIdFromChild:"DeviceModelId" PointIdReplace:"false"`

		DeviceModelId   valueTypes.Integer `json:"device_model_id"`
		DeviceModel     valueTypes.String  `json:"device_model"`
		DeviceModelCode valueTypes.String  `json:"device_model_code"`
	} `json:"modelList" PointId:"model_list" DataTable:"true" DataTableSortOn:"DeviceModelId"`
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
