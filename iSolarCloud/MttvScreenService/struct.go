// Package MttvScreenService - API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package MttvScreenService

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/accumEnergyPsKpiData"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/addBuildProgressing"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/deleteBuildProgressing"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/energyEquivalentHoursRanking"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/energyGetRankBySortName"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/energyPovertyAlleviation"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/energyPowerGenerationTrends"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/findSingleStationPR"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getCapabilityTrend"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getKpiByUserIdAndAreaCode"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getMapByUser"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getOrgProByUserId"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPlanAndActualPower"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsDeviceListValue"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsInfoWithJoinGridByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsKpiForHoursByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsListByMapId"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getPsListByUserIdAndAreaCode"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/getTheoryAndActualPower"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/nextLevelOrgList"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/nextLevelOrgStatisticalDataList"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/orgPowerReport"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/queryBuildProgressing"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/queryBuildProgressingNew"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/queryBuildProgressingOne"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/queryNearlyTwoYearsGenAndPrft"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/querySaveEnergyPsSOH"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/queryScreenUserMdIfo"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/saveOrUpdatePosition"
	"github.com/anicoll/gosungrow/iSolarCloud/MttvScreenService/updateBuildProgressing"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
)

var _ api.Area = (*Area)(nil)

type Area api.AreaStruct

func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot api.Web) Area {
	area := Area{
		ApiRoot: apiRoot,
		Name:    api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints{
			api.GetName(accumEnergyPsKpiData.EndPoint{}):            accumEnergyPsKpiData.Init(apiRoot),
			api.GetName(addBuildProgressing.EndPoint{}):             addBuildProgressing.Init(apiRoot),
			api.GetName(deleteBuildProgressing.EndPoint{}):          deleteBuildProgressing.Init(apiRoot),
			api.GetName(energyEquivalentHoursRanking.EndPoint{}):    energyEquivalentHoursRanking.Init(apiRoot),
			api.GetName(energyGetRankBySortName.EndPoint{}):         energyGetRankBySortName.Init(apiRoot),
			api.GetName(energyPovertyAlleviation.EndPoint{}):        energyPovertyAlleviation.Init(apiRoot),
			api.GetName(energyPowerGenerationTrends.EndPoint{}):     energyPowerGenerationTrends.Init(apiRoot),
			api.GetName(findSingleStationPR.EndPoint{}):             findSingleStationPR.Init(apiRoot),
			api.GetName(getCapabilityTrend.EndPoint{}):              getCapabilityTrend.Init(apiRoot),
			api.GetName(getKpiByUserIdAndAreaCode.EndPoint{}):       getKpiByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(getMapByUser.EndPoint{}):                    getMapByUser.Init(apiRoot),
			api.GetName(getOrgProByUserId.EndPoint{}):               getOrgProByUserId.Init(apiRoot),
			api.GetName(getPlanAndActualPower.EndPoint{}):           getPlanAndActualPower.Init(apiRoot),
			api.GetName(getPsDeviceListValue.EndPoint{}):            getPsDeviceListValue.Init(apiRoot),
			api.GetName(getPsInfoWithJoinGridByPsId.EndPoint{}):     getPsInfoWithJoinGridByPsId.Init(apiRoot),
			api.GetName(getPsKpiForHoursByPsId.EndPoint{}):          getPsKpiForHoursByPsId.Init(apiRoot),
			api.GetName(getPsListByMapId.EndPoint{}):                getPsListByMapId.Init(apiRoot),
			api.GetName(getPsListByUserIdAndAreaCode.EndPoint{}):    getPsListByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(getTheoryAndActualPower.EndPoint{}):         getTheoryAndActualPower.Init(apiRoot),
			api.GetName(nextLevelOrgList.EndPoint{}):                nextLevelOrgList.Init(apiRoot),
			api.GetName(nextLevelOrgStatisticalDataList.EndPoint{}): nextLevelOrgStatisticalDataList.Init(apiRoot),
			api.GetName(orgPowerReport.EndPoint{}):                  orgPowerReport.Init(apiRoot),
			api.GetName(queryBuildProgressing.EndPoint{}):           queryBuildProgressing.Init(apiRoot),
			api.GetName(queryBuildProgressingNew.EndPoint{}):        queryBuildProgressingNew.Init(apiRoot),
			api.GetName(queryBuildProgressingOne.EndPoint{}):        queryBuildProgressingOne.Init(apiRoot),
			api.GetName(queryNearlyTwoYearsGenAndPrft.EndPoint{}):   queryNearlyTwoYearsGenAndPrft.Init(apiRoot),
			api.GetName(querySaveEnergyPsSOH.EndPoint{}):            querySaveEnergyPsSOH.Init(apiRoot),
			api.GetName(queryScreenUserMdIfo.EndPoint{}):            queryScreenUserMdIfo.Init(apiRoot),
			api.GetName(saveOrUpdatePosition.EndPoint{}):            saveOrUpdatePosition.Init(apiRoot),
			api.GetName(updateBuildProgressing.EndPoint{}):          updateBuildProgressing.Init(apiRoot),
		},
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

func (a Area) Call(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

func (a Area) GetRequest(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetResponse(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetData(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
