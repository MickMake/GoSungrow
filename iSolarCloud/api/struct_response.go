package api

import (
	"GoSungro/Only"
	"errors"
	"fmt"
	"strings"
)


type Response struct {
	ResponseCommon
}

type ResponseCommon struct {
	ReqSerialNum string        `json:"req_serial_num"`
	ResultCode   string        `json:"result_code"`
	ResultData   []interface{} `json:"result_data"`
	ResultMsg    string        `json:"result_msg"`
}


func (req ResponseCommon) IsValid() error {
	var err error
	for range Only.Once {
		err = CheckResultMessage(req.ResultMsg)
		if err != nil {
			break
		}
		err = CheckResultCode(req.ResultCode)
		if err != nil {
			break
		}
		err = CheckString("ReqSerialNum", req.ReqSerialNum)
		if err != nil {
			break
		}
		// if req.ResultData == nil {
		// 	err = errors.New("zero results")
		// 	break
		// }
	}
	return err
}

// ResponseCommon checks
func CheckResultCode(rc string) error {
	var err error
	for range Only.Once {
		switch rc {
			case "1":
				err = nil
			case "-1":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "010":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "000":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "201":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "E00003":
				err = errors.New(fmt.Sprintf("need to login again '%s'", rc))
			default:
				err = errors.New(fmt.Sprintf("unknown error '%s'", rc))
		}
	}
	return err
}

func CheckResultMessage(rc string) error {
	var err error
	for range Only.Once {
		switch {
			case rc == "success":
				err = nil
			case rc == "er_invalid_appkey":
				err = errors.New(fmt.Sprintf("appkey is incorrect '%s'", rc))
			case rc == "er_token_login_invalid":
				err = errors.New(fmt.Sprintf("need to login again '%s'", rc))
			case rc == "er_parameter_value_invalid":
				err = errors.New(fmt.Sprintf("incorrect request data '%s'", rc))
			case rc == "er_unknown_exception":
				err = errors.New(fmt.Sprintf("API error '%s'", rc))
			case strings.HasPrefix(rc, "Parameter:"):
				err = errors.New(fmt.Sprintf("incorrect request data '%s'", rc))
			// case err == nil:
			// 	err = errors.New(fmt.Sprintf("unknown '%s'", rc))
			default:
				err = errors.New(fmt.Sprintf("unknown error '%s'", rc))
		}
	}
	return err
}
