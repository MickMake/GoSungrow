package getKpiInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"sort"
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
	ActualEnergy             []string      `json:"actual_energy"`
	ActualEnergyUnit         string        `json:"actual_energy_unit"`
	ChargeTotalEnergy        string        `json:"charge_total_energy"`
	ChargeTotalEnergyUnit    string        `json:"charge_total_energy_unit"`
	DisChargeTotalEnergy     string        `json:"dis_charge_total_energy"`
	DisChargeTotalEnergyUnit string        `json:"dis_charge_total_energy_unit"`
	MonthEnergy              api.UnitValue `json:"month_energy"`
	OrgName                  string        `json:"org_name"`
	P83024                   api.Float     `json:"p83024"`
	PercentPlanMonth         string        `json:"percent_plan_month"`
	PercentPlanYear          string        `json:"percent_plan_year"`
	PlanEnergy               []string      `json:"plan_energy"`
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Date",
			"Point Id",
			// "Parents",
			"Group Name",
			"Description",
			"Value",
			"Unit",
		)

		data := e.GetData()
		var sorted []string
		for p := range data.DataPoints {
			sorted = append(sorted, string(p))
		}
		sort.Strings(sorted)

		for _, p := range sorted {
			entries := data.DataPoints[api.PointId(p)]
			for _, de := range entries {
				if de.Hide {
					continue
				}

				_ = table.AddRow(
					de.Date.Format(api.DtLayout),
					// api.NameDevicePointInt(de.Point.Parents, p.PointID.Value()),
					// de.Point.Id,
					p,
					// de.Point.Parents.String(),
					de.Point.GroupName,
					de.Point.Name,
					de.Value,
					de.Point.Unit,
				)
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(4),
		// 	ValueColumn:  output.SetInteger(5),
		// 	UnitsColumn:  output.SetInteger(6),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}
	return table
}
