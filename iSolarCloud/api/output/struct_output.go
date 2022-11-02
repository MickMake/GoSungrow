package output

import (
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const (
	TypeNone  = iota
	TypeJson  = iota
	TypeCsv   = iota
	TypeList  = iota
	TypeTable = iota
	TypeRaw   = iota
	TypeGraph = iota

	StringTypeNone  = ""
	StringTypeJson  = "json"
	StringTypeCsv   = "csv"
	StringTypeList  = "list"
	StringTypeTable = "table"
	StringTypeRaw   = "raw"
	StringTypeGraph = "graph"
)

type OutputType int


func (out *OutputType) SetNone() {
	*out = TypeNone
}
func (out *OutputType) SetJson() {
	*out = TypeJson
}
func (out *OutputType) SetCsv() {
	*out = TypeCsv
}
func (out *OutputType) SetList() {
	*out = TypeList
}
func (out *OutputType) SetTable() {
	*out = TypeTable
}
func (out *OutputType) SetRaw() {
	*out = TypeRaw
}
func (out *OutputType) SetGraph() {
	*out = TypeGraph
}

func (out *OutputType) IsNone() bool {
	if *out == TypeNone {
		return true
	}
	return false
}
func (out *OutputType) IsJson() bool {
	if *out == TypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsCsv() bool {
	if *out == TypeCsv {
		return true
	}
	return false
}
func (out *OutputType) IsList() bool {
	if *out == TypeList {
		return true
	}
	return false
}
func (out *OutputType) IsTable() bool {
	if *out == TypeTable {
		return true
	}
	return false
}
func (out *OutputType) IsRaw() bool {
	if *out == TypeRaw {
		return true
	}
	return false
}
func (out *OutputType) IsGraph() bool {
	if *out == TypeGraph {
		return true
	}
	return false
}

func (out *OutputType) IsStrNone(t string) bool {
	if t == StringTypeNone {
		return true
	}
	return false
}
func (out *OutputType) IsStrJson(t string) bool {
	if t == StringTypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsStrCsv(t string) bool {
	if t == StringTypeCsv {
		return true
	}
	return false
}
func (out *OutputType) IsStrTable(t string) bool {
	if t == StringTypeTable {
		return true
	}
	return false
}
func (out *OutputType) IsStrList(t string) bool {
	if t == StringTypeList {
		return true
	}
	return false
}
func (out *OutputType) IsStrRaw(t string) bool {
	if t == StringTypeRaw {
		return true
	}
	return false
}
func (out *OutputType) IsStrGraph(t string) bool {
	if t == StringTypeGraph {
		return true
	}
	return false
}


func (out *OutputType) Set(outputType string) {
	for range Only.Once {
		// re := regexp.MustCompile(`^(\w+)\s`)
		// ot := re.FindStringSubmatch(outputType)
		// if len(ot) == 0 {
		// 	break
		// }
		// fmt.Printf("%s\n", ot[0])

		switch strings.ToLower(outputType) {
			case StringTypeJson:
				out.SetJson()
			case StringTypeCsv:
				out.SetCsv()
			case StringTypeRaw:
				out.SetRaw()
			case StringTypeGraph:
				out.SetGraph()
			case StringTypeTable:
				out.SetTable()
			case StringTypeList:
				fallthrough
			default:
				out.SetList()
		}
	}
}
