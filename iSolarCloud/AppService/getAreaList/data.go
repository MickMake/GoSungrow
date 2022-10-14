package getAreaList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/powerStationService/getAreaList"
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
	PageList []struct {
		FaultStationCount valueTypes.Integer   `json:"fault_station_count"`
		IsHaveEsPs        valueTypes.Bool      `json:"is_have_es_ps"`
		IsLeaf            valueTypes.Bool      `json:"is_leaf"`
		OrgID             valueTypes.Integer   `json:"org_id"`
		OrgName           valueTypes.String    `json:"org_name"`
		P83048            valueTypes.UnitValue `json:"p83048"`
		P83049            valueTypes.UnitValue `json:"p83049"`
		P83050            valueTypes.UnitValue `json:"p83050"`
		P83051            valueTypes.UnitValue `json:"p83051"`
		PlanMonth         string        `json:"plan_month"`
		StationCount      valueTypes.Integer   `json:"station_count"`
		TodayEnergy       valueTypes.UnitValue `json:"today_energy"`
		TotalEnergy       valueTypes.UnitValue `json:"total_energy"`
	} `json:"pageList"`
	RowCount valueTypes.Integer `json:"rowCount"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
