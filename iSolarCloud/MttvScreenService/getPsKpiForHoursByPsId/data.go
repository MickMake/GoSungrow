package getPsKpiForHoursByPsId

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
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
	Hours map[string]Hour `json:"hours"`
}
type Hour struct {
	CurrPower         float64            `json:"curr_power"`
	CurrPowerUnit     valueTypes.String  `json:"curr_power_unit"`
	CurrRadiation     valueTypes.Integer `json:"curr_radiation"`
	CurrRadiationUnit valueTypes.String  `json:"curr_radiation_unit"`
	TodayEnergy       valueTypes.Integer `json:"today_energy"`
	TodayEnergyUnit   valueTypes.String  `json:"today_energy_unit"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
