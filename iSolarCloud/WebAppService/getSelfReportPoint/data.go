package getSelfReportPoint

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/reportService/getSelfReportPoint"
const Disabled = false
const EndPointName = "WebAppService.getSelfReportPoint"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"PointId"`

	PointId           valueTypes.PointId `json:"point_id" PointDevice:""`
	PointName         valueTypes.String  `json:"point_name"`
	Id                valueTypes.Integer `json:"id"`
	DeviceType        valueTypes.Integer `json:"device_type"`
	Period            valueTypes.Integer `json:"period"`
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
