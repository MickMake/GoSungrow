package GoStruct

import (
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


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

// MarshalJSON - Convert value to JSON
func (e EndPointPath) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error
	for range Only.Once {
		// data, err = json.Marshal(e.path)
		// if err!= nil {
		// 	break
		// }
		data = []byte(`"` + e.String() +`"`)
	}

	return data, err
}

func (e *EndPointPath) Clear() {
	e.path = []string{}
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

func (e *EndPointPath) First() string {
	if len(e.path) == 0 {
		return ""
	}
	return e.path[0]
}

func (e *EndPointPath) Last() string {
	l := len(e.path)
	if l == 0 {
		return ""
	}
	return e.path[l-1]
}

func (e *EndPointPath) ReplaceFirst(s string) {
	if len(e.path) > 0 {
		e.path[0] = s
	}
}

func (e *EndPointPath) InsertFirst(s string) {
	p := []string{s}
	p = append(p, e.path...)
	e.path = p
}

func (e *EndPointPath) ShiftLeft(s int64) {
	if s < int64(len(e.path)) {
		e.path = e.path[s:]
	}
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
