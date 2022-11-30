package valueTypes

import (
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"regexp"
	"strconv"
	"strings"
)


// @TODO - Consider standardizing points to a known format.
// @TODO - queryUserCurveTemplateData is a good example of the structure.
// type Point struct {
//	PointId    valueTypes.PointId `json:"point_id"`
//	PointName  valueTypes.String  `json:"point_name"`
//	PsId       valueTypes.PsId    `json:"ps_id"`
//	PsKey      valueTypes.PsKey   `json:"ps_key"`
//	Color      valueTypes.String  `json:"color"`
//	DetailId   valueTypes.String  `json:"detail_id"`
//	PsName     valueTypes.String  `json:"ps_name"`
//	Statistics valueTypes.String  `json:"statistics"`
//	Style      valueTypes.String  `json:"style"`
//	Unit       valueTypes.String  `json:"unit"`
// }


type PsKey struct {
	string `json:"string,omitempty"`

	deviceType string `json:"device_type,omitempty"`
	psId       string `json:"ps_id,omitempty"`

	Valid   bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKey) UnmarshalJSON(data []byte) error {

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error != nil {
			break
		}
		t.SetValue(t.string)
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PsKey) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.string)
		if t.Error != nil {
			break
		}
		t.Valid = true
	}

	return data, t.Error
}

func (t *PsKey) Value() string {
	return t.string
}

func (t PsKey) String() string {
	// PsKey can only handle one. So strip the others out.
	a := strings.Split(t.string, ",")
	if len(a) == 0 {
		return t.string
	}
	return a[0]
}

func (t *PsKey) Match(comp string) bool {
	if t.string == comp {
		return true
	}
	return false
}

func (t *PsKey) PsKey() string {
	return t.string
}

func (t *PsKey) SetValue(value string) PsKey {
	for range Only.Once {
		t.string = value
		t.deviceType = ""
		t.psId = ""
		t.Valid = true

		s := strings.Split(value, "_")
		switch {
			case len(s) == 1:
				t.psId = s[0]
			case len(s) >= 2:
				t.psId = s[0]
				t.deviceType = s[1]
		}
	}

	return *t
}

func (t *PsKey) GetDeviceType() string {
	return t.deviceType
}

func (t *PsKey) GetPsId() string {
	return t.psId
}

func SetPsKeyString(value string) PsKey {
	var t PsKey
	return t.SetValue(value)
}


type PsId struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	Valid  bool  `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsId) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from int
		t.Error = json.Unmarshal(data, &t.int64)
		if t.Error == nil {
			t.SetValue(t.int64)
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error == nil {
			t.SetString(t.string)
			break
		}

		t.SetString(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PsId) MarshalJSON() ([]byte, error) {
	var data []byte
	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.int64)
		if t.Error != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatInt(t.int64, 10)
	}

	return data, t.Error
}

func (t PsId) Value() int64 {
	return t.int64
}

func (t PsId) String() string {
	return t.string
}

func (t PsId) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *PsId) SetString(value string) PsId {
	for range Only.Once {
		t.string = value
		t.int64 = 0
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}

		var v int
		v, t.Error = strconv.Atoi(t.string)
		if t.Error != nil {
			break
		}
		t.int64 = int64(v)
		t.Valid = true
	}

	return *t
}

func (t *PsId) SetValue(value int64) PsId {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}


func SetPsIdString(value string) PsId {
	var t PsId
	return t.SetString(value)
}

func SetPsIdValue(value int64) PsId {
	var t PsId
	return t.SetValue(value)
}

func SetPsIdStrings(values []string) PsIds {
	var t PsIds
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range values {
			if psId == "" {
				continue
			}
			t = append(t, SetPsIdString(psId))
		}
	}
	return t
}

func SetPsIdValues(values []int64) PsIds {
	var t PsIds
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range values {
			if psId == 0 {
				continue
			}
			t = append(t, SetPsIdValue(psId))
		}
	}
	return t
}


type PsIds []PsId

func (t *PsIds) Strings() []string {
	var ret []string
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range *t {
			ret = append(ret, psId.String())
		}
	}
	return ret
}


type PointId struct {
	// string `json:"string,omitempty"`
	// int64  `json:"integer,omitempty"`
	// isInt  bool

	Point     string `json:"point"`
	PsKey     PsKey `json:"ps_key"`

	Valid  bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointId) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// pid := strings.TrimPrefix(string(data), "p")
		var d string
		t.Error = json.Unmarshal(data, &d)
		if t.Error != nil {
			break
		}

		t.Set(d)
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PointId) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Error = nil
		if t.PsKey.String() != "" {
			d := fmt.Sprintf(`"%s.%s"`, t.PsKey.String(), t.Point)
			data = []byte(d)
			break
		}
		data = []byte(t.Point)
	}

	return data, t.Error
}

func (t PointId) String() string {
	return t.Point
}

func (t *PointId) Full() string {
	var ret string
	for range Only.Once {
		if t.PsKey.String() != "" {
			ret = fmt.Sprintf(`"%s.%s"`, t.PsKey.String(), t.Point)
			break
		}
		ret = t.Point
	}
	return ret
}

func (t *PointId) Set(value string) PointId {
	for range Only.Once {
		t.PsKey = PsKey{}
		t.Point = ""
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}

		a := strings.Split(value, ".")
		switch {
			case len(a) == 0:
			case len(a) == 1:
				t.Point = a[0]
				t.Valid = true
			case len(a) >= 2:
				t.PsKey.SetValue(a[0])
				t.Point = a[1]
				t.Valid = true
		}
	}

	return *t
}

// func (t *PointId) SetValue(value int64) PointId {
// 	for range Only.Once {
// 		t.string = ""
// 		t.int64 = value
// 		t.Valid = true
// 		t.isInt = true
// 		t.string = strconv.FormatInt(t.int64, 10)
// 	}
//
// 	return *t
// }


// func (t *PointId) Fix() PointId {
// 	for range Only.Once {
// 		p := strings.TrimPrefix(t.string, "p")
// 		_, t.Error = strconv.ParseInt(p, 10, 64)
// 		if t.Error != nil {
// 			t.Valid = false
// 			break
// 		}
// 	}
// 	return *t
// }

func (t *PointId) PointToName() string {
	return PointToName(t.String())
}

func CleanString(s string) string {
	// var ret string
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			(b == '-') ||
			(b == '_') ||
			(b == '.') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	// dupes := regexp.MustCompile(`\s+`)
	// ret = dupes.ReplaceAllString(result.String(), )
	return result.String()
}

func PointToName(ret string) string {
	ret = CleanString(ret)
	ret = strings.ReplaceAll(ret, "_", " ")
	ret = strings.Title(ret)

	// Add space between lowercase and uppercase letters.
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	ret = re.ReplaceAllString(ret, "$1 $2")

	// Lowercase point ids.
	re = regexp.MustCompile("P([0-9]+)")
	ret = re.ReplaceAllString(ret, "p$1")

	return ret
}


func SetPointIdString(value string) PointId {
	var t PointId
	return t.Set(value)
}

// func SetPointIdValue(value int64) PointId {
// 	var t PointId
// 	return t.SetValue(value)
// }


type PsKeys struct {
	PsKeys []PsKey `json:"ps_keys,omitempty"`

	Valid  bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKeys) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		var pid string
		t.Error = json.Unmarshal(data, &pid)
		if t.Error != nil {
			break
		}

		t.Set(strings.Split(pid, ",")...)
		t.Valid = true
	}
	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PsKeys) MarshalJSON() ([]byte, error) {
	var a []string
	for _, pid := range t.PsKeys {
		a = append(a, pid.String())
	}
	ret := `"` + strings.Join(a, ",") + `"`

	return []byte(ret), nil
}

func (t PsKeys) String() string {
	var ret string
	for index := range t.PsKeys {
		ret += t.PsKeys[index].String() + ","
	}
	ret = strings.TrimSuffix(ret, ",")
	return ret
}

func (t *PsKeys) Set(values ...string) PsKeys {
	for _, value := range values {
		for _, v := range strings.Split(value, ",") {
			var pid PsKey
			pid.SetValue(v)
			t.PsKeys = append(t.PsKeys, pid)
		}
	}

	return *t
}

func SetPsKeysString(values string) PsKeys {
	var t PsKeys
	t.Set(strings.Split(values, ",")...)
	return t
}


type PointIds struct {
	PointIds []PointId `json:"points,omitempty"`

	Valid  bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointIds) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		var pid string
		t.Error = json.Unmarshal(data, &pid)
		if t.Error != nil {
			break
		}

		t.Set(strings.Split(pid, ",")...)
		t.Valid = true
	}
	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PointIds) MarshalJSON() ([]byte, error) {
	var a []string
	// Some endpoints can't handle the pskey, so we need to strip it out.
	for _, pid := range t.PointIds {
		a = append(a, pid.String())
	}
	ret := `"` + strings.Join(a, ",") + `"`

	// Can't use this method as some endpoints can't handle the pskey.
	// data, err := json.Marshal(t.PointIds)
	return []byte(ret), nil
}

func (t PointIds) String() string {
	var ret string
	for index := range t.PointIds {
		ret += t.PointIds[index].String() + ","
	}
	ret = strings.TrimSuffix(ret, ",")
	return ret
}

func (t *PointIds) Set(values ...string) PointIds {
	for _, value := range values {
		for _, v := range strings.Split(value, ",") {
			var pid PointId
			pid.Set(v)
			t.PointIds = append(t.PointIds, pid)
		}
	}

	return *t
}

func SetPointIdsString(values string) PointIds {
	var t PointIds
	t.Set(strings.Split(values, ",")...)
	return t
}
