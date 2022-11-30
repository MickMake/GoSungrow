package getPsValue

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsValue"
const Disabled = false

type RequestData struct {
	Size    valueTypes.Integer `json:"size" required:"true"`
	CurPage valueTypes.Integer `json:"curPage" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" PointIdFromChild:"PsId" PointIdReplace:"true"`

	PsId          valueTypes.PsId     `json:"ps_id"`
	Id            valueTypes.Integer  `json:"id"`
	Name          valueTypes.String   `json:"name"`
	PsName        valueTypes.String   `json:"ps_name"`
	PsStatus      valueTypes.Integer  `json:"ps_status"`
	PsType        valueTypes.Integer  `json:"ps_type"`
	State         valueTypes.Bool     `json:"state"`
	UpdateTime    valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
	PsFaultStatus valueTypes.Integer  `json:"ps_fault_status"`

	ArrearsStatus  valueTypes.Integer `json:"arrears_status"`
	DesignCapacity struct {
		Unit  valueTypes.String `json:"unit" PointIgnore:"true"`
		Value valueTypes.Float  `json:"value" PointUnitFrom:"Unit"`
	} `json:"design_capacity"`
	DesignCapacityOriginal valueTypes.Float    `json:"design_capacity_original"`
	DeviceStatusUpdateTime valueTypes.DateTime `json:"device_status_update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
	EquivalentHour         valueTypes.Integer  `json:"equivalent_hour"`
	GcjLatitude            valueTypes.Float    `json:"gcj_latitude"`
	GcjLongitude           valueTypes.Float    `json:"gcj_longitude"`
	InstallerPsFaultStatus valueTypes.Integer  `json:"installer_ps_fault_status"`
	InverterLoadSum        interface{}         `json:"inverter_load_sum"`
	Load                   valueTypes.Float    `json:"load"`
	MaxPr                  valueTypes.Float    `json:"max_pr"`
	MinPr                  valueTypes.Float    `json:"min_pr"`
	OrderNum               valueTypes.Integer  `json:"order_num"`
	OwnerPsFaultStatus     valueTypes.Integer  `json:"owner_ps_fault_status"`
	PowerType              valueTypes.Integer  `json:"power_type"`
	PrramPr                valueTypes.Float    `json:"prram_pr"`
	PrramPrValue           valueTypes.Float    `json:"prram_pr_value"`
	PsCountryID            valueTypes.Integer  `json:"ps_country_id"`
	RadiationMax           valueTypes.Float    `json:"radiation_max"`
	Res                    valueTypes.Integer  `json:"res"`
	ShareType              valueTypes.String   `json:"share_type"`
	Wgs84Latitude          valueTypes.Float    `json:"wgs84_latitude"`
	Wgs84Longitude         valueTypes.Float    `json:"wgs84_longitude"`
	RowCount               valueTypes.Integer  `json:"rowCount"`

	P83012  valueTypes.Float `json:"p83012"`
	P83013  valueTypes.Float `json:"p83013"`
	P83020  valueTypes.Float `json:"p83020"`
	P83022  valueTypes.Float `json:"p83022"`
	P83022y valueTypes.Float `json:"p83022y"`
	P83023  valueTypes.Float `json:"p83023"`
	P83024  valueTypes.Float `json:"p83024"`
	P83025  valueTypes.Float `json:"p83025"`
	P83033  valueTypes.Float `json:"p83033"`
	P83046  valueTypes.Float `json:"p83046"`
	P83048  valueTypes.Float `json:"p83048"`
	P83049  valueTypes.Float `json:"p83049"`
	P83050  valueTypes.Float `json:"p83050"`
	P83051  valueTypes.Float `json:"p83051"`
	P83067  valueTypes.Float `json:"p83067"`
	P83068  valueTypes.Float `json:"p83068"`
	P83069  valueTypes.Float `json:"p83069"`
	P83070  valueTypes.Float `json:"p83070"`
	P83202  valueTypes.Float `json:"p83202"`
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
