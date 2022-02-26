package output


// type Table struct {
// 	Data   [][]interface{}
// 	Header []string
// 	Error  error
// }
//
//
// func NewTable() Table {
// 	return Table{}
// }
//
// func (c Table) String() string {
// 	var ret string
// 	ret += c.HeaderString()
// 	ret += c.DataString()
// 	return ret
// }
//
// func (c Table) Print() {
// 	fmt.Println(c)
// }
//
// func (c Table) HeaderString() string {
// 	var ret string
// 	for _, h := range c.Header {
// 		ret += fmt.Sprintf("\"%s\",", h)
// 	}
// 	ret += fmt.Sprintln()
// 	return ret
// }
//
// func (c Table) DataString() string {
// 	var ret string
// 	for _, r := range c.Data {
// 		for _, r := range r {
// 			ret += fmt.Sprintf("\"%s\",", r)
// 		}
// 		ret += fmt.Sprintln()
// 	}
// 	return ret
// }
//
// func (c Table) AddRow(row []string) Table {
// 	foo := Cell{
// 		string: "",
// 		float64:      0,
// 	}
// 	c.Data = append(c.Data, foo)
// 	for i := range row {
//
// 	}
// 	return c
// }
//
// func (c Table) SetHeader(header []string) Table {
// 	c.Header = header
// 	return c
// }
//
// func (c *Table) WriteFile(fn string, perm os.FileMode) error {
// 	for range Only.Once {
// 		fmt.Printf("Writing file '%s'\n", fn)
// 		c.Error = os.WriteFile(fn, []byte(c.String()), perm)
// 		if c.Error != nil {
// 			c.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, c.Error))
// 			break
// 		}
// 	}
//
// 	return c.Error
// }
