package getPsTreeMenu

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getPsTreeMenu"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPsTreeMenu"
)

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	List []Ps `json:"list" DataTable:"true"`
}
type Ps struct {
	PsKey         valueTypes.PsKey   `json:"ps_key"`
	PsId          valueTypes.PsId    `json:"ps_id"`
	PsName        valueTypes.String  `json:"ps_name"`
	DevPsId       valueTypes.Integer `json:"dev_ps_id"`
	DeviceType    valueTypes.Integer `json:"device_type"`
	DeviceName    valueTypes.String  `json:"device_name"`
	IsPublic      valueTypes.Bool    `json:"is_public"`
	IsVirtualUnit valueTypes.Bool    `json:"is_virtual_unit"`
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
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
