package getPsUser

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/userService/getPsUser"
const Disabled = false
const EndPointName = "AppService.getPsUser"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BackgroundColor              valueTypes.Integer  `json:"background_color"`
	BloodType                    valueTypes.String   `json:"blood_type"`
	CountryId                    valueTypes.Integer  `json:"country_id"`
	CountryName                  valueTypes.String   `json:"country_name"`
	CreateDate                   valueTypes.DateTime `json:"create_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	CreateUserId                 valueTypes.Integer  `json:"create_user_id"`
	DiseaseHis                   valueTypes.String   `json:"disease_his"`
	Email                        valueTypes.String   `json:"email"`
	EnglishName                  valueTypes.String   `json:"english_name"`
	IsAgreeGdpr                  valueTypes.Bool     `json:"is_agree_gdpr"`
	IsCanModifyUserAccount       valueTypes.Bool     `json:"is_can_modify_user_account"`
	IsDst                        valueTypes.Bool     `json:"is_dst"`
	IsGdpr                       valueTypes.Bool     `json:"is_gdpr"`
	IsNewVersion                 valueTypes.Bool     `json:"is_new_version"`
	IsOpenProtocol               valueTypes.Bool     `json:"is_open_protocol"`
	IsReceiveNotice              valueTypes.Bool     `json:"is_receive_notice"`
	IsSharePosition              valueTypes.Bool     `json:"is_share_position"`
	IsValidMobileEmail           valueTypes.Bool     `json:"is_valid_mobile_email"`
	LackOfInformation            valueTypes.Integer  `json:"lack_of_information"`
	Language                     valueTypes.String   `json:"language"`
	MobileTel                    valueTypes.String   `json:"moble_tel" PointId:"mobile_tel"`
	OrgId                        valueTypes.Integer  `json:"org_id"`
	PhotoId                      valueTypes.Integer  `json:"photo_id"`
	PhotoURL                     valueTypes.String   `json:"photo_url"`
	RootOrgId                    valueTypes.Integer  `json:"root_org_id"`
	Sex                          valueTypes.String   `json:"sex"`
	Stylename                    valueTypes.String   `json:"stylename"`
	TimeZoneId                   valueTypes.Integer  `json:"time_zone_id"`
	Timezone                     valueTypes.String   `json:"timezone"`
	UserAccount                  valueTypes.String   `json:"user_account"`
	UserAccountModifyCount       valueTypes.Integer  `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes valueTypes.Integer  `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            valueTypes.String   `json:"user_dealer_org_code"`
	UserId                       valueTypes.Integer  `json:"user_id"`
	UserLanguage                 valueTypes.String   `json:"user_language"`
	UserLevel                    valueTypes.Integer  `json:"user_level"`
	UserName                     valueTypes.String   `json:"user_name"`
	UserOrgInfo                  struct {
		DealerOrgCode                 valueTypes.String  `json:"dealer_org_code"`
		Email                         valueTypes.String  `json:"email"`
		Installer                     valueTypes.String  `json:"installer"`
		InstallerEmail                valueTypes.String  `json:"installer_email"`
		InstallerPhone                valueTypes.String  `json:"installer_phone"`
		IsCountryDefaultDealerOrgCode valueTypes.Bool    `json:"is_country_default_dealer_org_code"`
		MobileTel                     valueTypes.String  `json:"moble_tel" PointId:"mobile_tel"`
		OrgId                         valueTypes.Integer `json:"org_id"`
		OrgIndexCode                  valueTypes.String  `json:"org_index_code"`
		OrgName                       valueTypes.String  `json:"org_name"`
		UpDealerEmail                 valueTypes.String  `json:"up_dealer_email"`
		UpDealerMobileTel             valueTypes.String  `json:"up_dealer_moble_tel" PointId:up_dealer_mobile_tel"`
		UpDealerOrgCode               valueTypes.String  `json:"up_dealer_org_code"`
		UpDealerOrgInstaller          valueTypes.String  `json:"up_dealer_org_installer"`
		UpDealerOrgInstallerEmail     valueTypes.String  `json:"up_dealer_org_installer_email"`
		UpDealerOrgInstallerPhone     valueTypes.String  `json:"up_dealer_org_installer_phone"`
		UpDealerUserName              valueTypes.String  `json:"up_dealer_user_name"`
		UpDealerUserTelNationCode     valueTypes.String  `json:"up_dealer_user_tel_nation_code"`
		UpOrgId                       valueTypes.Integer `json:"up_org_id"`
		UpOrgName                     valueTypes.String  `json:"up_org_name"`
		UserName                      valueTypes.String  `json:"user_name"`
		UserTelNationCode             valueTypes.String  `json:"user_tel_nation_code"`
	} `json:"user_org_info"`
	UserPassword      valueTypes.String  `json:"user_password"`
	UserTelNationCode valueTypes.String  `json:"user_tel_nation_code"`
	ValidFlag         valueTypes.Integer `json:"valid_flag"`
	Voice             valueTypes.Integer `json:"voice"`
	WorkTel           valueTypes.String  `json:"work_tel"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
