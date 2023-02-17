package getPowerTrendDayData

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"

	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerTrendDayData"
const Disabled = false
const EndPointName = "WebAppService.getPowerTrendDayData"

type RequestData struct {
	BeginTime2 valueTypes.String `json:"beginTime" required:"true"`
	EndTime2   valueTypes.String `json:"endTime" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	// Dummy valueTypes.String `json:"dummy"`
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
