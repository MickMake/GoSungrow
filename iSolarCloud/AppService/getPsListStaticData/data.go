package getPsListStaticData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsListStaticData"
const Disabled = false
const EndPointName = "AppService.getPsListStaticData"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		GoStruct               GoStruct.GoStruct   `json:"GoStruct" PointIdFrom:"PsId" PointIdReplace:"true" PointDeviceFrom:"PsId"`

		PsId                   valueTypes.PsId     `json:"ps_id"`
		PsName                 valueTypes.String   `json:"ps_name"`
		PsShortName            valueTypes.String   `json:"ps_short_name"`
		PsStatus               valueTypes.Bool     `json:"ps_status"`
		PsType                 valueTypes.Integer  `json:"ps_type"`

		AreaId                 interface{}         `json:"area_id"`
		DesignCapacity         valueTypes.Float    `json:"design_capacity" PointUnit:"W"`
		GprsLatitude           valueTypes.Float    `json:"gprs_latitude"`
		GprsLongitude          valueTypes.Float    `json:"gprs_longitude"`
		InstallDate            valueTypes.DateTime `json:"install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
		InstallerPsFaultStatus valueTypes.Integer  `json:"installer_ps_fault_status"`
		Latitude               valueTypes.Float    `json:"latitude"`
		Location               valueTypes.String   `json:"location"`
		Longitude              valueTypes.Float    `json:"longitude"`
		MapLatitude            valueTypes.Float    `json:"map_latitude"`
		MapLongitude           valueTypes.Float    `json:"map_longitude"`
		OwnerPsFaultStatus     valueTypes.Integer  `json:"owner_ps_fault_status"`
		PsFaultStatus          valueTypes.Integer  `json:"ps_fault_status"`
		ValidFlag              valueTypes.Bool     `json:"valid_flag"`
		WaitAssignOrderCount   valueTypes.Integer  `json:"wait_assign_order_count"`
	} `json:"pageList" PointId:"page_list" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
