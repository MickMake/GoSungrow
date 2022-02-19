// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package MttvScreenService

import (
	"GoSungro/iSolarCloud/api"
	"fmt"
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


var _ api.Area = (*Area)(nil)

type Area api.AreaStruct


func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot *api.Web) Area {
	area := Area {
		ApiRoot:   apiRoot,
		Name:      api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints {},
	}

	return area
}


// ****************************************
// Methods not scoped by api.EndPoint interface type

func GetAreaName() string {
	return string(api.GetArea(Area{}))
}

func (a Area) GetEndPoint(name api.EndPointName) api.EndPoint {
	var ret api.EndPoint
	for _, e := range a.EndPoints {
		// fmt.Printf("endpoint: %v\n", e)
		if e.GetName() == name {
			ret = e
			break
		}
	}
	return ret
}


// ****************************************
// Methods scoped by api.Area interface type

func (a Area) Init(apiRoot *api.Web) api.AreaStruct {
	panic("implement me")
}

func (a Area) GetAreaName() api.AreaName {
	return a.Name
}

func (a Area) GetEndPoints() api.TypeEndPoints {
	for _, endpoint := range a.EndPoints {
		fmt.Printf("endpoint: %v\n", endpoint)
	}
	return a.EndPoints
}

func (a Area) Call(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

func (a Area) GetRequest(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) GetResponse(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) GetData(name api.EndPointName) api.Json {
	panic("implement me")
}

func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
