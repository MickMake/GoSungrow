package getKpiInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/powerStationService/getKpiInfo"
const Disabled = false
const EndPointName = "AppService.getKpiInfo"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ActualEnergy             []valueTypes.Float   `json:"actual_energy" PointUnitFrom:"ActualEnergyUnit" DataTable:"true" DataTableMerge:"true" DataTableIndex:"true"`
	ActualEnergyUnit         valueTypes.String    `json:"actual_energy_unit" PointIgnore:"true"`
	PlanEnergy               []valueTypes.Float   `json:"plan_energy" PointUnitFrom:"PlanEnergyUnit" DataTable:"true" DataTableMerge:"true" DataTableIndex:"true"`
	PlanEnergyUnit           valueTypes.String    `json:"plan_energy_unit" PointIgnore:"true"`

	OrgName                  valueTypes.String    `json:"org_name"`
	PsCount                  valueTypes.Integer   `json:"ps_count"`
	TotalCapacity            valueTypes.UnitValue `json:"total_capcity" PointId:"total_capacity" PointUpdateFreq:"UpdateFreqTotal"`
	TotalDesignCapacity      valueTypes.UnitValue `json:"total_design_capacity" PointUpdateFreq:"UpdateFreqTotal"`

	P83024                   valueTypes.Float     `json:"p83024" PointUnit:"Wh" PointVirtual:"true"`
	TodayEnergy              valueTypes.UnitValue `json:"today_energy" PointUpdateFreq:"UpdateFreqTotal"`
	MonthEnergy              valueTypes.UnitValue `json:"month_energy" PointUpdateFreq:"UpdateFreqMonth"`
	YearEnergy               valueTypes.UnitValue `json:"year_energy" PointUpdateFreq:"UpdateFreqYear"`
	TotalEnergy              valueTypes.UnitValue `json:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`

	ChargeTotalEnergy        valueTypes.Float     `json:"charge_total_energy" PointUnitFrom:"ChargeTotalEnergyUnit" PointUpdateFreq:"UpdateFreqTotal"`
	ChargeTotalEnergyUnit    valueTypes.String    `json:"charge_total_energy_unit" PointIgnore:"true"`
	DisChargeTotalEnergy     valueTypes.Float     `json:"dis_charge_total_energy" PointUnitFrom:"DisChargeTotalEnergyUnit" PointUpdateFreq:"UpdateFreqTotal"`
	DisChargeTotalEnergyUnit valueTypes.String    `json:"dis_charge_total_energy_unit" PointIgnore:"true"`
	PercentPlanYear          valueTypes.Float     `json:"percent_plan_year" PointUnit:"%" PointUpdateFreq:"UpdateFreqYear"`
	PercentPlanMonth         valueTypes.Float     `json:"percent_plan_month" PointUnit:"%" PointUpdateFreq:"UpdateFreqMonth"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
