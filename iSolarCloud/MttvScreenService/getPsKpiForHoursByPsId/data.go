package getPsKpiForHoursByPsId

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsKpiForHoursByPsId"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId     `json:"ps_id" required:"true"`
	Day  valueTypes.DateTime `json:"day" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Hours map[string]Hour `json:"hours" DataTable:"true" DataTableIndex:"true"`	// DataTableSortOn:"Index"`
}

type Hour struct {
	CurrPower         valueTypes.Float  `json:"curr_power" PointUnitFrom:"CurrPowerUnit"`
	CurrPowerUnit     valueTypes.String `json:"curr_power_unit" PointIgnore:"true"`
	CurrRadiation     valueTypes.Float  `json:"curr_radiation" PointUnitFrom:"CurrRadiationUnit"`
	CurrRadiationUnit valueTypes.String `json:"curr_radiation_unit" PointIgnore:"true"`
	TodayEnergy       valueTypes.Float  `json:"today_energy" PointUnitFrom:"TodayEnergyUnit"`
	TodayEnergyUnit   valueTypes.String `json:"today_energy_unit" PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		err = json.Unmarshal(data, &e.Hours)
	}

	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
