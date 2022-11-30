package output

import (
	"GoSungrow/iSolarCloud/api/GoStruct/gojson"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	tabular "github.com/agrison/go-tablib"
	"os"
	"sort"
	"strings"
)
// "github.com/agrison/go-tablib"
// "go.pennock.tech/tabular"
// "github.com/jbub/tabular"


type Tables map[string]Table

func NewTables() Tables {
	return make(Tables)
}

func (t *Tables) Sort() []string {
	var sorted []string
	for range Only.Once {
		for p := range *t {
			sorted = append(sorted, p)
		}
		sort.Strings(sorted)
	}
	return sorted
}


type Table struct {
	name       string
	filePrefix string
	title      string
	table      *tabular.Dataset
	graph      *Chart
	json       []byte
	raw        []byte
	OutputType OutputType
	saveAsFile bool
	graphFilter string
	Error      error
}

func NewTable(headers ...string) Table {
	return Table {
		filePrefix: "",
		title:      "",
		table:      tabular.NewDataset(headers),
		Error: nil,
	}
}

func (t *Table) String() string {
	var ret string
	for range Only.Once {
		if t == nil {
			break
		}

		if !t.table.Valid() {
			break
		}

		ret = t.table.Tabular("condensed").String()
	}
	return ret
}

func (t *Table) IsValid() bool {
	var yes bool
	for range Only.Once {
		if t.table == nil {
			break
		}
		if !t.table.Valid() {
			break
		}
		if t.table.Height() == 0 {
			break
		}
		if t.table.Width() == 0 {
			break
		}
		yes = true
	}
	return yes
}

func (t *Table) IsNotValid() bool {
	return !t.IsValid()
}

func (t *Table) GetHeaders() []string {
	return t.table.Headers()
}

func (t *Table) GetSortedHeaders() []string {
	var sorted []string

	for range Only.Once {
		for _, h := range t.table.Headers() {
			sorted = append(sorted, h)
		}
		sort.Strings(sorted)
	}
	return sorted
}

func (t *Table) Sort(sort string) {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		// Make sure we have a header.
		for _, header := range t.table.Headers() {
			if header == sort {
				t.table = t.table.Sort(sort)
				break
			}
		}
	}
}

func (t *Table) RowLength() int {
	return t.table.Height()
}

func (t *Table) GetCell(row int, colName string) (interface{}, error) {
	var ret interface{}
	for range Only.Once {
		var r map[string]interface{}
		r, t.Error = t.table.Row(row)
		if t.Error != nil {
			break
		}
		ret = r[colName]
	}
	return ret, t.Error
}

func (t *Table) AddRow(row ...interface{}) error {
	t.Error = t.table.Append(row)
	return t.Error
}

func (t *Table) writeFile(data string, perm os.FileMode) error {
	for range Only.Once {
		fmt.Printf("Writing file '%s'\n", t.filePrefix)
		t.Error = os.WriteFile(t.filePrefix, []byte(data), perm)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", t.filePrefix, t.Error))
			break
		}
	}

	return t.Error
}

func (t *Table) SetTitle(title string, args ...interface{}) {
	t.title = fmt.Sprintf(title, args...)
}

func (t *Table) GetTitle() string {
	return t.title
}

func (t *Table) GetFilePrefix() string {
	return t.filePrefix
}

func (t *Table) SetRaw(data []byte) {
	t.raw = data
}

func (t *Table) AppendRaw(data []byte) {
	t.raw = append(t.raw, data...)
}

func (t *Table) SetJson(data []byte) {
	t.json = data
}

func (t *Table) SetSaveFile(ok bool) {
	t.saveAsFile = ok
}

func (t *Table) SetGraphFilter(filter string) {
	t.graphFilter = filter
}

func (t *Table) SetFilePrefix(prefix string, args ...interface{}) {
	for range Only.Once {
		if prefix == "" {
			break
		}
		if len(args) == 0 {
			t.filePrefix = prefix
			t.filePrefix = strings.ReplaceAll(t.filePrefix, "[", "")
			t.filePrefix = strings.ReplaceAll(t.filePrefix, "]", "")
			break
		}
		t.filePrefix = fmt.Sprintf(prefix, args...)
		t.filePrefix = strings.ReplaceAll(t.filePrefix, "[", "")
		t.filePrefix = strings.ReplaceAll(t.filePrefix, "]", "")
	}
}

func (t *Table) AppendFilePrefix(prefix string, args ...interface{}) {
	for range Only.Once {
		if prefix == "" {
			break
		}
		if len(args) == 0 {
			t.filePrefix += "-" + prefix
			break
		}
		t.filePrefix += "-" + fmt.Sprintf(prefix, args...)
	}
}

func (t *Table) SetOutputType(outputType string) {
	t.OutputType.Set(outputType)
}

func (t *Table) SetName(name string) {
	t.name = name
}

func (t *Table) GetName() string {
	return t.name
}

func (t *Table) Output() error {
	for range Only.Once {
		if t == nil {
			break
		}

		switch {
			case t.OutputType.IsNone():

			case t.OutputType.IsTable():
				t.Error = t.WriteTable()

			case t.OutputType.IsList():
				t.Error = t.WriteList()

			case t.OutputType.IsCsv():
				t.Error = t.WriteCsv()

			case t.OutputType.IsXML():
				t.Error = t.WriteXml()

			case t.OutputType.IsXLSX():
				t.Error = t.WriteXLSX()

			case t.OutputType.IsRaw():
				t.Error = t.WriteRaw()

			case t.OutputType.IsJson():
				t.Error = t.WriteJson()

			case t.OutputType.IsGraph():
				t.Error = t.SetGraphFromJson(Json(t.graphFilter))
				if t.Error != nil {
					break
				}
				t.Error = t.CreateGraph()
				if t.Error != nil {
					break
				}

			case t.OutputType.IsStruct():
				t.Error = t.WriteStruct()

			default:
		}
	}

	return t.Error
}


func (t *Table) GetTable() string {
	return t.String()
}

func (t *Table) WriteTable() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeTable
			t.Error = t.writeFile(t.String(), DefaultFileMode)
			break
		}
		fmt.Printf("# %s\n", t.title)
		fmt.Print(t.String())
	}
	return t.Error
}

func (t *Table) WriteList() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeList
			t.Error = t.writeFile(t.String(), DefaultFileMode)
			break
		}
		fmt.Printf("# %s\n", t.title)
		fmt.Print(t.String())
	}
	return t.Error
}


func (t *Table) GetCsv() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tabular.Exportable
		result, t.Error = t.table.CSV()
		if t.Error != nil {
			break
		}
		ret = result.String()
	}
	return ret
}

func (t *Table) WriteCsv() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeCsv
			t.Error = t.writeFile(t.GetCsv(), DefaultFileMode)
			break
		}
		fmt.Print(t.GetCsv())
	}
	return t.Error
}


func (t *Table) GetXml() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tabular.Exportable
		result, t.Error = t.table.XML()
		if t.Error != nil {
			break
		}
		ret = result.String()
	}
	return ret
}

func (t *Table) WriteXml() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeXML
			t.Error = t.writeFile(t.GetXml(), DefaultFileMode)
			break
		}
		fmt.Print(t.GetXml())
	}
	return t.Error
}


func (t *Table) GetXLSX() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tabular.Exportable
		result, t.Error = t.table.XLSX()
		if t.Error != nil {
			break
		}
		ret = result.String()
	}
	return ret
}

func (t *Table) WriteXLSX() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		// if t.saveAsFile {
			t.filePrefix += "." + StringTypeXLSX
			t.Error = t.writeFile(t.GetXLSX(), DefaultFileMode)
		// 	break
		// }
		// fmt.Print(t.GetXml())
	}
	return t.Error
}


func (t *Table) GetJson() string {
	return string(t.raw)
}

func (t *Table) WriteJson() error {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeJson
			t.Error = t.writeFile(string(t.json), DefaultFileMode)
			break
		}
		fmt.Printf("%s", t.json)
	}
	return t.Error
}


func (t *Table) GetRaw() string {
	return string(t.json)
}

func (t *Table) GetRawBytes() []byte {
	return t.json
}

func (t *Table) WriteRaw() error {
	for range Only.Once {
		// if t.IsNotValid() {
		// 	break
		// }

		if t.saveAsFile {
			t.filePrefix += "." + StringTypeRaw
			t.Error = t.writeFile(string(t.raw), DefaultFileMode)
			break
		}
		fmt.Printf("%s", t.raw)
	}
	return t.Error
}


func (t *Table) GetStruct() string {
	return string(t.json)
}

func (t *Table) WriteStruct() error {
	for range Only.Once {
		// if t.IsNotValid() {
		// 	break
		// }

		var data string
		var options gojson.Options
		if t.name == "" {
			t.name = "Package"
		}
		options.PackageName(t.name)

		if string(t.json) == "{}" {
			options.StructureName("Response")
			data, t.Error = gojson.Parse(options, t.raw)
		} else {
			options.StructureName("ResultData")
			data, t.Error = gojson.Parse(options, t.json)
		}

		if t.saveAsFile {
			t.filePrefix += ".go"
			t.Error = t.writeFile(data, DefaultFileMode)
			break
		}
		fmt.Printf("%s", data)
	}
	return t.Error
}
