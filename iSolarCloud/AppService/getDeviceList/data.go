package getDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceList"
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
		AttrID                  api.Integer `json:"attr_id"`
		ChannelId               api.Integer `json:"chnnl_id"`
		CommandStatus           api.Integer `json:"command_status"`
		ConnectState            api.Integer `json:"connect_state"`
		DataFlag                api.Integer `json:"data_flag"`
		DataFlagDetail          api.Integer `json:"data_flag_detail"`
		DevFaultStatus          string      `json:"dev_fault_status"`
		DevStatus               string      `json:"dev_status"`
		DeviceArea              string      `json:"device_area"`
		DeviceCode              api.Integer `json:"device_code"`
		DeviceFactoryDate       interface{} `json:"device_factory_date"`
		DeviceID                api.Integer `json:"device_id"`
		DeviceModel             string      `json:"device_model"`
		DeviceModelCode         string      `json:"device_model_code"`
		DeviceModelID           api.Integer `json:"device_model_id"`
		DeviceName              string      `json:"device_name"`
		DeviceProSn             string      `json:"device_pro_sn"`
		DeviceState             string      `json:"device_state"`
		DeviceSubType           interface{} `json:"device_sub_type"`
		DeviceSubTypeName       interface{} `json:"device_sub_type_name"`
		DeviceType              api.Integer `json:"device_type"`
		FactoryName             string      `json:"factory_name"`
		InstallerDevFaultStatus string      `json:"installer_dev_fault_status"`
		InverterModelType       api.Integer `json:"inverter_model_type"`
		IsCountryCheck          api.Integer `json:"is_country_check"`
		IsHasFunctionEnum       api.Integer `json:"is_has_function_enum"`
		IsHasTheAbility         api.Integer `json:"is_has_the_ability"`
		IsInit                  api.Integer `json:"is_init"`
		IsReadSet               api.Integer `json:"is_read_set"`
		IsReplacing             api.Integer `json:"is_replacing"`
		IsReset                 api.Integer `json:"is_reset"`
		IsSecond                api.Integer `json:"is_second"`
		IsThirdParty            api.Integer `json:"is_third_party"`
		ModuleUUID              api.Integer `json:"module_uuid"`
		OwnerDevFaultStatus     string      `json:"owner_dev_fault_status"`
		P24                     interface{} `json:"p24"`
		Posx                    interface{} `json:"posx"`
		Posy                    interface{} `json:"posy"`
		PsID                    api.Integer `json:"ps_id"`
		PsKey                   string      `json:"ps_key"`
		RelState                api.Integer `json:"rel_state"`
		Sn                      string      `json:"sn"`
		TypeName                string      `json:"type_name"`
		UUID                    api.Integer `json:"uuid"`
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Ps Key",
			"Ps Id",
			"Type",
			"Code",
			"Id",
			"Type Name",
			"Serial Number",
			"Model",
			"Model Id",
			"Name",
			"State",
			"Status",
			// "Factory Date",
		)
		for _, d := range e.Response.ResultData.PageList {
			_ = table.AddRow(
				d.PsKey,
				d.PsID,
				d.DeviceType,
				d.DeviceCode,
				d.ChannelId,
				d.TypeName,
				d.DeviceProSn,
				d.DeviceModel,
				d.DeviceModelID,
				d.DeviceName,
				d.DeviceState,
				d.DevStatus,
				// d.DeviceFactoryDate,
			)
		}
	}
	return table
}
