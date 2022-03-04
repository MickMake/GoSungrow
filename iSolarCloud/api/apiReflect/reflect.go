package apiReflect

import (
	"GoSungrow/Only"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"

	// "github.com/google/uuid"
	"path/filepath"
	"reflect"
	"strings"
)

// GetArea Return an Area name if we are given an Area or EndPoint struct.
func GetArea(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret1 = strings.TrimPrefix(ret1, trim)
		ret2 := val.Type().Name()
		// ret3 := val.Type().String()
		// fmt.Printf("%s\t%s\t%s\n", ret1, ret2, "")

		if ret2 == "Area" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-1]
			break
		}

		if ret2 == "EndPoint" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-2]
			break
		}

		ret = ret1
	}
	return ret
}

// GetName Return an endpoint name if we are given an Area or EndPoint struct.
func GetName(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret1 = strings.TrimPrefix(ret1, trim)
		ret2 := val.Type().Name()
		// ret3 := val.Type().String()
		// fmt.Printf("%s\t%s\t%s\n", ret1, ret2, "")

		if ret2 == "Area" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-2]
			break
		}

		if ret2 == "EndPoint" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-1]
			break
		}

		ret = ret1
	}
	return ret
}

func GetType(v interface{}) string {
	return reflect.ValueOf(v).Type().Name()
	// var ret string
	// for range Only.Once {
	// 	if v == nil {
	// 		break
	// 	}
	//
	// 	val := reflect.ValueOf(v)
	// 	ret = val.Type().Name()
	// }
	// return ret
}

func GetPkgType(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
	// var ret string
	// for range Only.Once {
	// 	if v == nil {
	// 		break
	// 	}
	//
	// 	val := reflect.ValueOf(v)
	// 	// ret = val.Type().Name()
	// 	ret = val.Type().String()
	// }
	// return ret
}

func GetStructName(v interface{}) (string, string) {
	var area string
	var endpoint string
	for range Only.Once {
		val := reflect.ValueOf(v)
		// ret = val.Type().Name()		// Returns structure, (EndPoint name).
		// ret = val.Type().PkgPath()	// Returns structure path.
		// ret = val.Type().String()	// Returns

		// @TODO - Need to check for pointers to struct
		// 	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
		// 		ret = strings.ToLower(t.Elem().Name())
		// 	} else {
		// 		ret = strings.ToLower(t.Name())
		// 	}

		s := strings.Split(val.Type().String(), ".")
		if len(s) < 2 {
			break
		}
		area = s[0]
		endpoint = s[1]
	}
	return area, endpoint
}

func DoTypesMatch(a interface{}, b interface{}) error {
	var err error
	for range Only.Once {
		aName := GetType(a)
		bName := GetType(b)
		if aName == bName {
			break
		}
		err = errors.New(fmt.Sprintf("interface '%s' doesn't match '%s'", aName, bName))
	}
	return err
}

func DoPkgTypesMatch(a interface{}, b interface{}) error {
	var err error
	for range Only.Once {
		aName := GetPkgType(a)
		bName := GetPkgType(b)
		if aName == bName {
			break
		}
		err = errors.New(fmt.Sprintf("interface '%s' doesn't match '%s'", aName, bName))
	}
	return err
}

func PackageName(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			ret = val.Elem().Type().PkgPath()
			break
		}

		ret2 := val.Type().Name()
		ret3 := val.Type().String()
		ret = val.Type().PkgPath()
		ret = strings.TrimPrefix(ret, trim)

		fmt.Printf("%s\t%s\t%s\n", ret, ret2, ret3)
		fmt.Println("")
	}
	return ret
}

func Query(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	for id := 0; id < s.NumField(); id++ {
		value := fmt.Sprintf("%v", s.Field(id).Interface())
		if value == "" {
			continue
		}
		ret += fmt.Sprintf("&%s=%s",
			typeOf.Field(id).Tag.Get("json"),
			value,
		)
		// fmt.Printf("%d: %s %s = %v\n",
		//	i,
		//	typeOfT.Field(i).Name,
		//	s.Field(i).Type(),
		//	s.Field(i).Interface(),
		// )
	}

	return ret
}

func PrintHeader(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
	case "string":
		ret = fmt.Sprintf("%v", s)
	default:
		for id := 0; id < s.NumField(); id++ {
			// value := fmt.Sprintf("%v", s.Field(id).Interface())
			// if value == "" {
			//	continue
			// }
			ret += fmt.Sprintf("%s (%s),",
				typeOf.Field(id).Name,
				typeOf.Field(id).Tag.Get("json"),
			)
			// fmt.Printf("%d: %s %s = %v\n",
			//	i,
			//	typeOfT.Field(i).Name,
			//	s.Field(i).Type(),
			//	s.Field(i).Interface(),
			// )
		}
	}

	return ret
}

func PrintValue(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = fmt.Sprintf("%v", s)
	default:
		for id := 0; id < s.NumField(); id++ {
			ret += fmt.Sprintf("%v,", s.Field(id).Interface())
		}
	}

	return ret
}

func PrintAsJson(ref interface{}) (string, error) {
	var j string
	var err error

	for range Only.Once {
		var ret []byte
		ret, err = json.MarshalIndent(ref, "", "\t")
		if err != nil {
			break
		}

		j = string(ret)
	}

	return j, err
}

func HeaderAsArray(i interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, "Name")
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%s (%s)",
				typeOf.Field(id).Name,
				typeOf.Field(id).Tag.Get("json"),
			))
		}
	}

	return ret
}
func AsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, fmt.Sprintf("%v", s))
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
		}
	}

	return ret
}

//goland:noinspection GoUnusedFunction
func rowAsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	for id := 0; id < s.NumField(); id++ {
		ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
	}

	return ret
}

// var headerStyleTable = map[string]json2csv.KeyStyle{
//	"jsonpointer": json2csv.JSONPointerStyle,
//	"slash":       json2csv.SlashStyle,
//	"dot":         json2csv.DotNotationStyle,
//	"dot-bracket": json2csv.DotBracketStyle,
// }
//
// func PrintAsCsv(ref interface{}) (string, error) {
//	var c string
//	var err error
//
//	for range Once.Once {
//		var ret []byte
//		ret, err = json.Marshal(ref)
//		if err != nil {
//			break
//		}
//
//		var results []json2csv.KeyValue
//		results, err = json2csv.JSON2CSV(string(ret))
//		if err != nil {
//			log.Fatal(err)
//		}
//		if len(results) == 0 {
//			break
//		}
//
//		csv := json2csv.NewCSVWriter(os.Stdout)
//		// header style (jsonpointer, slash, dot, dot-bracket
//		csv.HeaderStyle = headerStyleTable["dot"]
//		csv.Transpose = true
//
//		err = csv.WriteCSV(results)
//		if err != nil {
//			break
//		}
//
//	}
//
//	return c, err
// }

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func ReflectAsJson(ref interface{}) string {
	var ret string

	for range Only.Once {
		switch reflect.TypeOf(ref).Kind() {
		case reflect.Slice:
		case reflect.Array:
			fmt.Println("The interface is a slice.")
			s := reflect.ValueOf(ref)
			ret += "["
			for i := 0; i < s.Len(); i++ {
				ret += ReflectAsJson(s.Index(i))
			}
			ret += "]"

		case reflect.Struct:
			s := reflect.ValueOf(ref) // .Elem()
			typeOf := s.Type()
			for i := 0; i < s.NumField(); i++ {
				value := fmt.Sprintf("%v", s.Field(i).Interface())
				if value == "" {
					continue
				}
				ret += fmt.Sprintf("%s:%s\n",
					typeOf.Field(i).Tag.Get("json"),
					value,
				)
				// fmt.Printf("%d: %s %s = %v\n",
				//	i,
				//	typeOfT.Field(i).Name,
				//	s.Field(i).Type(),
				//	s.Field(i).Interface(),
				// )
			}
		}
	}

	return ret
}

// FindInStruct Search a given structure for the State object and return its pointer.
//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func FindInStruct(ref interface{}, name string) interface{} {
	var ret interface{}

	for range Only.Once {
		v := reflect.ValueOf(ref)
		var e reflect.Value

		kind := v.Kind()
		// We're doing these checks to ensure ease of future expansion.
		if kind == reflect.Ptr {
			e = v.Elem()
			if e.Kind().String() == "ptr" {
				// PrintflnCyan("POINTER TO POINTER")	@TODO - DEBUG
				ret = FindInStruct(e.Addr().Elem().Interface(), name)
				break
			}
		} else if kind == reflect.Struct {
			// We can't handle a non-pointer, otherwise we get this...
			// reflect.flag.mustBeAssignable using unaddressable value
			e = v
			// Panic(PanicErrorNotGivenAPointer, v.String())
		} else {
			break
		}

		typeOfT := e.Type()
		for i := 0; i < e.NumField(); i++ {
			if typeOfT.Field(i).Name == name {
				ret = typeOfT.Field(i)
				break
			}

			// if typeOfT.Field(i).Name == name {
			//	state = e.Field(i).Interface().(*State)
			//	if state == nil {
			//		e.Field(i).Set(reflect.ValueOf(state))
			//	}
			//	break
			// }
		}
	}

	return ret
}

func GetNameOld(ref interface{}) (string, string) {
	var packageName string
	var structName string

	str := reflect.TypeOf(ref).String()
	str = strings.TrimPrefix(str, "*")
	str = strings.ToLower(str)
	sa := strings.SplitN(str, ".", 2)
	switch len(sa) {
	case 0:
	case 1:
		packageName = sa[0]
	case 2:
		packageName = sa[0]
		structName = sa[1]
	}
	return packageName, structName
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetPackageName(ref interface{}) string {
	str, _ := GetNameOld(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetStructName2(ref interface{}) string {
	str, _ := GetNameOld(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetCmdName2(ref interface{}) string {
	str := reflect.TypeOf(ref).PkgPath()
	str = filepath.Base(str)
	str = strings.ToLower(str)
	return str
}

// func GetStructName(ref interface{}) string {
// 	var ret string
//
// 	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
// 		ret = strings.ToLower(t.Elem().Name())
// 	} else {
// 		ret = strings.ToLower(t.Name())
// 	}
//
// 	return ret
// }

// VerifyOptionsRequired Verify fields within the structure that are required.
func VerifyOptionsRequired(ref interface{}) error {
	var err error

	for range Only.Once {
		// required := GetOptionsRequired(ref)

		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			required := fieldTo.Tag.Get("required")
			if required == "" {
				continue
			}

			fieldVo := vo.Field(i)
			value := fmt.Sprintf("%v", fieldVo.Interface())
			if value == "" {
				err = errors.New(fmt.Sprintf("option '%s' is empty", fieldTo.Name))
				break
			}
		}
	}

	return err
}

func HelpOptions(ref interface{}) string {
	var ret string

	for range Only.Once {
		t := reflect.TypeOf(ref)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			required := field.Tag.Get("required")
			if required == "" {
				ret += fmt.Sprintf("%s: optional\n", field.Name)
				continue
			}

			ret += fmt.Sprintf("%s: required\n", field.Name)
		}
	}

	return ret
}

func FindRequestData(ref interface{}) string {
	var ret string

	for range Only.Once {
		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			// required := fieldTo.Tag.GetByJson("required")
			fmt.Printf(">%s\t", fieldTo.Name)

			fieldVo := vo.Field(i)

			fmt.Printf(">%s\n", fieldVo.String())
			value := fmt.Sprintf("%v", fieldVo.Interface())
			if value == "" {
				break
			}
		}
	}

	return ret
}

func GetRequestString(ref interface{}) string {
	var ret string

	for range Only.Once {
		vo := reflect.ValueOf(ref)
		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldVo := vo.Field(i)
			ret += fmt.Sprintf("-%v", fieldVo.Interface())
		}
	}

	return ret
}

func GetFingerprint(ref interface{}) string {
	var ret string

	for range Only.Once {
		// h := hash(GetRequestString(ref))
		h := md5.Sum([]byte(GetRequestString(ref)))
		ret = fmt.Sprintf("%x", h)
	}

	return ret
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

type Required []string

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetOptionsRequired(ref interface{}) Required {
	var ret []string

	for range Only.Once {
		t := reflect.TypeOf(ref)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			required := field.Tag.Get("required")
			if required == "" {
				continue
			}

			ret = append(ret, field.Name)
		}
	}

	return ret
}

type StructKey struct {
	Name      string
	JsonName  string
	JsonValue string
	Required  string
	Value     string
	Type      reflect.Type
	Interface interface{}
}
type StructKeys map[string]StructKey

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetStructKeys(ref interface{}, keys ...string) StructKeys {
	ret := make(StructKeys)

	s := New(ref)
	if len(keys) == 0 {
		keys = s.Names()
	}

	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	// n := s.Names()
	// fmt.Printf("%v\n", n)
	//
	// n2 := s.Name()
	// fmt.Printf("%v\n", n2)
	//
	// n3 := s.TagName
	// fmt.Printf("%v\n", n3)
	//
	// n4 := s.Fields()
	// fmt.Printf("%v\n", n4)
	//
	// n5 := s.Map()
	// fmt.Printf("%v\n", n5)
	//
	// n6 := s.Values()
	// fmt.Printf("%v\n", n6)
	//
	// n7 := s.Field("")
	// fmt.Printf("%v\n", n7)

	for _, k := range New(ref).Fields() {
		if _, ok := keyMap[k.Name()]; !ok {
			continue
		}

		jValue := ""
		value := ""
		switch k.value.Type().Name() {
			case "string":
				value = k.Value().(string)
			case "int":
				v := k.Value().(int)
				value = strconv.FormatInt(int64(v), 10)
			case "int64":
				value = strconv.FormatInt(k.Value().(int64), 10)
			case "float64":
				value = strconv.FormatFloat(k.Value().(float64), 'f', 6, 64)
			default:
				j, e := json.Marshal(k.Value())
				if e == nil {
					jValue = string(j)
				} else {
					jValue = fmt.Sprintf("%v", k.Value())
				}
		}

		ret[k.Name()] = StructKey {
			Name:      k.Name(),
			JsonName:  k.Tag("json"),
			JsonValue: jValue,
			Value:     value,
			Required:  k.Tag("required"),
			Type:      k.value.Type(),
			Interface: k.Value(),
		}
	}

	return ret
}

func (r *Required) IsRequired(field string) bool {
	var ok bool
	for _, f := range *r {
		if f == field {
			ok = true
		}
	}
	return ok
}
func (r *Required) IsNotRequired(field string) bool {
	return !r.IsRequired(field)
}
