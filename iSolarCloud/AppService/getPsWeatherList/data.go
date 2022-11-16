package getPsWeatherList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsWeatherList"
const Disabled = false

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
		Chill      valueTypes.Float    `json:"chill"`
		Code       valueTypes.Float    `json:"code"`
		CodeName   valueTypes.String   `json:"code_name"`
		DateTime   valueTypes.DateTime `json:"date_time" PointIgnore:"true"`
		Direction  valueTypes.Float    `json:"direction"`
		High       valueTypes.Float    `json:"high" PointUnit:"F"`
		Highc      valueTypes.Float    `json:"highc" PointUnit:"C"`
		Humidity   valueTypes.Float    `json:"humidity"`
		Low        valueTypes.Float    `json:"low" PointUnit:"F"`
		Lowc       valueTypes.Float    `json:"lowc" PointUnit:"C"`
		Pressure   valueTypes.Float    `json:"pressure" PointUnit:"hPa"`
		PsId       valueTypes.PsId     `json:"ps_id"`
		Rising     valueTypes.Float    `json:"rising"`
		Speed      valueTypes.Float    `json:"speed"`
		Sunrise    valueTypes.Time     `json:"sunrise"`
		Sunset     valueTypes.Time     `json:"sunset"`
		Visibility valueTypes.Float    `json:"visibility"`
	} `json:"weather_list" PointIdFromChild:"DateTime" PointNameDateFormat:"20060102" DataTable:"true"`
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
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
