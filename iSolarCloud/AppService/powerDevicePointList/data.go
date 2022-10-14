package powerDevicePointList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/powerDevicePointList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CurrentPage  valueTypes.Integer `json:"currentPage"`
	PageDataList []struct {
		CreateTime   valueTypes.DateTime `json:"create_time"`
		DeviceType   valueTypes.Integer  `json:"device_type"`
		ID           valueTypes.Integer  `json:"id"`
		Period       valueTypes.Integer  `json:"period"`		// 0, 1, 2, 3, 4
		PointID      valueTypes.Integer  `json:"point_id"`
		PointName    valueTypes.String `json:"point_name"`
		PointNameNew valueTypes.String `json:"point_name_new"`
		TypeName     valueTypes.String `json:"type_name"`
	} `json:"pageDataList"`
	PageSize    valueTypes.Integer `json:"pageSize"`
	TotalCounts valueTypes.Integer `json:"totalCounts"`
	TotalPages  valueTypes.Integer `json:"totalPages"`
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

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		entries.StructToPoints(e.Response.ResultData, pkg, "system", valueTypes.NewDateTime(""))
		for _, v := range e.Response.ResultData.PageDataList {
			entries.StructToPoints(v, fmt.Sprintf("%s.%.3d", pkg, v.ID.Value()), "system", valueTypes.NewDateTime(""))
		}
	}

	return entries
}
