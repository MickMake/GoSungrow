package queryPsStructureList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryPsStructureList"
const Disabled = false
const EndPointName = "AppService.queryPsStructureList"

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
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"PsKey"`
	GoStruct       GoStruct.GoStruct        `json:"-" PointDeviceFrom:"PsId"`

	PsId          valueTypes.Integer `json:"ps_id"`
	PsName        valueTypes.String  `json:"ps_name"`
	PsKey         valueTypes.String  `json:"ps_key"`
	DeviceName    valueTypes.String  `json:"device_name"`
	DeviceType    valueTypes.Integer `json:"device_type"`
	IsVirtualUnit valueTypes.Integer `json:"is_virtual_unit"`
	UUID          valueTypes.Integer `json:"uuid"`
	UpUUID        valueTypes.Integer `json:"up_uuid"`
	UUIDIndexCode valueTypes.String  `json:"uuid_index_code"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
