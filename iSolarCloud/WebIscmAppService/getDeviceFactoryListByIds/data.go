package getDeviceFactoryListByIds

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getDeviceFactoryListByIds"
const Disabled = false
const EndPointName = "WebIscmAppService.getDeviceFactoryListByIds"

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
	DeviceFactoryList []struct {
		Id           valueTypes.Integer `json:"id"`
		CustomerCode valueTypes.String  `json:"customer_code"`
		CustomerName valueTypes.String  `json:"customer_name"`
	} `json:"deviceFactoryList" PointId:"device_factory_list" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"Id"`
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
