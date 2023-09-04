package getDeviceList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceList"
const Disabled = false
const EndPointName = "AppService.getDeviceList"

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
	PageList []Device `json:"pageList" PointId:"devices" DataTable:"true" DataTableSortOn:"PsKey"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

type Device struct {
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
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

// type Device struct {
// 	Vendor        valueTypes.String
//
// 	PsKey         valueTypes.PsKey
// 	PsId          valueTypes.PsId
// 	DeviceType    valueTypes.Integer
// 	DeviceCode    valueTypes.Integer
// 	ChannelId     valueTypes.Integer
//
// 	DeviceName    valueTypes.String
// 	DeviceProSn   valueTypes.String
// 	DeviceModel   valueTypes.String
// 	DeviceModelId valueTypes.Integer
// 	TypeName      valueTypes.String
// 	DeviceState   valueTypes.Integer
// 	DevStatus     valueTypes.Bool
// 	Uuid          valueTypes.Integer
// }

type Devices []Device

func (e *EndPoint) GetDevices() Devices {
	return e.Response.ResultData.PageList
}

func (e *EndPoint) GetDevicesTable() output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsId.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelId.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }
		//
		// data := e.GetDevices()
		table = GetDevicesTable(e.Response.ResultData.PageList)
	}
	return table
}

func GetDevicesTable(data Devices) output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsId.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelId.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }

		table = output.NewTable(
			"Vendor",
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
			"UUID",
		)
		table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Vendor",
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	"UUID",
		// )

		for _, d := range data {
			_ = table.AddRow(d.FactoryName.String(),
				d.PsKey.String(),
				d.PsId.String(),
				d.DeviceType.String(),
				d.DeviceCode.String(),
				d.ChannelId.String(),
				d.TypeName.String(),
				d.DeviceProSn.String(),
				d.DeviceModel.String(),
				d.DeviceModelId.String(),
				d.DeviceName.String(),
				d.DeviceState.String(),
				d.DevStatus.String(),
				d.UUID.String(),
			)
		}
	}
	return table
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
