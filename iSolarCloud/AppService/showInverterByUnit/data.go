package showInverterByUnit

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/showInverterByUnit"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

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

	for range Only.Once {
		entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
