package getDeviceList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceList"
const Disabled = false

type RequestData struct {
	PsId int64 `json:"ps_id" required:"true"`
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
		AttrID                  int64       `json:"attr_id"`
		ChnnlID                 int64       `json:"chnnl_id"`
		CommandStatus           int64       `json:"command_status"`
		ConnectState            int64       `json:"connect_state"`
		DataFlag                int64       `json:"data_flag"`
		DataFlagDetail          int64       `json:"data_flag_detail"`
		DevFaultStatus          string      `json:"dev_fault_status"`
		DevStatus               string      `json:"dev_status"`
		DeviceArea              string      `json:"device_area"`
		DeviceCode              int64       `json:"device_code"`
		DeviceFactoryDate       interface{} `json:"device_factory_date"`
		DeviceID                int64       `json:"device_id"`
		DeviceModel             string      `json:"device_model"`
		DeviceModelCode         string      `json:"device_model_code"`
		DeviceModelID           int64       `json:"device_model_id"`
		DeviceName              string      `json:"device_name"`
		DeviceProSn             string      `json:"device_pro_sn"`
		DeviceState             string      `json:"device_state"`
		DeviceSubType           interface{} `json:"device_sub_type"`
		DeviceSubTypeName       interface{} `json:"device_sub_type_name"`
		DeviceType              int64       `json:"device_type"`
		FactoryName             string      `json:"factory_name"`
		InstallerDevFaultStatus string      `json:"installer_dev_fault_status"`
		InverterModelType       int64       `json:"inverter_model_type"`
		IsCountryCheck          int64       `json:"is_country_check"`
		IsHasFunctionEnum       int64       `json:"is_has_function_enum"`
		IsHasTheAbility         int64       `json:"is_has_the_ability"`
		IsInit                  int64       `json:"is_init"`
		IsReadSet               int64       `json:"is_read_set"`
		IsReplacing             int64       `json:"is_replacing"`
		IsReset                 int64       `json:"is_reset"`
		IsSecond                int64       `json:"is_second"`
		IsThirdParty            int64       `json:"is_third_party"`
		ModuleUUID              int64       `json:"module_uuid"`
		OwnerDevFaultStatus     string      `json:"owner_dev_fault_status"`
		P24                     interface{} `json:"p24"`
		Posx                    interface{} `json:"posx"`
		Posy                    interface{} `json:"posy"`
		PsID                    int64       `json:"ps_id"`
		PsKey                   string      `json:"ps_key"`
		RelState                int64       `json:"rel_state"`
		Sn                      string      `json:"sn"`
		TypeName                string      `json:"type_name"`
		UUID                    int64       `json:"uuid"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
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
				d.ChnnlID,
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
