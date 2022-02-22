package getPsListByName

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPsListByName"
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
	PsID         int64  `json:"ps_id"`
	PsName       string `json:"ps_name"`
	PsShortName  string `json:"ps_short_name"`
	PsTimezone   string `json:"ps_timezone"`
	PsTimezoneID int64  `json:"ps_timezone_id"`
	ShareType    string `json:"share_type"`
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
