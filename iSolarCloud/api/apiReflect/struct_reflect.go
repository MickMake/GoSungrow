package apiReflect

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"crypto/md5"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)


// Reflect - Combines all the common reflect work into one package
// So we don't have to keep repeating ourselves.
type Reflect struct {
	Valid         bool
	DataStructure DataStructure
	Parent        interface{}
	Interface     interface{}
	IsNil         bool
	IsExported    bool
	isUnknown     bool
	Kind          reflect.Kind
	TypeOf        reflect.Type
	ValueOf       reflect.Value

	Length        int
	FieldName     string
	FieldTo       reflect.StructField
	FieldVo       reflect.Value
}
type ReflectMap map[string]Reflect

func (r Reflect) String() string {
	var ret string
	for range Only.Once {
		fn := r.FieldName
		if strings.HasPrefix(fn, "<struct {") {
			fn = "struct {}"
		}
		switch r.Kind {
			case reflect.Array:
				fallthrough
			case reflect.Map:
				fallthrough
			case reflect.Slice:
				ret = fmt.Sprintf("%s (%s)\tFieldName:%s Kind:%s Length:%d (IsNil:%v IsUnknown:%v IsExported:%v)",
					strings.Join(r.DataStructure.Endpoint, "."),
					r.DataStructure.PointId,
					fn,
					r.Kind.String(),
					r.Length,
					r.IsNil, r.isUnknown, r.IsExported,
				)

			default:
				ret = fmt.Sprintf("%s (%s)\tFieldName:%s Kind:%s (IsNil:%v IsUnknown:%v IsExported:%v)",
					strings.Join(r.DataStructure.Endpoint, "."),
					r.DataStructure.PointId,
					fn,
					r.Kind.String(),
					r.IsNil, r.isUnknown, r.IsExported,
				)
		}
	}
	return ret
}

func (r *Reflect) IsKnown() bool {
	return !r.isUnknown
}

func (r *Reflect) IsUnknown() bool {
	return r.isUnknown
}

func (r *Reflect) SetByFieldName(parent interface{}, current interface{}, name string) {
	for range Only.Once {
		r.Valid = true
		r.Parent = parent
		r.Interface = current
		r.IsNil = valueTypes.IsNil(current)
		r.isUnknown = valueTypes.IsUnknownStruct(current)
		r.TypeOf = reflect.TypeOf(current)
		r.ValueOf = reflect.ValueOf(current)
		r.Kind = r.TypeOf.Kind()
		r.FieldName = name
		r.DataStructure.PointNameAppend = true

		if r.Kind == reflect.Struct {
			r.Length = r.ValueOf.NumField()
		}

		if r.Kind == reflect.Slice {
			r.Length = r.ValueOf.Len()
		}

		if r.Kind == reflect.Array {
			r.Length = r.ValueOf.Len()
		}

		if r.Kind == reflect.Map {
			r.Length = len(r.ValueOf.MapKeys())
		}

		p := reflect.TypeOf(current)
		if p.Kind() != reflect.Struct {
			break
		}


		sf, ok := p.FieldByName(name)
		if !ok {
			break
		}

		r.FieldTo = sf
		r.IsExported = r.FieldTo.IsExported()
		r.FieldVo = reflect.ValueOf(current).FieldByName(name)

		r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
	}
}

func (r *Reflect) SetByIndex(parent Reflect, current Reflect, index int, name EndPointPath) {
	for range Only.Once {
		// Get child interface from parent.
		// pt := current.TypeOf
		// pv := current.ValueOf
		switch current.Kind {
			case reflect.Struct:
				// if !r.FieldTo.IsExported() {
				// 	fmt.Printf("YES")
				// }
				// if !current.TypeOf.Field(index).IsExported() {
				// 	fmt.Printf("YES")
				// }
				r.Interface = current.ValueOf.Field(index).Interface()
			case reflect.Slice:
				r.Interface = current.ValueOf.Index(index).Interface()
			case reflect.Array:
				r.Interface = current.ValueOf.Index(index).Interface()
			case reflect.Map:
				mk := current.ValueOf.MapKeys()
				r.Interface = current.ValueOf.MapIndex(mk[index]).Interface()
		}

		r.Valid = true
		r.Parent = parent.Interface
		r.IsNil = valueTypes.IsNil(r.Interface)
		r.isUnknown = valueTypes.IsUnknownStruct(r.Interface)
		r.TypeOf = reflect.TypeOf(r.Interface)
		r.ValueOf = reflect.ValueOf(r.Interface)
		if r.IsNil {
			r.Kind = reflect.Invalid
		} else {
			r.Kind = r.TypeOf.Kind()
		}

		r.Length = -1
		if r.Kind == reflect.Struct {
			r.Length = r.ValueOf.NumField()
		}
		if r.Kind == reflect.Slice {
			r.Length = r.ValueOf.Len()
		}
		if r.Kind == reflect.Array {
			r.Length = r.ValueOf.Len()
		}
		if r.Kind == reflect.Map {
			r.Length = len(r.ValueOf.MapKeys())
		}

		if current.TypeOf.Kind() == reflect.Struct {
			r.FieldTo = current.TypeOf.Field(index)
			r.IsExported = r.FieldTo.IsExported()
			r.FieldVo = current.ValueOf.Field(index)
			r.FieldName = r.FieldTo.Name

			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, r.DataStructure.PointId)
			break
		}

		if (current.TypeOf.Kind() == reflect.Array) || (current.TypeOf.Kind() == reflect.Slice) {
			// pt = reflect.TypeOf(parent)
			// r.FieldTo = current.TypeOf.Field(index)
			// r.IsExported = r.FieldTo.IsExported()
			r.FieldTo = reflect.StructField{}
			r.IsExported = true
			r.FieldVo = current.ValueOf.Index(index)
			r.FieldName = r.FieldVo.String()
			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			if r.Length == 0 {
				r.DataStructure.PointNameAppend = false
			}

			f := "%d"
			if current.Length > 0 {
				f = fmt.Sprintf("%%.%dd", valueTypes.SizeOfInt(current.Length))
			}
			f = fmt.Sprintf(f, index)

			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, f)

			if r.isUnknown {
				r.DataStructure.Json = current.DataStructure.PointId
				r.DataStructure.PointId = current.DataStructure.PointId
				r.DataStructure.PointUnit = current.DataStructure.PointUnit
				break
			}
			r.DataStructure.Json = current.DataStructure.PointId + "_" + f
			r.DataStructure.PointId = current.DataStructure.PointId + "_" + f
			r.DataStructure.PointUnit = current.DataStructure.PointUnit
			break
		}

		if current.TypeOf.Kind() == reflect.Map {
			// pt = reflect.TypeOf(parent)
			// r.FieldTo = current.TypeOf.Field(index)
			// r.IsExported = r.FieldTo.IsExported()
			r.FieldTo = reflect.StructField{}
			r.IsExported = true
			mk := current.ValueOf.MapKeys()
			r.FieldVo = current.ValueOf.MapIndex(mk[index])
			r.FieldName = mk[index].String()

			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			r.DataStructure.Json = mk[index].String()
			r.DataStructure.PointId = mk[index].String()

			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, r.DataStructure.PointId)
			break
		}
	}
}

// setPointName - Are we using an index number for name or field key value?
func (r *Reflect) setPointName(parent Reflect, current Reflect, name []string, index int) []string {
	for range Only.Once {
		var pn string
		var intSize int
		switch current.Kind {
			case reflect.Array:
				fallthrough
			case reflect.Slice:
				ft := valueTypes.GetIntFormatForPrintf(r.Length)
				pn = fmt.Sprintf(ft, index)
				intSize = valueTypes.SizeOfInt(r.Length)

			case reflect.Map:
				// mk := reflect.ValueOf(current).MapKeys()
				// pn = mk[index].String()
				pn = r.DataStructure.PointId

			case reflect.Struct:
				pn = r.DataStructure.PointId
		}

		switch {
			case r.DataStructure.PointNameFromChild != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(current.Interface, r.DataStructure.PointNameFromChild, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}

			case r.DataStructure.PointNameFromParent != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(parent.Interface, r.DataStructure.PointNameFromParent, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}

			default:
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}
		}
	}

	return name
}

func (r *Reflect) PointNameFromChild(child Reflect, name EndPointPath) []string {
	for range Only.Once {
		if r.DataStructure.PointNameFromChild != "" {
			// PointNameFromChild - In this case points to a field within a CHILD struct.
			pn := GetPointNameFrom(child.Interface, r.DataStructure.PointNameFromChild, 0, r.DataStructure.PointNameDateFormat)
			name = append(name, pn)
		}
	}
	return name
}


func GetStructFields(ref interface{}) map[string]string {
	ret := make(map[string]string)

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref, Ref, i, EndPointPath{})

				if !Child.IsExported {
					continue
				}

				ret[Child.FieldName] = fmt.Sprintf("%v", Child.DataStructure.Required)
			}
			break
		}
	}

	return ret
}

func GetStructFieldsAsArray(ref interface{}) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref, Ref, i, EndPointPath{})

				if !Child.IsExported {
					continue
				}

				ret = append(ret, Child.FieldName)
			}
		}
	}

	return ret
}

func GetStructValuesAsArray(ref interface{}) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref, Ref, i, EndPointPath{})

				if !Child.IsExported {
					continue
				}

				ret = append(ret, fmt.Sprintf("%v", Child.Interface))
			}
		}
	}

	return ret
}

func GetStringFrom(ref interface{}, name string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)

		switch vo.Kind() {
			case reflect.Struct:
				// Iterate over all available fields, looking for the field name.
				for i := 0; i < vo.NumField(); i++ {
					if vo.Type().Field(i).Name == name {
						ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
						break
					}
				}

			case reflect.Map:
				// Iterate over all available fields, looking for the field name.
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

func getJsonTag(fieldTo reflect.StructField) string {
	var ret string

	for range Only.Once {
		ret = fieldTo.Tag.Get("json")
		ret = strings.ReplaceAll(ret, "omitempty", "")
		ret = strings.TrimSuffix(ret, ",")
	}

	return ret
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

func GetType(v interface{}) string {
	return reflect.ValueOf(v).Type().Name()
}

func GetPkgType(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
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

// VerifyOptionsRequired Verify fields within the structure that are required.
func VerifyOptionsRequired(ref interface{}) error {
	var err error

	for range Only.Once {
		// @TODO - Move over to using Reflect structure.
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
			// value := valueTypes.AnyToValueString(fieldVo.Interface(), 0, valueTypes.DateTimeLayoutSecond)
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


type Required []string

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
