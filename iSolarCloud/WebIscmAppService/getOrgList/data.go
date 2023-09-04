package getOrgList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/orgService/getOrgList"
const Disabled = false
const EndPointName = "WebIscmAppService.getOrgList"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent          GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"OrgId"`

	OrgId                   valueTypes.Integer  `json:"org_id"`
	OrgName                 valueTypes.String   `json:"org_name"`
	OrgType                 valueTypes.String   `json:"org_type"`
	OrgScale                valueTypes.String   `json:"org_scale"`
	AreaCode                interface{}         `json:"area_code"`
	BankName                valueTypes.String   `json:"bank_name"`
	BankNum                 valueTypes.String   `json:"bank_num"`
	BankOrgNum              valueTypes.String   `json:"bank_org_num"`
	Business                valueTypes.String   `json:"business"`
	BusinessLicence         valueTypes.String   `json:"business_licence"`
	BusinessScope           valueTypes.String   `json:"business_scope"`
	ChildSize               valueTypes.Integer  `json:"childsize" PointId:"child_size"`
	CorporateIdentityNum    valueTypes.String   `json:"corporate_identity_num"`
	CorporateIdentityType   valueTypes.String   `json:"corporate_identity_type"`
	CorporateName           valueTypes.String   `json:"corporate_name"`
	Country                 valueTypes.String   `json:"country"`
	CountryEnUs             valueTypes.String   `json:"country_en_us"`
	Currency                valueTypes.String   `json:"currency"`
	FaxNo                   valueTypes.String   `json:"fax_no"`
	HotLine                 valueTypes.String   `json:"hot_line"`
	Introduction            valueTypes.String   `json:"introduction"`
	Logo                    interface{}         `json:"logo"`
	NetProperty             interface{}         `json:"net_property"`
	OrganizationCertificate valueTypes.String   `json:"organization_certificate"`
	PostalCode              valueTypes.String   `json:"postalcode" PointId:"postal_code"`
	RegisterNum             valueTypes.String   `json:"register_num"`
	RegisteredAddress       valueTypes.String   `json:"registered_address"`
	RegisteredArea          valueTypes.String   `json:"registered_area"`
	RegisteredCapital       valueTypes.Integer  `json:"registered_capital"`
	RegisteredDate          valueTypes.DateTime `json:"registered_date" PointNameDateFormat:"DateTimeLayout"`
	ServiceEmail            valueTypes.String   `json:"service_emaill" PointId:"service_email"`
	TaxCertificate          valueTypes.String   `json:"tax_certificate"`
	TelNo                   valueTypes.String   `json:"tel_no"`
	TimezoneID              valueTypes.String   `json:"timezone_id"`
	WebSite                 valueTypes.String   `json:"web_site"`
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
