package api

import (
	"github.com/MickMake/GoUnify/Only"
)

type DataEntries struct {
	Entries []DataEntry
}

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

func (de *DataEntries) Len() int {
	return len(de.Entries)
}

func (de *DataEntries) Add(ref DataEntry) *DataEntries {
	for range Only.Once {
		if de == nil {
			break
		}
		de.Entries = append(de.Entries, ref)
	}
	return de
}
