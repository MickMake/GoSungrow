package showInverterByUnit

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/showInverterByUnit"
	Disabled     = false
	EndPointName = "AppService.showInverterByUnit"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	AndName      valueTypes.String  `json:"and_name"`
	AndUUID      valueTypes.Integer `json:"and_uuid"`
	DeviceType   valueTypes.Integer `json:"device_type"`
	InverterInfo valueTypes.String  `json:"inverterinfo" PointId:"inverter_info"`
	Name         valueTypes.String  `json:"name"`
	UnitName     valueTypes.String  `json:"unit_name"`
	UnitUUID     valueTypes.Integer `json:"unit_uuid"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
