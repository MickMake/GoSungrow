package getFaultMsgListWithYYYYMM

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/faultService/getFaultMsgListWithYYYYMM"
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


type ResultData   []struct {
	CreateTime    int64  `json:"create_time"`
	FaultCode     string `json:"fault_code"`
	FaultLevel    int64  `json:"fault_level"`
	FaultReason   string `json:"fault_reason"`
	FaultType     int64  `json:"fault_type"`
	FaultTypeCode int64  `json:"fault_type_code"`
	ID            int64  `json:"id"`
	PsID          int64  `json:"ps_id"`
	PsKey         string `json:"ps_key"`
	UUID          int64  `json:"uuid"`
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
