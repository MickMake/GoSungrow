//goland:noinspection GoMixedReceiverTypes
package GoStruct

import (
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)


type DataTags struct {
	Endpoint                  EndPointPath `json:"endpoint,omitempty"`

	Required                  bool      `json:"required"`

	Json                      string    `json:"json,omitempty"`

	PointDevice               string    `json:"point_device,omitempty"`
	PointDeviceFrom           string    `json:"point_device_from,omitempty"`
	PointDeviceFromParent     string    `json:"point_device_from_parent,omitempty"`
	PointId                   string    `json:"point_id,omitempty"`

	PointUpdateFreq           string    `json:"point_update_freq,omitempty"`
	PointValueType            string    `json:"point_value_type,omitempty"`

	PointAliasTo              string    `json:"point_alias_to,omitempty"`
	PointVirtual              bool      `json:"point_virtual,omitempty"`
	PointVirtualShift         int64     `json:"point_virtual_shift,omitempty"`

	PointTimestamp            time.Time `json:"point_timestamp,omitempty"`
	PointTimestampFrom        string    `json:"point_timestamp_from,omitempty"`

	PointUnit                 string    `json:"point_unit,omitempty"`
	PointUnitFrom             string    `json:"point_unit_from,omitempty"`
	PointUnitFromParent       string    `json:"point_unit_from_parent,omitempty"`
	PointVariableUnit         bool      `json:"point_variable_unit,omitempty"`

	PointName                 string    `json:"point_name,omitempty"`
	PointIdReplace            bool      `json:"point_name_append,omitempty"`
	PointIdFrom               string    `json:"point_id_from,omitempty"`
	PointIdFromChild          string    `json:"point_name_from_child,omitempty"`
	// PointIdFromParent         string    `json:"point_name_from_parent,omitempty"`
	PointNameDateFormat       string    `json:"point_name_date_format,omitempty"`

	PointIgnore               bool      `json:"point_ignore,omitempty"`
	PointIgnoreZero           bool      `json:"point_ignore_zero,omitempty"`
	PointIgnoreIfNil          string    `json:"point_ignore_if_nil,omitempty"`
	PointIgnoreIfChildFromNil string    `json:"point_ignore_if_child_nil,omitempty"`

	PointGroupName            string    `json:"point_group_name,omitempty"`
	PointGroupNameFrom        string    `json:"point_group_name_from,omitempty"`
	PointArrayFlatten         bool      `json:"point_array_flatten,omitempty"`
	PointListFlatten          bool      `json:"point_list_flatten,omitempty"`

	PointSplitOn              string    `json:"point_split_on,omitempty"`
	PointSplitOnType          string    `json:"point_split_on_type,omitempty"`

	PointValueReplace         string    `json:"point_value_replace,omitempty"`
	PointValueReplaceWith     string    `json:"point_value_replace_with,omitempty"`

	DataTable                 bool      `json:"data_table,omitempty"`
	DataTableChild            bool      `json:"data_table_child,omitempty"`
	DataTablePivot            bool      `json:"data_table_pivot,omitempty"`
	DataTableId               string    `json:"data_table_id,omitempty"`
	DataTableName             string    `json:"data_table_name,omitempty"`
	DataTableTitle            string    `json:"data_table_title,omitempty"`
	DataTableMerge            bool      `json:"data_table_merge,omitempty"`
	DataTableIndex            bool      `json:"data_table_show_index,omitempty"`
	DataTableIndexNames       []string  `json:"data_table_index_names,omitempty"`
	DataTableSortOn           string    `json:"data_table_sort_on,omitempty"`
	DataTableIndexTitle       string    `json:"data_table_index_title,omitempty"`

	ValueType                 string    `json:"value_type,omitempty"`
	ValueKind                 string    `json:"value_kind,omitempty"`
	// MqttFlatten               bool

	StrBool tagStrings `json:"-"`
}
type tagStrings struct {
	Required           string `json:"-"`
	PointIdReplace     string `json:"-"`
	PointIgnore        string `json:"-"`
	PointIgnoreZero    string `json:"-"`
	PointArrayFlatten  string `json:"-"`
	PointListFlatten   string `json:"-"`
	PointVariableUnit  string `json:"-"`
	PointVirtual       string `json:"-"`
	PointVirtualShift  string `json:"-"`
	DataTable          string `json:"-"`
	DataTableChild     string `json:"-"`
	DataTablePivot     string `json:"-"`
	DataTableMerge     string `json:"-"`
	DataTableIndex     string `json:"-"`
	DataTableIndexNames     string `json:"-"`
}

//goland:noinspection GoMixedReceiverTypes
func (ds DataTags) String() string {
	var ret string
	for range Only.Once {
		j, err := json.Marshal(ds)
		if err != nil {
			break
		}
		ret = string(j)
	}
	return ret
}

// GetTags -
func (ds *DataTags) GetTags(fieldTo reflect.StructField, fieldVo reflect.Value) *DataTags {
// func (ds *DataTags) GetTags(parent *Reflect, current *Reflect, child *Reflect, fieldTo reflect.StructField, fieldVo reflect.Value) *DataTags {
	var ret *DataTags
	// , fieldTo reflect.StructField, fieldVo reflect.Value
	for range Only.Once {
		var valueType string
		// if child == nil {
		// 	fieldTo = child.FieldTo
		// 	fieldVo = child.FieldVo
		// }

		if fieldTo.Type != nil {
			valueType = fieldTo.Type.String()
		}

		*ds = DataTags {
			Endpoint:              EndPointPath{},

			Json:                  reflection.GetJsonTag(fieldTo),

			PointDevice:           fieldTo.Tag.Get(PointDevice),
			PointDeviceFrom:       fieldTo.Tag.Get(PointDeviceFrom),
			PointDeviceFromParent: fieldTo.Tag.Get(PointDeviceFromParent),

			PointId:               fieldTo.Tag.Get(PointId),

			PointTimestampFrom:    fieldTo.Tag.Get(PointTimestampFrom),

			PointGroupName:        fieldTo.Tag.Get(PointGroupName),
			PointGroupNameFrom:    fieldTo.Tag.Get(PointGroupNameFrom),

			PointName:           fieldTo.Tag.Get(PointName),
			PointIdFrom:         fieldTo.Tag.Get(PointIdFrom),
			PointIdFromChild:    fieldTo.Tag.Get(PointIdFromChild),
			// PointIdFromParent:   fieldTo.Tag.Get(PointIdFromParent),
			PointNameDateFormat: fieldTo.Tag.Get(PointNameDateFormat),

			PointUpdateFreq:       fieldTo.Tag.Get(PointUpdateFreq),

			PointValueType:        fieldTo.Tag.Get(PointValueType),
			PointValueReplace:     fieldTo.Tag.Get(PointValueReplace),
			PointValueReplaceWith: fieldTo.Tag.Get(PointValueReplaceWith),
			ValueType:             valueType,
			ValueKind:             fieldVo.Kind().String(),

			PointUnit:             fieldTo.Tag.Get(PointUnit),
			PointUnitFrom:         fieldTo.Tag.Get(PointUnitFrom),
			PointUnitFromParent:   fieldTo.Tag.Get(PointUnitFromParent),

			PointAliasTo:          fieldTo.Tag.Get(PointAliasTo),
			// PointVirtual:          fieldTo.Tag.Get(PointVirtual),

			PointIgnoreIfNil:          fieldTo.Tag.Get(PointIgnoreIfNil),
			PointIgnoreIfChildFromNil: fieldTo.Tag.Get(PointIgnoreIfChildFromNil),

			PointSplitOn:          fieldTo.Tag.Get(PointSplitOn),
			PointSplitOnType:      fieldTo.Tag.Get(PointSplitOnType),

			DataTableId:           fieldTo.Tag.Get(DataTableId),
			DataTableName:         fieldTo.Tag.Get(DataTableName),
			DataTableTitle:        fieldTo.Tag.Get(DataTableTitle),
			DataTableSortOn:       fieldTo.Tag.Get(DataTableSortOn),
			DataTableIndexTitle:   fieldTo.Tag.Get(DataTableIndexTitle),
			// DataTableIndexNames:   fieldTo.Tag.Get(DataTableIndexNames),

			StrBool:               tagStrings{
				Required:           fieldTo.Tag.Get("required"),
				PointIdReplace:     fieldTo.Tag.Get(PointIdReplace),
				PointIgnore:        fieldTo.Tag.Get(PointIgnore),
				PointIgnoreZero:    fieldTo.Tag.Get(PointIgnoreZero),
				PointArrayFlatten:  fieldTo.Tag.Get(PointArrayFlatten),
				PointListFlatten:   fieldTo.Tag.Get(PointListFlatten),
				PointVariableUnit:  fieldTo.Tag.Get(PointVariableUnit),
				PointVirtual:       fieldTo.Tag.Get(PointVirtual),
				PointVirtualShift:  fieldTo.Tag.Get(PointVirtualShift),
				DataTable:          fieldTo.Tag.Get(IsDataTable),
				DataTableChild:     fieldTo.Tag.Get(DataTableChild),
				DataTablePivot:     fieldTo.Tag.Get(DataTablePivot),
				DataTableMerge:     fieldTo.Tag.Get(DataTableMerge),
				DataTableIndex:     fieldTo.Tag.Get(DataTableIndex),
				DataTableIndexNames: fieldTo.Tag.Get(DataTableIndexNames),
			},
		}
		ret = ds
	}
	return ret
}

func (ds *DataTags) UpdateTags(parent *Reflect, current *Reflect) *DataTags {
	for range Only.Once {
		ds.PointIgnore = false
		if ds.StrBool.PointIgnore != "" {
			ds.PointIgnore = true
		}

		if ds.PointIgnoreIfNil != "" {
			ret := reflection.GetStringFrom(current.Interface, current.Index, ds.PointIgnoreIfNil, valueTypes.IgnoreLength, ds.PointNameDateFormat)
			if (ret == "") || (ret == "--") {
				ds.PointIgnore = true
			}
		}

		switch ds.StrBool.PointIgnoreZero {
			case "false":
				ds.PointIgnoreZero = false
			case "true":
				ds.PointIgnoreZero = true
			default:
				ds.PointIgnoreZero = true
		}

		ds.Required = false
		if ds.StrBool.Required == "true" {
			ds.Required = true
		}

		ds.PointIdReplace = false
		if ds.StrBool.PointIdReplace == "true" {
			ds.PointIdReplace = true
		}

		ds.PointVirtual = false
		if ds.StrBool.PointVirtual == "true" {
			ds.PointVirtual = true
		}

		ds.PointVirtualShift = 0
		if ds.StrBool.PointVirtualShift != "" {
			var err error
			ds.PointVirtualShift, err = strconv.ParseInt(ds.StrBool.PointVirtualShift, 10, 64)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Invalid number for PointVirtualShift: %s - %s\n", ds.StrBool.PointVirtualShift, err)
			}
		}

		ds.PointArrayFlatten = false
		if ds.StrBool.PointArrayFlatten == "true" {
			ds.PointArrayFlatten = true
		}

		ds.PointListFlatten = false
		if ds.StrBool.PointListFlatten == "true" {
			ds.PointListFlatten = true
		}

		ds.PointVariableUnit = false
		if ds.StrBool.PointVariableUnit == "true" {
			ds.PointVariableUnit = true
		}

		ds.DataTable = false
		if ds.StrBool.DataTable == "true" {
			ds.DataTable = true
		}

		ds.DataTableChild = false
		if ds.StrBool.DataTableChild == "true" {
			ds.DataTableChild = true
		}

		ds.DataTablePivot = false
		if ds.StrBool.DataTablePivot == "true" {
			ds.DataTablePivot = true
		}

		ds.DataTableMerge = false
		if ds.StrBool.DataTableMerge == "true" {
			ds.DataTableMerge = true
		}

		ds.DataTableIndex = false
		if ds.StrBool.DataTableIndex == "true" {
			ds.DataTableIndex = true
		}

		ds.DataTableIndexNames = []string{}
		if ds.StrBool.DataTableIndexNames != "" {
			ds.DataTableIndexNames = strings.Split(ds.StrBool.DataTableIndexNames, ",")
		}

		if ds.PointId == "" {
			ds.PointId = ds.Json
		}

		if ds.PointUnitFrom != "" {
			if ds.PointUnit == "" {
				ds.PointUnit = reflection.GetStringFrom(current.Interface, current.Index, ds.PointUnitFrom, valueTypes.IgnoreLength, ds.PointNameDateFormat)
			}
		}
		if ds.PointUnitFromParent != "" {
			if ds.PointUnit == "" {
				ds.PointUnit = reflection.GetStringFrom(current.ParentReflect.Interface, current.ParentReflect.Index, ds.PointUnitFromParent, valueTypes.IgnoreLength, ds.PointNameDateFormat)
			}
		}

		if ds.PointGroupNameFrom != "" {
			ds.PointGroupName = reflection.GetStringFrom(current.Interface, current.Index, ds.PointGroupNameFrom, valueTypes.IgnoreLength, ds.PointNameDateFormat)
		}

		// if ds.PointTimestamp.IsZero() {
		// 	ds.PointTimestamp = time.Now()
		// }
		if ds.PointTimestampFrom != "" {
			ds.PointTimestamp = reflection.GetTimestampFrom(current.Interface, ds.PointTimestampFrom, valueTypes.DateTimeLayout)
		}

		if ds.PointDeviceFrom != "" {
			ds.PointDevice = reflection.GetStringFrom(current.Interface, current.Index, ds.PointDeviceFrom, valueTypes.IgnoreLength, ds.PointNameDateFormat)
		}
		if ds.PointDeviceFromParent != "" {
			ds.PointDevice = reflection.GetStringFrom(current.ParentReflect.Interface, current.ParentReflect.Index, ds.PointDeviceFromParent, valueTypes.IgnoreLength, ds.PointNameDateFormat)
		}

		if ds.PointName == "" {
			ds.PointName = valueTypes.PointToName(ds.PointId)
		}

		if parent.DataStructure.PointIgnoreIfChildFromNil != "" {
			// Ignore a child if it has a zero/empty value.
			ret := reflection.GetStringFrom(current.Interface, current.Index, parent.DataStructure.PointIgnoreIfChildFromNil, valueTypes.IgnoreLength, ds.PointNameDateFormat)
			if ret != "" {
				ds.PointIgnore = true
			}
		}

		if ds.PointNameDateFormat == "" {
			// ds.PointNameDateFormat = valueTypes.DateTimeLayout
		}

		switch ds.PointUpdateFreq {
			case "UpdateFreqInstant":
				ds.PointUpdateFreq = UpdateFreqInstant
			case "UpdateFreq5Mins":
				ds.PointUpdateFreq = UpdateFreq5Mins
			case "UpdateFreqBoot":
				ds.PointUpdateFreq = UpdateFreqBoot
			case "UpdateFreqDay":
				ds.PointUpdateFreq = UpdateFreqDay
			case "UpdateFreqMonth":
				ds.PointUpdateFreq = UpdateFreqMonth
			case "UpdateFreqYear":
				ds.PointUpdateFreq = UpdateFreqYear
			case "UpdateFreqTotal":
				ds.PointUpdateFreq = UpdateFreqTotal
		}
	}
	return ds
}

// SetFrom - Copy DataTags structure from src to dst with zero/nil checking.
func (ds *DataTags) SetFrom(from *DataTags) error {
	var err error
	for range Only.Once {
		// Don't overwrite important flags.
		from.Json = ""
		from.PointId = ""
		from.PointName = ""
		from.ValueType = ""
		from.ValueKind = ""
		from.Endpoint.Clear()

		voFrom := reflect.ValueOf(*from)
		for index := 0; index < reflect.ValueOf(*from).NumField(); index++ {
			FieldVoFrom := voFrom.Field(index)
			FieldToFrom := voFrom.Type().Field(index)

			if FieldToFrom.IsExported() == false {
				// err = errors.New(fmt.Sprintf("NOT Exported: FieldFrom.%s\n", FieldToFrom.Name))
				continue
			}

			if FieldVoFrom.IsZero() {
				// if reflection.IsRefZero(FieldVoSrc.Interface()) {
				// err = errors.New(fmt.Sprintf("Is Zero: FieldFrom.%s (%v)\n", FieldToFrom.Name, FieldVoFrom.Interface()))
				continue
			}

			if !FieldVoFrom.IsValid() {
				// err = errors.New(fmt.Sprintf("Is NOT Valid: FieldFrom.%s (%v)\n", FieldToFrom.Name, FieldVoFrom.Interface()))
				continue
			}

			// fmt.Printf("TAGS: %s\n", FieldToFrom.Tag)

			FieldVoTo := reflect.ValueOf(ds).Elem().Field(index)
			FieldToTo := reflect.TypeOf(ds).Elem().Field(index)
			if !FieldVoTo.CanSet() {
				err = errors.New(fmt.Sprintf("Cannot set: FieldTo.%s (%v)\n", FieldToTo.Name, FieldVoTo.Interface()))
				break
			}

			switch FieldToFrom.Type.String() { // FieldVoFrom.Kind().String()
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

				case "[]string":
					// We're not updating this field.

				default:
					_, _ = fmt.Fprintf(os.Stderr,"DataTags.SetFrom() Unknown type %s (%s) for field '%s' from '%v' to '%v'\n",
					FieldToFrom.Type, FieldVoFrom.Kind().String(),
					FieldToFrom.Name,
					FieldVoTo.Interface(), FieldVoFrom.Interface())
			}
		}

		if !from.PointTimestamp.IsZero() {
			ds.PointTimestamp = from.PointTimestamp
		}

		// @TODO - Fix this up.
		ds.StrBool.Required = StrSet(from.StrBool.Required, ds.StrBool.Required)
		ds.StrBool.PointIdReplace = StrSet(from.StrBool.PointIdReplace, ds.StrBool.PointIdReplace)
		ds.StrBool.PointIgnore = StrSet(from.StrBool.PointIgnore, ds.StrBool.PointIgnore)
		ds.StrBool.PointIgnoreZero = StrSet(from.StrBool.PointIgnoreZero, ds.StrBool.PointIgnoreZero)
		ds.StrBool.PointArrayFlatten = StrSet(from.StrBool.PointArrayFlatten, ds.StrBool.PointArrayFlatten)
		ds.StrBool.PointListFlatten = StrSet(from.StrBool.PointListFlatten, ds.StrBool.PointListFlatten)
		ds.StrBool.PointVariableUnit = StrSet(from.StrBool.PointVariableUnit, ds.StrBool.PointVariableUnit)
		ds.StrBool.PointVirtual = StrSet(from.StrBool.PointVirtual, ds.StrBool.PointVirtual)
		ds.StrBool.DataTable = StrSet(from.StrBool.DataTable, ds.StrBool.DataTable)
		ds.StrBool.DataTableChild = StrSet(from.StrBool.DataTableChild, ds.StrBool.DataTableChild)
		ds.StrBool.DataTablePivot = StrSet(from.StrBool.DataTablePivot, ds.StrBool.DataTablePivot)
		ds.StrBool.DataTableMerge = StrSet(from.StrBool.DataTableMerge, ds.StrBool.DataTableMerge)
		ds.StrBool.DataTableIndex = StrSet(from.StrBool.DataTableIndex, ds.StrBool.DataTableIndex)
		ds.StrBool.DataTableIndexNames = StrSet(from.StrBool.DataTableIndexNames, ds.StrBool.DataTableIndexNames)
		// ds.Endpoint = src.Endpoint.Copy()
	}

	return err
}

func StrSet(src string, dst string) string {
	if dst != "" {
		return dst
	}
	return src
}


// Reflect - Combines all the common reflect work into one package
// So we don't have to keep repeating ourselves.
type Reflect struct {
	FieldPath       EndPointPath
	ParentReflect   *Reflect
	CurrentReflect  *Reflect
	ChildReflect    []*Reflect
	ChildReflectMap map[string]*Reflect

	Valid           bool
	DataStructure   DataTags
	GoStructs       struct {
		Parent  *DataTags
		Current *DataTags
	}

	Interface      interface{}
	IsNil          bool
	IsExported     bool
	isUnknown      bool
	IsOk           bool
	Value          valueTypes.UnitValues
	InterfaceValue interface{}
	IsStart        bool

	Index         int
	Length        int
	FieldName     string
	Kind          reflect.Kind
	TypeOf        reflect.Type
	ValueOf       reflect.Value
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
		if strings.HasPrefix(fn, "<struct []{") {
			fn = "struct []{}"
		}

		var pid string
		if r.DataStructure.Endpoint.Last() != r.DataStructure.PointId {
			pid = " [" + r.DataStructure.PointId + "]"
		}
		ret = fmt.Sprintf("EndPoint:%s%s\t- FieldPath:%s Kind:%s",
			r.DataStructure.Endpoint.String(),
			pid,
			r.FieldPath.String(),
			r.Kind.String(),
		)
		switch r.Kind {
			case reflect.Array:
				fallthrough
			case reflect.Map:
				fallthrough
			case reflect.Slice:
				ret += fmt.Sprintf(" Length:%d",
					r.Length,
				)
		}
		ret += "\t- "

		if r.IsOk { ret += "IsOk "}
		if r.IsNil { ret += "IsNil "}
		if r.isUnknown { ret += "IsUnknown "}
		if r.IsExported { ret += "IsExported "}

		if r.DataStructure.DataTable { ret += "DataTable "}
		if r.DataStructure.DataTableChild { ret += "DataTableChild "}
		if r.DataStructure.DataTablePivot { ret += "DataTablePivot "}
		if r.DataStructure.PointIdReplace { ret += "PointIdReplace "}
		if r.DataStructure.PointArrayFlatten { ret += "PointArrayFlatten "}
		if r.DataStructure.PointListFlatten { ret += "PointListFlatten "}
		if r.DataStructure.PointVariableUnit { ret += "PointVariableUnit "}
		if r.DataStructure.PointVirtual { ret += "PointVirtual "}
		if r.DataStructure.PointIgnore { ret += "PointIgnore "}
		if r.DataStructure.PointIgnoreZero { ret += "PointIgnoreZero "}
	}
	return ret
}

func (r *Reflect) Name() string {
	var ret string
	for range Only.Once {
		ret = r.DataStructure.DataTableId
		if ret == "" {
			ret = r.FieldPath.String()
		}
		if ret == "" {
			ret = r.DataStructure.Endpoint.String()
		}
	}
	return ret
}

func (r *Reflect) Copy() Reflect {
	var ref Reflect
	for range Only.Once {
		ref = *r
		ref.DataStructure.Endpoint = r.DataStructure.Endpoint.Copy()
	}
	return ref
}

func (r *Reflect) IsKnown() bool {
	return !r.isUnknown
}

func (r *Reflect) IsUnknown() bool {
	return r.isUnknown
}

func (r *Reflect) Init(parent interface{}, current interface{}, name EndPointPath) {
	for range Only.Once {
		r.ParentReflect = &Reflect {
			ParentReflect:  nil,
			CurrentReflect: r,
			ChildReflect:   nil,
			ChildReflectMap: make(map[string]*Reflect, 0),
			Valid:          false,
			DataStructure:  DataTags{},
			Interface:      parent,
			IsNil:          false,
			IsExported:     true,
			isUnknown:      false,
			IsOk:           true,
			Value:          valueTypes.UnitValues{},
			InterfaceValue: parent,
			Index:          0,
			Length:         0,
			FieldName:      "",
			Kind:           0,
			TypeOf:         nil,
			ValueOf:        reflect.Value{},
			FieldTo:        reflect.StructField{},
			FieldVo:        reflect.Value{},
		}
		r.CurrentReflect = r

		r.Valid = true
		r.Interface = current
		r.IsNil = valueTypes.IsNil(r.Interface)
		r.isUnknown = valueTypes.IsUnknownStruct(r.Interface, true)
		r.TypeOf = reflect.TypeOf(r.Interface)
		r.ValueOf = reflect.ValueOf(r.Interface)
		if r.IsNil {
			r.Kind = reflect.Invalid
		} else {
			r.Kind = r.TypeOf.Kind()
		}

		if !name.IsZero() {
			r.DataStructure.Endpoint = name.Copy()
		}
		r.FieldPath = name.Copy()
		r.Index = valueTypes.IgnoreLength
		r.Length = valueTypes.IgnoreLength

		switch r.Kind {
			case reflect.Struct:
				r.Length = r.ValueOf.NumField()
				r.FieldTo = reflect.StructField{}
				r.FieldVo = reflect.Value{}
				r.IsExported = r.FieldTo.IsExported()
				r.FieldName = r.FieldTo.Name

				r.DataStructure.GetTags(reflect.StructField{}, reflect.Value{})	// , r.FieldTo, r.FieldVo)
				r.SetGoStructOptions(1)
				r.DataStructure.UpdateTags(r.ParentReflect, r.CurrentReflect)

				r.InterfaceValue = current
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)

			case reflect.Slice:
				fallthrough
			case reflect.Array:
				r.Length = r.ValueOf.Len()
				r.FieldTo = reflect.StructField{}
				r.FieldVo = reflect.Value{}
				r.IsExported = true
				r.FieldName = r.FieldTo.Name
				r.DataStructure.GetTags(reflect.StructField{}, reflect.Value{})	// , r.FieldTo, r.FieldVo)
				r.SetGoStructOptions(2)
				r.DataStructure.UpdateTags(r.ParentReflect, r.CurrentReflect)
				if r.Length == 0 {
					r.DataStructure.PointIdReplace = true
				}

				r.InterfaceValue = r.FieldVo.Interface()
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)

			case reflect.Map:
				r.Length = len(r.ValueOf.MapKeys())
				r.FieldTo = reflect.StructField{}
				r.FieldVo = reflect.Value{}
				r.IsExported = true
				r.FieldName = r.FieldTo.Name
				// mk := current.ValueOf.MapKeys()
				// r.FieldVo = current.ValueOf.MapIndex(mk[index])
				// r.FieldVo = current.ValueOf.MapIndex(indexName)
				// r.FieldName = current.FieldName		// mk[index].String()

				r.DataStructure.GetTags(reflect.StructField{}, reflect.Value{})	// , r.FieldTo, r.FieldVo)
				r.SetGoStructOptions(1)
				r.DataStructure.UpdateTags(r.ParentReflect, r.CurrentReflect)

				r.InterfaceValue = r.FieldVo.Interface()
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
		}

		// r.SetGoStructOptions()
		// r.DataStructure.UpdateTags(r.ParentReflect, r.CurrentReflect)
		r.SetPointId()
	}
}

func (r *Reflect) SetByIndex(parent *Reflect, current *Reflect, index int, indexName reflect.Value) {
	for range Only.Once {
		r.ParentReflect = parent
		r.CurrentReflect = current
		if current.ChildReflect == nil {
			current.ChildReflect = make([]*Reflect, 0)
		}
		if current.ChildReflectMap == nil {
			current.ChildReflectMap = make(map[string]*Reflect, 0)
		}

		switch current.TypeOf.Kind() {
			case reflect.Struct:
				if !current.TypeOf.Field(index).IsExported() {
					return
				}
				r.Interface = current.ValueOf.Field(index).Interface()

			case reflect.Slice:
				fallthrough
			case reflect.Array:
				r.Interface = current.ValueOf.Index(index).Interface()

			case reflect.Map:
				r.Interface = current.ValueOf.MapIndex(indexName).Interface()
		}

		r.isUnknown = valueTypes.IsUnknownStruct(r.Interface, true)
		r.TypeOf = reflect.TypeOf(r.Interface)
		r.ValueOf = reflect.ValueOf(r.Interface)
		r.IsNil = valueTypes.IsNil(r.Interface)
		if r.IsNil {
			r.Kind = reflect.Invalid
		} else {
			r.Kind = r.TypeOf.Kind()
		}
		r.Index = index

		if r.Kind == reflect.Pointer {
			// Special case:
			// We're going to change the pointer to a proper object reference.
			if !r.IsNil {
				r.Interface = r.ValueOf.Elem().Interface()
				if valueTypes.IsNil(r.Interface) {
					break
				}
				r.isUnknown = valueTypes.IsUnknownStruct(r.Interface, true)
				r.TypeOf = reflect.TypeOf(r.Interface)
				r.ValueOf = reflect.ValueOf(r.Interface)
				r.IsNil = valueTypes.IsNil(r.Interface)
				if r.IsNil {
					r.Kind = reflect.Invalid
				} else {
					r.Kind = r.TypeOf.Kind()
				}
			}
			// DO NOT BREAK!
			// KEEP FIRST!
		}

		switch r.Kind {
			case reflect.Struct:
				r.Length = r.ValueOf.NumField()

			case reflect.Slice:
				fallthrough
			case reflect.Array:
				r.Length = r.ValueOf.Len()

			case reflect.Map:
				r.Length = len(r.ValueOf.MapKeys())

			default:
				r.Length = valueTypes.IgnoreLength
		}
		r.Valid = true

		switch current.TypeOf.Kind() {
			case reflect.Struct:
				r.FieldTo = current.TypeOf.Field(index)
				r.IsExported = r.FieldTo.IsExported()
				r.FieldVo = current.ValueOf.Field(index)
				r.FieldName = r.FieldTo.Name
				r.FieldPath = r.CurrentReflect.FieldPath.Copy()
				r.FieldPath.Append(r.FieldName)
				current.ChildReflect = append(current.ChildReflect, r)
				current.ChildReflectMap[r.FieldName] = r

				// DataStructure
				r.DataStructure.GetTags(r.FieldTo, r.FieldVo)
				r.SetGoStructOptions(1)
				r.DataStructure.UpdateTags(parent, current)

				// Value
				r.InterfaceValue = r.FieldVo.Interface()
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
				r.SetUnit()

			case reflect.Slice:
				fallthrough
			case reflect.Array:
				r.FieldTo = reflect.StructField{}
				if current.Kind == reflect.Struct {
					r.FieldTo = current.FieldTo	// Point to the parent if a struct.
				}
				r.IsExported = true
				r.FieldVo = current.ValueOf.Index(index)
				r.FieldName = current.FieldName + ":" + strconv.Itoa(index)		// r.FieldVo.String()
				r.FieldPath = r.CurrentReflect.FieldPath.Copy()
				r.FieldPath.Append("[" + strconv.Itoa(r.Index) + "]")
				current.ChildReflect = append(current.ChildReflect, r)
				current.ChildReflectMap[r.FieldName] = r

				// DataStructure
				r.DataStructure.GetTags(r.FieldTo, r.FieldVo)
				if r.DataStructure.PointUnit == "" {
					// Only on arrays and maps.
					if current.DataStructure.PointUnit != "" {
						r.DataStructure.PointUnit = current.DataStructure.PointUnit
					}
				}
				if r.DataStructure.PointUnitFromParent == "" {
					// Only on arrays and maps.
					if current.DataStructure.PointUnitFrom != "" {
						r.DataStructure.PointUnitFromParent = current.DataStructure.PointUnitFrom
					}
				}
				r.SetGoStructOptions(2)
				r.DataStructure.UpdateTags(parent, current)
				if r.Length == 0 {
					r.DataStructure.PointIdReplace = true
				}
				r.DataStructure.Json = current.DataStructure.PointId
				r.DataStructure.PointId = current.DataStructure.PointId
				if !r.isUnknown {
					f := "%d"
					if current.Length > 0 {
						f = fmt.Sprintf("%%.%dd", valueTypes.SizeOfInt(current.Length))
					}
					// f = fmt.Sprintf(f, index)
					r.DataStructure.Json = current.DataStructure.PointId + "_" + f
					r.DataStructure.PointId = current.DataStructure.PointId + "_" + f
				}

				// Value
				r.InterfaceValue = r.FieldVo.Interface()
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
				r.SetUnit()

			case reflect.Map:
				r.FieldTo = reflect.StructField{}
				if current.Kind == reflect.Struct {
					r.FieldTo = current.FieldTo	// Point to the parent if a struct.
				}
				r.IsExported = true
				r.FieldVo = current.ValueOf.MapIndex(indexName)
				name := valueTypes.AnyToValueString(indexName.Interface(), valueTypes.IgnoreLength, "")	// map key could be anything.
				r.FieldName = name		// current.FieldName + ":" + indexName.String()





				r.FieldPath = r.CurrentReflect.FieldPath.Copy()
				r.FieldPath.Append("[" + name + "]")		// indexName.String()
				current.ChildReflect = append(current.ChildReflect, r)
				current.ChildReflectMap[r.FieldName] = r

				// DataStructure
				r.DataStructure.GetTags(r.FieldTo, r.FieldVo)
				if r.DataStructure.PointUnit == "" {
					// Only on arrays and maps.
					if current.DataStructure.PointUnit != "" {
						r.DataStructure.PointUnit = current.DataStructure.PointUnit
					}
				}
				if r.DataStructure.PointUnitFromParent == "" {
					// Only on arrays and maps.
					if current.DataStructure.PointUnitFrom != "" {
						r.DataStructure.PointUnitFromParent = current.DataStructure.PointUnitFrom
					}
				}
				r.SetGoStructOptions(1)
				r.DataStructure.Json = name	// indexName.String()		// current.ValueOf.MapIndex(indexName).String() || r.FieldVo.String()
				r.DataStructure.PointId = name	// indexName.String()	// current.ValueOf.MapIndex(indexName).String() || r.FieldVo.String()
				r.DataStructure.UpdateTags(parent, current)

				// Value
				r.InterfaceValue = r.FieldVo.Interface()
				r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
					// map[string]interface{}{ indexName.String(): r.InterfaceValue }, r.DataStructure.PointUnit,
					r.InterfaceValue, "", r.DataStructure.PointUnit,
					r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
				r.SetUnit()


			default:
				r.Interface = current.Interface
		}

		r.SetPointId()
	}
}

func (r *Reflect) SetValue(value interface{}) {
	for range Only.Once {
		r.InterfaceValue = value
		r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
			value, "", r.Value.GetUnit(),
			r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
	}
}

func (r *Reflect) SetUnit() {
	for range Only.Once {
		switch {
			case r.Value.GetUnit() == "":
				r.Value.SetUnit(r.DataStructure.PointUnit)
			default:
				r.DataStructure.PointUnit = r.Value.GetUnit()
				r.DataStructure.PointUnitFrom = ""
				r.DataStructure.PointUnitFromParent = ""
		}
	}
}

// func (r *Reflect) ConvertToState() {
// 	for range Only.Once {
// 		var ok bool
//
// 		switch {
// 			case r.Value.First().ValueFloat() > 0:
// 				ok = true
// 			case r.Value.First().ValueFloat() < 0:
// 				ok = true
// 			case r.Value.First().ValueInt() > 0:
// 				ok = true
// 			case r.Value.First().ValueInt() < 0:
// 				ok = true
// 		}
// 		r.Value = valueTypes.SetUnitValueBool(ok)
// 	}
// }

func (r *Reflect) ValuesRange() []valueTypes.UnitValue {
	return r.Value.Range(valueTypes.SortOrder)
}

// func (r *Reflect) SetValues(values ...interface{}) {
// 	for range Only.Once {
// 		r.InterfaceValue = values
// 		var uvs valueTypes.UnitValues
// 		for _, value := range values {
// 			var uvs2 valueTypes.UnitValues
// 			uvs2, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
// 				value, "", r.DataStructure.PointUnit,
// 				r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
// 			uvs.AddUnitValues(uvs2)
// 		}
// 	}
// }

func (r *Reflect) SetUnitValue(value valueTypes.UnitValue) {
	for range Only.Once {
		r.InterfaceValue = value
		r.Value.Reset()
		r.Value.AddUnitValue(value.ValueKey(), value)
		// r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
		// 	value, "", r.DataStructure.PointUnit,
		// 	r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
	}
}

func (r *Reflect) SetUnitValues(value valueTypes.UnitValues) {
	for range Only.Once {
		r.InterfaceValue = value
		r.Value = value
		// r.Value, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
		// 	value, "", r.DataStructure.PointUnit,
		// 	r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
	}
}

// SetPointId - Sets the EndPointPath based off struct tags?
func (r *Reflect) SetPointId() EndPointPath {
	for range Only.Once {
		r.DataStructure.Endpoint = r.CurrentReflect.CopyEndPointPath()
		// fmt.Printf("EPP(BEFORE): %s\n", r.DataStructure.Endpoint.String())
		// fmt.Printf("[                 ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
		var pn string
		if strings.Contains(r.EndPointPath().String(), "Points") || strings.Contains(r.DataStructure.Endpoint.String(), "values") {
			fmt.Printf("")
		}

		switch {
			case r.DataStructure.PointIdFromChild != "":
				// PointIdFromChild - In this case points to a field within a CHILD struct.
				// @TODO - This needs fixing for arrays! Will always return the first entry found instead of the correct one.
				// fmt.Printf("[PointIdFromChild1]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				var pns []string
				for _, pid := range strings.Split(r.DataStructure.PointIdFromChild, ".") {
					p := reflection.GetStringFrom(r.Interface, r.Index, pid, valueTypes.IgnoreLength, r.DataStructure.PointNameDateFormat)	// Look forward into structure.
					if p == "" {
						continue
					}
					pns = append(pns, p)
				}
				pn = strings.Join(pns, ".")	// pns = append(pns, r.DataStructure.PointId)
				r.DataStructure.Endpoint.Append(r.DataStructure.PointId)
				// fmt.Printf("[PointIdFromChild2]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				if r.DataStructure.PointIdReplace {
					r.DataStructure.Endpoint.PopLast()
					// fmt.Printf("[PointIdFromChild3]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				}
				if pn != "" {
					r.DataStructure.Endpoint.Append(pn)
					// fmt.Printf("[PointIdFromChild4]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				}

			case r.DataStructure.PointIdFrom != "":
				// PointIdFromChild - In this case points to a field within a CHILD struct.
				// fmt.Printf("[PointIdFrom1     ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				var pns []string
				// al := valueTypes.SizeOfInt(r.Length)
				for _, pid := range strings.Split(r.DataStructure.PointIdFrom, ".") {
					// p := reflection.GetStringFromStruct(r.Interface, pid, valueTypes.IgnoreLength, r.DataStructure.PointNameDateFormat)	// Look forward into structure.
					p := reflection.GetStringFromStruct(r.Interface, pid, valueTypes.IgnoreLength, r.DataStructure.PointNameDateFormat)	// Look forward into structure.
					if p == "" {
						p = reflection.GetStringFromStruct(r.CurrentReflect.Interface, pid, valueTypes.IgnoreLength, r.DataStructure.PointNameDateFormat)
					}
					if p == "" {
						continue
					}
					pns = append(pns, p)
				}

				if r.DataStructure.PointId != "" {
					if !r.DataStructure.PointIdReplace {
						pns = append(pns, r.DataStructure.PointId)
					}
				}
				pn = strings.Join(pns, ".")
				r.DataStructure.Endpoint.Append(r.DataStructure.PointId)
				// fmt.Printf("[PointIdFrom2     ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				if r.DataStructure.PointIdReplace {
					r.DataStructure.Endpoint.PopLast()
					// fmt.Printf("[PointIdFrom3     ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				}
				if pn != "" {
					r.DataStructure.Endpoint.Append(pn)
					// fmt.Printf("[PointIdFrom4     ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				}

			// case r.DataStructure.PointIdFromParent != "":
			// 	// PointIdFromParent - In this case points to a field within a PARENT struct.
			// 	// We also already have the parent.
			// 	// @TODO - Need to fix this up.
			// 	for _, child := range r.ParentReflect.ChildReflect {
			// 		if child.FieldName == r.DataStructure.PointIdFromChild {
			// 			dateFormat := child.DataStructure.PointNameDateFormat
			// 			if dateFormat == "" {
			// 				dateFormat = valueTypes.DateTimeAltLayout
			// 			}
			// 			intSize := valueTypes.SizeOfInt(r.CurrentReflect.Length)
			// 			pn = valueTypes.AnyToValueString(child.InterfaceValue, intSize, dateFormat)
			// 			break
			// 		}
			// 	}

			default:
				// pn = r.DataStructure.PointId
				// fmt.Printf("[default          ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				switch r.CurrentReflect.Kind {
					case reflect.Array:
						fallthrough
					case reflect.Slice:
						// var pns []string
						// for _, pid := range strings.Split(ds.PointIdFrom, ".") {
						// 	pns = append(pns, reflection.GetStringFromStruct(r.CurrentReflect.Interface, pid))
						// }
						// pn = strings.Join(pns, ".")
						// @TODO - A real hack.
						if (r.CurrentReflect.DataStructure.PointIdFrom == "") && (r.CurrentReflect.DataStructure.PointIdFromChild == "") {
							ft := valueTypes.GetIntFormatForPrintf(r.CurrentReflect.Length)
							pn = fmt.Sprintf(ft, r.Index)
							if r.IsPointIdReplace() {
								pn = r.DataStructure.PointId + "_" + pn
							}
							if pn == "" {
								pn = r.DataStructure.PointId
							}
							// fmt.Printf("[default3a        ]	pn: %s - EPP: %s\n", pn, r.DataStructure.Endpoint)
						}

					default:
						pn = r.DataStructure.PointId
						// fmt.Printf("[default3b        ]	pn: %s - EPP: %s\n", pn, r.DataStructure.Endpoint)
				}
				// r.DataStructure.Endpoint.Append(r.DataStructure.PointId)
				// if r.DataStructure.PointIdReplace {
				// 	r.DataStructure.Endpoint.PopLast()
				// }
				if pn != "" {
					r.DataStructure.Endpoint.Append(pn)
					// fmt.Printf("[default4         ]	EPP: %s	- FP: %s\n", r.DataStructure.Endpoint, r.FieldPath)
				}
		}
	}

	// fmt.Printf("EPP(AFTER): %s\n", r.DataStructure.Endpoint.String())
	return r.DataStructure.Endpoint
}


func (r *Reflect) IsTable() bool {
	return r.DataStructure.DataTable
}

func (r *Reflect) IsNotTable() bool {
	return !r.DataStructure.DataTable
}

func (r *Reflect) IsTableChild() (bool, int) {
	var yes bool
	var safeIterate int
	for range Only.Once {
		yes = r.isTableChild(&safeIterate)
	}
	return yes, safeIterate
}

func (r *Reflect) isTableChild(safeIterate *int) bool {
	var yes bool
	for range Only.Once {
		*safeIterate++
		if *safeIterate > 42 {
			// Arbitrary limit of 42.
			break
		}

		// Don't include this!
		// if r.DataStructure.DataTableChild {
		// 	yes = true
		// 	break
		// }

		if r.DataStructure.DataTable {
			yes = true
			break
		}

		if r.CurrentReflect == nil {
			break
		}

		if r.IsStart {
			break
		}

		yes = r.CurrentReflect.isTableChild(safeIterate)
	}
	return yes
}

func (r *Reflect) IsPointArrayFlatten() bool {
	return r.DataStructure.PointArrayFlatten
}

func (r *Reflect) IsPointListFlatten() bool {
	return r.DataStructure.PointListFlatten
}

func (r *Reflect) IsPointVariableUnit() bool {
	return r.DataStructure.PointVariableUnit
}

func (r *Reflect) IsPointIdReplace() bool {
	return r.DataStructure.PointIdReplace
}

func (r *Reflect) IsNotPointIdReplace() bool {
	return !r.DataStructure.PointIdReplace
}

func (r *Reflect) IsPointIgnore() bool {
	return r.DataStructure.PointIgnore
}

func (r *Reflect) IsPointIgnoreZero() bool {
	return r.DataStructure.PointIgnoreZero
}

func (r *Reflect) IsPointVirtual() bool {
	return r.DataStructure.PointVirtual
}

func (r *Reflect) IsDataTableMerge() bool {
	return r.DataStructure.DataTableMerge
}

func (r *Reflect) IsDataTableIndex() bool {
	return r.DataStructure.DataTableIndex
}

func (r *Reflect) GetDataTableIndexNames() []string {
	return r.DataStructure.DataTableIndexNames
}

func (r *Reflect) SetDataTableIndexNames(args ...string) {
	r.DataStructure.DataTableIndexNames = args
}

func (r *Reflect) SetDataTableIndexTitle(args string) {
	r.DataStructure.DataTableIndexTitle = args
}

func (r *Reflect) CountChildren() (int, int) {
	var rows int
	var cols int
	for range Only.Once {
		rows = len(r.ChildReflect)
		for row := range r.ChildReflect {
			col := len(r.ChildReflect[row].ChildReflect)
			if cols < col {
				cols = col
			}
		}
	}
	return rows, cols
}

func (r *Reflect) IsPointTimestampZero() bool {
	return r.DataStructure.PointTimestamp.IsZero()
}
func (r *Reflect) IsPointTimestampNotZero() bool {
	return !r.DataStructure.PointTimestamp.IsZero()
}

func (r *Reflect) AsJson() string {
	var ret string
	for range Only.Once {
		j, err := json.Marshal(r.InterfaceValue)
		if err != nil {
			break
		}
		ret = string(j)
	}
	return ret
}

// SetGoStructOptions - Copies and updates DataTags from a GoStruct object.
func (r *Reflect) SetGoStructOptions(limit int) bool {
	var yes bool

	// fmt.Printf("Current(enter): %s\n", r.DataStructure)
	for range Only.Once {
		r.GoStructs.Parent, r.GoStructs.Current = GetChildGoStruct(r.Interface, limit)
		if r.GoStructs.Current == nil {
			r.GoStructs.Current = r.GetGoStruct()
		}

		if r.GoStructs.Current != nil {
			yes = true
			r.GoStructs.Current.UpdateTags(r.CurrentReflect, r)

			err := r.DataStructure.SetFrom(r.GoStructs.Current)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}
		}

		if r.GoStructs.Parent != nil {
			if r.CurrentReflect.IsGroup() {
				yes = true
				r.GoStructs.Parent.UpdateTags(r.CurrentReflect.CurrentReflect, r.CurrentReflect)

				err := r.DataStructure.SetFrom(r.GoStructs.Parent)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}

				if r.GoStructs.Parent.DataTable {
					r.DataStructure.DataTable = false
					r.DataStructure.StrBool.DataTable = "false"

					r.DataStructure.DataTableChild = true
					r.DataStructure.StrBool.DataTableChild = "true"

					r.DataStructure.DataTableMerge = false
					r.DataStructure.StrBool.DataTableMerge = "false"

					r.DataStructure.DataTablePivot = false
					r.DataStructure.StrBool.DataTablePivot = "false"

					r.DataStructure.DataTableIndex = false
					r.DataStructure.StrBool.DataTableIndex = "false"

					r.DataStructure.DataTableIndexNames = []string{}
					r.DataStructure.StrBool.DataTableIndexNames = ""

					r.DataStructure.DataTableId = ""
					r.DataStructure.DataTableName = ""
					r.DataStructure.DataTableTitle = ""
					r.DataStructure.DataTableSortOn = ""
					r.DataStructure.DataTableIndexTitle = ""
				}
				break
			}

			if r.IsGroup() {
				// fmt.Printf("IsGroup(): %t / %t\n", r.CurrentReflect.IsGroup(), r.IsGroup())
				yes = true
				r.GoStructs.Parent.UpdateTags(r.CurrentReflect, r)

				err := r.DataStructure.SetFrom(r.GoStructs.Parent)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}

			if r.FieldName == "ResultData" {
				yes = true
				r.GoStructs.Parent.UpdateTags(r.CurrentReflect, r)

				// fmt.Printf("SetFrom1: %s\n", r.DataStructure)
				err := r.DataStructure.SetFrom(r.GoStructs.Parent)
				// fmt.Printf("SetFrom2: %s\n", r.DataStructure)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				}
			}
		}
	}
	// fmt.Printf("Current(exit): %s\n", r.DataStructure)
	// fmt.Printf("GoStructs.Parent(exit): %s\n", r.GoStructs.Parent)
	// fmt.Printf("GoStructs.Current(exit): %s\n", r.GoStructs.Current)

	return yes
}

func (r *Reflect) IsGoStruct() bool {
	var yes bool
	for range Only.Once {
		if r.FieldName == NameGoStruct {
			yes = true
			r.DataStructure.PointIgnore = true
			break
		}
		if r.FieldName == NameGoStructParent {
			yes = true
			r.DataStructure.PointIgnore = true
			break
		}
		yes = false
	}
	return yes
}

func (r *Reflect) IsGroup() bool {
	var yes bool

	switch {
		case r == nil:
			break

		case r.Kind == reflect.Map:
			yes = true

		case r.Kind == reflect.Array:
			yes = true

		case r.Kind == reflect.Slice:
			yes = true
	}

	return yes
}

func GetChildGoStruct(ref interface{}, limit int) (*DataTags, *DataTags) {
	var parent *DataTags
	var current *DataTags

	for range Only.Once {
		// Only scan the immediate children.
		if limit < 0 {
			break
		}
		limit--

		vo := reflect.ValueOf(ref)

		if vo.Kind() == reflect.Pointer {
			if valueTypes.IsNil(ref) {
				break
			}
			ref2 := vo.Elem().Interface()
			if valueTypes.IsNil(ref2) {
				break
			}
			limit++
			parent, current = GetChildGoStruct(ref2, limit)
			break
		}

		if vo.Kind() == reflect.Struct {
			// Iterate over all available fields, looking for the parent GoStruct.
			for i := 0; i < vo.NumField(); i++ {
				fieldTo := vo.Type().Field(i)
				fieldVo := vo.Field(i)

				if fieldTo.Name != NameGoStructParent {
					// Field name has to be GoStructParent
					continue
				}
				if fieldTo.Type.String() != PkgGoStructParent {
					// Package type has to be GoStruct.GoStructParent
					continue
				}
				if !fieldTo.IsExported() {
					continue
				}

				var dt DataTags
				dt.GetTags(fieldTo, fieldVo)	// &Reflect{Interface: ref}, &Reflect{Interface: fieldVo.Interface()}, fieldTo, fieldVo
				parent = &dt
				break
			}

			// Iterate over all available fields, looking for the current GoStruct.
			for i := 0; i < vo.NumField(); i++ {
				fieldTo := vo.Type().Field(i)
				fieldVo := vo.Field(i)

				if fieldTo.Name != NameGoStruct {
					// Field name has to be GoStruct
					continue
				}
				if fieldTo.Type.String() != PkgGoStruct {
					// Package type has to be GoStruct.GoStruct
					continue
				}
				if !fieldTo.IsExported() {
					continue
				}

				var dt DataTags
				dt.GetTags(fieldTo, fieldVo)	// &Reflect{Interface: ref}, &Reflect{Interface: fieldVo.Interface()}, fieldTo, fieldVo
				current = &dt
				break
			}
		}

		if (vo.Kind() == reflect.Slice) || (vo.Kind() == reflect.Array) {
			if limit < 0 {
				break // Count a recurse as one extra.
			}
			limit--

			// Iterate over all available fields, looking for the field name.
			for i := 0; i < vo.Len(); i++ {
				parent, current = GetChildGoStruct(vo.Index(i).Interface(), limit)
				if current != nil {
					break
				}
				if parent != nil {
					break
				}
			}
			break
		}
	}

	return parent, current
}

func (r *Reflect) HasGoStruct() bool {
	var yes bool

	for range Only.Once {
		if r.CurrentReflect.CurrentReflect.GoStructs.Current != nil {
			yes = true
			break
		}
		if r.CurrentReflect.GoStructs.Current != nil {
			yes = true
			break
		}
	}

	return yes
}

func (r *Reflect) GetGoStruct() *DataTags {
	var tags *DataTags

	for range Only.Once {
		// if r.CurrentReflect.GoStruct != nil {
		// 	fmt.Printf("GOSTRUCT r.CurrentReflect.GoStruct[%s] - %s\n", r.CurrentReflect.FieldPath, r.CurrentReflect.GoStruct)
		// 	tags = r.CurrentReflect.GoStruct
		// 	break
		// }
		if r.CurrentReflect.CurrentReflect.GoStructs.Current != nil {
			// fmt.Printf("GOSTRUCT r.CurrentReflect.CurrentReflect.GoStruct[%s] - %s\n", r.CurrentReflect.CurrentReflect.FieldPath, r.CurrentReflect.GoStruct)
			tags = r.CurrentReflect.CurrentReflect.GoStructs.Current
			break
		}
	}

	return tags
}

func (r *Reflect) GetGoStructCurrent() *Reflect {
	var current *Reflect

	for range Only.Once {
		// if r.CurrentReflect.GoStruct != nil {
		// 	fmt.Printf("GOSTRUCT r.CurrentReflect.GoStruct[%s] - %s\n", r.CurrentReflect.FieldPath, r.CurrentReflect.GoStruct)
		// 	current = r.CurrentReflect
		// 	break
		// }
		if r.CurrentReflect.CurrentReflect.GoStructs.Current != nil {
			// fmt.Printf("GOSTRUCT r.CurrentReflect.CurrentReflect.GoStruct[%s] - %s\n", r.CurrentReflect.CurrentReflect.FieldPath, r.CurrentReflect.GoStruct)
			current = r.CurrentReflect.CurrentReflect
			break
		}
	}

	return current
}

func (r *Reflect) IsGoStructForParent() bool {
	var yes bool

	for range Only.Once {
		voSrc := reflect.ValueOf(*r)
		t := voSrc.Type()
		s := t.Kind().String()
		if s == NameGoStructParent {
			yes = true
		}
		if s == PkgGoStructParent {
			yes = true
		}
	}

	return yes
}

func (r *Reflect) PointIgnoreIfChildFromNil() bool {
	var yes bool
	for range Only.Once {
		if r.ParentReflect.DataStructure.PointIgnoreIfChildFromNil == "" {
			// If parent doesn't care about a zero value.
			yes = false
			break
		}

		ret := reflection.GetStringFrom(r.Interface, r.Index, r.ParentReflect.DataStructure.PointIgnoreIfChildFromNil, valueTypes.IgnoreLength, r.DataStructure.PointNameDateFormat)
		if ret == "" {
			yes = false
			break
		}

		yes = r.IsNil
	}
	return yes
}

func (r *Reflect) EndPointPath() *EndPointPath {
	return &r.DataStructure.Endpoint
}

func (r *Reflect) GetFieldPath() EndPointPath {
	return r.FieldPath
}

func (r *Reflect) CopyEndPointPath() EndPointPath {
	return r.DataStructure.Endpoint.Copy()
}

func (r *Reflect) SetEndPointPath(epp EndPointPath) *EndPointPath {
	r.DataStructure.Endpoint = epp
	return &r.DataStructure.Endpoint
}

func (r *Reflect) PointId() string {
	return r.DataStructure.PointId
}

func (r *Reflect) PointName() string {
	if r.DataStructure.PointName == "" {
		return valueTypes.PointToName(r.DataStructure.PointId)
	}
	return r.DataStructure.PointName
}

func (r *Reflect) PointGroupName() string {
	return r.DataStructure.PointGroupName
}

func (r *Reflect) PointUpdateFreq() string {
	return r.DataStructure.PointUpdateFreq
}

func (r *Reflect) ValueUnit() string {
	return r.Value.GetUnit()
}

func (r *Reflect) ValueType() string {
	return r.Value.Type()
}

func (r *Reflect) ValueFirst() *valueTypes.UnitValue {
	return r.Value.First()
}

func (r *Reflect) ValueLast() *valueTypes.UnitValue {
	return r.Value.Last()
}

func (r *Reflect) ValueLength() int {
	return r.Value.Length()
}


func GetStructFields(ref interface{}) map[string]string {
	ret := make(map[string]string)

	for range Only.Once {
		var Ref Reflect
		Ref.Init(ref, ref, EndPointPath{})

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(&Ref, &Ref, i, reflect.Value{})

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

func GetStructFieldsAsArray(ref *Reflect) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.Init(ref, ref, EndPointPath{})

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(&Ref, &Ref, i, reflect.Value{})

				if !Child.IsExported {
					continue
				}

				ret = append(ret, Child.FieldName)
			}
		}
	}

	return ret
}

func GetStructValuesAsArray(ref *Reflect) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.Init(ref, ref, EndPointPath{})

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(&Ref, &Ref, i, reflect.Value{})

				if !Child.IsExported {
					continue
				}

				ret = append(ret, fmt.Sprintf("%v", Child.Interface))
			}
		}
	}

	return ret
}

func AddFloatValues(refs ...*Reflect) float64 {
	var ret float64
	for _, ref := range refs {
		for _, value := range ref.Value.Range(true) {
			ret += value.ValueFloat()
		}
	}
	return ret
}

func AddIntegerValues(refs ...*Reflect) int64 {
	var ret int64
	for _, ref := range refs {
		for _, value := range ref.Value.Range(true) {
			ret += value.ValueInt()
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
