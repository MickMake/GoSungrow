package getReportInfoByReportId

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/getReportInfoByReportId"
const Disabled = false

type RequestData struct {
	ReportId valueTypes.String `json:"report_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	CreateUserID         valueTypes.Integer `json:"create_user_id"`
	ID                   valueTypes.Integer `json:"id"`
	ReportID             valueTypes.Integer `json:"report_id"`
	ReportName           valueTypes.String  `json:"report_name"`
	ReportPsType         valueTypes.Integer `json:"report_ps_type"`
	ReportTemplatePeriod valueTypes.Integer `json:"report_template_period"`
	ReportTemplateType   valueTypes.Integer `json:"report_template_type"`
	PsIDList             interface{}        `json:"ps_id_list"`
	SamplingPeriod       interface{}        `json:"sampling_period"`
	TimeDimension        valueTypes.Integer `json:"time_dimension"`
	UserIds              valueTypes.Integer `json:"user_ids"`
	PointColumn          valueTypes.String  `json:"point_column" PointSplitOn:","`
	PointName            valueTypes.String  `json:"point_name" PointSplitOn:","`
	Ratio                valueTypes.String  `json:"ratio" PointSplitOn:","`
	ReportTemplatePoint  valueTypes.String  `json:"report_template_point" PointSplitOn:","`
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
		// pkg := apiReflect.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.ReportId.String()
		entries.StructToDataMap(*e, e.Request.ReportId.String(), apiReflect.NewEndPointPath(e.Request.ReportId.String()))
	}

	return entries
}
