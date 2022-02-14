// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package PowerPointService

import (
	"GoSungro/iSolarCloud/api"
	"reflect"
	"strings"
)


var endpoints = [][]string {
	{"PowerPointService","queryPointDataTopOne","/v1/devService/queryPointDataTopOne"},
}


type AreaStruct struct {
	Name api.AreaName
	EndPoints api.TypeEndPoints
}
var Area AreaStruct

func init() {
	ref := reflect.TypeOf(AreaStruct{})
	p := ref.String()
	s := ref.Name()
	Area.Name = api.AreaName(strings.TrimSuffix(p, "." + s))
	// Area.EndPoints = api.CreateEndPoints(endpoints)
}
