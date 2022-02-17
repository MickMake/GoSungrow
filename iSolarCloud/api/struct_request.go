package api

import (
	"GoSungro/Only"
	"errors"
	"strings"
)


type Request struct {
	RequestCommon
}

type RequestCommon struct {
	Appkey     string `json:"appkey"`
	Lang       string `json:"lang"`
	SysCode    string `json:"sys_code"`
	Token      string `json:"token"`
	UserID     string `json:"user_id"`
	ValidFlag  string `json:"valid_flag"`
	// DeviceType string `json:"device_type"`
}


func (req RequestCommon) IsValid() error {
	var err error
	for range Only.Once {
		err = CheckString("Appkey", req.Appkey)
		if err != nil {
			break
		}
		err = CheckString("Lang", req.Lang)
		if err != nil {
			break
		}
		err = CheckString("SysCode", req.SysCode)
		if err != nil {
			break
		}
		err = CheckString("Auth", req.Token)
		if err != nil {
			break
		}
		err = CheckString("UserID", req.UserID)
		if err != nil {
			break
		}
		err = CheckString("ValidFlag", req.ValidFlag)
		if err != nil {
			break
		}
	}
	return err
}

// RequestCommon checks
func CheckString(name string, rc string) error {
	var err error
	for range Only.Once {
		if rc == "" {
			err = errors.New(name + ": empty string")
			break
		}
		if strings.TrimSpace(rc) == "" {
			err = errors.New(name + ": empty string with spaces")
			break
		}
	}
	return err
}
