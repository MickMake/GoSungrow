// EndPoint
package nullEndPoint

import (
	"GoSungro/iSolarCloud/api"
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

	foo := EndPoint {
		Area:     api.GetArea(EndPoint{}),
		Name:     api.GetName(EndPoint{}),
		Url:      api.GetUrl(Url),
		Request:  Request{},
		Response: Response{},
		Error:    nil,
	}

	fmt.Printf("endpoint: %v\n", foo)

	return foo
}


func (g EndPoint) GetArea() api.AreaName {
	fmt.Println("g.GetArea() implement me")
	return g.Area
}

func (g EndPoint) GetName() api.EndPointName {
	fmt.Println("g.GetName() implement me")
	return g.Name
}

func (g EndPoint) GetUrl() *url.URL {
	fmt.Println("g.GetUrl() implement me")
	return g.Url
}

func (g EndPoint) SetRequest(ref interface{}) error {
	fmt.Println("g.SetRequest() implement me")
	fmt.Printf("ref == %v\n", ref)
	return nil
}

func (g EndPoint) GetRequest() api.Json {
	return api.GetAsJson(g.Request)
}

func (g EndPoint) GetResponse() api.Json {
	return api.GetAsJson(g.Response)
}

func (g EndPoint) GetData() api.Json {
	fmt.Println("g.GetData() implement me")
	return api.GetAsJson(g.Response.(Response).ResultData)
}

func (g EndPoint) IsValid() error {
	fmt.Println("g.IsValid() implement me")
	return nil
}

func (g EndPoint) Call() api.Json {
	fmt.Println("g.Call() implement me")
	return ""
}

func (g EndPoint) Init() *EndPoint {
	fmt.Println("g.Init() implement me")
	ret := Init()
	return &ret
}

func (g EndPoint) GetError() error {
	fmt.Println("g.IsValid() implement me")
	return g.Error
}
