package output

import (
	"GoSungrow/Only"
	"errors"
	"fmt"
	"go.pennock.tech/tabular"
	datatable "go.pennock.tech/tabular/auto"
	"os"
)


type Table struct {
	filePrefix string
	title      string
	table      datatable.RenderTable
	graph      *Chart
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
func (t *Table) GetData()  {
	headers := t.GetHeaders()
	for _, r := range t.AllRows() {
		for i, c := range r.Cells() {
			fmt.Printf("Header: %s\tValue: %v\n", headers[i], c)
			// err = foo.SensorPublishConfig()
		}
		// r.Cells()
		// err = foo.SensorPublishConfig()
	}
}

func (t *Table) Print() {
	for range Only.Once {
		t.Error = t.table.RenderTo(os.Stdout)
		if t.Error != nil {
			break
		}
		// fmt.Printf("Headers: %v\n", t.table.Headers())
		// fmt.Printf("NRows: %v\n", t.table.NRows())
		//
		// for ri, r := range t.table.AllRows() {
		// 	fmt.Println(ri)
		// 	for _, cr := range r.Cells() {
		// 		cr.
		// 	}
		// }
		// fmt.Printf("AllRows: %v\n", t.table.AllRows())
		//
		// fmt.Println(t)
	}
}

func (t *Table) SetTitle(title string) error {
	t.title = title
	return t.Error
}

func (t *Table) SetFilePrefix(prefix string) error {
	t.filePrefix = prefix
	return t.Error
}

func (t *Table) SetHeader(header...interface{}) error {
	t.table.AddHeaders(header...)
	t.Error = t.getErrors()
	return t.Error
}

func (t Table) AddRow(row ...interface{}) error {
	t.table.AddRowItems(row...)
	t.Error = t.getErrors()
	return t.Error
}

func (t Table) getErrors() error {
	if errs := t.table.Errors(); errs != nil {
		for _, err := range errs {
			t.Error = err
			break
		}
	}
	return t.Error
}

func (t *Table) writeFile(fn string, perm os.FileMode) error {
	for range Only.Once {
		fmt.Printf("Writing file '%s'\n", fn)
		t.Error = os.WriteFile(fn, []byte(t.String()), perm)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, t.Error))
			break
		}
	}

	return t.Error
}

func (t *Table) WriteFile(fn string, perm os.FileMode) error {
	for range Only.Once {
		fmt.Printf("Writing file '%s'\n", fn)
		t.Error = os.WriteFile(fn, []byte(t.String()), perm)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, t.Error))
			break
		}
	}

	return t.Error
}

func (t *Table) WriteCsvFile() error {
	return t.writeFile(t.filePrefix + ".csv", DefaultFileMode)
}
