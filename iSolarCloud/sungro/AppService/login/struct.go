// EndPoint
package login

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)


var Url = "/v1/userService/login"

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint struct {
	api.EndPointStruct
	Token
	Request Request
	Response Response
}

type Request struct {
	// api.RequestCommon	// login is special as it doesn't have the usual fields.
	Appkey       string `json:"appkey" required:"true"`
	SysCode      string `json:"sys_code" required:"true"`
	UserAccount  string `json:"user_account" required:"true"`
	UserPassword string `json:"user_password" required:"true"`
}

type Response struct {
	api.ResponseCommon
	ResultData ResultData `json:"result_data"`
}

type ResultData struct {
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
}


func Init() EndPoint {
	fmt.Println("Init()")
	return EndPoint {
		EndPointStruct: api.EndPointStruct {
			Area: api.GetArea(EndPoint{}),
			Name:     api.GetName(EndPoint{}),
			Url:      api.GetUrl(Url),
			Request:  Request{},
			Response: Response{},
			Error:    nil,
		},
		Token: Token {
			TokenFile:   DefaultAuthTokenFile,
			TokenExpiry: time.Time{},
			newToken:    true,
			retry:       0,
		},
	}
}


// ****************************************
// Methods not scoped by api.EndPoint interface type

func (e EndPoint) Init() *EndPoint {
	ret := Init()
	return &ret
}

func (e EndPoint) GetRequest() Request {
	// return e.Request.(Request)
	return e.Request
}

func (e EndPoint) GetResponse() Response {
	// return e.Response.(Response)
	return e.Response
}

func (e EndPoint) CallMe() api.EndPoint {
	for range Only.Once {
		// e.Error = api.DoTypesMatch(e.Response, ref)
		// if e.Error != nil {
		// 	break
		// }
		// e.Response = ref.(Response)
	}
	return e
}


// ****************************************
// Methods defined by api.EndPoint interface type

func (e EndPoint) GetArea() api.AreaName {
	return e.Area
}

func (e EndPoint) GetName() api.EndPointName {
	return e.Name
}

func (e EndPoint) GetUrl() *url.URL {
	return e.Url
}

func (e EndPoint) Call() api.Json {
	fmt.Println("e.Call() implement me")
	return ""
}

func (e EndPoint) GetData() api.Json {
	// return api.GetAsJson(e.Response.(Response).ResultData)
	return api.GetAsJson(e.Response.ResultData)
}

func (e EndPoint) GetError() error {
	return e.Error
}

func (e EndPoint) IsError() bool {
	if e.Error != nil {
		return true
	}
	return false
}

func (e EndPoint) SetRequest(ref interface{}) api.EndPoint {
	for range Only.Once {
		e.Error = api.DoTypesMatch(e.Request, ref)
		if e.Error != nil {
			break
		}
		e.Request = ref.(Request)
	}
	return e
}

func (e EndPoint) RequestRef() interface{} {
	return e.Request
}

func (e EndPoint) GetRequestJson() api.Json {
	return api.GetAsJson(e.Request)
}

func (e EndPoint) IsRequestValid() error {
	for range Only.Once {
		req := e.GetRequest()
		e.Error = api.CheckString("SysCode", req.SysCode)
		if e.Error != nil {
			break
		}
		e.Error = api.CheckString("Appkey", req.Appkey)
		if e.Error != nil {
			break
		}
		e.Error = api.CheckString("UserAccount", req.UserAccount)
		if e.Error != nil {
			break
		}
		e.Error = api.CheckString("UserPassword", req.UserPassword)
		if e.Error != nil {
			break
		}
	}
	return e.Error
}

func (e EndPoint) SetResponse(ref []byte) api.EndPoint {
	for range Only.Once {
		r := e.GetResponse()
		e.Error = json.Unmarshal(ref, &r)
		if e.Error != nil {
			break
		}
		e.Response = r
	}
	return e
}

func (e EndPoint) GetResponseJson() api.Json {
	return api.GetAsJson(e.Response)
}

func (e EndPoint) ResponseRef() interface{} {
	return e.Response
}

func (e EndPoint) IsResponseValid() error {
	for range Only.Once {
		resp := e.GetResponse()
		e.Error = resp.ResponseCommon.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}
