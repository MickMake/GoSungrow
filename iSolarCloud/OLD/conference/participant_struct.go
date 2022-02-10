package conference

const ParticipantCountUrl = "?format=json&object=participant&action=count"

type ParticipantCountRequest struct {
}
type ParticipantCountResponse struct {
	Total int `json:"total"` // Total number of Participants for the specified ApiAppKey/User.
}

const ParticipantReadUrl = "?format=json&object=participant&action=read"

type ParticipantReadRequest struct {
}
type ParticipantReadResponse struct {
}
