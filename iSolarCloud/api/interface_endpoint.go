package api


type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() EndPointUrl
	Help() string

	Call() EndPoint
	GetData() Json
	String() string
	SetError(string, ...interface{}) EndPoint
	GetError() error
	IsError() bool
	MarshalJSON() ([]byte, error)

	SetRequest(ref interface{}) EndPoint	// EndPointStruct
	RequestRef() interface{}
	GetRequestJson() Json
	IsRequestValid() error
	RequestString() string

	SetResponse([]byte) EndPoint	// EndPointStruct
	ResponseRef() interface{}
	GetResponseJson() Json
	IsResponseValid() error
	ResponseString() string
}
