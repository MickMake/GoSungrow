package getPsList

import (
	"GoSungrow/iSolarCloud/Common"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/powerStationService/getPsList"
const Disabled = false

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
		Images                    Common.PowerStationImages        `json:"images" PointArrayFlatten:"false"`

		AlarmCount                valueTypes.Count     `json:"alarm_count" PointId:"alarm_count" PointUpdateFreq:"UpdateFreqBoot"`
		AlarmDevCount             valueTypes.Count     `json:"alarm_dev_count" PointId:"alarm_dev_count" PointUpdateFreq:"UpdateFreqBoot"`
		AreaId                    interface{}          `json:"area_id" PointId:"area_id"`
		AreaType                  valueTypes.Integer   `json:"area_type" PointId:"area_type"`
		ArrearsStatus             valueTypes.Integer   `json:"arrears_status" PointId:"arrears_status"`
		BuildDate                 valueTypes.DateTime  `json:"build_date" PointId:"build_date" PointUpdateFreq:"UpdateFreqBoot"`
		BuildStatus               valueTypes.Integer   `json:"build_status" PointId:"build_status" PointUpdateFreq:"UpdateFreqBoot"`
		Co2Reduce                 valueTypes.UnitValue `json:"co2_reduce" PointId:"co2_reduce"`
		Co2ReduceTotal            valueTypes.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
		CurrPower                 valueTypes.UnitValue `json:"curr_power" PointId:"curr_power"`
		DailyIrradiation          valueTypes.UnitValue `json:"daily_irradiation" PointId:"daily_irradiation" PointUpdateFreq:"UpdateFreqDay"`
		DailyIrradiationVirgin    valueTypes.Float     `json:"daily_irradiation_virgin" PointIgnore:"true"`
		DesignCapacity            valueTypes.Float     `json:"design_capacity" PointId:"design_capacity" PointUnitFrom:"design_capacity_unit"`
		DesignCapacityUnit        valueTypes.String    `json:"design_capacity_unit" PointId:"design_capacity_unit"`
		DesignCapacityVirgin      valueTypes.Float     `json:"design_capacity_virgin" PointIgnore:"true"`
		EquivalentHour            valueTypes.UnitValue `json:"equivalent_hour" PointId:"equivalent_hour" PointUpdateFreq:"UpdateFreqDay"`
		ExpectInstallDate         valueTypes.DateTime  `json:"expect_install_date" PointId:"expect_install_date"`
		FaultAlarmOfflineDevCount valueTypes.Count     `json:"fault_alarm_offline_dev_count" PointId:"fault_alarm_offline_dev_count"`
		FaultCount                valueTypes.Count     `json:"fault_count" PointId:"fault_count"`
		FaultDevCount             valueTypes.Count     `json:"fault_dev_count" PointId:"fault_dev_count"`
		GcjLatitude               valueTypes.Float     `json:"gcj_latitude" PointId:"gcj_latitude"`
		GcjLongitude              valueTypes.Float     `json:"gcj_longitude" PointId:"gcj_longitude"`
		GprsLatitude              valueTypes.Float     `json:"gprs_latitude" PointId:"gprs_latitude"`
		GprsLongitude             valueTypes.Float     `json:"gprs_longitude" PointId:"gprs_longitude"`
		InstallDate               valueTypes.DateTime  `json:"install_date" PointId:"install_date"`
		InstalledPowerMap         valueTypes.UnitValue `json:"installed_power_map" PointId:"installed_power_map"`
		InstalledPowerVirgin      valueTypes.Float     `json:"installed_power_virgin" PointIgnore:"true"`
		InstallerAlarmCount       valueTypes.Count     `json:"installer_alarm_count" PointId:"installer_alarm_count"`
		InstallerFaultCount       valueTypes.Count     `json:"installer_fault_count" PointId:"installer_fault_count"`
		InstallerPsFaultStatus    valueTypes.Integer   `json:"installer_ps_fault_status" PointId:"installer_ps_fault_status"`
		IsBankPs                  valueTypes.Bool      `json:"is_bank_ps" PointId:"is_bank_ps"`
		IsTuv                     valueTypes.Bool      `json:"is_tuv" PointId:"is_tuv"`
		JoinYearInitElec          valueTypes.Float     `json:"join_year_init_elec" PointId:"join_year_init_elec"`
		Latitude                  valueTypes.Float     `json:"latitude" PointId:"latitude"`
		Location                  valueTypes.String    `json:"location" PointId:"location"`
		Longitude                 valueTypes.Float     `json:"longitude" PointId:"longitude"`
		MapLatitude               valueTypes.Float     `json:"map_latitude" PointId:"map_latitude"`
		MapLongitude              valueTypes.Float     `json:"map_longitude" PointId:"map_longitude"`
		MlpeFlag                  valueTypes.Integer   `json:"mlpe_flag" PointId:"mlpe_flag"`
		Nmi                       valueTypes.String    `json:"nmi" PointId:"nmi"`
		OfflineDevCount           valueTypes.Count     `json:"offline_dev_count" PointId:"offline_dev_count"`
		OperateYear               interface{}          `json:"operate_year" PointId:"operate_year"`
		OperationBusName          valueTypes.String    `json:"operation_bus_name" PointId:"operation_bus_name"`
		OwnerAlarmCount           valueTypes.Count     `json:"owner_alarm_count" PointId:"owner_alarm_count"`
		OwnerFaultCount           valueTypes.Count     `json:"owner_fault_count" PointId:"owner_fault_count"`
		OwnerPsFaultStatus        valueTypes.Integer   `json:"owner_ps_fault_status" PointId:"owner_ps_fault_status"`

		P83022y                valueTypes.String    `json:"p83022y" PointId:"p83022" PointUpdateFreq:"UpdateFreq5Mins"`
		P83046                 valueTypes.Float     `json:"p83046" PointId:"p83046" PointUpdateFreq:"UpdateFreq5Mins"`
		P83048                 valueTypes.Float     `json:"p83048" PointId:"p83048" PointUpdateFreq:"UpdateFreq5Mins"`
		P83049                 valueTypes.Float     `json:"p83049" PointId:"p83049" PointUpdateFreq:"UpdateFreq5Mins"`
		P83050                 valueTypes.Float     `json:"p83050" PointId:"p83050" PointUpdateFreq:"UpdateFreq5Mins"`
		P83051                 valueTypes.Float     `json:"p83051" PointId:"p83051" PointUpdateFreq:"UpdateFreq5Mins"`
		P83054                 valueTypes.Float     `json:"p83054" PointId:"p83054" PointUpdateFreq:"UpdateFreq5Mins"`
		P83055                 valueTypes.Float     `json:"p83055" PointId:"p83055" PointUpdateFreq:"UpdateFreq5Mins"`
		P83067                 valueTypes.Float     `json:"p83067" PointId:"p83067" PointUpdateFreq:"UpdateFreq5Mins"`
		P83070                 valueTypes.Float     `json:"p83070" PointId:"p83070" PointUpdateFreq:"UpdateFreq5Mins"`
		P83076                 valueTypes.Float     `json:"p83076" PointId:"_p83076" PointName:"Pv Power" PointIgnore:"true"`                  // Dupe of PvPower
		P83077                 valueTypes.Float     `json:"p83077" PointId:"_p83077" PointName:"Pv Energy" PointIgnore:"true"`                 // Dupe of PvEnergy
		P83081                 valueTypes.Float     `json:"p83081" PointId:"_p83081" PointName:"Es Power" PointIgnore:"true"`                  // Dupe of EsPower
		P83089                 valueTypes.Float     `json:"p83089" PointId:"_p83089" PointName:"Es Discharge Energy" PointIgnore:"true"`       // Dupe of EsDischargeEnergy
		P83095                 valueTypes.Float     `json:"p83095" PointId:"_p83095" PointName:"Es Total Discharge Energy" PointIgnore:"true"` // Dupe of EsTotalDischargeEnergy
		P83118                 valueTypes.Float     `json:"p83118" PointId:"_p83118" PointName:"Use Energy" PointIgnore:"true"`                // Dupe of UseEnergy
		P83120                 valueTypes.Float     `json:"p83120" PointId:"_p83120" PointName:"Es Energy" PointIgnore:"true"`                 // Dupe of EsEnergy
		P83127                 valueTypes.Float     `json:"p83127" PointId:"_p83127" PointName:"Es Total Energy" PointIgnore:"true"`           // Dupe of EsTotalEnergy
		PvPower                valueTypes.UnitValue `json:"pv_power" PointId:"p83076" PointName:"Pv Power" PointUpdateFreq:"UpdateFreq5Mins"`
		PvEnergy               valueTypes.UnitValue `json:"pv_energy" PointId:"p83077" PointName:"Pv Energy" PointUpdateFreq:"UpdateFreq5Mins"`
		EsPower                valueTypes.UnitValue `json:"es_power" PointId:"p83081" PointName:"ES Power" PointUpdateFreq:"UpdateFreq5Mins"`
		EsDischargeEnergy      valueTypes.UnitValue `json:"es_disenergy" PointId:"p83089" PointName:"ES Discharge Energy" PointUpdateFreq:"UpdateFreq5Mins"`
		EsTotalDischargeEnergy valueTypes.UnitValue `json:"es_total_disenergy" PointId:"p83095" PointName:"ES Total Discharge Energy" PointUpdateFreq:"UpdateFreqTotal"`
		UseEnergy              valueTypes.UnitValue `json:"use_energy" PointId:"p83118" PointName:"Use Energy" PointUpdateFreq:"UpdateFreq5Mins"`
		EsEnergy               valueTypes.UnitValue `json:"es_energy" PointId:"p83120" PointName:"ES Energy" PointUpdateFreq:"UpdateFreq5Mins"`
		EsTotalEnergy          valueTypes.UnitValue `json:"es_total_energy" PointId:"p83127" PointName:"ES Total Energy" PointUpdateFreq:"UpdateFreqTotal"`

		ParamCo2               valueTypes.Float     `json:"param_co2" PointId:"param_co2"`
		ParamCoal              valueTypes.Float     `json:"param_coal" PointId:"param_coal"`
		ParamIncome            valueTypes.Float     `json:"param_income" PointId:"param_income"`
		ParamMeter             valueTypes.Float     `json:"param_meter" PointId:"param_meter"`
		ParamNox               valueTypes.Float     `json:"param_nox" PointId:"param_nox"`
		ParamPowder            valueTypes.Float     `json:"param_powder" PointId:"param_powder"`
		ParamSo2               valueTypes.Float     `json:"param_so2" PointId:"param_so2"`
		ParamTree              valueTypes.Float     `json:"param_tree" PointId:"param_tree"`
		ParamWater             valueTypes.Float     `json:"param_water" PointId:"param_water"`
		PrScale                string               `json:"pr_scale" PointId:"pr_scale"`
		Producer               interface{}          `json:"producer" PointId:"producer"`
		PsCountryId            valueTypes.Integer   `json:"ps_country_id" PointId:"ps_country_id"`
		PsFaultStatus          valueTypes.Integer   `json:"ps_fault_status" PointId:"ps_fault_status"`
		PsHealthStatus         valueTypes.Integer   `json:"ps_health_status" PointId:"ps_health_status"`
		PsHolder               valueTypes.String    `json:"ps_holder" PointId:"ps_holder"`
		PsId                   valueTypes.PsId      `json:"ps_id" PointId:"ps_id"`
		PsIsNotInit            valueTypes.Bool      `json:"ps_is_not_init" PointId:"ps_is_not_init"`
		PsName                 valueTypes.String    `json:"ps_name" PointId:"ps_name"`
		PsShortName            valueTypes.String    `json:"ps_short_name" PointId:"ps_short_name"`
		PsStatus               valueTypes.Integer   `json:"ps_status" PointId:"ps_status"`
		PsTimezone             valueTypes.String    `json:"ps_timezone" PointId:"ps_timezone"`
		PsType                 valueTypes.Integer   `json:"ps_type" PointId:"ps_type"`
		Radiation              valueTypes.UnitValue `json:"radiation" PointId:"radiation"`
		RadiationVirgin        valueTypes.Float     `json:"radiation_virgin" PointIgnore:"true"`
		RecordCreateTime       valueTypes.DateTime  `json:"recore_create_time" PointId:"record_create_time"`
		SafeStartDate          valueTypes.DateTime  `json:"safe_start_date" PointId:"safe_start_date"`
		ShareType              valueTypes.Integer   `json:"share_type" PointId:"share_type"`
		ShippingAddress        valueTypes.String    `json:"shipping_address" PointId:"shipping_address"`
		ShippingZipCode        valueTypes.String    `json:"shipping_zip_code" PointId:"shipping_zip_code"`
		TodayEnergy            valueTypes.UnitValue `json:"today_energy" PointId:"today_energy" PointUpdateFreq:"UpdateFreqDay"`
		TodayIncome            valueTypes.UnitValue `json:"today_income" PointId:"today_income" PointUpdateFreq:"UpdateFreqDay"`
		TotalCapacity          valueTypes.UnitValue `json:"total_capcity" PointId:"total_capacity" PointUpdateFreq:"UpdateFreqTotal"`
		TotalEnergy            valueTypes.UnitValue `json:"total_energy" PointId:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`
		TotalIncome            valueTypes.UnitValue `json:"total_income" PointId:"total_income" PointUpdateFreq:"UpdateFreqTotal"`
		TotalInitCo2Accelerate valueTypes.Float     `json:"total_init_co2_accelerate" PointId:"total_init_co2_accelerate" PointUpdateFreq:"UpdateFreqTotal"`
		TotalInitElec          valueTypes.Float     `json:"total_init_elec" PointId:"total_init_elec" PointUpdateFreq:"UpdateFreqTotal"`
		TotalInitProfit        valueTypes.Float     `json:"total_init_profit" PointId:"total_init_profit" PointUpdateFreq:"UpdateFreqTotal"`
		ValidFlag              valueTypes.Bool      `json:"valid_flag" PointId:"valid_flag"`
		WgsLatitude            valueTypes.Float     `json:"wgs_latitude" PointId:"wgs_latitude"`
		WgsLongitude           valueTypes.Float     `json:"wgs_longitude" PointId:"wgs_longitude"`
		ZipCode                valueTypes.String    `json:"zip_code" PointId:"zip_code"`
	} `json:"pageList" PointId:"page_list" PointNameFromChild:"PsId" PointNameAppend:"false" PointArrayFlatten:"false" DataTable:"true"`
	RowCount valueTypes.Count `json:"rowCount" PointId:"row_count" PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		entries.StructToDataMap(*e,  "system", GoStruct.EndPointPath{})
	}
	return entries
}


type Device struct {
	PsFaultStatus          valueTypes.Integer
	PsHealthStatus         valueTypes.Integer
	PsHolder               valueTypes.String
	PsId                   valueTypes.PsId
	PsName                 valueTypes.String
	PsShortName            valueTypes.String
	PsStatus               valueTypes.Integer
	PsType                 valueTypes.Integer
}
type Devices []Device

func (e *ResultData) GetPsDevices() Devices {
	var ret Devices
	for _, d := range e.PageList {
		ret = append(ret, Device{
			PsFaultStatus:  d.PsFaultStatus,
			PsHealthStatus: d.PsHealthStatus,
			PsHolder:       d.PsHolder,
			PsId:           d.PsId,
			PsName:         d.PsName,
			PsShortName:    d.PsShortName,
			PsStatus:       d.PsStatus,
			PsType:         d.PsType,
		})
	}
	return ret
}

func (e *ResultData) GetPsIds() []valueTypes.PsId {
	var ret []valueTypes.PsId
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsId)
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsName() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsName.Value())
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsSerial() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsShortName.Value())
			}
		}
	}
	return ret
}

func (e *EndPoint) GetPsIds() []valueTypes.PsId {
	return e.Response.ResultData.GetPsIds()
}
