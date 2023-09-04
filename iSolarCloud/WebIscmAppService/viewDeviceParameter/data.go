package viewDeviceParameter

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/viewDeviceParameter"
const Disabled = false
const EndPointName = "WebIscmAppService.viewDeviceParameter"

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
	PageDataList []struct {
		AapProtocolNum           valueTypes.String  `json:"aap_protocol_num"`
		AapProtocolVersion       valueTypes.String  `json:"aap_protocol_version"`
		ComType                  valueTypes.String  `json:"com_type"`
		CountryId                valueTypes.Integer `json:"country_id"`
		CountryName              valueTypes.String  `json:"country_name"`
		DeviceModel              valueTypes.String  `json:"device_model"`
		DeviceModelId            valueTypes.String  `json:"device_model_id"`
		GridTypeId               valueTypes.Integer `json:"grid_type_id"`
		IacProtocolNum           valueTypes.String  `json:"iac_protocol_num"`
		IacProtocolVersion       valueTypes.String  `json:"iac_protocol_version"`
		Id                       valueTypes.Integer `json:"id"`
		IsCountryCheck           valueTypes.Bool    `json:"is_country_check"`
		IsReadSet                valueTypes.Bool    `json:"is_read_set"`
		IsReset                  valueTypes.Bool    `json:"is_reset"`
		IsThirdParty             valueTypes.Bool    `json:"is_third_party"`
		IsUseProtocolGetTemplate valueTypes.Bool    `json:"is_use_protocol_get_template"`
		MdspVersion              valueTypes.String  `json:"mdsp_version"`
		Remark                   valueTypes.String  `json:"remark"`
		SdspVersion              valueTypes.String  `json:"sdsp_version"`
		Series                   valueTypes.String  `json:"series"`
		SetId                    valueTypes.Integer `json:"set_id"`
		TemplateType             valueTypes.Integer `json:"template_type"`
		Version                  valueTypes.String  `json:"version"`
		SetName                  valueTypes.String  `json:"set_name"`
	} `json:"pageDataList" PointId:"page_data_list" DataTable:"true" DataTableSort:"CountryName"`
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
