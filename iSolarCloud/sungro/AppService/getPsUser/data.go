package getPsUser

import (
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
	BackgroundColor              int64       `json:"background_color"`
	BloodType                    string      `json:"blood_type"`
	CountryID                    int64       `json:"country_id"`
	CountryName                  string      `json:"country_name"`
	CreateDate                   string      `json:"create_date"`
	CreateUserID                 int64       `json:"create_user_id"`
	DiseaseHis                   string      `json:"disease_his"`
	Email                        string      `json:"email"`
	EnglishName                  string      `json:"english_name"`
	IsAgreeGdpr                  int64       `json:"is_agree_gdpr"`
	IsCanModifyUserAccount       int64       `json:"is_can_modify_user_account"`
	IsDst                        int64       `json:"is_dst"`
	IsGdpr                       int64       `json:"is_gdpr"`
	IsNewVersion                 int64       `json:"is_new_version"`
	IsOpenProtocol               int64       `json:"is_open_protocol"`
	IsReceiveNotice              int64       `json:"is_receive_notice"`
	IsSharePosition              int64       `json:"is_share_position"`
	IsValidMobileEmail           int64       `json:"is_valid_mobile_email"`
	LackOfInformation            int64       `json:"lack_of_information"`
	Language                     string      `json:"language"`
	MobleTel                     string      `json:"moble_tel"`
	OrgID                        int64       `json:"org_id"`
	PhotoID                      int64       `json:"photo_id"`
	PhotoURL                     string      `json:"photo_url"`
	RootOrgID                    string      `json:"root_org_id"`
	Sex                          string      `json:"sex"`
	Stylename                    string      `json:"stylename"`
	TimeZoneID                   int64       `json:"time_zone_id"`
	Timezone                     string      `json:"timezone"`
	UserAccount                  string      `json:"user_account"`
	UserAccountModifyCount       int64       `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes int64       `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            interface{} `json:"user_dealer_org_code"`
	UserID                       int64       `json:"user_id"`
	UserLanguage                 string      `json:"user_language"`
	UserLevel                    int64       `json:"user_level"`
	UserName                     string      `json:"user_name"`
	UserOrgInfo                  struct {
		DealerOrgCode                 string      `json:"dealer_org_code"`
		Email                         string      `json:"email"`
		Installer                     string      `json:"installer"`
		InstallerEmail                string      `json:"installer_email"`
		InstallerPhone                string      `json:"installer_phone"`
		IsCountryDefaultDealerOrgCode int64       `json:"is_country_default_dealer_org_code"`
		MobleTel                      interface{} `json:"moble_tel"`
		OrgID                         int64       `json:"org_id"`
		OrgIndexCode                  string      `json:"org_index_code"`
		OrgName                       string      `json:"org_name"`
		UpDealerEmail                 string      `json:"up_dealer_email"`
		UpDealerMobleTel              string      `json:"up_dealer_moble_tel"`
		UpDealerOrgCode               string      `json:"up_dealer_org_code"`
		UpDealerOrgInstaller          string      `json:"up_dealer_org_installer"`
		UpDealerOrgInstallerEmail     string      `json:"up_dealer_org_installer_email"`
		UpDealerOrgInstallerPhone     string      `json:"up_dealer_org_installer_phone"`
		UpDealerUserName              string      `json:"up_dealer_user_name"`
		UpDealerUserTelNationCode     string      `json:"up_dealer_user_tel_nation_code"`
		UpOrgID                       int64       `json:"up_org_id"`
		UpOrgName                     string      `json:"up_org_name"`
		UserName                      string      `json:"user_name"`
		UserTelNationCode             interface{} `json:"user_tel_nation_code"`
	} `json:"user_org_info"`
	UserPassword      string      `json:"user_password"`
	UserTelNationCode interface{} `json:"user_tel_nation_code"`
	ValidFlag         int64       `json:"valid_flag"`
	Voice             int64       `json:"voice"`
	WorkTel           string      `json:"work_tel"`
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
