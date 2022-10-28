package getRemoteUpgradeTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getRemoteUpgradeTaskList"
const Disabled = false

type RequestData struct {
	PsIdList valueTypes.String `json:"ps_id_list" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	// ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	// Dummy string `json:"dummy"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, pkg, "system", dt)
	}

	return entries
}
