package getRemoteParamSettingList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)


const Url = ""


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


type ResultData struct {
}
