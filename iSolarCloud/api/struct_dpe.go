package api

import (
	"github.com/MickMake/GoUnify/Only"
)


type DataEntries struct {
	Entries []DataEntry
}

// func NewDataPointEntries() DataEntries {
// 	return DataEntries{
// 		Entries: []DataEntry{},
// 		// Map: &GoStruct.StructMap{},
// 	}
// }

// func (de *DataEntries) Hide() {
// 	for range Only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Hide = true
// 		}
// 	}
// }

const LastEntry = -1
func (de *DataEntries) GetEntry(index int) *DataEntry {
	for range Only.Once {
		if de == nil {
			return nil
		}
		if de.Entries == nil {
			return nil
		}
		l := de.Len() - 1
		if index > l {
			index = l
			break
		}
		if index < 0 {
			index = l + index + 1
			if index < 0 {
				index = 0
			}
		}
	}
	return &(de.Entries[index])
	// return &(de.Entries[index])
}

// func (de *DataEntries) GetEntryValue(index int) valueTypes.UnitValue {
// 	var ret valueTypes.UnitValue
// 	for range Only.Once {
// 		ref := de.GetEntry(index)
// 		if ref == nil {
// 			break
// 		}
// 		ret = ref.Value
// 	}
// 	return ret
// }

// func (de *DataEntries) GetFloat() float64 {
// 	var ret float64
// 	for range Only.Once {
// 		ref := de.GetEntry(0)
// 		if ref == nil {
// 			break
// 		}
// 		ret = ref.Value.Value()
// 	}
// 	return ret
// }

// func (de *DataEntries) MatchPointId(pointId string) bool {
// 	var yes bool
// 	for range Only.Once {
// 		for _, v := range de.Entries {
// 			if v.Point.Id.String() == pointId {
// 				yes = true
// 				break
// 			}
// 		}
// 	}
// 	return yes
// }

func (de *DataEntries) Len() int {
	return len(de.Entries)
}

// func (de *DataEntries) GetUnits() string {
// 	var unit string
// 	for range Only.Once {
// 		for _, v := range de.Entries {
// 			unit = v.Point.Unit
// 			break
// 		}
// 	}
// 	return unit
// }

func (de *DataEntries) Add(ref DataEntry) *DataEntries {
	for range Only.Once {
		if de == nil {
			break
		}
		de.Entries = append(de.Entries, ref)
	}
	return de
}

// func (de *DataEntries) SetUnits(units string) *DataEntries {
// 	for range Only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Point.Unit = units
// 		}
// 	}
// 	return de
// }

// func (de *DataEntries) SetGroupName(groupName string) *DataEntries {
// 	for range Only.Once {
// 		for i := range de.Entries {
// 			de.Entries[i].Point.GroupName = groupName
// 		}
// 	}
// 	return de
// }

// func (de *DataEntries) SetTimestamp(timeStamp valueTypes.DateTime) *DataEntries {
// 	for range Only.Once {
// 		for i := range de.Entries {
// 			// dt := valueTypes.SetDateTimeString(timeStamp)
// 			de.Entries[i].Date = timeStamp
// 		}
// 	}
// 	return de
// }

// func (de *DataEntries) Copy() DataEntries {
// 	var ret DataEntries
// 	for _, d := range de.Entries {
// 		// var point Point
// 		// point = *d.Point
// 		// d.Point = &point
// 		ret.Entries = append(ret.Entries, d.Copy())
// 	}
// 	return ret
// }

// func (de *DataEntries) MakeState(state bool) *DataEntries {
// 	for i := range de.Entries {
// 		de.Entries[i].MakeState(state)
// 	}
// 	return de
// }

// func (de *DataEntries) SetFloat(value float64, unit string, Type string) *DataEntries {
// 	for i := range de.Entries {
// 		de.Entries[i].MakeFloat(value, unit, Type)
// 	}
// 	return de
// }

// func (de *DataEntries) FloatToState(value float64) *DataEntries {
// 	for i := range de.Entries {
// 		if value == 0 {
// 			de.Entries[i].MakeState(false)
// 			break
// 		}
// 		de.Entries[i].MakeState(true)
// 	}
// 	return de
// }
