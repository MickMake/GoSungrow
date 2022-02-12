package api

import (
	"GoSungro/Only"
	"errors"
	"sort"
)


type TypeAreaNames map[AreaName]TypeEndPoints		// Map of EndPoints by area name.
type AreaName string


func (an *TypeAreaNames) Exists(area AreaName, name EndPointName) error {
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

func (an *TypeAreaNames) SortAreas() []AreaName {
	keys := make([]string, 0, len(*an))
	for k := range *an {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	ret := make([]AreaName, 0, len(keys))
	for _, r := range keys {
		ret = append(ret, AreaName(r))
	}
	return ret
}

func (an *TypeAreaNames) GetEndPoint(area AreaName, name EndPointName) *TypeEndPoint {
	var ret *TypeEndPoint
	for range Only.Once {
		if _, ok := (*an)[area]; !ok {
			break
		}
		point := (*an)[area]
		ret = point.GetEndPoint(name)
	}
	return ret
}

func (an *TypeAreaNames) SetFuncPut(area AreaName, endpoint EndPointName, fn SetFunc) error {
	var err error
	for range Only.Once {
		err = an.Exists(area, endpoint)
		if err != nil {
			break
		}
		point := (*an)[area][endpoint]
		err = point.SetFuncPut(fn)
	}
	return err
}

func (an *TypeAreaNames) SetFuncGet(area AreaName, endpoint EndPointName, fn GetFunc) error {
	var err error
	for range Only.Once {
		err = an.Exists(area, endpoint)
		if err != nil {
			break
		}
		point := (*an)[area][endpoint]
		err = point.SetFuncGet(fn)
	}
	return err
}

func (an *TypeAreaNames) SetRequest(area AreaName, endpoint EndPointName, ref interface{}) error {
	var err error

	for range Only.Once {
		err = an.Exists(area, endpoint)
		if err != nil {
			break
		}

		point := (*an)[area][endpoint]
		err = point.SetRequest(ref)
	}

	return err
}

func (an *TypeAreaNames) SetResponse(area AreaName, endpoint EndPointName, ref interface{}) error {
	var err error

	for range Only.Once {
		err = an.Exists(area, endpoint)
		if err != nil {
			break
		}

		point := (*an)[area][endpoint]
		err = point.SetResponse(ref)
	}

	return err
}
