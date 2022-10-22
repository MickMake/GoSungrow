package findPsType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/findPsType"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	PsType    valueTypes.Integer `json:"ps_type"`
	SysScheme valueTypes.Integer `json:"sys_scheme"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		entries.StructToPoints(e.Response.ResultData, pkg, e.Request.PsId.String(), dt)

		// for _, d := range e.Response.ResultData {
		// 	name := fmt.Sprintf("%s.%s", pkg, e.Request.PsId.String())
		// 	entries.StructToPoints(d, name, e.Request.PsId.String(), dt)
		// }
	}

	return entries
}
