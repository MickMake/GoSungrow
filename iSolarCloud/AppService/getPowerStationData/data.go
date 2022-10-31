package getPowerStationData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationData"
const Disabled = false

type RequestData struct {
	PsId     valueTypes.PsId `json:"ps_id" required:"true"`
	DateId   valueTypes.DateTime `json:"date_id" required:"true"`
	DateType valueTypes.String `json:"date_type" required:"true"`
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
	DayPowerQuantityTotal valueTypes.UnitValue `json:"day_power_quantity_total" PointUpdateFreq:"UpdateFreqTotal"`
	DayTotalProfit        valueTypes.UnitValue `json:"day_total_profit" PointUpdateFreq:"UpdateFreqTotal"`
	P34048List            []valueTypes.Float   `json:"p34048List" PointId:"p34048" PointUnitFrom:"P34048Unit"`
	P34048Unit            valueTypes.String    `json:"p34048_unit" PointIgnore:"true"`
	P83012List            []valueTypes.Float   `json:"p83012List" PointId:"p83012" PointUnitFrom:"P83012Unit"`
	P83012Unit            valueTypes.String    `json:"p83012_unit" PointIgnore:"true"`
	P83033List            []valueTypes.Float   `json:"p83033List" PointId:"p83033" PointUnitFrom:"P83033Unit"`
	P83033Unit            valueTypes.String    `json:"p83033_unit" PointIgnore:"true"`
	P83106List            []valueTypes.Float   `json:"p83106List" PointId:"p83106" PointUnitFrom:"P83106Unit"`
	P83106Unit            valueTypes.String    `json:"p83106_unit" PointIgnore:"true"`
	P83022List            []valueTypes.Float   `json:"p83022List" PointId:"p83022" PointUnitFrom:"P83022Unit"` // Used for Year, Month, Day
	P83022Unit            valueTypes.String    `json:"p83022_unit" PointIgnore:"true"`                         // Used for Year, Month, Day
	P83118List            []valueTypes.Float   `json:"p83118List" PointId:"p83118" PointUnitFrom:"P83118Unit"` // Used for Year, Month, Day
	P83118Unit            valueTypes.String    `json:"p83118_unit" PointIgnore:"true"`                         // Used for Year, Month, Day
}

func (e *ResultData) IsValid() error {
	var err error
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
