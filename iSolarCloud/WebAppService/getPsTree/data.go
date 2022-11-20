package getPsTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPsTree"
const Disabled = false

type RequestData struct {
	PsId     valueTypes.PsId `json:"psid" required:"true"`
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
	GoStruct       GoStruct.GoStruct        `json:"GoStruct" PointDeviceFrom:"PsId"`

	PsId      valueTypes.Integer `json:"psid"`
	Id        valueTypes.Integer `json:"id"`
	ChannelId valueTypes.Integer `json:"chnnlid" PointId:"channel_id"`
	Pid       valueTypes.Integer `json:"pid"`
	Name      valueTypes.String  `json:"name"`
	IsParent  valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	Level     valueTypes.Integer `json:"level"`
	Unit      valueTypes.String  `json:"unit"`
	AType     interface{}        `json:"atype"`
	CType     interface{}        `json:"ctype"`
	NodeKey   valueTypes.Integer `json:"nodekey" PointId:"node_key"`
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
