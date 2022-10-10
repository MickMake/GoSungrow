package getPsUser

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/userService/getPsUser"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BackgroundColor              api.Integer  `json:"background_color"`
	BloodType                    api.String   `json:"blood_type"`
	CountryID                    api.Integer  `json:"country_id"`
	CountryName                  api.String   `json:"country_name"`
	CreateDate                   api.DateTime `json:"create_date"`
	CreateUserID                 api.Integer  `json:"create_user_id"`
	DiseaseHis                   api.String   `json:"disease_his"`
	Email                        api.String   `json:"email"`
	EnglishName                  api.String   `json:"english_name"`
	IsAgreeGdpr                  api.Bool     `json:"is_agree_gdpr"`
	IsCanModifyUserAccount       api.Bool     `json:"is_can_modify_user_account"`
	IsDst                        api.Bool     `json:"is_dst"`
	IsGdpr                       api.Bool     `json:"is_gdpr"`
	IsNewVersion                 api.Bool     `json:"is_new_version"`
	IsOpenProtocol               api.Bool     `json:"is_open_protocol"`
	IsReceiveNotice              api.Bool     `json:"is_receive_notice"`
	IsSharePosition              api.Bool     `json:"is_share_position"`
	IsValidMobileEmail           api.Bool     `json:"is_valid_mobile_email"`
	LackOfInformation            api.Integer  `json:"lack_of_information"`
	Language                     api.String   `json:"language"`
	MobileTel                    api.String   `json:"moble_tel"`
	OrgID                        api.Integer  `json:"org_id"`
	PhotoID                      api.Integer  `json:"photo_id"`
	PhotoURL                     api.String   `json:"photo_url"`
	RootOrgID                    api.Integer  `json:"root_org_id"`
	Sex                          api.String   `json:"sex"`
	Stylename                    api.String   `json:"stylename"`
	TimeZoneID                   api.Integer  `json:"time_zone_id"`
	Timezone                     api.String   `json:"timezone"`
	UserAccount                  api.String   `json:"user_account"`
	UserAccountModifyCount       api.Integer  `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes api.Integer  `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            api.String   `json:"user_dealer_org_code"`
	UserID                       api.Integer  `json:"user_id"`
	UserLanguage                 api.String   `json:"user_language"`
	UserLevel                    api.Integer  `json:"user_level"`
	UserName                     api.String   `json:"user_name"`
	UserOrgInfo                  struct {
		DealerOrgCode                 api.String  `json:"dealer_org_code"`
		Email                         api.String  `json:"email"`
		Installer                     api.String  `json:"installer"`
		InstallerEmail                api.String  `json:"installer_email"`
		InstallerPhone                api.String  `json:"installer_phone"`
		IsCountryDefaultDealerOrgCode api.Bool    `json:"is_country_default_dealer_org_code"`
		MobileTel                     api.String  `json:"moble_tel"`
		OrgID                         api.Integer `json:"org_id"`
		OrgIndexCode                  api.String  `json:"org_index_code"`
		OrgName                       api.String  `json:"org_name"`
		UpDealerEmail                 api.String  `json:"up_dealer_email"`
		UpDealerMobileTel             api.String  `json:"up_dealer_moble_tel"`
		UpDealerOrgCode               api.String  `json:"up_dealer_org_code"`
		UpDealerOrgInstaller          api.String  `json:"up_dealer_org_installer"`
		UpDealerOrgInstallerEmail     api.String  `json:"up_dealer_org_installer_email"`
		UpDealerOrgInstallerPhone     api.String  `json:"up_dealer_org_installer_phone"`
		UpDealerUserName              api.String  `json:"up_dealer_user_name"`
		UpDealerUserTelNationCode     api.String  `json:"up_dealer_user_tel_nation_code"`
		UpOrgID                       api.Integer `json:"up_org_id"`
		UpOrgName                     api.String  `json:"up_org_name"`
		UserName                      api.String  `json:"user_name"`
		UserTelNationCode             api.String  `json:"user_tel_nation_code"`
	} `json:"user_org_info"`
	UserPassword      api.String  `json:"user_password"`
	UserTelNationCode api.String  `json:"user_tel_nation_code"`
	ValidFlag         api.Integer `json:"valid_flag"`
	Voice             api.Integer `json:"voice"`
	WorkTel           api.String  `json:"work_tel"`
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
