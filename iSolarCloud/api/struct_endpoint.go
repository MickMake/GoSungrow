package api

import (
	"GoSungro/Only"
	"errors"
	"net/url"
)


type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() *url.URL
	Call() EndPoint
	GetData() Json

	SetRequest(ref interface{}) EndPoint	// EndPointStruct
	RequestRef() interface{}
	GetRequestJson() Json
	IsRequestValid() error

	SetResponse([]byte) EndPoint	// EndPointStruct
	ResponseRef() interface{}
	GetResponseJson() Json
	IsResponseValid() error

	SetError(string, ...interface{}) EndPoint
	GetError() error
	IsError() bool
}

type EndPointStruct struct {
	ApiRoot *Web

	Area     AreaName     `json:"area"`
	Name     EndPointName `json:"name"`
	Url      *url.URL     `json:"url"`
	Request  interface{}  `json:"request"`
	Response interface{}  `json:"response"`
	Error    error        `json:"error"`
}

type Common struct {
	Get GetFunc
	Set SetFunc
	// ApiRoot *web.ApiRoot
}

type EndPointName string

type Json string
type GetFunc func() Json
type SetFunc func(j Json)


func (p *EndPointStruct) GetArea() AreaName {
	return p.Area
}

func (p *EndPointStruct) GetName() EndPointName {
	return p.Name
}

func (p *EndPointStruct) GetUrl() *url.URL {
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


// func (p *EndPointStruct) SetFuncPut(fn SetFunc) error {
// 	for range Only.Once {
// 		if fn == nil {
// 			p.Error = errors.New("endpoint put function is nil")
// 			break
// 		}
// 		p.Put = fn
// 	}
// 	return p.Error
// }
//
// func (p *EndPointStruct) SetFuncGet(fn GetFunc) error {
// 	for range Only.Once {
// 		if fn == nil {
// 			p.Error = errors.New("endpoint get function is nil")
// 			break
// 		}
// 		p.Get = fn
// 	}
// 	return p.Error
// }
//
// func (p *EndPointStruct) SetResponse(ref interface{}) error {
// 	for range Only.Once {
// 		if ref == nil {
// 			p.Error = errors.New("endpoint has a nil response structure")
// 			break
// 		}
// 		p.ResponseCommon = ref
// 	}
// 	return p.Error
// }
//
// func (p *EndPointStruct) SetResource(ref interface{}) EndPoint {
// 	for range Only.Once {
// 		p.EndPoint = ref
// 	}
// 	return p.EndPoint
// }
