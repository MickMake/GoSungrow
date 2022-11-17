package iSolarCloud

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
	"time"
)


// SunGrowDataRequest - Collection of all possible request args.
type SunGrowDataRequest struct {
	args     RequestArgs
	Required map[string]string
	aPsId    valueTypes.PsIds
}
type SunGrowDataRequests map[string]SunGrowDataRequest

type RequestArgs struct {
	// Be careful with types. If you see a general "error' response,
	// then it's more likely a type mismatch.
	PsId           *valueTypes.PsId     `json:"ps_id,omitempty"`
	PsIds          *valueTypes.PsIds    `json:"psIds,omitempty"`
	ReportType     *valueTypes.String   `json:"report_type,omitempty"`
	DateId         *valueTypes.DateTime `json:"date_id,omitempty"`
	DateType       *string              `json:"date_type,omitempty"`
	FaultTypeCode  *valueTypes.Integer  `json:"fault_type_code,omitempty"`
	Size           *valueTypes.Integer  `json:"size,omitempty"`
	CurPage        *valueTypes.Integer  `json:"curPage,omitempty"`
	DeviceType     *valueTypes.String   `json:"device_type,omitempty"`
	ReportId       *valueTypes.String   `json:"report_id,omitempty"`
	CodeType       *valueTypes.String   `json:"code_type,omitempty"`
	OrgIds         *valueTypes.String   `json:"orgIds,omitempty"`
	PsIdList       *valueTypes.String   `json:"ps_id_list,omitempty"`
	Uuid           *valueTypes.String   `json:"uuid,omitempty"`
	TemplateId     *valueTypes.String   `json:"template_id,omitempty"`
	DeviceModelId  *valueTypes.String   `json:"device_model_id,omitempty"`
	PsKey          *valueTypes.PsKey    `json:"ps_key,omitempty"`

	// UNVERIFIED
	Day            *valueTypes.DateTime  `json:"day,omitempty"`
	AppKey         *valueTypes.String    `json:"app_key,omitempty"`
	BeginTime      *valueTypes.DateTime  `json:"beginTime,omitempty"`
	DealerOrgCode  *valueTypes.String    `json:"dealer_org_code,omitempty"`
	DeviceSn       *valueTypes.String    `json:"device_sn,omitempty"`
	EndTimeStamp   *valueTypes.DateTime  `json:"end_time_stamp,omitempty"`
	FaultCode      *valueTypes.Integer   `json:"fault_code,omitempty"`
	FaultName      *valueTypes.String    `json:"fault_name,omitempty"`
	Id             *valueTypes.Integer   `json:"id,omitempty"`
	MinuteInterval *valueTypes.Integer   `json:"minute_interval,omitempty"`
	OrderId        *valueTypes.String    `json:"order_id,omitempty"`
	OrgId          *valueTypes.String    `json:"org_id,omitempty"`
	PointId        *valueTypes.PointId   `json:"point_id,omitempty"`
	Points         *valueTypes.String    `json:"points,omitempty"`
	Prefix         *valueTypes.String    `json:"prefix,omitempty"`
	PrimaryKey     *valueTypes.String    `json:"primaryKey,omitempty"`
	PsKeyList      *valueTypes.String    `json:"ps_key_list,omitempty"`
	QueryType      *valueTypes.String    `json:"query_type,omitempty"`
	Sn             *valueTypes.String    `json:"sn,omitempty"`
	StartTimeStamp *valueTypes.DateTime  `json:"start_time_stamp,omitempty"`
	Table          *valueTypes.String    `json:"table,omitempty"`
	TaskId         *valueTypes.String    `json:"task_id,omitempty"`

	// PsId valueTypes.PsId `json:"id"`
	// PsId valueTypes.PsId `json:"ps_id"`
	// DeviceType valueTypes.String `json:"device_type"`
	// DateType      *valueTypes.String   `json:"date_type,omitempty"`

	// AppKey string `json:"app_key"`
	// BeginTime string `json:"beginTime"`
	// DealerOrgCode string `json:"dealer_org_code"`
	// DeviceSn valueTypes.String `json:"device_sn"`
	// EndTimeStamp string `json:"end_time_stamp"`
	// FaultCode string `json:"fault_code"`
	// FaultName string `json:"fault_name"`
	// FaultTypeCode string `json:"fault_type_code"`
	// Id string `json:"id"`
	// MinuteInterval string `json:"minute_interval"`
	// OrderId string `json:"order_id"`
	// OrgId string `json:"org_id"`
	// PointId string `json:"point_id"`
	// Points string `json:"points"`
	// Prefix string `json:"prefix"`
	// PrimaryKey string `json:"primaryKey"`
	// PsKeyList string `json:"ps_key_list"`
	// QueryType valueTypes.String `json:"query_type"`
	// Sn valueTypes.String `json:"sn"`
	// StartTimeStamp string `json:"start_time_stamp"`
	// Table string `json:"table"`
	// TaskId valueTypes.String `json:"task_id"`
}

const (
	NamePsId           = "PsId"
	NamePsIds          = "PsIds"
	NameReportType     = "ReportType"
	NameDateId         = "DateId"
	NameDateType       = "DateType"
	NameFaultTypeCode  = "FaultTypeCode"
	NameSize           = "Size"
	NameCurPage        = "CurPage"
	NameDeviceType     = "DeviceType"
	NameReportId       = "ReportId"
	NameCodeType       = "CodeType"
	NameOrgIds         = "OrgIds"
	NamePsIdList       = "PsIdList"
	NameUuid           = "Uuid"
	NameTemplateId     = "TemplateId"
	NameDeviceModelId  = "DeviceModelId"
	NamePsKey          = "PsKey"

	// UNVERIFIED
	NameDay            = "Day"
	NameAppKey         = "AppKey"
	NameBeginTime      = "BeginTime"
	NameDealerOrgCode  = "DealerOrgCode"
	NameDeviceSn       = "DeviceSn"
	NameEndTimeStamp   = "EndTimeStamp"
	NameFaultCode      = "FaultCode"
	NameFaultName      = "FaultName"
	NameId             = "Id"
	NameMinuteInterval = "MinuteInterval"
	NameOrderId        = "OrderId"
	NameOrgId          = "OrgId"
	NamePointId        = "PointId"
	NamePoints         = "Points"
	NamePrefix         = "Prefix"
	NamePrimaryKey     = "PrimaryKey"
	NamePsKeyList      = "PsKeyList"
	NameQueryType      = "QueryType"
	NameSn             = "Sn"
	NameStartTimeStamp = "StartTimeStamp"
	NameTable          = "Table"
	NameTaskId         = "TaskId"
)

var Help = map[string]string{
	NamePsId:           "PsId - valid ps_id",
	NamePsIds:          "PsIds - list of ps_id",
	NameReportType:     "ReportType - ",
	NameDateId:         "DateId - Date in format YYYYMMDD or YYYMM or YYYY",
	NameDateType:       "DateType - ",
	NameFaultTypeCode:  "FaultTypeCode - ",
	NameSize:           "Size - ",
	NameCurPage:        "CurPage - ",
	NameDeviceType:     "DeviceType - ",
	NameReportId:       "ReportId - ",
	NameCodeType:       "CodeType - ",
	NameOrgIds:         "OrgIds - ",
	NamePsIdList:       "PsIdList - ",
	NameUuid:           "Uuid - ",
	NameTemplateId:     "TemplateId - ",
	NameDeviceModelId:  "DeviceModelId - ",
	NamePsKey:          "PsKey - ",
	NameDay:            "Day - Date in format YYYYMMDD",

	// UNVERIFIED
	NameAppKey:         "AppKey - ",
	NameBeginTime:      "BeginTime - ",
	NameDealerOrgCode:  "DealerOrgCode - ",
	NameDeviceSn:       "DeviceSn - ",
	NameEndTimeStamp:   "EndTimeStamp - ",
	NameFaultCode:      "FaultCode - ",
	NameFaultName:      "FaultName - ",
	NameId:             "Id - ",
	NameMinuteInterval: "MinuteInterval - ",
	NameOrderId:        "OrderId - ",
	NameOrgId:          "OrgId - ",
	NamePointId:        "PointId - ",
	NamePoints:         "Points - ",
	NamePrefix:         "Prefix - ",
	NamePrimaryKey:     "PrimaryKey - ",
	NamePsKeyList:      "PsKeyList - ",
	NameQueryType:      "QueryType - ",
	NameSn:             "Sn - ",
	NameStartTimeStamp: "StartTimeStamp - ",
	NameTable:          "Table - ",
	NameTaskId:         "TaskId - ",
}

// MarshalJSON - Convert value to JSON
func (sgd SunGrowDataRequest) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		var dt *string
		if sgd.args.DateId != nil {
			dt = &sgd.args.DateId.DateType
		}

		type Parse RequestArgs
		// Store result from string
		data, err = json.Marshal(Parse {
			PsId:           sgd.args.PsId,
			PsIds:          sgd.args.PsIds,
			ReportType:     sgd.args.ReportType,
			DateId:         sgd.args.DateId,
			DateType:       dt,
			FaultTypeCode:  sgd.args.FaultTypeCode,
			Size:           sgd.args.Size,
			CurPage:        sgd.args.CurPage,
			DeviceType:     sgd.args.DeviceType,
			ReportId:       sgd.args.ReportId,
			CodeType:       sgd.args.CodeType,
			OrgIds:         sgd.args.OrgIds,
			PsIdList:       sgd.args.PsIdList,
			Uuid:           sgd.args.Uuid,
			TemplateId:     sgd.args.TemplateId,
			DeviceModelId:  sgd.args.DeviceModelId,
			PsKey:          sgd.args.PsKey,

			// UNVERIFIED
			Day:            sgd.args.Day,
			AppKey:         sgd.args.AppKey,
			BeginTime:      sgd.args.BeginTime,
			DealerOrgCode:  sgd.args.DealerOrgCode,
			DeviceSn:       sgd.args.DeviceSn,
			EndTimeStamp:   sgd.args.EndTimeStamp,
			FaultCode:      sgd.args.FaultCode,
			FaultName:      sgd.args.FaultName,
			Id:             sgd.args.Id,
			MinuteInterval: sgd.args.MinuteInterval,
			OrderId:        sgd.args.OrderId,
			OrgId:          sgd.args.OrgId,
			PointId:        sgd.args.PointId,
			Points:         sgd.args.Points,
			Prefix:         sgd.args.Prefix,
			PrimaryKey:     sgd.args.PrimaryKey,
			PsKeyList:      sgd.args.PsKeyList,
			QueryType:      sgd.args.QueryType,
			Sn:             sgd.args.Sn,
			StartTimeStamp: sgd.args.StartTimeStamp,
			Table:          sgd.args.Table,
			TaskId:         sgd.args.TaskId,
		})
		if err == nil {
			break
		}
	}

	return data, err
}

func (sgd *SunGrowDataRequest) Set(arg string, value string) {
	for range Only.Once {
		switch arg {
			case NamePsId:
				sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value))	// strings.Split(a[1], ","))
				sgd.SetPsId(value)
				// v, err := strconv.ParseInt(value, 10, 64)
				// if err != nil {
				// 	fmt.Printf("Error: %s - %s\n", NamePsId, err)
				// 	fmt.Printf("%s\n", Help[arg])
				// }
				// val := valueTypes.SetPsIdValue(v); sgd.args.PsId = &val
			case NamePsIds:
				sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value))	// strings.Split(a[1], ","))
				val := valueTypes.SetPsIdStrings(SplitArg(value)); sgd.args.PsIds = &val
			case NameReportType:
				val := valueTypes.SetStringValue(value); sgd.args.ReportType = &val
			case NameDateId:
				sgd.SetDateId(value)
			case NameFaultTypeCode:
				val := valueTypes.SetIntegerString(value); sgd.args.FaultTypeCode = &val
			case NameSize:
				val := valueTypes.SetIntegerString(value); sgd.args.Size = &val
			case NameCurPage:
				val := valueTypes.SetIntegerString(value); sgd.args.CurPage = &val
			case NameDeviceType:
				val := valueTypes.SetStringValue(value); sgd.args.DeviceType = &val
			case NameReportId:
				val := valueTypes.SetStringValue(value); sgd.args.ReportId = &val
			case NameCodeType:
				val := valueTypes.SetStringValue(value); sgd.args.CodeType = &val
			case NameOrgIds:
				val := valueTypes.SetStringValue(value); sgd.args.OrgIds = &val
			case NamePsIdList:
				val := valueTypes.SetStringValue(value); sgd.args.PsIdList = &val
			case NameUuid:
				val := valueTypes.SetStringValue(value); sgd.args.Uuid = &val
			case NameTemplateId:
				val := valueTypes.SetStringValue(value); sgd.args.TemplateId = &val
			case NameDeviceModelId:
				val := valueTypes.SetStringValue(value); sgd.args.DeviceModelId = &val
			case NamePsKey:
				val := valueTypes.SetPsKeyValue(value); sgd.args.PsKey = &val


			// UNVERIFIED
			case NameDay:
				sgd.SetDay(value)
			case NameAppKey:
				val := valueTypes.SetStringValue(value); sgd.args.AppKey = &val
			case NameBeginTime:
				val := valueTypes.SetDateTimeString(value); sgd.args.BeginTime = &val
			case NameDealerOrgCode:
				val := valueTypes.SetStringValue(value); sgd.args.DealerOrgCode = &val
			case NameDeviceSn:
				val := valueTypes.SetStringValue(value); sgd.args.DeviceSn = &val
			case NameEndTimeStamp:
				val := valueTypes.SetDateTimeString(value); sgd.args.EndTimeStamp = &val
			case NameFaultCode:
				val := valueTypes.SetIntegerString(value); sgd.args.FaultCode = &val
			case NameFaultName:
				val := valueTypes.SetStringValue(value); sgd.args.FaultName = &val
			case NameId:
				val := valueTypes.SetIntegerString(value); sgd.args.Id = &val
			case NameMinuteInterval:
				val := valueTypes.SetIntegerString(value); sgd.args.MinuteInterval = &val
			case NameOrderId:
				val := valueTypes.SetStringValue(value); sgd.args.OrderId = &val
			case NameOrgId:
				val := valueTypes.SetStringValue(value); sgd.args.OrgId = &val
			case NamePointId:
				val := valueTypes.SetPointIdString(value); sgd.args.PointId = &val
			case NamePoints:
				val := valueTypes.SetStringValue(value); sgd.args.Points = &val
			case NamePrefix:
				val := valueTypes.SetStringValue(value); sgd.args.Prefix = &val
			case NamePrimaryKey:
				val := valueTypes.SetStringValue(value); sgd.args.PrimaryKey = &val
			case NamePsKeyList:
				val := valueTypes.SetStringValue(value); sgd.args.PsKeyList = &val
			case NameQueryType:
				val := valueTypes.SetStringValue(value); sgd.args.QueryType = &val
			case NameSn:
				val := valueTypes.SetStringValue(value); sgd.args.Sn = &val
			case NameStartTimeStamp:
				val := valueTypes.SetDateTimeString(value); sgd.args.StartTimeStamp = &val
			case NameTable:
				val := valueTypes.SetStringValue(value); sgd.args.Table = &val
			case NameTaskId:
				val := valueTypes.SetStringValue(value); sgd.args.TaskId = &val
		}
	}
}

func (sgd *SunGrowDataRequest) IsSet(arg string) bool {
	var ok bool
	for range Only.Once {
		switch arg {
			case NamePsId:
				if sgd.args.PsId != nil { ok = true }
			case NamePsIds:
				if sgd.args.PsIds != nil { ok = true }
			case NameReportType:
				if sgd.args.ReportType != nil { ok = true }
			case NameDateId:
				if sgd.args.DateId != nil { ok = true }
			case NameFaultTypeCode:
				if sgd.args.FaultTypeCode != nil { ok = true }
			case NameSize:
				if sgd.args.Size != nil { ok = true }
			case NameCurPage:
				if sgd.args.CurPage != nil { ok = true }
			case NameDeviceType:
				if sgd.args.DeviceType != nil { ok = true }
			case NameReportId:
				if sgd.args.ReportId != nil { ok = true }
			case NameCodeType:
				if sgd.args.CodeType != nil { ok = true }
			case NameOrgIds:
				if sgd.args.OrgIds != nil { ok = true }
			case NamePsIdList:
				if sgd.args.PsIdList != nil { ok = true }
			case NameUuid:
				if sgd.args.Uuid != nil { ok = true }
			case NameTemplateId:
				if sgd.args.TemplateId != nil { ok = true }
			case NameDeviceModelId:
				if sgd.args.DeviceModelId != nil { ok = true }
			case NamePsKey:
				if sgd.args.PsKey != nil { ok = true }


			// UNVERIFIED
			case NameDay:
				if sgd.args.Day != nil { ok = true }
			case NameAppKey:
				if sgd.args.AppKey != nil { ok = true }
			case NameBeginTime:
				if sgd.args.BeginTime != nil { ok = true }
			case NameDealerOrgCode:
				if sgd.args.DealerOrgCode != nil { ok = true }
			case NameDeviceSn:
				if sgd.args.DeviceSn != nil { ok = true }
			case NameEndTimeStamp:
				if sgd.args.EndTimeStamp != nil { ok = true }
			case NameFaultCode:
				if sgd.args.FaultCode != nil { ok = true }
			case NameFaultName:
				if sgd.args.FaultName != nil { ok = true }
			case NameId:
				if sgd.args.Id != nil { ok = true }
			case NameMinuteInterval:
				if sgd.args.MinuteInterval != nil { ok = true }
			case NameOrderId:
				if sgd.args.OrderId != nil { ok = true }
			case NameOrgId:
				if sgd.args.OrgId != nil { ok = true }
			case NamePointId:
				if sgd.args.PointId != nil { ok = true }
			case NamePoints:
				if sgd.args.Points != nil { ok = true }
			case NamePrefix:
				if sgd.args.Prefix != nil { ok = true }
			case NamePrimaryKey:
				if sgd.args.PrimaryKey != nil { ok = true }
			case NamePsKeyList:
				if sgd.args.PsKeyList != nil { ok = true }
			case NameQueryType:
				if sgd.args.QueryType != nil { ok = true }
			case NameSn:
				if sgd.args.Sn != nil { ok = true }
			case NameStartTimeStamp:
				if sgd.args.StartTimeStamp != nil { ok = true }
			case NameTable:
				if sgd.args.Table != nil { ok = true }
			case NameTaskId:
				if sgd.args.TaskId != nil { ok = true }
		}
	}
	return ok
}

func (sgd *SunGrowDataRequest) SetArgs(args ...string) {
	for range Only.Once {
		for _, arg := range args {
			a := strings.Split(arg, ":")
			if len(a) == 0 {
				continue
			}

			if len(a) == 1 {
				a = append(a, "")
			}

			sgd.Set(a[0], a[1])
		}
	}
}

func (sgd *SunGrowDataRequest) Validate(endpoint api.EndPoint) bool {
	ok := true
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			if value != "true" {
				continue
			}

			if key == NamePsId {
				// Handled differently.
				continue
			}
			if key == NamePsIds {
				// Handled differently.
				continue
			}

			if sgd.IsSet(key) {
				continue
			}

			ok = false
			fmt.Printf("%s is required\n", key)
			fmt.Printf("%s\n", Help[key])
		}
	}
	return ok
}

func (sgd *SunGrowDataRequest) Help(endpoint api.EndPoint) {
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			if key == NamePsId {
				continue
			}
			if key == NamePsIds {
				continue
			}

			if key == NameDateType {
				continue
			}

			if value != "true" {
				fmt.Printf("optional - %s:value\n", key)
				continue
			}

			fmt.Printf("required - %s:value\n", key)
		}
	}
}

func (sgd *SunGrowDataRequest) RequestAsFilePrefix() string {
	var ret string
	for range Only.Once {
		j, err := json.Marshal(*sgd)
		if err != nil {
			break
		}
		ret = string(j)
		ret = strings.TrimPrefix(ret, "{")
		ret = strings.TrimSuffix(ret, "}")
		ret = strings.ReplaceAll(ret, `"`, ``)
		ret = strings.ReplaceAll(ret, `,`, `-`)
	}
	return ret
}

func (sgd *SunGrowDataRequest) SetDateId(date string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameDateId) {
		// 	break
		// }
		did := valueTypes.SetDateTimeString(date)
		sgd.args.DateId = &did
		if sgd.args.DateId.IsZero() {
			did = valueTypes.NewDateTime(valueTypes.Now)
			sgd.args.DateId = &did
		}
	}
}

func (sgd *SunGrowDataRequest) SetDay(date string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameDay) {
		// 	break
		// }
		did := valueTypes.SetDateTimeString(date)
		sgd.args.Day = &did
		if sgd.args.Day.IsZero() {
			now := time.Now().Format(valueTypes.DateLayoutDay)
			did = valueTypes.NewDateTime(now)
			sgd.args.Day = &did
		}
	}
}

func (sgd *SunGrowDataRequest) SetFaultTypeCode(ftc string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameFaultTypeCode) {
		// 	break
		// }
		aftc := valueTypes.SetIntegerString(ftc)
		sgd.args.FaultTypeCode = &aftc
	}
}

func (sgd *SunGrowDataRequest) SetReportType(rt string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameReportType) {
		// 	break
		// }
		art := valueTypes.SetStringValue(rt)
		sgd.args.ReportType = &art
	}
}

func (sgd *SunGrowDataRequest) SetPSIDs(psIds []string) {
	sgd.aPsId = valueTypes.SetPsIdStrings(psIds)
}

func (sgd *SunGrowDataRequest) GetPSIDs() valueTypes.PsIds {
	return sgd.aPsId
}

func (sgd *SunGrowDataRequest) SetRequired(req map[string]string) {
	sgd.Required = req
}

func (sgd *SunGrowDataRequest) IsRequiredAndNotSet(arg string) bool {
	var yes bool
	for range Only.Once {
		if _, ok := sgd.Required[arg]; !ok {
			break
		}
		if sgd.IsNotRequired(arg) {
			break
		}
		if sgd.IsSet(arg) {
			break
		}
		yes = true
	}
	return yes
}

func (sgd *SunGrowDataRequest) IsRequiredAndSet(arg string) bool {
	return !sgd.IsRequiredAndNotSet(arg)
}

func (sgd *SunGrowDataRequest) SetIfRequired(arg string, value string) {
	for range Only.Once {
		if sgd.IsNotRequired(arg) {
			break
		}
		if sgd.IsRequiredAndSet(arg) {
			break
		}
		sgd.Set(arg, value)
	}
}

func (sgd *SunGrowDataRequest) SetPsId(psId string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NamePsId) {
		// 	break
		// }
		if psId == "" {
			sgd.args.PsId = nil
			break
		}

		// v, err := strconv.ParseInt(value, 10, 64)
		// if err != nil {
		// 	fmt.Printf("Error: %s - %s\n", NamePsId, err)
		// 	fmt.Printf("%s\n", Help[arg])
		// }
		// val := valueTypes.SetPsIdValue(v); sgd.args.PsId = &val

		pid := valueTypes.SetPsIdString(psId)
		if pid.Error != nil {
			fmt.Printf("Error: %s - %s\n", NamePsId, pid.Error)
		}
		sgd.args.PsId = &pid
	}
}

func (sgd *SunGrowDataRequest) IsRequired(req string) bool {
	var yes bool
	if _, ok := sgd.Required[req]; ok {
		yes = true
	}
	return yes
}

func (sgd *SunGrowDataRequest) IsNotRequired(req string) bool {
	return !sgd.IsRequired(req)
}
