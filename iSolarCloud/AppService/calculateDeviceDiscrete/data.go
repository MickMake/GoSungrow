package calculateDeviceDiscrete

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"

	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/calculateDeviceDiscrete"
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
