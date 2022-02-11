package api

import (
	"GoSungro/Only"
	"errors"
	"sort"
)


type TypeEndPoints map[EndPointName]TypeEndPoint	// Map of EndPoints by endpoint name.
// type EndPointNames []EndPointName


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

func (ps *TypeEndPoints) SortEndPoints() []string {
	keys := make([]string, 0, len(*ps))
	for k := range *ps {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return keys
}

func (ps *TypeEndPoints) GetEndPoint(name EndPointName) *TypeEndPoint {
	ret, _ := (*ps)[name]
	return &ret
}
