package getOrgListForUser

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/orgService/getOrgListForUser"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData []struct {
	GcjLatitude    interface{} `json:"gcj_latitude"`
	GcjLongitude   interface{} `json:"gcj_longitude"`
	ID             int64       `json:"id"`
	IsLeaf         int64       `json:"is_leaf"`
	MapLevel       interface{} `json:"map_level"`
	OrgID          int64       `json:"org_id"`
	OrgIndexCode   string      `json:"org_index_code"`
	OrgIsShow      int64       `json:"org_is_show"`
	OrgLevel       int64       `json:"org_level"`
	OrgName        string      `json:"org_name"`
	SizeChild      int64       `json:"size_child"`
	UpOrgID        int64       `json:"up_org_id"`
	Wgs84Latitude  interface{} `json:"wgs84_latitude"`
	Wgs84Longitude interface{} `json:"wgs84_longitude"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
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
