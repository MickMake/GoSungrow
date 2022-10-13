package login

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"

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

	AcceptOrderNum         api.Integer `json:"accept_order_num"`
	BackgroundColor        api.Integer `json:"background_color"`
	Countryid              string      `json:"countryid"`
	Createdate             string      `json:"createdate"`
	Createuserid           string      `json:"createuserid"`
	CurrentOrderNum        api.Integer `json:"current_order_num"`
	DisableTime            string      `json:"disable_time"`
	Email                  string      `json:"email"`
	Englishname            interface{} `json:"englishname"`
	GcjLatitude            api.Float   `json:"gcj_latitude"`
	GcjLongitude           api.Float   `json:"gcj_longitude"`
	ImToken                interface{} `json:"im_token"`
	IsDST                  api.Bool    `json:"isDST"`
	IsAfsFlag              api.Bool    `json:"is_afs_flag"`
	IsAgreeGdpr            api.Bool    `json:"is_agree_gdpr"`
	IsAu                   api.Bool    `json:"is_au"`
	IsCanModifyUserAccount api.Bool    `json:"is_can_modify_user_account"`
	IsDisableMap           api.Bool    `json:"is_disable_map"`
	IsGdpr                 api.Bool    `json:"is_gdpr"`
	IsHaveIm               api.Bool    `json:"is_have_im"`
	IsNewVersion           api.Bool    `json:"is_new_version"`
	IsOnline               api.Bool    `json:"is_online"`
	IsOpenProtocol         api.Bool    `json:"is_open_protocol"`
	IsReceiveNotice        api.Bool    `json:"is_receive_notice"`
	IsSharePosition        api.Bool    `json:"is_share_position"`
	IsUploadLocation       api.Bool    `json:"is_upload_location"`
	IsUseSungrowBrand      api.Bool    `json:"is_use_sungrow_brand"`
	IsValidMobileEmail     api.Bool    `json:"is_valid_mobile_email"`
	Isdst                  api.Bool    `json:"isdst"`
	Jobs                   interface{} `json:"jobs"`
	Language               string      `json:"language"`
	LoginFirstDate         api.DateTime      `json:"loginFirstDate"`
	LoginFirstDate2        api.DateTime      `json:"loginFristDate"`
	LoginLastDate          api.DateTime      `json:"loginLastDate"`
	LoginLastIP            string      `json:"loginLastIp"`
	LoginTimes             api.Integer `json:"loginTimes"`
	Logo                   interface{} `json:"logo"`
	LogoHTTPSURL           interface{} `json:"logo_https_url"`
	MapType                string      `json:"map_type"`
	MinDate                api.DateTime      `json:"min_date"`
	MobileTel              interface{} `json:"mobile_tel"`
	OrgID                  string      `json:"org_id"`
	OrgName                string      `json:"org_name"`
	OrgTimezone            string      `json:"org_timezone"`
	PasswordIsSimple       api.Integer `json:"password_is_simple"`
	PhotoID                interface{} `json:"photo_id"`
	PhotoURL               interface{} `json:"photo_url"`
	Privileges             []struct {
		FatherID        api.Integer `json:"father_id"`
		IconURL         interface{} `json:"icon_url"`
		IsOpen          api.Bool    `json:"is_open"`
		IsThirdPlatform api.Bool    `json:"is_third_platform"`
		MenuCode        string      `json:"menu_code"`
		MenuLevel       api.Integer `json:"menu_level"`
		MenuName        string      `json:"menu_name"`
		MenuOrder       interface{} `json:"menu_order"`
		MenuType        string      `json:"menu_type"`
		MenuURL         string      `json:"menu_url"`
		PrivilegeID     api.Integer `json:"privilege_id"`
		RoleID          api.Integer `json:"role_id"`
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
	UnlockLaveMinute             api.Integer   `json:"unlock_lave_minute"`
	UploadTime                   interface{}   `json:"upload_time"`
	UserAccount                  string        `json:"user_account"`
	UserAccountModifyCount       api.Integer   `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes api.Integer   `json:"user_account_modify_remain_times"`
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
	Wgs84Latitude                api.Float     `json:"wgs84_latitude"`
	Wgs84Longitude               api.Float     `json:"wgs84_longitude"`
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
func (e *EndPoint) IsOnline() bool {
	return e.Response.ResultData.IsOnline.Value()
}
func (e *EndPoint) LoginLastDate() string {
	return e.Response.ResultData.LoginLastDate.String()
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


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", api.NewDateTime(""))
	}

	return entries
}
