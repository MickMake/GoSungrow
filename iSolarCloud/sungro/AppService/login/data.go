package login

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"errors"
	"fmt"
)

const Url = "/v1/userService/login"
const Disabled = false

type RequestData struct {
	UserAccount  string `json:"user_account" required:"true"`
	UserPassword string `json:"user_password" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ErrTimes    string `json:"err_times"`
	LoginState  string `json:"login_state"`
	Msg         string `json:"msg"`
	RemainTimes string `json:"remain_times"`

	AcceptOrderNum         int64       `json:"accept_order_num"`
	BackgroundColor        int64       `json:"background_color"`
	Countryid              string      `json:"countryid"`
	Createdate             string      `json:"createdate"`
	Createuserid           string      `json:"createuserid"`
	CurrentOrderNum        int64       `json:"current_order_num"`
	DisableTime            string      `json:"disable_time"`
	Email                  string      `json:"email"`
	Englishname            interface{} `json:"englishname"`
	GcjLatitude            interface{} `json:"gcj_latitude"`
	GcjLongitude           interface{} `json:"gcj_longitude"`
	ImToken                interface{} `json:"im_token"`
	IsDST                  string      `json:"isDST"`
	IsAfsFlag              string      `json:"is_afs_flag"`
	IsAgreeGdpr            int64       `json:"is_agree_gdpr"`
	IsAu                   int64       `json:"is_au"`
	IsCanModifyUserAccount int64       `json:"is_can_modify_user_account"`
	IsDisableMap           string      `json:"is_disable_map"`
	IsGdpr                 int64       `json:"is_gdpr"`
	IsHaveIm               int64       `json:"is_have_im"`
	IsNewVersion           int64       `json:"is_new_version"`
	IsOnline               string      `json:"is_online"`
	IsOpenProtocol         int64       `json:"is_open_protocol"`
	IsReceiveNotice        int64       `json:"is_receive_notice"`
	IsSharePosition        int64       `json:"is_share_position"`
	IsUploadLocation       int64       `json:"is_upload_location"`
	IsUseSungrowBrand      string      `json:"is_use_sungrow_brand"`
	IsValidMobileEmail     int64       `json:"is_valid_mobile_email"`
	Isdst                  string      `json:"isdst"`
	Jobs                   interface{} `json:"jobs"`
	Language               string      `json:"language"`
	LoginFirstDate         string      `json:"loginFirstDate"`
	LoginFristDate         string      `json:"loginFristDate"`
	LoginLastDate          string      `json:"loginLastDate"`
	LoginLastIP            string      `json:"loginLastIp"`
	LoginTimes             int64       `json:"loginTimes"`
	Logo                   interface{} `json:"logo"`
	LogoHTTPSURL           interface{} `json:"logo_https_url"`
	MapType                string      `json:"map_type"`
	MinDate                string      `json:"min_date"`
	MobileTel              interface{} `json:"mobile_tel"`
	OrgID                  string      `json:"org_id"`
	OrgName                string      `json:"org_name"`
	OrgTimezone            string      `json:"org_timezone"`
	PasswordIsSimple       int64       `json:"password_is_simple"`
	PhotoID                interface{} `json:"photo_id"`
	PhotoURL               interface{} `json:"photo_url"`
	Privileges             []struct {
		FatherID        int64       `json:"father_id"`
		IconURL         interface{} `json:"icon_url"`
		IsOpen          interface{} `json:"is_open"`
		IsThirdPlatform int64       `json:"is_third_platform"`
		MenuCode        string      `json:"menu_code"`
		MenuLevel       int64       `json:"menu_level"`
		MenuName        string      `json:"menu_name"`
		MenuOrder       interface{} `json:"menu_order"`
		MenuType        string      `json:"menu_type"`
		MenuURL         string      `json:"menu_url"`
		PrivilegeID     int64       `json:"privilege_id"`
		RoleID          int64       `json:"role_id"`
		URLTarget       string      `json:"url_target"`
		VueIcon         interface{} `json:"vue_icon"`
		VuePath         interface{} `json:"vue_path"`
	} `json:"privileges"`
	RoleID                       string        `json:"role_id"`
	SecondaryOrgIds              []interface{} `json:"secondaryOrgIds"`
	ServerTel                    string        `json:"server_tel"`
	ServiceVersion               string        `json:"service_version"`
	Sex                          string        `json:"sex"`
	Stylename                    string        `json:"stylename"`
	TimeZone                     string        `json:"timeZone"`
	Timezone                     string        `json:"timezone"`
	Timezoneid                   string        `json:"timezoneid"`
	Toggleflag                   string        `json:"toggleflag"`
	Token                        string        `json:"token"`
	UnlockLaveMinute             int64         `json:"unlock_lave_minute"`
	UploadTime                   interface{}   `json:"upload_time"`
	UserAccount                  string        `json:"user_account"`
	UserAccountModifyCount       int64         `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes int64         `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            interface{}   `json:"user_dealer_org_code"`
	UserID                       string        `json:"user_id"`
	UserLevel                    string        `json:"user_level"`
	UserMasterOrgID              string        `json:"user_master_org_id"`
	UserMasterOrgName            string        `json:"user_master_org_name"`
	UserMasterOrgTimeZoneID      string        `json:"user_master_org_time_zone_id"`
	UserMasterOrgTimeZoneName    string        `json:"user_master_org_time_zone_name"`
	UserName                     string        `json:"user_name"`
	UserRoleIDList               []string      `json:"user_role_id_list"`
	UserTelNationCode            interface{}   `json:"user_tel_nation_code"`
	UserauthorURL                []interface{} `json:"userauthorURL"`
	Userauthorbutto              []string      `json:"userauthorbutto"`
	Userdesc                     interface{}   `json:"userdesc"`
	Userpassword                 string        `json:"userpassword"`
	Validflag                    string        `json:"validflag"`
	Voice                        string        `json:"voice"`
	Welcometext                  string        `json:"welcometext"`
	Wgs84Latitude                interface{}   `json:"wgs84_latitude"`
	Wgs84Longitude               interface{}   `json:"wgs84_longitude"`
	WorkTel                      interface{}   `json:"work_tel"`
}

func (e *ResultData) IsValid() error {
	var err error
	for range Only.Once {
		switch {
		case e.Msg == `账号不存在`:
			err = errors.New(fmt.Sprintf("Account does not exist '%s'", e.Msg))
		case e.Msg == fmt.Sprintf(`账号或密码输入错误，还剩%s次机会`, e.RemainTimes):
			err = errors.New(fmt.Sprintf("The account number or password is entered incorrectly, there are %s chances left '%s'", e.RemainTimes, e.Msg))
		case e.Msg == "":
			break
		default:
			err = errors.New(fmt.Sprintf("unknown error '%s'", e.Msg))
		}
	}
	return err
}

func (e *EndPoint) AppKey() string {
	return e.Request.RequestCommon.Appkey
}

func (e *EndPoint) UserId() string {
	return e.Response.ResultData.UserID
}

func (e *EndPoint) Email() string {
	return e.Response.ResultData.Email
}
func (e *EndPoint) CreateDate() string {
	return e.Response.ResultData.Createdate
}
func (e *EndPoint) IsOnline() string {
	return e.Response.ResultData.IsOnline
}
func (e *EndPoint) LoginLastDate() string {
	return e.Response.ResultData.LoginLastDate
}
func (e *EndPoint) LoginLastIP() string {
	return e.Response.ResultData.LoginLastIP
}
func (e *EndPoint) LoginState() string {
	return e.Response.ResultData.LoginState
}
func (e *EndPoint) Token() string {
	return e.Response.ResultData.Token
}
func (e *EndPoint) UserAccount() string {
	return e.Response.ResultData.UserAccount
}
func (e *EndPoint) UserID() string {
	return e.Response.ResultData.UserID
}
func (e *EndPoint) UserName() string {
	return e.Response.ResultData.UserName
}
