package showPSView

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/showPSView"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BatteryPlateCom       string        `json:"batteryPlateCom"`
	EnvironmentCom        string        `json:"environmentCom"`
	IsPlatformDefaultUnit string        `json:"is_platform_default_unit"`
	Normalization         string        `json:"normalization"`
	Todayvalue            string        `json:"todayvalue"`
	Todayweather          string        `json:"todayweather"`
	TotalAllPower         api.UnitValue `json:"totalAllPower"`
	Yestedayvalue         string        `json:"yestedayvalue"`
	Yestedayweather       string        `json:"yestedayweather"`
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
