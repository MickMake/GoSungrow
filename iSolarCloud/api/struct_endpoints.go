package api

import (
	"GoSungro/Only"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
)


// type TypeEndPoints map[EndPointName]EndPointStruct	// Map of EndPoints by endpoint name.
type TypeEndPoints map[EndPointName]EndPoint	// Map of EndPoints by endpoint name.


func (ps TypeEndPoints) Exists(name EndPointName) error {
	var err error
	for range Only.Once {
		if _, ok := ps[name]; !ok {
			err = errors.New("unknown endpoint")
			break
		}
	}

	return err
}

func (ps *TypeEndPoints) SortEndPoints() []EndPointName {
	keys := make([]string, 0, len(*ps))
	for k := range *ps {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	ret := make([]EndPointName, 0, len(*ps))
	for k := range keys {
		keys = append(keys, string(k))
	}
	return ret
}

func (ps TypeEndPoints) ListEndpoints() {
	for range Only.Once {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"EndPoint", "Url"})
		table.SetBorder(true)
		for _, endpoint := range ps.SortEndPoints() {
			table.Append([]string{
				string(ps[endpoint].GetName()),
				ps[endpoint].GetUrl().String(),
			})
		}
		table.Render()
	}
}

func (ps *TypeEndPoints) CountEndpoints() int {
	return len(*ps)
}

func (ps *TypeEndPoints) GetEndPoint(name EndPointName) EndPoint {
	ret, _ := (*ps)[name]
	fmt.Printf("ok:%v\tref:%v\n")
	return ret
}
