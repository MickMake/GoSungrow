package getPowerDeviceChannl

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceChannl"
const Disabled = false

type RequestData struct {
	Id valueTypes.Integer `json:"id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Id                 valueTypes.Integer `json:"id"`
	PsId               valueTypes.Integer `json:"psid" PointId:"ps_id"`
	ChannelDescription interface{}        `json:"channeldesc" PointId:"channel_desc"`
	ChannelId          valueTypes.Integer `json:"channelid" PointId:"channel_id"`
	ChannelName        valueTypes.String  `json:"channelname" PointId:"channel_name"`
	ChannelId2         valueTypes.Integer `json:"chnnl_id" PointId:"channel_id2"`
	CrtUsername        interface{}        `json:"crtusername" PointId:"crt_username"`
	DataFlag           valueTypes.Integer `json:"dataflag" PointId:"data_flag"`
	FlagServer         interface{}        `json:"flagserver" PointId:"flag_server"`
	ProtocolType       interface{}        `json:"protocoltype" PointId:"protocol_type"`
	HostIp             interface{}        `json:"hostip" PointId:"host_ip"`
	TcpMode            interface{}        `json:"tcpmode" PointId:"tcp_mode"`
	TcpPort            interface{}        `json:"tcpport" PointId:"tcp_port"`
	IsSure             valueTypes.Bool    `json:"is_sure"`
	IsEnable           valueTypes.Bool    `json:"isenable" PointId:"is_enable"`
	PsGuid             valueTypes.String  `json:"psguid" PointId:"ps_guid"`
	SnCode             valueTypes.String  `json:"sncode" PointId:"sn_code"`
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
