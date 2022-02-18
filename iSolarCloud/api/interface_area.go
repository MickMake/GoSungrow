package api


type Area interface {
	Init(*Web) AreaStruct
	GetAreaName() AreaName
	GetEndPoints() TypeEndPoints
	Call(name EndPointName) Json
	SetRequest(name EndPointName, ref interface{}) error
	GetRequest(name EndPointName) Json
	GetResponse(name EndPointName) Json
	GetData(name EndPointName) Json
	IsValid(name EndPointName) error
	GetError(name EndPointName) error
}

