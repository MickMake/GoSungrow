package queryDeviceInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/devService/queryDeviceInfo"
const Disabled = false

type RequestData struct {
	DeviceSn string `json:"device_sn,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ArmVersion           interface{} `json:"arm_version"`
	BatVersion           interface{} `json:"bat_version"`
	BatteryVersion       interface{} `json:"battery_version"`
	ChannelId            api.Integer `json:"chnnl_id"`
	CommandStatus        api.Integer `json:"command_status"`
	CommunicationDevUUID api.Integer `json:"communication_dev_uuid"`
	CountryID            interface{} `json:"country_id"`
	CpldVersion          interface{} `json:"cpld_version"`
	DataFlagDetail       api.Integer `json:"data_flag_detail"`
	DevFaultStatus       string      `json:"dev_fault_status"`
	DevStatus            string      `json:"dev_status"`
	DeviceModelList      []struct {
		DeviceModel     api.String  `json:"device_model"`
		DeviceModelCode api.String  `json:"device_model_code"`
		ModelID         api.Integer `json:"model_id"`
	} `json:"deviceModelList"`
	DevicePropertyValueList []struct {
		DeviceType     api.Integer `json:"device_type"`
		PropertyCode   api.Integer `json:"property_code"`
		PropertyDefVal api.String  `json:"property_def_val"`
		PropertyName   api.String  `json:"property_name"`
		PropertyValue  interface{} `json:"property_value"`
		UUID           api.String  `json:"uuid"`
	} `json:"devicePropertyValueList"`
	DeviceArea              api.String    `json:"device_area"`
	DeviceAreaName          api.String    `json:"device_area_name"`
	DeviceCode              api.Integer   `json:"device_code"`
	DeviceFactoryID         interface{}   `json:"device_factory_id"`
	DeviceModel             api.String    `json:"device_model"`
	DeviceModelCode         api.String    `json:"device_model_code"`
	DeviceModelID           api.Integer   `json:"device_model_id"`
	DeviceName              api.String    `json:"device_name"`
	DeviceProSn             api.String    `json:"device_pro_sn"`
	DeviceStatus            int64         `json:"device_status"`
	DeviceType              int64         `json:"device_type"`
	FactoryName             api.String    `json:"factory_name"`
	GridTypeID              interface{}   `json:"grid_type_id"`
	InstallerDevFaultStatus interface{}   `json:"installer_dev_fault_status"`
	InverterModelType       int64         `json:"inverter_model_type"`
	IsCountryCheck          api.Bool      `json:"is_country_check"`
	IsEnable                api.Bool      `json:"is_enable"`
	IsG2point5Module        api.Bool      `json:"is_g2point5_module"`
	IsHaveversion           api.Bool      `json:"is_haveversion"`
	IsInit                  api.Bool      `json:"is_init"`
	IsNeedModbusAddress     api.Bool      `json:"is_need_modbus_address"`
	IsReadSet               api.Bool      `json:"is_read_set"`
	IsReplacing             api.Bool      `json:"is_replacing"`
	IsReset                 api.Bool      `json:"is_reset"`
	IsSupportParamSet       api.Bool      `json:"is_support_param_set"`
	IsSupportRemoteUpgrade  api.Bool      `json:"is_support_remote_upgrade"`
	IsThirdParty            api.Bool      `json:"is_third_party"`
	Latitude                api.Float     `json:"latitude"`
	LcdVersion              interface{}   `json:"lcd_version"`
	LoggerCode              api.Integer   `json:"logger_code"`
	Longitude               api.Float     `json:"longitude"`
	MVersion                api.String    `json:"m_version"`
	MachineVersion          interface{}   `json:"machine_version"`
	MasterNodeVersion       interface{}   `json:"master_node_version"`
	McuVersion              interface{}   `json:"mcu_version"`
	MdspVersion             interface{}   `json:"mdsp_version"`
	ModelTechList           []interface{} `json:"modelTechList"`
	NodeTimestamps          interface{}   `json:"node_timestamps"`
	OwnerDevFaultStatus     interface{}   `json:"owner_dev_fault_status"`
	P2                      interface{}   `json:"p2"`
	P24                     interface{}   `json:"p24"`
	PsID                    api.Integer   `json:"ps_id"`
	PsKey                   api.String    `json:"ps_key"`
	PsName                  api.String    `json:"ps_name"`
	PsShortName             api.String    `json:"ps_short_name"`
	PvdVersion              interface{}   `json:"pvd_version"`
	SdspVersion             interface{}   `json:"sdsp_version"`
	Sn                      api.String    `json:"sn"`
	SubTypeList             []interface{} `json:"subTypeList"`
	TempVersion             interface{}   `json:"temp_version"`
	TypeName                api.String    `json:"type_name"`
	UpUUID                  api.Integer   `json:"up_uuid"`
	UpgradeVersion          interface{}   `json:"upgrade_version"`
	UUID                    api.Integer   `json:"uuid"`
	Version1                interface{}   `json:"version1"`
	Version10               interface{}   `json:"version10"`
	Version11               interface{}   `json:"version11"`
	Version12               interface{}   `json:"version12"`
	Version2                interface{}   `json:"version2"`
	Version3                interface{}   `json:"version3"`
	Version4                interface{}   `json:"version4"`
	Version5                interface{}   `json:"version5"`
	Version6                interface{}   `json:"version6"`
	Version7                interface{}   `json:"version7"`
	Version8                interface{}   `json:"version8"`
	Version9                interface{}   `json:"version9"`
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
