package login

import (
	"GoSungro/Only"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)


const (
	DateTimeFormat       = "2006-01-02T15:04:05"
	DefaultAuthTokenFile = "SunGroAuthToken.json"
)

type SunGroAuth struct {
	AppKey       string
	UserAccount  string
	UserPassword string
	TokenFile    string
	Token        string

	expiry       time.Time
	newToken     bool
	retry        int
	err          error
}

// type Token struct {
// 	File     string
// 	Expiry   time.Time
// 	newToken bool
// 	retry    int
// }


func (a *SunGroAuth) Verify() error {
	var err error
	for range Only.Once {
		if a == nil {
			err = errors.New("no auth details")
			break
		}

		if a.AppKey == "" {
			err = errors.New("API AppKey")
			a.err = err
			break
		}
		if a.UserAccount == "" {
			err = errors.New("empty API Username")
			a.err = err
			break
		}
		if a.UserPassword == "" {
			err = errors.New("empty API Password")
			a.err = err
			break
		}
	}

	return err
}

func (e *EndPoint) Login(auth *SunGroAuth) error {
	for range Only.Once {
		e.Auth = auth

		e.Error = e.Auth.Verify()
		if e.Error != nil {
			break
		}

		e.Error = e.readTokenFile()
		if e.Error != nil {
			break
		}

		e.Error = e.Auth.Verify()
		if e.Error != nil {
			break
		}

		if e.GetToken() != "" {
			break
			// e.Auth.newToken = true
		}

		// if auth.Expiry == "" {
		// 	auth.Expiry = time.Now().Format(DateTimeFormat)
		// }
		// e.Expiry, e.Error = time.Parse(DateTimeFormat, auth.Expiry)
		// if e.Error != nil {
		// 	e.newToken = true
		// }

		e.Request = Request {
			Appkey:   auth.AppKey,
			SysCode:  "900",
			UserAccount: auth.UserAccount,
			UserPassword: auth.UserPassword,
		}

		foo := Assert(e.Call())
		e.Response = foo.Response

		e.Error = e.saveToken()
		if e.Error != nil {
			break
		}

		e.Error = e.RetrieveToken()
		if e.Error != nil {
			break
		}
	}

	return e.Error
}

func (e *EndPoint) GetToken() string {
	return e.GetResponse().ResultData.Token
}

func (e *EndPoint) GetUserEmail() string {
	return e.GetResponse().ResultData.Email
}

func (e *EndPoint) GetUserName() string {
	return e.GetResponse().ResultData.UserName
}

func (e *EndPoint) GetUserId() string {
	return e.GetResponse().ResultData.UserID
}

func (e *EndPoint) RetrieveToken() error {
	for range Only.Once {
		e.HasTokenExpired()
		if !e.Auth.newToken {
			break
		}

		u := fmt.Sprintf("%s",
			e.GetUrl(),
		)
		p, _ := json.Marshal(e.Request)

		var response *http.Response
		response, e.Error = http.Post(u, "application/json", bytes.NewBuffer(p))
		if e.Error != nil {
			break
		}
		//goland:noinspection GoUnhandledErrorResult
		defer response.Body.Close()
		if response.StatusCode != 200 {
			e.Error = errors.New(fmt.Sprintf("Status Code is %d", response.StatusCode))
			break
		}

		var body []byte
		body, e.Error = ioutil.ReadAll(response.Body)
		if e.Error != nil {
			break
		}

		e.Error = json.Unmarshal(body, &e.Response)
		if e.Error != nil {
			break
		}

		e.Auth.expiry = time.Now()

		e.Error = e.saveToken()
		if e.Error != nil {
			break
		}
	}

	return e.Error
}

func (e *EndPoint) HasTokenExpired() bool {
	for range Only.Once {
		if e.Auth.expiry.Before(time.Now()) {
			e.Auth.newToken = true
			break
		}

		if e.GetToken() == "" {
			e.Auth.newToken = true
			break
		}
	}

	return e.Auth.newToken
}

func (e *EndPoint) HasTokenChanged() bool {
	ok := e.Auth.newToken
	if e.Auth.newToken {
		e.Auth.newToken = false
	}
	return ok
}

func (e *EndPoint) GetTokenExpiry() time.Time {
	return e.Auth.expiry
}

// Retrieves a token from a local file.
func (e *EndPoint) readTokenFile() error {
	for range Only.Once {
		if e.Auth.TokenFile == "" {
			e.Auth.TokenFile, e.Error = os.UserHomeDir()
			if e.Error != nil {
				e.Auth.TokenFile = ""
				break
			}
			e.Auth.TokenFile = filepath.Join(e.Auth.TokenFile, ".GoSungro", DefaultAuthTokenFile)
		}

		var f *os.File
		f, e.Error = os.Open(e.Auth.TokenFile)
		if e.Error != nil {
			if os.IsNotExist(e.Error) {
				e.Error = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		e.Error = json.NewDecoder(f).Decode(&e.Response.ResultData)

		e.Auth.Token = e.GetToken()
	}

	return e.Error
}

// Saves a token to a file path.
func (e *EndPoint) saveToken() error {
	for range Only.Once {
		if e.Auth.TokenFile == "" {
			e.Auth.TokenFile, e.Error = os.UserHomeDir()
			if e.Error != nil {
				e.Auth.TokenFile = ""
				break
			}
			e.Auth.TokenFile = filepath.Join(e.Auth.TokenFile, ".GoSungro", DefaultAuthTokenFile)
		}

		fmt.Printf("Saving token file to: %s\n", e.Auth.TokenFile)
		var f *os.File
		f, e.Error = os.OpenFile(e.Auth.TokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if e.Error != nil {
			e.Error = errors.New(fmt.Sprintf("Unable to cache SUNGRO oauth token: %v", e.Error))
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		e.Error = json.NewEncoder(f).Encode(e.Response.ResultData)
	}

	return e.Error
}
