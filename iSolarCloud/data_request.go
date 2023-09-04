package iSolarCloud

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
	"time"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
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
	PsId          *valueTypes.PsId    `json:"ps_id,omitempty"`
	PsId2         *valueTypes.PsId    `json:"psId,omitempty"` // Specifically for WebIscmAppService.getPowerStationInfo & AppService/querySysAdvancedParam & WebIscmAppService/getContactMessage & WebAppService/getDevicePointAttrs
	PsId3         *valueTypes.PsId    `json:"psid,omitempty"` // Specifically for WebAppService.getPsTree
	PsIds         *valueTypes.PsIds   `json:"psIds,omitempty"`
	PsIdList      *valueTypes.String  `json:"ps_id_list,omitempty"`
	ReportType    *valueTypes.Integer `json:"report_type,omitempty"`
	FaultTypeCode *valueTypes.Integer `json:"fault_type_code,omitempty"`
	Size          *valueTypes.Integer `json:"size,omitempty"`
	CurPage       *valueTypes.Integer `json:"curPage,omitempty"`
	DeviceType    *valueTypes.Integer `json:"device_type,omitempty"`
	DeviceType2   *valueTypes.Integer `json:"deviceType,omitempty"` // Specifically for WebAppService.getDevicePointAttrs
	ReportId      *valueTypes.String  `json:"report_id,omitempty"`
	CodeType      *valueTypes.String  `json:"code_type,omitempty"`
	OrgIds        *valueTypes.String  `json:"orgIds,omitempty"`
	TemplateId    *valueTypes.String  `json:"template_id,omitempty"`
	DeviceModelId *valueTypes.String  `json:"device_model_id,omitempty"`
	Uuid          *valueTypes.Integer `json:"uuid,omitempty"`
	UuidList      *valueTypes.String  `json:"uuid_list,omitempty"`
	SetType       *valueTypes.String  `json:"set_type,omitempty"`
	Type          *valueTypes.String  `json:"type,omitempty"`
	DataType      *valueTypes.String  `json:"data_type,omitempty"`

	// Points
	PsKeyList *valueTypes.String   `json:"ps_key_list,omitempty"`
	PsKeys    *valueTypes.PsKeys   `json:"-,omitempty"` // Used by queryMutiPointDataList
	PsKey     *valueTypes.PsKey    `json:"ps_key,omitempty"`
	Point     *valueTypes.PointId  `json:"point_id,omitempty"`
	Points    *valueTypes.PointIds `json:"points,omitempty"`
	DataPoint *valueTypes.String   `json:"data_point,omitempty"`

	// DateTime
	DateId         *valueTypes.DateTime `json:"date_id,omitempty"`
	DateType       *string              `json:"date_type,omitempty"`
	MonthDate      *valueTypes.DateTime `json:"month_date,omitempty"`
	MonthDate2     *valueTypes.DateTime `json:"monthDate,omitempty"`
	Day            *valueTypes.DateTime `json:"day,omitempty"`
	BeginTime1     *valueTypes.DateTime `json:"beginTime,omitempty"` // valueTypes.Time
	EndTime1       *valueTypes.DateTime `json:"endTime,omitempty"`   // valueTypes.Time
	StartTimeStamp *valueTypes.DateTime `json:"start_time_stamp,omitempty"`
	EndTimeStamp   *valueTypes.DateTime `json:"end_time_stamp,omitempty"`
	StartTime      *valueTypes.DateTime `json:"start_time,omitempty"` // valueTypes.Time
	EndTime        *valueTypes.DateTime `json:"end_time,omitempty"`   // valueTypes.Time

	// UNVERIFIED
	AppKey         *valueTypes.String  `json:"app_key,omitempty"`
	DealerOrgCode  *valueTypes.String  `json:"dealer_org_code,omitempty"`
	DeviceSn       *valueTypes.String  `json:"device_sn,omitempty"`
	FaultCode      *valueTypes.Integer `json:"fault_code,omitempty"`
	FaultName      *valueTypes.String  `json:"fault_name,omitempty"`
	Id             *valueTypes.Integer `json:"id,omitempty"`
	MinuteInterval *valueTypes.Integer `json:"minute_interval,omitempty"`
	OrderId        *valueTypes.String  `json:"order_id,omitempty"`
	OrgId          *valueTypes.Integer `json:"org_id,omitempty"`
	Prefix         *valueTypes.String  `json:"prefix,omitempty"`
	PrimaryKey     *valueTypes.String  `json:"primaryKey,omitempty"`
	QueryType      *valueTypes.String  `json:"query_type,omitempty"`
	Sn             *valueTypes.String  `json:"sn,omitempty"`
	Table          *valueTypes.String  `json:"table,omitempty"`
	TaskId         *valueTypes.String  `json:"task_id,omitempty"`
	UserId         *valueTypes.String  `json:"userId,omitempty"`
	MenuId         *valueTypes.Integer `json:"menuId,omitempty"`
	// UserId2        *valueTypes.String    `json:"userId,omitempty"`

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
	NamePsId          = "PsId"
	NamePsId2         = "PsId2"
	NamePsId3         = "PsId3"
	NamePsIds         = "PsIds"
	NamePsIdList      = "PsIdList"
	NameReportType    = "ReportType"
	NameFaultTypeCode = "FaultTypeCode"
	NameSize          = "Size"
	NameCurPage       = "CurPage"
	NameDeviceType    = "DeviceType"
	NameDeviceType2   = "DeviceType2"
	NameReportId      = "ReportId"
	NameCodeType      = "CodeType"
	NameOrgIds        = "OrgIds"
	NameTemplateId    = "TemplateId"
	NameDeviceModelId = "DeviceModelId"
	NameUuid          = "Uuid"
	NameUuidList      = "UuidList"
	NameSetType       = "SetType"
	NameType          = "Type"
	NameDataType      = "DataType"

	NamePsKeyList = "PsKeyList"
	NamePsKeys    = "PsKeys"
	NamePsKey     = "PsKey"
	NamePointId   = "PointId"
	NamePoints    = "Points"
	NameDataPoint = "DataPoint"

	// DateTime
	NameDateId         = "DateId"
	NameDateType       = "DateType"
	NameMonthDate      = "MonthDate"
	NameDay            = "Day"
	NameBeginTime1     = "BeginTime1"
	NameEndTime1       = "EndTime1"
	NameStartTimeStamp = "StartTimeStamp"
	NameEndTimeStamp   = "EndTimeStamp"
	NameStartTime      = "StartTime"
	NameEndTime        = "EndTime"

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
	NamePrefix         = "Prefix"
	NamePrimaryKey     = "PrimaryKey"
	NameQueryType      = "QueryType"
	NameSn             = "Sn"
	NameTable          = "Table"
	NameTaskId         = "TaskId"
	NameUserId         = "UserId"
	NameMenuId         = "MenuId"
)

var Help = map[string]string{
	NamePsId:          "valid ps_id",
	NamePsId2:         "valid ps_id",
	NamePsId3:         "valid ps_id",
	NamePsIds:         "list of ps_id",
	NamePsIdList:      "",
	NameReportType:    api.HelpReportType(),
	NameFaultTypeCode: "",
	NameSize:          "Page size",
	NameCurPage:       "Current page",
	NameDeviceType:    "",
	NameDeviceType2:   "",
	NameReportId:      "",
	NameCodeType:      "",
	NameOrgIds:        "",
	NameTemplateId:    "",
	NameDeviceModelId: "",
	NameUuid:          "UUID",
	NameUuidList:      "Comma separated list of UUIDs",
	NameSetType:       "",
	NameType:          api.HelpReportType(),
	NameDataType:      "",

	NamePsKeyList: "",
	NamePsKey:     "",
	NamePointId:   "",
	NamePoints:    "",
	NameDataPoint: "",

	// DateTime
	NameDateId:         "Date in format YYYYMMDD or YYYYMM or YYYY",
	NameDateType:       "",
	NameMonthDate:      "Date in format YYYYMM",
	NameDay:            "Date in format YYYYMMDD",
	NameBeginTime1:     "",
	NameEndTime1:       "",
	NameStartTimeStamp: "Date in format YYYYMMDDHHMMSS",
	NameEndTimeStamp:   "Date in format YYYYMMDDHHMMSS",
	NameStartTime:      "Date in format YYYYMMDD or YYYYMM or YYYY",
	NameEndTime:        "Date in format YYYYMMDD or YYYYMM or YYYY",

	// UNVERIFIED
	NameAppKey:         "",
	NameDealerOrgCode:  "",
	NameDeviceSn:       "",
	NameFaultCode:      "",
	NameFaultName:      "",
	NameId:             "",
	NameMinuteInterval: "",
	NameOrderId:        "",
	NameOrgId:          "",
	NamePrefix:         "",
	NamePrimaryKey:     "",
	NameQueryType:      api.HelpQueryType(),
	NameSn:             "",
	NameTable:          "",
	NameTaskId:         "",
	NameUserId:         "",
	NameMenuId:         "",
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

		if sgd.IsRequired(NamePsKeys) {
			// queryMutiPointDataList - expects multiple pskeys separated by comma.
			// getDevicePointMinuteDataList - expects a single pskey.
			pskey := valueTypes.SetPsKeyString(sgd.args.PsKeys.String())
			sgd.args.PsKey = &pskey
			sgd.args.PsKeys = nil
		}

		type Parse RequestArgs
		// Store result from string
		data, err = json.Marshal(Parse{
			PsId:          sgd.args.PsId,
			PsId2:         sgd.args.PsId2,
			PsId3:         sgd.args.PsId3,
			PsIds:         sgd.args.PsIds,
			PsIdList:      sgd.args.PsIdList,
			ReportType:    sgd.args.ReportType,
			FaultTypeCode: sgd.args.FaultTypeCode,
			Size:          sgd.args.Size,
			CurPage:       sgd.args.CurPage,
			DeviceType:    sgd.args.DeviceType,
			DeviceType2:   sgd.args.DeviceType2,
			ReportId:      sgd.args.ReportId,
			CodeType:      sgd.args.CodeType,
			OrgIds:        sgd.args.OrgIds,
			TemplateId:    sgd.args.TemplateId,
			DeviceModelId: sgd.args.DeviceModelId,
			Uuid:          sgd.args.Uuid,
			UuidList:      sgd.args.UuidList,
			SetType:       sgd.args.SetType,
			Type:          sgd.args.Type,
			DataType:      sgd.args.DataType,

			// Points
			PsKeyList: sgd.args.PsKeyList,
			PsKey:     sgd.args.PsKey,
			PsKeys:    sgd.args.PsKeys,
			Point:     sgd.args.Point,
			Points:    sgd.args.Points,
			DataPoint: sgd.args.DataPoint,

			// DateTime
			DateId:         sgd.args.DateId,
			DateType:       dt,
			MonthDate:      sgd.args.MonthDate,
			BeginTime1:     sgd.args.BeginTime1,
			EndTime1:       sgd.args.EndTime1,
			StartTimeStamp: sgd.args.StartTimeStamp,
			EndTimeStamp:   sgd.args.EndTimeStamp,
			StartTime:      sgd.args.StartTime,
			EndTime:        sgd.args.EndTime,

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
			Prefix:         sgd.args.Prefix,
			PrimaryKey:     sgd.args.PrimaryKey,
			QueryType:      sgd.args.QueryType,
			Sn:             sgd.args.Sn,
			Table:          sgd.args.Table,
			TaskId:         sgd.args.TaskId,
			UserId:         sgd.args.UserId,
			MenuId:         sgd.args.MenuId,
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
			sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value)) // strings.Split(a[1], ","))
			sgd.SetPsId(value)
		case NamePsIds:
			sgd.aPsId = valueTypes.SetPsIdStrings(SplitArg(value)) // strings.Split(a[1], ","))
			val := valueTypes.SetPsIdStrings(SplitArg(value))
			sgd.args.PsIds = &val
		case NameReportType:
			val := valueTypes.SetIntegerString(value)
			sgd.args.ReportType = &val
		case NameFaultTypeCode:
			val := valueTypes.SetIntegerString(value)
			sgd.args.FaultTypeCode = &val
		case NameSize:
			val := valueTypes.SetIntegerString(value)
			sgd.args.Size = &val
		case NameCurPage:
			val := valueTypes.SetIntegerString(value)
			sgd.args.CurPage = &val
		case NameDeviceType:
			val := valueTypes.SetIntegerString(value)
			sgd.args.DeviceType = &val
		case NameDeviceType2:
			val := valueTypes.SetIntegerString(value)
			sgd.args.DeviceType2 = &val
		case NameReportId:
			val := valueTypes.SetStringValue(value)
			sgd.args.ReportId = &val
		case NameCodeType:
			val := valueTypes.SetStringValue(value)
			sgd.args.CodeType = &val
		case NameOrgIds:
			val := valueTypes.SetStringValue(value)
			sgd.args.OrgIds = &val
		case NamePsIdList:
			val := valueTypes.SetStringValue(value)
			sgd.args.PsIdList = &val
		case NameTemplateId:
			val := valueTypes.SetStringValue(value)
			sgd.args.TemplateId = &val
		case NameDeviceModelId:
			val := valueTypes.SetStringValue(value)
			sgd.args.DeviceModelId = &val
		case NameUuid:
			val := valueTypes.SetIntegerString(value)
			sgd.args.Uuid = &val
		case NameUuidList:
			val := valueTypes.SetStringValue(value)
			sgd.args.UuidList = &val
		case NameSetType:
			val := valueTypes.SetStringValue(value)
			sgd.args.SetType = &val
		case NameType:
			val := valueTypes.SetStringValue(value)
			sgd.args.Type = &val
		case NameDataType:
			val := valueTypes.SetStringValue(value)
			sgd.args.DataType = &val

		// Points
		case NamePsKeyList:
			val := valueTypes.SetStringValue(value)
			sgd.args.PsKeyList = &val
		case NamePsKeys:
			val := valueTypes.SetPsKeysString(value)
			sgd.args.PsKeys = &val
		case NamePsKey:
			val := valueTypes.SetPsKeyString(value)
			sgd.args.PsKey = &val
		case NamePointId:
			val := valueTypes.SetPointIdString(value)
			sgd.args.Point = &val
		case NamePoints:
			sgd.SetPoints(value)
			// val := valueTypes.SetPointIdsString(value); sgd.args.Points = &val
		case NameDataPoint:
			// sgd.SetPoints(value)
			val := valueTypes.SetStringValue(value)
			sgd.args.DataPoint = &val

		// DateTime
		case NameDateId:
			sgd.SetDateId(value)
		case NameMonthDate:
			sgd.SetMonthDate(value)
		case NameDay:
			sgd.SetDay(value)
		case NameBeginTime1:
			val := valueTypes.SetDateTimeString(value)
			sgd.args.BeginTime1 = &val
		case NameEndTime1:
			val := valueTypes.SetDateTimeString(value)
			sgd.args.EndTime1 = &val
		case NameStartTimeStamp:
			val := valueTypes.SetDateTimeString(value)
			sgd.args.StartTimeStamp = &val
		case NameEndTimeStamp:
			val := valueTypes.SetDateTimeString(value)
			sgd.args.EndTimeStamp = &val
		case NameStartTime:
			sgd.SetStartTime(value)
			// val := valueTypes.SetDateTimeString(value); sgd.args.StartTime = &val
		case NameEndTime:
			sgd.SetEndTime(value)
			// val := valueTypes.SetDateTimeString(value); sgd.args.EndTime = &val

		// UNVERIFIED
		case NameAppKey:
			val := valueTypes.SetStringValue(value)
			sgd.args.AppKey = &val
		case NameDealerOrgCode:
			val := valueTypes.SetStringValue(value)
			sgd.args.DealerOrgCode = &val
		case NameDeviceSn:
			val := valueTypes.SetStringValue(value)
			sgd.args.DeviceSn = &val
		case NameFaultCode:
			val := valueTypes.SetIntegerString(value)
			sgd.args.FaultCode = &val
		case NameFaultName:
			val := valueTypes.SetStringValue(value)
			sgd.args.FaultName = &val
		case NameId:
			val := valueTypes.SetIntegerString(value)
			sgd.args.Id = &val
		case NameMinuteInterval:
			val := valueTypes.SetIntegerString(value)
			sgd.args.MinuteInterval = &val
		case NameOrderId:
			val := valueTypes.SetStringValue(value)
			sgd.args.OrderId = &val
		case NameOrgId:
			val := valueTypes.SetIntegerString(value)
			sgd.args.OrgId = &val
		case NamePrefix:
			val := valueTypes.SetStringValue(value)
			sgd.args.Prefix = &val
		case NamePrimaryKey:
			val := valueTypes.SetStringValue(value)
			sgd.args.PrimaryKey = &val
		case NameQueryType:
			val := valueTypes.SetStringValue(value)
			sgd.args.QueryType = &val
		case NameSn:
			val := valueTypes.SetStringValue(value)
			sgd.args.Sn = &val
		case NameTable:
			val := valueTypes.SetStringValue(value)
			sgd.args.Table = &val
		case NameTaskId:
			val := valueTypes.SetStringValue(value)
			sgd.args.TaskId = &val
		case NameUserId:
			val := valueTypes.SetStringValue(value)
			sgd.args.UserId = &val
		case NameMenuId:
			val := valueTypes.SetIntegerString(value)
			sgd.args.MenuId = &val
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
		case NameDeviceType2:
			value = sgd.args.DeviceType2.String()
		case NameReportId:
			value = sgd.args.ReportId.String()
		case NameCodeType:
			value = sgd.args.CodeType.String()
		case NameOrgIds:
			value = sgd.args.OrgIds.String()
		case NamePsIdList:
			value = sgd.args.PsIdList.String()
		case NameTemplateId:
			value = sgd.args.TemplateId.String()
		case NameDeviceModelId:
			value = sgd.args.DeviceModelId.String()
		case NameUuid:
			value = sgd.args.Uuid.String()
		case NameUuidList:
			value = sgd.args.UuidList.String()
		case NameSetType:
			value = sgd.args.SetType.String()
		case NameType:
			value = sgd.args.Type.String()
		case NameDataType:
			value = sgd.args.DataType.String()

		// Points
		case NamePsKeyList:
			value = sgd.args.PsKeyList.String()
		case NamePsKeys:
			value = sgd.args.PsKeys.String()
		case NamePsKey:
			value = sgd.args.PsKey.String()
		case NamePointId:
			value = sgd.args.Point.String()
		case NamePoints:
			value = sgd.args.Points.String()
		case NameDataPoint:
			value = sgd.args.DataPoint.String()

		// DateTime
		case NameMonthDate:
			value = sgd.args.MonthDate.Original()
		case NameDateId:
			value = sgd.args.DateId.Original()
		case NameDay:
			value = sgd.args.Day.Original()
		case NameBeginTime1:
			value = sgd.args.BeginTime1.Original()
		case NameEndTime1:
			value = sgd.args.EndTime1.Original()
		case NameStartTimeStamp:
			value = sgd.args.StartTimeStamp.Format(valueTypes.DateTimeLayoutSecond) // .Original()
		case NameEndTimeStamp:
			value = sgd.args.EndTimeStamp.Format(valueTypes.DateTimeLayoutSecond) // .Original()
		case NameStartTime:
			value = sgd.args.StartTime.Original()
		case NameEndTime:
			value = sgd.args.EndTime.Original()

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
		case NamePrefix:
			value = sgd.args.Prefix.String()
		case NamePrimaryKey:
			value = sgd.args.PrimaryKey.String()
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
		case NameMenuId:
			value = sgd.args.MenuId.String()
		}
	}
	return value
}

func (sgd *SunGrowDataRequest) IsSet(arg string) bool {
	var ok bool
	for range Only.Once {
		switch arg {
		case NamePsId:
			// Handled differently.
			ok = true
			// if sgd.args.PsId != nil { ok = true }
		case NamePsId2:
			// Handled differently.
			ok = true
			// if sgd.args.PsId2 != nil { ok = true }
		case NamePsId3:
			// Handled differently.
			ok = true
			// if sgd.args.PsId3 != nil { ok = true }
		case NamePsIds:
			// Handled differently.
			ok = true
			// if sgd.args.PsIds != nil { ok = true }
		case NameReportType:
			if sgd.args.ReportType != nil {
				ok = true
			}
		case NameFaultTypeCode:
			if sgd.args.FaultTypeCode != nil {
				ok = true
			}
		case NameSize:
			if sgd.args.Size != nil {
				ok = true
			}
		case NameCurPage:
			if sgd.args.CurPage != nil {
				ok = true
			}
		case NameDeviceType:
			if sgd.args.DeviceType != nil {
				ok = true
			}
		case NameDeviceType2:
			if sgd.args.DeviceType2 != nil {
				ok = true
			}
		case NameReportId:
			if sgd.args.ReportId != nil {
				ok = true
			}
		case NameCodeType:
			if sgd.args.CodeType != nil {
				ok = true
			}
		case NameOrgIds:
			if sgd.args.OrgIds != nil {
				ok = true
			}
		case NamePsIdList:
			if sgd.args.PsIdList != nil {
				ok = true
			}
		case NameTemplateId:
			if sgd.args.TemplateId != nil {
				ok = true
			}
		case NameDeviceModelId:
			if sgd.args.DeviceModelId != nil {
				ok = true
			}
		case NameUuid:
			if sgd.args.Uuid != nil {
				ok = true
			}
		case NameUuidList:
			if sgd.args.UuidList != nil {
				ok = true
			}
		case NameSetType:
			if sgd.args.SetType != nil {
				ok = true
			}
		case NameType:
			if sgd.args.Type != nil {
				ok = true
			}
		case NameDataType:
			if sgd.args.DataType != nil {
				ok = true
			}

		// Points
		case NamePsKeyList:
			if sgd.args.PsKeyList != nil {
				ok = true
			}
		case NamePsKeys:
			if sgd.args.PsKeys != nil {
				ok = true
			}
		case NamePsKey:
			if sgd.args.PsKey != nil {
				ok = true
			}
		case NamePointId:
			if sgd.args.Point != nil {
				ok = true
			}
		case NamePoints:
			if sgd.args.Points != nil {
				ok = true
			}
		case NameDataPoint:
			if sgd.args.DataPoint != nil {
				ok = true
			}

		// DateTime
		case NameMonthDate:
			if sgd.args.MonthDate != nil {
				ok = true
			}
		case NameDateId:
			if sgd.args.DateId != nil {
				ok = true
			}
		case NameDay:
			if sgd.args.Day != nil {
				ok = true
			}
		case NameBeginTime1:
			if sgd.args.BeginTime1 != nil {
				ok = true
			}
		case NameEndTime1:
			if sgd.args.EndTime1 != nil {
				ok = true
			}
		case NameStartTimeStamp:
			if sgd.args.StartTimeStamp != nil {
				ok = true
			}
		case NameEndTimeStamp:
			if sgd.args.EndTimeStamp != nil {
				ok = true
			}
		case NameStartTime:
			if sgd.args.StartTime != nil {
				ok = true
			}
		case NameEndTime:
			if sgd.args.EndTime != nil {
				ok = true
			}

		// UNVERIFIED
		case NameAppKey:
			if sgd.args.AppKey != nil {
				ok = true
			}
		case NameDealerOrgCode:
			if sgd.args.DealerOrgCode != nil {
				ok = true
			}
		case NameDeviceSn:
			if sgd.args.DeviceSn != nil {
				ok = true
			}
		case NameFaultCode:
			if sgd.args.FaultCode != nil {
				ok = true
			}
		case NameFaultName:
			if sgd.args.FaultName != nil {
				ok = true
			}
		case NameId:
			if sgd.args.Id != nil {
				ok = true
			}
		case NameMinuteInterval:
			if sgd.args.MinuteInterval != nil {
				ok = true
			}
		case NameOrderId:
			if sgd.args.OrderId != nil {
				ok = true
			}
		case NameOrgId:
			if sgd.args.OrgId != nil {
				ok = true
			}
		case NamePrefix:
			if sgd.args.Prefix != nil {
				ok = true
			}
		case NamePrimaryKey:
			if sgd.args.PrimaryKey != nil {
				ok = true
			}
		case NameQueryType:
			// Handled differently.
			ok = true
			// if sgd.args.QueryType != nil { ok = true }
		case NameSn:
			if sgd.args.Sn != nil {
				ok = true
			}
		case NameTable:
			if sgd.args.Table != nil {
				ok = true
			}
		case NameTaskId:
			if sgd.args.TaskId != nil {
				ok = true
			}
		case NameUserId:
			if sgd.args.UserId != nil {
				ok = true
			}
		case NameMenuId:
			if sgd.args.MenuId != nil {
				ok = true
			}
		case NameDateType:
			// Handled differently.
			ok = true
			// if sgd.args.DateType != nil { ok = true }
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

			if sgd.IsSet(key) {
				continue
			}

			ok = false
			fmt.Printf("(required) %s - %s\n", key, Help[key])
		}
	}
	return ok
}

func (sgd *SunGrowDataRequest) GetArgs(endpoint api.EndPoint) string {
	var ret string

	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		var sorted []string
		for key := range args {
			sorted = append(sorted, key)
		}
		sort.Strings(sorted)

		for _, key := range sorted {
			if args[key] != "true" {
				continue
			}

			if key == NameDateType {
				continue // Handled differently.
			}

			if sgd.IsNotSet(key) {
				continue
			}

			ret += fmt.Sprintf("%s:%s ", key, sgd.Get(key))
		}
	}

	return ret
}

func (sgd *SunGrowDataRequest) GetArgsHash(endpoint api.EndPoint) string {
	var ret string

	for range Only.Once {
		ret = sgd.GetArgs(endpoint)
		h := fnv.New32a()
		_, _ = h.Write([]byte(ret))
		ret = fmt.Sprintf("%X", h.Sum32())
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

			if sgd.IsRequiredAndNotSet(key) {
				fmt.Printf("required - %s:value\n", key)
			}
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
			did.DateType = "4"
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

func (sgd *SunGrowDataRequest) SetTimeStamp(date string) {
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

func (sgd *SunGrowDataRequest) SetStartTime(date string) {
	for range Only.Once {
		did := valueTypes.SetDateTimeString(date)
		sgd.args.StartTime = &did
		if sgd.args.StartTime.IsZero() {
			now := time.Now().Format(valueTypes.DateLayoutDay)
			did = valueTypes.NewDateTime(now)
			sgd.args.StartTime = &did
		}
		var qt valueTypes.String
		switch len(date) {
		case 8:
			qt = qt.SetString("1")
		case 6:
			qt = qt.SetString("2")
		case 4:
			qt = qt.SetString("3")
		}
		sgd.args.QueryType = &qt
	}
}

func (sgd *SunGrowDataRequest) SetEndTime(date string) {
	for range Only.Once {
		did := valueTypes.SetDateTimeString(date)
		sgd.args.EndTime = &did
		if sgd.args.EndTime.IsZero() {
			now := time.Now().Format(valueTypes.DateLayoutDay)
			did = valueTypes.NewDateTime(now)
			sgd.args.EndTime = &did
		}
		var qt valueTypes.String
		switch len(date) {
		case 8:
			qt = qt.SetString("1")
		case 6:
			qt = qt.SetString("2")
		case 4:
			qt = qt.SetString("3")
		}
		sgd.args.QueryType = &qt
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
		art := valueTypes.SetIntegerString(rt)
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
			yes = false
			break
		}
		if sgd.IsNotRequired(arg) {
			yes = false
			break
		}
		if sgd.IsSet(arg) {
			yes = false
			break
		}
		yes = true
	}
	return yes
}

func (sgd *SunGrowDataRequest) IsRequiredAndSet(arg string) bool {
	var yes bool
	for range Only.Once {
		if _, ok := sgd.Required[arg]; !ok {
			yes = false
			break
		}
		if sgd.IsNotRequired(arg) {
			yes = false
			break
		}
		if sgd.IsNotSet(arg) {
			yes = false
			break
		}
		yes = true
	}
	return yes
}

// GetPrimaryArg - Fetch the primary arg that's set. Typically used in filename generation.
func (sgd *SunGrowDataRequest) GetPrimaryArg() string {
	var yes string
	for range Only.Once {
		switch {
		case sgd.IsRequiredAndSet(NamePsKey):
			yes = sgd.Get(NamePsKey)
		case sgd.IsRequiredAndSet(NamePsId):
			yes = sgd.Get(NamePsId)
		case sgd.IsRequiredAndSet(NamePsId2):
			yes = sgd.Get(NamePsId2)
		case sgd.IsRequiredAndSet(NamePsId3):
			yes = sgd.Get(NamePsId3)
		case sgd.IsRequiredAndSet(NameUuid):
			yes = sgd.Get(NameUuid)
		case sgd.IsRequiredAndSet(NameFaultCode):
			yes = sgd.Get(NameFaultCode)
		case sgd.IsRequiredAndSet(NameDeviceType):
			yes = sgd.Get(NameDeviceType)
		case sgd.IsRequiredAndSet(NameOrgId):
			yes = sgd.Get(NameOrgId)
		case sgd.IsRequiredAndSet(NameTaskId):
			yes = sgd.Get(NameTaskId)
		case sgd.IsRequiredAndSet(NameSn):
			yes = sgd.Get(NameSn)
		case sgd.IsRequiredAndSet(NameCodeType):
			yes = sgd.Get(NameCodeType)
		case sgd.IsRequiredAndSet(NameReportType):
			yes = sgd.Get(NameReportType)
		case sgd.IsRequiredAndSet(NameUserId):
			yes = sgd.Get(NameUserId)
		}
	}
	return yes
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
			sgd.args.PsId2 = &pid // Specifically for WebIscmAppService.getPowerStationInfo
		}
		if sgd.IsRequired(NamePsId3) {
			sgd.args.PsId3 = &pid // Specifically for WebAppService.getPsTree
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

func (sgd *SunGrowDataRequest) SetPoints(points string) {
	for range Only.Once {
		pids := valueTypes.SetPointIdsString(points)
		if pids.Error != nil {
			fmt.Printf("Error: %s - %s\n", NamePsId, sgd.args.Points.Error)
			break
		}

		// var pskeys []string
		// for _, pskey := range pids.PointIds {
		// 	pskeys = append(pskeys, pskey.PsKey.String())
		// }
		// psk := valueTypes.SetPsKeysString(strings.Join(pskeys, ","))
		psk := pids.PsKeys()

		sgd.args.Points = &pids
		if sgd.IsRequired(NamePsId) {
			p := valueTypes.SetPsIdString(psk.PsKeys[0].PsId)
			sgd.args.PsId = &p
			if len(sgd.aPsId) == 0 {
				sgd.aPsId = append(sgd.aPsId, p)
			}
		}
		if sgd.IsRequired(NamePsKey) {
			if psk.PsKeys[0].String() != "" {
				sgd.args.PsKey = &psk.PsKeys[0]
			}
		}
		if sgd.IsRequired(NamePsKeys) {
			sgd.args.PsKeys = psk
		}
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
