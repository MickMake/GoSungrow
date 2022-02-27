package output

import (
	"GoSungrow/Only"
	"encoding/json"
	"errors"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"go.pennock.tech/tabular"
	"os"
	"strconv"
	"strings"
	"time"
)


type GraphRequest struct {
	Title        string  `json:"title"`

	TimeColumn   *int     `json:"time_column"`
	ValueColumn  *int     `json:"value_column"`
	UnitsColumn  *int     `json:"units_column"`
	SearchColumn *int     `json:"search_column"`
	SearchString *string  `json:"search_string"`

	MinLeftAxis  *float64 `json:"min_left_axis"`
	MaxLeftAxis  *float64 `json:"max_left_axis"`

	Error        error   `json:"-"`
}

func SetInteger(v int) *int {
	return &v
}
func SetString(v string) *string {
	return &v
}
func SetFloat(v float64) *float64 {
	return &v
}

// type Integer int
// func (c *Integer) Set(f int) (*Integer, bool) {
// 	var changed bool
// 	for range Only.Once {
// 		ff := Integer(f)
// 		if c == nil {
// 			c = &ff
// 			break
// 		}
//
// 		if*c == ff {
// 			break
// 		}
//
// 		changed = true
// 	}
// 	return c, changed
// }
//
// type Float float64
// func (c *Float) Set(f float64) (*Float, bool) {
// 	var changed bool
// 	for range Only.Once {
// 		ff := Float(f)
// 		if c == nil {
// 			c = &ff
// 			break
// 		}
//
// 		if*c == ff {
// 			break
// 		}
//
// 		changed = true
// 	}
// 	return c, changed
// }
//
// type String string
// func (c *String) Set(s string) (*String, bool) {
// 	var changed bool
// 	for range Only.Once {
// 		ss := String(s)
// 		if c == nil {
// 			c = &ss
// 			break
// 		}
//
// 		if*c == ss {
// 			break
// 		}
//
// 		if ss == "" {
// 			break
// 		}
//
// 		changed = true
// 	}
// 	return c, changed
// }


func JsonToGraphRequest(j Json) GraphRequest {
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

func (t *Table) InitGraph(req GraphRequest) {
	if t.graph == nil {
		t.graph = New(req.Title)
	}

	t.graph.req = req
}

func (t *Table) SetGraph(req GraphRequest) error {
	for range Only.Once {
		if t.graph == nil {
			t.graph = New(req.Title)
		}

		t.Error = t.graph.SetFilename(t.filePrefix + ".png")
		if t.Error != nil {
			break
		}

		t.graph.SetRangeY(req.MinLeftAxis, req.MaxLeftAxis)
		if t.Error != nil {
			break
		}

		changed := t.graph.SetGraphSearch(req)
		if !changed {
			break
		}
		if t.Error != nil {
			break
		}

		t.Error = t.ProcessGraphData()
		if t.Error != nil {
			break
		}
	}

	return t.Error
}

func (t *Table) SetGraphFromJson(j Json) error {
	for range Only.Once {
		gr := JsonToGraphRequest(j)
		if gr.Error != nil {
			t.Error = gr.Error
			break
		}

		t.Error = t.SetGraph(gr)
		if t.Error != nil {
			break
		}
	}

	return t.Error
}


func (t *Table) GetSearchColumn() SearchStrings {
	return t.graph.otherSearch
}

func (t *Table) ProcessGraphData() error {
	for range Only.Once {
		req := t.graph.req

		var units string
		var times []time.Time
		var values []float64
		for row := 0; row < t.table.NRows(); row++ {
			// Get the search column
			var cell *tabular.Cell
			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.SearchColumn})
			if t.Error != nil {
				continue
			}
			if !strings.Contains(cell.String(), *req.SearchString) {
				continue
			}

			if req.Title == "" {
				_ = t.SetTitle(cell.String())
			}

			//
			if units == "" {
				if *req.UnitsColumn > 0 {
					cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.UnitsColumn})
					if t.Error != nil {
						continue
					}
					units = cell.String()
				}
			}

			//
			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.TimeColumn})
			if t.Error != nil {
				continue
			}
			var tim time.Time
			tim, t.Error = time.Parse(DateTimeSearchLayout, cell.String())
			if t.Error != nil {
				continue
			}
			times = append(times, tim)

			//
			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.ValueColumn})
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

		t.Error = t.graph.SetX("Date", times...)
		if t.Error != nil {
			break
		}

		t.Error = t.graph.SetY(units, values...)
		if t.Error != nil {
			break
		}
	}

	return t.Error
}

type SearchStrings map[string]int
func (t *Table) FindSearchStrings() error {
	for range Only.Once {
		t.graph.otherSearch = make(SearchStrings)

		for row := 0; row < t.table.NRows(); row++ {
			// Get the search column
			var cell *tabular.Cell
			cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *t.graph.req.SearchColumn})
			if t.Error != nil {
				continue
			}
			if _, ok := t.graph.otherSearch[cell.String()]; ok {
				t.graph.otherSearch[cell.String()] += 1
			} else {
				t.graph.otherSearch[cell.String()] = 0
			}
		}
	}

	return t.Error
}

func (t *Table) SearchStrings() SearchStrings {
	return t.graph.otherSearch
}

func (t *Table) CreateGraph() error {
	return t.graph.Generate()
}

func (t *Table) CreateGraphAll() error {
	return t.graph.Generate()
}


type Chart struct {
	Error error `json:"-"`

	otherSearch SearchStrings
	req GraphRequest
	filename string
	timeSeries1 chart.TimeSeries
	timeSeries2 chart.TimeSeries
	graph chart.Chart
}


// func Randy(min int, max int) float64 {
// 	r := rand.Intn(max - min) + min
// 	return float64(r)
// }
//
// func Test() {
// 	foo := New("Testing 1. 2. 3.")
//
// 	err := foo.SetFilename("HelloWorld.png")
// 	fmt.Println(err)
//
// 	now := time.Now()
// 	var times []time.Time
// 	for i := 0; i < 16; i++ {
// 		now = now.Add(time.Minute * 5)
// 		times = append(times, now)
// 	}
// 	err = foo.SetX("Date", times...)
// 	fmt.Println(err)
//
// 	var values []float64
// 	for i := 0; i < 16; i++ {
// 		then := (float64(i) * Randy(-200, 500)) +  Randy(-5000, 10000)
// 		values = append(values, then)
// 	}
//
// 	err = foo.SetY("Power", values...)
// 	fmt.Println(err)
//
// 	foo.SetRangeY(-6000, 12000)
//
// 	err = foo.Generate()
// 	fmt.Println(err)
// }


func New(title string) *Chart {
	var c Chart

	for range Only.Once {
		c.timeSeries1 = chart.TimeSeries {
			Name:    "",
			Style:   chart.Style{},
			YAxis:   chart.YAxisPrimary,
			XValues: []time.Time{},
			YValues: []float64{},
		}
		c.timeSeries2 = chart.TimeSeries{
			Name:    "",
			Style:   chart.Style{},
			YAxis:   chart.YAxisSecondary,
			XValues: []time.Time{},
			YValues: []float64{},
		}

		c.graph = chart.Chart {
			Title:          title,
			TitleStyle:     chart.Style{},
			ColorPalette:   nil,
			Width:          0,
			Height:         0,
			DPI:            0,
			Background:     chart.Style{},
			Canvas:         chart.Style{},
			XAxis:          chart.XAxis{},
			YAxis:          chart.YAxis{},
			YAxisSecondary: chart.YAxis{},
			Font:           nil,
			Series:         []chart.Series{c.timeSeries1},	// , c.timeSeries2},
			Elements:       nil,
			Log:            nil,
		}

	}
	return &c
}


func (c *Chart) SetFilename(fn string) error {

	for range Only.Once {
		if fn == "" {
			c.Error = errors.New("empty filename")
			break
		}
		fn = strings.TrimSuffix(fn, ".png")
		fn = strings.TrimSuffix(fn, ".PNG")
		c.filename = fn + ".png"
	}
	return c.Error
}

func (c *Chart) SetTitle(fn string) error {
	c.graph.Title = fn
	c.req.Title = fn
	return c.Error
}

func (c *Chart) SetRangeY(min *float64, max *float64) bool {
	var changed bool
	for range Only.Once {
		if min == nil {
			min = c.req.MinLeftAxis
		}

		if max == nil {
			max = c.req.MaxLeftAxis
		}

		c.graph.YAxis.Range = &chart.ContinuousRange {
			Min:        *c.req.MinLeftAxis,
			Max:        *c.req.MaxLeftAxis,
			Domain:     0,
			Descending: false,
		}

		changed = true
	}

	return changed
}

func (c *Chart) SetGraphSearch(req GraphRequest) bool {
	var changed bool
	for range Only.Once {
		if req.SearchString != nil {
			c.req.SearchString = req.SearchString
			changed = true
		}

		if req.TimeColumn != nil {
			c.req.TimeColumn = req.TimeColumn
			changed = true
		}

		if req.ValueColumn != nil {
			c.req.ValueColumn = req.ValueColumn
			changed = true
		}

		if req.UnitsColumn != nil {
			c.req.UnitsColumn = req.UnitsColumn
			changed = true
		}

		if req.SearchColumn != nil {
			c.req.SearchColumn = req.SearchColumn
			changed = true
		}
	}

	return changed
}

func (c *Chart) SetX(name string, values ...time.Time) error {

	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("no X values")
			break
		}

		c.graph.XAxis = chart.XAxis {
			Name:           name,
			NameStyle:      chart.Style{},
			Style:          chart.Style{},
			ValueFormatter: nil,
			Range:          &chart.ContinuousRange {
				Min:        0,
				Max:        0,
				Domain:     0,
				Descending: false,
			},
			TickStyle:      chart.Style{},
			Ticks:          nil,
			TickPosition:   chart.TickPositionUnderTick,
			GridLines:      nil,
			GridMajorStyle: chart.Style{},
			GridMinorStyle: chart.Style{},
		}

		// c.graph.XAxis.ValueFormatter = func(v interface{}) string {
		// 	typed := v.(float64)
		// 	typedDate := chart.TimeFromFloat64(typed)
		// 	return typedDate.Format(DateTimeLayout)
		// 	// t := v.(time.Time)
		// 	// return t.Format(DateTimeLayout)
		// }

		// c.graph.XAxis.TickPosition = chart.TickPositionUnderTick
		// c.graph.XAxis.GridLines = []chart.GridLine {
		// 	{IsMinor: true, Style: chart.Style{}, Value: 600},
		// }
		// c.graph.XAxis.Range = &chart.ContinuousRange {
		// 	Min:        0,
		// 	Max:        0,
		// 	Domain:     0,
		// 	Descending: false,
		// }

		c.timeSeries1.XValues = append(c.timeSeries1.XValues, values...)
		// c.timeSeries2.XValues = append(c.timeSeries2.XValues, values...)
	}
	return c.Error
}

func (c *Chart) SetY(name string, values ...float64) error {

	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("no Y values")
			break
		}

		// c.graph.YAxis = chart.YAxis {
		// 	Name:           name,
		// 	NameStyle:      chart.Style{},
		// 	Style:          chart.Style{},
		// 	Zero:           chart.GridLine{},
		// 	AxisType:       0,
		// 	Ascending:      false,
		// 	ValueFormatter: nil,
		// 	Range:          &chart.ContinuousRange {
		// 		Min:        0,
		// 		Max:        0,
		// 		Domain:     0,
		// 		Descending: false,
		// 	},
		// 	TickStyle:      chart.Style{},
		// 	Ticks:          nil,
		// 	GridLines:      nil,
		// 	GridMajorStyle: chart.Style{},
		// 	GridMinorStyle: chart.Style{},
		// }

		c.graph.YAxis.Name = name

		c.timeSeries1.YValues = append(c.timeSeries1.YValues, values...)
		// c.graph.Series = []chart.Series{ c.timeSeries }
		// c.graph.YAxis.Range = &chart.ContinuousRange {
		// 	Min:        0,
		// 	Max:        0,
		// 	Domain:     0,
		// 	Descending: false,
		// }
	}
	return c.Error
}

func (c *Chart) SetY2(name string, values ...float64) error {

	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("no Y values")
			break
		}

		// c.graph.YAxis = chart.YAxis {
		// 	Name:           name,
		// 	NameStyle:      chart.Style{},
		// 	Style:          chart.Style{},
		// 	Zero:           chart.GridLine{},
		// 	AxisType:       0,
		// 	Ascending:      false,
		// 	ValueFormatter: nil,
		// 	Range:          &chart.ContinuousRange {
		// 		Min:        0,
		// 		Max:        0,
		// 		Domain:     0,
		// 		Descending: false,
		// 	},
		// 	TickStyle:      chart.Style{},
		// 	Ticks:          nil,
		// 	GridLines:      nil,
		// 	GridMajorStyle: chart.Style{},
		// 	GridMinorStyle: chart.Style{},
		// }

		c.graph.YAxis.Name = name

		c.timeSeries1.YValues = append(c.timeSeries1.YValues, values...)
		// c.graph.Series = []chart.Series{ c.timeSeries }
		// c.graph.YAxis.Range = &chart.ContinuousRange {
		// 	Min:        0,
		// 	Max:        0,
		// 	Domain:     0,
		// 	Descending: false,
		// }
	}
	return c.Error
}

func (c *Chart) Generate() error {

	for range Only.Once {
		if c.filename == "" {
			c.Error = errors.New("empty filename")
			break
		}

		c.timeSeries1.Style = chart.Style {
			StrokeColor: drawing.ColorBlue,
			FillColor:   drawing.ColorBlue.WithAlpha(64),
		}

		c.graph.Series = []chart.Series {
			c.timeSeries1,
			// c.timeSeries2,
		}

		c.graph.Title = c.req.Title

		var f *os.File
		f, c.Error = os.Create(c.filename)
		if c.Error != nil {
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		c.graph.Width = 1024
		c.graph.Height = 768
		c.Error = c.graph.Render(chart.PNG, f)
	}
	return c.Error
}



// var lock sync.Mutex
// var graph *chart.Chart
// var ts *chart.TimeSeries
//
// func addData(t time.Time, e time.Duration) {
// 	lock.Lock()
// 	ts.XValues = append(ts.XValues, t)
// 	ts.YValues = append(ts.YValues, chart.TimeMillis(e))
// 	lock.Unlock()
// }
//
// func drawChart(res http.ResponseWriter, req *http.Request) {
// 	start := time.Now()
// 	defer func() {
// 		addData(start, time.Since(start))
// 	}()
// 	if len(ts.XValues) == 0 {
// 		http.Error(res, "no data (yet)", http.StatusBadRequest)
// 		return
// 	}
// 	res.Header().Set("Content-Type", "image/png")
// 	if err := graph.Render(chart.PNG, res); err != nil {
// 		log.Printf("%v", err)
// 	}
// }
//
//
// func Server() {
// 	ts = &chart.TimeSeries{
// 		XValues: []time.Time{},
// 		YValues: []float64{},
// 	}
// 	graph = &chart.Chart{
// 		Series: []chart.Series{ts},
// 	}
// 	http.HandleFunc("/", drawChart)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
