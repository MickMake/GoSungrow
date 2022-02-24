package getPsDetailWithPsType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPsDetailWithPsType"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BatteryLevelPercent         string        `json:"battery_level_percent"`
	ChargingDischargingPowerMap api.UnitValue `json:"charging_discharging_power_map"`
	Co2ReduceTotal              api.UnitValue `json:"co2_reduce_total"`
	CoalReduceTotal             api.UnitValue `json:"coal_reduce_total"`
	ConnectType                 string        `json:"connect_type"`
	CurrPower                   api.UnitValue `json:"curr_power"`
	DesignCapacity              api.UnitValue `json:"design_capacity"`
	EnergyScheme                interface{}   `json:"energy_scheme"`
	GcjLatitude                 string        `json:"gcj_latitude"`
	GcjLongitude                string        `json:"gcj_longitude"`
	HasAmmeter                  int64         `json:"has_ammeter"`
	HouseholdInverterData       interface{}   `json:"household_inverter_data"`
	InstallerPsFaultStatus      string        `json:"installer_ps_fault_status"`
	IsHaveEsInverter            int64         `json:"is_have_es_inverter"`
	IsSingleInverter            int64         `json:"is_single_inverter"`
	IsTransformSystem           string        `json:"is_transform_system"`
	Latitude                    float64       `json:"latitude"`
	LoadPowerMap                api.UnitValue `json:"load_power_map"`
	LoadPowerMapVirgin          api.UnitValue `json:"load_power_map_virgin"`
	Longitude                   float64       `json:"longitude"`
	MapLatitude                 string        `json:"map_latitude"`
	MapLongitude                string        `json:"map_longitude"`
	MeterReduceTotal            api.UnitValue `json:"meter_reduce_total"`
	MobleTel                    string        `json:"moble_tel"`
	MonthEnergy                 api.UnitValue `json:"month_energy"`
	MonthEnergyVirgin           api.UnitValue `json:"month_energy_virgin"`
	MonthIncome                 api.UnitValue `json:"month_income"`
	NegativeLoadMsg             interface{}   `json:"negative_load_msg"`
	OwnerPsFaultStatus          string        `json:"owner_ps_fault_status"`
	P83081Map                   api.UnitValue `json:"p83081_map"`
	P83081MapVirgin             api.UnitValue `json:"p83081_map_virgin"`
	P83102Map                   api.UnitValue `json:"p83102_map"`
	P83102MapVirgin             api.UnitValue `json:"p83102_map_virgin"`
	P83102Percent               string        `json:"p83102_percent"`
	P83118Map                   api.UnitValue `json:"p83118_map"`
	P83118MapVirgin             api.UnitValue `json:"p83118_map_virgin"`
	P83119Map                   api.UnitValue `json:"p83119_map"`
	P83119MapVirgin             api.UnitValue `json:"p83119_map_virgin"`
	P83120Map                   api.UnitValue `json:"p83120_map"`
	P83120MapVirgin             api.UnitValue `json:"p83120_map_virgin"`
	P83122                      string        `json:"p83122"`
	P83124Map                   api.UnitValue `json:"p83124_map"`
	P83124MapVirgin             api.UnitValue `json:"p83124_map_virgin"`
	P83202Map                   api.UnitValue `json:"p83202_map"`
	P83202MapVirgin             api.UnitValue `json:"p83202_map_virgin"`
	P83532MapVirgin             api.UnitValue `json:"p83532_map_virgin"`
	PowerChargeSetted           int64         `json:"power_charge_setted"`
	PowerGridPowerMap           api.UnitValue `json:"power_grid_power_map"`
	PowerGridPowerMapVirgin     api.UnitValue `json:"power_grid_power_map_virgin"`
	PsCountryID                 int64         `json:"ps_country_id"`
	PsDeviceType                string        `json:"ps_device_type"`
	PsFaultStatus               string        `json:"ps_fault_status"`
	PsHealthStatus              string        `json:"ps_health_status"`
	PsLocation                  string        `json:"ps_location"`
	PsName                      string        `json:"ps_name"`
	PsPsKey                     string        `json:"ps_ps_key"`
	PsState                     string        `json:"ps_state"`
	PsType                      int64         `json:"ps_type"`
	PvPowerMap                  api.UnitValue `json:"pv_power_map"`
	PvPowerMapVirgin            api.UnitValue `json:"pv_power_map_virgin"`
	RobotNumSweepCapacity       struct {
		Num           int64   `json:"num"`
		SweepCapacity float64 `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`
	SelfConsumptionOffsetReminder int64         `json:"self_consumption_offset_reminder"`
	So2ReduceTotal                api.UnitValue `json:"so2_reduce_total"`
	StorageInverterData           []struct {
		CommunicationDevSn      string        `json:"communication_dev_sn"`
		DevStatus               int64         `json:"dev_status"`
		DeviceCode              int64         `json:"device_code"`
		DeviceModelCode         string        `json:"device_model_code"`
		DeviceName              string        `json:"device_name"`
		DeviceState             string        `json:"device_state"`
		DeviceType              int64         `json:"device_type"`
		DrmStatus               string        `json:"drm_status"`
		DrmStatusName           string        `json:"drm_status_name"`
		EnergyFlow              []string      `json:"energy_flow"`
		HasAmmeter              int64         `json:"has_ammeter"`
		InstallerDevFaultStatus int64         `json:"installer_dev_fault_status"`
		InverterSn              string        `json:"inverter_sn"`
		OwnerDevFaultStatus     int64         `json:"owner_dev_fault_status"`
		P13003Map               api.UnitValue `json:"p13003_map"`
		P13003MapVirgin         api.UnitValue `json:"p13003_map_virgin"`
		P13119Map               api.UnitValue `json:"p13119_map"`
		P13119MapVirgin         api.UnitValue `json:"p13119_map_virgin"`
		P13121Map               api.UnitValue `json:"p13121_map"`
		P13121MapVirgin         api.UnitValue `json:"p13121_map_virgin"`
		P13126Map               api.UnitValue `json:"p13126_map"`
		P13126MapVirgin         api.UnitValue `json:"p13126_map_virgin"`
		P13141                  string        `json:"p13141"`
		P13149Map               api.UnitValue `json:"p13149_map"`
		P13149MapVirgin         api.UnitValue `json:"p13149_map_virgin"`
		P13150Map               api.UnitValue `json:"p13150_map"`
		P13150MapVirgin         api.UnitValue `json:"p13150_map_virgin"`
		PsKey                   string        `json:"ps_key"`
		UUID                    int64         `json:"uuid"`
	} `json:"storage_inverter_data"`
	TodayEnergy       api.UnitValue `json:"today_energy"`
	TodayEnergyVirgin api.UnitValue `json:"today_energy_virgin"`
	TodayIncome       api.UnitValue `json:"today_income"`
	TotalEnergy       api.UnitValue `json:"total_energy"`
	TotalEnergyVirgin api.UnitValue `json:"total_energy_virgin"`
	TotalIncome       api.UnitValue `json:"total_income"`
	TreeReduceTotal   api.UnitValue `json:"tree_reduce_total"`
	ValidFlag         int64         `json:"valid_flag"`
	WgsLatitude       float64       `json:"wgs_latitude"`
	WgsLongitude      float64       `json:"wgs_longitude"`
	ZfzyMap           api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin     api.UnitValue `json:"zfzy_map_virgin"`
	ZjzzMap           api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin     api.UnitValue `json:"zjzz_map_virgin"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}