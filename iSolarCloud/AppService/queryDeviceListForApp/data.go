package queryDeviceListForApp

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryDeviceListForApp"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		GoStruct                GoStruct.GoStruct   `json:"-" PointDeviceFrom:"PsKey"`

		AttrId                  valueTypes.Integer  `json:"attr_id"`
		ChannelId               valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		CommandStatus           valueTypes.Integer  `json:"command_status"`
		ConnectState            valueTypes.Integer  `json:"connect_state"`
		DataFlag                valueTypes.Integer  `json:"data_flag"`
		DataFlagDetail          valueTypes.Integer  `json:"data_flag_detail"`
		DevFaultStatus          valueTypes.Integer  `json:"dev_fault_status"`
		DevStatus               valueTypes.Bool     `json:"dev_status"`
		DeviceArea              valueTypes.String   `json:"device_area"`
		DeviceCode              valueTypes.Integer  `json:"device_code"`
		DeviceFactoryDate       valueTypes.DateTime `json:"device_factory_date" PointNameDateFormat:"2006/01/02 15:04:05"`
		DeviceId                valueTypes.Integer  `json:"device_id"`
		DeviceModel             valueTypes.String   `json:"device_model"`
		DeviceModelCode         valueTypes.String   `json:"device_model_code"`
		DeviceModelId           valueTypes.Integer  `json:"device_model_id"`
		DeviceName              valueTypes.String   `json:"device_name"`
		DeviceProSn             valueTypes.String   `json:"device_pro_sn"`
		DeviceState             valueTypes.Integer  `json:"device_state"`
		DeviceSubType           interface{}         `json:"device_sub_type"`
		DeviceSubTypeName       interface{}         `json:"device_sub_type_name"`
		DeviceType              valueTypes.Integer  `json:"device_type"`
		FactoryName             valueTypes.String   `json:"factory_name"`
		InstallerDevFaultStatus valueTypes.Integer  `json:"installer_dev_fault_status"`
		InverterModelType       valueTypes.Integer  `json:"inverter_model_type"`
		IsCountryCheck          valueTypes.Bool     `json:"is_country_check"`
		IsHasFunctionEnum       valueTypes.Bool     `json:"is_has_function_enum"`
		IsHasTheAbility         valueTypes.Bool     `json:"is_has_the_ability"`
		IsInit                  valueTypes.Bool     `json:"is_init"`
		IsReadSet               valueTypes.Bool     `json:"is_read_set"`
		IsReplacing             valueTypes.Bool     `json:"is_replacing"`
		IsReset                 valueTypes.Bool     `json:"is_reset"`
		IsSecond                valueTypes.Bool     `json:"is_second"`
		IsThirdParty            valueTypes.Bool     `json:"is_third_party"`
		ModuleUUID              valueTypes.Integer  `json:"module_uuid"`
		OwnerDevFaultStatus     valueTypes.Integer  `json:"owner_dev_fault_status"`
		P24                     interface{}         `json:"p24"`
		Posx                    interface{}         `json:"posx"`
		Posy                    interface{}         `json:"posy"`
		PsId                    valueTypes.PsId     `json:"ps_id"`
		PsKey                   valueTypes.PsKey    `json:"ps_key"`
		RelState                valueTypes.Integer  `json:"rel_state"`
		Sn                      valueTypes.String   `json:"sn" PointName:"Serial Number"`
		TypeName                valueTypes.String   `json:"type_name"`
		UUID                    valueTypes.Integer  `json:"uuid"`
	} `json:"pageList" PointId:"page_list" PointIdFromChild:"PsKey" PointIdReplace:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
