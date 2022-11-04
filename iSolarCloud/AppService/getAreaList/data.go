package getAreaList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getAreaList"
const Disabled = false

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
	PageList []struct {
		FaultStationCount valueTypes.Integer   `json:"fault_station_count"`
		IsHaveEsPs        valueTypes.Bool      `json:"is_have_es_ps"`
		IsLeaf            valueTypes.Bool      `json:"is_leaf"`
		OrgId             valueTypes.Integer   `json:"org_id"`
		OrgName           valueTypes.String    `json:"org_name"`
		P83048            valueTypes.UnitValue `json:"p83048"`
		P83049            valueTypes.UnitValue `json:"p83049"`
		P83050            valueTypes.UnitValue `json:"p83050"`
		P83051            valueTypes.UnitValue `json:"p83051"`
		PlanMonth         valueTypes.String    `json:"plan_month"`
		StationCount      valueTypes.Integer   `json:"station_count"`
		TodayEnergy       valueTypes.UnitValue `json:"today_energy"`
		TotalEnergy       valueTypes.UnitValue `json:"total_energy"`
	} `json:"pageList" PointId:"page_list" PointNameFromChild:"OrgId" PointNameAppend:"false" PointArrayFlatten:"false" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count" PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
