package getPowerDevicePointNames

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)


const Url = "/v1/reportService/getPowerDevicePointNames"
const Disabled = false

const (
	DeviceType1 = "1"
	DeviceType3 = "3"
	DeviceType4 = "4"
	DeviceType5 = "5"
	DeviceType7 = "7"
	DeviceType11 = "11"
	DeviceType14 = "14"
	DeviceType17 = "17"
)


type RequestData struct {
	DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("device_type can be one of:\n")
	ret += fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s\n",
		DeviceType1,
		DeviceType3,
		DeviceType4,
		DeviceType5,
		DeviceType7,
		DeviceType11,
		DeviceType14,
		DeviceType17,
		)
	return ret
}


type ResultData   []struct {
	PointCalType int64  `json:"point_cal_type"`
	PointID      int64  `json:"point_id"`
	PointName    string `json:"point_name"`
}
