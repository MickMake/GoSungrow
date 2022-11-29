package getPointInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPointInfo"
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
	DeviceTypeList []struct {
		DeviceType valueTypes.Integer `json:"device_type"`
		DeviceName valueTypes.String  `json:"device_name"`
	} `json:"deviceTypeList" PointId:"device_type_list" DataTable:"true" DataTableSortOn:"DeviceType"`
	DisplayModeList []struct {
		PointType valueTypes.String `json:"point_type" DataTableSortOn:"PointType"`
		CodeName  valueTypes.String `json:"code_name"`
	} `json:"displayModeList" PointId:"display_mode_list" DataTable:"true"`
	PointCalTypeList []struct {
		CodeValue  valueTypes.String `json:"code_value"`
		CodeName   valueTypes.String `json:"code_name"`
		CodeValue2 interface{}       `json:"code_value2"`
	} `json:"pointCalTypeList" PointId:"point_cal_type_list" DataTable:"true" DataTableSortOn:"CodeValue"`
	PointTypeList []struct {
		PointType valueTypes.String `json:"point_type"`
		CodeName  valueTypes.String `json:"code_name"`
	} `json:"pointTypeList" PointId:"point_type_list" DataTable:"true" DataTableSortOn:"PointType"`
	PolymerizationModeList []struct {
		PointType valueTypes.String `json:"point_type"`
		CodeName  valueTypes.String `json:"code_name"`
	} `json:"polymerizationModeList" PointId:"polymerization_mode_list" DataTable:"true" DataTableSortOn:"PointType"`
	PowerPointManage interface{} `json:"powerPointManage"`
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
