package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// 	"go.pennock.tech/tabular"
// 	tabular "github.com/agrison/go-tablib"

type GraphRequest struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`

	TimeColumn  *string  `json:"time_column"`
	DataColumn  *string  `json:"value_column"`
	UnitsColumn *string  `json:"units_column"`
	NameColumn  *string  `json:"name_column"`
	DataUnit    *string  `json:"unit"`
	DataMin     *float64 `json:"min"`
	DataMax     *float64 `json:"max"`
	Width       *int     `json:"width"`
	Height      *int     `json:"height"`

	// DataColumn   *string  `json:"data_column"`
	// SearchString *string  `json:"search_string"`

	Error error `json:"-"`
}

func JsonToGraphRequest(j Json) GraphRequest {
	var ret GraphRequest
	for range Only.Once {
		if j == "" {
			break
		}
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
		t.graph = New()
		_ = t.graph.SetTitle(req.Title)
	}

	t.graph.req = req
}

func (t *Table) SetGraph(req GraphRequest) error {
	for range Only.Once {
		if t.graph == nil {
			t.graph = New()
		}

		t.graph.req = req

		_ = t.graph.SetTitle(req.Title)
		if t.Error != nil {
			break
		}

		t.graph.SetWidth(req.Width)
		if t.Error != nil {
			break
		}

		t.graph.SetHeight(req.Height)
		if t.Error != nil {
			break
		}

		changed := t.graph.SetColumns(req)
		if !changed {
			break
		}
		if t.Error != nil {
			break
		}

		t.graph.SetRangeY(req.DataMin, req.DataMax)
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
		if j == "" {
			break
		}

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

func (t *Table) ProcessGraphData() error {
	t.Error = t.graph.ProcessGraphData(t)
	return t.Error

	// for range Only.Once {
	// 	req := t.graph.req
	//
	// 	t.graph.searchName = ""
	// 	var units string
	// 	var times []time.Time
	// 	var values []float64
	//
	// 	if t.graph.req.DataUnit != nil {
	// 		units = *t.graph.req.DataUnit
	// 	}
	//
	// 	for row := 0; row < t.RowLength(); row++ {
	// 		var cell interface{}
	//
	// 		// Get the search column
	// 		// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.SearchColumn})
	// 		// cell, t.Error = t.GetCell(row, *req.DataColumn)
	// 		// if t.Error != nil {
	// 		// 	continue
	// 		// }
	// 		// if !strings.Contains(cell.String(), *req.SearchString) {
	// 		// 	continue
	// 		// }
	//
	//
	// 		if req.Title == "" {
	// 			req.Title = cell.(string)
	// 			t.SetTitle(cell.(string))
	// 		}
	//
	//
	// 		// if t.graph.searchName == "" {
	// 			if req.NameColumn != nil {
	// 				// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.NameColumn})
	// 				cell, t.Error = t.GetCell(row, *req.NameColumn)
	// 				if t.Error != nil {
	// 					continue
	// 				}
	// 				t.graph.searchName = cell.(string)
	// 			}
	// 		// }
	//
	//
	// 		// Get units
	// 		if units == "" {
	// 			if req.UnitsColumn != nil {
	// 				// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.UnitsColumn})
	// 				cell, t.Error = t.GetCell(row, *req.UnitsColumn)
	// 				if t.Error != nil {
	// 					continue
	// 				}
	// 				units = cell.(string)
	// 			}
	// 		}
	//
	//
	// 		// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.TimeColumn})
	// 		cell, t.Error = t.GetCell(row, *req.TimeColumn)
	// 		if t.Error != nil {
	// 			continue
	// 		}
	// 		var tim time.Time
	// 		tim, t.Error = time.ParseInLocation(DateTimeSearchLayout, cell.(string), time.Local)	// @TODO - May have to revisit this!
	// 		if t.Error != nil {
	// 			continue
	// 		}
	// 		times = append(times, tim)
	//
	//
	// 		// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.ValueColumn})
	// 		cell, t.Error = t.GetCell(row, *req.DataColumn)
	// 		if t.Error != nil {
	// 			continue
	// 		}
	// 		var val float64
	// 		val, t.Error = strconv.ParseFloat(cell.(string), 64)
	// 		if t.Error != nil {
	// 			val = 0
	// 		}
	// 		values = append(values, val)
	// 	}
	//
	// 	var filename string
	// 	// filename = fmt.Sprintf("%s-%s.png", t.filePrefix, strings.ReplaceAll(t.graph.searchName, " ", ""))
	// 	filename = fmt.Sprintf("%s.png", t.filePrefix)
	// 	t.Error = t.graph.SetFilename(filename)
	// 	if t.Error != nil {
	// 		break
	// 	}
	//
	// 	t.Error = t.graph.SetX("Date", times...)
	// 	if t.Error != nil {
	// 		break
	// 	}
	//
	// 	t.Error = t.graph.SetY(units, values...)
	// 	if t.Error != nil {
	// 		break
	// 	}
	// }
	// return t.Error
}

func (t *Table) CreateGraph() error {
	for range Only.Once {
		t.Error = t.SetGraphFromJson(Json(t.graphFilter))
		if t.Error != nil {
			break
		}

		t.Error = t.graph.Generate()

		// if t.graph.req.SearchString == nil {
		// 	break
		// }
		// if *t.graph.req.SearchString != "" {
		// 	t.Error = t.graph.Generate()
		// 	break
		// }
		//
		// t.Error = t.FindSearchStrings()
		// if t.Error != nil {
		// 	break
		// }
		//
		// for s := range t.graph.otherSearch {
		// 	t.graph.req.SearchString = &s
		// 	t.Error = t.ProcessGraphData()
		// 	if t.Error != nil {
		// 		continue
		// 	}
		// 	t.Error = t.graph.Generate()
		// 	if t.Error != nil {
		// 		continue
		// 	}
		// }
	}

	return t.Error
}

type Chart struct {
	Error error `json:"-"`

	searchName  string
	req         GraphRequest
	filename    string
	timeSeries1 chart.TimeSeries
	timeSeries2 chart.TimeSeries
	annotation  chart.AnnotationSeries
	minSeries   *chart.MinSeries
	maxSeries   *chart.MaxSeries
	graph       chart.Chart
}

func New() *Chart {
	var c Chart

	for range Only.Once {
		c.timeSeries1 = chart.TimeSeries{
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

		c.annotation = chart.AnnotationSeries{
			Annotations: []chart.Value2{},
		}

		c.graph = chart.Chart{
			Title:        "",
			TitleStyle:   chart.Style{},
			ColorPalette: nil,
			Width:        0,
			Height:       0,
			DPI:          0,
			Background: chart.Style{
				Padding: chart.Box{
					Top:    50,
					Left:   25,
					Right:  25,
					Bottom: 10,
				},
				FillColor: drawing.ColorFromHex("efefef"),
			},
			Canvas:         chart.Style{},
			XAxis:          chart.XAxis{},
			YAxis:          chart.YAxis{},
			YAxisSecondary: chart.YAxis{},
			Font:           nil,
			Series:         []chart.Series{}, // c.timeSeries1, c.annotation },	// , c.timeSeries2},
			Elements:       nil,
			Log:            nil,
		}

	}
	return &c
}

func (c *Chart) ProcessGraphData(table *Table) error {
	for range Only.Once {
		// req := c.req

		c.searchName = ""
		// var units string
		var times []time.Time
		var values []float64

		// if c.req.DataUnit != nil {
		// 	units = *c.req.DataUnit
		// }
		for row := 0; row < table.RowLength(); row++ {
			var cell interface{}
			var cellType string

			// Get the search column
			// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.SearchColumn})
			// cell, t.Error = t.GetCell(row, *req.DataColumn)
			// if t.Error != nil {
			// 	continue
			// }
			// if !strings.Contains(cell.String(), *req.SearchString) {
			// 	continue
			// }

			// if req.Title == "" {
			// 	req.Title = cell.(string)
			// 	c.SetTitle(cell.(string))
			// }

			// if t.graph.searchName == "" {
			if c.req.NameColumn != nil {
				// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.NameColumn})
				cellType, cell, c.Error = table.GetCell(row, *c.req.NameColumn)
				if c.Error != nil {
					continue
				}
				val, isNil, ok := valueTypes.AnyToUnitValue(cell, "", *c.req.DataUnit, "", "")
				if isNil {
					continue
				}
				if !ok {
					continue
				}
				units := val.String()
				c.req.NameColumn = &units

				// switch cellType {
				// 	case "string":
				// 		c.searchName = cell.(string)
				// 	case valueTypes.TypeUnitValue:
				// 		c.searchName = cell.(valueTypes.UnitValue).String()
				// 	default:
				// 		c.searchName = fmt.Sprintf("%s", cell)
				// }
			}
			// }

			// Get units
			if c.req.DataUnit == nil {
				if c.req.UnitsColumn != nil {
					// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.UnitsColumn})
					cellType, cell, c.Error = table.GetCell(row, *c.req.UnitsColumn)
					if c.Error != nil {
						continue
					}
					val, isNil, ok := valueTypes.AnyToUnitValue(cell, "", *c.req.DataUnit, "", "")
					if isNil {
						continue
					}
					if !ok {
						continue
					}
					units := val.String()
					c.req.DataUnit = &units

					// switch cellType {
					// 	case "string":
					// 		units := cell.(string)
					// 		c.req.DataUnit = &units
					// 	case valueTypes.TypeUnitValue:
					// 		units := cell.(valueTypes.UnitValue).String()
					// 		c.req.DataUnit = &units
					// 	default:
					// 		units := fmt.Sprintf("%s", cell)
					// 		c.req.DataUnit = &units
					// }
				}
			}

			// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.TimeColumn})
			cellType, cell, c.Error = table.GetCell(row, *c.req.TimeColumn)
			if c.Error != nil {
				continue
			}
			// val, isNil, ok := valueTypes.AnyToUnitValue(cell, "", *c.req.DataUnit, "", "")
			// if isNil {
			// 	continue
			// }
			// if !ok {
			// 	continue
			// }
			// if !val.First().IsTypeDateTime() {
			// 	continue
			// }
			// times = append(times, val.First().IsTypeDateTime())
			switch cellType {
			case valueTypes.TypeUnitValue:
				var tim time.Time
				tim, c.Error = time.ParseInLocation(DateTimeSearchLayout, cell.(valueTypes.UnitValue).String(), time.Local) // @TODO - May have to revisit this!
				if c.Error != nil {
					continue
				}
				times = append(times, tim)
			case "string":
				var tim time.Time
				tim, c.Error = time.ParseInLocation(DateTimeSearchLayout, cell.(string), time.Local) // @TODO - May have to revisit this!
				if c.Error != nil {
					continue
				}
				times = append(times, tim)
			case valueTypes.TypeDateTime:
				times = append(times, cell.(valueTypes.DateTime).Time)
			default:
				var tim time.Time
				tim, c.Error = time.ParseInLocation(DateTimeSearchLayout, fmt.Sprintf("%s", cell), time.Local) // @TODO - May have to revisit this!
				if c.Error != nil {
					continue
				}
				times = append(times, tim)
			}

			// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *req.ValueColumn})
			cellType, cell, c.Error = table.GetCell(row, *c.req.DataColumn)
			if c.Error != nil {
				continue
			}
			u := ""
			if c.req.DataUnit != nil {
				u = *c.req.DataUnit
			}
			val, isNil, ok := valueTypes.AnyToUnitValue(cell, "", u, "", "")
			if isNil {
				continue
			}
			if !ok {
				continue
			}
			values = append(values, val.First().Value())

			// switch cellType {
			// 	case "string":
			// 		var val float64
			// 		val, c.Error = strconv.ParseFloat(cell.(string), 64)
			// 		if c.Error != nil {
			// 			val = 0
			// 		}
			// 		values = append(values, val)
			// 	case "float32":
			// 		values = append(values, float64(cell.(float32)))
			// 	case "float64":
			// 		values = append(values, cell.(float64))
			// 	case valueTypes.TypeUnitValue:
			// 		units := cell.(valueTypes.UnitValue).String()
			// 		c.req.DataUnit = &units
			// 	default:
			// 		units := fmt.Sprintf("%s", cell)
			// 		c.req.DataUnit = &units
			// }
		}

		c.Error = c.SetFilename(path.Join(table.directory, table.filePrefix))
		if c.Error != nil {
			break
		}

		c.Error = c.SetX("Date", times...)
		if c.Error != nil {
			break
		}

		var unit string
		if c.req.DataUnit != nil {
			unit = *c.req.DataUnit
		}
		c.Error = c.SetY(unit, values...)
		if c.Error != nil {
			break
		}

		c.timeSeries1.Style = chart.Style{
			StrokeColor: drawing.ColorBlue,
			FillColor:   drawing.ColorBlue.WithAlpha(64),
		}

		c.minSeries = &chart.MinSeries{
			Name: "Min",
			Style: chart.Style{
				// StrokeColor:     chart.ColorAlternateGray,
				// StrokeColor:     drawing.ColorBlue,
				// FillColor:       drawing.ColorBlue.WithAlpha(50),
				StrokeColor:     drawing.ColorBlue,
				StrokeDashArray: []float64{5.0, 5.0},
				DotWidth:        0.5,
			},
			InnerSeries: c.timeSeries1,
		}

		c.maxSeries = &chart.MaxSeries{
			Name: "Max",
			Style: chart.Style{
				// StrokeColor:     chart.ColorAlternateGray,
				// StrokeColor:     drawing.ColorGreen,
				// FillColor:       drawing.ColorGreen.WithAlpha(64),
				StrokeColor:     drawing.ColorBlue,
				StrokeDashArray: []float64{5.0, 5.0},
				DotWidth:        1.0,
			},
			InnerSeries: c.timeSeries1,
		}

		c.graph.Series = []chart.Series{
			c.timeSeries1,
			// c.annotation,
			c.minSeries,
			c.maxSeries,
			chart.LastValueAnnotationSeries(c.minSeries),
			chart.LastValueAnnotationSeries(c.maxSeries),
			// c.timeSeries2,
		}

		c.graph.DPI = 150
		c.graph.Elements = []chart.Renderable{chart.Legend(&c.graph)}
	}

	return c.Error
}

func (c *Chart) SetWidth(v *int) {
	for range Only.Once {
		if v == nil {
			c.graph.Width = 3840
			break
		}
		if *v == 0 {
			c.graph.Width = 3840
			break
		}
		c.graph.Width = *v
	}
}

func (c *Chart) SetHeight(v *int) {
	for range Only.Once {
		if v == nil {
			c.graph.Height = 1080
			break
		}
		if *v == 0 {
			c.graph.Height = 1080
			break
		}
		c.graph.Height = *v
	}
}

func (c *Chart) SetFilename(fn string) error {
	for range Only.Once {
		if fn == "" {
			c.Error = errors.New("empty filename")
			break
		}

		fn = strings.TrimSuffix(fn, ".png")
		fn = strings.TrimSuffix(fn, ".PNG")
		fn = strings.ReplaceAll(fn, " ", "_")
		c.filename = fn + ".png"
	}
	return c.Error
}

func (c *Chart) SetTitle(title string) error {
	c.graph.Title = title
	// c.title = title
	return c.Error
}

func (c *Chart) SetRangeY(min *float64, max *float64) bool {
	var changed bool
	for range Only.Once {
		if min == nil {
			min = c.req.DataMin
		}
		if min == nil {
			break
		}

		if max == nil {
			max = c.req.DataMax
		}

		if max == nil {
			break
		}

		c.graph.YAxis.Range = &chart.ContinuousRange{
			Min:        *min,
			Max:        *max,
			Domain:     0,
			Descending: false,
		}

		changed = true
	}

	return changed
}

func (c *Chart) SetColumns(req GraphRequest) bool {
	var changed bool
	for range Only.Once {
		// if req.SearchString != nil {
		// 	c.req.SearchString = req.SearchString
		// 	changed = true
		// }

		if req.TimeColumn != nil {
			c.req.TimeColumn = req.TimeColumn
			changed = true
		}

		if req.DataColumn != nil {
			c.req.DataColumn = req.DataColumn
			changed = true
		}

		if req.UnitsColumn != nil {
			c.req.UnitsColumn = req.UnitsColumn
			changed = true
		}

		// if req.DataColumn != nil {
		// 	c.req.DataColumn = req.DataColumn
		// 	changed = true
		// }
	}

	return changed
}

func (c *Chart) SetX(name string, values ...time.Time) error {
	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("no X values")
			break
		}

		c.graph.XAxis = chart.XAxis{
			Name:           name,
			NameStyle:      chart.Style{},
			Style:          chart.Style{},
			ValueFormatter: chart.TimeValueFormatterWithFormat(DateTimeLayout),
			Range: &chart.ContinuousRange{
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

		c.timeSeries1.XValues = values
		// c.timeSeries1.XValues = append(c.timeSeries1.XValues, values...)
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

		// var minX float64
		// var maxX float64
		var minY float64
		var maxY float64
		for _, value := range values {
			if value > maxY {
				maxY = value
				// maxX = c.timeSeries1.XValues[index].
				// maxX = 0.0
			}
			if value < minY {
				minY = value
				// minX = index
				// minX = 50.0
			}
		}

		if (c.req.DataMin == nil) && (c.req.DataMax == nil) {
			rng := ((maxY - minY) / 2) * 0.25
			nMin := minY - rng
			if (nMin < 0.0) && (minY >= 0.0) {
				// If the subtraction has sent us negative.
				nMin = 0.0
			}
			nMax := maxY + rng
			c.SetRangeY(&nMin, &nMax)
		}

		// c.annotation.Annotations = append(c.annotation.Annotations, chart.Value2 {
		// 	Style:  chart.Style {
		// 		StrokeColor: drawing.ColorGreen,
		// 		FillColor:   drawing.ColorGreen.WithAlpha(64),
		// 	},
		// 	Label:  "Min " + strconv.FormatFloat(minY, 'f', -1, 64),
		// 	XValue: minX,
		// 	YValue: minY,
		// })
		//
		// c.annotation.Annotations = append(c.annotation.Annotations, chart.Value2 {
		// 	Style:  chart.Style {
		// 		StrokeColor: drawing.ColorGreen,
		// 		FillColor:   drawing.ColorGreen.WithAlpha(64),
		// 	},
		// 	Label:  "Max " + strconv.FormatFloat(maxY, 'f', -1, 64),
		// 	XValue: maxX,
		// 	YValue: maxY,
		// })

		c.graph.YAxis.Name = name
		c.timeSeries1.YValues = values
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

		fmt.Printf("Creating graph file '%s'\n", c.filename)
		var f *os.File
		f, c.Error = os.Create(c.filename)
		if c.Error != nil {
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		c.Error = c.graph.Render(chart.PNG, f)
	}
	return c.Error
}
