package getPowerDeviceModelTechList

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceModelTechList"
const Disabled = false

const (
	DeviceType1  = "1"
	DeviceType3  = "3"
	DeviceType4  = "4"
	DeviceType5  = "5"
	DeviceType7  = "7"
	DeviceType11 = "11"
	DeviceType14 = "14"
	DeviceType17 = "17"
	DeviceType22 = "22"
	DeviceType23 = "23"
	DeviceType26 = "26"
	DeviceType37 = "37"
	DeviceType41 = "41"
	DeviceType43 = "43"
	DeviceType47 = "47"
)

var DeviceTypes = []string{
	DeviceType1,
	DeviceType3,
	DeviceType4,
	DeviceType5,
	DeviceType7,
	DeviceType11,
	DeviceType14,
	DeviceType17,
	DeviceType22,
	DeviceType23,
	DeviceType26,
	DeviceType37,
	DeviceType41,
	DeviceType43,
	DeviceType47,
}

type RequestData struct {
	DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	CodeID          api.Integer `json:"code_id"`
	CodeName        string      `json:"code_name"`
	CodeValue       string      `json:"code_value"`
	DefaultValue    interface{} `json:"default_value"`
	TechDescription string      `json:"tech_description"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}
