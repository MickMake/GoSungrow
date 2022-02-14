// EndPoint
package getPowerDevicePointNames

import (
	"GoSungro/iSolarCloud/api"
	"fmt"
	"net/url"
)


var Url = "/v1/reportService/getPowerDevicePointNames"

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint api.EndPointStruct

type Request struct {
	api.RequestCommon
	DeviceType string `json:"device_type"`
}

type Response struct {
	api.ResponseCommon
	ResultData ResultData `json:"result_data"`
}

type ResultData   []struct {
	PointCalType int64  `json:"point_cal_type"`
	PointID      int64  `json:"point_id"`
	PointName    string `json:"point_name"`
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
	fmt.Println("g.GetArea()")
	return g.Area
}

func (g EndPoint) GetName() api.EndPointName {
	fmt.Println("g.GetName()")
	return g.Name
}

func (g EndPoint) GetUrl() *url.URL {
	fmt.Println("g.GetUrl()")
	return g.Url
}

func (g EndPoint) SetRequest(ref interface{}) error {
	fmt.Println("g.SetRequest()")
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
	fmt.Println("g.Init()")
	ret := Init()
	return &ret
}

func (g EndPoint) GetError() error {
	return g.Error
}
