package api


// From struct_data.go

// func (dm *DataMap) GetEntry(entry string, index int) *DataEntry {
// 	var ret *DataEntry
// 	for range Only.Once {
// 		pe := dm.Map[entry]
// 		if pe.Entries != nil {
// 			ret = pe.GetEntry(index)
// 			break
// 		}
//
// 		for k, v := range dm.Map {
// 			if strings.HasSuffix(k, "." + entry) {
// 				ret = v.GetEntry(index)
// 				break
// 			}
// 		}
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetFloatValue(entry string, index int) float64 {
// 	var ret float64
// 	for range Only.Once {
// 		pe := dm.GetEntry(entry, index)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: GetFloatValue('%s', '%d')\n", entry, index)
// 			break
// 		}
// 		ret = pe.Value.ValueFloat()
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetValue(entry string, index int) float64 {
// 	var ret float64
// 	for range Only.Once {
// 		v := dm.GetEntry(entry, index)
// 		if v.IsNotValid() {
// 			fmt.Printf("ERROR: GetValue('%s', %d)\n", entry, index)
// 			break
// 		}
//
// 		ret = v.Value.ValueFloat()
// 	}
// 	return ret
// }
//
// func (dm *DataMap) GetEntryFromPointId(pointId string) *DataEntries {
// 	var ret *DataEntries
// 	for range Only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				ret = dm.Map[i]
// 				break
// 			}
// 		}
// 	}
// 	return ret
// }
//
// func (dm *DataMap) SetEntryUnits(pointId string, unit string) {
// 	for range Only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetUnits(unit)
// 				dm.Map[i].SetUnits(unit)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) SetEntryGroupName(pointId string, groupName string) {
// 	for range Only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetGroupName(groupName)
// 				dm.Map[i].SetGroupName(groupName)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) SetEntryTimestamp(pointId string, timeStamp valueTypes.DateTime) {
// 	for range Only.Once {
// 		for i, v := range dm.Map {
// 			if v.MatchPointId(pointId) {
// 				// e := dm.Map[i]
// 				// dm.Map[i] = e.SetTimestamp(timeStamp)
// 				dm.Map[i].SetTimestamp(timeStamp)
// 				break
// 			}
// 		}
// 	}
// }
//
// func (dm *DataMap) FromRefAddAlias(ref string, parentId string, pid string, name string) {
// 	for range Only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddAlias('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		dm.Add(de)
// 	}
// }
//
// func (dm *DataMap) FromRefAddState(ref string, parentId string, pid string, name string) {
// 	for range Only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddState('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		de.MakeState(pe.Value.ValueBool())
// 		// de := pe.CreateState(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name)
// 		dm.Add(de)
// 	}
// }
//
// func (dm *DataMap) FromRefAddFloat(ref string, parentId string, pid string, name string, value float64) {
// 	for range Only.Once {
// 		pe := dm.GetEntry(ref, 0)
// 		if pe.IsNotValid() {
// 			fmt.Printf("ERROR: FromRefAddFloat('%s', '%s', '%s', '%s')\n", ref, parentId, pid, name)
// 			break
// 		}
//
// 		de := CopyDataEntry(*pe, pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, pe.Point.GroupName, pe.Point.Unit, pe.Point.ValueType)
// 		de.MakeFloat(value, "", "")
// 		// de := pe.CreateFloat(pe.EndPoint, parentId, valueTypes.SetPointIdString(pid), name, value)
// 		dm.Add(de)
// 	}
// }
//
// func CopyDataEntry(ref DataEntry, endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string) DataEntry {
// 	var ret DataEntry
// 	for range Only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		point := CopyPoint(*ref.Point, parentId, pid, name, groupName, unit, Type)
// 		// point = &Point {
// 		// 	Parents:   de.Point.Parents,
// 		// 	Id:        pid,
// 		// 	GroupName: "alias",
// 		// 	Name:      name,
// 		// 	Unit:      de.Point.Unit,
// 		// 	UpdateFreq:  de.Point.UpdateFreq,
// 		// 	ValueType: de.Point.ValueType,
// 		// 	Valid:     true,
// 		// 	States:    de.Point.States,
// 		// }
// 		// var parent ParentDevice
// 		// parent.Set(parentId)
// 		// point.Parents.Add(parent)
// 		// point.Unit = "binary"
// 		// if point.Unit == "" {
// 		// 	point.Unit = ref.Unit()
// 		// }
// 		// point.Name = name
// 		// if point.Name == "" {
// 		// 	point.Name = pid.PointToName()
// 		// }
// 		// // if de2.Point.GroupName == "" {
// 		// // 	de2.Point.GroupName = groupName
// 		// // }
// 		// point.FixUnitType()
// 		// point.Valid = true
//
// 		ret.Point = point
// 		ret.EndPoint = endpoint
// 		ret.Parent.Set(parentId)
// 		ret.Valid = true
// 		ret.Hide = false
// 	}
//
// 	return ret
// }
//
// func CopyPoint(ref Point, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string) *Point {
// 	for range Only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		ref.Parents.Add(parent)
// 		ref.Id = pid
// 		ref.Unit = unit
// 		ref.Name = name
// 		ref.UpdateFreq = ""
// 		ref.GroupName = groupName
// 		ref.ValueType = Type
// 		ref.Valid = true
// 		ref.States = nil
//
// 		ref.FixUnitType()
// 	}
//
// 	return &ref
// }
//
// func (dm *DataMap) HideEntry(pointId valueTypes.PointId) {
// 	for range Only.Once {
// 		de := dm.GetEntryFromPointId(pointId)
// 		de.Hide()
// 	}
// }
//
// func (dm *DataMap) AddEntry(endpoint string, parentId string, point Point, date valueTypes.DateTime, value string) {
// 	for range Only.Once {
// 		unit := point.Unit	// Save unit.
// 		vType := point.ValueType	// Save type.
//
// 		// Match to a previously defined point.
// 		p := GetPoint(point.Id.String())
// 		if p != nil {
// 			// No point found. Create one.
// 			p = CreatePoint(parentId, pid, name, groupName, unit, Type)
// 		}
// 		point = *p
//
// 		// var parents ParentDevices
// 		// parents.Add(ParentDevice{Key: device})
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
//
// 		if point.Name == "" {
// 			point.Name = point.Id.PointToName()
// 		}
// 		// fid := JoinDevicePoint(parent.Key, point.Id)
// 		ref := valueTypes.SetUnitValueString(value, unit, vType)
// 		point.Unit = ref.Unit()
// 		point.Valid = true
//
// 		if _, ok := dm.DataPoints[point.Id.String()]; ok {
// 			fmt.Printf("BARF: %s\n", point.Id)
// 		}
//
// 		// dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
// 		dm.Add(DataEntry {
// 			EndPoint:   endpoint,
// 			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Point:      &point,
// 			Date:       date,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 			ValueBool:  ref.ValueBool(),
// 			Index:      0,
// 			Valid:      true,
// 			Hide:       false,
// 		})
// 	}
// }
//
// func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, ref valueTypes.UnitValue) {
// 	for range Only.Once {
// 		if endpoint == "" {
// 			endpoint = GoStruct.GetCallerPackage(2)
// 		}
//
// 		ref = ref.UnitValueFix()
//
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		point := GetPoint(pid.String())
// 		if point == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentId, pid, name, groupName, ref.Unit(), ref.Type())
// 			// de := CreateDataEntry(endpoint, parentId, pid, name, groupName, date, ref)
// 			// dm.Add(de)
// 			// break
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
// 		if point.Unit == "" {
// 			point.Unit = ref.Unit()
// 		}
// 		if point.Name == "" {
// 			point.Name = name
// 		}
// 		if point.Name == "" {
// 			point.Name = pid.PointToName()
// 		}
// 		if point.GroupName == "" {
// 			point.GroupName = groupName
// 		}
// 		point.FixUnitType()
// 		point.Valid = true
//
// 		dm.Add(DataEntry {
// 			EndPoint:   endpoint,
// 			// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Point:      point,
// 			Date:       date,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 			ValueBool:  ref.ValueBool(),
// 			Index:      0,
// 			Valid:      true,
// 			Hide:       false,
// 		})
// 	}
// }
//
// func (dm *DataMap) AddFloat(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value float64) {
// 	for range Only.Once {
// 		// fvs := Float64ToString(value)
// 		point := GetPoint(parentId, pid)
// 		if point == nil {
// 			// No UV found. Create one.
// 			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
// 			break
// 		}
//
// 		ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
// 		if ref.Unit() != point.Unit {
// 			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
// 			point.Unit = ref.Unit()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		point.Parents.Add(parent)
//
// 		dm.Add(pid, DataEntry {
// 			EndPoint:   endpoint,
// 			FullId:     JoinDevicePoint(endpoint, point.Id),
// 			// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 			Parent:     parent,
//
// 			Date:       date,
// 			Point:      point,
// 			Value:      ref.String(),
// 			ValueFloat: ref.Value(),
// 		})
// 	}
//
// 	uv := valueTypes.SetUnitValueFloat(value, "", "float")
// 	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 	// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 	// 	Unit:       "float",
// 	// 	Value:      fmt.Sprintf("%f", value),
// 	// 	ValueFloat: 0,
// 	// })
// 	dm.Add(pid, de)
// }
//
// func (dm *DataMap) AddString(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value string) {
// 	dm.Add(pid, CreateDataEntryString(date, endpoint, parentId, pid, name, value))
// }
//
// func (dm *DataMap) AddInt(endpoint string, parentId string, pid PointId, name string, date valueTypes.DateTime, value int64) {
//
// 	for range Only.Once {
// 		uvs, ok := valueTypes.AnyToUnitValue(value, "", "")
// 		if !ok {
// 			fmt.Printf("ERROR: AddInt(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value %d)",
// 				endpoint, parentId, pid, name, date, value)
// 			break
// 		}
// 		for _, uv := range uvs {
// 			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 			dm.Add(pid, de)
// 		}
//
// 		// uv := valueTypes.SetUnitValueInteger(value, "", "int")
// 		// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 		// // de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
// 		// // 	Unit:       "int",
// 		// // 	Value:      fmt.Sprintf("%d", value),
// 		// // 	ValueFloat: float64(value),
// 		// // })
// 		// dm.Add(pid, de)
// 	}
// }
//
// func (dm *DataMap) AddAny(endpoint string, parentId string, pid valueTypes.PointId, name string, date valueTypes.DateTime, value interface{}) {
//
// 	for range Only.Once {
// 		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, "", "")
// 		if !ok {
// 			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
// 				endpoint, parentId, pid, name, date, value)
// 			break
// 		}
//
// 		point := GetPoint(parentId + "." + pid.String())
// 		if point == nil {
// 			// No UV found. Create one.
// 			for _, uv := range uvs {
// 				de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 				if isNil {
// 					de.Point.ValueType += "(NIL)"
// 				}
// 				dm.Add(de)
// 			}
// 			// dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name,
// 			// 	valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)))
// 			break
// 		}
//
// 		// ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
// 		// if ref.Unit() != point.Unit {
// 		// 	fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
// 		// 	point.Unit = ref.Unit()
// 		// }
//
// 		if isNil {
// 			point.ValueType += "(NIL)"
// 		}
//
// 		for _, uv := range uvs {
// 			if uv.Unit() != point.Unit {
// 				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.Unit())
// 				point.Unit = uv.Unit()
// 			}
//
// 			var parent ParentDevice
// 			parent.Set(parentId)
// 			point.Parents.Add(parent)
//
// 			// CreateDataEntry
// 			de := DataEntry {
// 				EndPoint: endpoint,
// 				// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 				Parent:   parent,
//
// 				Date:       date,
// 				Point:      point,
// 				Value:      uv.String(),
// 				ValueFloat: uv.Value(),
// 				ValueBool:  uv.ValueBool(),
// 				Index:      0,
// 				Valid:      true,
// 				Hide:       false,
// 			}
// 			dm.Add(de)
// 		}
//
// 		for _, uv := range uvs {
// 			de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
// 			dm.Add(de)
// 		}
// 	}
// }
//
// func (de *DataEntry) CreateFloat(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string, value float64) DataEntry {
// 	var ret DataEntry
// 	for range Only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		ret = de.CreateDataEntry(endpoint, parentId, pid, name, groupName, unit, Type)
// 		uv := valueTypes.SetUnitValueFloat(value, ret.Point.Unit, ret.Point.ValueType)
// 		ret.Value = uv.String()
// 		ret.ValueFloat = uv.Value()
// 		ret.Valid = true
// 		ret.Hide = false
// 	}
// 	return ret
// }
//
// func (de *DataEntry) CreateState(endpoint string, parentId string, pid valueTypes.PointId, name string) DataEntry {
// 	var ret DataEntry
// 	for range Only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
//
// 		de2 := de.CreateDataEntry(endpoint, parentId, pid, name)
// 		if de2.ValueFloat == 0 {
// 			de2.Value = "false"
// 			de2.ValueBool = false
// 			de2.ValueFloat = 0
// 		} else {
// 			de2.Value = "true"
// 			de2.ValueBool = true
// 			de2.ValueFloat = 1
// 		}
// 		de2.Valid = true
// 		de2.Hide = false
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
// 		de2.Point.Parents.Add(parent)
// 		de2.Point.Unit = "binary"
// 		if de2.Point.Unit == "" {
// 			de2.Point.Unit = ref.Unit()
// 		}
// 		de2.Point.Name = name
// 		if de2.Point.Name == "" {
// 			de2.Point.Name = pid.PointToName()
// 		}
// 		// if de2.Point.GroupName == "" {
// 		// 	de2.Point.GroupName = groupName
// 		// }
// 		de2.Point.FixUnitType()
// 		de2.Point.Valid = true
// 	}
//
// 	return ret
// }
//
// func CreateDataEntryActive(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value float64) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "state")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, point.Id),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      fmt.Sprintf("%v", IsActive(value)),
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }
//
// func CreateDataEntryString(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value string) DataEntry {
// 	point := GetPoint(parentId, pid)
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, "string")
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, pid),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      value,
// 		ValueFloat: 0,
// 		Index:      0,
// 	}
// }
//
// func CreateDataEntryUnitValue(date valueTypes.DateTime, endpoint string, parentId string, pid valueTypes.PointId, name string, value valueTypes.UnitValue) DataEntry {
// 	value = value.UnitValueFix()
//
// 	point := GetPoint(parentId + "." + pid.String())
// 	if point == nil {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point = CreatePoint(parentId, pid, name, value.Unit())
// 	}
//
// 	var parent ParentDevice
// 	parent.Set(parentId)
// 	point.Parents.Add(parent)
// 	point.Valid = true
//
// 	return DataEntry {
// 		EndPoint:   endpoint,
// 		// FullId:     valueTypes.JoinDataPoint(endpoint, pid.String()),
// 		// FullId:     JoinDevicePoint(parent.Key, pid),
// 		Parent:     parent,
//
// 		Point:      point,
// 		Date:       date,
// 		Value:      value.String(),
// 		ValueFloat: value.Value(),
// 		ValueBool:  value.ValueBool(),
// 		Index:      0,
// 		Valid:      true,
// 		Hide:       false,
// 	}
// }
//
// func CreatePoint(parentId string, pid valueTypes.PointId, name string, unit string) *Point {
// 	if name == "" {
// 		name = pid.PointToName()
// 	}
//
// 	var parents ParentDevices
// 	parents.Add(ParentDevice{Key: parentId})
//
// 	ret := &Point {
// 		Parents:   parents,
// 		Id:        pid,
// 		GroupName: parentId,
// 		Name:      name,
// 		Unit:      unit,
// 		UpdateFreq:  "",
// 		ValueType: "",
// 		Valid:     true,
// 		States:    nil,
// 	}
// 	ret.FixUnitType()
//
// 	return ret
// }
//
// func IsActive(value float64) bool {
// 	if (value > 0.01) || (value < -0.01) {
// 		return true
// 	}
// 	return false
// }
//
// func JoinDevicePoint(endpoint string, pid valueTypes.PointId) valueTypes.PointId {
// 	var ret valueTypes.PointId
// 	for range Only.Once {
// 		if endpoint == "" {
// 			endpoint = "virtual"
// 		}
// 		ret = valueTypes.PointId(JoinWithDots(0, "", endpoint, pid))
// 	}
// 	return ret
// }
//
// func JoinStringsWithDots(args ...string) string {
// 	return strings.Join(args, ".")
// }
