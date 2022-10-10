package psForcastInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/psForcastInfo"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	AreaForcastList []struct {
		Chill             api.String   `json:"chill"`
		City              api.String   `json:"city"`
		Code              string       `json:"code"`
		CodeName          api.String   `json:"code_name"`
		DateTime          api.DateTime `json:"date_time"`
		Direction         api.Float    `json:"direction"`
		High              api.Float    `json:"high"`
		Highc             api.Float    `json:"highc"`
		Humidity          api.Float    `json:"humidity"`
		Low               api.Float    `json:"low"`
		Lowc              api.Float    `json:"lowc"`
		Pressure          api.Float    `json:"pressure"`
		PsKnowledge       api.String   `json:"ps_knowledge"`
		Rising            string       `json:"rising"`
		Speed             api.Float    `json:"speed"`
		SpeedOriginal     api.Float    `json:"speed_original"`
		SpeedOriginalUnit api.String   `json:"speed_original_unit"`
		SpeedUnit         api.String   `json:"speed_unit"`
		Sunrise           api.String   `json:"sunrise"`
		Sunset            api.String   `json:"sunset"`
		Visibility        api.Float    `json:"visibility"`
		WeatherDesc       api.String   `json:"weather_desc"`
		WeatherURL        api.String   `json:"weather_url"`
	} `json:"areaForcastList"`
	StationsCityCode []struct {
		City   api.String  `json:"city"`
		PsID   api.Integer `json:"ps_id"`
		PsName api.String  `json:"ps_name"`
	} `json:"stationsCityCode"`
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
