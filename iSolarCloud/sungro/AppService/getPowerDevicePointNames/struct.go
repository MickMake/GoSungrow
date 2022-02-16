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

func (e EndPoint) Init() *EndPoint {
	ret := Init()
	return &ret
}


func (e EndPoint) SetRequest(ref interface{}) error {
	e.Request = ref.(Request)
	return nil
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
