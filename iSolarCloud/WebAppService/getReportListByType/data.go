package getReportListByType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/getReportListByType"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData []struct {
	CreateTime valueTypes.DateTime `json:"create_time"`
	Cycle      valueTypes.Integer  `json:"cycle"`
	ID         valueTypes.Integer  `json:"id"`
	MonthDate  valueTypes.String   `json:"month_date"`
	PsID       valueTypes.PsId     `json:"ps_id"`
	ReportName valueTypes.String   `json:"report_name"`
	ReportType valueTypes.Integer  `json:"report_type"`
	Status     valueTypes.Bool     `json:"status"`
	UpdateTime valueTypes.DateTime `json:"update_time"`
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
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.DateId.Format(valueTypes.DateTimeLayoutDay)
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
