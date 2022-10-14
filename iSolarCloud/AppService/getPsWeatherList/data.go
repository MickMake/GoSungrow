package getPsWeatherList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsWeatherList"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
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
		Chill      valueTypes.Float    `json:"chill"`
		Code       valueTypes.Float    `json:"code"`
		CodeName   valueTypes.Float    `json:"code_name"`
		DateTime   valueTypes.DateTime `json:"date_time"`
		Direction  valueTypes.Float    `json:"direction"`
		High       valueTypes.Float    `json:"high"`
		Highc      valueTypes.Float    `json:"highc"`
		Humidity   valueTypes.Float    `json:"humidity"`
		Low        valueTypes.Float    `json:"low"`
		Lowc       valueTypes.Float    `json:"lowc"`
		Pressure   valueTypes.Float    `json:"pressure"`
		PsID       valueTypes.Float    `json:"ps_id"`
		Rising     valueTypes.Float    `json:"rising"`
		Speed      valueTypes.Float    `json:"speed"`
		Sunrise    valueTypes.Float    `json:"sunrise"`
		Sunset     valueTypes.Float    `json:"sunset"`
		Visibility valueTypes.Float    `json:"visibility"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := api.JoinWithDots(0, "", apiReflect.GetName("", *e), e.Request.PsId)
		// entries.StructToPoints(e.Response.ResultData, pkg, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		for _, v := range e.Response.ResultData.WeatherList {
			entries.StructToPoints(v, api.JoinWithDots(0, valueTypes.DateTimeLayoutDay, pkg, v.DateTime), e.Request.PsId.String(), valueTypes.NewDateTime(""))
		}
	}

	return entries
}
