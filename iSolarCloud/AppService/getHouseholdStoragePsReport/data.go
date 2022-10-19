package getHouseholdStoragePsReport

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const Url = "/v1/powerStationService/getHouseholdStoragePsReport"
const Disabled = false

type RequestData struct {
	DateID   string `json:"date_id" required:"true"`
	DateType string `json:"date_type" required:"true"`
	PsId     valueTypes.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("ps_id: Can be fetched from getPsList.")
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	ConnectType       string          `json:"connect_type"`
	HasAmmeter        valueTypes.Bool `json:"has_ammeter"`
	IsHaveEsInverter  valueTypes.Bool `json:"is_have_es_inverter"`
	IsTransformSystem valueTypes.Bool `json:"is_transform_system"`

	DayData   *DayData   `json:"day_data,omitempty" PointName:"Day"`
	MonthData *MonthData `json:"month_data,omitempty" PointName:"Month"`
	YearData  *YearData  `json:"year_data,omitempty" PointName:"Year"`
	TotalData *TotalData `json:"total_data,omitempty" PointName:"Total"`
}

type DayData struct {
	JthdMap         valueTypes.UnitValue `json:"jthd_map"`
	JthdMapVirgin   valueTypes.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap         valueTypes.UnitValue `json:"jtyd_map"`
	JtydMapVirgin   valueTypes.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83072Map       valueTypes.UnitValue `json:"p83072_map"  PointId:"p83072" PointUpdateFreq:"UpdateFreqDay"`
	P83072MapVirgin valueTypes.UnitValue `json:"p83072_map_virgin"  PointIgnore:"true"`
	P83077Map       valueTypes.UnitValue `json:"p83077_map"  PointId:"p83077" PointUpdateFreq:"UpdateFreqDay"`
	P83077MapVirgin valueTypes.UnitValue `json:"p83077_map_virgin"  PointIgnore:"true"`
	P83088Map       valueTypes.UnitValue `json:"p83088_map"  PointId:"p83088" PointUpdateFreq:"UpdateFreqDay"`
	P83088MapVirgin valueTypes.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83089Map       valueTypes.UnitValue `json:"p83089_map"  PointId:"p83089" PointUpdateFreq:"UpdateFreqDay"`
	P83089MapVirgin valueTypes.UnitValue `json:"p83089_map_virgin"  PointIgnore:"true"`
	P83097Map       valueTypes.UnitValue `json:"p83097_map"  PointId:"p83097" PointUpdateFreq:"UpdateFreqDay"`
	P83097MapVirgin valueTypes.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83102Map       valueTypes.UnitValue `json:"p83102_map"  PointId:"p83102" PointUpdateFreq:"UpdateFreqDay"`
	P83102MapVirgin valueTypes.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83118Map       valueTypes.UnitValue `json:"p83118_map"  PointId:"p83118" PointUpdateFreq:"UpdateFreqDay"`
	P83118MapVirgin valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map       valueTypes.UnitValue `json:"p83119_map"  PointId:"p83119" PointUpdateFreq:"UpdateFreqDay"`
	P83119MapVirgin valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map       valueTypes.UnitValue `json:"p83120_map"  PointId:"p83120" PointUpdateFreq:"UpdateFreqDay"`
	P83120MapVirgin valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121          valueTypes.Float     `json:"p83121"  PointId:"p83121" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqDay"`
	P83122          valueTypes.Float     `json:"p83122"  PointId:"p83122" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqDay"`
	PointData15List []struct {
		P83076     valueTypes.Float    `json:"p83076"  PointId:"p83076" PointName:"PV Power" PointUnitFrom:"P83076Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83076Unit valueTypes.String   `json:"p83076_unit"  PointIgnore:"true"`
		P83080     valueTypes.Float    `json:"p83080"  PointId:"p83080" PointName:"PV Power To Grid" PointUnitFrom:"P83080Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83080Unit valueTypes.String   `json:"p83080_unit"  PointIgnore:"true"`
		P83086     valueTypes.Float    `json:"p83086"  PointId:"p83086" PointUnitFrom:"P83086Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83086Unit valueTypes.String   `json:"p83086_unit"  PointIgnore:"true"`
		P83087     valueTypes.Float    `json:"p83087"  PointId:"p83087" PointName:"PV Power To Battery" PointUnitFrom:"P83087Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83087Unit valueTypes.String   `json:"p83087_unit"  PointIgnore:"true"`
		P83096     valueTypes.Float    `json:"p83096"  PointId:"p83096" PointUnitFrom:"P83096Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83096Unit valueTypes.String   `json:"p83096_unit"  PointIgnore:"true"`
		P83101     valueTypes.Float    `json:"p83101"  PointId:"p83101" PointUnitFrom:"P83101Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83101Unit valueTypes.String   `json:"p83101_unit"  PointIgnore:"true"`
		P83106     valueTypes.Float    `json:"p83106"  PointId:"p83106" PointName:"PV Power To Load" PointUnitFrom:"P83106Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83106Unit valueTypes.String   `json:"p83106_unit"  PointIgnore:"true"`
		P83128     valueTypes.Float    `json:"p83128"  PointId:"p83128" PointName:"Grid Power" PointUnitFrom:"P83128Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqDay"`
		P83128Unit valueTypes.String   `json:"p83128_unit"  PointIgnore:"true"`
		TimeStamp  valueTypes.DateTime `json:"time_stamp"`
		Zfzy       valueTypes.Float    `json:"zfzy" PointName:"PV Power To Battery" PointUnitFrom:"ZfzyUnit" PointTimestampFrom:"TimeStamp"`
		ZfzyUnit   valueTypes.String   `json:"zfzy_unit"  PointIgnore:"true"`
	} `json:"point_data_15_list" PointNameFromChild:"TimeStamp"`
	ZfzyMap       valueTypes.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       valueTypes.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

type MonthData struct {
	JthdMap          valueTypes.UnitValue `json:"jthd_map"`
	JthdMapVirgin    valueTypes.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap          valueTypes.UnitValue `json:"jtyd_map"`
	JtydMapVirgin    valueTypes.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83073Map       valueTypes.UnitValue `json:"p83073_map"  PointId:"p83073" PointUpdateFreq:"UpdateFreqMonth"`
	P83073MapVirgin valueTypes.UnitValue `json:"p83073_map_virgin"  PointIgnore:"true"`
	P83078Map       valueTypes.UnitValue `json:"p83078_map"  PointId:"p83078" PointUpdateFreq:"UpdateFreqMonth"`
	P83078MapVirgin valueTypes.UnitValue `json:"p83078_map_virgin"  PointIgnore:"true"`
	P83088Map       valueTypes.UnitValue `json:"p83088_map"  PointId:"p83088" PointUpdateFreq:"UpdateFreqMonth"`
	P83088MapVirgin valueTypes.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83091Map       valueTypes.UnitValue `json:"p83091_map"  PointId:"p83091" PointUpdateFreq:"UpdateFreqMonth"`
	P83091MapVirgin valueTypes.UnitValue `json:"p83091_map_virgin"  PointIgnore:"true"`
	P83097Map       valueTypes.UnitValue `json:"p83097_map"  PointId:"p83097" PointUpdateFreq:"UpdateFreqMonth"`
	P83097MapVirgin valueTypes.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83103Map       valueTypes.UnitValue `json:"p83103_map"  PointId:"p83103" PointUpdateFreq:"UpdateFreqMonth"`
	P83103MapVirgin valueTypes.UnitValue `json:"p83103_map_virgin"  PointIgnore:"true"`
	P83118Map       valueTypes.UnitValue `json:"p83118_map"  PointId:"p83118" PointUpdateFreq:"UpdateFreqMonth"`
	P83118MapVirgin valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map       valueTypes.UnitValue `json:"p83119_map"  PointId:"p83119" PointUpdateFreq:"UpdateFreqMonth"`
	P83119MapVirgin valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map       valueTypes.UnitValue `json:"p83120_map"  PointId:"p83120" PointUpdateFreq:"UpdateFreqMonth"`
	P83120MapVirgin valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121          valueTypes.Float     `json:"p83121"  PointId:"p83121" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqMonth"`
	P83122          valueTypes.Float     `json:"p83122"  PointId:"p83122" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqMonth"`
	ZfzyMap         valueTypes.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin   valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap         valueTypes.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin   valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
	MonthDataDayList []struct {
		DateID                   valueTypes.Integer  `json:"date_id"`
		Jthd                     valueTypes.Float    `json:"jthd" PointUnitFrom:"JthdUnit" PointTimestampFrom:"TimeStamp"`
		JthdUnit                 valueTypes.String   `json:"jthd_unit"  PointIgnore:"true"`
		Jtyd                     valueTypes.Float    `json:"jtyd" PointUnitFrom:"JtydUnit" PointTimestampFrom:"TimeStamp"`
		JtydUnit                 valueTypes.String   `json:"jtyd_unit"  PointIgnore:"true"`
		P83022                   valueTypes.Float    `json:"p83022"  PointId:"p83022" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83072                   valueTypes.Float    `json:"p83072"  PointId:"p83072" PointUnitFrom:"p83072_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83072Unit               valueTypes.String   `json:"p83072_unit"  PointIgnore:"true"`
		P83077                   valueTypes.Float    `json:"p83077"  PointId:"p83077" PointUnitFrom:"p83077_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83077Unit               valueTypes.String   `json:"p83077_unit"  PointIgnore:"true"`
		P83088                   valueTypes.Float    `json:"p83088"  PointId:"p83088" PointUnitFrom:"p83088_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83088Unit               valueTypes.String   `json:"p83088_unit"  PointIgnore:"true"`
		P83089                   valueTypes.Float    `json:"p83089"  PointId:"p83089" PointUnitFrom:"p83089_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83089Unit               valueTypes.String   `json:"p83089_unit"  PointIgnore:"true"`
		P83097                   valueTypes.Float    `json:"p83097"  PointId:"p83097" PointUnitFrom:"p83097_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83097Unit               valueTypes.String   `json:"p83097_unit"  PointIgnore:"true"`
		P83102                   valueTypes.Float    `json:"p83102"  PointId:"p83102" PointUnitFrom:"p83102_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83102Unit               valueTypes.String   `json:"p83102_unit"  PointIgnore:"true"`
		P83118                   valueTypes.Float    `json:"p83118"  PointId:"p83118" PointUnitFrom:"p83118_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83118Unit               valueTypes.String   `json:"p83118_unit"  PointIgnore:"true"`
		P83119                   valueTypes.Float    `json:"p83119"  PointId:"p83119" PointUnitFrom:"p83119_unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83119Unit               valueTypes.String   `json:"p83119_unit"  PointIgnore:"true"`
		P83120                   valueTypes.Float    `json:"p83120"  PointId:"p83120" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83121                   valueTypes.Float    `json:"p83121"  PointId:"p83121" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		P83122                   valueTypes.Float    `json:"p83122"  PointId:"p83122" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqMonth"`
		PsID                     valueTypes.Integer  `json:"ps_id"`
		SelfConsumptionYield     valueTypes.Float    `json:"self_consumption_yield" PointUnitFrom:"SelfConsumptionYieldUnit" PointTimestampFrom:"TimeStamp"`
		SelfConsumptionYieldUnit valueTypes.String   `json:"self_consumption_yield_unit"  PointIgnore:"true"`
		TimeStamp                valueTypes.DateTime `json:"time_stamp"`
	} `json:"month_data_day_list" PointNameFromChild:"TimeStamp" PointNameDateFormat:"20060102"`
}

type YearData struct {
	JthdMap           valueTypes.UnitValue `json:"jthd_map"`
	JthdMapVirgin     valueTypes.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap           valueTypes.UnitValue `json:"jtyd_map"`
	JtydMapVirgin     valueTypes.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83074            valueTypes.UnitValue `json:"p83074_map"  PointId:"p83074" PointUpdateFreq:"UpdateFreqYear"`
	P83074MapVirgin   valueTypes.UnitValue `json:"p83074_map_virgin"  PointIgnore:"true"`
	P83079            valueTypes.UnitValue `json:"p83079_map"  PointId:"p83079" PointUpdateFreq:"UpdateFreqYear"`
	P83079MapVirgin   valueTypes.UnitValue `json:"p83079_map_virgin"  PointIgnore:"true"`
	P83088            valueTypes.UnitValue `json:"p83088_map"  PointId:"p83088" PointUpdateFreq:"UpdateFreqYear"`
	P83088MapVirgin   valueTypes.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83093            valueTypes.UnitValue `json:"p83093_map"  PointId:"p83093" PointUpdateFreq:"UpdateFreqYear"`
	P83093MapVirgin   valueTypes.UnitValue `json:"p83093_map_virgin"  PointIgnore:"true"`
	P83097            valueTypes.UnitValue `json:"p83097_map"  PointId:"p83097" PointUpdateFreq:"UpdateFreqYear"`
	P83097MapVirgin   valueTypes.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83104            valueTypes.UnitValue `json:"p83104_map"  PointId:"p83104" PointUpdateFreq:"UpdateFreqYear"`
	P83104MapVirgin   valueTypes.UnitValue `json:"p83104_map_virgin"  PointIgnore:"true"`
	P83118            valueTypes.UnitValue `json:"p83118_map"  PointId:"p83118" PointUpdateFreq:"UpdateFreqYear"`
	P83118MapVirgin   valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119            valueTypes.UnitValue `json:"p83119_map"  PointId:"p83119" PointUpdateFreq:"UpdateFreqYear"`
	P83119MapVirgin   valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120            valueTypes.UnitValue `json:"p83120_map"  PointId:"p83120" PointUpdateFreq:"UpdateFreqYear"`
	P83120MapVirgin   valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121            valueTypes.Float     `json:"p83121"  PointId:"p83121" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqYear"`
	P83122            valueTypes.Float     `json:"p83122"  PointId:"p83122" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqYear"`
	ZfzyMap       valueTypes.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       valueTypes.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
	YearDataMonthList []struct {
		DateID                   valueTypes.Integer  `json:"date_id"`
		Jthd                     valueTypes.Float    `json:"jthd" PointUnitFrom:"JthdUnit" PointTimestampFrom:"TimeStamp"`
		JthdUnit                 valueTypes.String   `json:"jthd_unit"  PointIgnore:"true"`
		Jtyd                     valueTypes.Float    `json:"jtyd" PointUnitFrom:"JtydUnit" PointTimestampFrom:"TimeStamp"`
		JtydUnit                 valueTypes.String   `json:"jtyd_unit"  PointIgnore:"true"`
		P83037                   valueTypes.Float    `json:"p83037"  PointId:"p83037" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83073                   valueTypes.Float    `json:"p83073"  PointId:"p83073" PointUnitFrom:"P83073Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83073Unit               valueTypes.String   `json:"p83073_unit"  PointIgnore:"true"`
		P83078                   valueTypes.Float    `json:"p83078"  PointId:"p83078" PointUnitFrom:"P83078Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83078Unit               valueTypes.String   `json:"p83078_unit"  PointIgnore:"true"`
		P83088                   valueTypes.Float    `json:"p83088"  PointId:"p83088" PointUnitFrom:"P83088Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83088Unit               valueTypes.String   `json:"p83088_unit"  PointIgnore:"true"`
		P83091                   valueTypes.Float    `json:"p83091"  PointId:"p83091" PointUnitFrom:"P83091Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83091Unit               valueTypes.String   `json:"p83091_unit"  PointIgnore:"true"`
		P83098                   valueTypes.Float    `json:"p83098"  PointId:"p83098" PointUnitFrom:"P83098Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83098Unit               valueTypes.String   `json:"p83098_unit"  PointIgnore:"true"`
		P83103                   valueTypes.Float    `json:"p83103"  PointId:"p83103" PointUnitFrom:"P83103Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83103Unit               valueTypes.String   `json:"p83103_unit"  PointIgnore:"true"`
		P83118                   valueTypes.Float    `json:"p83118"  PointId:"p83118" PointUnitFrom:"P83118Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83118Unit               valueTypes.String   `json:"p83118_unit"  PointIgnore:"true"`
		P83119                   valueTypes.Float    `json:"p83119"  PointId:"p83119" PointUnitFrom:"P83119Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83119Unit               valueTypes.String   `json:"p83119_unit"  PointIgnore:"true"`
		P83120                   valueTypes.Float    `json:"p83120"  PointId:"p83120" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83121                   valueTypes.Float    `json:"p83121"  PointId:"p83121" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		P83122                   valueTypes.Float    `json:"p83122"  PointId:"p83122" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqYear"`
		PsID                     valueTypes.Integer  `json:"ps_id"`
		SelfConsumptionYield     valueTypes.Float    `json:"self_consumption_yield" PointUnitFrom:"SelfConsumptionYieldUnit" PointTimestampFrom:"TimeStamp"`
		SelfConsumptionYieldUnit valueTypes.String   `json:"self_consumption_yield_unit"  PointIgnore:"true"`
		TimeStamp                valueTypes.DateTime `json:"time_stamp"`
	} `json:"year_data_month_list" PointNameFromChild:"TimeStamp" PointNameDateFormat:"200601"`
}

type TotalData struct {
	JthdMap           valueTypes.UnitValue `json:"jthd_map"`
	JthdMapVirgin     valueTypes.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap           valueTypes.UnitValue `json:"jtyd_map"`
	JtydMapVirgin     valueTypes.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83075            valueTypes.UnitValue `json:"p83075_map"  PointId:"p83075" PointUpdateFreq:"UpdateFreqTotal"`
	P83075MapVirgin   valueTypes.UnitValue `json:"p83075_map_virgin"  PointIgnore:"true"`
	P83094            valueTypes.UnitValue `json:"p83094_map"  PointId:"p83094" PointUpdateFreq:"UpdateFreqTotal"`
	P83094MapVirgin   valueTypes.UnitValue `json:"p83094_map_virgin"  PointIgnore:"true"`
	P83095            valueTypes.UnitValue `json:"p83095_map"  PointId:"p83095" PointUpdateFreq:"UpdateFreqTotal"`
	P83095MapVirgin   valueTypes.UnitValue `json:"p83095_map_virgin"  PointIgnore:"true"`
	P83105            valueTypes.UnitValue `json:"p83105_map"  PointId:"p83105" PointUpdateFreq:"UpdateFreqTotal"`
	P83105MapVirgin   valueTypes.UnitValue `json:"p83105_map_virgin"  PointIgnore:"true"`
	P83107            valueTypes.UnitValue `json:"p83107_map"  PointId:"p83107" PointUpdateFreq:"UpdateFreqTotal"`
	P83107MapVirgin   valueTypes.UnitValue `json:"p83107_map_virgin"  PointIgnore:"true"`
	P83123            valueTypes.UnitValue `json:"p83123_map"  PointId:"p83123" PointUpdateFreq:"UpdateFreqTotal"`
	P83123MapVirgin   valueTypes.UnitValue `json:"p83123_map_virgin"  PointIgnore:"true"`
	P83124            valueTypes.UnitValue `json:"p83124_map"  PointId:"p83124" PointUpdateFreq:"UpdateFreqTotal"`
	P83124MapVirgin   valueTypes.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83125            valueTypes.Float     `json:"p83125"  PointId:"p83125" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqTotal"`
	P83126            valueTypes.Float     `json:"p83126"  PointId:"p83126" PointUnitFrom:"FOO" PointUpdateFreq:"UpdateFreqTotal"`
	P83127            valueTypes.UnitValue `json:"p83127_map"  PointId:"p83127" PointUpdateFreq:"UpdateFreqTotal"`
	P83127MapVirgin   valueTypes.UnitValue `json:"p83127_map_virgin"  PointIgnore:"true"`
	ZfzyMap       valueTypes.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       valueTypes.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
	TotalDataYearList []struct {
		DateID                   valueTypes.Integer  `json:"date_id"`
		Jthd                     valueTypes.Float    `json:"jthd" PointUnitFrom:"JthdUnit" PointTimestampFrom:"TimeStamp"`
		JthdUnit                 valueTypes.String   `json:"jthd_unit"  PointIgnore:"true"`
		Jtyd                     valueTypes.Float    `json:"jtyd" PointUnitFrom:"JtydUnit" PointTimestampFrom:"TimeStamp"`
		JtydUnit                 valueTypes.String   `json:"jtyd_unit"  PointIgnore:"true"`
		P83038                   valueTypes.Float    `json:"p83038"  PointId:"p83038" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83074                   valueTypes.Float    `json:"p83074"  PointId:"p83074" PointUnitFrom:"P83074Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83074Unit               valueTypes.String   `json:"p83074_unit"  PointIgnore:"true"`
		P83079                   valueTypes.Float    `json:"p83079"  PointId:"p83079" PointUnitFrom:"P83079Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83079Unit               valueTypes.String   `json:"p83079_unit"  PointIgnore:"true"`
		P83088                   valueTypes.Float    `json:"p83088"  PointId:"p83088" PointUnitFrom:"P83088Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83088Unit               valueTypes.String   `json:"p83088_unit"  PointIgnore:"true"`
		P83093                   valueTypes.Float    `json:"p83093"  PointId:"p83093" PointUnitFrom:"P83093Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83093Unit               valueTypes.String   `json:"p83093_unit"  PointIgnore:"true"`
		P83099                   valueTypes.Float    `json:"p83099"  PointId:"p83099" PointUnitFrom:"P83099Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83099Unit               valueTypes.String   `json:"p83099_unit"  PointIgnore:"true"`
		P83104                   valueTypes.Float    `json:"p83104"  PointId:"p83104" PointUnitFrom:"P83104Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83104Unit               valueTypes.String   `json:"p83104_unit"  PointIgnore:"true"`
		P83118                   valueTypes.Float    `json:"p83118"  PointId:"p83118" PointUnitFrom:"P83118Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83118Unit               valueTypes.String   `json:"p83118_unit"  PointIgnore:"true"`
		P83119                   valueTypes.Float    `json:"p83119"  PointId:"p83119" PointUnitFrom:"P83119Unit" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83119Unit               valueTypes.String   `json:"p83119_unit"  PointIgnore:"true"`
		P83120                   valueTypes.Float    `json:"p83120"  PointId:"p83120" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83121                   valueTypes.Float    `json:"p83121"  PointId:"p83121" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		P83122                   valueTypes.Float    `json:"p83122"  PointId:"p83122" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqTotal"`
		PsID                     valueTypes.Integer  `json:"ps_id"`
		SelfConsumptionYield     valueTypes.Float    `json:"self_consumption_yield" PointUnitFrom:"SelfConsumptionYieldUnit" PointTimestampFrom:"TimeStamp"`
		SelfConsumptionYieldUnit valueTypes.String   `json:"self_consumption_yield_unit"  PointIgnore:"true"`
		TimeStamp                valueTypes.DateTime `json:"time_stamp"`
	} `json:"total_data_year_list" PointNameFromChild:"TimeStamp" PointNameDateFormat:"20060102"`
}

type DecodeResultData ResultData

func (p *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		var pd DecodeResultData

		// Store ResultData
		_ = json.Unmarshal(data, &pd)
		p.ConnectType = pd.ConnectType
		p.HasAmmeter = pd.HasAmmeter
		p.IsHaveEsInverter = pd.IsHaveEsInverter
		p.IsTransformSystem = pd.IsTransformSystem

		if strings.Contains(string(data), "day_data") {
			// Store ResultData.DayData
			err = json.Unmarshal(data, &pd.DayData)
			p.DayData = pd.DayData
			break
		}

		if strings.Contains(string(data), "month_data") {
			// Store ResultData.MonthData
			_ = json.Unmarshal(data, &pd.MonthData)
			p.MonthData = pd.MonthData
			break
		}

		if strings.Contains(string(data), "year_data") {
			// Store ResultData.YearData
			_ = json.Unmarshal(data, &pd.YearData)
			p.YearData = pd.YearData
			break
		}

		if strings.Contains(string(data), "total_data") {
			// Store ResultData.TotalData
			_ = json.Unmarshal(data, &pd.TotalData)
			p.TotalData = pd.TotalData
			break
		}
	}

	return err
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


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		name := api.JoinWithDots(0, "", pkg, e.Request.PsId)
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), valueTypes.NewDateTime(valueTypes.Now))
		break

		// if e.Response.ResultData.DayData != nil {
		// 	name := api.JoinWithDots(0, "", pkg, e.Request.PsId, "Day")
		// 	entries.StructToPoints(*e.Response.ResultData.DayData, name, e.Request.PsId.String(), valueTypes.NewDateTime(valueTypes.Now))
		//
		// 	// s := valueTypes.SizeOfArrayLength(e.Response.ResultData.DayData.PointData15List)
		// 	// for i, d := range e.Response.ResultData.DayData.PointData15List {
		// 	// 	name2 := api.JoinWithDots(s, "", name, i)
		// 	// 	entries.StructToPoints(*e.Response.ResultData.DayData, name2, e.Request.PsId.String(), d.TimeStamp)
		// 	// }
		// }
		//
		// if e.Response.ResultData.MonthData != nil {
		// 	name := api.JoinWithDots(0, "", pkg, e.Request.PsId, "Month")
		// 	entries.StructToPoints(*e.Response.ResultData.MonthData, name, e.Request.PsId.String(), valueTypes.NewDateTime(valueTypes.Now))
		//
		// 	// s := valueTypes.SizeOfArrayLength(e.Response.ResultData.MonthData.MonthDataDayList)
		// 	// for i, d := range e.Response.ResultData.MonthData.MonthDataDayList {
		// 	// 	name2 := api.JoinWithDots(s, "", name, i)
		// 	// 	entries.StructToPoints(*e.Response.ResultData.MonthData, name2, e.Request.PsId.String(), d.TimeStamp)
		// 	// }
		// }
		//
		// if e.Response.ResultData.YearData != nil {
		// 	name := api.JoinWithDots(0, "", pkg, e.Request.PsId, "Year")
		// 	entries.StructToPoints(*e.Response.ResultData.YearData, name, e.Request.PsId.String(), valueTypes.NewDateTime(valueTypes.Now))
		//
		// 	// s := valueTypes.SizeOfArrayLength(e.Response.ResultData.YearData.YearDataMonthList)
		// 	// for i, d := range e.Response.ResultData.YearData.YearDataMonthList {
		// 	// 	name2 := api.JoinWithDots(s, "", name, i)
		// 	// 	entries.StructToPoints(*e.Response.ResultData.YearData, name2, e.Request.PsId.String(), d.TimeStamp)
		// 	// }
		// }
		//
		// if e.Response.ResultData.TotalData != nil {
		// 	name := api.JoinWithDots(0, "", pkg, e.Request.PsId, "Total")
		// 	entries.StructToPoints(*e.Response.ResultData.TotalData, name, e.Request.PsId.String(), valueTypes.NewDateTime(valueTypes.Now))
		//
		// 	// s := valueTypes.SizeOfArrayLength(e.Response.ResultData.TotalData.TotalDataYearList)
		// 	// for i, d := range e.Response.ResultData.TotalData.TotalDataYearList {
		// 	// 	name2 := api.JoinWithDots(s, "", name, i)
		// 	// 	entries.StructToPoints(*e.Response.ResultData.TotalData, name2, e.Request.PsId.String(), d.TimeStamp)
		// 	// }
		// }
	}

	return entries
}
