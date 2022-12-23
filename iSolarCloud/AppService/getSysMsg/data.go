package getSysMsg

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/otherService/getSysMsg"
const Disabled = false
const EndPointName = "AppService.getSysMsg"

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
	GoStructParent        GoStruct.GoStructParent   `json:"-" DataTable:"true" DataTableSortOn:"CreateTime" PointIdReplace:"false"`

	MsgId           valueTypes.Integer  `json:"msg_id"`
	CreateTime      valueTypes.DateTime `json:"create_time" PointNameDateFormat:"DateTimeLayout"`
	PublishTime     valueTypes.DateTime `json:"publish_time" PointNameDateFormat:"DateTimeLayout"`
	StartTime       valueTypes.DateTime `json:"start_time" PointNameDateFormat:"DateTimeLayout"`
	EndTime         valueTypes.DateTime `json:"end_time" PointNameDateFormat:"DateTimeLayout"`
	Ispublish       valueTypes.Bool     `json:"ispublish" PointId:"is_publish"`
	UserName        valueTypes.String   `json:"user_name"`
	UserId          valueTypes.Integer  `json:"user_id"`
	MsgUserIds      valueTypes.String   `json:"msg_user_ids"`
	MsgLevel        valueTypes.Integer  `json:"msg_level"`
	MsgTitleEnUs    valueTypes.String   `json:"msg_title_en_us"`
	MsgContentsEnUs valueTypes.String   `json:"msg_contents_en_us"`

	TypeId          valueTypes.Integer  `json:"type_id"`
	BusinessId      valueTypes.Integer  `json:"business_id"`
	ClickNum        valueTypes.Integer  `json:"click_num"`
	RemindType      valueTypes.Integer  `json:"remind_type"`
	OperateType     valueTypes.Integer  `json:"operate_type"`
	OperateURL      valueTypes.String   `json:"operate_url"`
	AttachNames     interface{}         `json:"attach_names"`
	AttachPath      interface{}         `json:"attach_path"`

	MsgContents     valueTypes.String   `json:"msg_contents" PointIgnore:"true"`
	MsgContentsDeDe valueTypes.String   `json:"msg_contents_de_de" PointIgnore:"true"`
	MsgContentsEsEs valueTypes.String   `json:"msg_contents_es_es" PointIgnore:"true"`
	MsgContentsFrFr valueTypes.String   `json:"msg_contents_fr_fr" PointIgnore:"true"`
	MsgContentsIDID valueTypes.String   `json:"msg_contents_id_id" PointIgnore:"true"`
	MsgContentsItIt valueTypes.String   `json:"msg_contents_it_it" PointIgnore:"true"`
	MsgContentsJaJp valueTypes.String   `json:"msg_contents_ja_jp" PointIgnore:"true"`
	MsgContentsKoKr valueTypes.String   `json:"msg_contents_ko_kr" PointIgnore:"true"`
	MsgContentsNlNl valueTypes.String   `json:"msg_contents_nl_nl" PointIgnore:"true"`
	MsgContentsPlPl valueTypes.String   `json:"msg_contents_pl_pl" PointIgnore:"true"`
	MsgContentsPtBr valueTypes.String   `json:"msg_contents_pt_br" PointIgnore:"true"`
	MsgContentsPtPt valueTypes.String   `json:"msg_contents_pt_pt" PointIgnore:"true"`
	MsgContentsTrTr valueTypes.String   `json:"msg_contents_tr_tr" PointIgnore:"true"`
	MsgContentsViVn valueTypes.String   `json:"msg_contents_vi_vn" PointIgnore:"true"`
	MsgContentsZhTw valueTypes.String   `json:"msg_contents_zh_tw" PointIgnore:"true"`

	MsgTitle        valueTypes.String   `json:"msg_title" PointIgnore:"true"`
	MsgTitleDeDe    valueTypes.String   `json:"msg_title_de_de" PointIgnore:"true"`
	MsgTitleEsEs    valueTypes.String   `json:"msg_title_es_es" PointIgnore:"true"`
	MsgTitleFrFr    valueTypes.String   `json:"msg_title_fr_fr" PointIgnore:"true"`
	MsgTitleIDID    valueTypes.String   `json:"msg_title_id_id" PointIgnore:"true"`
	MsgTitleItIt    valueTypes.String   `json:"msg_title_it_it" PointIgnore:"true"`
	MsgTitleJaJp    valueTypes.String   `json:"msg_title_ja_jp" PointIgnore:"true"`
	MsgTitleKoKr    valueTypes.String   `json:"msg_title_ko_kr" PointIgnore:"true"`
	MsgTitleNlNl    valueTypes.String   `json:"msg_title_nl_nl" PointIgnore:"true"`
	MsgTitlePlPl    valueTypes.String   `json:"msg_title_pl_pl" PointIgnore:"true"`
	MsgTitlePtBr    valueTypes.String   `json:"msg_title_pt_br" PointIgnore:"true"`
	MsgTitlePtPt    valueTypes.String   `json:"msg_title_pt_pt" PointIgnore:"true"`
	MsgTitleTrTr    valueTypes.String   `json:"msg_title_tr_tr" PointIgnore:"true"`
	MsgTitleViVn    valueTypes.String   `json:"msg_title_vi_vn" PointIgnore:"true"`
	MsgTitleZhTw    valueTypes.String   `json:"msg_title_zh_tw" PointIgnore:"true"`
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
