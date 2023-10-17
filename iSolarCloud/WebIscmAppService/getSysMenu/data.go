package getSysMenu

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/userService/getSysMenu"
	Disabled     = false
	EndPointName = "WebIscmAppService.getSysMenu"
)

type RequestData struct {
	MenuId valueTypes.Integer `json:"menuId" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true"`

	MenuId           valueTypes.Integer `json:"menuid" PointId:"menu_id"`
	MenuLevel        valueTypes.Integer `json:"menulevel" PointId:"menu_level"`
	MenuName         valueTypes.String  `json:"menuname" PointId:"menu_name"`
	MenuOrder        valueTypes.Integer `json:"menuorder" PointId:"menu_order"`
	MenuType         valueTypes.String  `json:"menutype" PointId:"menu_type"`
	MenuUrl          valueTypes.String  `json:"menuurl" PointId:"menu_url"`
	MenuCode         valueTypes.String  `json:"menucode" PointId:"menu_code"`
	MenuDesc         valueTypes.String  `json:"menudesc" PointId:"menu_desc"`
	MenuValidFlag    valueTypes.Bool    `json:"menuvalidflag" PointId:"menu_valid_flag"`
	IsThirdPlatform  valueTypes.Bool    `json:"isthirdplatform" PointId:"is_third_platform"`
	OpenType         valueTypes.String  `json:"opentype" PointId:"open_type"`
	UrlTarget        valueTypes.String  `json:"urltarget" PointId:"url_target"`
	FatherMenu       valueTypes.String  `json:"fathermenu" PointId:"father_menu"`
	FatherMenuId     valueTypes.Integer `json:"fathermenuid" PointId:"father_menu_id"`
	Belongs          interface{}        `json:"belongs" PointId:"belongs"`
	IconFileId       interface{}        `json:"iconfileid" PointId:"iconfile_id"`
	IconUrl          interface{}        `json:"iconurl" PointId:"icon_url"`
	IsOpen           interface{}        `json:"isopen" PointId:"is_open"`
	PrivilegeCodeStr interface{}        `json:"privilegecodestr" PointId:"privilege_code_str"`
	ReportId         interface{}        `json:"report_id" PointId:"report_id"`
	VueIcon          interface{}        `json:"vueIcon" PointId:"vue_icon"`
	VuePath          interface{}        `json:"vuePath" PointId:"vue_path"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.MenuId.String(), GoStruct.NewEndPointPath(e.Request.MenuId.String()))
	return entries
}
