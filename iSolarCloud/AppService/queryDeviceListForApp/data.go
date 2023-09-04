package queryDeviceListForApp

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/queryDeviceListForApp"
const Disabled = false
const EndPointName = "AppService.queryDeviceListForApp"

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

// ResultData - Both getDeviceList & queryDeviceListForApp have the same output.
type ResultData struct {
	PageList []struct {
		GoStruct                GoStruct.GoStruct   `json:"-" PointDeviceFrom:"PsKey" PointIdFrom:"PsKey" PointIdReplace:"true"`

		PsKey                   valueTypes.PsKey    `json:"ps_key"`
		PsId                    valueTypes.PsId     `json:"ps_id"`
		DeviceType              valueTypes.Integer  `json:"device_type"`
		DeviceCode              valueTypes.Integer  `json:"device_code"`
		ChannelId               valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		Sn                      valueTypes.String   `json:"sn" PointName:"Serial Number"`
		FactoryName             valueTypes.String   `json:"factory_name"`

		AttrId                  valueTypes.Integer  `json:"attr_id"`
		CommandStatus           valueTypes.Integer  `json:"command_status"`
		ConnectState            valueTypes.Bool     `json:"connect_state"`
		DataFlag                valueTypes.Integer  `json:"data_flag"`
		DataFlagDetail          valueTypes.Integer  `json:"data_flag_detail"`

		DevFaultStatus          valueTypes.Integer  `json:"dev_fault_status"`
		DevStatus               valueTypes.Bool     `json:"dev_status"`
		DeviceArea              valueTypes.String   `json:"device_area"`
		DeviceFactoryDate       valueTypes.DateTime `json:"device_factory_date" PointNameDateFormat:"DateTimeLayout"`
		DeviceId                valueTypes.Integer  `json:"device_id"`
		DeviceModel             valueTypes.String   `json:"device_model"`
		DeviceModelCode         valueTypes.String   `json:"device_model_code"`
		DeviceModelId           valueTypes.Integer  `json:"device_model_id"`
		DeviceName              valueTypes.String   `json:"device_name"`
		DeviceProSn             valueTypes.String   `json:"device_pro_sn"`
		DeviceState             valueTypes.Integer  `json:"device_state"`
		DeviceSubType           interface{}         `json:"device_sub_type"`
		DeviceSubTypeName       interface{}         `json:"device_sub_type_name"`

		IsCountryCheck          valueTypes.Bool     `json:"is_country_check"`
		IsHasFunctionEnum       valueTypes.Bool     `json:"is_has_function_enum"`
		IsHasTheAbility         valueTypes.Bool     `json:"is_has_the_ability"`
		IsInit                  valueTypes.Bool     `json:"is_init"`
		IsReadSet               valueTypes.Bool     `json:"is_read_set"`
		IsReplacing             valueTypes.Bool     `json:"is_replacing"`
		IsReset                 valueTypes.Bool     `json:"is_reset"`
		IsSecond                valueTypes.Bool     `json:"is_second"`
		IsThirdParty            valueTypes.Bool     `json:"is_third_party"`

		InstallerDevFaultStatus valueTypes.Integer  `json:"installer_dev_fault_status"`
		InverterModelType       valueTypes.Integer  `json:"inverter_model_type"`
		ModuleUUID              valueTypes.Integer  `json:"module_uuid"`
		OwnerDevFaultStatus     valueTypes.Integer  `json:"owner_dev_fault_status"`
		P24                     interface{}         `json:"p24"`
		Posx                    interface{}         `json:"posx"`
		Posy                    interface{}         `json:"posy"`
		RelState                valueTypes.Bool     `json:"rel_state"`
		TypeName                valueTypes.String   `json:"type_name"`
		UUID                    valueTypes.Integer  `json:"uuid"`
	} `json:"pageList" PointId:"devices" DataTable:"true" DataTableSortOn:"PsKey"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
