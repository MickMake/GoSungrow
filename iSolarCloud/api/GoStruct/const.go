package GoStruct

// These are tags that can be added to a Go structure that GoStruct uses to process the structure.


const (
	// NameGoStruct - Name of field within structure that allows for assigning tags to the parent.
	// Add like this:
	// type ResultData []struct {
	//     GoStruct      GoStruct.GoStruct  `json:"-" DataTable:"true" DataTableSortOn:"UnitConvertId"`
	//
	// ...
	//
	// }
	NameGoStruct = "GoStruct"
	PkgGoStruct = "GoStruct.GoStruct"
	NameGoStructParent = "GoStructParent"
	PkgGoStructParent = "GoStruct.GoStructParent"

	// PointId - Point id in the form p\d+ or \d+ or free-form text.
	PointId = "PointId"
	// PointIdFrom -  Searches current structure for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointIdFrom = "PointIdFrom"
	// PointIdFromChild -  Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointIdFromChild = "PointIdFromChild"
	// PointIdReplace -  Replace PointNameFrom instead of append.
	PointIdReplace = "PointIdReplace"
	// PointIdFromParent -  Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	// PointIdFromParent = "PointIdFromParent"	- Not supported.


	// PointParentId -  Associated parent of point.
	PointParentId = "PointParentId"

	// PointUpdateFreq -  Point update frequency - Total, Yearly, Monthly, Day.
	PointUpdateFreq = "PointUpdateFreq"
	UpdateFreqInstant = "instant"
	UpdateFreq5Mins   = "5mins"
	UpdateFreqBoot    = "boot"
	UpdateFreqDay     = "daily"
	UpdateFreqMonth   = "monthly"
	UpdateFreqYear    = "yearly"
	UpdateFreqTotal   = "total"


	// PointValueType -  Value type of point: energy, date, battery, temperature.
	PointValueType = "PointValueType"
	// PointUnit -  Units: Wh, kWh, C, h.
	PointUnit = "PointUnit"
	// PointUnitFrom -  Get PointUnit from another field structure.
	PointUnitFrom = "PointUnitFrom"
	// PointUnitFromParent -  Get PointUnit from another parent field structure.
	PointUnitFromParent = "PointUnitFromParent"


	// PointIgnore -  Ignore this point.
	PointIgnore = "PointIgnore"
	// PointIgnoreIfNil -  Ignore this point if a child is nil or empty.
	PointIgnoreIfNil = "PointIgnoreIfNil"
	// PointIgnoreChildIfFromNil -  Ignore this point if a child is nil or empty.
	PointIgnoreChildIfFromNil = "PointIgnoreChildIfFromNil"
	// PointIgnoreZero -  Ignore arrays with zero size, (default true).
	PointIgnoreZero = "PointIgnoreZero"


	// PointAliasTo -  Alias this point to another point.
	PointAliasTo = "PointAliasTo"
	// PointAliasFrom -  Alias this point from another point.
	PointAliasFrom = "PointAliasFrom"


	// PointGroupName -  Point group name.
	PointGroupName = "PointGroupName"
	// PointGroupNameFrom -  Get PointGroupName from another field structure.
	PointGroupNameFrom = "PointGroupNameFrom"


	// PointName -  Human-readable name of point.
	PointName = "PointName"


	// PointNameDateFormat -  Date format when using PointNameFrom, (if the field is a time.Time type).
	PointNameDateFormat = "PointNameDateFormat"

	// PointArrayFlatten -  Flatten an array into a string. EG: ["one", "two", "three"]
	PointArrayFlatten = "PointArrayFlatten"

	// PointSplitOn -  Split a point into an array separating by defined string.
	PointSplitOn = "PointSplitOn"

	// PointSplitOnType -  What valueTypes will be used for a split.
	PointSplitOnType = "PointSplitOnType"

	// PointTimestampFrom -  Pull timestamp from another field structure.
	PointTimestampFrom = "PointTimestampFrom"

	// PointValueReplace - Search endpoint value and replace with string contained in PointValueReplaceWith
	PointValueReplace = "PointValueReplace"
	// PointValueReplaceWith -
	PointValueReplaceWith = "PointValueReplaceWith"


	// PointDevice - Define a PointDevice for this point.
	PointDevice = "PointDevice"
	// PointDeviceFrom - Reference PointDevice from another sibling.
	PointDeviceFrom = "PointDeviceFrom"
	// PointDeviceFromParent - Reference PointDevice from a parent.
	PointDeviceFromParent = "PointDeviceFromParent"


	// IsDataTable -  This entity is a data table - Will only traverse down one child.
	IsDataTable = "DataTable"
	// DataTableChild - Will be included in table listings.
	DataTableChild = "DataTableChild"
	// DataTableId -  Table id, (defaults to Json tag).
	DataTableId = "DataTableId"
	// DataTableName -  Table Name, (defaults to DataTableId).
	DataTableName = "DataTableName"
	// DataTableTitle -  Table Title, (defaults to DataTableId in name format).
	DataTableTitle = "DataTableTitle"
	// DataTableMerge -  Merge rows together - useful for when we use, for EG: []valueTypes.Float
	DataTableMerge = "DataTableMerge"
	// DataTableShowIndex -  Show index on table.
	DataTableShowIndex = "DataTableShowIndex"
	// DataTableSortOn -  Sort table using this Field.
	DataTableSortOn = "DataTableSortOn"
)

type GoStruct bool
type GoStructParent bool
