package viewDeviceParameter

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/viewDeviceParameter"
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


type ResultData struct {
	PageDataList []struct {
		AapProtocolNum           valueTypes.String  `json:"aap_protocol_num"`
		AapProtocolVersion       valueTypes.String  `json:"aap_protocol_version"`
		ComType                  valueTypes.String  `json:"com_type"`
		CountryID                valueTypes.Integer `json:"country_id"`
		CountryName              valueTypes.String  `json:"country_name"`
		DeviceModel              valueTypes.String  `json:"device_model"`
		DeviceModelID            valueTypes.String  `json:"device_model_id"`
		GridTypeID               valueTypes.Integer `json:"grid_type_id"`
		IacProtocolNum           valueTypes.String  `json:"iac_protocol_num"`
		IacProtocolVersion       valueTypes.String  `json:"iac_protocol_version"`
		ID                       valueTypes.Integer `json:"id"`
		IsCountryCheck           valueTypes.Integer `json:"is_country_check"`
		IsReadSet                valueTypes.Integer `json:"is_read_set"`
		IsReset                  valueTypes.Integer `json:"is_reset"`
		IsThirdParty             valueTypes.Integer `json:"is_third_party"`
		IsUseProtocolGetTemplate valueTypes.Integer `json:"is_use_protocol_get_template"`
		MdspVersion              valueTypes.String  `json:"mdsp_version"`
		Remark                   valueTypes.String  `json:"remark"`
		SdspVersion              valueTypes.String  `json:"sdsp_version"`
		Series                   valueTypes.String  `json:"series"`
		SetID                    valueTypes.Integer `json:"set_id"`
		SetName                  valueTypes.String  `json:"set_name"`
		TemplateType             valueTypes.Integer `json:"template_type"`
		Version                  valueTypes.String  `json:"version"`
	} `json:"pageDataList" PointId:"page_data_list" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
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
		entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
