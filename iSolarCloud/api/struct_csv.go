package api

import (
	"GoSungrow/Only"
	"errors"
	"fmt"
	"os"
)

type Csv struct {
	Data   [][]string
	Header []string
	Error  error
}

func NewCsv() Csv {
	return Csv{}
}

func (c Csv) String() string {
	var ret string
	ret += c.HeaderString()
	ret += c.DataString()
	return ret
}

func (c Csv) Print() {
	fmt.Println(c)
}

func (c Csv) HeaderString() string {
	var ret string
	for _, h := range c.Header {
		ret += fmt.Sprintf("\"%s\",", h)
	}
	ret += fmt.Sprintln()
	return ret
}

func (c Csv) DataString() string {
	var ret string
	for _, r := range c.Data {
		for _, r := range r {
			ret += fmt.Sprintf("\"%s\",", r)
		}
		ret += fmt.Sprintln()
	}
	return ret
}

func (c Csv) AddRow(row []string) Csv {
	c.Data = append(c.Data, row)
	return c
}

func (c Csv) SetHeader(header []string) Csv {
	c.Header = header
	return c
}

func (c *Csv) WriteFile(fn string, perm os.FileMode) error {
	for range Only.Once {
		fmt.Printf("Writing file '%s'\n", fn)
		c.Error = os.WriteFile(fn, []byte(c.String()), perm)
		if c.Error != nil {
			c.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, c.Error))
			break
		}
	}

	return c.Error
}
