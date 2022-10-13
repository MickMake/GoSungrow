package getSysMsg

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/otherService/getSysMsg"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	AttachNames     interface{}  `json:"attach_names"`
	AttachPath      interface{}  `json:"attach_path"`
	BusinessID      api.Integer  `json:"business_id"`
	ClickNum        api.Integer  `json:"click_num"`
	CreateTime      api.DateTime `json:"create_time"`
	EndTime         api.DateTime `json:"end_time"`
	Ispublish       api.Bool     `json:"ispublish"`
	MsgContents     api.String   `json:"msg_contents"`
	MsgContentsDeDe api.String   `json:"msg_contents_de_de"`
	MsgContentsEnUs api.String   `json:"msg_contents_en_us"`
	MsgContentsEsEs api.String   `json:"msg_contents_es_es"`
	MsgContentsFrFr api.String   `json:"msg_contents_fr_fr"`
	MsgContentsIDID api.String   `json:"msg_contents_id_id"`
	MsgContentsItIt api.String   `json:"msg_contents_it_it"`
	MsgContentsJaJp api.String   `json:"msg_contents_ja_jp"`
	MsgContentsKoKr api.String   `json:"msg_contents_ko_kr"`
	MsgContentsNlNl api.String   `json:"msg_contents_nl_nl"`
	MsgContentsPlPl api.String   `json:"msg_contents_pl_pl"`
	MsgContentsPtBr api.String   `json:"msg_contents_pt_br"`
	MsgContentsPtPt api.String   `json:"msg_contents_pt_pt"`
	MsgContentsTrTr api.String   `json:"msg_contents_tr_tr"`
	MsgContentsViVn api.String   `json:"msg_contents_vi_vn"`
	MsgContentsZhTw api.String   `json:"msg_contents_zh_tw"`
	MsgID           api.Integer  `json:"msg_id"`
	MsgLevel        api.Integer  `json:"msg_level"`
	MsgTitle        api.String   `json:"msg_title"`
	MsgTitleDeDe    api.String   `json:"msg_title_de_de"`
	MsgTitleEnUs    api.String   `json:"msg_title_en_us"`
	MsgTitleEsEs    api.String   `json:"msg_title_es_es"`
	MsgTitleFrFr    api.String   `json:"msg_title_fr_fr"`
	MsgTitleIDID    api.String   `json:"msg_title_id_id"`
	MsgTitleItIt    api.String   `json:"msg_title_it_it"`
	MsgTitleJaJp    api.String   `json:"msg_title_ja_jp"`
	MsgTitleKoKr    api.String   `json:"msg_title_ko_kr"`
	MsgTitleNlNl    api.String   `json:"msg_title_nl_nl"`
	MsgTitlePlPl    api.String   `json:"msg_title_pl_pl"`
	MsgTitlePtBr    api.String   `json:"msg_title_pt_br"`
	MsgTitlePtPt    api.String   `json:"msg_title_pt_pt"`
	MsgTitleTrTr    api.String   `json:"msg_title_tr_tr"`
	MsgTitleViVn    api.String   `json:"msg_title_vi_vn"`
	MsgTitleZhTw    api.String   `json:"msg_title_zh_tw"`
	MsgUserIds      api.String   `json:"msg_user_ids"`
	OperateType     api.Integer  `json:"operate_type"`
	OperateURL      api.String   `json:"operate_url"`
	PublishTime     api.DateTime `json:"publish_time"`
	RemindType      api.Integer  `json:"remind_type"`
	StartTime       api.DateTime `json:"start_time"`
	TypeID          api.Integer  `json:"type_id"`
	UserID          api.Integer  `json:"user_id"`
	UserName        api.String   `json:"user_name"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}
