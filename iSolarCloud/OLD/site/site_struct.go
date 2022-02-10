package site

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Site struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=site&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain string `json:"domain" required:"true"` // Name of Domain.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Domains in the specified domain.
}

//const CreateUrl = "?format=json&object=site&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=site&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type DeleteResponse struct {
}

//const ReadUrl = "?format=json&object=site&action=read"
//const ReadBillingUrl = "?format=json&object=site&action=read&billing=yes"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain  string `json:"domain" required:"true"` // Identifies Domain from which to read info by name.
	Site    string `json:"site"`                   // Identifies site from which to read info.
	Billing string `json:"billing:"`               // If set to "yes", then returns billing info.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	CurrentUser             int `json:"current_user"`              // The current user count that the ApiAppKey can have.
	CurrentQueues           int `json:"current_queues"`            // The current call queue count that the ApiAppKey can have.
	CurrentRegisteredDevice int `json:"current_registered_device"` // The current count of registered devices.
	CurrentDevice           int `json:"current_device"`            // The current total device count.
	CurrentCalls            int `json:"current_calls"`             // The current count of active calls in the site
}

//const ListUrl = "?format=json&object=site&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain" required:"true"` // Domain containing Users to list.
}
type ListResponse []ListResponseSingle
type ListResponseSingle string

//const UpdateUrl = "?format=json&object=site&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type UpdateResponse struct {
}
