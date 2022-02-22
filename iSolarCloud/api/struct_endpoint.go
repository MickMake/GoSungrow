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

func (ep *EndPointStruct) GetArea() AreaName {
	return ep.Area
}

func (ep *EndPointStruct) GetName() EndPointName {
	return ep.Name
}

func (ep *EndPointStruct) GetUrl() EndPointUrl {
	return ep.Url
}

func (ep *EndPointStruct) Call() Json {
	panic("implement me")
}

func (ep *EndPointStruct) SetRequest(ref interface{}) error {
	for range Only.Once {
		if ref == nil {
			ep.Error = errors.New("endpoint has a nil request structure")
			break
		}
		ep.Request = ref
	}
	return ep.Error
}

func (ep *EndPointStruct) GetRequest() Json {
	panic("implement me")
}

func (ep *EndPointStruct) GetResponse() Json {
	panic("implement me")
}

func (ep *EndPointStruct) IsValid() error {
	var err error
	for range Only.Once {
		if ep == nil {
			ep.Error = errors.New("endpoint has a nil structure")
			break
		}
		if ep.Request == nil {
			ep.Error = errors.New("endpoint has a nil request structure")
			break
		}
		if ep.Response == nil {
			ep.Error = errors.New("endpoint has a nil response structure")
			break
		}
	}
	return err
}

func (ep EndPointStruct) String() string {
	var ret string
	for range Only.Once {
		if ep.Name == NullEndPoint {
			break
		}

		ret += fmt.Sprintf("Area:\t%s\nEndPoint:\t%s\nUrl:\t%s\n",
			ep.Area,
			ep.Name,
			ep.Url,
		)

		foo := ep.GetRequest()
		ret += fmt.Sprintf("Request JSON:\t%s\n",
			foo,
		)

		foo = ep.GetResponse()
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
