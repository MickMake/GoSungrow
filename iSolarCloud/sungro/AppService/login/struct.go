// EndPoint
package login

import (
	"GoSungro/iSolarCloud/api"
	"fmt"
	"net/url"
)


var Url = "/v1/userService/login"

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint api.EndPointStruct

type Request struct {
	// api.RequestCommon	// login is special as it doesn't have the usual fields.
	Appkey       string `json:"appkey"`
	SysCode      string `json:"sys_code"`
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}

type Response struct {
	api.ResponseCommon
	ResultData   struct {
		AcceptOrderNum         int64       `json:"accept_order_num"`
		BackgroundColor        int64       `json:"background_color"`
		Countryid              string      `json:"countryid"`
		Createdate             string      `json:"createdate"`
		Createuserid           string      `json:"createuserid"`
		CurrentOrderNum        int64       `json:"current_order_num"`
		DisableTime            string      `json:"disable_time"`
		Email                  string      `json:"email"`
		Englishname            interface{} `json:"englishname"`
		ErrTimes               string      `json:"err_times"`
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
		LoginState             string      `json:"login_state"`
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
	} `json:"result_data"`
}


func Init() EndPoint {
	fmt.Println("Init()")

	foo := EndPoint {
		Area:     api.GetArea(EndPoint{}),
		Name:     api.GetName(EndPoint{}),
		Url:      api.GetUrl(Url),
		Request:  Request{},
		Response: Response{},
		Error:    nil,
	}

	fmt.Printf("endpoint: %v\n", foo)

	return foo
}


func (g EndPoint) GetArea() api.AreaName {
	fmt.Println("g.GetArea()")
	return g.Area
}

func (g EndPoint) GetName() api.EndPointName {
	fmt.Println("g.GetName()")
	return g.Name
}

func (g EndPoint) GetUrl() *url.URL {
	fmt.Println("g.GetUrl()")
	return g.Url
}

func (g EndPoint) SetRequest(ref interface{}) error {
	fmt.Println("g.SetRequest()")
	g.Request = ref.(Request)
	fmt.Printf("ref == %v\n", g.Request)
	return nil
}

func (g EndPoint) GetRequest() api.Json {
	return api.GetAsJson(g.Request)
}

func (g EndPoint) GetResponse() api.Json {
	return api.GetAsJson(g.Response)
}

func (g EndPoint) GetData() api.Json {
	return api.GetAsJson(g.Response.(Response).ResultData)
}

func (g EndPoint) IsValid() error {
	fmt.Println("g.IsValid() implement me")
	return nil
}

func (g EndPoint) Call() api.Json {
	fmt.Println("g.Call() implement me")
	return ""
}

func (g EndPoint) Init() *EndPoint {
	fmt.Println("g.Init()")
	ret := Init()
	return &ret
}

func (g EndPoint) GetError() error {
	return g.Error
}
