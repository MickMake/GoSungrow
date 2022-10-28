package showPSView

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/showPSView"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BatteryPlateCom       valueTypes.String    `json:"batteryPlateCom" PointId:"battery_plate_com"`
	EnvironmentCom        valueTypes.String    `json:"environmentCom" PointId:"environment_com"`
	IsPlatformDefaultUnit valueTypes.Bool      `json:"is_platform_default_unit"`
	Normalization         valueTypes.String    `json:"normalization"`
	Todayvalue            valueTypes.String    `json:"todayvalue" PointId:"today_value"`
	Todayweather          valueTypes.Integer   `json:"todayweather" PointId:"today_weather"`
	TotalAllPower         valueTypes.UnitValue `json:"totalAllPower" PointId:"total_all_power"`
	Yestedayvalue         valueTypes.String    `json:"yestedayvalue" PointId:"yesterday_value"`
	Yestedayweather       valueTypes.Integer   `json:"yestedayweather" PointId:"yesterday_weather"`
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
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		name := pkg + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), dt)
	}

	return entries
}
