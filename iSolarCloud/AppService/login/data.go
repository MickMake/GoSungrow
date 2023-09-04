package login

import (
	"errors"
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
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

	AcceptOrderNum         valueTypes.Integer  `json:"accept_order_num"`
	BackgroundColor        valueTypes.Integer  `json:"background_color"`
	CountryId              valueTypes.String   `json:"countryid"`
	CreateDate             valueTypes.String   `json:"createdate"`
	CreateUserId           valueTypes.String   `json:"createuserid"`
	CurrentOrderNum        valueTypes.Integer  `json:"current_order_num"`
	DisableTime            valueTypes.String   `json:"disable_time"`
	Email                  valueTypes.String   `json:"email"`
	EnglishName            interface{}         `json:"englishname"`
	GcjLatitude            valueTypes.Float    `json:"gcj_latitude"`
	GcjLongitude           valueTypes.Float    `json:"gcj_longitude"`
	ImToken                interface{}         `json:"im_token"`
	IsDST                  valueTypes.Bool     `json:"isDST"`
	IsAfsFlag              valueTypes.Bool     `json:"is_afs_flag"`
	IsAgreeGdpr            valueTypes.Bool     `json:"is_agree_gdpr"`
	IsAu                   valueTypes.Bool     `json:"is_au"`
	IsCanModifyUserAccount valueTypes.Bool     `json:"is_can_modify_user_account"`
	IsDisableMap           valueTypes.Bool     `json:"is_disable_map"`
	IsGdpr                 valueTypes.Bool     `json:"is_gdpr"`
	IsHaveIm               valueTypes.Bool     `json:"is_have_im"`
	IsNewVersion           valueTypes.Bool     `json:"is_new_version"`
	IsOnline               valueTypes.Bool     `json:"is_online"`
	IsOpenProtocol         valueTypes.Bool     `json:"is_open_protocol"`
	IsReceiveNotice        valueTypes.Bool     `json:"is_receive_notice"`
	IsSharePosition        valueTypes.Bool     `json:"is_share_position"`
	IsUploadLocation       valueTypes.Bool     `json:"is_upload_location"`
	IsUseSungrowBrand      valueTypes.Bool     `json:"is_use_sungrow_brand"`
	IsValidMobileEmail     valueTypes.Bool     `json:"is_valid_mobile_email"`
	Isdst                  valueTypes.Bool     `json:"isdst"`
	Jobs                   interface{}         `json:"jobs"`
	Language               valueTypes.String   `json:"language"`
	LoginFirstDate         valueTypes.DateTime `json:"loginFirstDate" PointNameDateFormat:"DateTimeLayout"`
	LoginFirstDate2        valueTypes.DateTime `json:"loginFristDate" PointNameDateFormat:"DateTimeLayout"`
	LoginLastDate          valueTypes.DateTime `json:"loginLastDate" PointNameDateFormat:"DateTimeLayout"`
	LoginLastIP            valueTypes.String   `json:"loginLastIp"`
	LoginTimes             valueTypes.Integer  `json:"loginTimes"`
	Logo                   interface{}         `json:"logo"`
	LogoHTTPSURL           interface{}         `json:"logo_https_url"`
	MapType                valueTypes.String   `json:"map_type"`
	MinDate                valueTypes.DateTime `json:"min_date" PointNameDateFormat:"DateTimeLayout"`
	MobileTel              interface{}         `json:"mobile_tel"`
	OrgId                  valueTypes.Integer  `json:"org_id"`
	OrgName                valueTypes.String   `json:"org_name"`
	OrgTimezone            valueTypes.String   `json:"org_timezone"`
	PasswordIsSimple       valueTypes.Integer  `json:"password_is_simple"`
	PhotoId                interface{}         `json:"photo_id"`
	PhotoURL               interface{}         `json:"photo_url"`
	Privileges             []struct {
		FatherId        valueTypes.Integer `json:"father_id"`
		IconURL         interface{}        `json:"icon_url"`
		IsOpen          valueTypes.Bool    `json:"is_open"`
		IsThirdPlatform valueTypes.Bool    `json:"is_third_platform"`
		MenuCode        valueTypes.String  `json:"menu_code"`
		MenuLevel       valueTypes.Integer `json:"menu_level"`
		MenuName        valueTypes.String  `json:"menu_name"`
		MenuOrder       interface{}        `json:"menu_order"`
		MenuType        valueTypes.String  `json:"menu_type"`
		MenuURL         valueTypes.String  `json:"menu_url"`
		PrivilegeId     valueTypes.Integer `json:"privilege_id"`
		RoleId          valueTypes.Integer `json:"role_id"`
		URLTarget       valueTypes.String  `json:"url_target"`
		VueIcon         interface{}        `json:"vue_icon"`
		VuePath         interface{}        `json:"vue_path"`
	} `json:"privileges" DataTable:"true"`
	RoleId                       valueTypes.String  `json:"role_id"`
	SecondaryOrgIds              []interface{}      `json:"secondaryOrgIds"`
	ServerTel                    valueTypes.String  `json:"server_tel"`
	ServiceVersion               valueTypes.String  `json:"service_version"`
	Sex                          valueTypes.String  `json:"sex"`
	Stylename                    valueTypes.String  `json:"stylename"`
	TimeZone                     valueTypes.String  `json:"timeZone"`
	Timezone                     valueTypes.String  `json:"timezone"`
	Timezoneid                   valueTypes.String  `json:"timezoneid"`
	Toggleflag                   valueTypes.String  `json:"toggleflag"`
	Token                        valueTypes.String  `json:"token"`
	UnlockLaveMinute             valueTypes.Integer `json:"unlock_lave_minute"`
	UploadTime                   interface{}        `json:"upload_time"`
	UserAccount                  valueTypes.String  `json:"user_account"`
	UserAccountModifyCount       valueTypes.Integer `json:"user_account_modify_count"`
	UserAccountModifyRemainTimes valueTypes.Integer `json:"user_account_modify_remain_times"`
	UserDealerOrgCode            interface{}        `json:"user_dealer_org_code"`
	UserId                       valueTypes.String  `json:"user_id"`
	UserLevel                    valueTypes.String  `json:"user_level"`
	UserMasterOrgId              valueTypes.Integer `json:"user_master_org_id"`
	UserMasterOrgName            valueTypes.String  `json:"user_master_org_name"`
	UserMasterOrgTimeZoneId      valueTypes.String  `json:"user_master_org_time_zone_id"`
	UserMasterOrgTimeZoneName    valueTypes.String  `json:"user_master_org_time_zone_name"`
	UserName                     valueTypes.String  `json:"user_name"`
	UserRoleIdList               []string           `json:"user_role_id_list" DataTable:"true"`
	UserTelNationCode            interface{}        `json:"user_tel_nation_code"`
	UserauthorURL                []interface{}      `json:"userauthorURL"`
	Userauthorbutto              []string           `json:"userauthorbutto" DataTable:"true"`
	Userdesc                     interface{}        `json:"userdesc"`
	Userpassword                 valueTypes.String  `json:"userpassword"`
	Validflag                    valueTypes.String  `json:"validflag"`
	Voice                        valueTypes.String  `json:"voice"`
	Welcometext                  valueTypes.String  `json:"welcometext"`
	Wgs84Latitude                valueTypes.Float   `json:"wgs84_latitude"`
	Wgs84Longitude               valueTypes.Float   `json:"wgs84_longitude"`
	WorkTel                      interface{}        `json:"work_tel"`
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
	return e.Response.ResultData.Email.String()
}
func (e *EndPoint) CreateDate() string {
	return e.Response.ResultData.CreateDate.String()
}
func (e *EndPoint) IsOnline() bool {
	return e.Response.ResultData.IsOnline.Value()
}
func (e *EndPoint) LoginLastDate() string {
	return e.Response.ResultData.LoginLastDate.String()
}
func (e *EndPoint) LoginLastIP() string {
	return e.Response.ResultData.LoginLastIP.String()
}
func (e *EndPoint) LoginState() string {
	return e.Response.ResultData.LoginState.String()
}
func (e *EndPoint) Token() string {
	return e.Response.ResultData.Token.String()
}
func (e *EndPoint) UserAccount() string {
	return e.Response.ResultData.UserAccount.String()
}
func (e *EndPoint) UserId() string {
	return e.Response.ResultData.UserId.String()
}
func (e *EndPoint) UserName() string {
	return e.Response.ResultData.UserName.String()
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
