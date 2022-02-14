package api

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/web"
	"errors"
	"net/url"
)


type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() *url.URL
	Call() Json
	SetRequest(ref interface{}) error
	GetRequest() Json
	GetResponse() Json
	GetData() Json
	IsValid() error
	GetError() error
}


type EndPointStruct struct {
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
	Web *web.Web
}

type RequestCommon struct {
	Appkey     string `json:"appkey"`
	Lang       string `json:"lang"`
	SysCode    string `json:"sys_code"`
	Token      string `json:"token"`
	UserID     string `json:"user_id"`
	ValidFlag  string `json:"valid_flag"`
	// DeviceType string `json:"device_type"`
}

type ResponseCommon struct {
	ReqSerialNum string        `json:"req_serial_num"`
	ResultCode   string        `json:"result_code"`
	ResultData   []interface{} `json:"result_data"`
	ResultMsg    string        `json:"result_msg"`
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
// 		p.Response = ref
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
