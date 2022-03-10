package iSolarCloud

import (
	"GoSungrow/Only"
)


func fillArray(count int, args []string) []string {
	var ret []string
	for range Only.Once {
		if count < len(args) {
			count = len(args)
		}
		ret = make([]string, count)
		for i, e := range args {
			ret[i] = e
		}
	}
	return ret
}
