package queryDeviceInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryDeviceInfo"
const Disabled = false

type RequestData struct {
	DeviceSn valueTypes.String `json:"device_sn,omitempty"`
	Uuid     valueTypes.String `json:"uuid,omitempty"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceModelList []struct {
		DeviceModel     valueTypes.String  `json:"device_model"`
		DeviceModelCode valueTypes.String  `json:"device_model_code"`
		ModelId         valueTypes.Integer `json:"model_id"`
	} `json:"deviceModelList" PointId:"device_model_list" PointIdFromChild:"ModelId" PointArrayFlatten:"false" DataTable:"true"`
	DevicePropertyValueList []struct {
		DeviceType     valueTypes.Integer `json:"device_type"`
		PropertyCode   valueTypes.Integer `json:"property_code"`
		PropertyDefVal valueTypes.String  `json:"property_def_val"`
		PropertyName   valueTypes.String  `json:"property_name"`
		PropertyValue  valueTypes.String  `json:"property_value"`
		UUID           valueTypes.String  `json:"uuid"`
	} `json:"devicePropertyValueList" PointId:"device_property_value_list" PointArrayFlatten:"false" DataTable:"true"`

	ArmVersion              interface{}        `json:"arm_version"`
	BatVersion              valueTypes.String        `json:"bat_version"`
	BatteryVersion          interface{}        `json:"battery_version"`
	ChannelId               valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
	CommandStatus           valueTypes.Integer `json:"command_status"`
	CommunicationDevUUID    valueTypes.Integer `json:"communication_dev_uuid"`
	CountryId               valueTypes.Integer        `json:"country_id"`
	CpldVersion             interface{}        `json:"cpld_version"`
	DataFlagDetail          valueTypes.Integer `json:"data_flag_detail"`
	DevFaultStatus          valueTypes.Integer `json:"dev_fault_status"`
	DevStatus               valueTypes.Integer `json:"dev_status"`
	DeviceArea              valueTypes.String  `json:"device_area"`
	DeviceAreaName          valueTypes.String  `json:"device_area_name"`
	DeviceCode              valueTypes.Integer `json:"device_code"`
	DeviceFactoryId         interface{}        `json:"device_factory_id"`
	DeviceModel             valueTypes.String  `json:"device_model"`
	DeviceModelCode         valueTypes.String  `json:"device_model_code"`
	DeviceModelId           valueTypes.Integer `json:"device_model_id"`
	DeviceName              valueTypes.String  `json:"device_name"`
	DeviceProSn             valueTypes.String  `json:"device_pro_sn"`
	DeviceStatus            valueTypes.Integer `json:"device_status"`
	DeviceType              valueTypes.Integer `json:"device_type"`
	FactoryName             valueTypes.String  `json:"factory_name"`
	GridTypeId              valueTypes.Float   `json:"grid_type_id"`
	InstallerDevFaultStatus valueTypes.Float        `json:"installer_dev_fault_status"`
	InverterModelType       valueTypes.Integer `json:"inverter_model_type"`
	IsCountryCheck          valueTypes.Bool    `json:"is_country_check"`
	IsEnable                valueTypes.Bool    `json:"is_enable"`
	IsG2point5Module        valueTypes.Bool    `json:"is_g2point5_module"`
	IsHaveversion           valueTypes.Bool    `json:"is_haveversion" PointId:"is_have_version"`
	IsInit                  valueTypes.Bool    `json:"is_init"`
	IsNeedModbusAddress     valueTypes.Bool    `json:"is_need_modbus_address"`
	IsReadSet               valueTypes.Bool    `json:"is_read_set"`
	IsReplacing             valueTypes.Bool    `json:"is_replacing"`
	IsReset                 valueTypes.Bool    `json:"is_reset"`
	IsSupportParamSet       valueTypes.Bool    `json:"is_support_param_set"`
	IsSupportRemoteUpgrade  valueTypes.Bool    `json:"is_support_remote_upgrade"`
	IsThirdParty            valueTypes.Bool    `json:"is_third_party"`
	Latitude                valueTypes.Float   `json:"latitude"`
	LcdVersion              valueTypes.String        `json:"lcd_version"`
	LoggerCode              valueTypes.Integer `json:"logger_code"`
	Longitude               valueTypes.Float   `json:"longitude"`
	MVersion                valueTypes.String  `json:"m_version"`
	MachineVersion          interface{}        `json:"machine_version"`
	MasterNodeVersion       interface{}        `json:"master_node_version"`
	McuVersion              interface{}        `json:"mcu_version"`
	MdspVersion             valueTypes.String        `json:"mdsp_version"`
	ModelTechList           []interface{}      `json:"modelTechList"`
	NodeTimestamps          interface{}        `json:"node_timestamps"`
	OwnerDevFaultStatus     valueTypes.Integer        `json:"owner_dev_fault_status"`
	P2                      interface{}        `json:"p2"`
	P24                     interface{}        `json:"p24"`
	PsId                    valueTypes.PsId    `json:"ps_id"`
	PsKey                   valueTypes.String  `json:"ps_key"`
	PsName                  valueTypes.String  `json:"ps_name"`
	PsShortName             valueTypes.String  `json:"ps_short_name"`
	PvdVersion              interface{}        `json:"pvd_version"`
	SdspVersion             valueTypes.String        `json:"sdsp_version"`
	Sn                      valueTypes.String  `json:"sn" PointName:"Serial Number"`
	SubTypeList             []interface{}      `json:"subTypeList"`
	TempVersion             interface{}        `json:"temp_version"`
	TypeName                valueTypes.String  `json:"type_name"`
	UpUUID                  valueTypes.Integer `json:"up_uuid"`
	UpgradeVersion          interface{}        `json:"upgrade_version"`
	UUID                    valueTypes.Integer `json:"uuid"`
	Version1                valueTypes.String        `json:"version1"`
	Version10               valueTypes.String        `json:"version10"`
	Version11               valueTypes.String        `json:"version11"`
	Version12               valueTypes.String        `json:"version12"`
	Version2                valueTypes.String        `json:"version2"`
	Version3                valueTypes.String        `json:"version3"`
	Version4                valueTypes.String        `json:"version4"`
	Version5                valueTypes.String        `json:"version5"`
	Version6                valueTypes.String        `json:"version6"`
	Version7                valueTypes.String        `json:"version7"`
	Version8                valueTypes.String        `json:"version8"`
	Version9                valueTypes.String        `json:"version9"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Response.ResultData.PsId.String() + "." + e.Response.ResultData.PsKey.String()
		entries.StructToDataMap(*e,  e.Response.ResultData.PsId.String(), GoStruct.NewEndPointPath(e.Response.ResultData.PsId.String()))
	}

	return entries
}
