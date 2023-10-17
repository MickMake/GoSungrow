package getInstallerInfoByDealerOrgCodeOrId

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/orgService/getInstallerInfoByDealerOrgCodeOrId"
	Disabled     = false
	EndPointName = "AppService.getInstallerInfoByDealerOrgCodeOrId"
)

type RequestData struct {
	DealerOrgCode valueTypes.String  `json:"dealer_org_code"` // required:"true"`
	OrgId         valueTypes.Integer `json:"org_id"`          // required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	UserInfoList []struct {
		DealerOrgCode     valueTypes.String  `json:"dealer_org_code"`
		Email             valueTypes.String  `json:"email"`
		Installer         valueTypes.String  `json:"installer"`
		InstallerEmail    valueTypes.String  `json:"installer_email"`
		InstallerPhone    valueTypes.String  `json:"installer_phone"`
		MobleTel          interface{}        `json:"moble_tel" PointId:"mobile_tel"`
		OrgId             valueTypes.Integer `json:"org_id"`
		OrgName           valueTypes.String  `json:"org_name"`
		UserId            valueTypes.Integer `json:"user_id"`
		UserName          valueTypes.String  `json:"user_name"`
		UserTelNationCode interface{}        `json:"user_tel_nation_code"`
	} `json:"user_info_list" DataTable:"true"`
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
