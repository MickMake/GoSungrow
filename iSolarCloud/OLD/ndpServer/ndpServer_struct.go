package ndpServer

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Ndp struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

// const CountUrl = "?format=json&object=ndpserver&action=count"

type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type CountResponseSingle struct {
	Total int `json:"total"` // Total number of Devices for the specified ApiAppKey/User.
}
type CountResponse []CountResponseSingle

// const CreateUrl = "?format=json&object=ndpserver&action=create"

type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type CreateResponse struct {
}

// const DeleteUrl = "?format=json&object=ndpserver&action=delete"

type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type DeleteResponse struct {
}

// const ListUrl = "?format=json&object=ndpserver&action=list"

type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
}

// const ReadUrl = "?format=json&object=ndpserver&action=read"

type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
}

// const UpdateUrl = "?format=json&object=ndpserver&action=update"

type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Server string `json:"server"` // The SipBx server name to query against. Also accepts:
	// - 'lookup' which will reference the API configuration 'NsNdpSiSunGroServer'
	// - 'readconfig' which will reference the API database record for 'NdpDomain.
}
type UpdateResponse struct {
}
