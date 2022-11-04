package getOrgList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/orgService/getOrgList"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	AreaCode                interface{}        `json:"area_code"`
	BankName                valueTypes.String  `json:"bank_name"`
	BankNum                 valueTypes.String  `json:"bank_num"`
	BankOrgNum              valueTypes.String  `json:"bank_org_num"`
	Business                valueTypes.String  `json:"business"`
	BusinessLicence         valueTypes.String  `json:"business_licence"`
	BusinessScope           valueTypes.String  `json:"business_scope"`
	Childsize               valueTypes.Integer `json:"childsize"`
	CorporateIdentityNum    valueTypes.String  `json:"corporate_identity_num"`
	CorporateIdentityType   valueTypes.String  `json:"corporate_identity_type"`
	CorporateName           valueTypes.String  `json:"corporate_name"`
	Country                 valueTypes.String  `json:"country"`
	CountryEnUs             valueTypes.String  `json:"country_en_us"`
	Currency                valueTypes.String  `json:"currency"`
	FaxNo                   valueTypes.String  `json:"fax_no"`
	HotLine                 valueTypes.String  `json:"hot_line"`
	Introduction            valueTypes.String  `json:"introduction"`
	Logo                    interface{}        `json:"logo"`
	NetProperty             interface{}        `json:"net_property"`
	OrgID                   valueTypes.Integer `json:"org_id"`
	OrgName                 valueTypes.String  `json:"org_name"`
	OrgScale                valueTypes.String  `json:"org_scale"`
	OrgType                 valueTypes.String  `json:"org_type"`
	OrganizationCertificate valueTypes.String  `json:"organization_certificate"`
	Postalcode              valueTypes.String  `json:"postalcode"`
	RegisterNum             valueTypes.String  `json:"register_num"`
	RegisteredAddress       valueTypes.String  `json:"registered_address"`
	RegisteredArea          valueTypes.String  `json:"registered_area"`
	RegisteredCapital       valueTypes.Integer `json:"registered_capital"`
	RegisteredDate          valueTypes.String  `json:"registered_date"`
	ServiceEmaill           valueTypes.String  `json:"service_emaill"`
	TaxCertificate          valueTypes.String  `json:"tax_certificate"`
	TelNo                   valueTypes.String  `json:"tel_no"`
	TimezoneID              valueTypes.String  `json:"timezone_id"`
	WebSite                 valueTypes.String  `json:"web_site"`
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
