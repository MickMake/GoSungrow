package getAllPsIdByOrgIds

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getAllPsIdByOrgIds"
	Disabled     = false
	EndPointName = "AppService.getAllPsIdByOrgIds"
)

type RequestData struct {
	OrgIds valueTypes.String `json:"orgIds" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []string

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.OrgIds.String()
		entries.StructToDataMap(*e, e.Request.OrgIds.String(), GoStruct.NewEndPointPath(e.Request.OrgIds.String()))
	}

	return entries
}
