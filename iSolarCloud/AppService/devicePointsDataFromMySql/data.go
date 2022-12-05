package devicePointsDataFromMySql

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/devicePointsDataFromMySql"
const Disabled = false
const EndPointName = "AppService.devicePointsDataFromMySql"

type RequestData struct {
	PsKey          valueTypes.PsKey `json:"ps_key" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	DevicePointData []struct {
		AlarmCount              valueTypes.Count       `json:"alarm_count"`
		DevFaultStatus          valueTypes.Integer       `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer       `json:"dev_status"`
		DevUpdateTime           valueTypes.DateTime `json:"dev_update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
		FaultCount              valueTypes.Count       `json:"fault_count"`
		InstallerAlarmCount     valueTypes.Count       `json:"installer_alarm_count"`
		InstallerDevFaultStatus valueTypes.Integer `json:"installer_dev_fault_status"`
		InstallerFaultCount     valueTypes.Count       `json:"installer_fault_count"`
		OwnerAlarmCount         valueTypes.Count       `json:"owner_alarm_count"`
		OwnerDevFaultStatus     valueTypes.Integer `json:"owner_dev_fault_status"`
		OwnerFaultCount         valueTypes.Integer       `json:"owner_fault_count"`
		P1                      interface{} `json:"p1"`
		P1001                   interface{} `json:"p1001"`
		P1002                   interface{} `json:"p1002"`
		P1005                   interface{} `json:"p1005"`
		P1302                   interface{} `json:"p1302"`
		P14                     interface{} `json:"p14"`
		P2                      interface{} `json:"p2"`
		P202                    interface{} `json:"p202"`
		P210                    interface{} `json:"p210"`
		P24                     interface{} `json:"p24"`
		P66                     interface{} `json:"p66"`
		P8018                   interface{} `json:"p8018"`
		P8030                   interface{} `json:"p8030"`
		P8031                   interface{} `json:"p8031"`
		P8062                   interface{} `json:"p8062"`
		P8063                   interface{} `json:"p8063"`
		P81002                  interface{} `json:"p81002"`
		P81022                  interface{} `json:"p81022"`
		P81023                  interface{} `json:"p81023"`
		P81202                  interface{} `json:"p81202"`
		P81203                  interface{} `json:"p81203"`
		PsKey                   valueTypes.PsKey `json:"ps_key"`
		UpdateTime              valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
		UpdateUser              valueTypes.Integer `json:"update_user"`
		UUID                    valueTypes.Integer `json:"uuid"`
	} `json:"device_point_data" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsKey.String()
		entries.StructToDataMap(*e,  e.Request.PsKey.String(), GoStruct.NewEndPointPath(e.Request.PsKey.String()))
	}

	return entries
}
