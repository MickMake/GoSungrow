package conference

const MeetingCountUrl = "?format=json&object=meeting&action=count"

type MeetingCountRequest struct {
}
type MeetingCountResponse struct {
	Total int `json:"total"` // Total number of Meetings for the specified ApiAppKey/User.
}

const MeetingReadUrl = "?format=json&object=meeting&action=read"

type MeetingReadRequest struct {
}
type MeetingReadResponse struct {
}
