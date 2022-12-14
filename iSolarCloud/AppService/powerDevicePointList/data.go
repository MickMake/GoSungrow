package powerDevicePointList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/reportService/powerDevicePointList"
const Disabled = false
const EndPointName = "AppService.powerDevicePointList"

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
	PageDataList []Point            `json:"pageDataList" PointId:"points" PointArrayFlatten:"false" DataTable:"true"` // DataTableSortOn:"Id"`
	CurrentPage  valueTypes.Integer `json:"currentPage" PointIgnore:"true"`
	PageSize     valueTypes.Integer `json:"pageSize" PointIgnore:"true"`
	TotalCounts  valueTypes.Integer `json:"totalCounts" PointId:"total_count"`
	TotalPages   valueTypes.Integer `json:"totalPages" PointIgnore:"true"`
}

type Point struct {
	CreateTime   valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	DeviceType   valueTypes.Integer  `json:"device_type"`
	Id           valueTypes.Integer  `json:"id"`
	Period       valueTypes.Integer  `json:"period"` // 0, 1, 2, 3, 4
	PointId      valueTypes.PointId  `json:"point_id"`
	PointName    valueTypes.String   `json:"point_name" PointIgnore:"true"` // Old name of point.
	PointNameNew valueTypes.String   `json:"point_name_new" PointId:"name"`
	TypeName     valueTypes.String   `json:"type_name"`
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
