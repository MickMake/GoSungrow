package getDeviceModel

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getDeviceModel"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	AllFactoryList []struct {
		Id             valueTypes.Integer `json:"id"`
		FactoryName    valueTypes.String  `json:"factory_name"`
		FactoryAddress valueTypes.String  `json:"factory_address"`
		FactoryDesc    valueTypes.String  `json:"factory_desc"`
		FactoryLogo    valueTypes.String  `json:"factory_logo"`
		LinkMan        valueTypes.String  `json:"link_man"`
		Remark         valueTypes.String  `json:"remark"`
	} `json:"all_factory_list" DataTable:"true" DataTableSortOn:"Id"`
	DeviceFactoryList []struct {
		Id              valueTypes.Integer `json:"id"`
		CustomerCode    valueTypes.String  `json:"customer_code"`
		CustomerName    valueTypes.String  `json:"customer_name"`
		FactoryAddress  valueTypes.String  `json:"factory_address"`
		FactoryDesc     valueTypes.String  `json:"factory_desc"`
		FactoryLogo     valueTypes.String  `json:"factory_logo"`
		FactoryName     valueTypes.String  `json:"factory_name"`
		FactoryNameEnUs valueTypes.String  `json:"factory_name_en_us"`
		FileName        valueTypes.String  `json:"file_name"`
		Industry        interface{}        `json:"industry"`
		IndustryName    interface{}        `json:"industry_name"`
		LinkMan         valueTypes.String  `json:"link_man"`
		LinkMethod      valueTypes.String  `json:"link_method"`
		OweMonitor      interface{}        `json:"owe_monitor"`
		OweRemind       interface{}        `json:"owe_remind"`
		Remark          valueTypes.String  `json:"remark"`
	} `json:"deviceFactoryList" PointId:"device_factory_list" DataTable:"true" DataTableSortOn:"Id"`
	DeviceTypeList []struct {
		TypeId          valueTypes.Integer  `json:"type_id"`
		TypeCode        valueTypes.Integer  `json:"type_code"`
		TypeName        valueTypes.String   `json:"type_name"`
		TypeNameEn      valueTypes.String   `json:"type_name_en"`
		SysId           valueTypes.String   `json:"sys_id"`
		SysName         valueTypes.String   `json:"sys_name"`
		IsRemoteUpgrade valueTypes.Bool     `json:"is_remote_upgrade"`
		UpdateDate      valueTypes.DateTime `json:"update_date"`
		ValidFlag       valueTypes.Bool     `json:"valid_flag"`
	} `json:"deviceTypeList" PointId:"device_type_list" DataTable:"true" DataTableSortOn:"TypeId"`
	PowerDeviceModel struct {
		Id                        valueTypes.Integer  `json:"id"`
		SysId                     valueTypes.String   `json:"sys_id"`
		SysType                   valueTypes.String   `json:"sys_type"`
		ComType                   valueTypes.String   `json:"com_type"`
		DeviceFactoryId           valueTypes.String   `json:"device_factory_id"`
		DeviceFactoryName         valueTypes.String   `json:"device_factory_name"`
		DeviceModel               valueTypes.String   `json:"device_model"`
		DeviceModelCode           valueTypes.String   `json:"device_model_code"`
		DeviceSupplier            valueTypes.String   `json:"device_supplier"`
		DeviceType                valueTypes.Integer  `json:"device_type"`
		DeviceTypeName            valueTypes.String   `json:"device_typename"`
		DeviceVersion             valueTypes.String   `json:"device_version"`
		Edition                   valueTypes.Integer  `json:"edition"`
		InverterModelType         valueTypes.Integer  `json:"inverter_model_type"`
		IsCountryCheck            valueTypes.Bool     `json:"is_country_check"`
		IsDefaultModel            valueTypes.Bool     `json:"is_default_model"`
		IsHaveBatteryWarranty     valueTypes.Bool     `json:"is_have_battery_warranty"`
		IsReadSet                 valueTypes.Bool     `json:"is_read_set"`
		IsRemoteUpgrade           valueTypes.Bool     `json:"is_remote_upgrade"`
		IsReset                   valueTypes.Bool     `json:"is_reset"`
		IsSupportParamSet         valueTypes.Bool     `json:"is_support_paramset" PointId:"is_support_param_set"`
		IsThirdParty              valueTypes.Bool     `json:"is_third_party"`
		IsValid                   valueTypes.Bool     `json:"isvalid" PointId:"is_valid"`
		ModelIsSupportFaultRecord valueTypes.Integer  `json:"model_is_support_fault_record"`
		Phone                     valueTypes.String   `json:"phone"`
		Protocol                  valueTypes.String   `json:"protocol"`
		Remark                    valueTypes.String   `json:"remark"`
		UpdateDate                valueTypes.DateTime `json:"updatedate" PointId:"update_date"`
		UpdateUserCode            valueTypes.String   `json:"updateusercode" PointId:"update_user_code"`
	} `json:"powerDeviceModel"`
	SysTypeList []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"sysTypeList" PointId:"sys_type_list" DataTable:"true" DataTableSortOn:"CodeValue"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
