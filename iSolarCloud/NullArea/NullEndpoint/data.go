package NullEndpoint

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)


const Url = "%URL%"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

// IsValid Checks for validity of results data.
func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

// Help provides more info to the user on request JSON fields.
func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

// ResultData holds data returned from the API.
type ResultData struct {
	Dummy valueTypes.String `json:"dummy"`
}

// IsValid Checks for validity of results data.
func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
