package powerDevicePointList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/powerDevicePointList"
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
	CurrentPage  valueTypes.Integer `json:"currentPage" PointIgnore:"true"`
	PageDataList []struct {
		CreateTime   valueTypes.DateTime `json:"create_time" PointIgnore:"true"`
		DeviceType   valueTypes.Integer  `json:"device_type"`
		Id           valueTypes.Integer  `json:"id"`
		Period       valueTypes.Integer  `json:"period"`		// 0, 1, 2, 3, 4
		PointId      valueTypes.PointId  `json:"point_id"`
		PointName    valueTypes.String `json:"point_name" PointIgnore:"true"`	// Old name of point.
		PointNameNew valueTypes.String `json:"point_name_new" PointId:"name"`
		TypeName     valueTypes.String `json:"type_name"`
	} `json:"pageDataList" PointId:"points" PointArrayFlatten:"false" DataTable:"true"`
	PageSize    valueTypes.Integer `json:"pageSize" PointIgnore:"true"`
	TotalCounts valueTypes.Integer `json:"totalCounts" PointId:"total_counts"`
	TotalPages  valueTypes.Integer `json:"totalPages" PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e,  "system", GoStruct.EndPointPath{})
	}

	return entries
}
