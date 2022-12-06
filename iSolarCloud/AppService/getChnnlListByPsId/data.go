package getChnnlListByPsId

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getChnnlListByPsId"
const Disabled = false
const EndPointName = "AppService.getChnnlListByPsId"

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		GoStruct       GoStruct.GoStruct   `json:"-" PointIdFrom:"PsId" PointIdReplace:"true" PointIdReplace:"true"`

		PsId           valueTypes.Integer  `json:"ps_id"`
		Sn             valueTypes.String   `json:"sn"`
		ChannelId      valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		ChannelName    valueTypes.String   `json:"chnnl_name" PointId:"channel_name"`
		ChannelDesc    interface{}         `json:"chnnl_desc" PointId:"channel_description"`
		CreateTime     valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		UpdateDate     valueTypes.DateTime `json:"update_date" PointNameDateFormat:"2006-01-02 15:04:05"`
		DataFlag       valueTypes.Integer  `json:"data_flag"`
		DataFlagDetail valueTypes.Integer  `json:"data_flag_detail"`
		IsAuto         valueTypes.Bool     `json:"is_auto"`
		IsEnable       valueTypes.Bool     `json:"is_enable"`
		IsSure         valueTypes.Bool     `json:"is_sure"`
		ProtocolType   interface{}         `json:"protocol_type"`
	} `json:"pageList" PointId:"devices" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.EndPointPath{})
	return entries
}
