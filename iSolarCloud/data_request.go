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
	// Be careful with types. If you see a general 'error' response,
	// then it's more likely a type mismatch.
	PsId           *valueTypes.PsId     `json:"ps_id,omitempty"`
	PsId2          *valueTypes.PsId     `json:"psId,omitempty"`		// Specifically for WebIscmAppService.getPowerStationInfo
	PsId3          *valueTypes.PsId     `json:"psid,omitempty"`		// Specifically for WebAppService.getPsTree
	PsIds          *valueTypes.PsIds    `json:"psIds,omitempty"`
	PsIdList       *valueTypes.String   `json:"ps_id_list,omitempty"`
	ReportType     *valueTypes.String   `json:"report_type,omitempty"`
	FaultTypeCode  *valueTypes.Integer  `json:"fault_type_code,omitempty"`
	Size           *valueTypes.Integer  `json:"size,omitempty"`
	CurPage        *valueTypes.Integer  `json:"curPage,omitempty"`
	DeviceType     *valueTypes.String   `json:"device_type,omitempty"`
	ReportId       *valueTypes.String   `json:"report_id,omitempty"`
	CodeType       *valueTypes.String   `json:"code_type,omitempty"`
	OrgIds         *valueTypes.String   `json:"orgIds,omitempty"`
	Uuid           *valueTypes.String   `json:"uuid,omitempty"`
	TemplateId     *valueTypes.String   `json:"template_id,omitempty"`
	DeviceModelId  *valueTypes.String   `json:"device_model_id,omitempty"`
	PsKey          *valueTypes.PsKey    `json:"ps_key,omitempty"`

	// DateTime
	DateId         *valueTypes.DateTime  `json:"date_id,omitempty"`
	DateType       *string               `json:"date_type,omitempty"`
	MonthDate      *valueTypes.DateTime  `json:"month_date,omitempty"`
	Day            *valueTypes.DateTime  `json:"day,omitempty"`
	BeginTime      *valueTypes.DateTime  `json:"beginTime,omitempty"`	// valueTypes.Time
	EndTime        *valueTypes.DateTime  `json:"endTime,omitempty"`		// valueTypes.Time
	StartTimeStamp *valueTypes.DateTime  `json:"start_time_stamp,omitempty"`
	EndTimeStamp   *valueTypes.DateTime  `json:"end_time_stamp,omitempty"`

	// UNVERIFIED
	AppKey         *valueTypes.String    `json:"app_key,omitempty"`
	DealerOrgCode  *valueTypes.String    `json:"dealer_org_code,omitempty"`
	DeviceSn       *valueTypes.String    `json:"device_sn,omitempty"`
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
	Table          *valueTypes.String    `json:"table,omitempty"`
	TaskId         *valueTypes.String    `json:"task_id,omitempty"`
	UserId         *valueTypes.String    `json:"user_id,omitempty"`

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

//goland:noinspection GoCommentStart
const (
	NamePsId           = "PsId"
	NamePsId2          = "PsId2"
	NamePsId3          = "PsId3"
	NamePsIds          = "PsIds"
	NamePsIdList       = "PsIdList"
	NameReportType     = "ReportType"
	NameFaultTypeCode  = "FaultTypeCode"
	NameSize           = "Size"
	NameCurPage        = "CurPage"
	NameDeviceType     = "DeviceType"
	NameReportId       = "ReportId"
	NameCodeType       = "CodeType"
	NameOrgIds         = "OrgIds"
	NameUuid           = "Uuid"
	NameTemplateId     = "TemplateId"
	NameDeviceModelId  = "DeviceModelId"
	NamePsKey          = "PsKey"

	// DateTime
	NameDateId         = "DateId"
	NameDateType       = "DateType"
	NameMonthDate      = "MonthDate"
	NameDay            = "Day"
	NameBeginTime      = "BeginTime"
	NameEndTime        = "EndTime"
	NameStartTimeStamp = "StartTimeStamp"
	NameEndTimeStamp   = "EndTimeStamp"

	// UNVERIFIED
	NameAppKey         = "AppKey"
	NameDealerOrgCode  = "DealerOrgCode"
	NameDeviceSn       = "DeviceSn"
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
	NameTable          = "Table"
	NameTaskId         = "TaskId"
	NameUserId         = "UserId"
)

var Help = map[string]string{
	NamePsId:           "PsId - valid ps_id",
	NamePsId2:          "PsId - valid ps_id",
	NamePsId3:          "PsId - valid ps_id",
	NamePsIds:          "PsIds - list of ps_id",
	NamePsIdList:       "PsIdList - ",
	NameReportType:     "ReportType - ",
	NameFaultTypeCode:  "FaultTypeCode - ",
	NameSize:           "Size - ",
	NameCurPage:        "CurPage - ",
	NameDeviceType:     "DeviceType - ",
	NameReportId:       "ReportId - ",
	NameCodeType:       "CodeType - ",
	NameOrgIds:         "OrgIds - ",
	NameUuid:           "Uuid - ",
	NameTemplateId:     "TemplateId - ",
	NameDeviceModelId:  "DeviceModelId - ",
	NamePsKey:          "PsKey - ",

	// DateTime
	NameDateId:         "DateId - Date in format YYYYMMDD or YYYYMM or YYYY",
	NameDateType:       "DateType - ",
	NameMonthDate:      "MonthDate - Date in format YYYYMM",
	NameDay:            "Day - Date in format YYYYMMDD",
	NameBeginTime:      "BeginTime - ",
	NameEndTime:        "EndTime - ",
	NameStartTimeStamp: "StartTimeStamp - ",
	NameEndTimeStamp:   "EndTimeStamp - ",

	// UNVERIFIED
	NameAppKey:         "AppKey - ",
	NameDealerOrgCode:  "DealerOrgCode - ",
	NameDeviceSn:       "DeviceSn - ",
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
	NameTable:          "Table - ",
	NameTaskId:         "TaskId - ",
	NameUserId:         "UserId - ",
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
			PsId2:          sgd.args.PsId2,
			PsId3:          sgd.args.PsId3,
			PsIds:          sgd.args.PsIds,
			PsIdList:       sgd.args.PsIdList,
			ReportType:     sgd.args.ReportType,
			FaultTypeCode:  sgd.args.FaultTypeCode,
			Size:           sgd.args.Size,
			CurPage:        sgd.args.CurPage,
			DeviceType:     sgd.args.DeviceType,
			ReportId:       sgd.args.ReportId,
			CodeType:       sgd.args.CodeType,
			OrgIds:         sgd.args.OrgIds,
			Uuid:           sgd.args.Uuid,
			TemplateId:     sgd.args.TemplateId,
			DeviceModelId:  sgd.args.DeviceModelId,
			PsKey:          sgd.args.PsKey,

			// DateTime
			DateId:         sgd.args.DateId,
			DateType:       dt,
			MonthDate:      sgd.args.MonthDate,
			BeginTime:      sgd.args.BeginTime,
			EndTime:        sgd.args.EndTime,
			StartTimeStamp: sgd.args.StartTimeStamp,
			EndTimeStamp:   sgd.args.EndTimeStamp,

			// UNVERIFIED
			Day:            sgd.args.Day,
			AppKey:         sgd.args.AppKey,
			DealerOrgCode:  sgd.args.DealerOrgCode,
			DeviceSn:       sgd.args.DeviceSn,
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
			Table:          sgd.args.Table,
			TaskId:         sgd.args.TaskId,
			UserId:         sgd.args.UserId,
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
				fallthrough
			case NamePsId2:
				fallthrough
			case NamePsId3:
				sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value))	// strings.Split(a[1], ","))
				sgd.SetPsId(value)
			case NamePsIds:
				sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value))	// strings.Split(a[1], ","))
				val := valueTypes.SetPsIdStrings(SplitArg(value)); sgd.args.PsIds = &val
			case NameReportType:
				val := valueTypes.SetStringValue(value); sgd.args.ReportType = &val
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

			// DateTime
			case NameDateId:
				sgd.SetDateId(value)
			case NameMonthDate:
				sgd.SetMonthDate(value)
			case NameDay:
				sgd.SetDay(value)
			case NameBeginTime:
				val := valueTypes.SetDateTimeString(value); sgd.args.BeginTime = &val
			case NameEndTime:
				val := valueTypes.SetDateTimeString(value); sgd.args.EndTime = &val
			case NameStartTimeStamp:
				val := valueTypes.SetDateTimeString(value); sgd.args.StartTimeStamp = &val
			case NameEndTimeStamp:
				val := valueTypes.SetDateTimeString(value); sgd.args.EndTimeStamp = &val

			// UNVERIFIED
			case NameAppKey:
				val := valueTypes.SetStringValue(value); sgd.args.AppKey = &val
			case NameDealerOrgCode:
				val := valueTypes.SetStringValue(value); sgd.args.DealerOrgCode = &val
			case NameDeviceSn:
				val := valueTypes.SetStringValue(value); sgd.args.DeviceSn = &val
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
			case NameTable:
				val := valueTypes.SetStringValue(value); sgd.args.Table = &val
			case NameTaskId:
				val := valueTypes.SetStringValue(value); sgd.args.TaskId = &val
			case NameUserId:
				val := valueTypes.SetStringValue(value); sgd.args.UserId = &val
		}
	}
}

func (sgd *SunGrowDataRequest) Get(arg string) string {
	var value string
	for range Only.Once {
		switch arg {
			case NamePsId:
				value = sgd.args.PsId.String()
			case NamePsId2:
				value = sgd.args.PsId2.String()
			case NamePsId3:
				value = sgd.args.PsId3.String()
			case NamePsIds:
				value = strings.Join(sgd.args.PsIds.Strings(), ",")
			case NameReportType:
				value = sgd.args.ReportType.String()
			case NameFaultTypeCode:
				value = sgd.args.FaultTypeCode.String()
			case NameSize:
				value = sgd.args.Size.String()
			case NameCurPage:
				value = sgd.args.CurPage.String()
			case NameDeviceType:
				value = sgd.args.DeviceType.String()
			case NameReportId:
				value = sgd.args.ReportId.String()
			case NameCodeType:
				value = sgd.args.CodeType.String()
			case NameOrgIds:
				value = sgd.args.OrgIds.String()
			case NamePsIdList:
				value = sgd.args.PsIdList.String()
			case NameUuid:
				value = sgd.args.Uuid.String()
			case NameTemplateId:
				value = sgd.args.TemplateId.String()
			case NameDeviceModelId:
				value = sgd.args.DeviceModelId.String()
			case NamePsKey:
				value = sgd.args.PsKey.String()


			// DateTime
			case NameMonthDate:
				value = sgd.args.MonthDate.Original()
			case NameDateId:
				value = sgd.args.DateId.Original()
			case NameDay:
				value = sgd.args.Day.Original()
			case NameBeginTime:
				value = sgd.args.BeginTime.Original()
			case NameEndTime:
				value = sgd.args.EndTime.Original()
			case NameStartTimeStamp:
				value = sgd.args.StartTimeStamp.Original()
			case NameEndTimeStamp:
				value = sgd.args.EndTimeStamp.Original()

			// UNVERIFIED
			case NameAppKey:
				value = sgd.args.AppKey.String()
			case NameDealerOrgCode:
				value = sgd.args.DealerOrgCode.String()
			case NameDeviceSn:
				value = sgd.args.DeviceSn.String()
			case NameFaultCode:
				value = sgd.args.FaultCode.String()
			case NameFaultName:
				value = sgd.args.FaultName.String()
			case NameId:
				value = sgd.args.Id.String()
			case NameMinuteInterval:
				value = sgd.args.MinuteInterval.String()
			case NameOrderId:
				value = sgd.args.OrderId.String()
			case NameOrgId:
				value = sgd.args.OrgId.String()
			case NamePointId:
				value = sgd.args.PointId.String()
			case NamePoints:
				value = sgd.args.Points.String()
			case NamePrefix:
				value = sgd.args.Prefix.String()
			case NamePrimaryKey:
				value = sgd.args.PrimaryKey.String()
			case NamePsKeyList:
				value = sgd.args.PsKeyList.String()
			case NameQueryType:
				value = sgd.args.QueryType.String()
			case NameSn:
				value = sgd.args.Sn.String()
			case NameTable:
				value = sgd.args.Table.String()
			case NameTaskId:
				value = sgd.args.TaskId.String()
			case NameUserId:
				value = sgd.args.UserId.String()
		}
	}
	return value
}

func (sgd *SunGrowDataRequest) IsSet(arg string) bool {
	var ok bool
	for range Only.Once {
		switch arg {
			case NamePsId:
				if sgd.args.PsId != nil { ok = true }
			case NamePsId2:
				if sgd.args.PsId2 != nil { ok = true }
			case NamePsId3:
				if sgd.args.PsId3 != nil { ok = true }
			case NamePsIds:
				if sgd.args.PsIds != nil { ok = true }
			case NameReportType:
				if sgd.args.ReportType != nil { ok = true }
			case NameDateId:
				if sgd.args.DateId != nil { ok = true }
			case NameFaultTypeCode:
				if sgd.args.FaultTypeCode != nil { ok = true }
			case NameMonthDate:
				if sgd.args.MonthDate != nil { ok = true }
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
			case NameEndTime:
				if sgd.args.EndTime != nil { ok = true }
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
			case NameUserId:
				if sgd.args.UserId != nil { ok = true }
		}
	}
	return ok
}

func (sgd *SunGrowDataRequest) IsNotSet(arg string) bool {
	return !sgd.IsSet(arg)
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
			if key == NamePsId2 {
				// Handled differently.
				continue
			}
			if key == NamePsId3 {
				// Handled differently.
				continue
			}
			if key == NamePsIds {
				// Handled differently.
				continue
			}

			if key == NameDateType {
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

func (sgd *SunGrowDataRequest) GetArgs(endpoint api.EndPoint) string {
	var ret string
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			if value != "true" {
				continue
			}

			if key == NameDateType {
				continue	// Handled differently.
			}

			if sgd.IsNotSet(key) {
				continue
			}

			ret += fmt.Sprintf("%s:%s ", key, sgd.Get(key))
		}
	}
	return ret
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
		var did valueTypes.DateTime
		if date == "total" {
			did = valueTypes.NewDateTime(valueTypes.Now)
			did.DateType = "4"
		} else {
			did = valueTypes.SetDateTimeString(date)
		}
		sgd.args.DateId = &did
		if sgd.args.DateId.IsZero() {
			did = valueTypes.NewDateTime(valueTypes.Now)
			sgd.args.DateId = &did
		}

		sgd.args.DateType = &did.DateType
	}
}

func (sgd *SunGrowDataRequest) SetMonthDate(date string) {
	for range Only.Once {
		did := valueTypes.SetDateTimeString(date)
		sgd.args.MonthDate = &did
		if sgd.args.MonthDate.IsZero() {
			did = valueTypes.NewDateTime(valueTypes.Now)
			sgd.args.MonthDate = &did
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

func (sgd *SunGrowDataRequest) SetStartTimeStamp(date string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameDay) {
		// 	break
		// }
		did := valueTypes.SetDateTimeString(date)
		sgd.args.StartTimeStamp = &did
		if sgd.args.StartTimeStamp.IsZero() {
			now := time.Now().Format(valueTypes.DateTimeLayoutSecond)
			did = valueTypes.NewDateTime(now)
			sgd.args.StartTimeStamp = &did
			// Use valueTypes.GetDayStartTimestamp
		}
	}
}

func (sgd *SunGrowDataRequest) SetEndTimeStamp(date string) {
	for range Only.Once {
		// if sgd.IsNotRequired(NameDay) {
		// 	break
		// }
		did := valueTypes.SetDateTimeString(date)
		sgd.args.EndTimeStamp = &did
		if sgd.args.EndTimeStamp.IsZero() {
			now := time.Now().Format(valueTypes.DateTimeLayoutSecond)
			did = valueTypes.NewDateTime(now)
			sgd.args.EndTimeStamp = &did
			// Use valueTypes.GetDayEndTimestamp
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

		pid := valueTypes.SetPsIdString(psId)
		if pid.Error != nil {
			fmt.Printf("Error: %s - %s\n", NamePsId, pid.Error)
			break
		}

		if sgd.IsRequired(NamePsId) {
			sgd.args.PsId = &pid
		}
		if sgd.IsRequired(NamePsId2) {
			sgd.args.PsId2 = &pid		// Specifically for WebIscmAppService.getPowerStationInfo
		}
		if sgd.IsRequired(NamePsId3) {
			sgd.args.PsId3 = &pid		// Specifically for WebAppService.getPsTree
		}
		if sgd.IsRequired(NamePsIds) {
			// sgd.args.PsIds = &pid
		}

		sgd.SetIfRequired(NamePsId, psId)
		// sgd.SetIfRequired(NamePsId2, psId)
		// sgd.SetIfRequired(NamePsId3, psId)
		// sgd.SetIfRequired(NamePsIds, psId)
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

func (sgd *SunGrowDataRequest) IsPsIdRequired() bool {
	if sgd.IsRequired(NamePsId) {
		return true
	}
	if sgd.IsRequired(NamePsId2) {
		return true
	}
	if sgd.IsRequired(NamePsId3) {
		return true
	}
	return false
}

func (sgd *SunGrowDataRequest) IsPsIdNotRequired() bool {
	return !sgd.IsPsIdRequired()
}
