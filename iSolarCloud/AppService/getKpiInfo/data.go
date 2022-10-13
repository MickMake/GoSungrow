package getKpiInfo

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"time"
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
	ActualEnergy             []api.Float   `json:"actual_energy" PointUnitFrom:"actual_energy_unit"`
	ActualEnergyUnit         string        `json:"actual_energy_unit"`
	ChargeTotalEnergy        api.Float     `json:"charge_total_energy" PointUnitFrom:"charge_total_energy_unit"`
	ChargeTotalEnergyUnit    string        `json:"charge_total_energy_unit"`
	DisChargeTotalEnergy     api.Float     `json:"dis_charge_total_energy" PointUnitFrom:"dis_charge_total_energy_unit"`
	DisChargeTotalEnergyUnit string        `json:"dis_charge_total_energy_unit"`
	MonthEnergy              api.UnitValue `json:"month_energy"`
	OrgName                  api.String    `json:"org_name"`
	P83024                   api.Float     `json:"p83024"`
	PercentPlanMonth         api.Float     `json:"percent_plan_month"`
	PercentPlanYear          api.Float     `json:"percent_plan_year"`
	PlanEnergy               []api.Float   `json:"plan_energy" PointUnitFrom:"plan_energy_unit"`
	PlanEnergyUnit           string        `json:"plan_energy_unit"`
	PsCount                  api.Integer   `json:"ps_count"`
	TodayEnergy              api.UnitValue `json:"today_energy"`
	TotalCapcity             api.UnitValue `json:"total_capcity" PointId:"total_capacity"`
	TotalDesignCapacity      api.UnitValue `json:"total_design_capacity"`
	TotalEnergy              api.UnitValue `json:"total_energy"`
	YearEnergy               api.UnitValue `json:"year_energy"`
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

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
// }

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		entries.StructToPoints(e.Response.ResultData, pkg, "system", time.Time{})
	}

	return entries
}
