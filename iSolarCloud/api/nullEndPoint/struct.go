// Package nullEndPoint
// - This file is auto-generated from the update_all.sh script.
// Do not modify anything here. Any changes to this EndPoint should be made in the data.go file.
// The only exception is the AppService.login package.
package nullEndPoint

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)


// api.EndPoint - Import API endpoint interface
var _ api.EndPoint = (*EndPoint)(nil)

// EndPoint - Holds the request, response and web method structures.
type EndPoint struct {
	api.EndPointStruct
	Request  Request
	Response Response
	RawResponse []byte
}

// Request - Holds the api.RequestCommon and user RequestData structures. See data.go for request fields.
type Request struct {
	api.RequestCommon
	RequestData
}

// Response - Holds the api.ResponseCommon and endpoint specific ResultData structures. See data.go for response fields.
type Response struct {
	api.ResponseCommon
	ResultData ResultData `json:"result_data"`
}

// Init - Used to initialize a new endpoint instance. Usually called from an area.
func Init(apiRoot api.Web) EndPoint {
	return EndPoint{
		EndPointStruct: api.EndPointStruct{
			ApiRoot:  apiRoot,
			Area:     api.GetArea(EndPoint{}),
			Name:     api.GetName(EndPoint{}),
			Url:      api.SetUrl(Url),
			Request:  Request{},
			Response: Response{},
			Error:    nil,
		},
	}
}


// ******************************************************************************** //

// Methods not scoped by api.EndPoint interface type

// Init - If the endpoint needs to be re-initialized.
func (e EndPoint) Init(apiRoot api.Web) *EndPoint {
	ret := Init(apiRoot)
	return &ret
}

// GetRequest - Get the Request structure as scoped by this endpoint.
func (e EndPoint) GetRequest() Request {
	return e.Request
}

// GetResponse - Get the Response structure as scoped by this endpoint.
func (e EndPoint) GetResponse() Response {
	return e.Response
}

// Assert - Used to obtain locally scoped EndPoint methods, (not visible from api.EndPoint).
//goland:noinspection GoUnusedExportedFunction
func Assert(e api.EndPoint) EndPoint {
	return e.(EndPoint)
}

// AssertResultData - Used to obtain locally scoped ResultData methods, (not visible from api.EndPoint).
//goland:noinspection GoUnusedExportedFunction
func AssertResultData(e api.EndPoint) ResultData {
	return e.(EndPoint).Response.ResultData
}


// ******************************************************************************** //

// Methods defined by api.EndPoint interface type

// Help - Return help information on the JSON structure used to populate RequestData.
func (e EndPoint) Help() string {
	ret := apiReflect.HelpOptions(e.Request.RequestData)
	ret += fmt.Sprintf("JSON request:\t%s\n", e.GetRequestJson())
	ret += e.Request.Help()
	return ret
}

// IsDisabled - Is this endpoint disabled? See data.go Disabled constant.
func (e EndPoint) IsDisabled() bool {
	return Disabled
}

// GetArea - Returns the API area that this EndPoint is located.
func (e EndPoint) GetArea() api.AreaName {
	return e.Area
}

// GetName - Returns the API EndPoint name.
func (e EndPoint) GetName() api.EndPointName {
	return e.Name
}

// GetUrl - Returns the API EndPoint url.
func (e EndPoint) GetUrl() api.EndPointUrl {
	return e.Url
}

// Call - Once RequestData is populated, this will access the iSolarCloud API and populate ResultData.
func (e EndPoint) Call() api.EndPoint {
	return e.ApiRoot.Get(e)
}

// GetJsonData - Get the JSON representation of ResultData, either as condensed or "pretty".
func (e EndPoint) GetJsonData(raw bool) output.Json {
	if raw {
		// return output.GetAsPrettyJson(string(e.RawResponse))
		return output.Json(e.RawResponse)
	} else {
		return output.GetAsPrettyJson(e.Response.ResultData)
	}
}

// SetError - Set the error code for this EndPoint.
func (e EndPoint) SetError(format string, a ...interface{}) api.EndPoint {
	e.EndPointStruct.Error = errors.New(fmt.Sprintf(format, a...))
	return e
}

// GetError - Get the error code for this EndPoint.
func (e EndPoint) GetError() error {
	return e.EndPointStruct.Error
}

// IsError - Is there an error?
func (e EndPoint) IsError() bool {
	if e.Error != nil {
		return true
	}
	return false
}

// ReadDataFile - Read a JSON file and populate the ResultData structure.
// (File names will default to AREA-ENDPOINT.json )
func (e EndPoint) ReadDataFile() error {
	// return e.FileRead("", &e.Response.ResultData)
	return e.ApiReadDataFile(&e.Response.ResultData)
}

// WriteDataFile - Write to a file, the contents of ResultData as JSON.
// (File names will default to AREA-ENDPOINT.json )
func (e EndPoint) WriteDataFile() error {
	// return e.FileWrite("", e.Response.ResultData, output.DefaultFileMode)
	return e.ApiWriteDataFile(e.Response.ResultData)
}


// ********************************************************************************

// SetRequest - Save an interface reference as either api.RequestCommon or RequestData.
func (e EndPoint) SetRequest(ref interface{}) api.EndPoint {
	for range Only.Once {
		if apiReflect.GetPkgType(ref) == "api.RequestCommon" {
			e.Request.RequestCommon = ref.(api.RequestCommon)
			break
		}

		if apiReflect.GetType(ref) == "RequestData" {
			e.Request.RequestData = ref.(RequestData)
			e.Error = e.IsRequestValid()
			break
		}

		e.Error = apiReflect.DoPkgTypesMatch(e.Request, ref)
		if e.Error != nil {
			break
		}

		e.Request = ref.(Request)
	}
	return e
}

// SetRequestByJson - Save RequestData from a JSON string.
func (e EndPoint) SetRequestByJson(j output.Json) api.EndPoint {
	for range Only.Once {
		e.Error = json.Unmarshal([]byte(j), &e.Request.RequestData)
		if e.Error != nil {
			break
		}
		e.Error = e.IsRequestValid()
		if e.Error != nil {
			break
		}
	}
	return e
}

// RequestRef - Return the locally scoped Request structure.
func (e EndPoint) RequestRef() interface{} {
	return e.Request
}

// GetRequestJson - Return the Request structure as a JSON string.
func (e EndPoint) GetRequestJson() output.Json {
	return output.GetAsJson(e.Request.RequestData)
}

// // GetFingerprint - Used to formulate cache filenames.
// func (e EndPoint) GetFingerprint() string {
// 	return apiReflect.GetFingerprint(e.Request.RequestData)
// }

// IsRequestValid - Is api.RequestCommon and RequestData valid?
func (e EndPoint) IsRequestValid() error {
	for range Only.Once {
		// req := e.GetRequest()
		// req := e.Request.RequestCommon
		e.Error = e.Request.RequestCommon.IsValid()
		if e.Error != nil {
			break
		}
		e.Error = e.Request.RequestData.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}


// ********************************************************************************

// SetResponse - Save a JSON string to the Response structure.
// (Used by the web call method.)
func (e EndPoint) SetResponse(ref []byte) api.EndPoint {
	for range Only.Once {
		e.RawResponse = ref
		e.Error = json.Unmarshal(ref, &e.Response)
		if e.Error != nil {
			break
		}
	}
	return e
}

// GetResponseJson - Return the Response structure as a JSON string.
func (e EndPoint) GetResponseJson() output.Json {
	return output.GetAsPrettyJson(e.Response)
}

// ResponseRef - Return the locally scoped Response structure.
func (e EndPoint) ResponseRef() interface{} {
	return e.Response
}

// IsResponseValid - Is api.ResponseCommon and ResultData valid?
func (e EndPoint) IsResponseValid() error {
	for range Only.Once {
		e.Error = e.Response.ResponseCommon.IsValid()
		if e.Error != nil {
			break
		}
		e.Error = e.Response.ResultData.IsValid()
		if e.Error != nil {
			break
		}
	}
	return e.Error
}

// String - Stringer method for this EndPoint.
func (e EndPoint) String() string {
	return output.GetEndPointString(e)
}

// RequestString - Return the Request structure as a human-readable string.
func (e EndPoint) RequestString() string {
	return output.GetRequestString(e.Request)
}

// ResponseString - Return the Response structure as a human-readable string.
func (e EndPoint) ResponseString() string {
	return output.GetRequestString(e.Response)
}


// ********************************************************************************

// MarshalJSON - Marshall the EndPoint.
func (e EndPoint) MarshalJSON() ([]byte, error) {
	return api.MarshalJSON(e)

	// return json.Marshal(&struct {
	// 	Area     string   `json:"area"`
	// 	EndPoint string   `json:"endpoint"`
	// 	Host     string   `json:"api_host"`
	// 	Url      string   `json:"endpoint_url"`
	// 	Request  interface{}  `json:"request"`
	// 	Response interface{} `json:"response"`
	// }{
	// 	Area:     string(e.Area),
	// 	EndPoint: string(e.Name),
	// 	Host:     e.ApiRoot.Url.String(),
	// 	Url:      e.Url.String(),
	// 	Request:  e.Request,
	// 	Response: e.Response,
	// })
}


// ********************************************************************************

// RequestFingerprint - Check if a cache file exists for this EndPoint.
func (e EndPoint) RequestFingerprint() string {
	return e.ApiFingerprint(e.Request.RequestData)
}

// CacheFilename - Check if a cache file exists for this EndPoint.
func (e EndPoint) CacheFilename() string {
	return e.ApiCacheFilename(e.Request.RequestData)
}

// // CheckCache - Check if a cache file exists for this EndPoint.
// func (e EndPoint) CheckCache() bool {
// 	return e.ApiCheckCache(e.Request.RequestData)
// }
//
// // ReadCache - Read a cache file and return it as an EndPoint structure.
// func (e EndPoint) ReadCache() api.EndPoint {
// 	e.Error = e.ApiReadCache(e.Request.RequestData, &e)
// 	return e
// }
//
// // WriteCache - Write this EndPoint structure out to a cache file.
// func (e EndPoint) WriteCache() error {
// 	return e.ApiWriteCache(e.Request.RequestData, e)
// }

// SetCacheTimeout - Set the cache timeout for this EndPoint. (Defaults to 1 hour.)
func (e EndPoint) SetCacheTimeout(duration time.Duration) api.EndPoint {
	e.ApiRoot.SetCacheTimeout(duration)
	return e
}

// GetCacheTimeout - Return the cache timeout for this EndPoint.
func (e EndPoint) GetCacheTimeout() time.Duration {
	return e.ApiRoot.GetCacheTimeout()
}
