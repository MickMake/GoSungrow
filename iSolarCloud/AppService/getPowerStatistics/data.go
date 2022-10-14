package getPowerStatistics

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStatistics"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("ps_id: Can be fetched from getPsList")
	return ret
}

type ResultData struct {
	PRVlaue        valueTypes.Float     `json:"PRVlaue" PointId:"PRValue"`
	City           interface{}   `json:"city"`
	DayPower       valueTypes.UnitValue `json:"dayPower"`
	DesignCapacity valueTypes.UnitValue `json:"design_capacity"`
	EqVlaue        valueTypes.Float     `json:"eqVlaue" PointId:"EQValue"`
	NowCapacity    valueTypes.UnitValue `json:"nowCapacity"`
	PsName         valueTypes.String    `json:"ps_name"`
	PsShortName    valueTypes.String    `json:"ps_short_name"`
	Status1        valueTypes.Integer   `json:"status1"`
	Status2        valueTypes.Integer   `json:"status2"`
	Status3        valueTypes.Integer   `json:"status3"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
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
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e) + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, pkg, e.Request.PsId.String(), valueTypes.NewDateTime(""))
	}

	return entries
}
