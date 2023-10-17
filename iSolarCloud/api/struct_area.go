package api

import (
	"errors"
	"fmt"

	"github.com/anicoll/gosungrow/pkg/only"
)

type AreaStruct struct {
	ApiRoot   interface{} // *web.Web
	Name      AreaName
	EndPoints TypeEndPoints
	Error     error
}

func (as AreaStruct) Exists(name EndPointName) error {
	var err error
	for range only.Once {
		if _, ok := as.EndPoints[name]; !ok {
			err = errors.New("unknown endpoint")
			break
		}
	}

	return err
}

func (as *AreaStruct) SortEndPoints() []EndPointName {
	return as.EndPoints.SortEndPoints()
}

func (as *AreaStruct) GetEndPoint(name EndPointName) EndPoint {
	return as.EndPoints.GetEndPoint(name)
}

func (as AreaStruct) ListEndpoints() {
	for range only.Once {
		fmt.Printf("Listing all endpoints from area '%s':\n", as.Name)
		// as.EndPoints.ListEndpoints()
		fmt.Printf("%v", as.EndPoints)
	}
}

func (as AreaStruct) CountEnabled() int {
	return len(as.EndPoints.GetEnabled())
}

func (as AreaStruct) CountDisabled() int {
	return len(as.EndPoints.GetDisabled())
}

func (as AreaStruct) CoveragePercent() float64 {
	d := len(as.EndPoints.GetDisabled())
	e := len(as.EndPoints.GetEnabled())
	return (float64(e) / (float64(e) + float64(d))) * 100
}
