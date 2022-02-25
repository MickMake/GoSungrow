package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)


type Api struct{}

// This is used to trim the sub-packages imported under the API.
var thisPackagePath string

func init() {
	for range Only.Once {
		val := reflect.ValueOf(Api{})
		if val.Kind() == reflect.Ptr {
			thisPackagePath = val.Elem().Type().PkgPath()
			break
		}
		thisPackagePath = strings.TrimSuffix(val.Type().PkgPath(), "api")
	}
}

//goland:noinspection GoUnusedExportedFunction
func AppendUrl(host string, endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", host, endpoint)
		ret, _ = url.Parse(endpoint)
	}
	return ret
}

func GetArea(v interface{}) AreaName {
	return AreaName(apiReflect.GetArea(thisPackagePath, v))
}

func GetName(v interface{}) EndPointName {
	return EndPointName(apiReflect.GetName(thisPackagePath, v))
}

func GetUrl(u string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		var err error
		ret, err = url.Parse(u)
		if err != nil {
			ret = nil
		}
	}
	return ret
}

func GetAsJson(r interface{}) Json {
	var ret Json
	for range Only.Once {
		j, err := json.Marshal(r)
		if err != nil {
			ret = Json(fmt.Sprintf("{ \"error\": \"%s\"", err))
			break
		}
		ret = Json(j)
	}
	return ret
}

func GetAsPrettyJson(r interface{}) Json {
	var ret Json
	for range Only.Once {
		j, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			ret = Json(fmt.Sprintf("{ \"error\": \"%s\"", err))
			break
		}
		ret = Json(j)
	}
	return ret
}

//goland:noinspection GoUnusedExportedFunction
func GetAsString(r interface{}) string {
	var ret string
	for range Only.Once {
		j, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			ret = fmt.Sprintf("Error: %s\n", err)
			break
		}

		a, e := apiReflect.GetStructName(r)
		ret += fmt.Sprintf(`"%s.%s": %s`, a, e, j)
	}
	return ret
}

func GetRequestString(r interface{}) string {
	var ret string
	for range Only.Once {
		j, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			ret = fmt.Sprintf("Error: %s\n", err)
			break
		}

		a, e := apiReflect.GetStructName(r)
		ret += fmt.Sprintf(`"%s.%s": %s`, a, e, j)
	}
	return ret
}

func GetEndPointString(r interface{}) string {
	var ret string
	for range Only.Once {
		// endpoint := r.(EndPointStruct)
		j, err := json.MarshalIndent(r, "", "\t")
		// j, err := json.Marshal(r)
		if err != nil {
			ret = fmt.Sprintf("Error: %s\n", err)
			break
		}
		ret = string(j)
	}
	return ret
}
