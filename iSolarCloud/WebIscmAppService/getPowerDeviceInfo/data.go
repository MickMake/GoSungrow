package getPowerDeviceInfo

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getPowerDeviceInfo"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPowerDeviceInfo"
)

type RequestData struct {
	Uuid       valueTypes.String  `json:"uuid" required:"true"`
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceList []struct {
		GoStruct GoStruct.GoStruct `json:"-" PointIdFrom:"PsId" PointIdReplace:"true" PointDeviceFrom:"PsId"`

		PsId             valueTypes.PsId    `json:"ps_id"`
		PsName           valueTypes.String  `json:"ps_name"`
		BatVersion       valueTypes.String  `json:"bat_version"`
		BatteryVersion   interface{}        `json:"battery_version"`
		ChannelId        valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
		ClassName        valueTypes.String  `json:"class_name"`
		ComType          valueTypes.String  `json:"com_type"`
		CommandStatus    valueTypes.Integer `json:"command_status"`
		CountryId        valueTypes.Integer `json:"country_id"`
		CpldVersion      interface{}        `json:"cpld_version"`
		DeviceCode       valueTypes.Integer `json:"device_code"`
		DeviceFactoryId  interface{}        `json:"device_factory_id"`
		DeviceId         valueTypes.Integer `json:"device_id"`
		DeviceModel      valueTypes.String  `json:"device_model"`
		DeviceModelCode  valueTypes.String  `json:"device_model_code"`
		DeviceModelId    valueTypes.Integer `json:"device_model_id"`
		DeviceName       valueTypes.String  `json:"device_name"`
		DeviceProFactory interface{}        `json:"device_pro_factry" PointId:"device_pro_factory"`
		DeviceProSn      valueTypes.String  `json:"device_pro_sn"`
		DeviceStatus     valueTypes.Integer `json:"device_status"`
		DeviceSubType    interface{}        `json:"device_sub_type"`
		DeviceType       valueTypes.Integer `json:"device_type"`
		FactoryName      valueTypes.String  `json:"factory_name"`
		GridTypeId       valueTypes.Integer `json:"grid_type_id"`
		IsDisplay        valueTypes.Bool    `json:"is_display"`
		IsHaveVersion    valueTypes.Bool    `json:"is_haveversion" PointId:"is_have_version"`
		IsLoad           valueTypes.Bool    `json:"is_load"`
		IsPublic         valueTypes.Bool    `json:"is_public"`
		IsTotalDc        valueTypes.Bool    `json:"is_total_dc"`
		IsUse            valueTypes.Bool    `json:"is_use"`
		IsVirtualUnit    valueTypes.Bool    `json:"is_virtual_unit"`
		LcdVersion       valueTypes.String  `json:"lcd_version"`
		MachineVersion   interface{}        `json:"machine_version"`
		McuVersion       interface{}        `json:"mcu_version"`
		MdspVersion      valueTypes.String  `json:"mdsp_version"`
		PvdVersion       interface{}        `json:"pvd_version"`
		RelState         valueTypes.Integer `json:"rel_state"`
		SdspVersion      valueTypes.String  `json:"sdsp_version"`
		SyncDate         valueTypes.String  `json:"sync_date"`
		TempVersion      interface{}        `json:"temp_version"`
		UpdateDate       valueTypes.String  `json:"update_date"`
		UUID             valueTypes.Integer `json:"uuid"`
		Version1         valueTypes.String  `json:"version1"`
		Version2         valueTypes.String  `json:"version2"`
		Version3         valueTypes.String  `json:"version3"`
		Version4         valueTypes.String  `json:"version4"`
		Version5         interface{}        `json:"version5"`
		Version6         interface{}        `json:"version6"`
		Version7         interface{}        `json:"version7"`
		Version8         interface{}        `json:"version8"`
		Version9         interface{}        `json:"version9"`
		Version10        interface{}        `json:"version10"`
		Version11        interface{}        `json:"version11"`
		Version12        interface{}        `json:"version12"`
	} `json:"deviceList" PointId:"device_list" PointIdReplace:"true" DataTable:"false"`
	DeviceProValue []struct {
		ProCode            valueTypes.String  `json:"pro_code"`
		ProName            valueTypes.String  `json:"pro_name"`
		ProValue           valueTypes.String  `json:"pro_value"`
		PropertyOrderId    valueTypes.Integer `json:"property_order_id"`
		IsParamSetProperty valueTypes.Bool    `json:"is_param_set_property"`
		CheckRule          valueTypes.String  `json:"check_rule"`
	} `json:"deviceProValue" PointId:"device_products" PointIdReplace:"true" DataTable:"true"`
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
