package api

import (
	"GoSungrow/Only"
	"encoding/json"
	"errors"
	"fmt"
)

type EndPointName string

type EndPointStruct struct {
	ApiRoot *Web `json:"-"`

	Area     AreaName     `json:"area"`
	Name     EndPointName `json:"name"`
	Url      EndPointUrl  `json:"url"`
	Request  interface{}  `json:"-"`
	Response interface{}  `json:"-"`
	Error    error        `json:"-"`
}

func (p *EndPointStruct) GetArea() AreaName {
	return p.Area
}

func (p *EndPointStruct) GetName() EndPointName {
	return p.Name
}

func (p *EndPointStruct) GetUrl() EndPointUrl {
	return p.Url
}

func (p *EndPointStruct) Call() Json {
	panic("implement me")
}

func (p *EndPointStruct) SetRequest(ref interface{}) error {
	for range Only.Once {
		if ref == nil {
			p.Error = errors.New("endpoint has a nil request structure")
			break
		}
		p.Request = ref
	}
	return p.Error
}

func (p *EndPointStruct) GetRequest() Json {
	panic("implement me")
}

func (p *EndPointStruct) GetResponse() Json {
	panic("implement me")
}

func (p *EndPointStruct) IsValid() error {
	var err error
	for range Only.Once {
		if p == nil {
			p.Error = errors.New("endpoint has a nil structure")
			break
		}
		if p.Request == nil {
			p.Error = errors.New("endpoint has a nil request structure")
			break
		}
		if p.Response == nil {
			p.Error = errors.New("endpoint has a nil response structure")
			break
		}
	}
	return err
}

func (p EndPointStruct) String() string {
	var ret string
	for range Only.Once {
		if p.Name == NullEndPoint {
			break
		}

		ret += fmt.Sprintf("Area:\t%s\nEndPoint:\t%s\nUrl:\t%s\n",
			p.Area,
			p.Name,
			p.Url,
		)

		foo := p.GetRequest()
		ret += fmt.Sprintf("Request JSON:\t%s\n",
			foo,
		)

		foo = p.GetResponse()
		ret += fmt.Sprintf("Response JSON:\t%s\n",
			foo,
		)
	}
	return ret
}

// func (p EndPointStruct) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(&struct {
// 		Area     string      `json:"area"`
// 		EndPoint string      `json:"endpoint"`
// 		Host     string      `json:"api_host"`
// 		Url      string      `json:"endpoint_url"`
// 		Request  interface{} `json:"request"`
// 		Response interface{} `json:"response"`
// 	}{
// 		Area:     string(p.Area),
// 		EndPoint: string(p.Name),
// 		Host:     p.ApiRoot.Url.String(),
// 		Url:      p.Url.String(),
// 		Request:  p.Request,
// 		Response: p.Response,
// 	})
// }

func MarshalJSON(endpoint EndPoint) ([]byte, error) {
	e := endpoint.SetError("")
	j, err := json.Marshal(&struct {
		Area     string      `json:"area"`
		EndPoint string      `json:"endpoint"`
		Host     string      `json:"api_host"`
		Url      string      `json:"endpoint_url"`
		Request  interface{} `json:"request"`
		Response interface{} `json:"response"`
	}{
		Area:     string(e.GetArea()),
		EndPoint: string(e.GetName()),
		Host:     e.GetUrl().String(),
		Url:      e.GetUrl().String(),
		Request:  e.RequestRef(),
		Response: e.ResponseRef(),
	})
	return j, err
}
