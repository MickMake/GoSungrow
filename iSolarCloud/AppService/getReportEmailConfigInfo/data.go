package getReportEmailConfigInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/reportService/getReportEmailConfigInfo"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	ReportEmailConfigInfoList []struct {
		Email      api.String `json:"email"`
		ReportList []struct {
			CreateTime                 api.DateTime `json:"create_time"`
			CreateUserID               api.Integer  `json:"create_user_id"`
			EmailAddTime               api.DateTime `json:"email_add_time"`
			ID                         api.Integer  `json:"id"`
			IsAllowEmailSend           api.Bool     `json:"is_allow_email_send"`
			IsBank                     api.Bool     `json:"is_bank"`
			IsCanRenewSendConfirmEmail api.Bool     `json:"is_can_renew_send_confirm_email"`
			IsNewWeb                   api.Bool     `json:"is_new_web"`
			OrderID                    api.Integer  `json:"order_id"`
			ReSendConfirmEmailTime     api.DateTime `json:"re_send_confirm_email_time"`
			ReportID                   api.Integer  `json:"report_id"`
			ReportName                 api.String   `json:"report_name"`
			SendEmail                  api.String   `json:"send_email"`
			Status                     api.Bool     `json:"status"`
			TimeDimension              api.Integer  `json:"time_dimension"`
			Type                       api.Integer  `json:"type"`
			UpdateTime                 api.DateTime `json:"update_time"`
			UserID                     api.Integer  `json:"user_id"`
		} `json:"report_list"`
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
