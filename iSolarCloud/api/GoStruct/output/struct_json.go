package output

import (
	"encoding/json"
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"github.com/MickMake/GoUnify/Only"
)

type Json string

func (req Json) String() string {
	return string(req)
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

		a, e := reflection.GetStructName(r)
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

		a, e := reflection.GetStructName(r)
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
