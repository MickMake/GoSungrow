package graph

import (
	"GoSungrow/Only"
	"errors"
	"github.com/wcharczuk/go-chart/v2"
	"math/rand"
	"os"
	"strings"
	"time"
)


type Chart struct {
	Error error `json:"-"`

	filename string
	timeSeries chart.TimeSeries
	graph chart.Chart
}


func Randy(min int, max int) float64 {
	r := rand.Intn(max - min) + min
	return float64(r)
}

func Test() {
	// foo := New("Testing 1. 2. 3.")
	//
	// err := foo.SetFilename("HelloWorld.png")
	// fmt.Println(err)
	//
	// now := time.Now()
	// var times []time.Time
	// for i := 0; i < 16; i++ {
	// 	now = now.Add(time.Minute * 5)
	// 	times = append(times, now)
	// }
	// err = foo.SetX("Date", times...)
	// fmt.Println(err)
	//
	// var values []float64
	// for i := 0; i < 16; i++ {
	// 	then := (float64(i) * Randy(-200, 500)) +  Randy(-5000, 10000)
	// 	values = append(values, then)
	// }
	//
	// err = foo.SetY("Power", values...)
	// fmt.Println(err)
	//
	// foo.SetRangeY(-6000, 12000)
	//
	// err = foo.Generate()
	// fmt.Println(err)
}

func New(title string) *Chart {
	var c Chart

	for range Only.Once {
		if title == "" {
			c.Error = errors.New("empty title")
			break
		}

		c.timeSeries = chart.TimeSeries {
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
			Series:         []chart.Series{c.timeSeries},
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

func (c *Chart) SetRangeY(min float64, max float64) {
	c.graph.YAxis.Range = &chart.ContinuousRange {
		Min:        min,
		Max:        max,
		Domain:     0,
		Descending: false,
	}
	// c.graph.YAxis.Range.SetMin(min)
	// c.graph.YAxis.Range.SetMax(max)
}

func (c *Chart) SetX(name string, values ...time.Time) error {

	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("empty X values")
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
			TickPosition:   0,
			GridLines:      nil,
			GridMajorStyle: chart.Style{},
			GridMinorStyle: chart.Style{},
		}

		c.timeSeries.XValues = append(c.timeSeries.XValues, values...)
		c.graph.Series = []chart.Series{ c.timeSeries }
	}
	return c.Error
}

func (c *Chart) SetY(name string, values ...float64) error {

	for range Only.Once {
		if len(values) == 0 {
			c.Error = errors.New("empty Y values")
			break
		}

		c.graph.YAxis = chart.YAxis{
			Name:           name,
			NameStyle:      chart.Style{},
			Style:          chart.Style{},
			Zero:           chart.GridLine{},
			AxisType:       0,
			Ascending:      false,
			ValueFormatter: nil,
			Range:          &chart.ContinuousRange {
				Min:        0,
				Max:        0,
				Domain:     0,
				Descending: false,
			},
			TickStyle:      chart.Style{},
			Ticks:          nil,
			GridLines:      nil,
			GridMajorStyle: chart.Style{},
			GridMinorStyle: chart.Style{},
		}

		c.timeSeries.YValues = append(c.timeSeries.YValues, values...)
		c.graph.Series = []chart.Series{ c.timeSeries }
	}
	return c.Error
}

func (c *Chart) Generate() error {

	for range Only.Once {
		if c.filename == "" {
			c.Error = errors.New("empty filename")
			break
		}

		f, _ := os.Create(c.filename)

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
