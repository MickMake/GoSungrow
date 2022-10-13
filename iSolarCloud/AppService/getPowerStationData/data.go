package getPowerStationData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationData"
const Disabled = false

type RequestData struct {
	PsId     api.Integer `json:"ps_id" required:"true"`
	DateID   string `json:"date_id" required:"true"`
	DateType string `json:"date_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += fmt.Sprintln("date_type: Day = 1")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("date_type: Month = 2")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYYmm")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYYmm")
	ret += fmt.Sprintln("date_type: Year = 3")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYY")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYY")
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	DayPowerQuantityTotal api.UnitValue `json:"day_power_quantity_total"`
	DayTotalProfit        api.UnitValue `json:"day_total_profit"`
	P34048List            []api.Float   `json:"p34048List" PointId:"p34048" PointUnitFrom:"p34048_unit"`
	P34048Unit            string        `json:"p34048_unit"`
	P83012List            []api.Float   `json:"p83012List" PointId:"p83012" PointUnitFrom:"p83012_unit"`
	P83012Unit            string        `json:"p83012_unit"`
	P83033List            []api.Float   `json:"p83033List" PointId:"p83033" PointUnitFrom:"p83033_unit"`
	P83033Unit            string        `json:"p83033_unit"`
	P83106List            []api.Float   `json:"p83106List" PointId:"p83106" PointUnitFrom:"p83106_unit"`
	P83106Unit            string        `json:"p83106_unit"`

	P83022List []api.Float `json:"p83022List" PointId:"p83022" PointUnitFrom:"p83022_unit"` // Used for Year, Month, Day
	P83022Unit string      `json:"p83022_unit"`                                             // Used for Year, Month, Day
	P83118List []api.Float `json:"p83118List" PointId:"p83118" PointUnitFrom:"p83118_unit"` // Used for Year, Month, Day
	P83118Unit string      `json:"p83118_unit"`                                             // Used for Year, Month, Day
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		name := apiReflect.GetName("", *e) + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), api.NewDateTime(""))

		// uv := api.SetUnitValueString(e.Response.ResultData.DayPowerQuantityTotal.Value, e.Response.ResultData.DayPowerQuantityTotal.Unit)
		// entries.AddUnitValue(name + ".DayPowerQuantityTotal", e.Request.PsId.String(), "DayPowerQuantityTotal", "", "", api.NewDateTime(""), uv)
		//
		// uv = api.SetUnitValueString(e.Response.ResultData.DayTotalProfit.Value, e.Response.ResultData.DayTotalProfit.Unit)
		// entries.AddUnitValue(name + ".DayTotalProfit", e.Request.PsId.String(), "DayTotalProfit", "", "", api.NewDateTime(""), uv)
		//
		// for _, d := range e.Response.ResultData.P34048List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P34048Unit)
		// 	entries.AddUnitValue(name + ".p34048", e.Request.PsId.String(), "p34048", "", "", api.NewDateTime(""), uv)
		// }
		//
		// for _, d := range e.Response.ResultData.P83012List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P83012Unit)
		// 	entries.AddUnitValue(name + ".p83012", e.Request.PsId.String(), "p83012", "", "", api.NewDateTime(""), uv)
		// }
		//
		// for _, d := range e.Response.ResultData.P83033List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P83033Unit)
		// 	entries.AddUnitValue(name + ".p83033", e.Request.PsId.String(), "p83033", "", "", api.NewDateTime(""), uv)
		// }
		//
		// for _, d := range e.Response.ResultData.P83106List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P83106Unit)
		// 	entries.AddUnitValue(name + ".p83106", e.Request.PsId.String(), "p83106", "", "", api.NewDateTime(""), uv)
		// }
		//
		// for _, d := range e.Response.ResultData.P83022List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P83022Unit)
		// 	entries.AddUnitValue(name + ".p83022", e.Request.PsId.String(), "p83022", "", "", api.NewDateTime(""), uv)
		// }
		//
		// for _, d := range e.Response.ResultData.P83118List {
		// 	uv = api.SetUnitValueString(d, e.Response.ResultData.P83118Unit)
		// 	entries.AddUnitValue(name + ".p83118", e.Request.PsId.String(), "p83118", "", "", api.NewDateTime(""), uv)
		// }
	}

	return entries
}
