package getPowerDevicePointInfo

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/reportService/getPowerDevicePointInfo"
const Disabled = false

type RequestData struct {
	Id string `json:"id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	DeviceType    int64  `json:"device_type" PointId:"device_type" PointType:""`
	ID            int64  `json:"id" PointId:"id" PointType:""`
	Period        int64  `json:"period" PointId:"period" PointType:""`
	PointID       int64  `json:"point_id" PointId:"point_id" PointType:""`
	PointName     string `json:"point_name" PointId:"point_name" PointType:""`
	ShowPointName string `json:"show_point_name" PointId:"show_point_name" PointType:""`
	TranslationID int64  `json:"translation_id" PointId:"translation_id" PointType:""`
}

func (e *ResultData) IsValid() error {
	var err error
		// switch {
		// 	case e.Dummy == "":
		// 		break
		// 	default:
		// 		err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
		// }
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
