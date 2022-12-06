package getDeviceTypeInfoById

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getDeviceTypeInfoById"
const Disabled = false
const EndPointName = "WebIscmAppService.getDeviceTypeInfoById"

type RequestData struct {
	CodeType       valueTypes.String `json:"code_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData   struct {
	PowerDeviceTypeMap struct {
		UpdateDate      valueTypes.DateTime  `json:"update_date" PointNameDateFormat:"2006-01-02 15:04:05"`
		SysId           valueTypes.String  `json:"sys_id"`
		SysName         valueTypes.String  `json:"sys_name"`
		TypeCode        valueTypes.Integer `json:"type_code"`
		TypeId          valueTypes.Integer `json:"type_id"`
		TypeName        valueTypes.String  `json:"type_name"`
		TypeNameEn      valueTypes.String  `json:"type_name_en"`
		IsRemoteUpgrade valueTypes.Bool `json:"is_remote_upgrade"`
		ValidFlag       valueTypes.Bool `json:"valid_flag"`
	} `json:"powerDeviceTypeMap" PointId:"power_device_type_map" DataTable:"true"`
	SysList []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"sysList" PointId:"sys_list" DataTable:"true" DataTableSort:"CodeValue"`
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
