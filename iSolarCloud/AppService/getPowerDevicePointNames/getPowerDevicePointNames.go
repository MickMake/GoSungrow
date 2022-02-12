// Resource
package getPowerDevicePointNames

import (
	"GoSungro/iSolarCloud/api"
	"fmt"
)

var Url = "/v1/reportService/getPowerDevicePointNames"

var _ api.Resource = (*Resource)(nil)

type Resource struct {
	api.TypeEndPoint
	// api.Common
	// Request
	// Response
}

type Request struct {
	api.RequestCommon
	Extra string
}
type Response struct {
	api.ResponseCommon
	Another string
}

func (g Resource) Call() api.Json {
	fmt.Println("Call() implement me")
	return ""
}

func (g Resource) SetRequest(ref interface{}) error {
	fmt.Println("Call() implement me")
	fmt.Printf("ref == %v\n", ref)
	return nil
}

func (g Resource) GetResource() interface{} {
	fmt.Println("Call() implement me")
	return &Resource{}
}

func (g Resource) Init() *Resource {
	fmt.Println("Init()")
	return &Resource{}
}

func Init() *Resource {
	fmt.Println("Init()")

	foo := Resource {
		api.TypeEndPoint {
			Area:     api.GetArea(Resource{}),
			Name:     api.GetEndPoint(Resource{}),
			Url:      api.GetUrl(Url),
			// Resource: nil,
			Request:  Request{},
			Response: Response{},
			// Get:      nil,
			// Put:      nil,
			Error:    nil,
		},
	}
	fmt.Printf("endpoint: %v\n", foo)

	return &foo
}
