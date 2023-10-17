package getPsDetailForSinglePage

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/powerStationService/getPsDetailForSinglePage"
	Disabled     = false
	EndPointName = "AppService.getPsDetailForSinglePage"
)

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Co2Reduce struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"co2_reduce"`
	CoalReduce struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"coal_reduce"`
	CoalReduceTotal struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"coal_reduce_total"`
	CurrPower struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"curr_power"`
	DesignCapacity struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"design_capacity"`
	P83023      valueTypes.String `json:"p83023"`
	TodayEnergy struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"today_energy"`
	TodayEnergyVirgin struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"today_energy_virgin"`
	TotalEnergy struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"total_energy"`
	TotalEnergyVirgin struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"total_energy_virgin"`
	TreeReduce struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.String `json:"value" PointUnitFrom:"Unit"`
	} `json:"tree_reduce"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
