package getAuthKey

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"errors"
	"fmt"
)

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint struct {
	api.EndPointStruct
	Request  Request
	Response Response
}

type Request struct {
	api.RequestCommon
	RequestData
}

type Response struct {
	api.ResponseCommon
	ResultData ResultData `json:"result_data"`
}

func Init(apiRoot *api.Web) EndPoint {
	return EndPoint{
		EndPointStruct: api.EndPointStruct{
			ApiRoot:  apiRoot,
			Area:     api.GetArea(EndPoint{}),
			Name:     api.GetName(EndPoint{}),
			Url:      api.SetUrl(Url),
			Request:  Request{},
			Response: Response{},
			Error:    nil,
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

func AssertResultData(e api.EndPoint) ResultData {
	return e.(EndPoint).Response.ResultData
}

// ****************************************
// Methods defined by api.EndPoint interface type

func (e EndPoint) Help() string {
	ret := apiReflect.HelpOptions(e.Request.RequestData)
	ret += fmt.Sprintf("JSON request:\t%s\n", e.GetRequestJson())
	ret += e.Request.Help()
	return ret
}

func (e EndPoint) IsDisabled() bool {
	return Disabled
}

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

func (e EndPoint) GetData(raw bool) api.Json {
	if raw {
		return api.Json(e.ApiRoot.Body)
	} else {
		return api.GetAsPrettyJson(e.Response.ResultData)
	}
}

func (e EndPoint) SetError(format string, a ...interface{}) api.EndPoint {
	e.EndPointStruct.Error = errors.New(fmt.Sprintf(format, a...))
	return e
}

func (e EndPoint) GetError() error {
	return e.EndPointStruct.Error
}

func (e EndPoint) IsError() bool {
	if e.Error != nil {
		return true
	}
	return false
}

func (e EndPoint) ReadFile() error {
	return e.FileRead("", &e.Response.ResultData)
}

func (e EndPoint) WriteFile() error {
	return e.FileWrite("", e.Response.ResultData, api.DefaultFileMode)
}

func (e EndPoint) SetRequest(ref interface{}) api.EndPoint {
	for range Only.Once {
		if apiReflect.GetPkgType(ref) == "api.RequestCommon" {
			e.Request.RequestCommon = ref.(api.RequestCommon)
			break
		}

		if apiReflect.GetType(ref) == "RequestData" {
			e.Request.RequestData = ref.(RequestData)
			e.Error = e.IsRequestValid()
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

func (e EndPoint) SetRequestByJson(j api.Json) api.EndPoint {
	for range Only.Once {
		e.Error = json.Unmarshal([]byte(j), &e.Request.RequestData)
		if e.Error != nil {
			break
		}
		e.Error = e.IsRequestValid()
		if e.Error != nil {
			break
		}
	}
	return e
}

func (e EndPoint) RequestRef() interface{} {
	return e.Request
}

func (e EndPoint) GetRequestJson() api.Json {
	return api.GetAsJson(e.Request.RequestData)
}

func (e EndPoint) IsRequestValid() error {
	for range Only.Once {
		// req := e.GetRequest()
		// req := e.Request.RequestCommon
		e.Error = e.Request.RequestCommon.IsValid()
		if e.Error != nil {
			break
		}
		e.Error = e.Request.RequestData.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}

func (e EndPoint) SetResponse(ref []byte) api.EndPoint {
	for range Only.Once {
		e.Error = json.Unmarshal(ref, &e.Response)
		if e.Error != nil {
			break
		}
	}
	return e
}

func (e EndPoint) GetResponseJson() api.Json {
	return api.GetAsPrettyJson(e.Response)
}

func (e EndPoint) ResponseRef() interface{} {
	return e.Response
}

func (e EndPoint) IsResponseValid() error {
	for range Only.Once {
		e.Error = e.Response.ResponseCommon.IsValid()
		if e.Error != nil {
			break
		}
		e.Error = e.Response.ResultData.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}

func (e EndPoint) String() string {
	return api.GetEndPointString(e)
}

func (e EndPoint) RequestString() string {
	return api.GetRequestString(e.Request)
}

func (e EndPoint) ResponseString() string {
	return api.GetRequestString(e.Response)
}

func (e EndPoint) MarshalJSON() ([]byte, error) {
	return api.MarshalJSON(e)

	// return json.Marshal(&struct {
	// 	Area     string   `json:"area"`
	// 	EndPoint string   `json:"endpoint"`
	// 	Host     string   `json:"api_host"`
	// 	Url      string   `json:"endpoint_url"`
	// 	Request  interface{}  `json:"request"`
	// 	Response interface{} `json:"response"`
	// }{
	// 	Area:     string(e.Area),
	// 	EndPoint: string(e.Name),
	// 	Host:     e.ApiRoot.Url.String(),
	// 	Url:      e.Url.String(),
	// 	Request:  e.Request,
	// 	Response: e.Response,
	// })
}
