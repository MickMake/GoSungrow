package reflection

import (
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"reflect"
	"strings"
	"time"
)


func GetPointNameFrom(ref interface{}, name string, intSize int, dateFormat string) string {
	var ret string
	for range Only.Once {
		if dateFormat == "" {
			dateFormat = valueTypes.DateTimeAltLayout
		}
		vo := reflect.ValueOf(ref)

		var ra []string
		switch vo.Kind() {
			case reflect.Struct:
				for _, pnf := range strings.Split(name, ".") {
					// Iterate over all available fields, looking for the field name.
					for i := 0; i < vo.NumField(); i++ {
						fn := vo.Type().Field(i).Name
						if fn == pnf {
							ra = append(ra, valueTypes.AnyToValueString(vo.Field(i).Interface(), intSize, dateFormat))
							break
						}
					}
				}

			case reflect.Map:
				for _, pnf := range strings.Split(name, ".") {
					// Iterate over all available keys, looking for the key name.
					for _, key := range vo.MapKeys() {
						if key.String() == pnf {
							ra = append(ra, valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), intSize, dateFormat))
							break
						}
					}
				}
		}
		ret = strings.Join(ra, ".")
	}

	return ret
}

// func GetStringFrom(ref interface{}, name string) string {
// 	var ret string
// 	for range Only.Once {
// 		vo := reflect.ValueOf(ref)
//
// 		switch vo.Kind() {
// 		case reflect.Struct:
// 			// Iterate over all available fields, looking for the field name.
// 			for i := 0; i < vo.NumField(); i++ {
// 				if vo.Type().Field(i).Name == name {
// 					ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
// 					break
// 				}
// 			}
//
// 		case reflect.Array:
// 			fallthrough
// 		case reflect.Slice:
// 			// Iterate over all available entities, looking for the field name under struct or map.
// 			for i := 0; i < vo.Len(); i++ {
// 				ivo := reflect.ValueOf(vo.Index(i).Interface())
// 				foo := ivo.Interface()
// 				fmt.Sprintf("%s", foo)
// 				foo2 := vo.Index(i).Interface()
// 				fmt.Sprintf("%s", foo2)
// 				switch ivo.Kind() {
// 				case reflect.Struct:
// 					// Iterate over all available fields, looking for the field name.
// 					for ii := 0; ii < ivo.NumField(); ii++ {
// 						if ivo.Type().Field(ii).Name == name {
// 							ret = valueTypes.AnyToValueString(ivo.Field(i).Interface(), 0, "")
// 							break
// 						}
// 					}
// 				case reflect.Map:
// 					// Iterate over map, looking for the key name.
// 					for _, key := range ivo.MapKeys() {
// 						if key.String() == name {
// 							ret = valueTypes.AnyToValueString(ivo.MapIndex(key).Interface(), 0, "")
// 							break
// 						}
// 					}
// 				}
// 				// ret = GetStringFrom(vo.Index(i).Interface(), name)
// 				// if ret != "" {
// 				// 	break
// 				// }
// 			}
//
// 		case reflect.Map:
// 			// Iterate over map, looking for the key name.
// 			for _, key := range vo.MapKeys() {
// 				if key.String() == name {
// 					ret = valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), 0, "")
// 					break
// 				}
// 			}
// 		}
// 	}
//
// 	return ret
// }

const AnyIndex = -1

func GetStringFrom(ref interface{}, index int, name string) string {
	var ret string
	for range Only.Once {
		kind := reflect.ValueOf(ref).Kind()
		if kind == reflect.Struct {
			ret = GetStringFromStruct(ref, name)
			break
		}

		if kind == reflect.Array {
			ret = GetStringFromArray(ref, index, name)
			break
		}

		if kind == reflect.Slice {
			ret = GetStringFromArray(ref, index, name)
			break
		}

		if kind == reflect.Map {
			ret = GetStringFromMap(ref, name)
			break
		}
	}

	return ret
}

func GetStringFromArray(ref interface{}, index int, name string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if (vo.Kind() != reflect.Slice) && (vo.Kind() != reflect.Array) {
			break
		}

		if index == AnyIndex {
			for i := 0; i < vo.Len(); i++ {
				v := vo.Index(index).Interface()
				ivo := reflect.ValueOf(v)
				switch ivo.Kind() {
					case reflect.Struct:
						ret = GetStringFromStruct(v, name)
					case reflect.Map:
						ret = GetStringFromMap(v, name)
					default:
						// Don't descend anything else.
				}
				if ret != "" {
					break
				}
			}
			break
		}

		if index >= vo.Len() {
			break
		}

		v := vo.Index(index).Interface()
		ivo := reflect.ValueOf(v)
		switch ivo.Kind() {
			case reflect.Struct:
				ret = GetStringFromStruct(v, name)
			case reflect.Map:
				ret = GetStringFromMap(v, name)
			default:
				// Don't descend anything else.
		}
	}

	return ret
}

func GetStringFromStruct(ref interface{}, name string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Struct {
			break
		}

		// Iterate over all available fields, looking for the field name.
		for i := 0; i < vo.NumField(); i++ {
			if vo.Type().Field(i).Name == name {
				ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
				break
			}
		}
	}

	return ret
}

func GetStringFromMap(ref interface{}, name string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Map {
			break
		}

		// If the interface has 1 element, then recurse into it - whether map or struct.
		if len(vo.MapKeys()) == 1 {
			key := vo.MapKeys()[0]
			vo = reflect.ValueOf(vo.MapIndex(key).Interface())
		}

		if vo.Kind() == reflect.Struct {
			ret = GetStringFromStruct(vo.Interface(), name)
			break
		}

		if vo.Kind() == reflect.Map {
			// Iterate over map, looking for the key name.
			for _, key := range vo.MapKeys() {
				if key.String() == name {
					ret = valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), 0, "")
					break
				}
			}
		}
	}

	return ret
}


func GetJsonTag(fieldTo reflect.StructField) string {
	var ret string

	for range Only.Once {
		ret = fieldTo.Tag.Get("json")
		ret = strings.ReplaceAll(ret, "omitempty", "")
		ret = strings.TrimSuffix(ret, ",")
	}

	return ret
}

func GetTimestampFrom(ref interface{}, name string, dateFormat string) time.Time {
	var ret time.Time
	for range Only.Once {
		if dateFormat == "" {
			dateFormat = valueTypes.DateTimeAltLayout
		}
		vo := reflect.ValueOf(ref)

		switch vo.Kind() {
			case reflect.Struct:
				// Iterate over all available fields, looking for the field name.
				for i := 0; i < vo.NumField(); i++ {
					if vo.Type().Field(i).Name == name {
						v := fmt.Sprintf("%v", vo.Field(i).Interface())
						ret = valueTypes.SetDateTimeString(v).Time
						break
					}
				}

			case reflect.Map:
				// Iterate over all available fields, looking for the field name.
				for _, key := range vo.MapKeys() {
					if key.String() == name {
						v := fmt.Sprintf("%v", vo.MapIndex(key).Interface())
						ret = valueTypes.SetDateTimeString(v).Time
						break
					}
				}
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

func IsRefZero(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}


func SetFrom(to interface{}, from interface{}) error {
	var err error
	for range Only.Once {
		break
		// to has to be a pointer!
		voSrc := reflect.ValueOf(to)
		for index := 0; index < reflect.ValueOf(to).NumField(); index++ {
			FieldVoFrom := voSrc.Field(index)
			FieldToFrom := voSrc.Type().Field(index)

			if FieldToFrom.IsExported() == false {
				// err = errors.New(fmt.Sprintf("NOT Exported: FieldToSrc.%s\n", FieldToSrc.Name))
				continue
			}

			if FieldVoFrom.IsZero() {
				// if reflection.IsRefZero(FieldVoSrc.Interface()) {
				err = errors.New(fmt.Sprintf("Is Zero: FieldToSrc.%s (%v)\n", FieldToFrom.Name, FieldVoFrom.Interface()))
				continue
			}

			if !FieldVoFrom.IsValid() {
				err = errors.New(fmt.Sprintf("Is NOT Valid: FieldToSrc.%s (%v)\n", FieldToFrom.Name, FieldVoFrom.Interface()))
				continue
			}

			FieldVoTo := reflect.ValueOf(from).Elem().Field(index)
			FieldToTo := reflect.TypeOf(from).Elem().Field(index)
			if !FieldVoTo.CanSet() {
				err = errors.New(fmt.Sprintf("Cannot set: FieldVoDst.%s (%v)\n", FieldToTo.Name, FieldVoTo.Interface()))
				continue
			}

			switch FieldToFrom.Type.String() { // FieldVoSrc.Kind().String()
				case "bool":
					FieldVoTo.SetBool(FieldVoFrom.Bool())

				case "string":
					// if FieldVoSrc.String() == "" {
					// 	break
					// }
					FieldVoTo.SetString(FieldVoFrom.String())

				case "GoStruct.EndPointPath":
					// We're not updating this field.

				case "time.Time":
					// We're not updating this field.

				case "GoStruct.tagStrings":
					// We're not updating this field.

				default:
					_, _ = fmt.Fprintf(os.Stderr,"SetFrom() Unknown type %s (%s) for field '%s' from '%v' to '%v'\n",
						FieldToFrom.Type, FieldVoFrom.Kind().String(), FieldToFrom.Name, FieldVoTo.Interface(), FieldVoFrom.Interface())
			}
		}
	}

	return err
}


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
		sp := strings.Split(ret1, "/")
		ret1 = sp[len(sp)-1]
		ret1 = strings.TrimPrefix(ret1, trim)
		ret2 := val.Type().Name()

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

func GetType(v interface{}) string {
	return reflect.ValueOf(v).Type().Name()
}

func GetPkgType(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
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
