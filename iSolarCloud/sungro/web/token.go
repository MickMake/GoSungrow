package web

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/sungro/AppService/login"
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
	// GrantPassword   = "password"
	// GrantRefresh    = "refresh_token"
)

type SunGroAuth struct {
	TokenExpiry string

	AppKey      string
	Username    string
	Password    string
}

type Token struct {
	Url *url.URL

	Request  login.Request
	Response login.Response

	TokenFile   string
	TokenExpiry time.Time
	newToken    bool
	retry       int

	Error error
}


func (t *Token) Login(auth *SunGroAuth) error {
	for range Only.Once {
		_ = t.readTokenFile()

		t.Error = t.Verify(auth)
		if t.Error != nil {
			break
		}

		t.Error = t.RetrieveToken()
		if t.Error != nil {
			break
		}
	}

	return t.Error
}

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

		if t.Response.ResultData.Token == "" {
			t.newToken = true
		}

		if auth.TokenExpiry == "" {
			auth.TokenExpiry = time.Now().Format(DateTimeFormat)
		}
		t.TokenExpiry, t.Error = time.Parse(DateTimeFormat, auth.TokenExpiry)
		if t.Error != nil {
			t.newToken = true
		}

		t.Request = login.Request {
			Appkey:   auth.AppKey,
			SysCode:  "900",
			UserAccount: auth.Username,
			UserPassword: auth.Password,
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

// func (t *Token) GetAuthHeader() string {
// 	var ret string
//
// 	for range Only.Once {
// 		//if t.Response.TokenType == "" {
// 		//	break
// 		//}
// 		//
// 		//if t.Response.AccessToken == "" {
// 		//	break
// 		//}
// 		//
// 		//ret = t.Response.TokenType + " " + t.Response.AccessToken
// 	}
//
// 	return ret
// }

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

		fmt.Printf("Saving token file to: %s\n", t.TokenFile)
		var f *os.File
		f, t.Error = os.OpenFile(t.TokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to cache SUNGRO oauth token: %v", t.Error))
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		t.Error = json.NewEncoder(f).Encode(t.Response.ResultData)
	}

	return t.Error
}
