package web

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/api/reflect"
	"GoSungro/iSolarCloud/sungro/AppService/login"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type Web struct {
	Auth login.EndPoint

	Body  []byte
	Error error
	retry int

	Url      *url.URL
	client      http.Client
	httpRequest  *http.Request
	httpResponse *http.Response
}

func (w *Web) SetUrl(u string) error {
	w.Url, w.Error = url.Parse(u)
	if w.Error != nil {
		w.Url = nil
	}
	w.Auth.Url = w.Url

	return w.Error
}

// func (w *Web) GetOLD(s interface{}, url string, args ...interface{}) error {
// 	for range Only.Once {
// 		if w.Url == nil {
// 			w.Error = errors.New("SUNGRO API URL is invalid")
// 			break
// 		}
//
// 		url = fmt.Sprintf(w.Url.String()+url, args...)
//
// 		w.httpRequest, w.Error = http.NewRequest("GET", url, nil)
// 		if w.Error != nil {
// 			break
// 		}
//
// 		w.httpRequest.Header.Set("Authorization", w.Auth.GetAuthHeader())
//
// 		//w.httpResponse, w.Error = http.Get(url)
// 		w.httpResponse, w.Error = w.client.Do(w.httpRequest)
// 		if w.Error != nil {
// 			break
// 		}
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer w.httpResponse.Body.Close()
// 		if w.httpResponse.StatusCode != 200 {
// 			w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
// 			break
// 		}
//
// 		w.Body, w.Error = ioutil.ReadAll(w.httpResponse.Body)
// 		if w.Error != nil {
// 			break
// 		}
//
// 		if len(w.Body) == 0 {
// 			w.Error = errors.New("empty httpResponse")
// 			break
// 		}
//
// 		w.Error = json.Unmarshal(w.Body, &s)
// 		if w.Error != nil {
// 			fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
// 			break
// 		}
// 	}
//
// 	return w.Error
// }

// func (w *Web) Get() error {
// 	for range Only.Once {
// 		w.HasTokenExpired()
// 		if !w.newToken {
// 			break
// 		}
//
// 		u := fmt.Sprintf("%s%s",
// 			t.Url.String(),
// 			TokenRequestUrl,
// 		)
// 		//p, _ := json.Marshal(map[string]string {
// 		//	"user_account": t.Request.Username,
// 		//	"user_password": t.Request.Password,
// 		//	"appkey": t.Request.AppKey,
// 		//	"sys_code": "900",
// 		//})
// 		p, _ := json.Marshal(t.Request)
//
// 		var httpResponse *http.Response
// 		httpResponse, t.Error = http.Post(u, "application/json", bytes.NewBuffer(p))
// 		if t.Error != nil {
// 			break
// 		}
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer httpResponse.Body.Close()
// 		if httpResponse.StatusCode != 200 {
// 			t.Error = errors.New(fmt.Sprintf("Status Code is %d", httpResponse.StatusCode))
// 			break
// 		}
//
// 		var body []byte
// 		body, t.Error = ioutil.ReadAll(httpResponse.Body)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		t.Error = json.Unmarshal(body, &t.Response)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		t.TokenExpiry = time.Now()
//
// 		t.Error = t.saveToken()
// 		if t.Error != nil {
// 			break
// 		}
// 	}
//
// 	return t.Error
// }

func (w *Web) Get(endpoint api.EndPoint) (api.EndPoint, error) {
// func (w *Web) Get(u *url.URL, request interface{}) (interface{}, error) {
// func (w *Web) Get(action interface{}) error {
	for range Only.Once {
		if w.Url == nil {
			w.Error = errors.New("SUNGRO API URL is invalid")
			break
		}

		request := endpoint.RequestRef()
		w.Error = reflect.VerifyOptionsRequired(request)
		if w.Error != nil {
			break
		}

		// objectName, actionName := GetName(action)
		// httpRequest := FindInStruct(action, "Request")
		// httpResponse := FindInStruct(action, "Response")
		//
		// objectName := GetStructName(object)
		// if objectName == "" {
		// 	w.Error = errors.New("invalid object name to structure")
		// 	break
		// }
		//
		// actionName := GetStructName(action)
		// if objectName == "" {
		// 	w.Error = errors.New("invalid action name to structure")
		// 	break
		// }
		//
		// requestString := Query(httpRequest)
		// if objectName == "" {
		// 	w.Error = errors.New("invalid httpRequest string for structure")
		// 	break
		// }
		//
		// responseString := Query(httpResponse)
		// if objectName == "" {
		// 	w.Error = errors.New("invalid httpResponse string for structure")
		// 	break
		// }
		//
		// u := fmt.Sprintf("%s?format=json&object=%s&action=%s%s",
		// 	w.Url.String(),
		// 	objectName,
		// 	actionName,
		// 	queryString,
		// )

		p, _ := json.Marshal(request)

		// w.httpRequest, w.Error = http.NewRequest("GET", u, nil)
		// if w.Error != nil {
		// 	break
		// }
		//
		// w.httpRequest.Header.Set("Authorization", w.Auth.GetAuthHeader())

		postUrl := fmt.Sprintf("%s%s", w.Url.String(), endpoint.GetUrl())

		for range Only.Twice {
			w.httpResponse, w.Error = http.Post(postUrl, "application/json", bytes.NewBuffer(p))
			if w.Error != nil {
				break
			}

			if strings.Contains(w.httpResponse.Status, "The access token provided is invalid") {
				// 401 Unauthorized The access token provided is invalid.
				// Will first attempt a refresh of the token OR re-login.
				w.Error = w.Auth.Login(&login.SunGroAuth {
					TokenExpiry: "",
					AppKey:      "",
					Username:    "",
					Password:    "",
				})
				if w.Error != nil {
					w.Error = errors.New(w.httpResponse.Status)
					break
				}
				//w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
				continue
			}

			if w.httpResponse.StatusCode == 401 {
				w.Error = errors.New(w.httpResponse.Status)
				break
			}

			// All OK.
			break
		}
		//goland:noinspection GoUnhandledErrorResult
		defer w.httpResponse.Body.Close()
		if w.Error != nil {
			break
		}

		if w.httpResponse.StatusCode != 200 {
			w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
			break
		}

		w.Body, w.Error = ioutil.ReadAll(w.httpResponse.Body)
		if w.Error != nil {
			break
		}

		if len(w.Body) == 0 {
			w.Error = errors.New("empty httpResponse")
			break
		}

		endpoint = endpoint.SetResponse(w.Body)
		w.Error = endpoint.IsResponseValid()

		if w.Error != nil {
			fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}
	}

	return endpoint, w.Error
}

// ******************************************************************************** //

// func (w *Web) GetDomain() string {
// 	return w.Response.ResultMsg
// }
// func (w *Web) VerifyDomain(domain string) string {
// 	if domain == "" {
// 		domain = w.Response.ResultMsg
// 	}
// 	if domain == "." {
// 		domain = w.Response.ResultMsg
// 	}
// 	return domain
// }
//
// func (w *Web) GetUser() string {
// 	return w.Response.ResultMsg
// }
// func (w *Web) VerifyUser(user string) string {
// 	if user == "" {
// 		user = w.Response.ResultData.UserID
// 	}
// 	if user == "." {
// 		user = w.Response.ResultData.UserID
// 	}
// 	return user
// }
//
// func (w *Web) GetUserMac() string {
// 	return "."
// }
// func (w *Web) VerifyUserMac(user string) string {
// 	// @TODO - Check MAC is sane.
// 	return user
// }
//
// func (w *Web) GetUsername() string {
// 	return w.Response.ResultData.UserName
// }
// func (w *Web) VerifyUsername(name string) string {
// 	if name == "" {
// 		name = w.Response.ResultData.UserName
// 	}
// 	if name == "." {
// 		name = w.Response.ResultData.UserName
// 	}
// 	return name
// }
//
// func (w *Web) GetUserEmail() string {
// 	return w.Response.ResultData.Email
// }
// func (w *Web) VerifyUserEmail(email string) string {
// 	if email == "" {
// 		email = w.Response.ResultData.Email
// 	}
// 	if email == "." {
// 		email = w.Response.ResultData.Email
// 	}
// 	return email
// }
//
// func (w *Web) GetDisplayName() string {
// 	return w.Response.ResultData.UserAccount
// }
// func (w *Web) VerifyDisplayName(name string) string {
// 	if name == "" {
// 		name = w.Response.ResultData.UserAccount
// 	}
// 	if name == "." {
// 		name = w.Response.ResultData.UserAccount
// 	}
// 	return name
// }
