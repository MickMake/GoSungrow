package deviceModel

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Model struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

const CountUrl = "?format=json&object=devicemodel&action=count"

type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
}
type CountResponseSingle struct {
	Total int `json:"total"` // Total number of DeviceModels for the specified ApiAppKey/User.
}
type CountResponse []CountResponseSingle

const CreateUrl = "?format=json&object=devicemodel&action=create"

type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Brand          string `json:"brand"`            // Search for device by brand
	Model          string `json:"model"`            // Search for device by model
	PortalView     string `json:"portal_view"`      // Set to "yes" for a more limited list
	IncludeNdpDefs bool   `json:"include_ndp_defs"` // Include NDP-defined device features
	NdpSyntax      string `json:"ndp_syntax"`       // Search by the NDP code structure / syntax
}
type CreateResponse struct {
}

const DeleteUrl = "?format=json&object=devicemodel&action=delete"

type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
}
type DeleteResponse struct {
}

const ListUrl = "?format=json&object=devicemodel&action=list"

type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Brand          string `json:"brand"`            // Search for device by brand
	Model          string `json:"model"`            // Search for device by model
	PortalView     string `json:"portal_view"`      // Set to "yes" for a more limited list
	IncludeNdpDefs bool   `json:"include_ndp_defs"` // Include NDP-defined device features
	NdpSyntax      string `json:"ndp_syntax"`       // Search by the NDP code structure / syntax
}
type ListResponseSingle struct {
	PortalView string `json:"portal_view"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	DeviceType string `json:"device_type"`
	NdpSyntax  string `json:"ndp_syntax"`
	BrandModel string `json:"brand_model"`
}
type ListResponse []ListResponseSingle

// [
//    {
//        "portal_view": "yes",
//        "brand": "Aastra",
//        "model": "480i",
//        "device_type": "Device",
//        "ndp_syntax": "Aastra",
//        "brand_model": "Aastra 480i"
//    }
// ]
const ReadUrl = "?format=json&object=devicemodel&action=read"

type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Brand          string `json:"brand"`            // Search for device by brand
	Model          string `json:"model"`            // Search for device by model
	PortalView     string `json:"portal_view"`      // Set to "yes" for a more limited list
	IncludeNdpDefs bool   `json:"include_ndp_defs"` // Include NDP-defined device features
	NdpSyntax      string `json:"ndp_syntax"`       // Search by the NDP code structure / syntax
}
type ReadResponseSingle struct {
	PortalView  string      `json:"portal_view"`
	Brand       string      `json:"brand"`
	Model       string      `json:"model"`
	DeviceType  string      `json:"device_type"`
	NdpSyntax   string      `json:"ndp_syntax"`
	BrandModel  string      `json:"brand_model"`
	ButtonTypes ButtonTypes `json:"button_types"`
}
type ReadResponse []ReadResponseSingle

const UpdateUrl = "?format=json&object=devicemodel&action=update"

type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
}
type UpdateResponse struct {
}

// Additional discovered entities.
//type ButtonTypes []ButtonType
type ButtonTypes []struct {
	ButtonType string   `json:"button_type"`
	Features   Features `json:"features"`
	Name       string   `json:"name"`
}

//type Features []Feature
type Features []struct {
	Autocomplete  string        `json:"autocomplete"`
	ExtraSettings ExtraSettings `json:"extra_settings"`
	FeatureType   string        `json:"feature_type"`
	Name          string        `json:"name"`
	Notarget      bool          `json:"notarget"`
}

//type ExtraSettings []ExtraSetting
type ExtraSettings []struct {
	Label string `json:"Label"`
	Type  string `json:"type"`
}
