package getPowerDeviceSetTaskDetailList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceSetTaskDetailList"
const Disabled = false
const EndPointName = "AppService.getPowerDeviceSetTaskDetailList"

type RequestData struct {
	QueryType valueTypes.String  `json:"query_type" required:"true"`
	TaskId    valueTypes.String  `json:"task_id" required:"true"`
	Uuid      valueTypes.Integer `json:"uuid" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("query_type: 2, 3, 4")
	return ret
}

type ResultData struct {
	DeviceList []struct {
		ChannelId       valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
		CountryId       valueTypes.Integer `json:"country_id"`
		DataFlagDetail  valueTypes.Integer `json:"data_flag_detail"`
		DeviceArea      valueTypes.String  `json:"device_area"`
		DeviceAreaName  valueTypes.String  `json:"device_area_name"`
		DeviceModel     valueTypes.String  `json:"device_model"`
		DeviceModelCode valueTypes.String  `json:"device_model_code"`
		DeviceModelId   valueTypes.Integer `json:"device_model_id"`
		DeviceName      valueTypes.String  `json:"device_name"`
		DeviceProSn     valueTypes.String  `json:"device_pro_sn"`
		DeviceType      valueTypes.Integer `json:"device_type"`
		GridTypeId      valueTypes.Integer `json:"grid_type_id"`
		PsName          valueTypes.String  `json:"ps_name"`
		PsShortName     valueTypes.String  `json:"ps_short_name"`
		Sn              valueTypes.String  `json:"sn"`

		IsHaveversion  valueTypes.Bool   `json:"is_haveversion"`
		BatVersion     valueTypes.String `json:"bat_version"`
		LcdVersion     valueTypes.String `json:"lcd_version"`
		MachineVersion valueTypes.String `json:"machine_version"`
		MdspVersion    valueTypes.String `json:"mdsp_version"`
		SdspVersion    valueTypes.String `json:"sdsp_version"`
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
	} `json:"device_list" DataTable:"true" DataTableIndex:"true"`
	PageList []struct {
		// GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"PsId"`
		GoStruct.GoStruct  `json:"-" PointIdFrom:"PsId" PointIdReplace:"true"`

		PsId   valueTypes.Integer `json:"ps_id"`
		PsName valueTypes.String  `json:"ps_name"`
		ChannelId             valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		DeviceType            valueTypes.Integer  `json:"device_type"`
		Sn                    valueTypes.String   `json:"sn"`

		PointId     valueTypes.Integer `json:"point_id"`
		PointIdType valueTypes.Integer `json:"point_id_type"`
		PointName   valueTypes.String  `json:"point_name"`
		Unit        valueTypes.String  `json:"unit"`

		CommandStatus         valueTypes.Integer  `json:"command_status"`
		CountryId             valueTypes.Integer  `json:"country_id"`
		CreateTime            valueTypes.DateTime `json:"create_time"`
		DataFlagDetail        valueTypes.Integer  `json:"data_flag_detail"`
		DeviceArea            valueTypes.String   `json:"device_area"`
		DeviceAreaName        valueTypes.String   `json:"device_area_name"`
		DeviceModel           valueTypes.String   `json:"device_model"`
		DeviceModelCode       valueTypes.String   `json:"device_model_code"`
		DeviceModelId         valueTypes.Integer  `json:"device_model_id"`
		DeviceName            valueTypes.String   `json:"device_name"`
		DeviceProSn           valueTypes.String   `json:"device_pro_sn"`
		DeviceUUID            valueTypes.Integer  `json:"device_uuid"`
		GridTypeId            valueTypes.Integer  `json:"grid_type_id"`
		ModbusAddress         valueTypes.String   `json:"modbus_address"`
		Module                valueTypes.Integer  `json:"module"`
		PsShortName           valueTypes.String   `json:"ps_short_name"`
		ReturnValue           valueTypes.String   `json:"return_value"`
		SetPrecision          valueTypes.String   `json:"set_precision"`
		SetValName            valueTypes.String   `json:"set_val_name"`
		SetValNameVal         valueTypes.String   `json:"set_val_name_val"`
		SetValue              valueTypes.String   `json:"set_value"`
		TaskDetailCommandType valueTypes.String   `json:"task_detail_command_type"`
		TaskDetailId          valueTypes.Integer  `json:"task_detail_id"`
		TaskId                valueTypes.Integer  `json:"task_id"`
		TaskType              valueTypes.Integer  `json:"task_type"`
		UpdateTime            valueTypes.DateTime `json:"update_time"`
		ValType               valueTypes.String   `json:"val_type"`

		IsHaveversion  valueTypes.Bool   `json:"is_haveversion"`
		BatVersion     valueTypes.String `json:"bat_version"`
		LcdVersion     valueTypes.String `json:"lcd_version"`
		MachineVersion valueTypes.String `json:"machine_version"`
		MdspVersion    valueTypes.String `json:"mdsp_version"`
		SdspVersion    valueTypes.String `json:"sdsp_version"`
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
	} `json:"pageList" PointId:"devices" PointIdReplace:"true"`
	PsNameInfoList []struct {
		PsId        valueTypes.Integer `json:"ps_id"`
		PsName      valueTypes.String  `json:"ps_name"`
		PsShortName valueTypes.String  `json:"ps_short_name"`
	} `json:"ps_name_info_list" DataTable:"true" DataTableIndex:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
