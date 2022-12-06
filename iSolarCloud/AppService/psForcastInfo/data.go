package psForcastInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)


const Url = "/v1/powerStationService/psForcastInfo"
const Disabled = false
const EndPointName = "AppService.psForcastInfo"

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
	AreaForecastList []struct {
		GoStruct          GoStruct.GoStruct   `json:"-" PointIdReplace:"true" PointIdFrom:"DateTime" PointNameDateFormat:"20060102" PointTimestampFrom:"DateTime"`

		DateTime          valueTypes.DateTime `json:"date_time" PointNameDateFormat:"2006-01-02 15:04:05"`

		City              valueTypes.String   `json:"city"`
		Chill             valueTypes.Float    `json:"chill"`
		Code              valueTypes.Integer  `json:"code"`
		CodeName          valueTypes.String   `json:"code_name"`
		Direction         valueTypes.Float    `json:"direction"`
		HighF             valueTypes.Float    `json:"high" PointUnit:"F"`
		HighC             valueTypes.Float    `json:"highc" PointUnit:"C"`
		Humidity          valueTypes.Float    `json:"humidity"`
		LowF              valueTypes.Float    `json:"low" PointUnit:"F"`
		LowC              valueTypes.Float    `json:"lowc" PointUnit:"C"`
		Pressure          valueTypes.Float    `json:"pressure" PointUnit:"hPa"`
		Rising            valueTypes.Bool     `json:"rising"`

		Speed             valueTypes.Float    `json:"speed" PointUnitFrom:"SpeedUnit"`
		SpeedUnit         valueTypes.String   `json:"speed_unit"`

		SpeedOriginal     valueTypes.Float    `json:"speed_original" PointUnitFrom:"SpeedOriginalUnit"`
		SpeedOriginalUnit valueTypes.String   `json:"speed_original_unit" PointIgnore:"true"`

		Sunrise           valueTypes.Time     `json:"sunrise"`
		Sunset            valueTypes.Time     `json:"sunset"`
		Visibility        valueTypes.Float    `json:"visibility"`
		WeatherDesc       valueTypes.String   `json:"weather_desc"`
		WeatherURL        valueTypes.String   `json:"weather_url"`
		PsKnowledge       valueTypes.String   `json:"ps_knowledge"`
	} `json:"areaForcastList" PointId:"area_forcast_list" DataTable:"true"`
	StationsCityCode []struct {
		PsId   valueTypes.PsId `json:"ps_id"`
		PsName valueTypes.String  `json:"ps_name"`
		City   valueTypes.String  `json:"city"`
	} `json:"stationsCityCode" PointId:"stations_city_code" DataTable:"true"`
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
