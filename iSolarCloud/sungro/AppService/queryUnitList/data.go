package queryUnitList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/userService/queryUnitList"

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
	IsBasicUnit  int64  `json:"is_basic_unit"`
	TargetUnit   string `json:"target_unit"`
	UnitConverID int64  `json:"unit_conver_id"`
	UnitName     string `json:"unit_name"`
	UnitType     int64  `json:"unit_type"`
}
