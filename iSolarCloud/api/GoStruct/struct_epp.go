package GoStruct

import "strings"


type EndPointPath struct {
	path []string
}

func NewEndPointPath(path ...string) EndPointPath {
	var epp EndPointPath
	epp.Append(path...)
	return epp
}

func (e EndPointPath) String() string {
	return strings.Join(e.path, ".")
}

func (e *EndPointPath) Strings() []string {
	return e.path
}

func (e *EndPointPath) Index(index int) string {
	if index > len(e.path) {
		return ""
	}
	return e.path[index]
}

func (e *EndPointPath) IsZero() bool {
	if len(e.path) == 0 {
		return true
	}
	return false
}

func (e *EndPointPath) Copy() EndPointPath {
	ret := make([]string, len(e.path))
	copy(ret, e.path)
	return EndPointPath { path:ret }
}

func (e *EndPointPath) Append(path ...string) {
	// ret := make([]string, len(e.path))
	// copy(ret, e.path)
	// for _, p := range path {
	// 	ret = append(ret, p)
	// }
	// return EndPointPath { path:ret }
	e.path = append(e.path, path...)
}

func (e *EndPointPath) PopLast() {
	if len(e.path) == 0 {
		return
		// return *e
	}
	e.path = e.path[:len(e.path) - 1]
	return
	// return (*e)[:len(*e) - 1]
}
