package getPowerDeviceTypeList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)


const Url = "/v1/devService/getPowerDeviceTypeList"
const Disabled = false
const EndPointName = "WebIscmAppService.getPowerDeviceTypeList"

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
	GoStructParent        GoStruct.GoStructParent `json:"GoStruct" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"UpdateDate"`

	UpdateDate      valueTypes.DateTime     `json:"update_date" PointNameDateFormat:"DateTimeLayout"`
	SysId           valueTypes.String       `json:"sys_id"`
	SysName         valueTypes.String       `json:"sys_name"`
	TypeId          valueTypes.Integer      `json:"type_id"`
	TypeCode        valueTypes.Integer      `json:"type_code"`
	TypeName        valueTypes.String       `json:"type_name"`
	TypeNameEn      valueTypes.String       `json:"type_name_en"`
	IsRemoteUpgrade valueTypes.Bool         `json:"is_remote_upgrade"`
	ValidFlag       valueTypes.Bool         `json:"valid_flag"`
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
