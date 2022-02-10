package web

import (
	"GoSungro/Only"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	DateTimeFormat       = "2006-01-02T15:04:05"
	DefaultAuthTokenFile = "SunGroAuthToken.json"
)

var (
	TokenRequestUrl = "/v1/userService/login"
	GrantPassword   = "password"
	GrantRefresh    = "refresh_token"
)

type SunGroAuth struct {
	TokenExpiry string

	AppKey      string
	Username    string
	Password    string
}

type Token struct {
	Url *url.URL

	Request  TokenRequest
	Response TokenResponse

	TokenFile   string
	TokenExpiry time.Time
	newToken    bool
	retry       int

	Error error
}

type TokenRequest struct {
	Username string `json:"user_account"`
	Password string `json:"user_password"`
	AppKey   string `json:"appkey"`
	SysCode  string `json:"sys_code"`
}

type TokenResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
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
	ResultMsg string `json:"result_msg"`
}

//func NewToken(id string, secret string, username string, password string) token {
//	var t token
//	t.Error = t.SetAuth(id, secret, username, password)
//	return t
//}

func (t *Token) Login(auth *SunGroAuth) error {
	for range Only.Once {
		t.Error = t.readTokenFile()
		//if t.Error == nil {
		//	break
		//}

		t.Error = t.Verify(auth)
		if t.Error != nil {
			break
		}

		t.Error = t.RetrieveToken()
		if t.Error != nil {
			break
		}

		//t.newToken = true
	}

	return t.Error
}

//func (t *Token) Refresh() error {
//	for range Once.Once {
//		//t.Error = t.readTokenFile()
//		//if t.Error == nil {
//		//	break
//		//}
//
//		t.Error = t.RetrieveToken()
//		if t.Error != nil {
//			break
//		}
//
//		t.Error = t.saveToken()
//		if t.Error != nil {
//			break
//		}
//
//		t.newToken = true
//	}
//
//	return t.Error
//}

func (t *Token) Verify(auth *SunGroAuth) error {
	for range Only.Once {
		if auth == nil {
			// If nil, then assume we haven't set anything.
			break
		}

		if auth.AppKey == "" {
			t.Error = errors.New("API AppKey")
			break
		}
		if auth.Username == "" {
			t.Error = errors.New("empty Client ApiUsername")
			break
		}
		if auth.Password == "" {
			t.Error = errors.New("empty Client ApiPassword")
			break
		}

		//if auth.Token == "" {
		//	t.newToken = true
		//}

		if auth.TokenExpiry == "" {
			auth.TokenExpiry = time.Now().Format(DateTimeFormat)
		}
		t.TokenExpiry, t.Error = time.Parse(DateTimeFormat, auth.TokenExpiry)
		if t.Error != nil {
			t.newToken = true
		}

		t.Request = TokenRequest {
			AppKey:   auth.AppKey,
			SysCode:  "900",
			Username: auth.Username,
			Password: auth.Password,
		}

		t.HasTokenExpired()
	}

	return t.Error
}

func (t *Token) RetrieveToken() error {
	for range Only.Once {
		t.HasTokenExpired()
		if !t.newToken {
			break
		}

		u := fmt.Sprintf("%s%s",
			t.Url.String(),
			TokenRequestUrl,
		)
		//p, _ := json.Marshal(map[string]string {
		//	"user_account": t.Request.Username,
		//	"user_password": t.Request.Password,
		//	"appkey": t.Request.AppKey,
		//	"sys_code": "900",
		//})
		p, _ := json.Marshal(t.Request)

		var response *http.Response
		response, t.Error = http.Post(u, "application/json", bytes.NewBuffer(p))
		if t.Error != nil {
			break
		}
		//goland:noinspection GoUnhandledErrorResult
		defer response.Body.Close()
		if response.StatusCode != 200 {
			t.Error = errors.New(fmt.Sprintf("Status Code is %d", response.StatusCode))
			break
		}

		var body []byte
		body, t.Error = ioutil.ReadAll(response.Body)
		if t.Error != nil {
			break
		}

		t.Error = json.Unmarshal(body, &t.Response)
		if t.Error != nil {
			break
		}

		t.TokenExpiry = time.Now()

		t.Error = t.saveToken()
		if t.Error != nil {
			break
		}

		////_, _ = fmt.Fprintf(os.Stderr, "Now: %s\n", t.TokenExpiry.Format("Mon Jan _2 15:04:05 MST 2006"))
		////_, _ = fmt.Fprintf(os.Stderr, "Seconds: %d\n", t.Response.ExpiresIn)
		//t.TokenExpiry = t.TokenExpiry.Add(time.Duration(t.Response.ExpiresIn) * time.Second)
		//
		////_, _ = fmt.Fprintf(os.Stderr, "Expiry: %s\n", t.TokenExpiry.Format("Mon Jan _2 15:04:05 MST 2006"))
		////_, _ = fmt.Fprintf(os.Stderr, "Response:\n%s\n", response)
		////_, _ = fmt.Fprintf(os.Stderr, "Body:\n%s\n", body)
		////_, _ = fmt.Fprintf(os.Stderr, "t.Response:\n%v\n", t.Response)
		////_, _ = fmt.Fprintf(os.Stderr, "Result:\n%s\n", result)
		////_, _ = fmt.Fprintf(os.Stderr, "%s %s\n", t.Response.TokenType, t.Response.AccessToken)
	}

	return t.Error
}

func (t *Token) HasTokenExpired() bool {
	for range Only.Once {
		if t.TokenExpiry.Before(time.Now()) {
			t.newToken = true
			break
		}

		if t.Response.ResultData.Token == "" {
			t.newToken = true
			break
		}
	}

	return t.newToken
}

func (t *Token) HasTokenChanged() bool {
	ok := t.newToken
	if t.newToken {
		t.newToken = false
	}
	return ok
}

func (t *Token) GetAuthHeader() string {
	var ret string

	for range Only.Once {
		//if t.Response.TokenType == "" {
		//	break
		//}
		//
		//if t.Response.AccessToken == "" {
		//	break
		//}
		//
		//ret = t.Response.TokenType + " " + t.Response.AccessToken
	}

	return ret
}

func (t *Token) GetToken() string {
	//return fmt.Sprintf("%s %s", t.Response.TokenType, t.Response.AccessToken)
	return ""
}

func (t *Token) GetTokenExpiry() time.Time {
	return t.TokenExpiry
}

// Retrieves a token from a local file.
func (t *Token) readTokenFile() error {
	for range Only.Once {
		if t.TokenFile == "" {
			t.TokenFile, t.Error = os.UserHomeDir()
			if t.Error != nil {
				t.TokenFile = ""
				break
			}
			t.TokenFile = filepath.Join(t.TokenFile, ".GoSungro", DefaultAuthTokenFile)
		}

		var f *os.File
		f, t.Error = os.Open(t.TokenFile)
		if t.Error != nil {
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		//ret = &oauth2.token{}
		t.Error = json.NewDecoder(f).Decode(&t.Response)
	}

	return t.Error
}

// Saves a token to a file path.
func (t *Token) saveToken() error {
	for range Only.Once {
		if t.TokenFile == "" {
			t.TokenFile, t.Error = os.UserHomeDir()
			if t.Error != nil {
				t.TokenFile = ""
				break
			}
			t.TokenFile = filepath.Join(t.TokenFile, ".GoSungro", DefaultAuthTokenFile)
		}

		//fmt.Printf("Saving token file to: %t\n", t.TokenFile)
		var f *os.File
		f, t.Error = os.OpenFile(t.TokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to cache SUNGRO oauth token: %v", t.Error))
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		t.Error = json.NewEncoder(f).Encode(t.Response)
	}

	return t.Error
}
