// EndPoint
package login

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/api/apiReflect"
	"encoding/json"
	"errors"
	"fmt"
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
	api.RequestCommon	// login is special as it doesn't have the usual fields.
	RequestData
}

type Response struct {
	api.ResponseCommon
	ResultData ResultData `json:"result_data"`
}


func Init(apiRoot *api.Web) EndPoint {
	return EndPoint {
		EndPointStruct: api.EndPointStruct {
			ApiRoot:  apiRoot,

			Area:     api.GetArea(EndPoint{}),
			Name:     api.GetName(EndPoint{}),
			Url:      api.SetUrl(Url),
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
	return e.Request
}

func (e EndPoint) GetResponse() Response {
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

func (e EndPoint) GetUrl() api.EndPointUrl {
	return e.Url
}

func (e EndPoint) Call() api.EndPoint {
	return e.ApiRoot.Get(e)
}

func (e EndPoint) GetData() api.Json {
	return api.GetAsJson(e.Response.ResultData)
}

func (e EndPoint) SetError(format string, a ...interface{}) api.EndPoint {
	e.EndPointStruct.Error = errors.New(fmt.Sprintf(format, a...))
	return e
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
		if apiReflect.GetPkgType(ref) == "api.RequestCommon" {
			e.Request.RequestCommon = ref.(api.RequestCommon)
			break
		}

		if apiReflect.GetType(ref) == "RequestData" {
			e.Request.RequestData = ref.(RequestData)
			break
		}

		e.Error = apiReflect.DoPkgTypesMatch(e.Request, ref)
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
		// req := e.Request.RequestCommon
		// e.Error = req.RequestCommon.IsValid()
		// if e.Error != nil {
		// 	break
		// }
		e.Error = req.RequestData.IsValid()
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
		// e.ResponseCommon = r
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

func (e EndPoint) String() string {
	ret := e.RequestString()
	ret += e.ResponseString()
	return ret
}

func (e EndPoint) RequestString() string {
	ret := api.GetAsString(e.Request.RequestCommon)
	ret += api.GetAsString(e.Request.RequestData)
	return ret
}

func (e EndPoint) ResponseString() string {
	ret := api.GetAsString(e.Response.ResponseCommon)
	ret += api.GetAsString(e.Response.ResultData)
	return ret
}
