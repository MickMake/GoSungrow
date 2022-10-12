package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
)


type Areas map[AreaName]AreaStruct // TypeEndPoints		// Map of EndPoints by area name.
type AreaName string
type AreaNames []AreaName

func (an *Areas) Exists(area string) bool {
	var ok bool
	_, ok = (*an)[AreaName(area)]
	return ok
}
func (an *Areas) NotExists(area string) bool {
	return !an.Exists(area)
}

func (an *Areas) EndpointExists(area AreaName, name EndPointName) error {
	var err error
	for range Only.Once {
		if _, ok := (*an)[area]; !ok {
			err = errors.New("unknown endpoint area")
			break
		}
		if err = (*an)[area].Exists(name); err != nil {
			break
		}
	}
	return err
}

func (an *Areas) SortAreas() AreaNames {
	keys := make([]string, 0, len(*an))
	for _, k := range *an {
		keys = append(keys, string(k.Name))
	}
	sort.Strings(keys)
	ret := make(AreaNames, 0, len(keys))
	for _, r := range keys {
		ret = append(ret, AreaName(r))
	}
	return ret
}

func (an *Areas) GetArea(area AreaName) *Area {
	var ret *Area
	for range Only.Once {
		if _, ok := (*an)[area]; !ok {
			break
		}
	}
	return ret
}

func (an *Areas) GetEndPoint(area AreaName, endpoint EndPointName) EndPoint {
	var ret EndPoint
	for range Only.Once {
		if area == "" {
			ret = ret.SetError("empty endpoint area name")
			break
		}

		if endpoint == "" {
			ret = ret.SetError("empty endpoint name")
			break
		}

		// err := an.EndpointExists(area, endpoint)
		// if err != nil {
		// 	fmt.Printf("ERROR: %s\n", err)
		// 	break
		// }

		if _, ok := (*an)[area]; !ok {
			ret = (*an)[NullAreaName].EndPoints[NullEndPointName]
			ret = ret.SetError("unknown endpoint area '%s'", area)
			break
		}

		if _, ok := (*an)[area].EndPoints[endpoint]; !ok {
			ret = (*an)[NullAreaName].EndPoints[NullEndPointName]
			ret = ret.SetError("unknown endpoint '%s.%s'", area, endpoint)
			break
		}

		ret = (*an)[area].EndPoints[endpoint]

		if ret.IsDisabled() {
			ret.SetError("API EndPoint is not implemented")
			break
		}

	}
	return ret
}

func (an Areas) ListAreas() {
	for range Only.Once {
		fmt.Println("Listing all endpoint areas:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Areas", "Enabled EndPoints", "Disabled EndPoints", "Coverage %"})
		table.SetBorder(true)
		te := 0
		td := 0
		for _, area := range an.SortAreas() {
			e := an[area].CountEnabled()
			d := an[area].CountDisabled()
			p := (float64(e) / float64(d)) * 100
			table.Append([]string{
				string(area),
				fmt.Sprintf("%d", e),
				fmt.Sprintf("%d", d),
				fmt.Sprintf("%.1f %%", p),
			})
			te += e
			td += d
		}

		table.Append([]string{"----------------", "----------------", "-----------------", "---------"})

		p := (float64(te) / float64(td)) * 100
		table.Append([]string{
			"Total",
			fmt.Sprintf("%d", te),
			fmt.Sprintf("%d", td),
			fmt.Sprintf("%.1f %%", p),
		})
		table.Render()
	}
}

func (an Areas) ListEndpoints(area string) error {
	var err error
	for range Only.Once {
		if area == "" {
			fmt.Printf("Listing all areas:\n")
			for _, a := range an.SortAreas() {
				an[a].ListEndpoints()
			}
			break
		}

		if an.NotExists(area) {
			err = errors.New("unknown area name")
			break
		}

		an[AreaName(area)].ListEndpoints()
	}

	return err
}

func (an *Areas) SetRequest(area AreaName, endpoint EndPointName, ref interface{}) error {
	var err error

	for range Only.Once {
		ep := an.GetEndPoint(area, endpoint)
		if ep.IsError() {
			err = ep.GetError()
			break
		}

		ep = ep.SetRequest(ref)
		if ep.IsError() {
			err = ep.GetError()
			break
		}

		// point := (*an)[area].EndPoints[endpoint]
		// point = point.SetRequest(ref)
		// err = point.GetError()
	}

	return err
	// return an.GetEndPoint(area, endpoint).SetRequest()
}

func (an *Areas) GetRequest(area AreaName, endpoint EndPointName) output.Json {
	return an.GetEndPoint(area, endpoint).GetRequestJson()
}

func (an *Areas) GetResponse(area AreaName, endpoint EndPointName) output.Json {
	return an.GetEndPoint(area, endpoint).GetResponseJson()
}
