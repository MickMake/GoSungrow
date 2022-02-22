package getPsWeatherList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPsWeatherList"
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
	WeatherList []struct {
		Chill      string `json:"chill"`
		Code       string `json:"code"`
		CodeName   string `json:"code_name"`
		DateTime   string `json:"date_time"`
		Direction  string `json:"direction"`
		High       string `json:"high"`
		Highc      string `json:"highc"`
		Humidity   string `json:"humidity"`
		Low        string `json:"low"`
		Lowc       string `json:"lowc"`
		Pressure   string `json:"pressure"`
		PsID       string `json:"ps_id"`
		Rising     string `json:"rising"`
		Speed      string `json:"speed"`
		Sunrise    string `json:"sunrise"`
		Sunset     string `json:"sunset"`
		Visibility string `json:"visibility"`
	} `json:"weather_list"`
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
