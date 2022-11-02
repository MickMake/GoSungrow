package iSolarCloud

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


// SunGrowDataRequest - Collection of all possible request args.
type SunGrowDataRequest struct {
	// Be careful with types. If you see a general "error' response,
	// then it's more likely a type mismatch.
	PsId           *valueTypes.PsId     `json:"ps_id,omitempty"`
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
	Points         *valueTypes.Integer   `json:"points,omitempty"`
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

	aPsId valueTypes.PsIds
}

const (
	NamePsId           = "PsId"
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

// MarshalJSON - Convert value to JSON
func (sgd SunGrowDataRequest) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		var dt *string
		if sgd.DateId != nil {
			dt = &sgd.DateId.DateType
		}

		type Parse SunGrowDataRequest
		// Store result from string
		data, err = json.Marshal(Parse {
			PsId:           sgd.PsId,
			ReportType:     sgd.ReportType,
			DateId:         sgd.DateId,
			DateType:       dt,
			FaultTypeCode:  sgd.FaultTypeCode,
			Size:           sgd.Size,
			CurPage:        sgd.CurPage,
			DeviceType:     sgd.DeviceType,
			ReportId:       sgd.ReportId,
			CodeType:       sgd.CodeType,
			OrgIds:         sgd.OrgIds,
			PsIdList:       sgd.PsIdList,
			Uuid:           sgd.Uuid,
			TemplateId:     sgd.TemplateId,
			DeviceModelId:  sgd.DeviceModelId,
			PsKey:          sgd.PsKey,

			// UNVERIFIED
			AppKey:         sgd.AppKey,
			BeginTime:      sgd.BeginTime,
			DealerOrgCode:  sgd.DealerOrgCode,
			DeviceSn:       sgd.DeviceSn,
			EndTimeStamp:   sgd.EndTimeStamp,
			FaultCode:      sgd.FaultCode,
			FaultName:      sgd.FaultName,
			Id:             sgd.Id,
			MinuteInterval: sgd.MinuteInterval,
			OrderId:        sgd.OrderId,
			OrgId:          sgd.OrgId,
			PointId:        sgd.PointId,
			Points:         sgd.Points,
			Prefix:         sgd.Prefix,
			PrimaryKey:     sgd.PrimaryKey,
			PsKeyList:      sgd.PsKeyList,
			QueryType:      sgd.QueryType,
			Sn:             sgd.Sn,
			StartTimeStamp: sgd.StartTimeStamp,
			Table:          sgd.Table,
			TaskId:         sgd.TaskId,
		})
		if err == nil {
			break
		}
	}

	return data, err
}

func (sgd *SunGrowDataRequest) Set(args ...string) SunGrowDataRequest {
	var request SunGrowDataRequest
	for range Only.Once {
		for _, arg := range args {
			a := strings.Split(arg, ":")
			if len(a) == 0 {
				continue
			}

			if len(a) == 1 {
				a = append(a, "")
			}

			switch a[0] {
			case NamePsId:
				request.aPsId = valueTypes.SetPsIdStrings(SplitArg(a[1]))	// strings.Split(a[1], ","))
			case NameReportType:
				val := valueTypes.SetStringValue(a[1]); request.ReportType = &val
			case NameDateId:
				val := valueTypes.SetDateTimeString(a[1]); request.DateId = &val
			case NameFaultTypeCode:
				val := valueTypes.SetIntegerString(a[1]); request.FaultTypeCode = &val
			case NameSize:
				val := valueTypes.SetIntegerString(a[1]); request.Size = &val
			case NameCurPage:
				val := valueTypes.SetIntegerString(a[1]); request.CurPage = &val
			case NameDeviceType:
				val := valueTypes.SetStringValue(a[1]); request.DeviceType = &val
			case NameReportId:
				val := valueTypes.SetStringValue(a[1]); request.ReportId = &val
			case NameCodeType:
				val := valueTypes.SetStringValue(a[1]); request.CodeType = &val
			case NameOrgIds:
				val := valueTypes.SetStringValue(a[1]); request.OrgIds = &val
			case NamePsIdList:
				val := valueTypes.SetStringValue(a[1]); request.PsIdList = &val
			case NameUuid:
				val := valueTypes.SetStringValue(a[1]); request.Uuid = &val
			case NameTemplateId:
				val := valueTypes.SetStringValue(a[1]); request.TemplateId = &val
			case NameDeviceModelId:
				val := valueTypes.SetStringValue(a[1]); request.DeviceModelId = &val
			case NamePsKey:
				val := valueTypes.SetPsKeyValue(a[1]); request.PsKey = &val


			// UNVERIFIED
			case NameAppKey:
				val := valueTypes.SetStringValue(a[1]); request.AppKey = &val
			case NameBeginTime:
				val := valueTypes.SetDateTimeString(a[1]); request.BeginTime = &val
			case NameDealerOrgCode:
				val := valueTypes.SetStringValue(a[1]); request.DealerOrgCode = &val
			case NameDeviceSn:
				val := valueTypes.SetStringValue(a[1]); request.DeviceSn = &val
			case NameEndTimeStamp:
				val := valueTypes.SetDateTimeString(a[1]); request.EndTimeStamp = &val
			case NameFaultCode:
				val := valueTypes.SetIntegerString(a[1]); request.FaultCode = &val
			case NameFaultName:
				val := valueTypes.SetStringValue(a[1]); request.FaultName = &val
			case NameId:
				val := valueTypes.SetIntegerString(a[1]); request.Id = &val
			case NameMinuteInterval:
				val := valueTypes.SetIntegerString(a[1]); request.MinuteInterval = &val
			case NameOrderId:
				val := valueTypes.SetStringValue(a[1]); request.OrderId = &val
			case NameOrgId:
				val := valueTypes.SetStringValue(a[1]); request.OrgId = &val
			case NamePointId:
				val := valueTypes.SetPointIdString(a[1]); request.PointId = &val
			case NamePoints:
				val := valueTypes.SetIntegerString(a[1]); request.Points = &val
			case NamePrefix:
				val := valueTypes.SetStringValue(a[1]); request.Prefix = &val
			case NamePrimaryKey:
				val := valueTypes.SetStringValue(a[1]); request.PrimaryKey = &val
			case NamePsKeyList:
				val := valueTypes.SetStringValue(a[1]); request.PsKeyList = &val
			case NameQueryType:
				val := valueTypes.SetStringValue(a[1]); request.QueryType = &val
			case NameSn:
				val := valueTypes.SetStringValue(a[1]); request.Sn = &val
			case NameStartTimeStamp:
				val := valueTypes.SetDateTimeString(a[1]); request.StartTimeStamp = &val
			case NameTable:
				val := valueTypes.SetStringValue(a[1]); request.Table = &val
			case NameTaskId:
				val := valueTypes.SetStringValue(a[1]); request.TaskId = &val
			}
		}
	}
	return request
}

func (sgd *SunGrowDataRequest) Validate(endpoint api.EndPoint) bool {
	ok := true
	for range Only.Once {
		args := endpoint.GetRequestArgNames()
		for key, value := range args {
			msg := fmt.Sprintf("%s is required\n", key)

			if value != "true" {
				continue
			}
			switch key {
				case NamePsId:
					// if sgd.PsId == nil { fmt.Printf(msg); ok = false }
				case NameReportType:
					if sgd.ReportType == nil { fmt.Printf(msg); ok = false }
				case NameDateId:
					if sgd.DateId == nil { fmt.Printf(msg); ok = false }
				case NameFaultTypeCode:
					if sgd.FaultTypeCode == nil { fmt.Printf(msg); ok = false }
				case NameSize:
					if sgd.Size == nil { fmt.Printf(msg); ok = false }
				case NameCurPage:
					if sgd.CurPage == nil { fmt.Printf(msg); ok = false }
				case NameDeviceType:
					if sgd.DeviceType == nil { fmt.Printf(msg); ok = false }
				case NameReportId:
					if sgd.ReportId == nil { fmt.Printf(msg); ok = false }
				case NameCodeType:
					if sgd.CodeType == nil { fmt.Printf(msg); ok = false }
				case NameOrgIds:
					if sgd.OrgIds == nil { fmt.Printf(msg); ok = false }
				case NamePsIdList:
					if sgd.PsIdList == nil { fmt.Printf(msg); ok = false }
				case NameUuid:
					if sgd.Uuid == nil { fmt.Printf(msg); ok = false }
				case NameTemplateId:
					if sgd.TemplateId == nil { fmt.Printf(msg); ok = false }
				case NameDeviceModelId:
					if sgd.DeviceModelId == nil { fmt.Printf(msg); ok = false }
				case NamePsKey:
					if sgd.PsKey == nil { fmt.Printf(msg); ok = false }


				// UNVERIFIED
				case NameAppKey:
					if sgd.AppKey == nil { fmt.Printf(msg); ok = false }
				case NameBeginTime:
					if sgd.BeginTime == nil { fmt.Printf(msg); ok = false }
				case NameDealerOrgCode:
					if sgd.DealerOrgCode == nil { fmt.Printf(msg); ok = false }
				case NameDeviceSn:
					if sgd.DeviceSn == nil { fmt.Printf(msg); ok = false }
				case NameEndTimeStamp:
					if sgd.EndTimeStamp == nil { fmt.Printf(msg); ok = false }
				case NameFaultCode:
					if sgd.FaultCode == nil { fmt.Printf(msg); ok = false }
				case NameFaultName:
					if sgd.FaultName == nil { fmt.Printf(msg); ok = false }
				case NameId:
					if sgd.Id == nil { fmt.Printf(msg); ok = false }
				case NameMinuteInterval:
					if sgd.MinuteInterval == nil { fmt.Printf(msg); ok = false }
				case NameOrderId:
					if sgd.OrderId == nil { fmt.Printf(msg); ok = false }
				case NameOrgId:
					if sgd.OrgId == nil { fmt.Printf(msg); ok = false }
				case NamePointId:
					if sgd.PointId == nil { fmt.Printf(msg); ok = false }
				case NamePoints:
					if sgd.Points == nil { fmt.Printf(msg); ok = false }
				case NamePrefix:
					if sgd.Prefix == nil { fmt.Printf(msg); ok = false }
				case NamePrimaryKey:
					if sgd.PrimaryKey == nil { fmt.Printf(msg); ok = false }
				case NamePsKeyList:
					if sgd.PsKeyList == nil { fmt.Printf(msg); ok = false }
				case NameQueryType:
					if sgd.QueryType == nil { fmt.Printf(msg); ok = false }
				case NameSn:
					if sgd.Sn == nil { fmt.Printf(msg); ok = false }
				case NameStartTimeStamp:
					if sgd.StartTimeStamp == nil { fmt.Printf(msg); ok = false }
				case NameTable:
					if sgd.Table == nil { fmt.Printf(msg); ok = false }
				case NameTaskId:
					if sgd.TaskId == nil { fmt.Printf(msg); ok = false }
			}
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

func (sgd *SunGrowDataRequest) GetFilename(prefix string) string {
	var ret string
	for range Only.Once {
		// aret := []string{ prefix }
		// if sgd.PsId != nil {
		// 	a := sgd.PsId.String()
		// 	if a != "" { aret = append(aret, a) }
		// }
		//
		// if sgd.DateId != nil {
		// 	a := sgd.DateId.Original()
		// 	if a != "" { aret = append(aret, a) }
		// }
		//
		// if sgd.ReportType != nil {
		// 	a := sgd.ReportType.String()
		// 	if a != "" { aret = append(aret, a) }
		// }
		//
		// if sgd.FaultTypeCode != nil {
		// 	a := sgd.FaultTypeCode.String()
		// 	if a != "" { aret = append(aret, a) }
		// }
		//
		// if sgd.Uuid != nil {
		// 	a := sgd.Uuid.String()
		// 	if a != "" { aret = append(aret, a) }
		// }

		ret += prefix
		j, err := json.Marshal(*sgd)
		if err != nil {
			break
		}
		ret = string(j)
		ret = strings.TrimPrefix(ret, "{")
		ret = strings.TrimSuffix(ret, "}")
		ret = strings.ReplaceAll(ret, `"`, ``)
		ret = strings.ReplaceAll(ret, `,`, `-`)
		// ret = strings.Join(aret, "-")
	}
	return ret
}

func (sgd *SunGrowDataRequest) SetDate(date string) {
	did := valueTypes.SetDateTimeString(date)
	sgd.DateId = &did
	if sgd.DateId.IsZero() {
		did = valueTypes.NewDateTime(valueTypes.Now)
		sgd.DateId = &did
	}
}

func (sgd *SunGrowDataRequest) SetFaultTypeCode(ftc string) {
	aftc := valueTypes.SetIntegerString(ftc)
	sgd.FaultTypeCode = &aftc
}

func (sgd *SunGrowDataRequest) SetReportType(rt string) {
	art := valueTypes.SetStringValue(rt)
	sgd.ReportType = &art
}

func (sgd *SunGrowDataRequest) SetPsIds(psIds []string) {
	sgd.aPsId = valueTypes.SetPsIdStrings(psIds)
}

func (sgd *SunGrowDataRequest) GetPsIds() valueTypes.PsIds {
	return sgd.aPsId
}

func (sgd *SunGrowDataRequest) SetPsId(psId string) {
	if psId == "" {
		sgd.PsId = nil
		return
	}
	pid := valueTypes.SetPsIdString(psId)
	sgd.PsId = &pid
}
