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
	BatteryPlateCom       string               `json:"batteryPlateCom"`
	EnvironmentCom        string               `json:"environmentCom"`
	IsPlatformDefaultUnit valueTypes.Bool      `json:"is_platform_default_unit"`
	Normalization         string               `json:"normalization"`
	Todayvalue            string               `json:"todayvalue"`
	Todayweather          string               `json:"todayweather"`
	TotalAllPower         valueTypes.UnitValue `json:"totalAllPower"`
	Yestedayvalue         string               `json:"yestedayvalue"`
	Yestedayweather       string               `json:"yestedayweather"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
