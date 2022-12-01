package getRemoteParamSettingList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getRemoteParamSettingList"
const Disabled = false

type RequestData struct {
	CurPage    valueTypes.Integer `json:"curPage" required:"true"`
	Size       valueTypes.Integer `json:"size" required:"true"`
	DeviceType valueTypes.String  `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData   struct {
	PageList []struct {
		GoStruct               GoStruct.GoStruct  `json:"-" PointIdReplace:"true" PointIdFrom:"PsId"`

		PsId                   valueTypes.Integer `json:"ps_id"`
		PsName                 valueTypes.String  `json:"ps_name"`
		PsShortName            valueTypes.String  `json:"ps_short_name"`

		AapProtocolNum         valueTypes.String  `json:"aap_protocol_num"`
		AppProtocolNum         valueTypes.String  `json:"app_protocol_num"`
		CommandStatus          valueTypes.Integer `json:"command_status"`
		CommunicationDevSn     valueTypes.String  `json:"communication_dev_sn"`
		CommunicationDevUUID   valueTypes.Integer `json:"communication_dev_uuid"`
		CountryId              valueTypes.Integer `json:"country_id"`
		CountryName            valueTypes.String  `json:"country_name"`
		DataFlagDetail         valueTypes.Integer `json:"data_flag_detail"`
		DeviceArea             valueTypes.String  `json:"device_area"`
		DeviceCode             valueTypes.Integer `json:"device_code"`
		DeviceModel            valueTypes.String  `json:"device_model"`
		DeviceModelId          valueTypes.Integer `json:"device_model_id"`
		DeviceName             valueTypes.String  `json:"device_name"`
		DeviceProSn            valueTypes.String  `json:"device_pro_sn"`
		DeviceType             valueTypes.Integer `json:"device_type"`
		GridCompanyName        valueTypes.String  `json:"grid_company_name"`
		GridId                 interface{}        `json:"grid_id"`
		GridTypeId             valueTypes.Integer `json:"grid_type_id"`
		GridTypeName           valueTypes.String  `json:"grid_type_name"`
		GridTypeOriginName     valueTypes.String  `json:"grid_type_origin_name"`
		IacProtocolNum         valueTypes.String  `json:"iac_protocol_num"`
		InverterModelType      valueTypes.Integer `json:"inverter_model_type"`
		ParamSetCountryId      valueTypes.Integer `json:"param_set_country_id"`
		LoggerCode             valueTypes.Integer `json:"logger_code"`
		ShareType              valueTypes.String  `json:"share_type"`
		Sn                     valueTypes.String  `json:"sn"`
		RatedPower             valueTypes.String  `json:"rated_power"`
		UUID                   valueTypes.Integer `json:"uuid"`

		IsCountryCheck         valueTypes.Bool `json:"is_country_check"`
		IsHaveVersion          valueTypes.Bool `json:"is_haveversion"`
		IsInit                 valueTypes.Bool `json:"is_init"`
		IsNeedModbusAddress    valueTypes.Bool  `json:"is_need_modbus_address"`
		IsReadSet              valueTypes.Bool `json:"is_read_set"`
		IsReset                valueTypes.Bool `json:"is_reset"`
		IsSupportParamSet      valueTypes.Bool  `json:"is_support_param_set"`
		IsSupportRemoteUpgrade valueTypes.Bool  `json:"is_support_remote_upgrade"`
		IsThirdParty           valueTypes.Bool `json:"is_third_party"`
		IsUploadedSelfDesc     valueTypes.Bool `json:"is_uploaded_self_desc"`
		IsUseIacAppProtocol    valueTypes.Bool `json:"is_use_iac_app_protocol"`

		AapProtocolVersion     valueTypes.String  `json:"aap_protocol_version"`
		AppProtocolVersion     valueTypes.String  `json:"app_protocol_version"`
		ArmVersion             valueTypes.String  `json:"arm_version"`
		BatVersion             valueTypes.String  `json:"bat_version"`
		CpldVersion            valueTypes.String  `json:"cpld_version"`
		IacProtocolVersion     valueTypes.String  `json:"iac_protocol_version"`
		LcdVersion             valueTypes.String  `json:"lcd_version"`
		MVersion               valueTypes.String  `json:"m_version"`
		MachineVersion         valueTypes.String  `json:"machine_version"`
		MasterNodeVersion      valueTypes.String  `json:"master_node_version"`
		McuVersion             valueTypes.String  `json:"mcu_version"`
		MdspVersion            valueTypes.String  `json:"mdsp_version"`
		PvdVersion             valueTypes.String  `json:"pvd_version"`
		SdspVersion            valueTypes.String  `json:"sdsp_version"`
		TempVersion            valueTypes.String  `json:"temp_version"`
		UpgradeVersion         valueTypes.String  `json:"upgrade_version"`
		Version                valueTypes.String  `json:"version"`
		Version1               valueTypes.String  `json:"version1"`
		Version2               valueTypes.String  `json:"version2"`
		Version3               valueTypes.String  `json:"version3"`
		Version4               valueTypes.String  `json:"version4"`
		Version5               valueTypes.String  `json:"version5"`
		Version6               valueTypes.String  `json:"version6"`
		Version7               valueTypes.String  `json:"version7"`
		Version8               valueTypes.String  `json:"version8"`
		Version9               valueTypes.String  `json:"version9"`
		Version10              valueTypes.String  `json:"version10"`
		Version11              valueTypes.String  `json:"version11"`
		Version12              valueTypes.String  `json:"version12"`
	} `json:"pageList" PointId:"devices" DataTable:"false"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
