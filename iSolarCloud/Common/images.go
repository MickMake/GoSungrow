package Common

import (
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)


// PowerStationImage - `json:"images" PointArrayFlatten:"false"`
type PowerStationImage struct {
	FileId      valueTypes.Integer `json:"file_id"`
	Id          valueTypes.Integer `json:"id"`
	PicLanguage valueTypes.Integer `json:"pic_language"`
	PicType     valueTypes.Integer `json:"pic_type"`
	PictureName valueTypes.String  `json:"picture_name"`
	PictureURL  valueTypes.String  `json:"picture_url"`
	PsId        valueTypes.PsId    `json:"ps_id"`
	PsUnitUUID  interface{}        `json:"ps_unit_uuid"`
}
type PowerStationImages []PowerStationImage


// PsDirectOrgList - `json:"ps_direct_org_list" PointArrayFlatten:"false"`
type PsDirectOrgList []struct {
	OrgId        valueTypes.Integer `json:"org_id"`
	OrgIndexCode valueTypes.String  `json:"org_index_code"`
	OrgName      valueTypes.String  `json:"org_name"`
}


// PsOrgInfo - `json:"ps_org_info" PointArrayFlatten:"false"`
type PsOrgInfo []struct {
	DealerOrgCode   valueTypes.String  `json:"dealer_org_code"`
	Installer       valueTypes.String  `json:"installer"`
	InstallerEmail  valueTypes.String  `json:"installer_email"`
	InstallerPhone  valueTypes.String  `json:"installer_phone"`
	OrgId           valueTypes.Integer `json:"org_id"`
	OrgIndexCode    valueTypes.String  `json:"org_index_code"`
	OrgName         valueTypes.String  `json:"org_name"`
	PsDealerOrgCode valueTypes.String  `json:"ps_dealer_org_code"`
	UpOrgId         valueTypes.Integer `json:"up_org_id"`
}


// SelectedOrgList - `json:"selectedOrgList" PointId:"selected_org_list" PointArrayFlatten:"false"`
type SelectedOrgList []struct {
	OrgId        valueTypes.Integer `json:"org_id"`
	OrgIndexCode valueTypes.String  `json:"org_index_code"`
	OrgName      valueTypes.String  `json:"org_name"`
}


// SnDetailList - `json:"sn_detail_list" PointArrayFlatten:"false"`
type SnDetailList []struct {
	CommunicateDeviceType     valueTypes.Integer `json:"communicate_device_type"`
	CommunicateDeviceTypeName valueTypes.String  `json:"communicate_device_type_name"`
	Id                        valueTypes.Integer `json:"id"`
	IsEnable                  valueTypes.Bool    `json:"is_enable"`
	Sn                        valueTypes.String  `json:"sn" PointName:"Serial Number"`
}


// ReportInfo - `json:"info" PointArrayFlatten:"false"`
type ReportInfo []struct {
	DesignCapacity         valueTypes.Float   `json:"design_capacity" PointUnit:"W"`
	InstallerPsFaultStatus valueTypes.Integer `json:"installer_ps_fault_status"`
	OwnerPsFaultStatus     valueTypes.Integer `json:"owner_ps_fault_status"`
	PsFaultStatus          valueTypes.Integer `json:"ps_fault_status"`
	PsId                   valueTypes.PsId    `json:"ps_id"`
	PsName                 valueTypes.String  `json:"ps_name"`
	PsStatus               valueTypes.Integer `json:"ps_status"`
	PsType                 valueTypes.Integer `json:"ps_type"`
	PsTypeName             valueTypes.String  `json:"ps_type_name"`
	SysScheme              valueTypes.Integer `json:"sys_scheme"`
	SysSchemeName          valueTypes.String  `json:"sys_scheme_name"`
	ValidFlag              valueTypes.Bool    `json:"valid_flag"`
}

// type foo queryBatteryBoardsList.Disabled
// type fff queryDevicePointMinuteDataList.Disabled
