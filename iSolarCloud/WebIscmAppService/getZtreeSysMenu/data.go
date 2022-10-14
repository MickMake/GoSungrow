package getZtreeSysMenu

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"errors"
	"fmt"
)

const Url = "/v1/userService/getZtreeSysMenu"
const Disabled = true

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
	Dummy string `json:"dummy"`
}

func (e *ResultData) IsValid() error {
	var err error
	switch {
	case e.Dummy == "":
		break
	default:
		err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	}
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
