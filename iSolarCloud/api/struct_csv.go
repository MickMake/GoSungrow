package api

import (
	"GoSungrow/Only"
	"GoSungrow/graph"
	"encoding/json"
	"errors"
	"fmt"
	"go.pennock.tech/tabular"
	datatable "go.pennock.tech/tabular/auto"
	"os"
	"strconv"
	"strings"
	"time"
)


type Table struct {
	// Data   [][]string
	// Header []string
	table datatable.RenderTable
	Error  error
}


func NewCsv() Table {
	return Table{
		// Data:   nil,
		// Header: nil,
		table:  datatable.New("utf8-heavy"),
		Error:  nil,
	}
}

func (t Table) String() string {
	var ret string
	for range Only.Once {
		ret, t.Error = t.table.Render()
		if t.Error != nil {
			break
		}

		// ret += t.HeaderString()
		// ret += t.DataString()
	}
	return ret
}

func (t Table) Print() {
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

// func (t Table) HeaderString() string {
// 	var ret string
// 	for _, h := range t.Header {
// 		ret += fmt.Sprintf("\"%s\",", h)
// 	}
// 	ret += fmt.Sprintln()
// 	return ret
// }
//
// func (t Table) DataString() string {
// 	var ret string
// 	for _, r := range t.Data {
// 		for _, r := range r {
// 			ret += fmt.Sprintf("\"%s\",", r)
// 		}
// 		ret += fmt.Sprintln()
// 	}
// 	return ret
// }

func (t Table) AddRow(row ...interface{}) Table {
	// t.Data = append(t.Data, row)
	t.table.AddRowItems(row...)
	t.Error = t.getErrors()
	return t
}

func (t Table) SetHeader(header...interface{}) Table {
	// t.Header = header
	t.table.AddHeaders(header...)
	t.Error = t.getErrors()
	return t
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


type GraphRequest struct {
	Title        string `json:"title"`
	TimeColumn   int    `json:"time_column"`
	ValueColumn  int    `json:"value_column"`
	SearchColumn int    `json:"search_column"`
	SearchString string `json:"search_string"`
	FileName     string `json:"file_name"`
	Error        error `json:"-"`
}

func JsonToGraphRequest(j string) GraphRequest {
	var ret GraphRequest
	for range Only.Once {
		ret.Error = json.Unmarshal([]byte(j), &ret)
		if ret.Error != nil {
			break
		}
	}
	// {"title":"","time_column":"","value_column":"","search_column":"","search_string":"","file_name":""}
	return ret
}

func (t Table) Graph(req GraphRequest) error {
	// func (t Table) Graph(name string, timeColumn int, valueColumn int, searchString string, searchColumn int, fileName string) error {
	for range Only.Once {
		foo := graph.New(req.Title)

		t.Error = foo.SetFilename(req.FileName)
		if t.Error != nil {
			break
		}

		var times []time.Time
		var values []float64
		for row := 0; row < t.table.NRows(); row++ {
			// Get the search column
			var cell *tabular.Cell
			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: req.SearchColumn})
			if t.Error != nil {
				continue
			}
			if !strings.Contains(cell.String(), req.SearchString) {
				continue
			}


			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: req.TimeColumn})
			if t.Error != nil {
				continue
			}
			var tim time.Time
			tim, t.Error = time.Parse(DtLayout, cell.String())
			if t.Error != nil {
				continue
			}
			times = append(times, tim)


			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: req.ValueColumn})
			if t.Error != nil {
				continue
			}
			var val float64
			val, t.Error = strconv.ParseFloat(cell.String(), 64)
			if t.Error != nil {
				val = 0
			}
			values = append(values, val)
		}

		t.Error = foo.SetX("Date", times...)
		if t.Error != nil {
			break
		}

		t.Error = foo.SetY("kWh", values...)
		if t.Error != nil {
			break
		}

		foo.SetRangeY(-6000, 12000)
		if t.Error != nil {
			break
		}

		t.Error = foo.Generate()
		if t.Error != nil {
			break
		}
	}

	return t.Error
}
