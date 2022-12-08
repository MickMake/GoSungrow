package getDevicePointAttrs

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getDevicePointAttrs"
const Disabled = false
const EndPointName = "WebAppService.getDevicePointAttrs"

type RequestData struct {
	Uuid        valueTypes.Integer `json:"uuid,omitempty"`
	DeviceType2 valueTypes.Integer `json:"deviceType" required:"true"`
	PsId2       valueTypes.PsId    `json:"psId" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []Point

// type Points []Point

type Point struct {
	GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"Id"`
	// GoStruct.GoStruct `json:"-" PointIdFrom:"PsId" PointIdReplace:"false"`

	PsId             valueTypes.Integer `json:"psid" PointId:"ps_id"`
	DeviceType       valueTypes.Integer `json:"pid" PointId:"device_type"`
	ChannelId        valueTypes.Integer `json:"chnnlid" PointId:"channel_id"`

	StationName      valueTypes.String  `json:"stationname" PointId:"station_name"`
	StationShortName valueTypes.String  `json:"stationshortname" PointId:"station_short_name"`
	IsParent         valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	DeviceModelId    valueTypes.Integer `json:"device_model_id"`
	DeviceName       valueTypes.String  `json:"devicename" PointId:"device_name"`
	AType            valueTypes.Integer `json:"atype" PointId:"a_type"`
	CodeId           valueTypes.Integer `json:"code_id"`
	CType            valueTypes.Integer `json:"ctype" PointId:"c_type"`

	Id               valueTypes.PointId `json:"id"`
	NodeKey          valueTypes.PointId `json:"nodekey" PointId:"node_key"`
	Name             valueTypes.String  `json:"name"`
	TargetUnit       valueTypes.String  `json:"target_unit"`
	Unit             valueTypes.String  `json:"unit"`
	UnitType         valueTypes.Integer `json:"unit_type"`
	Level            valueTypes.Integer `json:"level"`
	OrderId          valueTypes.Integer `json:"orderid" PointId:"order_id"`
	PointGroupId     valueTypes.Integer `json:"point_group_id"`
	Relate           valueTypes.Integer `json:"relate"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId2.String(), GoStruct.NewEndPointPath(e.Request.PsId2.String()))
	return entries
}

func (e *EndPoint) Points() []Point {
	return e.Response.ResultData
}
