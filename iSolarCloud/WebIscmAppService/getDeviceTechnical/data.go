package getDeviceTechnical

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getDeviceTechnical"
const Disabled = false
const EndPointName = "WebIscmAppService.getDeviceTechnical"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceTypeList []struct {
		TypeId   valueTypes.Integer `json:"type_id"`
		TypeName valueTypes.String  `json:"type_name"`
	} `json:"deviceTypeList" PointId:"device_type_list" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"TypeCode"`
	PowerDeviceTechnicalMap struct {
		ComType            valueTypes.String   `json:"com_type"`
		DeviceModel        valueTypes.String   `json:"device_model"`
		DeviceModelCode    valueTypes.String   `json:"device_model_code"`
		ModelId            valueTypes.Integer  `json:"model_id"`
		TechCode           valueTypes.String   `json:"tech_code"`
		TechContent        valueTypes.String   `json:"tech_content"`
		TechContentTransId valueTypes.Integer  `json:"tech_content_trans_id"`
		TechCreateTime     valueTypes.DateTime `json:"tech_createtime" PointId:"tech_create_time" PointNameDateFormat:"DateTimeLayout"`
		TechCreator        valueTypes.String   `json:"tech_creator"`
		TechDescription    interface{}         `json:"tech_description"`
		TechId             valueTypes.Integer  `json:"tech_id"`
		TechModifier       interface{}         `json:"tech_modifier"`
		TechModifyTime     valueTypes.DateTime `json:"tech_modifytime" PointId:"tech_modify_time" PointNameDateFormat:"DateTimeLayout"`
		TechName           valueTypes.String   `json:"tech_name"`
		TypeId             valueTypes.Integer  `json:"type_id"`
		TypeName           valueTypes.String   `json:"type_name"`
	} `json:"powerDeviceTechnicalMap" PointId:"power_device_technical_map"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
