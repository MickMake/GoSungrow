// EndPoint
package nullEndPoint

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"encoding/json"
	"fmt"
	"net/url"
)


var Url = ""

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint api.EndPointStruct

type Request struct {
	api.RequestCommon
}

type Response struct {
	api.ResponseCommon
}

func Init() EndPoint {
	fmt.Println("Init()")
	return EndPoint {
		Area:     api.GetArea(EndPoint{}),
		Name:     api.GetName(EndPoint{}),
		Url:      api.GetUrl(Url),
		Request:  Request{},
		Response: Response{},
		Error:    nil,
	}
}


func (e EndPoint) GetArea() api.AreaName {
	return e.Area
}

func (e EndPoint) GetName() api.EndPointName {
	return e.Name
}

func (e EndPoint) GetUrl() *url.URL {
	return e.Url
}

func (e EndPoint) GetData() api.Json {
	return api.GetAsJson(e.Response.(Response).ResultData)
}

func (e EndPoint) Call() api.Json {
	fmt.Println("e.Call() implement me")
	return ""
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

func (e EndPoint) Init() *EndPoint {
	ret := Init()
	return &ret
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
	return e.GetRequest().RequestCommon.IsValid()
}

func (e EndPoint) GetRequest() Request {
	return e.Request.(Request)
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
	return e.GetResponse().ResponseCommon.IsValid()
}

func (e EndPoint) GetResponse() Response {
	return e.Response.(Response)
}
