package output

// type SearchStrings map[string]int
// func (t *Table) FindSearchStrings() error {
// 	for range only.Once {
// 		t.graph.otherSearch = make(SearchStrings)
//
// 		for row := 0; row < t.RowLength(); row++ {
// 			// Get the search column
// 			var cell interface{}
// 			// cell, t.Error = t.table.CellAt(tabular.CellLocation{Row: row, Column: *t.graph.req.SearchColumn})
// 			cell, t.Error = t.GetCell(row, *t.graph.req.SearchColumn)
// 			if t.Error != nil {
// 				continue
// 			}
// 			if _, ok := t.graph.otherSearch[cell.(string)]; ok {
// 				t.graph.otherSearch[cell.(string)] += 1
// 			} else {
// 				t.graph.otherSearch[cell.(string)] = 0
// 			}
// 		}
// 	}
//
// 	return t.Error
// }
//
// func (t *Table) SearchStrings() SearchStrings {
// 	return t.graph.otherSearch
// }
//
// func (t *Table) CreateGraphAll() error {
// 	return t.graph.Generate()
// }
//
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
//
// var zero = float64(0)
//
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
//
// TimeValueFormatterWithFormat is a ValueFormatter for timestamps with a given format.
// func formatTime(v interface{}, dateFormat string) string {
// 	if typed, isTyped := v.(time.Time); isTyped {
// 		return typed.Format(dateFormat)
// 	}
// 	if typed, isTyped := v.(int64); isTyped {
// 		return time.Unix(0, typed).Format(dateFormat)
// 	}
// 	if typed, isTyped := v.(float64); isTyped {
// 		return time.Unix(0, int64(typed)).Format(dateFormat)
// 	}
// 	return ""
// }
//
// func (t *Table) GetSearchColumn() SearchStrings {
// 	return t.graph.otherSearch
// }
//
// func SetInteger(v int) *int {
// 	return &v
// }
//
// func SetString(v string) *string {
// 	return &v
// }
//
// func SetFloat(v float64) *float64 {
// 	return &v
// }
//
