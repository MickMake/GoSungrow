package getKpiInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getKpiInfo"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ActualEnergy             []valueTypes.Float   `json:"actual_energy" PointUnitFrom:"ActualEnergyUnit"`
	ActualEnergyUnit         valueTypes.String    `json:"actual_energy_unit" PointIgnore:"true"`
	ChargeTotalEnergy        valueTypes.Float     `json:"charge_total_energy" PointUnitFrom:"ChargeTotalEnergyUnit" PointUpdateFreq:"UpdateFreqTotal"`
	ChargeTotalEnergyUnit    valueTypes.String    `json:"charge_total_energy_unit" PointIgnore:"true"`
	DisChargeTotalEnergy     valueTypes.Float     `json:"dis_charge_total_energy" PointUnitFrom:"DisChargeTotalEnergyUnit" PointUpdateFreq:"UpdateFreqTotal"`
	DisChargeTotalEnergyUnit valueTypes.String    `json:"dis_charge_total_energy_unit" PointIgnore:"true"`
	MonthEnergy              valueTypes.UnitValue `json:"month_energy" PointUpdateFreq:"UpdateFreqMonth"`
	OrgName                  valueTypes.String    `json:"org_name"`
	P83024                   valueTypes.Float     `json:"p83024"`
	PercentPlanMonth         valueTypes.Float     `json:"percent_plan_month" PointUnit:"%" PointUpdateFreq:"UpdateFreqMonth"`
	PercentPlanYear          valueTypes.Float     `json:"percent_plan_year" PointUnit:"%" PointUpdateFreq:"UpdateFreqYear"`
	PlanEnergy               []valueTypes.Float   `json:"plan_energy" PointUnitFrom:"PlanEnergyUnit"`
	PlanEnergyUnit           valueTypes.String    `json:"plan_energy_unit" PointIgnore:"true"`
	PsCount                  valueTypes.Integer   `json:"ps_count"`
	TodayEnergy              valueTypes.UnitValue `json:"today_energy" PointUpdateFreq:"UpdateFreqTotal"`
	TotalCapacity            valueTypes.UnitValue `json:"total_capcity" PointId:"total_capacity" PointUpdateFreq:"UpdateFreqTotal"`
	TotalDesignCapacity      valueTypes.UnitValue `json:"total_design_capacity" PointUpdateFreq:"UpdateFreqTotal"`
	TotalEnergy              valueTypes.UnitValue `json:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`
	YearEnergy               valueTypes.UnitValue `json:"year_energy" PointUpdateFreq:"UpdateFreqYear"`
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
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		entries.StructToPoints(e.Response.ResultData, pkg, "", dt)

		_ = entries.CopyPoint(pkg + ".p83024", "virtual.system", "p83024", "")
	}

	return entries
}
