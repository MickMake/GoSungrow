package getPowerStatistics

import (
	"GoSungro/iSolarCloud/api"
	"fmt"
	"net/url"
)


var Url = "/v1/powerStationService/getPowerStatistics"

var _ api.EndPoint = (*EndPoint)(nil)

type EndPoint api.EndPointStruct

type Request struct {
	api.RequestCommon
	PsId string `json:"ps_id"`
}

type Response struct {
	api.ResponseCommon
	ResultData struct {
		PRVlaue  string      `json:"PRVlaue"`
		City     interface{} `json:"city"`
		DayPower struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"dayPower"`
		DesignCapacity struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"design_capacity"`
		EqVlaue     string `json:"eqVlaue"`
		NowCapacity struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"nowCapacity"`
		PsName      string `json:"ps_name"`
		PsShortName string `json:"ps_short_name"`
		Status1     string `json:"status1"`
		Status2     string `json:"status2"`
		Status3     string `json:"status3"`
	} `json:"result_data"`
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


func (g EndPoint) GetArea() api.AreaName {
	return g.Area
}

func (g EndPoint) GetName() api.EndPointName {
	return g.Name
}

func (g EndPoint) GetUrl() *url.URL {
	return g.Url
}

func (g EndPoint) SetRequest(ref interface{}) error {
	g.Request = ref.(Request)
	return nil
}

func (g EndPoint) GetRequest() api.Json {
	return api.GetAsJson(g.Request)
}

func (g EndPoint) GetResponse() api.Json {
	return api.GetAsJson(g.Response)
}

func (g EndPoint) RequestRef() interface{} {
	return g.Request
}

func (g EndPoint) ResponseRef() interface{} {
	return g.Response
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
	ret := Init()
	return &ret
}

func (g EndPoint) GetError() error {
	return g.Error
}
