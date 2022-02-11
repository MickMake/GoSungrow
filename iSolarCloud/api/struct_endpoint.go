package api

import (
	"GoSungro/Only"
	"errors"
	"net/url"
)


type TypeEndPoint struct {
	Area       AreaName     `json:"area"`
	Name       EndPointName `json:"name"`
	Url        *url.URL     `json:"url"`

	Request    interface{}
	Response   interface{}

	Get GetFunc
	Put PutFunc

	Error      error
}

type EndPointName string

type Json string
type GetFunc func() Json
type PutFunc func(j Json)


func (p *TypeEndPoint) IsValid() error {
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

func (p *TypeEndPoint) SetFuncPut(fn PutFunc) error {
	for range Only.Once {
		if fn == nil {
			p.Error = errors.New("endpoint put function is nil")
			break
		}
		p.Put = fn
	}
	return p.Error
}

func (p *TypeEndPoint) SetFuncGet(fn GetFunc) error {
	for range Only.Once {
		if fn == nil {
			p.Error = errors.New("endpoint get function is nil")
			break
		}
		p.Get = fn
	}
	return p.Error
}

func (p *TypeEndPoint) SetResponse(ref interface{}) error {
	for range Only.Once {
		if ref == nil {
			p.Error = errors.New("endpoint has a nil response structure")
			break
		}
		p.Response = ref
	}
	return p.Error
}

func (p *TypeEndPoint) SetRequest(ref interface{}) error {
	for range Only.Once {
		if ref == nil {
			p.Error = errors.New("endpoint has a nil request structure")
			break
		}
		p.Request = ref
	}
	return p.Error
}
