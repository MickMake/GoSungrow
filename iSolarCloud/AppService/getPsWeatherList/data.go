package getPsWeatherList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/powerStationService/getPsWeatherList"
const Disabled = false
const EndPointName = "AppService.getPsWeatherList"

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	WeatherList []struct {
		GoStruct   GoStruct.GoStruct   `json:"-" PointIdReplace:"true" PointIdFrom:"DateTime" PointNameDateFormat:"20060102" PointTimestampFrom:"DateTime"`

		DateTime   valueTypes.DateTime `json:"date_time" PointNameDateFormat:"2006/01/02"`
		PsId       valueTypes.PsId     `json:"ps_id"`

		Chill      valueTypes.Float    `json:"chill"`
		Code       valueTypes.Float    `json:"code"`
		CodeName   valueTypes.String   `json:"code_name"`
		Direction  valueTypes.Float    `json:"direction"`
		HighF      valueTypes.Float    `json:"high" PointUnit:"F"`
		HighC      valueTypes.Float    `json:"highc" PointUnit:"C"`
		Humidity   valueTypes.Float    `json:"humidity"`
		LowF       valueTypes.Float    `json:"low" PointUnit:"F"`
		LowC       valueTypes.Float    `json:"lowc" PointUnit:"C"`
		Pressure   valueTypes.Float    `json:"pressure" PointUnit:"hPa"`
		Rising     valueTypes.Float    `json:"rising"`

		Speed      valueTypes.Float    `json:"speed"`

		Sunrise    valueTypes.Time     `json:"sunrise"`
		Sunset     valueTypes.Time     `json:"sunset"`

		Visibility valueTypes.Float    `json:"visibility"`
	} `json:"weather_list" PointIdFromChild:"DateTime" PointNameDateFormat:"20060102" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
