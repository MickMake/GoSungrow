package getPowerDevicePointNames

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api/apiReflect"
	"errors"
)


const Url = "/v1/reportService/getPowerDevicePointNames"

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

func (rd *RequestData) IsValid() error {
	var err error
	for range Only.Once {
		if rd == nil {
			err = errors.New("empty request data")
			break
		}
		err = apiReflect.VerifyOptionsRequired(*rd)
		if err != nil {
			break
		}
	}
	return err
}


type ResultData   []struct {
	PointCalType int64  `json:"point_cal_type"`
	PointID      int64  `json:"point_id"`
	PointName    string `json:"point_name"`
}
