package queryPsStructureList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryPsStructureList"
const Disabled = false

type RequestData struct {
	PsId      valueTypes.PsId   `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	DeviceName    valueTypes.String  `json:"device_name"`
	DeviceType    valueTypes.Integer `json:"device_type"`
	IsVirtualUnit valueTypes.Integer `json:"is_virtual_unit"`
	PsID          valueTypes.Integer `json:"ps_id"`
	PsKey         valueTypes.String  `json:"ps_key"`
	PsName        valueTypes.String  `json:"ps_name"`
	UpUUID        valueTypes.Integer `json:"up_uuid"`
	UUID          valueTypes.Integer `json:"uuid"`
	UUIDIndexCode valueTypes.String  `json:"uuid_index_code"`
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
