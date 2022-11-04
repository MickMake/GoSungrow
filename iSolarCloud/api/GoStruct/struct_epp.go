package GoStruct

import "strings"


type EndPointPath []string

func NewEndPointPath(path ...string) EndPointPath {
	var epp EndPointPath
	return epp.Append(path...)
}

func (e *EndPointPath) Copy() EndPointPath {
	ret := make(EndPointPath, len(*e))
	copy(ret, *e)
	return ret
}

func (e *EndPointPath) Append(path ...string) EndPointPath {
	ret := make(EndPointPath, len(*e))
	copy(ret, *e)
	for _, p := range path {
		ret = append(ret, p)
	}
	return ret
}

func (e *EndPointPath) PopLast() EndPointPath {
	if len(*e) == 0 {
		return *e
	}
	return (*e)[:len(*e) - 1]
}

func (e EndPointPath) String() string {
	return strings.Join(e, ".")
}
