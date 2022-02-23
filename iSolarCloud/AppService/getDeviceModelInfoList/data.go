package getDeviceModelInfoList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/devService/getDeviceModelInfoList"
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
	ComType           string `json:"com_type"`
	DeviceFactoryID   string `json:"device_factory_id"`
	DeviceFactoryName string `json:"device_factory_name"`
	DeviceModel       string `json:"device_model"`
	DeviceModelCode   string `json:"device_model_code"`
	DeviceModelID     int64  `json:"device_model_id"`
	DeviceType        int64  `json:"device_type"`
	IsRemoteUpgrade   int64  `json:"is_remote_upgrade"`
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
