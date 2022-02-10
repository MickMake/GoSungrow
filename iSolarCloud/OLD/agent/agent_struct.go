package agent

//goland:noinspection GoNameStartsWithPackageName

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Agent struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=agent&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Devices for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=agent&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type CreateResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const DeleteUrl = "?format=json&object=agent&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type DeleteResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const ListUrl = "?format=json&object=agent&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const ReadUrl = "?format=json&object=agent&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const UpdateUrl = "?format=json&object=device&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain string `json:"domain" required:"true"` //
	Device string `json:"device"`                 //
}
type UpdateResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}
