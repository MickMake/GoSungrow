package getCloudList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/commonService/getCloudList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CloudList []struct {
		CloudID    api.Integer  `json:"cloud_id"`
		CloudName  string `json:"cloud_name"`
		GatewayURL string `json:"gateway_url"`
		OrderID    api.Integer  `json:"order_id"`
		ServiceURL string `json:"service_url"`
		Value      string `json:"value"`
		ValueDeDe  string `json:"value_de_de"`
		ValueEnUs  string `json:"value_en_us"`
		ValueEsEs  string `json:"value_es_es"`
		ValueFrFr  string `json:"value_fr_fr"`
		ValueItIt  string `json:"value_it_it"`
		ValueJaJp  string `json:"value_ja_jp"`
		ValueKoKr  string `json:"value_ko_kr"`
		ValueNlNl  string `json:"value_nl_nl"`
		ValuePlPl  string `json:"value_pl_pl"`
		ValuePtBr  string `json:"value_pt_br"`
		ValuePtPt  string `json:"value_pt_pt"`
		ValueTrTr  string `json:"value_tr_tr"`
		ValueViVn  string `json:"value_vi_vn"`
		ValueZhCn  string `json:"value_zh_cn"`
		ValueZhTw  string `json:"value_zh_tw"`
		WebURL     string `json:"web_url"`
	} `json:"cloud_list"`
	CurrentCloud struct {
		CloudID    api.Integer  `json:"cloud_id"`
		CloudName  string `json:"cloud_name"`
		GatewayURL string `json:"gateway_url"`
		OrderID    api.Integer  `json:"order_id"`
		ServiceURL string `json:"service_url"`
		Value      string `json:"value"`
		ValueDeDe  string `json:"value_de_de"`
		ValueEnUs  string `json:"value_en_us"`
		ValueEsEs  string `json:"value_es_es"`
		ValueFrFr  string `json:"value_fr_fr"`
		ValueItIt  string `json:"value_it_it"`
		ValueJaJp  string `json:"value_ja_jp"`
		ValueKoKr  string `json:"value_ko_kr"`
		ValueNlNl  string `json:"value_nl_nl"`
		ValuePlPl  string `json:"value_pl_pl"`
		ValuePtBr  string `json:"value_pt_br"`
		ValuePtPt  string `json:"value_pt_pt"`
		ValueTrTr  string `json:"value_tr_tr"`
		ValueViVn  string `json:"value_vi_vn"`
		ValueZhCn  string `json:"value_zh_cn"`
		ValueZhTw  string `json:"value_zh_tw"`
		WebURL     string `json:"web_url"`
	} `json:"current_cloud"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
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
