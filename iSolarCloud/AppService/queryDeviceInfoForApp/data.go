package queryDeviceInfoForApp

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/devService/queryDeviceInfoForApp"
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
	ChannelId            valueTypes.Integer `json:"chnnl_id"`
	CommandStatus        valueTypes.Integer `json:"command_status"`
	CommunicationDevUUID valueTypes.Integer `json:"communication_dev_uuid"`
	CountryId            interface{} `json:"country_id"`
	CpldVersion          interface{} `json:"cpld_version"`
	DataFlagDetail       valueTypes.Integer `json:"data_flag_detail"`
	DevFaultStatus       string      `json:"dev_fault_status"`
	DevStatus            string      `json:"dev_status"`
	DeviceModelList      []struct {
		DeviceModel     valueTypes.String  `json:"device_model"`
		DeviceModelCode valueTypes.String  `json:"device_model_code"`
		ModelId         valueTypes.Integer `json:"model_id"`
	} `json:"deviceModelList"`
	DevicePropertyValueList []struct {
		DeviceType     valueTypes.Integer `json:"device_type"`
		PropertyCode   valueTypes.Integer `json:"property_code"`
		PropertyDefVal valueTypes.String  `json:"property_def_val"`
		PropertyName   valueTypes.String  `json:"property_name"`
		PropertyValue  interface{} `json:"property_value"`
		UUID           valueTypes.String  `json:"uuid"`
	} `json:"devicePropertyValueList"`
	DeviceArea              valueTypes.String    `json:"device_area"`
	DeviceAreaName          valueTypes.String    `json:"device_area_name"`
	DeviceCode              valueTypes.Integer   `json:"device_code"`
	DeviceFactoryId         interface{}   `json:"device_factory_id"`
	DeviceModel             valueTypes.String    `json:"device_model"`
	DeviceModelCode         valueTypes.String    `json:"device_model_code"`
	DeviceModelId           valueTypes.Integer   `json:"device_model_id"`
	DeviceName              valueTypes.String    `json:"device_name"`
	DeviceProSn             valueTypes.String    `json:"device_pro_sn"`
	DeviceStatus            int64         `json:"device_status"`
	DeviceType              int64         `json:"device_type"`
	FactoryName             valueTypes.String    `json:"factory_name"`
	GridTypeId              interface{}   `json:"grid_type_id"`
	InstallerDevFaultStatus interface{}   `json:"installer_dev_fault_status"`
	InverterModelType       int64         `json:"inverter_model_type"`
	IsCountryCheck          valueTypes.Bool      `json:"is_country_check"`
	IsEnable                valueTypes.Bool      `json:"is_enable"`
	IsG2point5Module        valueTypes.Bool      `json:"is_g2point5_module"`
	IsHaveversion           valueTypes.Bool      `json:"is_haveversion"`
	IsInit                  valueTypes.Bool      `json:"is_init"`
	IsNeedModbusAddress     valueTypes.Bool      `json:"is_need_modbus_address"`
	IsReadSet               valueTypes.Bool      `json:"is_read_set"`
	IsReplacing             valueTypes.Bool      `json:"is_replacing"`
	IsReset                 valueTypes.Bool      `json:"is_reset"`
	IsSupportParamSet       valueTypes.Bool      `json:"is_support_param_set"`
	IsSupportRemoteUpgrade  valueTypes.Bool      `json:"is_support_remote_upgrade"`
	IsThirdParty            valueTypes.Bool      `json:"is_third_party"`
	Latitude                valueTypes.Float     `json:"latitude"`
	LcdVersion              interface{}   `json:"lcd_version"`
	LoggerCode              valueTypes.Integer   `json:"logger_code"`
	Longitude               valueTypes.Float     `json:"longitude"`
	MVersion                valueTypes.String    `json:"m_version"`
	MachineVersion          interface{}   `json:"machine_version"`
	MasterNodeVersion       interface{}   `json:"master_node_version"`
	McuVersion              interface{}   `json:"mcu_version"`
	MdspVersion             interface{}   `json:"mdsp_version"`
	ModelTechList           []interface{} `json:"modelTechList"`
	NodeTimestamps          interface{}   `json:"node_timestamps"`
	OwnerDevFaultStatus     interface{}   `json:"owner_dev_fault_status"`
	P2                      interface{}   `json:"p2"`
	P24                     interface{}   `json:"p24"`
	PsId                    valueTypes.PsId   `json:"ps_id"`
	PsKey                   valueTypes.String    `json:"ps_key"`
	PsName                  valueTypes.String    `json:"ps_name"`
	PsShortName             valueTypes.String    `json:"ps_short_name"`
	PvdVersion              interface{}   `json:"pvd_version"`
	SdspVersion             interface{}   `json:"sdsp_version"`
	Sn                      valueTypes.String    `json:"sn"`
	SubTypeList             []interface{} `json:"subTypeList"`
	TempVersion             interface{}   `json:"temp_version"`
	TypeName                valueTypes.String    `json:"type_name"`
	UpUUID                  valueTypes.Integer   `json:"up_uuid"`
	UpgradeVersion          interface{}   `json:"upgrade_version"`
	UUID                    valueTypes.Integer   `json:"uuid"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
