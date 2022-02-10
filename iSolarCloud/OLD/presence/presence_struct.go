package presence

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Presence struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=presence&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain string `json:"domain" required:"true"` // Identifies Domain from which to list Presences
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Presences for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=presence&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=presence&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	User   string `json:"user"`   //
}
type DeleteResponse struct {
}

//const ReadUrl = "?format=json&object=presence&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain string `json:"domain" required:"true"` // Identifies Domain from which to list Presences
}
type ReadResponse struct {
}

//const ListUrl = "?format=json&object=presence&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain     string `json:"domain" required:"true"` // Identifies Domain from which to list Presences
	LastUpdate string `json:"last_update"`            //
	Limit      string `json:"limit"`                  // Sets limit for number of Presences returned.
	Department string `json:"department"`             // Identifies Department from which to list Presences
	Order      string `json:"order"`                  //
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	FirstName string `json:"first_name"` //
	LastName  string `json:"last_name"`  //
	Group     string `json:"group"`      //
	User      string `json:"user"`       //
	Uid       string `json:"uid"`        //
	Domain    string `json:"domain"`     //
	Presence  string `json:"presence"`   //
	Site      string `json:"site"`       //
	Message   string `json:"message"`    //
	Email     string `json:"email"`      //
}

//const UpdateUrl = "?format=json&object=presence&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
}
type UpdateResponse struct {
}
