package output

import (
	"GoSungrow/Only"
	"errors"
	"fmt"
	"go.pennock.tech/tabular"
	datatable "go.pennock.tech/tabular/auto"
	"os"
	"reflect"
)


type Table struct {
	filePrefix string
	title      string
	table      datatable.RenderTable
	graph      *Chart
	json       []byte
	raw        []byte
	OutputType OutputType
	saveAsFile bool
	graphFilter string
	Error      error
}


func NewTable() Table {
	return Table {
		filePrefix: "",
		title: "",
		table: datatable.New("utf8-heavy"),
		// graph: graph.New(""),
		Error: nil,
	}
}

func (t *Table) String() string {
	var ret string
	for range Only.Once {
		if t == nil {
			break
		}

		ret, t.Error = t.table.Render()
		if t.Error != nil {
			break
		}
	}
	return ret
}

func (t *Table) GetHeaders() []tabular.Cell {
	return t.table.Headers()
}

func (t *Table) AllRows() []*tabular.Row {
	return t.table.AllRows()
}

type DataSet []DataRow
type DataRow map[string]string

func (t *Table) SetTitle(title string, args ...interface{}) {
	t.title = fmt.Sprintf(title, args...)
}

func (t *Table) SetRaw(data []byte) {
	t.raw = data
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

func (t *Table) SetFilePrefix(prefix string) {
	t.filePrefix = prefix
}

func (t *Table) SetOutputType(outputType string) {
	t.OutputType.Set(outputType)
}

func (t *Table) SetHeader(header...interface{}) error {
	t.table.AddHeaders(header...)
	t.Error = t.getErrors()
	return t.Error
}

func (t *Table) AddRow(row ...interface{}) error {
	t.table.AddRowItems(row...)
	t.Error = t.getErrors()
	return t.Error
}

func (t *Table) getErrors() error {
	if errs := t.table.Errors(); errs != nil {
		for _, err := range errs {
			t.Error = err
			break
		}
	}
	return t.Error
}

func (t *Table) writeFile(fn string, data string, perm os.FileMode) error {
	for range Only.Once {
		fmt.Printf("Writing file '%s'\n", fn)
		t.Error = os.WriteFile(fn, []byte(data), perm)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, t.Error))
			break
		}
	}

	return t.Error
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

			case t.OutputType.IsCsv():
				t.Error = t.WriteCsv()

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

			default:
		}
	}

	return t.Error
}


func (t *Table) GetTable() string {
	var ret string
	for range Only.Once {
		if t == nil {
			break
		}

		ret, t.Error = t.table.Render()
		if t.Error != nil {
			break
		}
	}
	return ret
}

func (t *Table) WriteTable() error {
	if t.saveAsFile {
		return t.writeFile(t.filePrefix+".txt", t.String(), DefaultFileMode)
	}
	fmt.Print(t.GetTable())
	return nil
}


func (t *Table) GetCsv() string {
	var ret string
	for range Only.Once {
		if t == nil {
			break
		}

		for _, h := range t.GetHeaders() {
			ret += fmt.Sprintf("%s,", h)
		}
		ret += fmt.Sprintf("\n")

		for _, r := range t.AllRows() {
			for _, c := range r.Cells() {
				switch reflect.ValueOf(c.Item()).Type().Name() {
					case "string":
						ret += fmt.Sprintf("\"%s\",", c)
					default:
						ret += fmt.Sprintf("%s,", c)
				}
			}
			ret += fmt.Sprintf("\n")
		}
	}
	return ret
}

func (t *Table) WriteCsv() error {
	if t.saveAsFile {
		return t.writeFile(t.filePrefix+".csv", t.GetCsv(), DefaultFileMode)
	}
	fmt.Print(t.GetCsv())
	return nil
}


func (t *Table) GetJson() string {
	return string(t.raw)
}

func (t *Table) WriteJson() error {
	if t.saveAsFile {
		return t.writeFile(t.filePrefix + ".json", string(t.json), DefaultFileMode)
	}
	fmt.Printf("%s", t.json)
	return nil
}


func (t *Table) GetRaw() string {
	return string(t.json)
}

func (t *Table) WriteRaw() error {
	if t.saveAsFile {
		return t.writeFile(t.filePrefix+".raw", string(t.raw), DefaultFileMode)
	}
	fmt.Printf("%s", t.raw)
	return nil
}
