package getSysMsg

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/otherService/getSysMsg"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	AttachNames     interface{}  `json:"attach_names"`
	AttachPath      interface{}  `json:"attach_path"`
	BusinessId      valueTypes.Integer  `json:"business_id"`
	ClickNum        valueTypes.Integer  `json:"click_num"`
	CreateTime      valueTypes.DateTime `json:"create_time"`
	EndTime         valueTypes.DateTime `json:"end_time"`
	Ispublish       valueTypes.Bool     `json:"ispublish"`
	MsgContents     valueTypes.String   `json:"msg_contents"`
	MsgContentsDeDe valueTypes.String   `json:"msg_contents_de_de"`
	MsgContentsEnUs valueTypes.String   `json:"msg_contents_en_us"`
	MsgContentsEsEs valueTypes.String   `json:"msg_contents_es_es"`
	MsgContentsFrFr valueTypes.String   `json:"msg_contents_fr_fr"`
	MsgContentsIDID valueTypes.String   `json:"msg_contents_id_id"`
	MsgContentsItIt valueTypes.String   `json:"msg_contents_it_it"`
	MsgContentsJaJp valueTypes.String   `json:"msg_contents_ja_jp"`
	MsgContentsKoKr valueTypes.String   `json:"msg_contents_ko_kr"`
	MsgContentsNlNl valueTypes.String   `json:"msg_contents_nl_nl"`
	MsgContentsPlPl valueTypes.String   `json:"msg_contents_pl_pl"`
	MsgContentsPtBr valueTypes.String   `json:"msg_contents_pt_br"`
	MsgContentsPtPt valueTypes.String   `json:"msg_contents_pt_pt"`
	MsgContentsTrTr valueTypes.String   `json:"msg_contents_tr_tr"`
	MsgContentsViVn valueTypes.String   `json:"msg_contents_vi_vn"`
	MsgContentsZhTw valueTypes.String   `json:"msg_contents_zh_tw"`
	MsgId           valueTypes.Integer  `json:"msg_id"`
	MsgLevel        valueTypes.Integer  `json:"msg_level"`
	MsgTitle        valueTypes.String   `json:"msg_title"`
	MsgTitleDeDe    valueTypes.String   `json:"msg_title_de_de"`
	MsgTitleEnUs    valueTypes.String   `json:"msg_title_en_us"`
	MsgTitleEsEs    valueTypes.String   `json:"msg_title_es_es"`
	MsgTitleFrFr    valueTypes.String   `json:"msg_title_fr_fr"`
	MsgTitleIDID    valueTypes.String   `json:"msg_title_id_id"`
	MsgTitleItIt    valueTypes.String   `json:"msg_title_it_it"`
	MsgTitleJaJp    valueTypes.String   `json:"msg_title_ja_jp"`
	MsgTitleKoKr    valueTypes.String   `json:"msg_title_ko_kr"`
	MsgTitleNlNl    valueTypes.String   `json:"msg_title_nl_nl"`
	MsgTitlePlPl    valueTypes.String   `json:"msg_title_pl_pl"`
	MsgTitlePtBr    valueTypes.String   `json:"msg_title_pt_br"`
	MsgTitlePtPt    valueTypes.String   `json:"msg_title_pt_pt"`
	MsgTitleTrTr    valueTypes.String   `json:"msg_title_tr_tr"`
	MsgTitleViVn    valueTypes.String   `json:"msg_title_vi_vn"`
	MsgTitleZhTw    valueTypes.String   `json:"msg_title_zh_tw"`
	MsgUserIds      valueTypes.String   `json:"msg_user_ids"`
	OperateType     valueTypes.Integer  `json:"operate_type"`
	OperateURL      valueTypes.String   `json:"operate_url"`
	PublishTime     valueTypes.DateTime `json:"publish_time"`
	RemindType      valueTypes.Integer  `json:"remind_type"`
	StartTime       valueTypes.DateTime `json:"start_time"`
	TypeId          valueTypes.Integer  `json:"type_id"`
	UserId          valueTypes.Integer  `json:"user_id"`
	UserName        valueTypes.String   `json:"user_name"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
