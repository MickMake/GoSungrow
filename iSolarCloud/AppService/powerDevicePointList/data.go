package powerDevicePointList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
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
	CurrentPage  int64 `json:"currentPage"`
	PageDataList []struct {
		CreateTime   string `json:"create_time"`
		DeviceType   int64  `json:"device_type"`
		ID           int64  `json:"id"`
		Period       int64  `json:"period"`
		PointID      int64  `json:"point_id"`
		PointName    string `json:"point_name"`
		PointNameNew string `json:"point_name_new"`
		TypeName     string `json:"type_name"`
	} `json:"pageDataList"`
	PageSize    int64 `json:"pageSize"`
	TotalCounts int64 `json:"totalCounts"`
	TotalPages  int64 `json:"totalPages"`
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
