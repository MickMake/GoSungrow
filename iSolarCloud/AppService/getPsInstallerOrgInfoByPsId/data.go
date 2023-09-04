package getPsInstallerOrgInfoByPsId

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/powerStationService/getPsInstallerOrgInfoByPsId"
const Disabled = false
const EndPointName = "AppService.getPsInstallerOrgInfoByPsId"

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


type ResultData   struct {
	PsOrgInfoList []struct {
		OrgId           valueTypes.Integer `json:"org_id"`
		OrgName         valueTypes.String  `json:"org_name"`
		OrgIndexCode    valueTypes.String  `json:"org_index_code"`
		DealerOrgCode   valueTypes.String  `json:"dealer_org_code"`
		PsDealerOrgCode valueTypes.String  `json:"ps_dealer_org_code"`
		Installer       valueTypes.String  `json:"installer"`
		InstallerEmail  valueTypes.String  `json:"installer_email"`
		InstallerPhone  valueTypes.String  `json:"installer_phone"`
		UpOrgID         valueTypes.Integer `json:"up_org_id"`
	} `json:"ps_org_info_list" DataTable:"true" DataTableSortOn:"PsKey"`
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
