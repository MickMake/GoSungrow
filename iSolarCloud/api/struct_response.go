package api

import (
	"github.com/MickMake/GoUnify/Only"
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
	ResultMsg    string        `json:"result_msg"`
}


func (req ResponseCommon) IsValid() error {
	var err error
	for range Only.Once {
		err = req.CheckResultMessage()
		if err != nil {
			break
		}

		err = req.CheckResultCode()
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

func (req ResponseCommon) IsTokenValid() bool {
	var ok bool
	for range Only.Once {
		switch {
			case req.ResultMsg == "success":
				ok = true
			case req.ResultMsg == "er_token_login_invalid":
				ok = false
			default:
				ok = false
		}
	}
	return ok
}

func (req ResponseCommon) IsTokenInvalid() bool {
	return !req.IsTokenValid()
}

func (req ResponseCommon) String() string {
	var ret string
	ret = fmt.Sprintf("ReqSerialNum:\t%s\n", req.ReqSerialNum)
	ret += fmt.Sprintf("ResultCode:\t%s\n", req.ResultCode)
	ret += fmt.Sprintf("ResultMsg:\t%s\n", req.ResultMsg)
	return ret
}

func (req ResponseCommon) CheckResultCode() error {
	var err error
	for range Only.Once {
		switch req.ResultCode {
		case "1":
			err = nil
		case "-1":
			err = errors.New(fmt.Sprintf("error '%s'", req.ResultCode))
		case "010":
			err = errors.New(fmt.Sprintf("error '%s'", req.ResultCode))
		case "000":
			err = errors.New(fmt.Sprintf("error '%s'", req.ResultCode))
		case "201":
			err = errors.New(fmt.Sprintf("error '%s'", req.ResultCode))
		case "E00003":
			err = errors.New(fmt.Sprintf("need to login again '%s'", req.ResultCode))
		default:
			err = errors.New(fmt.Sprintf("unknown error '%s'", req.ResultCode))
		}
	}
	return err
}

func (req ResponseCommon) CheckResultMessage() error {
	var err error
	for range Only.Once {
		switch {
		case req.ResultMsg == "success":
			err = nil
		case req.ResultMsg == `账号不存在`:
			err = errors.New(fmt.Sprintf("Account does not exist '%s'", req.ResultMsg))
		case req.ResultMsg == "er_invalid_appkey":
			err = errors.New(fmt.Sprintf("appkey is incorrect '%s'", req.ResultMsg))
		case req.ResultMsg == "er_token_login_invalid":
			err = errors.New(fmt.Sprintf("need to login again '%s'", req.ResultMsg))
		case req.ResultMsg == "er_parameter_value_invalid":
			err = errors.New(fmt.Sprintf("incorrect request data '%s'", req.ResultMsg))
		case req.ResultMsg == "er_unknown_exception":
			err = errors.New(fmt.Sprintf("API error '%s'", req.ResultMsg))
		case strings.HasPrefix(req.ResultMsg, "Parameter:"):
			err = errors.New(fmt.Sprintf("incorrect request data '%s'", req.ResultMsg))
		default:
			err = errors.New(fmt.Sprintf("unknown error '%s'", req.ResultMsg))
		}
	}
	return err
}
