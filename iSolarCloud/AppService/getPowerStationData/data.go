package getPowerStationData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"time"
)

const Url = "/v1/powerStationService/getPowerStationData"
const Disabled = false

type RequestData struct {
	PsId     valueTypes.PsId `json:"ps_id" required:"true"`
	DateId   valueTypes.DateTime `json:"date_id" required:"true"`
	DateType valueTypes.String `json:"date_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += api.HelpDateId()
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	GoStructParent        GoStruct.GoStructParent `json:"-" DataTable:"true"`

	DayPowerQuantityTotal valueTypes.UnitValue `json:"day_power_quantity_total" PointUpdateFreq:"UpdateFreqTotal"`
	DayTotalProfit        valueTypes.UnitValue `json:"day_total_profit" PointUpdateFreq:"UpdateFreqTotal"`

	P34048List            []valueTypes.Float   `json:"p34048List" PointId:"p34048" PointUnitFrom:"P34048Unit" PointVirtual:"true"`
	P34048Unit            valueTypes.String    `json:"p34048_unit" PointIgnore:"true"`
	P83012List            []valueTypes.Float   `json:"p83012List" PointId:"p83012" PointUnitFrom:"P83012Unit" PointVirtual:"true"`
	P83012Unit            valueTypes.String    `json:"p83012_unit" PointIgnore:"true"`
	P83033List            []valueTypes.Float   `json:"p83033List" PointId:"p83033" PointUnitFrom:"P83033Unit" PointVirtual:"true"`
	P83033Unit            valueTypes.String    `json:"p83033_unit" PointIgnore:"true"`
	P83106List            []valueTypes.Float   `json:"p83106List" PointId:"p83106" PointUnitFrom:"P83106Unit" PointVirtual:"true"`
	P83106Unit            valueTypes.String    `json:"p83106_unit" PointIgnore:"true"`
	P83022List            []valueTypes.Float   `json:"p83022List" PointId:"p83022" PointUnitFrom:"P83022Unit" PointVirtual:"true"`	// Used for Year, Month, Day
	P83022Unit            valueTypes.String    `json:"p83022_unit" PointIgnore:"true"`                         						// Used for Year, Month, Day
	P83118List            []valueTypes.Float   `json:"p83118List" PointId:"p83118" PointUnitFrom:"P83118Unit" PointVirtual:"true"`	// Used for Year, Month, Day
	P83118Unit            valueTypes.String    `json:"p83118_unit" PointIgnore:"true"`                         						// Used for Year, Month, Day
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

	me := entries.GetReflect("ResultData")

	// Find max # of values in []valueTypes.Float.
	var max int
	for index := range me.ChildReflect {
		if me.ChildReflect[index].Length > max {
			max = me.ChildReflect[index].Length
		}
	}

	var r []string
	var t string
	switch {
		case e.Request.DateType.Match(valueTypes.DateTypeDay):
			count := len(e.Response.ResultData.P34048List)
			dur := time.Duration(int(time.Hour * 24) / count)
			r = e.Request.DateId.GetRanges(count, dur, valueTypes.DateTimeLayout)
			t = "Hour"

		case e.Request.DateType.Match(valueTypes.DateTypeMonth):
			count := len(e.Response.ResultData.P83022List)
			r = e.Request.DateId.GetRanges(count, time.Hour * 24, valueTypes.DateLayout)
			t = "Day"

		case e.Request.DateType.Match(valueTypes.DateTypeYear):
			count := len(e.Response.ResultData.P83022List)
			r = e.Request.DateId.GetRanges(count, time.Hour * 24 * 31, "2006-01")
			t = "Month"

		case e.Request.DateType.Match(valueTypes.DateTypeTotal):
			// Not support for getPowerStationData.
			// 	count := len(e.Response.ResultData.P83022List)
			// 	r = e.Request.DateId.GetRanges(count, time.Hour * 24, valueTypes.DateLayout)
			// 	t = "Year"
	}
	me.SetDataTableIndexNames(r...)
	me.SetDataTableIndexTitle(t)

	return entries
}
