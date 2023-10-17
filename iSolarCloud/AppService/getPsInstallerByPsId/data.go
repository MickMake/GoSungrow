package getPsInstallerByPsId

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/orgService/getPsInstallerByPsId"
	Disabled     = false
	EndPointName = "AppService.getPsInstallerByPsId"
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
	PsId           valueTypes.Integer `json:"ps_id"`
	PsType         valueTypes.Integer `json:"ps_type"`
	RootOrgId      valueTypes.Integer `json:"root_org_id"`
	InstallerOrgId valueTypes.Integer `json:"installer_org_id"`
	Installer      valueTypes.String  `json:"installer"`
	InstallerEmail valueTypes.String  `json:"installer_email"`
	InstallerPhone valueTypes.String  `json:"installer_phone"`
	OrgIndexCode   valueTypes.String  `json:"org_index_code"`
	OrgURL         interface{}        `json:"org_url"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
