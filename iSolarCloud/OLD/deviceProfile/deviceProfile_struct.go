package deviceProfile

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Profile struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

const CountUrl = "?format=json&object=deviceprofile&action=count"

type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Territory string `json:"territory"` // The name of the device profile
}
type CountResponse struct {
	Total int `json:"total"` // Total number of DeviceProfiles for the specified ApiAppKey/User.
}

const CreateUrl = "?format=json&object=deviceprofile&action=create"

type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Territory   string `json:"territory"`   // The name of the device profile
	Description string `json:"description"` // Description for the device profile.
}
type CreateResponse struct {
}

const DeleteUrl = "?format=json&object=deviceprofile&action=delete"

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

const ListUrl = "?format=json&object=deviceprofile&action=list"

type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Brand          string `json:"brand"`            // Search for device by brand
	Profile        string `json:"model"`            // Search for device by model
	PortalView     string `json:"portal_view"`      // Set to "yes" for a more limited list
	IncludeNdpDefs bool   `json:"include_ndp_defs"` // Include NDP-defined device features
	NdpSyntax      string `json:"ndp_syntax"`       // Search by the NDP code structure / syntax
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
}

// [
//    {
//        "brand": "Avaya",
//        "model": "J100 Expansion Module",
//        "device_type": "sidecar",
//        "device_image_filetype": "image/png",
//        "device_image_filesize": "73630",
//        "device_image_height": "2480",
//        "device_image_width": "3508",
//        "device_image_default_position": "{\"x\":-179,\"y\":-115,\"zoom\":1.5999999999999999}",
//        "button_page_count": "3",
//        "sidecar_support": ""
//    }
// ]
const ReadUrl = "?format=json&object=deviceprofile&action=read"

type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Territory  string `json:"territory"`   //
	DeviceType string `json:"device_type"` //
	Brand      string `json:"brand"`       //
	Model      string `json:"model"`       //
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Brand                      string `json:"brand"`
	ButtonPageCount            string `json:"button_page_count"`
	DeviceImageDefaultPosition string `json:"device_image_default_position"`
	DeviceImageFilesize        string `json:"device_image_filesize"`
	DeviceImageFiletype        string `json:"device_image_filetype"`
	DeviceImageHeight          string `json:"device_image_height"`
	DeviceImageWidth           string `json:"device_image_width"`
	DeviceType                 string `json:"device_type"`
	Model                      string `json:"model"`
	SidecarSupport             string `json:"sidecar_support"`
}

const UpdateUrl = "?format=json&object=deviceprofile&action=update"

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
