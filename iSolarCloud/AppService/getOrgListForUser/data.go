package getOrgListForUser

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"
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
	GcjLatitude    api.Float   `json:"gcj_latitude"`
	GcjLongitude   api.Float   `json:"gcj_longitude"`
	ID             api.Integer `json:"id"`
	IsLeaf         api.Bool    `json:"is_leaf"`
	MapLevel       interface{} `json:"map_level"`
	OrgID          api.Integer `json:"org_id"`
	OrgIndexCode   api.String  `json:"org_index_code"`
	OrgIsShow      api.Integer `json:"org_is_show"`
	OrgLevel       api.Integer `json:"org_level"`
	OrgName        api.String  `json:"org_name"`
	SizeChild      api.Integer `json:"size_child"`
	UpOrgID        api.Integer `json:"up_org_id"`
	Wgs84Latitude  api.Float   `json:"wgs84_latitude"`
	Wgs84Longitude api.Float   `json:"wgs84_longitude"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", api.NewDateTime(""))
	}

	return entries
}
