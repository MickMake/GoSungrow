package api

type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() EndPointUrl
	IsDisabled() bool
	Help() string

	Call() EndPoint
	GetData(bool) Json
	String() string
	SetError(string, ...interface{}) EndPoint
	GetError() error
	IsError() bool
	MarshalJSON() ([]byte, error)
	ReadFile() error
	WriteFile() error

	SetRequest(ref interface{}) EndPoint // EndPointStruct
	SetRequestByJson(j Json) EndPoint
	RequestRef() interface{}
	GetRequestJson() Json
	IsRequestValid() error
	RequestString() string

	SetResponse([]byte) EndPoint // EndPointStruct
	ResponseRef() interface{}
	GetResponseJson() Json
	IsResponseValid() error
	ResponseString() string
}
