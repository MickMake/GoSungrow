package valueTypes

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
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

	PsId       string `json:"ps_id"`
	DeviceType string `json:"device_type"`
	DeviceCode string `json:"device_code"`
	ChannelId  string `json:"channel_id"`

	Valid bool  `json:"-"`
	Error error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKey) UnmarshalJSON(data []byte) error {
	for range only.Once {
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

	for range only.Once {
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

func (t *PsKey) PsIdDeviceType() string {
	return t.PsId + "_" + t.DeviceType
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
	for range only.Once {
		t.string = value
		t.DeviceType = ""
		t.DeviceCode = ""
		t.PsId = ""
		t.ChannelId = ""
		t.Valid = false

		s := strings.Split(value, "_")
		switch {
		case len(s) == 1:
			t.PsId = s[0]
			t.Valid = true
		case len(s) == 2:
			t.PsId = s[0]
			t.DeviceType = s[1]
			t.Valid = true
		case len(s) == 3:
			t.PsId = s[0]
			t.DeviceType = s[1]
			t.DeviceCode = s[2]
			t.Valid = true
		case len(s) >= 4:
			t.PsId = s[0]
			t.DeviceType = s[1]
			t.DeviceCode = s[2]
			t.ChannelId = s[3]
			t.Valid = true
		default:
			t.Error = errors.New("invalid ps_key")
		}
	}

	return *t
}

func (t *PsKey) GetPsId() string {
	return t.PsId
}

func (t *PsKey) GetDeviceType() string {
	return t.DeviceType
}

func (t *PsKey) GetDeviceCode() string {
	return t.DeviceCode
}

func (t *PsKey) GetChannelId() string {
	return t.ChannelId
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
	for range only.Once {
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
	for range only.Once {
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
	for range only.Once {
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
	for range only.Once {
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
	for range only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(PsId)
		for _, pids := range values {
			if pids == "" {
				continue
			}
			t = append(t, SetPsIdString(pids))
		}
	}
	return t
}

func SetPsIdValues(values []int64) PsIds {
	var t PsIds
	for range only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(PsId)
		for _, pids := range values {
			if pids == 0 {
				continue
			}
			t = append(t, SetPsIdValue(pids))
		}
	}
	return t
}

type PsIds []PsId

func (t PsIds) String() string {
	var ret string
	for _, pid := range t {
		ret += pid.String() + "\n"
	}
	return ret
}

func (t *PsIds) Strings() []string {
	var ret []string
	for _, pid := range *t {
		ret = append(ret, pid.String())
	}
	return ret
}

type PointId struct {
	// string `json:"string,omitempty"`
	// int64  `json:"integer,omitempty"`
	// isInt  bool

	Point string `json:"point"`
	PsKey PsKey  `json:"ps_key"`

	Valid bool  `json:"valid"`
	Error error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointId) UnmarshalJSON(data []byte) error {
	for range only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		pid := strings.TrimPrefix(string(data), `"`)
		pid = strings.TrimSuffix(pid, `"`)
		pid = strings.TrimPrefix(pid, "p")
		var di int64
		t.Error = json.Unmarshal([]byte(pid), &di)
		if t.Error == nil {
			ds := strconv.FormatInt(di, 10)
			t.Set("p" + ds)
			break
		}

		var ds string
		t.Error = json.Unmarshal([]byte(`"`+pid+`"`), &ds)
		if t.Error == nil {
			t.Set(ds)
			break
		}

		t.Set(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PointId) MarshalJSON() ([]byte, error) {
	var data []byte

	for range only.Once {
		t.Error = nil
		if t.PsKey.String() != "" {
			d := fmt.Sprintf(`"%s.%s"`, t.PsKey.String(), t.Point)
			data = []byte("\"" + d + "\"")
			break
		}
		data = []byte("\"" + t.Point + "\"")
	}

	return data, t.Error
}

func (t PointId) String() string {
	return t.Point
}

func (t *PointId) Full() string {
	var ret string
	for range only.Once {
		if t.PsKey.String() != "" {
			ret = fmt.Sprintf(`%s.%s`, t.PsKey.String(), t.Point)
			break
		}
		ret = t.Point
	}
	return ret
}

func (t *PointId) Set(values ...string) PointId {
	for range only.Once {
		t.PsKey = PsKey{}
		t.Point = ""
		t.Valid = false

		if len(values) == 0 {
			break
		}

		value := strings.Join(values, ".")

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
// 	for range only.Once {
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
// 	for range only.Once {
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

func SetPointIdString(value ...string) PointId {
	var t PointId
	return t.Set(value...)
}

// func SetPointIdValue(value int64) PointId {
// 	var t PointId
// 	return t.SetValue(value)
// }

type PsKeys struct {
	PsKeys []PsKey `json:"ps_keys,omitempty"`

	Valid bool  `json:"valid"`
	Error error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKeys) UnmarshalJSON(data []byte) error {
	for range only.Once {
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

func (t *PsKeys) Join(sep string) string {
	return strings.Join(t.Strings(), sep)
}

func (t *PsKeys) Strings() []string {
	var ret []string
	for _, pskey := range t.PsKeys {
		ret = append(ret, pskey.String())
	}
	return ret
}

func (t *PsKeys) PsIds() []string {
	var ret []string
	for _, pskey := range t.PsKeys {
		ret = append(ret, pskey.PsId)
	}
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

func (t *PsKeys) Match(pskey PsKey) bool {
	var yes bool
	for _, value := range t.PsKeys {
		if pskey.String() == value.String() {
			yes = true
			break
		}
	}
	return yes
}

func (t *PsKeys) MatchPsIdDeviceType(psid string, deviceType string) PsKey {
	var ret PsKey
	for _, value := range t.PsKeys {
		if value.PsId != psid {
			continue
		}
		if value.DeviceType != deviceType {
			continue
		}
		ret = value
		break
	}
	return ret
}

func (t *PsKeys) MatchString(pskey string) bool {
	var yes bool
	for _, value := range t.PsKeys {
		if pskey == value.String() {
			yes = true
			break
		}
	}
	return yes
}

func (t *PsKeys) Length() int {
	return len(t.PsKeys)
}

func SetPsKeysString(values string) PsKeys {
	var t PsKeys
	t.Set(strings.Split(values, ",")...)
	return t
}

type PointIds struct {
	PointIds []PointId `json:"points,omitempty"`

	Valid bool  `json:"valid"`
	Error error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointIds) UnmarshalJSON(data []byte) error {
	for range only.Once {
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
	if len(values) == 0 {
		t.Error = errors.New("no points defined")
		return *t
	}
	for _, value := range values {
		for _, v := range strings.Split(value, ",") {
			var pid PointId
			pid.Set(v)
			t.PointIds = append(t.PointIds, pid)
		}
	}

	return *t
}

func (t *PointIds) PsKeys() *PsKeys {
	var ret PsKeys
	var psks []string
	for _, pskey := range t.PointIds {
		psks = append(psks, pskey.PsKey.String())
	}
	ret = SetPsKeysString(strings.Join(psks, ","))
	return &ret
}

func (t *PointIds) PsIds() PsIds {
	var ret PsIds
	var psids []string
	for _, psid := range t.PointIds {
		psids = append(psids, psid.PsKey.PsId)
	}
	ret = SetPsIdStrings(psids)
	return ret
}

func SetPointIdsString(values ...string) PointIds {
	var t PointIds
	var vals []string
	for _, value := range values {
		vals = append(vals, strings.Split(value, ",")...)
	}
	t.Set(vals...)
	return t
}
