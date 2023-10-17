package getUserMenuLs

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/userService/getUserMenuLs"
	Disabled     = false
	EndPointName = "WebIscmAppService.getUserMenuLs"
)

type RequestData struct {
	UserId valueTypes.String `json:"userId" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"MenuId"`
	// GoStruct         GoStruct.GoStruct        `json:"-" PointIdFrom:"MenuId" PointIdReplace:"false"`

	MenuId        valueTypes.Integer `json:"menuid" PointId:"menu_id"`
	MenuLevel     valueTypes.Integer `json:"menulevel" PointId:"menu_level"`
	MenuName      valueTypes.String  `json:"menuname" PointId:"menu_name"`
	MenuOrder     valueTypes.Integer `json:"menuorder" PointId:"menu_order"`
	MenuCode      valueTypes.String  `json:"menucode" PointId:"menu_code"`
	MenuType      valueTypes.String  `json:"menutype" PointId:"menu_type"`
	MenuUrl       valueTypes.String  `json:"menuurl" PointId:"menu_url"`
	OpenType      valueTypes.String  `json:"opentype"`
	ParMenuId     valueTypes.Integer `json:"parmenuid" PointId:"par_menu_id"`
	PrivilegeCode valueTypes.String  `json:"privilegecode" PointId:"privilege_code"`
	UrlTarget     valueTypes.String  `json:"urltarget" PointId:"url_target"`
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
