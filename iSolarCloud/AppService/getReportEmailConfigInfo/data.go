package getReportEmailConfigInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/reportService/getReportEmailConfigInfo"
const Disabled = false
const EndPointName = "AppService.getReportEmailConfigInfo"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	ReportEmailConfigInfoList []struct {
		Email      valueTypes.String `json:"email"`
		ReportList []struct {
			CreateTime                 valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006/01/02 15:04:05"`
			CreateUserId               valueTypes.Integer  `json:"create_user_id"`
			EmailAddTime               valueTypes.DateTime `json:"email_add_time" PointNameDateFormat:"2006/01/02 15:04:05"`
			Id                         valueTypes.Integer  `json:"id"`
			IsAllowEmailSend           valueTypes.Bool     `json:"is_allow_email_send"`
			IsBank                     valueTypes.Bool     `json:"is_bank"`
			IsCanRenewSendConfirmEmail valueTypes.Bool     `json:"is_can_renew_send_confirm_email"`
			IsNewWeb                   valueTypes.Bool     `json:"is_new_web"`
			OrderId                    valueTypes.Integer  `json:"order_id"`
			ReSendConfirmEmailTime     valueTypes.DateTime `json:"re_send_confirm_email_time" PointNameDateFormat:"2006/01/02 15:04:05"`
			ReportId                   valueTypes.Integer  `json:"report_id"`
			ReportName                 valueTypes.String   `json:"report_name"`
			SendEmail                  valueTypes.String   `json:"send_email"`
			Status                     valueTypes.Bool     `json:"status"`
			TimeDimension              valueTypes.Integer  `json:"time_dimension"`
			Type                       valueTypes.Integer  `json:"type"`
			UpdateTime                 valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
			UserId                     valueTypes.Integer  `json:"user_id"`
		} `json:"report_list" DataTable:"true"`
	} `json:"report_email_config_info_list"`
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

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
