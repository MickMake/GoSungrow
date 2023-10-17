package output

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/gojson"
	"github.com/anicoll/gosungrow/tablib"
)

// "go.pennock.tech/tabular/auto"
// "github.com/olekukonko/tablewriter"
// "github.com/agrison/go-tablib"
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
	name        string
	directory   string
	filePrefix  string
	title       string
	graph       *Chart
	json        []byte
	raw         []byte
	OutputType  OutputType
	saveAsFile  bool
	graphFilter string
	Error       error

	tablib *tablib.Dataset
	// tabular     tabular.RenderTable
	// tablewriter *tablewriter.Table
	// buf         *bytes.Buffer
	// headers     []string
	// method      int8
}

func NewTable(headers ...string) Table {
	// buf := new(bytes.Buffer)
	t := Table{
		directory:  "",
		filePrefix: "",
		title:      "",
		tablib:     tablib.NewDataset(headers),
		// tabular:     tabular.New("utf8-heavy"),
		// tablewriter: tablewriter.NewWriter(buf),
		// buf:         buf,
		// headers:     []string{},
		// method:      MethodTablib,

		Error: nil,
	}

	t.tablib.SetAlign(tablib.AlignLeft)
	t.tablib.SetEmptyString(" ")
	t.tablib.SetMaxCellSize(128)
	t.tablib.SetWrapStrings(true)
	t.tablib.SetFloatFormat('f')
	t.tablib.SetWrapDelimiter(' ')
	t.tablib.SetDenseMode(true)
	t.tablib.SetSplitConcat(" ")

	// var h1 []interface{}
	// for _, h2 := range headers {
	// 	h1 = append(h1, h2)
	// }
	// t.tabular.AddHeaders(h1...)
	// t.tablewriter.SetHeader(headers)
	// t.headers = headers
	return t
}

// const (
// 	MethodTablib      = iota
// 	MethodTabular     = iota
// 	MethodTableWriter = iota
// )
// func (t *Table) SetTablib()  {
// 	t.method = MethodTablib
// }
// func (t *Table) Tablib() *tablib.Dataset {
// 	return t.tablib
// }
//
// func (t *Table) SetTabular()  {
// 	t.method = MethodTabular
// }
// func (t *Table) Tabular() tabular.RenderTable {
// 	return t.tabular
// }
//
// func (t *Table) SetTableWriter()  {
// 	t.method = MethodTableWriter
// }
// func (t *Table) TableWriter() *tablewriter.Table {
// 	return t.tablewriter
// }

func (t Table) String() string {
	var ret string
	for range Only.Once {
		// switch t.method {
		// 	case 2:
		// 		// Example: GoSungrow api ls endpoints / GoSungrow api ls areas
		// 		if t.buf == nil {
		// 			break
		// 		}
		// 		t.tablewriter.SetBorder(true)
		// 		t.tablewriter.Render()
		// 		ret = t.buf.String()
		// 		t.buf = nil
		//
		// 	case 1:
		// 		ret, t.Error = t.tabular.Render()
		// 		// t.Error = tabular.RenderTo(t.tabular, ret, "") // "csv", "html", "json", "markdown"
		//
		// 	case 0:
		// 		fallthrough
		// 	default:
		// 		if !t.tablib.Valid() {
		// 			break
		// 		}
		// 		ret = t.tablib.Tabular("utf8").String()
		// }

		ret = t.tablib.Tabular("utf8").String()
	}
	return ret
}

func (t *Table) IsValid() bool {
	var yes bool
	for range Only.Once {
		if t.tablib == nil {
			break
		}
		if !t.tablib.Valid() {
			break
		}
		if t.tablib.Height() == 0 {
			break
		}
		if t.tablib.Width() == 0 {
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
	var ret []string
	for range Only.Once {
		// switch t.method {
		// case 2:
		// 	if t.buf == nil {
		// 		break
		// 	}
		// 	ret = t.headers
		//
		// case 1:
		// 	for _, r2 := range t.tabular.Headers() {
		// 		ret = append(ret, r2.String())
		// 	}
		//
		// case 0:
		// 	fallthrough
		// default:
		// 	if !t.tablib.Valid() {
		// 		break
		// 	}
		// 	ret = t.tablib.Headers()
		// }

		if !t.tablib.Valid() {
			break
		}
		ret = t.tablib.Headers()
	}
	return ret
}

func (t *Table) GetSortedHeaders() []string {
	var sorted []string
	for range Only.Once {
		sorted = t.GetHeaders()
		sort.Strings(sorted)
	}
	return sorted
}

func (t *Table) Sort(sort string) {
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		// switch t.method {
		// 	case 2:
		// 		// @TODO -
		// 		// t.tablewriter.???()
		//
		// 	case 1:
		// 		// @TODO -
		// 		// t.tabular.???()
		//
		// 	case 0:
		// 		fallthrough
		// 	default:
		// 		// Make sure we have a header.
		// 		for _, header := range t.tablib.Headers() {
		// 			if header == sort {
		// 				t.tablib = t.tablib.Sort(sort)
		// 				break
		// 			}
		// 		}
		// }
		// Make sure we have a header.

		for _, header := range t.tablib.Headers() {
			if header == sort {
				t.tablib = t.tablib.Sort(sort)
				break
			}
		}
	}
}

func (t *Table) RowLength() int {
	// var ret int
	// switch t.method {
	// 	case 2:
	// 		ret = t.tablewriter.NumLines()
	//
	// 	case 1:
	// 		ret = t.tabular.NRows()
	//
	// 	case 0:
	// 		fallthrough
	// 	default:
	// 		ret = t.tablib.Height()
	// }
	// return ret

	return t.tablib.Height()
}

func (t *Table) GetCell(row int, colName string) (string, interface{}, error) {
	var ret interface{}
	var retType string
	for range Only.Once {
		// switch t.method {
		// 	case 2:
		// 		// @TODO -
		// 		// t.tablewriter.??(row)
		// 		fmt.Println("Not supported.")
		//
		// 	case 1:
		// 		// @TODO -
		// 		var h string
		// 		var i int
		// 		for i, h = range t.headers {
		// 			if h == colName {
		// 				break
		// 			}
		// 		}
		// 		c, err := t.tabular.CellAt(t2.CellLocation {
		// 			Row:    row,
		// 			Column: i,
		// 		})
		// 		if err != nil {
		// 			t.Error = err
		// 			break
		// 		}
		// 		ret = c.Item()
		//
		// 	case 0:
		// 		fallthrough
		// 	default:
		// 		var r map[string]interface{}
		// 		r, t.Error = t.tablib.Row(row)
		// 		ret = r[colName]
		// }

		var r map[string]interface{}
		r, t.Error = t.tablib.Row(row)
		if t.Error != nil {
			break
		}

		ret = r[colName]
		// retType = reflect.TypeOf(ret).String()
		retType = reflect.TypeOf(ret).Name()
	}
	return retType, ret, t.Error
}

func (t *Table) AddRow(row ...interface{}) error {
	// switch t.method {
	// 	case 2:
	// 		var ra []string
	// 		for _, r := range row {
	// 			ra = append(ra, fmt.Sprintf("%s", r))
	// 		}
	// 		t.tablewriter.Append(ra)
	//
	// 	case 1:
	// 		t.tabular.AddRowItems(row...)
	//
	// 	case 0:
	// 		fallthrough
	// 	default:
	// 		t.Error = t.tablib.Append(row)
	// }
	// return t.Error

	return t.tablib.Append(row)
}

func (t *Table) writeFile(data string, perm os.FileMode) error {
	for range Only.Once {
		Mkdir(t.directory)
		// if !DirExists(t.directory) {
		// 	Mkdir(t.directory)
		// }

		fn := path.Join(t.directory, t.filePrefix)
		fmt.Printf("Writing file '%s'\n", fn)
		t.Error = os.WriteFile(fn, []byte(data), perm)
		if t.Error != nil {
			t.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, t.Error))
			break
		}
	}

	return t.Error
}

func (t *Table) SetTitle(title string, args ...interface{}) {
	t.title = fmt.Sprintf(title, args...)
}

func (t *Table) AppendTitle(title string, args ...interface{}) {
	t.title += fmt.Sprintf(title, args...)
}

func (t *Table) GetTitle() string {
	return t.title
}

func (t *Table) GetFilePrefix() string {
	return path.Join(t.directory, t.filePrefix)
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

func (t *Table) SetDirectory(prefix string, args ...interface{}) {
	for range Only.Once {
		if prefix == "" {
			break
		}
		if len(args) == 0 {
			t.directory = prefix
			t.directory = strings.ReplaceAll(t.directory, "[", "")
			t.directory = strings.ReplaceAll(t.directory, "]", "")
			break
		}
		t.directory = fmt.Sprintf(prefix, args...)
		t.directory = strings.ReplaceAll(t.directory, "[", "")
		t.directory = strings.ReplaceAll(t.directory, "]", "")
	}
}

func (t *Table) AppendFilePrefix(prefix string, args ...interface{}) {
	for range Only.Once {
		if prefix == "" {
			break
		}
		if len(args) > 0 {
			prefix = fmt.Sprintf(prefix, args...)
		}
		t.filePrefix += "-" + prefix
	}
}

func (t *Table) PrependFilePrefix(prefix string, args ...interface{}) {
	for range Only.Once {
		if prefix == "" {
			break
		}
		if len(args) > 0 {
			prefix = fmt.Sprintf(prefix, args...)
		}
		t.filePrefix = prefix + "-" + t.filePrefix
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
			t.Error = t.CreateGraph()

		case t.OutputType.IsMarkDown():
			t.Error = t.WriteMarkDown()

		case t.OutputType.IsStruct():
			t.Error = t.WriteStruct()

		default:
			t.Error = t.WriteTable()
		}
	}

	return t.Error
}

func (t *Table) AsTable() string {
	return t.String()
}

func (t *Table) WriteTable() error {
	for range Only.Once {
		if t.IsNotValid() {
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		fmt.Printf("# %s\n", t.title)
		if t.saveAsFile {
			t.filePrefix += "." + StringTypeTable
			t.Error = t.writeFile(t.String(), DefaultFileMode)
			break
		}

		fmt.Print(t.String())
	}
	return t.Error
}

func (t *Table) WriteList() error {
	for range Only.Once {
		if t.IsNotValid() {
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		fmt.Printf("# %s\n", t.title)
		if t.saveAsFile {
			t.filePrefix += "." + StringTypeList
			t.Error = t.writeFile(t.String(), DefaultFileMode)
			break
		}

		fmt.Print(t.String())
	}
	return t.Error
}

func (t *Table) AsCsv() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tablib.Exportable
		result, t.Error = t.tablib.CSV()
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
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		if t.saveAsFile {
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += "." + StringTypeCsv
			t.Error = t.writeFile(t.AsCsv(), DefaultFileMode)
			break
		}

		fmt.Print(t.AsCsv())
	}
	return t.Error
}

func (t *Table) AsXml() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tablib.Exportable
		result, t.Error = t.tablib.XML()
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
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		if t.saveAsFile {
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += "." + StringTypeXML
			t.Error = t.writeFile(t.AsXml(), DefaultFileMode)
			break
		}

		fmt.Print(t.AsXml())
	}
	return t.Error
}

func (t *Table) AsXLSX() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tablib.Exportable
		result, t.Error = t.tablib.XLSX()
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
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		if t.saveAsFile {
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += "." + StringTypeXLSX
			t.Error = t.writeFile(t.AsXLSX(), DefaultFileMode)
			break
		}

		fmt.Print(t.AsXLSX())
	}
	return t.Error
}

func (t *Table) AsJson() string {
	return string(t.raw)
}

func (t *Table) WriteJson() error {
	for range Only.Once {
		// Don't check for valid table data.

		if t.saveAsFile {
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += "." + StringTypeJson
			t.Error = t.writeFile(string(t.json), DefaultFileMode)
			break
		}

		fmt.Printf("%s", t.json)
	}
	return t.Error
}

func (t *Table) AsRaw() string {
	return string(t.json)
}

func (t *Table) AsRawBytes() []byte {
	return t.json
}

func (t *Table) WriteRaw() error {
	for range Only.Once {
		// Don't check for valid table data.

		if t.saveAsFile {
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += "." + StringTypeRaw
			t.Error = t.writeFile(string(t.raw), DefaultFileMode)
			break
		}

		fmt.Printf("%s", t.raw)
	}
	return t.Error
}

func (t *Table) AsStruct() string {
	return string(t.json)
}

func (t *Table) WriteStruct() error {
	for range Only.Once {
		// Don't check for valid table data.

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
			fmt.Printf("# %s\n", t.title)
			t.filePrefix += ".go"
			t.Error = t.writeFile(data, DefaultFileMode)
			break
		}

		fmt.Printf("%s", data)
	}
	return t.Error
}

func (t *Table) AsMarkDown() string {
	var ret string
	for range Only.Once {
		if t.IsNotValid() {
			break
		}

		var result *tablib.Exportable
		result = t.tablib.Markdown()
		if t.Error != nil {
			break
		}
		ret = result.String()
	}
	return ret
}

func (t *Table) WriteMarkDown() error {
	for range Only.Once {
		if t.IsNotValid() {
			msg := fmt.Sprintf("# %s - has no data.", t.name)
			if t.saveAsFile {
				fmt.Println(msg)
			} else {
				_, _ = fmt.Fprintln(os.Stderr, msg)
			}
			break
		}

		if t.saveAsFile {
			_, _ = fmt.Fprintf(os.Stderr, "# %s\n", t.title)
			t.filePrefix += "." + StringTypeMarkDown
			t.Error = t.writeFile(t.AsMarkDown(), DefaultFileMode)
			break
		}

		fmt.Print(t.AsMarkDown())
	}
	return t.Error
}
