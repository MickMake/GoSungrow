package getKpiInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
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
	ActualEnergy             []string      `json:"actual_energy"`
	ActualEnergyUnit         string        `json:"actual_energy_unit"`
	ChargeTotalEnergy        string        `json:"charge_total_energy"`
	ChargeTotalEnergyUnit    string        `json:"charge_total_energy_unit"`
	DisChargeTotalEnergy     string        `json:"dis_charge_total_energy"`
	DisChargeTotalEnergyUnit string        `json:"dis_charge_total_energy_unit"`
	MonthEnergy              api.UnitValue `json:"month_energy"`
	OrgName                  string        `json:"org_name"`
	P83024                   float64       `json:"p83024"`
	PercentPlanMonth         string        `json:"percent_plan_month"`
	PercentPlanYear          string        `json:"percent_plan_year"`
	PlanEnergy               []string      `json:"plan_energy"`
	PlanEnergyUnit           string        `json:"plan_energy_unit"`
	PsCount                  int64         `json:"ps_count"`
	TodayEnergy              api.UnitValue `json:"today_energy"`
	TotalCapcity             api.UnitValue `json:"total_capcity"`
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
