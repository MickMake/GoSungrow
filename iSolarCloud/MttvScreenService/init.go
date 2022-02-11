// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package MttvScreenService

import (
	"GoSungro/iSolarCloud/api"
	"reflect"
	"strings"
)


var endpoints = [][]string {
	{"MttvScreenService","accumEnergyPsKpiData","/v1/powerStationService/accumEnergyPsKpiData"},
	{"MttvScreenService","addBuildProgressing","/v1/powerStationService/addBuildProgressing"},
	{"MttvScreenService","deleteBuildProgressing","/v1/powerStationService/deleteBuildProgressing"},
	{"MttvScreenService","energyEquivalentHoursRanking","/v1/powerStationService/energyEquivalentHoursRanking"},
	{"MttvScreenService","energyGetRankBySortName","/v1/powerStationService/energyGetRankBySortName"},
	{"MttvScreenService","energyPovertyAlleviation","/v1/orgService/energyPovertyAlleviation"},
	{"MttvScreenService","energyPowerGenerationTrends","/v1/orgService/energyPowerGenerationTrends"},
	{"MttvScreenService","findSingleStationPR","/v1/powerStationService/findSingleStationPR"},
	{"MttvScreenService","getCapabilityTrend","/v1/powerStationService/getCapabilityTrend"},
	{"MttvScreenService","getKpiByUserIdAndAreaCode","/v1/powerStationService/getKpiByUserIdAndAreaCode"},
	{"MttvScreenService","getMapByUser","/v1/powerStationService/getMapByUser"},
	{"MttvScreenService","getOrgProByUserId","/v1/orgService/getOrgProByUserId"},
	{"MttvScreenService","getPlanAndActualPower","/v1/powerStationService/getPlanAndActualPower"},
	{"MttvScreenService","getPsDeviceListValue","/v1/powerStationService/getPsDeviceListValue"},
	{"MttvScreenService","getPsInfoWithJoinGridByPsId","/v1/powerStationService/getPsInfoWithJoinGridByPsId"},
	{"MttvScreenService","getPsKpiForHoursByPsId","/v1/powerStationService/getPsKpiForHoursByPsId"},
	{"MttvScreenService","getPsListByMapId","/v1/powerStationService/getPsListByMapId"},
	{"MttvScreenService","getPsListByUserIdAndAreaCode","/v1/powerStationService/getPsListByUserIdAndAreaCode"},
	{"MttvScreenService","getTheoryAndActualPower","/v1/powerStationService/getTheoryAndActualPower"},
	{"MttvScreenService","nextLevelOrgList","/v1/orgService/nextLevelOrgList"},
	{"MttvScreenService","nextLevelOrgStatisticalDataList","/v1/orgService/nextLevelOrgStatisticalDataList"},
	{"MttvScreenService","orgPowerReport","/v1/orgService/orgPowerReport"},
	{"MttvScreenService","queryBuildProgressing","/v1/powerStationService/queryBuildProgressing"},
	{"MttvScreenService","queryBuildProgressingNew","/v1/powerStationService/queryBuildProgressingNew"},
	{"MttvScreenService","queryBuildProgressingOne","/v1/powerStationService/queryBuildProgressingOne"},
	{"MttvScreenService","queryNearlyTwoYearsGenAndPrft","/v1/orgService/queryNearlyTwoYearsGenAndPrft"},
	{"MttvScreenService","querySaveEnergyPsSOH","/v1/powerStationService/querySaveEnergyPsSOH"},
	{"MttvScreenService","queryScreenUserMdIfo","/v1/userService/queryScreenUserMdIfo"},
	{"MttvScreenService","saveOrUpdatePosition","/v1/devService/saveOrUpdatePosition"},
	{"MttvScreenService","updateBuildProgressing","/v1/powerStationService/updateBuildProgressing"},
}


type AreaStruct struct {
	Name api.AreaName
	EndPoints api.TypeEndPoints
}
var Area AreaStruct

func init() {
	ref := reflect.TypeOf(AreaStruct{})
	p := ref.String()
	s := ref.Name()
	Area.Name = api.AreaName(strings.TrimSuffix(p, "." + s))
	Area.EndPoints = api.CreateEndPoints(endpoints)
}
