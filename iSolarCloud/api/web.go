package api

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api/apiReflect"
	// "GoSungro/iSolarCloud/sungro/AppService/login"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


type Web struct {
	// Auth login.EndPoint

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
	// w.Auth.Url = w.Url

	return w.Error
}

func (w *Web) Get(endpoint EndPoint) EndPoint {
	for range Only.Once {
		if w.Url == nil {
			w.Error = errors.New("SUNGRO API URL is invalid")
			break
		}

		request := endpoint.RequestRef()
		w.Error = apiReflect.VerifyOptionsRequired(request)
		if w.Error != nil {
			break
		}

		w.Error = endpoint.IsRequestValid()
		if w.Error != nil {
			break
		}

		j, _ := json.Marshal(request)
		postUrl := fmt.Sprintf("%s%s", w.Url.String(), endpoint.GetUrl())

		w.httpResponse, w.Error = http.Post(postUrl, "application/json", bytes.NewBuffer(j))
		if w.Error != nil {
			break
		}

		if w.httpResponse.StatusCode == 401 {
			w.Error = errors.New(w.httpResponse.Status)
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

	endpoint.SetError("%s", w.Error)
	return endpoint
}

func (w *Web) GetUrl(u string) error {
	w.Url, w.Error = url.Parse(u)
	if w.Error != nil {
		w.Url = nil
	}
	// w.Auth.Url = w.Url

	return w.Error
}

func (w *Web) AppendUrl(endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", w.Url.String(), endpoint)
		ret, w.Error = url.Parse(endpoint)
		if w.Error != nil {
			break
		}
	}
	return ret
}


// ******************************************************************************** //

// func (w *ApiRoot) GetDomain() string {
// 	return w.ResponseCommon.ResultMsg
// }
// func (w *ApiRoot) VerifyDomain(domain string) string {
// 	if domain == "" {
// 		domain = w.ResponseCommon.ResultMsg
// 	}
// 	if domain == "." {
// 		domain = w.ResponseCommon.ResultMsg
// 	}
// 	return domain
// }
//
// func (w *ApiRoot) GetUser() string {
// 	return w.ResponseCommon.ResultMsg
// }
// func (w *ApiRoot) VerifyUser(user string) string {
// 	if user == "" {
// 		user = w.ResponseCommon.ResultData.UserID
// 	}
// 	if user == "." {
// 		user = w.ResponseCommon.ResultData.UserID
// 	}
// 	return user
// }
//
// func (w *ApiRoot) GetUserMac() string {
// 	return "."
// }
// func (w *ApiRoot) VerifyUserMac(user string) string {
// 	// @TODO - Check MAC is sane.
// 	return user
// }
//
// func (w *ApiRoot) GetUsername() string {
// 	return w.ResponseCommon.ResultData.UserName
// }
// func (w *ApiRoot) VerifyUsername(name string) string {
// 	if name == "" {
// 		name = w.ResponseCommon.ResultData.UserName
// 	}
// 	if name == "." {
// 		name = w.ResponseCommon.ResultData.UserName
// 	}
// 	return name
// }
//
// func (w *ApiRoot) GetUserEmail() string {
// 	return w.ResponseCommon.ResultData.Email
// }
// func (w *ApiRoot) VerifyUserEmail(email string) string {
// 	if email == "" {
// 		email = w.ResponseCommon.ResultData.Email
// 	}
// 	if email == "." {
// 		email = w.ResponseCommon.ResultData.Email
// 	}
// 	return email
// }
//
// func (w *ApiRoot) GetDisplayName() string {
// 	return w.ResponseCommon.ResultData.UserAccount
// }
// func (w *ApiRoot) VerifyDisplayName(name string) string {
// 	if name == "" {
// 		name = w.ResponseCommon.ResultData.UserAccount
// 	}
// 	if name == "." {
// 		name = w.ResponseCommon.ResultData.UserAccount
// 	}
// 	return name
// }
