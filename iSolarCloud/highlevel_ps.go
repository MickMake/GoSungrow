package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/findPsType"
	"GoSungrow/iSolarCloud/AppService/getChnnlListByPsId"
	"GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"GoSungrow/iSolarCloud/AppService/getPListinfoFromMysql"
	"GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailForSinglePage"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/AppService/getPsInstallerByPsId"
	"GoSungrow/iSolarCloud/AppService/getPsInstallerOrgInfoByPsId"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"GoSungrow/iSolarCloud/AppService/getReportData"
	"GoSungrow/iSolarCloud/AppService/psForcastInfo"
	"GoSungrow/iSolarCloud/AppService/psHourPointsValue"
	"GoSungrow/iSolarCloud/AppService/queryAllPsIdAndName"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/AppService/queryPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/queryPsIdList"
	"GoSungrow/iSolarCloud/AppService/queryPsProfit"
	"GoSungrow/iSolarCloud/AppService/queryPsStructureList"
	"GoSungrow/iSolarCloud/AppService/querySysAdvancedParam"
	"GoSungrow/iSolarCloud/AppService/reportList"
	"GoSungrow/iSolarCloud/MttvScreenService/getPsDeviceListValue"
	"GoSungrow/iSolarCloud/MttvScreenService/getPsKpiForHoursByPsId"
	"GoSungrow/iSolarCloud/WebAppService/getPsIdState"
	"GoSungrow/iSolarCloud/WebAppService/getReportPsTree"
	"GoSungrow/iSolarCloud/WebAppService/showPSView"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMaxDeviceIdByPsId"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPsTreeMenu"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
	"time"
)


func (sg *SunGrow) SetPsIds(args ...string) valueTypes.PsIds {
	var pids valueTypes.PsIds
	for range Only.Once {
		pids = valueTypes.SetPsIdStrings(args)
		if len(pids) > 0 {
			break
		}

		pids, sg.Error = sg.PsIds()
		if sg.Error != nil {
			break
		}
	}
	return pids
}

// func (sg *SunGrow) GetPsId() (valueTypes.PsId, error) {
// 	var ret valueTypes.PsId
//
// 	for range Only.Once {
//
// 		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		_getPsList := getPsList.AssertResultData(ep)
// 		psIds := _getPsList.GetPsIds()
// 		if len(psIds) == 0 {
// 			break
// 		}
//
// 		ret = psIds[0]
// 	}
//
// 	return ret, sg.Error
// }

func (sg *SunGrow) PsIds() (valueTypes.PsIds, error) {
	var ret valueTypes.PsIds

	for range Only.Once {
		ep := sg.GetByStruct(getPsList.EndPointName, nil, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsIds()
	}

	return ret, sg.Error
}


// PsTreeMenu - WebIscmAppService.getPsTreeMenu
func (sg *SunGrow) PsTreeMenu(psIds ...string) (PsTree, error) {
	var ret PsTree

	for range Only.Once {
		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, psId := range pids {
			ep := sg.GetByStruct(getPsTreeMenu.EndPointName,
				getPsTreeMenu.RequestData{PsId: psId},
				time.Hour,
			)
			if sg.IsError() {
				break
			}

			data := getPsTreeMenu.Assert(ep)

			ret.Scan(data.Response.ResultData.List, false)
		}
	}

	return ret, sg.Error
}

const Root = "0"
type Ps struct {
	Children []getPsTreeMenu.Ps
	Depth    int
	Print    bool
}
type PsTree struct {
	Devices  []getPsTreeMenu.Ps
	Map      map[string]*Ps
	Print    bool
	Stringer string
}

func (p PsTree) String() string {
	// p.loop(Root, 0, 0)
	return p.Stringer
}

func (p *PsTree) Root() *getPsTreeMenu.Ps {
	var ret *getPsTreeMenu.Ps
	for range Only.Once {
		if _, ok := p.Map[Root]; !ok {
			break
		}

		if len(p.Map[Root].Children) == 0 {
			break
		}

		ret = &(p.Map[Root].Children[0])
	}
	return ret
}

func (p *PsTree) Scan(devices []getPsTreeMenu.Ps, print bool) {
	p.Devices = devices
	p.Map = make(map[string]*Ps)
	for _, ps := range p.Devices {
		name := ps.UpUUID.String()
		// fmt.Printf("[%s]\tParent:%s\tSelf: - %s\t%s\t%s\t%s\n", name, ps.UpUUID, ps.UUID, ps.PsId, ps.PsKey, ps.DeviceName)
		if _, ok := p.Map[name]; !ok {
			p.Map[name] = &Ps {
				Children: []getPsTreeMenu.Ps{ ps},
			}
		} else {
			p.Map[name].Children = append(p.Map[name].Children, ps)
		}

		p.Print = print
		p.loop(Root, 0, 0)
	}
}

func (p *PsTree) loop(current string, count int, depth int) {
	if _, ok := p.Map[current]; !ok {
		return
	}
	if count >= len(p.Map) {
		return
	}
	count++
	p.Map[current].Depth = depth

	for _, child := range p.Map[current].Children {
		p.Stringer += fmt.Sprintf("%s\tPsId:%s\tPsName:%s\tPsKey:%s\tDeviceName:%s\tUuid:%s\n",
			"+" + strings.Repeat("--", depth),
			child.PsId, child.PsName,
			child.PsKey, child.DeviceName, child.UUID)

		name := child.UUID.String()
		depth++
		p.loop(name, count, depth)
		depth--
	}
}


// QueryPowerStationInfo - AppService.queryPowerStationInfo
func (sg *SunGrow) QueryPowerStationInfo(psId valueTypes.PsId, sn valueTypes.String) (queryPowerStationInfo.ResultData, error) {
	var ret queryPowerStationInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPowerStationInfo.EndPointName,
			queryPowerStationInfo.RequestData {
				PsId: psId,
				Sn: sn,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryPowerStationInfo.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// QueryPsIdList - AppService.queryPsIdList
func (sg *SunGrow) QueryPsIdList() ([]valueTypes.String, error) {
	var ret []valueTypes.String
	for range Only.Once {
		ep := sg.GetByStruct(queryPsIdList.EndPointName,
			queryPsIdList.RequestData {},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsIdList.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPsInstallerOrgInfoByPsId - AppService.getPsInstallerOrgInfoByPsId
func (sg *SunGrow) GetPsInstallerOrgInfoByPsId(psId valueTypes.PsId) (getPsInstallerOrgInfoByPsId.ResultData, error) {
	var ret getPsInstallerOrgInfoByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsInstallerOrgInfoByPsId.EndPointName,
			getPsInstallerOrgInfoByPsId.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsInstallerOrgInfoByPsId.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPListinfoFromMysql - AppService.getPListinfoFromMysql
func (sg *SunGrow) GetPListinfoFromMysql(psIds valueTypes.PsIds) (getPListinfoFromMysql.ResultData, error) {
	var ret getPListinfoFromMysql.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPListinfoFromMysql.EndPointName,
			getPListinfoFromMysql.RequestData {
				PsIds: psIds,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPListinfoFromMysql.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// QueryAllPsIdAndName - AppService.queryAllPsIdAndName
func (sg *SunGrow) QueryAllPsIdAndName() (queryAllPsIdAndName.ResultData, error) {
	var ret queryAllPsIdAndName.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryAllPsIdAndName.EndPointName,
			queryAllPsIdAndName.RequestData {},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryAllPsIdAndName.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// // GetRemoteUpgradeTaskList - AppService.getRemoteUpgradeTaskList
// func (sg *SunGrow) getRemoteUpgradeTaskList(psId valueTypes.PsIds) (getRemoteUpgradeTaskList.ResultData, error) {
// 	var ret getRemoteUpgradeTaskList.ResultData
// 	for range Only.Once {
// 		ep := sg.GetByStruct(
// 			"WebAppService.getRemoteUpgradeTaskList",
// 			getRemoteUpgradeTaskList.RequestData {
// 				PsIdList: psId,
// 			},
// 			time.Hour * 24,
// 		)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		data := getRemoteUpgradeTaskList.Assert(ep)
// 		ret = data.Response.ResultData
// 	}
// 	return ret, sg.Error
// }

// QuerySysAdvancedParam - AppService.querySysAdvancedParam
func (sg *SunGrow) QuerySysAdvancedParam(psId valueTypes.PsId, curPage valueTypes.Integer, size valueTypes.Integer) (querySysAdvancedParam.ResultData, error) {
	var ret querySysAdvancedParam.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(querySysAdvancedParam.EndPointName,
			querySysAdvancedParam.RequestData {
				PsId2: psId,
				CurPage: curPage,
				Size: size,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := querySysAdvancedParam.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetPsInstallerByPsId - AppService.getPsInstallerByPsId
func (sg *SunGrow) GetPsInstallerByPsId(psId valueTypes.PsId) (getPsInstallerByPsId.ResultData, error) {
	var ret getPsInstallerByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsInstallerByPsId.EndPointName,
			getPsInstallerByPsId.RequestData{
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsInstallerByPsId.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// FindPsType - AppService.findPsType
func (sg *SunGrow) FindPsType(psId valueTypes.PsId) (findPsType.ResultData, error) {
	var ret findPsType.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(findPsType.EndPointName,
			findPsType.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := findPsType.Assert(ep)
		ret = data.Response.ResultData
	}
	return ret, sg.Error
}

// GetChannelListByPsId - AppService.getChnnlListByPsId
func (sg *SunGrow) GetChannelListByPsId(psId valueTypes.PsId) (getChnnlListByPsId.ResultData, error) {
	var ret getChnnlListByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getChnnlListByPsId.EndPointName,
			getChnnlListByPsId.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getChnnlListByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetHouseholdStoragePsReport - AppService.getHouseholdStoragePsReport DateId:2022
func (sg *SunGrow) GetHouseholdStoragePsReport(psId valueTypes.PsId, dateId valueTypes.DateTime) (getHouseholdStoragePsReport.ResultData, error) {
	var ret getHouseholdStoragePsReport.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getHouseholdStoragePsReport.EndPointName,
			getHouseholdStoragePsReport.RequestData {
				PsId: psId,
				DateId: dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getHouseholdStoragePsReport.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetIncomeSettingInfos - AppService.getIncomeSettingInfos
func (sg *SunGrow) GetIncomeSettingInfos(psId valueTypes.PsId) (getIncomeSettingInfos.ResultData, error) {
	var ret getIncomeSettingInfos.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getIncomeSettingInfos.EndPointName,
			getIncomeSettingInfos.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getIncomeSettingInfos.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerChargeSettingInfo - AppService.getPowerChargeSettingInfo
func (sg *SunGrow) GetPowerChargeSettingInfo(psId valueTypes.PsId) (getPowerChargeSettingInfo.ResultData, error) {
	var ret getPowerChargeSettingInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerChargeSettingInfo.EndPointName,
			getPowerChargeSettingInfo.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerChargeSettingInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationBasicInfo - AppService.getPowerStationBasicInfo
func (sg *SunGrow) GetPowerStationBasicInfo(psId valueTypes.PsId) (getPowerStationBasicInfo.ResultData, error) {
	var ret getPowerStationBasicInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationBasicInfo.EndPointName,
			getPowerStationBasicInfo.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationBasicInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationData - AppService.getPowerStationData DateId:2022
func (sg *SunGrow) GetPowerStationData(psId valueTypes.PsId, dateId valueTypes.DateTime) (getPowerStationData.ResultData, error) {
	var ret getPowerStationData.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationData.EndPointName,
			getPowerStationData.RequestData {
				PsId: psId,
				DateId: dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationData.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationForHousehold - AppService.getPowerStationForHousehold
func (sg *SunGrow) GetPowerStationForHousehold(psId valueTypes.PsId) (getPowerStationForHousehold.ResultData, error) {
	var ret getPowerStationForHousehold.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationForHousehold.EndPointName,
			getPowerStationForHousehold.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationForHousehold.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStationInfo - AppService.getPowerStationInfo
func (sg *SunGrow) GetPowerStationInfo(psId valueTypes.PsId) (getPowerStationInfo.ResultData, error) {
	var ret getPowerStationInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStationInfo.EndPointName,
			getPowerStationInfo.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStationInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPowerStatistics - AppService.getPowerStatistics
func (sg *SunGrow) GetPowerStatistics(psId valueTypes.PsId) (getPowerStatistics.ResultData, error) {
	var ret getPowerStatistics.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPowerStatistics.EndPointName,
			getPowerStatistics.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPowerStatistics.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetail - AppService.getPsDetail
func (sg *SunGrow) GetPsDetail(psId valueTypes.PsId) (getPsDetail.ResultData, error) {
	var ret getPsDetail.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetail.EndPointName,
			getPsDetail.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetail.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetailForSinglePage - AppService.getPsDetailForSinglePage
func (sg *SunGrow) GetPsDetailForSinglePage(psId valueTypes.PsId) (getPsDetailForSinglePage.ResultData, error) {
	var ret getPsDetailForSinglePage.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetailForSinglePage.EndPointName,
			getPsDetailForSinglePage.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetailForSinglePage.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDetailWithPsType - AppService.getPsDetailWithPsType
func (sg *SunGrow) GetPsDetailWithPsType(psId valueTypes.PsId) (getPsDetailWithPsType.ResultData, error) {
	var ret getPsDetailWithPsType.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDetailWithPsType.EndPointName,
			getPsDetailWithPsType.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDetailWithPsType.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsHealthState - AppService.getPsHealthState
func (sg *SunGrow) GetPsHealthState(psId valueTypes.PsId) (getPsHealthState.ResultData, error) {
	var ret getPsHealthState.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsHealthState.EndPointName,
			getPsHealthState.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsHealthState.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsWeatherList - AppService.getPsWeatherList
func (sg *SunGrow) GetPsWeatherList(psId valueTypes.PsId) (getPsWeatherList.ResultData, error) {
	var ret getPsWeatherList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsWeatherList.EndPointName,
			getPsWeatherList.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsWeatherList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetReportData - AppService.getReportData ReportType:1
func (sg *SunGrow) GetReportData(psId valueTypes.PsId, reportType int64) (getReportData.ResultData, error) {
	var ret getReportData.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getReportData.EndPointName,
			getReportData.RequestData {
				PsId: psId,
				ReportType: valueTypes.SetIntegerValue(reportType),
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getReportData.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsForecastInfo - AppService.psForcastInfo
func (sg *SunGrow) GetPsForecastInfo(psId valueTypes.PsId) (psForcastInfo.ResultData, error) {
	var ret psForcastInfo.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(psForcastInfo.EndPointName,
			psForcastInfo.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := psForcastInfo.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsHourPointsValue - AppService.psHourPointsValue
func (sg *SunGrow) GetPsHourPointsValue(psId valueTypes.PsId) (psHourPointsValue.ResultData, error) {
	var ret psHourPointsValue.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(psHourPointsValue.EndPointName,
			psHourPointsValue.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := psHourPointsValue.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryDeviceList - AppService.queryDeviceList
func (sg *SunGrow) QueryDeviceList(psId valueTypes.PsId) (queryDeviceList.ResultData, error) {
	var ret queryDeviceList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryDeviceList.EndPointName,
			queryDeviceList.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryDeviceList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryDeviceListForApp - AppService.queryDeviceListForApp
func (sg *SunGrow) QueryDeviceListForApp(psId valueTypes.PsId) (queryDeviceListForApp.ResultData, error) {
	var ret queryDeviceListForApp.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryDeviceListForApp.EndPointName,
			queryDeviceListForApp.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryDeviceListForApp.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryPsProfit - AppService.queryPsProfit DateId:2022
func (sg *SunGrow) QueryPsProfit(psId valueTypes.PsId, dateId valueTypes.DateTime) (queryPsProfit.ResultData, error) {
	var ret queryPsProfit.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPsProfit.EndPointName,
			queryPsProfit.RequestData {
				PsId: psId,
				DateId: dateId,
				DateType: valueTypes.SetStringValue(dateId.DateType),
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsProfit.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryPsStructureList - AppService.queryPsStructureList
func (sg *SunGrow) QueryPsStructureList(psId valueTypes.PsId) (queryPsStructureList.ResultData, error) {
	var ret queryPsStructureList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(queryPsStructureList.EndPointName,
			queryPsStructureList.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := queryPsStructureList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// ReportList - AppService.reportList ReportType:1
func (sg *SunGrow) ReportList(psId valueTypes.PsId, reportType int64) (reportList.ResultData, error) {
	var ret reportList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(reportList.EndPointName,
			reportList.RequestData {
				PsId: psId,
				ReportType: valueTypes.SetIntegerValue(reportType),
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := reportList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsDeviceListValue - MttvScreenService.getPsDeviceListValue
func (sg *SunGrow) GetPsDeviceListValue(psId valueTypes.PsId) (getPsDeviceListValue.ResultData, error) {
	var ret getPsDeviceListValue.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsDeviceListValue.EndPointName,
			getPsDeviceListValue.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsDeviceListValue.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsKpiForHoursByPsId - MttvScreenService.getPsKpiForHoursByPsId Day:20221002
func (sg *SunGrow) GetPsKpiForHoursByPsId(psId valueTypes.PsId) (getPsKpiForHoursByPsId.ResultData, error) {
	var ret getPsKpiForHoursByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsKpiForHoursByPsId.EndPointName,
			getPsKpiForHoursByPsId.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsKpiForHoursByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsIdState - WebAppService.getPsIdState
func (sg *SunGrow) GetPsIdState(psId valueTypes.PsId) (getPsIdState.ResultData, error) {
	var ret getPsIdState.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsIdState.EndPointName,
			getPsIdState.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsIdState.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetReportPsTree - WebAppService.getReportPsTree DeviceType:14
func (sg *SunGrow) GetReportPsTree(psId valueTypes.PsId) (getReportPsTree.ResultData, error) {
	var ret getReportPsTree.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getReportPsTree.EndPointName,
			getReportPsTree.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getReportPsTree.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// ShowPSView - WebAppService.showPSView
func (sg *SunGrow) ShowPSView(psId valueTypes.PsId) (showPSView.ResultData, error) {
	var ret showPSView.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(showPSView.EndPointName,
			showPSView.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := showPSView.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetMaxDeviceIdByPsId - WebIscmAppService.getMaxDeviceIdByPsId
func (sg *SunGrow) GetMaxDeviceIdByPsId(psId valueTypes.PsId) (getMaxDeviceIdByPsId.ResultData, error) {
	var ret getMaxDeviceIdByPsId.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getMaxDeviceIdByPsId.EndPointName,
			getMaxDeviceIdByPsId.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getMaxDeviceIdByPsId.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// GetPsTreeMenu - WebIscmAppService.getPsTreeMenu
func (sg *SunGrow) GetPsTreeMenu(psId valueTypes.PsId) (getPsTreeMenu.ResultData, error) {
	var ret getPsTreeMenu.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getPsTreeMenu.EndPointName,
			getPsTreeMenu.RequestData {
				PsId: psId,
			},
			time.Hour * 24,
		)
		if sg.IsError() {
			break
		}

		data := getPsTreeMenu.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}


// // getRobotDynamicCleaningView - AppService.getRobotDynamicCleaningView
// func (sg *SunGrow) getRobotDynamicCleaningView(psId valueTypes.PsId) (getRobotDynamicCleaningView.ResultData, error) {
// 	var ret getRobotDynamicCleaningView.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetPowerRobotSweepAttrByPsId - AppService.getPowerRobotSweepAttrByPsId
// func (sg *SunGrow) GetPowerRobotSweepAttrByPsId(psId valueTypes.PsId) (getPowerRobotSweepAttrByPsId.ResultData, error) {
// 	var ret getPowerRobotSweepAttrByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // FindEnvironmentInfo - AppService.findEnvironmentInfo
// func (sg *SunGrow) FindEnvironmentInfo(psId valueTypes.PsId) (findEnvironmentInfo.ResultData, error) {
// 	var ret findEnvironmentInfo.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAllDeviceByPsId - AppService.getAllDeviceByPsId
// func (sg *SunGrow) GetAllDeviceByPsId(psId valueTypes.PsId) (getAllDeviceByPsId.ResultData, error) {
// 	var ret getAllDeviceByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAllPowerRobotViewInfoByPsId - AppService.getAllPowerRobotViewInfoByPsId
// func (sg *SunGrow) GetAllPowerRobotViewInfoByPsId(psId valueTypes.PsId) (getAllPowerRobotViewInfoByPsId.ResultData, error) {
// 	var ret getAllPowerRobotViewInfoByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // QueryPsNameByPsId - AppService.queryPsNameByPsId
// func (sg *SunGrow) QueryPsNameByPsId(psId valueTypes.PsId) (queryPsNameByPsId.ResultData, error) {
// 	var ret queryPsNameByPsId.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetTheoryAndActualPower - MttvScreenService.getTheoryAndActualPower
// func (sg *SunGrow) GetTheoryAndActualPower(psId valueTypes.PsId) (getTheoryAndActualPower.ResultData, error) {
// 	var ret getTheoryAndActualPower.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetInverterInfo - WebAppService.getInverterInfo
// func (sg *SunGrow) GetInverterInfo(psId valueTypes.PsId) (getInverterInfo.ResultData, error) {
// 	var ret getInverterInfo.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetMultiPowers - WebAppService.getMultiPowers
// func (sg *SunGrow) GetMultiPowers(psId valueTypes.PsId) (getMultiPowers.ResultData, error) {
// 	var ret getMultiPowers.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetPsCBoxDetail - WebAppService.getPsCBoxDetail
// func (sg *SunGrow) GetPsCBoxDetail(psId valueTypes.PsId) (getPsCBoxDetail.ResultData, error) {
// 	var ret getPsCBoxDetail.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // QueryBatteryBoardsList - WebAppService.queryBatteryBoardsList
// func (sg *SunGrow) QueryBatteryBoardsList(psId valueTypes.PsId) (queryBatteryBoardsList.ResultData, error) {
// 	var ret queryBatteryBoardsList.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // ShowTjReport - WebAppService.showTjReport
// func (sg *SunGrow) ShowTjReport(psId valueTypes.PsId) (showTjReport.ResultData, error) {
// 	var ret showTjReport.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }

// // GetAuthKeyList - WebIscmAppService.getAuthKeyList
// func (sg *SunGrow) GetAuthKeyList(psId valueTypes.PsId) (getAuthKeyList.ResultData, error) {
// 	var ret getAuthKeyList.ResultData
// 	for range Only.Once {
// 	}
// 	return ret, sg.Error
// }
