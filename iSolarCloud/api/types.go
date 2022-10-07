package api

import (
	"GoSungrow/Only"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)


type UnitValue struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
	ValueFloat float64 `json:"value_float,omitempty"`
	ValueInt int64 `json:"value_int,omitempty"`
}
type UnitValues []UnitValue
type UnitValueMap map[PointId]UnitValue



func (u *UnitValueMap) Sort() []string {
	var ret []string
	for n := range *u {
		ret = append(ret, string(n))
	}
	sort.Strings(ret)
	return ret
}


func JsonToUnitValue(j string) UnitValue {
	var ret UnitValue

	for range Only.Once {
		err := json.Unmarshal([]byte(j), &ret)
		if err != nil {
			break
		}
	}

	return ret
}

func Float32ToString(num float64) string {
	s := fmt.Sprintf("%.6f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func Float64ToString(num float64) string {
	s := fmt.Sprintf("%.6f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func DivideByThousand(num string) (string, error) {
	fv, err := strconv.ParseFloat(num, 64)
	if err == nil {
		num = Float64ToString(fv / 1000)
	}
	return num, err
}

// func (u *UnitValue) GetStructName() string {
// 	var ret string
// 	apiReflect.GetName()
// 	return ret
// }
//
// func (u *UnitValue) GetJsonName() string {
// 	var ret string
// 	apiReflect.GetOptionsRequired()
// 	return ret
// }
