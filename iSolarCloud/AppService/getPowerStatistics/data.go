package getPowerStatistics

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/powerStationService/getPowerStatistics"
	Disabled     = false
	EndPointName = "AppService.getPowerStatistics"
)

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("ps_id: Can be fetched from getPsList")
	return ret
}

type ResultData struct {
	City        valueTypes.String  `json:"city"`
	PsName      valueTypes.String  `json:"ps_name"`
	PsShortName valueTypes.String  `json:"ps_short_name"`
	Status1     valueTypes.Integer `json:"status1"`
	Status2     valueTypes.Integer `json:"status2"`
	Status3     valueTypes.Integer `json:"status3"`

	DayPower       valueTypes.UnitValue `json:"dayPower" PointId:"day_power"`
	DesignCapacity valueTypes.UnitValue `json:"design_capacity"`
	PRVlaue        valueTypes.Float     `json:"PRVlaue" PointId:"pr_value"`
	EqVlaue        valueTypes.Float     `json:"eqVlaue" PointId:"eq_value"`
	NowCapacity    valueTypes.UnitValue `json:"nowCapacity" PointId:"now_capacity"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
