package getVillageList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/commonService/getVillageList"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData []struct {
	Village     string `json:"village"`
	VillageCode string `json:"village_code"`
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
