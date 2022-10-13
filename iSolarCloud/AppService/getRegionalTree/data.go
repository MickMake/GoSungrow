package getRegionalTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/orgService/getRegionalTree"
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


type ResultData   struct {
	ResultList []struct {
		Checked    api.Bool   `json:"checked"`
		ID         api.String `json:"id"`
		IsFirstOrg api.Bool   `json:"isFirstOrg"`
		IsParent   api.Bool `json:"isParent"`
		Name       api.String `json:"name"`
		Open       api.Bool   `json:"open"`
		OrgID      api.Integer  `json:"org_id"`
		PID        api.Integer `json:"pId"`
		PsID       api.Integer `json:"ps_id"`
		ShareType  int64  `json:"share_type"`
	} `json:"resultList"`
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
