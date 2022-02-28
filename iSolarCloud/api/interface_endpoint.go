package api

import (
	"GoSungrow/iSolarCloud/api/output"
	"time"
)

type EndPoint interface {
	GetArea() AreaName
	GetName() EndPointName
	GetUrl() EndPointUrl
	IsDisabled() bool
	Help() string

	Call() EndPoint
	SetError(string, ...interface{}) EndPoint
	GetError() error
	IsError() bool
	MarshalJSON() ([]byte, error)
	ReadDataFile() error
	WriteDataFile() error
	String() string
	GetJsonData(bool) output.Json
	// GetTableData(filter interface{}) Table
	// GetData(bool) Json

	SetRequest(ref interface{}) EndPoint // EndPointStruct
	SetRequestByJson(j output.Json) EndPoint
	RequestRef() interface{}
	GetRequestJson() output.Json
	IsRequestValid() error
	RequestString() string
	RequestFingerprint() string

	SetResponse([]byte) EndPoint // EndPointStruct
	ResponseRef() interface{}
	GetResponseJson() output.Json
	IsResponseValid() error
	ResponseString() string

	// WriteCache() error
	// ReadCache() EndPoint
	// CheckCache() bool
	CacheFilename() string
	SetCacheTimeout(duration time.Duration) EndPoint
	GetCacheTimeout() time.Duration
}
