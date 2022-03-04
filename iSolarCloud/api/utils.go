package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
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

func GetStructKeys(ref interface{}, keys ...string) UnitValueMap {
	ret := make(UnitValueMap)

	for _, k := range apiReflect.GetStructKeys(ref, keys...) {
		p := UnitValue { Value: k.Value, Unit: "" }
		if k.Type.Name() == "UnitValue" {
			// v = JsonToUnitValue(k.JsonValue).Value
			// u = JsonToUnitValue(k.JsonValue).Unit
			p = JsonToUnitValue(k.JsonValue)

			p.Value, p.Unit = DivideByThousandIfRequired(p.Value, p.Unit)
		}

		k.JsonName = strings.TrimSuffix(k.JsonName, "_map")	// Bit of a hack, but hey... @TODO - Future self take note.
		ret[k.JsonName] = p
	}

	return ret
}

// DivideByThousandIfRequired Sigh.... Another dodgy one.
func DivideByThousandIfRequired(value string, unit string) (string, string) {
	switch unit {
		case "Wh":
			fallthrough
		case "W":
			fv, err := strconv.ParseFloat(value, 64)
			if err == nil {
				fv = fv / 1000
				value, _ = DivideByThousand(value)
				unit = "k" + unit
			}
	}

	return value, unit
}
