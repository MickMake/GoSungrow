package getPsTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPsTree"
const Disabled = false

type RequestData struct {
	// @TODO - Fixup this up for iSolarCloud/data_request.go
	PsId3     valueTypes.PsId `json:"psid" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" PointIdFromChild:"PsId" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"PsId"`
	GoStruct       GoStruct.GoStruct        `json:"-" PointDeviceFrom:"PsId"`

	PsId      valueTypes.Integer `json:"psid" PointId:"ps_id"`
	Id        valueTypes.Integer `json:"id"`
	ChannelId valueTypes.Integer `json:"chnnlid" PointId:"channel_id"`
	Pid       valueTypes.Integer `json:"pid"`
	Name      valueTypes.String  `json:"name"`
	IsParent  valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	Level     valueTypes.Integer `json:"level"`
	Unit      valueTypes.String  `json:"unit"`
	AType     interface{}        `json:"atype" PointId:"a_type"`
	CType     interface{}        `json:"ctype" PointId:"c_type"`
	NodeKey   valueTypes.Integer `json:"nodekey" PointId:"node_key"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId3.String(), GoStruct.NewEndPointPath(e.Request.PsId3.String()))
	return entries
}
