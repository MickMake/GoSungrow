package getPsDetailWithPsType

import (
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


type ResultData   struct {
	BatteryLevelPercent         string `json:"battery_level_percent"`
	ChargingDischargingPowerMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"charging_discharging_power_map"`
	Co2ReduceTotal struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"co2_reduce_total"`
	CoalReduceTotal struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"coal_reduce_total"`
	ConnectType string `json:"connect_type"`
	CurrPower   struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"curr_power"`
	DesignCapacity struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"design_capacity"`
	EnergyScheme           interface{} `json:"energy_scheme"`
	GcjLatitude            string      `json:"gcj_latitude"`
	GcjLongitude           string      `json:"gcj_longitude"`
	HasAmmeter             int64       `json:"has_ammeter"`
	HouseholdInverterData  interface{} `json:"household_inverter_data"`
	InstallerPsFaultStatus string      `json:"installer_ps_fault_status"`
	IsHaveEsInverter       int64       `json:"is_have_es_inverter"`
	IsSingleInverter       int64       `json:"is_single_inverter"`
	IsTransformSystem      string      `json:"is_transform_system"`
	Latitude               float64     `json:"latitude"`
	LoadPowerMap           struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"load_power_map"`
	LoadPowerMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"load_power_map_virgin"`
	Longitude        float64 `json:"longitude"`
	MapLatitude      string  `json:"map_latitude"`
	MapLongitude     string  `json:"map_longitude"`
	MeterReduceTotal struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"meter_reduce_total"`
	MobleTel    string `json:"moble_tel"`
	MonthEnergy struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"month_energy"`
	MonthEnergyVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"month_energy_virgin"`
	MonthIncome struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"month_income"`
	NegativeLoadMsg    interface{} `json:"negative_load_msg"`
	OwnerPsFaultStatus string      `json:"owner_ps_fault_status"`
	P83081Map          struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83081_map"`
	P83081MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83081_map_virgin"`
	P83102Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83102_map"`
	P83102MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83102_map_virgin"`
	P83102Percent string `json:"p83102_percent"`
	P83118Map     struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83118_map"`
	P83118MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83118_map_virgin"`
	P83119Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83119_map"`
	P83119MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83119_map_virgin"`
	P83120Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83120_map"`
	P83120MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83120_map_virgin"`
	P83122    string `json:"p83122"`
	P83124Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83124_map"`
	P83124MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83124_map_virgin"`
	P83202Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83202_map"`
	P83202MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83202_map_virgin"`
	P83532MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83532_map_virgin"`
	PowerChargeSetted int64 `json:"power_charge_setted"`
	PowerGridPowerMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"power_grid_power_map"`
	PowerGridPowerMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"power_grid_power_map_virgin"`
	PsCountryID    int64  `json:"ps_country_id"`
	PsDeviceType   string `json:"ps_device_type"`
	PsFaultStatus  string `json:"ps_fault_status"`
	PsHealthStatus string `json:"ps_health_status"`
	PsLocation     string `json:"ps_location"`
	PsName         string `json:"ps_name"`
	PsPsKey        string `json:"ps_ps_key"`
	PsState        string `json:"ps_state"`
	PsType         int64  `json:"ps_type"`
	PvPowerMap     struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"pv_power_map"`
	PvPowerMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"pv_power_map_virgin"`
	RobotNumSweepCapacity struct {
		Num           int64 `json:"num"`
		SweepCapacity float64 `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`
	SelfConsumptionOffsetReminder int64 `json:"self_consumption_offset_reminder"`
	So2ReduceTotal                struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"so2_reduce_total"`
	StorageInverterData []struct {
		CommunicationDevSn      string   `json:"communication_dev_sn"`
		DevStatus               int64    `json:"dev_status"`
		DeviceCode              int64    `json:"device_code"`
		DeviceModelCode         string   `json:"device_model_code"`
		DeviceName              string   `json:"device_name"`
		DeviceState             string   `json:"device_state"`
		DeviceType              int64    `json:"device_type"`
		DrmStatus               string   `json:"drm_status"`
		DrmStatusName           string   `json:"drm_status_name"`
		EnergyFlow              []string `json:"energy_flow"`
		HasAmmeter              int64    `json:"has_ammeter"`
		InstallerDevFaultStatus int64    `json:"installer_dev_fault_status"`
		InverterSn              string   `json:"inverter_sn"`
		OwnerDevFaultStatus     int64    `json:"owner_dev_fault_status"`
		P13003Map               struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13003_map"`
		P13003MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13003_map_virgin"`
		P13119Map struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13119_map"`
		P13119MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13119_map_virgin"`
		P13121Map struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13121_map"`
		P13121MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13121_map_virgin"`
		P13126Map struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13126_map"`
		P13126MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13126_map_virgin"`
		P13141    string `json:"p13141"`
		P13149Map struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13149_map"`
		P13149MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13149_map_virgin"`
		P13150Map struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13150_map"`
		P13150MapVirgin struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"p13150_map_virgin"`
		PsKey string `json:"ps_key"`
		UUID  int64  `json:"uuid"`
	} `json:"storage_inverter_data"`
	TodayEnergy struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"today_energy"`
	TodayEnergyVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"today_energy_virgin"`
	TodayIncome struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"today_income"`
	TotalEnergy struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"total_energy"`
	TotalEnergyVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"total_energy_virgin"`
	TotalIncome struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"total_income"`
	TreeReduceTotal struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"tree_reduce_total"`
	ValidFlag    int64   `json:"valid_flag"`
	WgsLatitude  float64 `json:"wgs_latitude"`
	WgsLongitude float64 `json:"wgs_longitude"`
	ZfzyMap      struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zfzy_map"`
	ZfzyMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zfzy_map_virgin"`
	ZjzzMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zjzz_map"`
	ZjzzMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zjzz_map_virgin"`
}
