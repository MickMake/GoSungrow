package api

import (
	"GoSungrow/Only"
	"errors"
	"fmt"
)

type AreaStruct struct {
	ApiRoot   interface{} // *web.Web
	Name      AreaName
	EndPoints TypeEndPoints
	Error     error
}

func (as AreaStruct) Exists(name EndPointName) error {
	var err error
	for range Only.Once {
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
	for range Only.Once {
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
