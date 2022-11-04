package GoStruct

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
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

func (r *Reflect) SetByIndex(parent Reflect, current Reflect, index int, indexName reflect.Value, name EndPointPath) {
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
				// mk := current.ValueOf.MapKeys()
				// r.Interface = current.ValueOf.MapIndex(mk[index]).Interface()
				r.Interface = current.ValueOf.MapIndex(indexName).Interface()
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
			r.FieldName = current.FieldName		// r.FieldVo.String()
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
			// mk := current.ValueOf.MapKeys()
			// r.FieldVo = current.ValueOf.MapIndex(mk[index])
			r.FieldVo = current.ValueOf.MapIndex(indexName)
			r.FieldName = current.FieldName		// mk[index].String()

			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			// r.DataStructure.Json = current.ValueOf.MapIndex(indexName).String()
			// r.DataStructure.PointId = current.ValueOf.MapIndex(indexName).String()
			r.DataStructure.Json = indexName.String()
			r.DataStructure.PointId = indexName.String()
			// r.DataStructure.Json = r.FieldVo.String()
			// r.DataStructure.PointId = r.FieldVo.String()

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
				pn = reflection.GetPointNameFrom(current.Interface, r.DataStructure.PointNameFromChild, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}

			case r.DataStructure.PointNameFromParent != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = reflection.GetPointNameFrom(parent.Interface, r.DataStructure.PointNameFromParent, intSize, r.DataStructure.PointNameDateFormat)
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
			pn := reflection.GetPointNameFrom(child.Interface, r.DataStructure.PointNameFromChild, 0, r.DataStructure.PointNameDateFormat)
			name = append(name, pn)
		}
	}
	return name
}


func FindStart(fieldName string, Parent Reflect, Current Reflect, name EndPointPath) *Reflect {
	var ret Reflect

	for range Only.Once {
		if Current.Kind == reflect.Pointer {
			// Special case:
			// We're going to change the pointer to a proper object reference.
			if Current.IsNil {
				break
			}
			ref2 := Current.ValueOf.Elem().Interface()
			if valueTypes.IsNil(ref2) {
				break
			}
			Current.SetByFieldName(Current.Interface, ref2, "")
			if Current.IsNil {
				break
			}
			// DO NOT BREAK!
			// KEEP FIRST!
		}

		if Current.Kind == reflect.Struct {
			// Iterate over all available fields and read the tag value
			for si := 0; si < Current.Length; si++ {
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, reflect.Value{}, name)
				if Child.FieldName == fieldName {
					// if !Child.DataStructure.PointNameAppend {
					// 	Child.DataStructure.Endpoint = Child.DataStructure.Endpoint.PopLast()
					// }
					name = name.Append(Child.DataStructure.PointId)
					ret = Child
					break
				}

				if Child.Kind != reflect.Struct {
					continue
				}

				if Child.IsKnown() {
					continue
				}

				Child = *FindStart(fieldName, Current, Child, name)
				// if Child.FieldName == fieldName {
				// 	name = name.Append(Child.DataStructure.PointId)
				// 	ret = Child
				// 	break
				// }
			}
			break
		}

		if Current.Kind == reflect.Slice {
			// Iterate over all available fields and read the tag value
			for si := 0; si < Current.Length; si++ {
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, reflect.Value{}, name)
				if Child.FieldName == fieldName {
					// if !Child.DataStructure.PointNameAppend {
					// 	Child.DataStructure.Endpoint = Child.DataStructure.Endpoint.PopLast()
					// }
					name = name.Append(Child.DataStructure.PointId)
					ret = Child
					break
				}

				if Child.Kind != reflect.Slice {
					continue
				}

				if Child.IsKnown() {
					continue
				}

				Child = *FindStart(fieldName, Current, Child, name)
				// if Child.FieldName == fieldName {
				// 	name = name.Append(Child.DataStructure.PointId)
				// 	ret = Child
				// 	break
				// }
			}
			break
		}

		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported: Type %s\n", Current.FieldName, Current.Kind.String())
	}

	return &ret
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
				Child.SetByIndex(Ref, Ref, i, reflect.Value{}, EndPointPath{})

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
				Child.SetByIndex(Ref, Ref, i, reflect.Value{}, EndPointPath{})

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
				Child.SetByIndex(Ref, Ref, i, reflect.Value{}, EndPointPath{})

				if !Child.IsExported {
					continue
				}

				ret = append(ret, fmt.Sprintf("%v", Child.Interface))
			}
		}
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

// GetOptionsRequired Get field options within the structure that are required.
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

// VerifyOptionsRequired Verify fields within the structure are required.
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
