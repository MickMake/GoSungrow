package getOrgListForUser

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/orgService/getOrgListForUser"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GcjLatitude    valueTypes.Float   `json:"gcj_latitude"`
	GcjLongitude   valueTypes.Float   `json:"gcj_longitude"`
	Id             valueTypes.Integer `json:"id"`
	IsLeaf         valueTypes.Bool    `json:"is_leaf"`
	MapLevel       interface{} `json:"map_level"`
	OrgId          valueTypes.Integer `json:"org_id"`
	OrgIndexCode   valueTypes.String  `json:"org_index_code"`
	OrgIsShow      valueTypes.Integer `json:"org_is_show"`
	OrgLevel       valueTypes.Integer `json:"org_level"`
	OrgName        valueTypes.String  `json:"org_name"`
	SizeChild      valueTypes.Integer `json:"size_child"`
	UpOrgId        valueTypes.Integer `json:"up_org_id"`
	Wgs84Latitude  valueTypes.Float   `json:"wgs84_latitude"`
	Wgs84Longitude valueTypes.Float   `json:"wgs84_longitude"`
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
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
