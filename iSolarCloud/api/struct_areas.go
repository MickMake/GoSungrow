package api

import (
	"GoSungro/Only"
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
	for k := range *an {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	ret := make(AreaNames, 0, len(keys))
	for _, r := range keys {
		ret = append(ret, AreaName(r))
	}
	return ret
}

func (an *Areas) GetEndPoint(area AreaName, name EndPointName) EndPoint {
	var ret EndPoint
	for range Only.Once {
		if _, ok := (*an)[area]; !ok {
			break
		}
		ret = (*an)[area].EndPoints[name]
	}
	return ret
}

func (an Areas) ListAreas() {
	for range Only.Once {
		fmt.Println("Listing all endpoint areas:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Areas", "EndPoints"})
		table.SetBorder(true)
		for _, area := range an.SortAreas() {
			size := fmt.Sprintf("%d", an[area].CountEndpoints())
			table.Append([]string{string(area), size})
		}
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

func (an *Areas) SetRequest(area AreaName, name EndPointName, ref interface{}) error {
	var err error

	for range Only.Once {
		err = an.EndpointExists(area, name)
		if err != nil {
			break
		}

		point := (*an)[area].EndPoints[name]
		point = point.SetRequest(ref)
		err = point.GetError()
	}

	return err
}

func (an *Areas) GetRequest(area AreaName, endpoint EndPointName) Json {
	var ret Json

	for range Only.Once {
		err := an.EndpointExists(area, endpoint)
		if err != nil {
			break
		}
		ret = an.GetEndPoint(area, endpoint).GetRequestJson()
	}

	return ret
}

func (an *Areas) GetResponse(area AreaName, endpoint EndPointName) Json {
	var ret Json

	for range Only.Once {
		err := an.EndpointExists(area, endpoint)
		if err != nil {
			break
		}
		ret = an.GetEndPoint(area, endpoint).GetResponseJson()
	}

	return ret
}
