package getCloudList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/commonService/getCloudList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CloudList []struct {
		CloudId    valueTypes.Integer  `json:"cloud_id"`
		CloudName  valueTypes.String `json:"cloud_name"`
		GatewayURL valueTypes.String `json:"gateway_url"`
		OrderId    valueTypes.Integer  `json:"order_id"`
		ServiceURL valueTypes.String `json:"service_url"`
		Value      valueTypes.String `json:"value"`
		ValueDeDe  valueTypes.String `json:"value_de_de"`
		ValueEnUs  valueTypes.String `json:"value_en_us"`
		ValueEsEs  valueTypes.String `json:"value_es_es"`
		ValueFrFr  valueTypes.String `json:"value_fr_fr"`
		ValueItIt  valueTypes.String `json:"value_it_it"`
		ValueJaJp  valueTypes.String `json:"value_ja_jp"`
		ValueKoKr  valueTypes.String `json:"value_ko_kr"`
		ValueNlNl  valueTypes.String `json:"value_nl_nl"`
		ValuePlPl  valueTypes.String `json:"value_pl_pl"`
		ValuePtBr  valueTypes.String `json:"value_pt_br"`
		ValuePtPt  valueTypes.String `json:"value_pt_pt"`
		ValueTrTr  valueTypes.String `json:"value_tr_tr"`
		ValueViVn  valueTypes.String `json:"value_vi_vn"`
		ValueZhCn  valueTypes.String `json:"value_zh_cn"`
		ValueZhTw  valueTypes.String `json:"value_zh_tw"`
		WebURL     valueTypes.String `json:"web_url"`
	} `json:"cloud_list" DataTable:"true"`
	CurrentCloud struct {
		CloudId    valueTypes.Integer  `json:"cloud_id"`
		CloudName  valueTypes.String `json:"cloud_name"`
		GatewayURL valueTypes.String `json:"gateway_url"`
		OrderId    valueTypes.Integer  `json:"order_id"`
		ServiceURL valueTypes.String `json:"service_url"`
		Value      valueTypes.String `json:"value"`
		ValueDeDe  valueTypes.String `json:"value_de_de"`
		ValueEnUs  valueTypes.String `json:"value_en_us"`
		ValueEsEs  valueTypes.String `json:"value_es_es"`
		ValueFrFr  valueTypes.String `json:"value_fr_fr"`
		ValueItIt  valueTypes.String `json:"value_it_it"`
		ValueJaJp  valueTypes.String `json:"value_ja_jp"`
		ValueKoKr  valueTypes.String `json:"value_ko_kr"`
		ValueNlNl  valueTypes.String `json:"value_nl_nl"`
		ValuePlPl  valueTypes.String `json:"value_pl_pl"`
		ValuePtBr  valueTypes.String `json:"value_pt_br"`
		ValuePtPt  valueTypes.String `json:"value_pt_pt"`
		ValueTrTr  valueTypes.String `json:"value_tr_tr"`
		ValueViVn  valueTypes.String `json:"value_vi_vn"`
		ValueZhCn  valueTypes.String `json:"value_zh_cn"`
		ValueZhTw  valueTypes.String `json:"value_zh_tw"`
		WebURL     valueTypes.String `json:"web_url"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
