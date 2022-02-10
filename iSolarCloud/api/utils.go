package api

import (
	"reflect"
	"strings"
)

var OnlyOnce = "1"

func MakeNewFrom(i interface{}) (new interface{}) {
	for range OnlyOnce {
		t := reflect.TypeOf(i)
		if t == nil {
			break
		}
		new = reflect.New(t.Elem()).Interface()
	}
	return new
}

func ToSnakeCase(s string) string {
	slen := len(s)
	b := []byte(s)
	pos := make([]int, 0)
	for i := 1; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			pos = append(pos, i)
			slen++
		}
	}
	if len(pos) > 0 {
		sc := make([]byte, slen)
		n := 0
		p := 0
		for i := 0; i < len(b); i++ {
			if pos[p] == i {
				sc[n] = '_'
				n++
			}
			sc[n] = b[i]
			n++
		}
		s = string(sc)
	}
	return strings.ToLower(s)
}
