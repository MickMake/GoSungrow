package getPListinfoFromMysql

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPListinfoFromMysql"
const Disabled = false
// ./goraw.sh AppService.getPListinfoFromMysql '{"psIds":1171348}'

type RequestData struct {
	PsIds      valueTypes.PsId   `json:"psIds" required:"true"`
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
		Unit  valueTypes.String `json:"unit"`
		Value valueTypes.String `json:"value"`
	} `json:"nowCapacity"`
	TodayPower struct {
		Unit  valueTypes.String `json:"unit"`
		Value valueTypes.String `json:"value"`
	} `json:"todayPower"`
	TotalCapacity struct {
		Unit  valueTypes.String `json:"unit"`
		Value valueTypes.String `json:"value"`
	} `json:"totalCapacity"`
	TotalPower struct {
		Unit  valueTypes.String `json:"unit"`
		Value valueTypes.String `json:"value"`
	} `json:"totalPower"`
	TotalStation valueTypes.Integer `json:"totalStation"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
