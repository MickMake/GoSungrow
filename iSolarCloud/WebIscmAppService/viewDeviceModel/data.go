package viewDeviceModel

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/viewDeviceModel"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DeviceFactoryList []struct {
		CustomerCode    valueTypes.String  `json:"customer_code"`
		CustomerName    valueTypes.String  `json:"customer_name"`
		FactoryAddress  valueTypes.String  `json:"factory_address"`
		FactoryDesc     valueTypes.String  `json:"factory_desc"`
		FactoryLogo     valueTypes.String  `json:"factory_logo"`
		FactoryName     valueTypes.String  `json:"factory_name"`
		FactoryNameEnUs valueTypes.String  `json:"factory_name_en_us"`
		FileName        valueTypes.String  `json:"file_name"`
		ID              valueTypes.Integer `json:"id"`
		Industry        valueTypes.Integer `json:"industry"`
		IndustryName    valueTypes.String  `json:"industry_name"`
		LinkMan         valueTypes.String  `json:"link_man"`
		LinkMethod      valueTypes.String  `json:"link_method"`
		OweMonitor      interface{}        `json:"owe_monitor"`
		OweRemind       interface{}        `json:"owe_remind"`
		Remark          valueTypes.String  `json:"remark"`
	} `json:"deviceFactoryList" PointId:"device_factory_list" DataTable:"true"`
	DeviceTypeList []struct {
		IsRemoteUpgrade valueTypes.Integer `json:"is_remote_upgrade"`
		SysID           valueTypes.String  `json:"sys_id"`
		SysName         valueTypes.String  `json:"sys_name"`
		TypeCode        valueTypes.Integer `json:"type_code"`
		TypeID          valueTypes.Integer `json:"type_id"`
		TypeName        valueTypes.String  `json:"type_name"`
		TypeNameEn      valueTypes.String  `json:"type_name_en"`
		UpdateDate      valueTypes.String  `json:"update_date"`
		ValidFlag       valueTypes.Integer `json:"valid_flag"`
	} `json:"deviceTypeList" PointId:"device_type_list" DataTable:"true"`
	SysTypeList []struct {
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue  valueTypes.String `json:"code_value"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"sysTypeList" PointId:"sys_type_list" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
