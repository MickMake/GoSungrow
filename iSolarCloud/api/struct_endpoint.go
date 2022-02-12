package api

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/web"
	"errors"
	"net/url"
)


// type ResourceType struct {
// 	Area       AreaName     `json:"area"`
// 	Name       EndPointName `json:"name"`
// 	Url        *url.URL     `json:"url"`
// }

type Resource interface {
	Call() Json
	SetRequest(ref interface{}) error
	GetResource() interface{}
}


// type Response interface {
// 	SetFuncPut(fn SetFunc) error
// 	SetFuncGet(fn GetFunc) error
// 	SetResponse(ref interface{})
// }

type Common struct {
	Get GetFunc
	Set SetFunc
	Web *web.Web
}

type RequestCommon struct {
	Get GetFunc
	Set SetFunc

	Appkey     string `json:"appkey"`
	DeviceType string `json:"device_type"`
	Lang       string `json:"lang"`
	SysCode    string `json:"sys_code"`
	Token      string `json:"token"`
	UserID     string `json:"user_id"`
	ValidFlag  string `json:"valid_flag"`
}

type ResponseCommon struct {
	Get GetFunc
	Set SetFunc

	ReqSerialNum string        `json:"req_serial_num"`
	ResultCode   string        `json:"result_code"`
	ResultData   []interface{} `json:"result_data"`
	ResultMsg    string        `json:"result_msg"`
}


type TypeEndPoint struct {
	Area       AreaName     `json:"area"`
	Name       EndPointName `json:"name"`
	Url        *url.URL     `json:"url"`

	Resource interface{}
	Request  interface{}
	Response   interface{}

	Get GetFunc
	Put SetFunc

	Error      error
}

type EndPointName string

type Json string
type GetFunc func() Json
type SetFunc func(j Json)


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

func (p *TypeEndPoint) SetFuncPut(fn SetFunc) error {
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

// func (p *TypeEndPoint) SetResource(ref interface{}) Resource {
// 	for range Only.Once {
// 		p.Resource = ref
// 	}
// 	return p.Resource
// }
