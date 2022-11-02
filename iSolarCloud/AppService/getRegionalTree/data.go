package getRegionalTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/orgService/getRegionalTree"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
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
		Checked    valueTypes.Bool   `json:"checked"`
		Id         valueTypes.String `json:"id"`
		IsFirstOrg valueTypes.Bool   `json:"isFirstOrg"`
		IsParent   valueTypes.Bool `json:"isParent"`
		Name       valueTypes.String `json:"name"`
		Open       valueTypes.Bool   `json:"open"`
		OrgId      valueTypes.Integer  `json:"org_id"`
		ParentId   valueTypes.Integer `json:"pId"`
		PsId       valueTypes.PsId `json:"ps_id"`
		ShareType  int64  `json:"share_type"`
	} `json:"resultList" DataTable:"true"`
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
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
