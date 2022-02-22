package communicationModuleDetail

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/devService/communicationModuleDetail"
const Disabled = false

type RequestData struct {
	Sn string `json:"sn" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CardFlowInfoList []interface{} `json:"cardFlowInfoList"`
	SetMealInfoList  []interface{} `json:"setMealInfoList"`
}
