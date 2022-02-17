// EndPoint
package login

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/api/apiReflect"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)


var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint struct {
	api.EndPointStruct
	Auth *SunGroAuth
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


func Init(apiRoot *api.Web) EndPoint {
	fmt.Println("Init()")
	return EndPoint {
		EndPointStruct: api.EndPointStruct {
			ApiRoot:  apiRoot,

			Area:     api.GetArea(EndPoint{}),
			Name:     api.GetName(EndPoint{}),
			Url:      api.GetUrl(Url),
			Request:  Request{},
			Response: Response{},
			Error:    nil,
		},
		Auth: &SunGroAuth {
			AppKey:       "",
			UserAccount:  "",
			UserPassword: "",
			TokenFile:    DefaultAuthTokenFile,
			expiry:       time.Now(),
			newToken:     false,
			retry:        0,
			err:          nil,
		},
	}
}


// ****************************************
// Methods not scoped by api.EndPoint interface type

func (e EndPoint) Init(apiRoot *api.Web) *EndPoint {
	ret := Init(apiRoot)
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

func Assert(e api.EndPoint) EndPoint {
	return e.(EndPoint)
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

func (e EndPoint) Call() api.EndPoint {
	fmt.Println("e.Call()")
	// foo := e.ApiRoot.Get(e)
	return e.ApiRoot.Get(e)
}

func (e EndPoint) GetData() api.Json {
	// return api.GetAsJson(e.Response.(Response).ResultData)
	return api.GetAsJson(e.Response.ResultData)
}

func (e EndPoint) SetError(format string, a ...interface{}) {
	e.Error = errors.New(fmt.Sprintf(format, a...))
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
		e.Error = apiReflect.DoTypesMatch(e.Request, ref)
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
		// r := e.GetResponse()
		// e.Error = json.Unmarshal(ref, &r)
		e.Error = json.Unmarshal(ref, &e.Response)
		if e.Error != nil {
			break
		}
		// e.Response = r
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
		// resp := e.GetResponse()
		// e.Error = resp.ResponseCommon.IsValid()
		e.Error = e.Response.ResponseCommon.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}
