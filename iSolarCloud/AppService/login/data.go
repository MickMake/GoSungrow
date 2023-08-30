package login

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"errors"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/userService/login"
const Disabled = false
const EndPointName = "AppService.login"

type RequestData struct {
	UserAccount  valueTypes.String `json:"user_account" required:"true"`
	UserPassword valueTypes.String `json:"user_password" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ErrTimes    valueTypes.String `json:"err_times"`
	LoginState  valueTypes.String `json:"login_state"`
	Msg         valueTypes.String `json:"msg"`
	RemainTimes valueTypes.String `json:"remain_times"`

	AcceptOrderNum         valueTypes.Integer `json:"accept_order_num"`
	BackgroundColor        valueTypes.Integer `json:"background_color"`
	CountryId              string      `json:"countryid"`
	CreateDate             string      `json:"createdate"`
	CreateUserId           string      `json:"createuserid"`
	CurrentOrderNum        valueTypes.Integer `json:"current_order_num"`
	DisableTime            string      `json:"disable_time"`
	Email                  string      `json:"email"`
	EnglishName            interface{} `json:"englishname"`
	GcjLatitude            valueTypes.Float   `json:"gcj_latitude"`
	GcjLongitude           valueTypes.Float   `json:"gcj_longitude"`
	ImToken                interface{} `json:"im_token"`
	IsDST                  valueTypes.Bool    `json:"isDST"`
	IsAfsFlag              valueTypes.Bool    `json:"is_afs_flag"`
	IsAgreeGdpr            valueTypes.Bool    `json:"is_agree_gdpr"`
	IsAu                   valueTypes.Bool    `json:"is_au"`
	IsCanModifyUserAccount valueTypes.Bool    `json:"is_can_modify_user_account"`
	IsDisableMap           valueTypes.Bool    `json:"is_disable_map"`
	IsGdpr                 valueTypes.Bool    `json:"is_gdpr"`
	IsHaveIm               valueTypes.Bool    `json:"is_have_im"`
	IsNewVersion           valueTypes.Bool    `json:"is_new_version"`
	IsOnline               valueTypes.Bool    `json:"is_online"`
	IsOpenProtocol         valueTypes.Bool    `json:"is_open_protocol"`
	IsReceiveNotice        valueTypes.Bool    `json:"is_receive_notice"`
	IsSharePosition        valueTypes.Bool    `json:"is_share_position"`
	IsUploadLocation       valueTypes.Bool    `json:"is_upload_location"`
	IsUseSungrowBrand      valueTypes.Bool    `json:"is_use_sungrow_brand"`
	IsValidMobileEmail     valueTypes.Bool    `json:"is_valid_mobile_email"`
	Isdst                  valueTypes.Bool    `json:"isdst"`
	Jobs                   interface{} `json:"jobs"`
	Language               string      `json:"language"`
	LoginFirstDate         valueTypes.DateTime      `json:"loginFirstDate" PointNameDateFormat:"DateTimeLayout"`
	LoginFirstDate2        valueTypes.DateTime      `json:"loginFristDate" PointNameDateFormat:"DateTimeLayout"`
	LoginLastDate          valueTypes.DateTime      `json:"loginLastDate" PointNameDateFormat:"DateTimeLayout"`
	LoginLastIP            string      `json:"loginLastIp"`
	LoginTimes             valueTypes.Integer `json:"loginTimes"`
	Logo                   interface{} `json:"logo"`
	LogoHTTPSURL           interface{} `json:"logo_https_url"`
	MapType                string      `json:"map_type"`
	MinDate                valueTypes.DateTime      `json:"min_date" PointNameDateFormat:"DateTimeLayout"`
	MobileTel              interface{} `json:"mobile_tel"`
	OrgId                  valueTypes.Integer      `json:"org_id"`
	OrgName                string      `json:"org_name"`
	OrgTimezone            string      `json:"org_timezone"`
	PasswordIsSimple       valueTypes.Integer `json:"password_is_simple"`
	PhotoId                interface{} `json:"photo_id"`
	PhotoURL               interface{} `json:"photo_url"`
	Privileges             []struct {
		FatherId        valueTypes.Integer `json:"father_id"`
		IconURL         interface{} `json:"icon_url"`
		IsOpen          valueTypes.Bool    `json:"is_open"`
		IsThirdPlatform valueTypes.Bool    `json:"is_third_platform"`
		MenuCode        string      `json:"menu_code"`
		MenuLevel       valueTypes.Integer `json:"menu_level"`
		MenuName        string      `json:"menu_name"`
		MenuOrder       interface{} `json:"menu_order"`
		MenuType        string      `json:"menu_type"`
		MenuURL         string      `json:"menu_url"`
		PrivilegeId     valueTypes.Integer `json:"privilege_id"`
		RoleId          valueTypes.Integer `json:"role_id"`
		URLTarget       string      `json:"url_target"`
		VueIcon         interface{} `json:"vue_icon"`
		VuePath         interface{} `json:"vue_path"`
	} `json:"privileges" DataTable:"true"`
	RoleId                       string        `json:"role_id"`
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
	UnlockLaveMinute             valueTypes.Integer   `json:"unlock_lave_minute"`
	UploadTime                   interface{}   `json:"upload_time"`
	UserAccount                  string        `json:"user_account"`
	UserAccountModifyCount       valueTypes.Integer   `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes valueTypes.Integer   `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            interface{}   `json:"user_dealer_org_code"`
	UserId                       string        `json:"user_id"`
	UserLevel                    string        `json:"user_level"`
	UserMasterOrgId              string        `json:"user_master_org_id"`
	UserMasterOrgName            string        `json:"user_master_org_name"`
	UserMasterOrgTimeZoneId      string        `json:"user_master_org_time_zone_id"`
	UserMasterOrgTimeZoneName    string        `json:"user_master_org_time_zone_name"`
	UserName                     string        `json:"user_name"`
	UserRoleIdList               []string      `json:"user_role_id_list" DataTable:"true"`
	UserTelNationCode            interface{}   `json:"user_tel_nation_code"`
	UserauthorURL                []interface{} `json:"userauthorURL"`
	Userauthorbutto              []string      `json:"userauthorbutto" DataTable:"true"`
	Userdesc                     interface{}   `json:"userdesc"`
	Userpassword                 string        `json:"userpassword"`
	Validflag                    string        `json:"validflag"`
	Voice                        string        `json:"voice"`
	Welcometext                  string        `json:"welcometext"`
	Wgs84Latitude                valueTypes.Float     `json:"wgs84_latitude"`
	Wgs84Longitude               valueTypes.Float     `json:"wgs84_longitude"`
	WorkTel                      interface{}   `json:"work_tel"`
}

func (e *ResultData) IsValid() error {
	var err error
	for range Only.Once {
		switch {
		case e.Msg.String() == `账号不存在`:
			err = errors.New(fmt.Sprintf("Account does not exist '%s'", e.Msg))
		case e.Msg.String() == fmt.Sprintf(`账号或密码输入错误，还剩%s次机会`, e.RemainTimes):
			err = errors.New(fmt.Sprintf("The account number or password is entered incorrectly, there are %s chances left '%s'", e.RemainTimes, e.Msg))
		case e.Msg.String() == "":
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

func (e *EndPoint) Email() string {
	return e.Response.ResultData.Email
}
func (e *EndPoint) CreateDate() string {
	return e.Response.ResultData.CreateDate
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
	return e.Response.ResultData.LoginState.String()
}
func (e *EndPoint) Token() string {
	return e.Response.ResultData.Token
}
func (e *EndPoint) UserAccount() string {
	return e.Response.ResultData.UserAccount
}
func (e *EndPoint) UserId() string {
	return e.Response.ResultData.UserId
}
func (e *EndPoint) UserName() string {
	return e.Response.ResultData.UserName
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
