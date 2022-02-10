package chart

//goland:noinspection GoNameStartsWithPackageName

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Chart struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=chart&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain      string `json:"domain" required:"true"` // Domain from which to read charts.
	DashboardId int    `json:"dashboard_id"`           // Dashboard id of the parent dashboard
	Type        string `json:"type"`                   // Chart type. Must be a component supported by the netsapiens-charts-library package
	Format      string `json:"format"`                 //
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Devices for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=chart&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain      string `json:"domain" required:"true"`       // Domain from which to read charts.
	Name        string `json:"name" required:"true"`         // Name/title of the chart
	DashboardId int    `json:"dashboard_id" required:"true"` // Dashboard id of the parent dashboard
	Type        string `json:"type" required:"true"`         // Chart type. Must be a component supported by the netsapiens-charts-library package
	Description string `json:"description"`                  // properties	string	// JSON object which describes visual and functional properties of the chart
	DataSources string `json:"data_sources"`                 // JSON object which describes the sources of the data to be visualized
	Status      string `json:"status"`                       //
}
type CreateResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const DeleteUrl = "?format=json&object=chart&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Id     int    `json:"id" required:"true"`     // Unique ID of the chart
	Domain string `json:"domain" required:"true"` // Domain from which to read charts.
}
type DeleteResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}

//const ListUrl = "?format=json&object=chart&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain      string `json:"domain" required:"true"` // Domain from which to read charts.
	DashboardId int    `json:"dashboard_id"`           // The unique id of the parent dashboard
	Format      string `json:"format"`                 //
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	Id   int    `json:"id"`   // Unique ID of the chart
	Name string `json:"name"` // User provided name/title of the chart
	Type string `json:"type"` //
}

//const ReadUrl = "?format=json&object=chart&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	DashboardId int    `json:"dashboard_id" required:"true"` // Dashboard from which to read charts.
	Domain      string `json:"domain" required:"true"`       // Domain from which to read charts.
	Id          int    `json:"id"`                           // Read a specific chart using the unique ID
	Name        string `json:"name"`                         // Read a specific chart by the name. Name is not inherently unique
	Format      string `json:"format"`                       //
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Id          int    `json:"id"`           // Unique ID of the chart
	DashboardId int    `json:"dashboard_id"` // Unique ID of the parent dashboard
	Name        string `json:"name"`         // User provided name/title of the chart
	Type        string `json:"type"`         //
	Description string `json:"description"`  //
	Properties  string `json:"properties"`   // JSON object which describes visual and functional properties of the chart
	DataSources string `json:"data_sources"` // JSON object which describes the sources of the data to be visualized
	Status      int    `json:"status"`       //
	Created     string `json:"created"`      //
	Modified    string `json:"modified"`     //
	DataToken   string `json:"data_token"`   // A specially generated oauthtoken used to request data for the chart
}

//const UpdateUrl = "?format=json&object=device&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Id          int    `json:"id" required:"true"`           // Read a specific chart using the unique ID
	Domain      string `json:"domain" required:"true"`       // Domain from which to read charts.
	DashboardId int    `json:"dashboard_id" required:"true"` // The unique id of the parent dashboard
	Name        string `json:"name"`                         // Name/title of the chart
	Description string `json:"description"`                  //
	Properties  string `json:"properties"`                   // JSON object which describes visual and functional properties of the chart
	DataSources string `json:"data_sources"`                 // JSON object which describes the sources of the data to be visualized
	Status      string `json:"status"`                       //
}
type UpdateResponse struct {
	Domain string `json:"domain"` //
	Device string `json:"device"` //
}
