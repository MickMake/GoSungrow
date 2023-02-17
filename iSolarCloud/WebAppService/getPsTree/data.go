package getPsTree

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/devService/getPsTree"
const Disabled = false
const EndPointName = "WebAppService.getPsTree"

type RequestData struct {
	// @TODO - Fixup this up for iSolarCloud/data_request.go
	PsId3 valueTypes.PsId `json:"psid" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"PsId"` // PointIdFrom:"PsId" PointIdReplace:"true"`

	PsId      valueTypes.Integer `json:"psid" PointId:"ps_id"`
	Id        valueTypes.Integer `json:"id"`
	ChannelId valueTypes.Integer `json:"chnnlid" PointId:"channel_id"`
	Pid       valueTypes.Integer `json:"pid"`

	Name     valueTypes.String  `json:"name"`
	IsParent valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	Level    valueTypes.Integer `json:"level"`
	Unit     valueTypes.String  `json:"unit"`
	AType    interface{}        `json:"atype" PointId:"a_type"`
	CType    interface{}        `json:"ctype" PointId:"c_type"`
	NodeKey  valueTypes.Integer `json:"nodekey" PointId:"node_key"`
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
