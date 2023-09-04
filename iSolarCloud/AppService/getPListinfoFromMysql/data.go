package getPListinfoFromMysql

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPListinfoFromMysql"
const Disabled = false
const EndPointName = "AppService.getPListinfoFromMysql"
// ./goraw.sh AppService.getPListinfoFromMysql '{"psIds":1171348}'

type RequestData struct {
	PsIds      valueTypes.PsIds   `json:"psIds" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	NowCapacity struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.Float  `json:"value" PointUnitFrom:"Unit"`
	} `json:"nowCapacity" PointId:"now_capacity"`
	TodayPower struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.Float  `json:"value" PointUnitFrom:"Unit"`
	} `json:"todayPower" PointId:"today_power"`
	TotalCapacity struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.Float  `json:"value" PointUnitFrom:"Unit"`
	} `json:"totalCapacity" PointId:"total_capacity"`
	TotalPower struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.Float  `json:"value" PointUnitFrom:"Unit"`
	} `json:"totalPower" PointId:"total_power"`
	TotalStation valueTypes.Integer `json:"totalStation" PointId:"total_station"`
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
