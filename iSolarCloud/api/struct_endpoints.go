package api

import (
	"GoSungrow/Only"
	"bytes"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"sort"
)


type TypeEndPoints map[EndPointName]EndPoint // Map of EndPoints by endpoint name.

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

func (ps TypeEndPoints) String() string {
	var ret string
	for range Only.Once {
		buf := new(bytes.Buffer)

		ep := ps.GetEnabled()
		if len(ep) > 0 {
			table := tablewriter.NewWriter(buf)
			table.SetHeader([]string{"EndPoint", "Url"})
			table.SetBorder(true)
			for _, endpoint := range ep.SortEndPoints() {
				u := ep[endpoint].GetUrl().String()
				table.Append([]string{
					string(ep[endpoint].GetName()),
					u,
				})
			}
			table.Render()
			ret += fmt.Sprintf("# Implemented endpoints\n")
			ret += buf.String()
		}

		dp := ps.GetDisabled()
		if len(dp) > 0 {
			buf = new(bytes.Buffer)
			table := tablewriter.NewWriter(buf)
			table.SetHeader([]string{"EndPoint", "Url"})
			table.SetBorder(true)
			for _, endpoint := range dp.SortEndPoints() {
				u := dp[endpoint].GetUrl().String()
				table.Append([]string{
					string(dp[endpoint].GetName()),
					u,
				})
			}
			table.Render()
			ret += fmt.Sprintf("# Non-implemented endpoints\n")
			ret += buf.String()
		}

		ret += fmt.Sprintf("\n# Implemented: %d Non-implemented: %d\n", len(ep), len(dp))

	}
	return ret
}

// func (ps TypeEndPoints) ListEndpoints() {
// 	fmt.Printf("%v", ps)
// }

func (ps *TypeEndPoints) SortEndPoints() []EndPointName {
	ret := make([]EndPointName, 0, len(*ps))
	for range Only.Once {
		keys := make([]string, 0, len(*ps))
		for k := range *ps {
			keys = append(keys, string(k))
		}
		sort.Strings(keys)

		for _, k := range keys {
			ret = append(ret, EndPointName(k))
		}
	}
	return ret
}

// func (ps *TypeEndPoints) CountEndpoints() int {
//	return len(*ps)
// }

func (ps *TypeEndPoints) GetEnabled() TypeEndPoints {
	ret := make(TypeEndPoints)
	for range Only.Once {
		s := ps.SortEndPoints()
		if len(s) == 0 {
			break
		}

		for _, endpoint := range s {
			if (*ps)[endpoint].IsDisabled() {
				continue
			}
			ret[(*ps)[endpoint].GetName()] = (*ps)[endpoint]
		}
	}
	return ret
}

func (ps *TypeEndPoints) GetDisabled() TypeEndPoints {
	ret := make(TypeEndPoints)
	for range Only.Once {
		s := ps.SortEndPoints()
		if len(s) == 0 {
			break
		}

		for _, endpoint := range s {
			if (*ps)[endpoint].IsDisabled() {
				ret[(*ps)[endpoint].GetName()] = (*ps)[endpoint]
			}
		}
	}
	return ret
}

func (ps *TypeEndPoints) GetEndPoint(name EndPointName) EndPoint {
	ret, _ := (*ps)[name]
	return ret
}
