package getPowerDeviceFactory

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceFactory"
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

type ResultData   struct {
	IndustryList []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"industryList" PointId:"industry_list" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"CodeValue"`
	PowerDeviceList []struct {
		Id              valueTypes.Integer `json:"id"`
		CustomerCode    valueTypes.String  `json:"customer_code"`
		CustomerName    valueTypes.String  `json:"customer_name"`
		FactoryAddress  valueTypes.String  `json:"factory_address"`
		FactoryDesc     valueTypes.String  `json:"factory_desc"`
		FactoryLogo     valueTypes.String  `json:"factory_logo"`
		FactoryName     valueTypes.String  `json:"factory_name"`
		FactoryNameEnUs valueTypes.String  `json:"factory_name_en_us"`
		FileName        valueTypes.String  `json:"file_name"`
		Industry        valueTypes.Integer `json:"industry"`
		IndustryName    valueTypes.String  `json:"industry_name"`
		LinkMan         valueTypes.String  `json:"link_man"`
		LinkMethod      valueTypes.String  `json:"link_method"`
		OweMonitor      interface{}        `json:"owe_monitor"`
		OweRemind       interface{}        `json:"owe_remind"`
		Remark          valueTypes.String  `json:"remark"`
	} `json:"powerDeviceList" PointId:"power_device_list" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"Id"`
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
