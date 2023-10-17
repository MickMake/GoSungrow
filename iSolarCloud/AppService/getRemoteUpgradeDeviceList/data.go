package getRemoteUpgradeDeviceList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/getRemoteUpgradeDeviceList"
	Disabled     = false
	EndPointName = "AppService.getRemoteUpgradeDeviceList"
)

type RequestData struct {
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	MModuleNameList []valueTypes.String `json:"m_module_name_list" PointId:"module_names"`
	PageList        []struct {
		GoStruct.GoStructParent `json:"-" PointIdFromChild:"PsKey" PointIdReplace:"true"`

		PsKey                valueTypes.PsKey   `json:"ps_key"`
		PsId                 valueTypes.PsId    `json:"ps_id"`
		DeviceType           valueTypes.Integer `json:"device_type"`
		DeviceCode           valueTypes.Integer `json:"device_code"`
		ChannelId            valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
		Sn                   valueTypes.String  `json:"sn" PointName:"Serial Number"`
		PsName               valueTypes.String  `json:"ps_name"`
		PsShortName          valueTypes.String  `json:"ps_short_name"`
		CommunicationDevUUId valueTypes.Integer `json:"communication_dev_uuid"`
		CommunicationModel   valueTypes.Integer `json:"communication_model"`
		CountryId            valueTypes.Integer `json:"country_id"`
		DataFlagDetail       valueTypes.Integer `json:"data_flag_detail"`
		DevStatus            valueTypes.Integer `json:"dev_status"`
		DeviceArea           valueTypes.String  `json:"device_area"`
		DeviceModel          valueTypes.String  `json:"device_model"`
		DeviceModelId        valueTypes.Integer `json:"device_model_id"`
		DeviceName           valueTypes.String  `json:"device_name"`
		DeviceSn             valueTypes.String  `json:"device_sn"`
		DeviceTypeName       valueTypes.String  `json:"device_type_name"`
		GridTypeId           valueTypes.Integer `json:"grid_type_id"`
		LoggerCode           valueTypes.Integer `json:"logger_code"`
		UUID                 valueTypes.Integer `json:"uuid"`

		IsEnable            valueTypes.Bool   `json:"is_enable"`
		IsHaveVersion       valueTypes.Bool   `json:"is_haveversion" PointId:"is_have_version"`
		ArmVersion          valueTypes.String `json:"arm_version"`
		BatVersion          valueTypes.String `json:"bat_version"`
		BatteryVersion      valueTypes.String `json:"battery_version"`
		CpldVersion         valueTypes.String `json:"cpld_version"`
		FirmwareVersionInfo struct {
			BatVersion  valueTypes.String `json:"bat_version"`
			LcdVersion  valueTypes.String `json:"lcd_version"`
			MdspVersion valueTypes.String `json:"mdsp_version"`
			SdspVersion valueTypes.String `json:"sdsp_version"`
		} `json:"firmware_version_info"`
		LcdVersion     valueTypes.String `json:"lcd_version"`
		MVersion       valueTypes.String `json:"m_version"`
		McuVersion     valueTypes.String `json:"mcu_version"`
		MdspVersion    valueTypes.String `json:"mdsp_version"`
		PvdVersion     valueTypes.String `json:"pvd_version"`
		SdspVersion    valueTypes.String `json:"sdsp_version"`
		SystemVersion  valueTypes.String `json:"system_version"`
		TempVersion    valueTypes.String `json:"temp_version"`
		UpgradeVersion valueTypes.String `json:"upgrade_version"`
		Version1       valueTypes.String `json:"version1"`
		Version2       valueTypes.String `json:"version2"`
		Version3       valueTypes.String `json:"version3"`
		Version4       valueTypes.String `json:"version4"`
		Version5       valueTypes.String `json:"version5"`
		Version6       valueTypes.String `json:"version6"`
		Version7       valueTypes.String `json:"version7"`
		Version8       valueTypes.String `json:"version8"`
		Version9       valueTypes.String `json:"version9"`
		Version10      valueTypes.String `json:"version10"`
		Version11      valueTypes.String `json:"version11"`
		Version12      valueTypes.String `json:"version12"`
	} `json:"pageList" PointId:"devices" PointIdReplace:"true" DataTable:"true" DataTableSortOn:"PsKey"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
