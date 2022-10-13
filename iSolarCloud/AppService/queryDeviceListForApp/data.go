package queryDeviceListForApp

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const Url = "/v1/devService/queryDeviceListForApp"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		AttrID                  api.Integer  `json:"attr_id"`
		ChannelId               api.Integer  `json:"chnnl_id" PointId:"channel_id"`
		CommandStatus           api.Integer  `json:"command_status"`
		ConnectState            api.Integer  `json:"connect_state"`
		DataFlag                api.Integer  `json:"data_flag"`
		DataFlagDetail          api.Integer  `json:"data_flag_detail"`
		DevFaultStatus          string       `json:"dev_fault_status"`
		DevStatus               string       `json:"dev_status"`
		DeviceArea              api.String   `json:"device_area"`
		DeviceCode              api.Integer  `json:"device_code"`
		DeviceFactoryDate       api.DateTime `json:"device_factory_date"`
		DeviceID                api.Integer  `json:"device_id"`
		DeviceModel             api.String   `json:"device_model"`
		DeviceModelCode         api.String   `json:"device_model_code"`
		DeviceModelID           api.Integer  `json:"device_model_id"`
		DeviceName              api.String   `json:"device_name"`
		DeviceProSn             api.String   `json:"device_pro_sn"`
		DeviceState             string       `json:"device_state"`
		DeviceSubType           interface{}  `json:"device_sub_type"`
		DeviceSubTypeName       interface{}  `json:"device_sub_type_name"`
		DeviceType              api.Integer  `json:"device_type"`
		FactoryName             api.String   `json:"factory_name"`
		InstallerDevFaultStatus string       `json:"installer_dev_fault_status"`
		InverterModelType       api.Integer  `json:"inverter_model_type"`
		IsCountryCheck          api.Bool     `json:"is_country_check"`
		IsHasFunctionEnum       api.Bool     `json:"is_has_function_enum"`
		IsHasTheAbility         api.Bool     `json:"is_has_the_ability"`
		IsInit                  api.Bool     `json:"is_init"`
		IsReadSet               api.Bool     `json:"is_read_set"`
		IsReplacing             api.Bool     `json:"is_replacing"`
		IsReset                 api.Bool     `json:"is_reset"`
		IsSecond                api.Bool     `json:"is_second"`
		IsThirdParty            api.Bool     `json:"is_third_party"`
		ModuleUUID              api.Integer  `json:"module_uuid"`
		OwnerDevFaultStatus     string       `json:"owner_dev_fault_status"`
		P24                     interface{}  `json:"p24"`
		Posx                    interface{}  `json:"posx"`
		Posy                    interface{}  `json:"posy"`
		PsID                    api.Integer  `json:"ps_id"`
		PsKey                   api.PsKey    `json:"ps_key"`
		RelState                api.Integer  `json:"rel_state"`
		Sn                      api.String   `json:"sn" PointName:"Serial Number"`
		TypeName                api.String   `json:"type_name"`
		UUID                    api.Integer  `json:"uuid"`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
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
		pkg := apiReflect.GetName("", *e)
		for _, d := range e.Response.ResultData.PageList {
			name := strings.Join([]string{pkg, d.PsKey.Value()}, ".")
			entries.StructToPoints(d, name, d.PsKey.Value(), api.NewDateTime(""))
		}
	}

	return entries
}
