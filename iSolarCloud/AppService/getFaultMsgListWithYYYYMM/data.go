package getFaultMsgListWithYYYYMM

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/faultService/getFaultMsgListWithYYYYMM"
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


type ResultData   []struct {
	CreateTime    valueTypes.Integer  `json:"create_time"`
	FaultCode     valueTypes.String   `json:"fault_code"`
	FaultLevel    valueTypes.Integer  `json:"fault_level"`
	FaultReason   valueTypes.String   `json:"fault_reason"`
	FaultType     valueTypes.Integer  `json:"fault_type"`
	FaultTypeCode valueTypes.Integer  `json:"fault_type_code"`
	Id            valueTypes.Integer  `json:"id"`
	PsId          valueTypes.PsId  `json:"ps_id"`
	PsKey         valueTypes.PsKey    `json:"ps_key"`
	UUID          valueTypes.Integer  `json:"uuid"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
