// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebAppService

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/WebAppService/showPSView"
	"fmt"
)


var endpoints = [][]string {
	{"WebAppService","addMaterial","/v1/otherService/addMaterial"},
	{"WebAppService","addOptTicketInfo","/v1/faultService/addOptTicketInfo"},
	{"WebAppService","addSpareParts","/v1/otherService/addSpareParts"},
	{"WebAppService","associateQueryFaultNames","/v1/faultService/associateQueryFaultNames"},
	{"WebAppService","auditPsDeviceCheck","/v1/devService/auditPsDeviceCheck"},
	{"WebAppService","calcOutputRankByDay","/v1/powerStationService/calcOutputRankByDay"},
	{"WebAppService","changeReadStatus","/v1/reportService/changeReadStatus"},
	{"WebAppService","checkMaterialName","/v1/otherService/checkMaterialName"},
	{"WebAppService","confirmFault","/v1/faultService/confirmOrCloseFault"},
	{"WebAppService","copeOperateTicket","/v1/faultService/copeOperateTicket"},
	{"WebAppService","copySecondTypeTicket","/v1/faultService/copySecondTypeTicket"},
	{"WebAppService","copyWorkTicket","/v1/faultService/copyWorkTicket"},
	{"WebAppService","delOptTicketInfo","/v1/faultService/delOptTicketInfo"},
	{"WebAppService","deleteDuty","/v1/otherService/deleteDuty"},
	{"WebAppService","deleteDutyMid","/v1/otherService/deleteDutyMid"},
	{"WebAppService","deleteMaterial","/v1/otherService/deleteMaterial"},
	{"WebAppService","deleteOrSharedSelfReport","/v1/reportService/deleteOrSharedSelfReport"},
	{"WebAppService","deleteSecondTypeTicket","/v1/faultService/deleteSecondTypeTicket"},
	{"WebAppService","deleteSparePartsById","/v1/otherService/deleteSparePartsById"},
	{"WebAppService","deleteWorkTicket","/v1/faultService/deleteWorkTicket"},
	{"WebAppService","deviceFactoryList","/v1/devService/deviceFactoryList"},
	{"WebAppService","dispartDataPageList","/v1/commonService/dispartDataPageList"},
	{"WebAppService","executeTask","/v1/faultService/executeTask"},
	{"WebAppService","findCurrentTask","/v1/faultService/findCurrentTask"},
	{"WebAppService","findDeviceMessageByPskey","/v1/devService/findDeviceMessageByPskey"},
	{"WebAppService","findFactoryMessage","/v1/devService/findFactoryMessage"},
	{"WebAppService","findImgResources","/v1/faultService/findImgResources"},
	{"WebAppService","findMateiralSubType","/v1/otherService/findMateiralSubType"},
	{"WebAppService","findMaterialById","/v1/otherService/findMaterialById"},
	{"WebAppService","findMyDealedCurrentTask","/v1/faultService/findMyDealedCurrentTask"},
	{"WebAppService","findMyDealedImgResources","/v1/faultService/findMyDealedImgResources"},
	{"WebAppService","findSeriesInverterData","/v1/devService/findSeriesInverterData"},
	{"WebAppService","findWebRole","/v1/userService/findWebRole"},
	{"WebAppService","getAllPsFaultCount","/v1/faultService/getAllPsFaultCount"},
	{"WebAppService","getAllPsFaultCountByUserId","/v1/faultService/getAllPsFaultCountByUserId"},
	{"WebAppService","getAllPsList","/v1/powerStationService/getAllPsList"},
	{"WebAppService","getAllStore","/v1/otherService/getAllStore"},
	{"WebAppService","getBaseDeviceInfo","/v1/devService/getBaseDeviceInfo"},
	{"WebAppService","getBoxData","/v1/devService/getBoxData"},
	{"WebAppService","getCBoxTree","/v1/devService/getCBoxTree"},
	{"WebAppService","getCheckDevTypeList","/v1/devService/getCheckDevTypeList"},
	{"WebAppService","getCheckUserList","/v1/userService/getCheckUserList"},
	{"WebAppService","getChildOrg","/v1/otherService/getChildOrg"},
	{"WebAppService","getCo","/v1/powerStationService/getCo"},
	{"WebAppService","getCodeTreeMap","/v1/commonService/getCodeTreeMap"},
	{"WebAppService","getDST","/v1/userService/getDST"},
	{"WebAppService","getDataCounts","/v1/commonService/getDataCounts"},
	{"WebAppService","getDataInfo","/v1/powerStationService/getDataInfo"},
	{"WebAppService","getDevList","/v1/devService/getDevList"},
	{"WebAppService","getDevName","/v1/devService/getDevName"},
	{"WebAppService","getDevTypeList","/v1/devService/getDevTypeList"},
	{"WebAppService","getDeviceDataList","/v1/reportService/getDeviceDataList"},
	{"WebAppService","getDeviceFactory","/v1/devService/getDeviceFactory"},
	{"WebAppService","getDeviceInfoForCheck","/v1/devService/getDeviceInfoForCheck"},
	{"WebAppService","getDevicePointAttrs","/v1/devService/getDevicePointAttrs"},
	{"WebAppService","getDeviceTreeChild","/v1/devService/getDeviceTreeChild"},
	{"WebAppService","getDeviceUuid","/v1/devService/getDeviceUuid"},
	{"WebAppService","getDutyInfoById","/v1/otherService/getDutyInfoById"},
	{"WebAppService","getDutyOrgZtree","/v1/otherService/getDutyOrgZtree"},
	{"WebAppService","getElecEffectList","/v1/faultService/getElecEffectList"},
	{"WebAppService","getEnvironmentInfo","/v1/devService/getEnvironmentInfo"},
	{"WebAppService","getFaultList","/v1/faultService/getFaultList"},
	{"WebAppService","getFaultName","/v1/faultService/getFaultName"},
	{"WebAppService","getFaultOrder","/v1/faultService/getFaultOrder"},
	{"WebAppService","getFaultOrderByOrderId","/v1/faultService/getFaultOrderByOrderId"},
	{"WebAppService","getFaultOrderList","/v1/faultService/getFaultOrderList"},
	{"WebAppService","getFaultOrderStepList","/v1/faultService/getFaultOrderStepList"},
	{"WebAppService","getHTRoleList","/v1/userService/getHTRoleList"},
	{"WebAppService","getHistoryComments","/v1/faultService/getHistoryComments"},
	{"WebAppService","getInfo","/v1/devService/getInfo"},
	{"WebAppService","getInverteTableListCount","/v1/devService/getInverteTableListCount"},
	{"WebAppService","getInverterDiscreteDistributioList","/v1/reportService/getInverterDiscreteDistributioList"},
	{"WebAppService","getInverterDiscreteList","/v1/reportService/getInverterDiscreteList"},
	{"WebAppService","getInverterFactoryList","/v1/reportService/getInverterFactoryList"},
	{"WebAppService","getInverterInfo","/v1/devService/getInverterInfo"},
	{"WebAppService","getLoadCurveList","/v1/reportService/getLoadCurveList"},
	{"WebAppService","getMultiPowers","/v1/devService/getMultiPowers"},
	{"WebAppService","getOndutyQuery","/v1/otherService/getOndutyQuery"},
	{"WebAppService","getOperateTicketUserList","/v1/faultService/getOperateTicketUserList"},
	{"WebAppService","getOptTicketsAttachments","/v1/faultService/getOptTicketsAttachments"},
	{"WebAppService","getOrderCount","/v1/faultService/getOrderCount"},
	{"WebAppService","getOrderData","/v1/faultService/getOrderData"},
	{"WebAppService","getOrderDataSize","/v1/faultService/getOrderDataSize"},
	{"WebAppService","getOrgPsEquipmentList","/v1/faultService/getOrgPsEquipmentList"},
	{"WebAppService","getOrgPsPowerGenerationSummaryReport","/v1/reportService/getOrgPsPowerGenerationSummaryReport"},
	{"WebAppService","getParentUidChain","/v1/devService/getParentUidChain"},
	{"WebAppService","getPowerKwhkwpList","/v1/powerStationService/getPowerKwhkwpList"},
	{"WebAppService","getPowerPrList","/v1/powerStationService/getPowerPrList"},
	{"WebAppService","getPowerPredictionInfo","/v1/reportService/getPowerPredictionInfo"},
	{"WebAppService","getPowerTrendDayData","/v1/powerStationService/getPowerTrendDayData"},
	{"WebAppService","getPowerTrendMonthData","/v1/powerStationService/getPowerTrendMonthData"},
	{"WebAppService","getPowerTrendYearData","/v1/powerStationService/getPowerTrendYearData"},
	{"WebAppService","getPowerValue","/v1/reportService/getPowerValue"},
	{"WebAppService","getPsBlock","/v1/devService/getPsBlock"},
	{"WebAppService","getPsBlockData","/v1/devService/getPsBlockData"},
	{"WebAppService","getPsBlockTree","/v1/devService/getPsBlockTree"},
	{"WebAppService","getPsBoxListCount","/v1/devService/getPsBoxListCount"},
	{"WebAppService","getPsCBoxDetail","/v1/devService/getPsCBoxDetail"},
	{"WebAppService","getPsContact","/v1/powerStationService/getPsContact"},
	{"WebAppService","getPsDataVal","/v1/devService/getPsDataVal"},
	{"WebAppService","getPsDeviceCheckList","/v1/devService/getPsDeviceCheckList"},
	{"WebAppService","getPsDeviceFaultList","/v1/faultService/getPsDeviceFaultList"},
	{"WebAppService","getPsFaultList","/v1/faultService/getPsFaultList"},
	{"WebAppService","getPsIdByUserId","/v1/powerStationService/getPsIdByUserId"},
	{"WebAppService","getPsIdState","/v1/powerStationService/getPsIdState"},
	{"WebAppService","getPsList","/v1/powerStationService/getPsListForWeb"},
	{"WebAppService","getPsListForWorkTicket","/v1/faultService/getPsListForWorkTicket"},
	{"WebAppService","getPsPictureMessage","/v1/powerStationService/getPsPictureMessage"},
	{"WebAppService","getPsTicketSizeAndClockNum","/v1/faultService/getPsTicketSizeAndClockNum"},
	{"WebAppService","getPsTree","/v1/devService/getPsTree"},
	{"WebAppService","getPsTreeChild","/v1/devService/getPsTreeChild"},
	{"WebAppService","getPsUserList","/v1/userService/getPsUserList"},
	{"WebAppService","getPsValue","/v1/powerStationService/getPsValue"},
	{"WebAppService","getPscSeriseData","/v1/devService/getPscSeriseData"},
	{"WebAppService","getReportInfoByReportId","/v1/reportService/getReportInfoByReportId"},
	{"WebAppService","getReportListByType","/v1/reportService/getReportListByType"},
	{"WebAppService","getReportPsTree","/v1/reportService/getReportPsTree"},
	{"WebAppService","getRoleList","/v1/userService/getRoleList"},
	{"WebAppService","getSafeEffectList","/v1/faultService/getSafeEffectList"},
	{"WebAppService","getSecondTypeTicketList","/v1/faultService/getSecondTypeTicketList"},
	{"WebAppService","getSecondTypeTicketListForTicketDetail","/v1/faultService/getSecondTypeTicketListForTicketDetail"},
	{"WebAppService","getSelfReportPoint","/v1/reportService/getSelfReportPoint"},
	{"WebAppService","getSparePartsDetail","/v1/otherService/getSparePartsDetail"},
	{"WebAppService","getStatementList","/v1/reportService/getStatementList"},
	{"WebAppService","getStoreByStationId","/v1/otherService/getStoreByStationId"},
	{"WebAppService","getSysUserList","/v1/userService/getSysUserList"},
	{"WebAppService","getTableList","/v1/devService/getTableList"},
	{"WebAppService","getUserList","/v1/userService/getUserList"},
	{"WebAppService","getWeather","/v1/powerStationService/getWeather"},
	{"WebAppService","getWorkTicketList","/v1/faultService/getWorkTicketList"},
	{"WebAppService","getWorkTicketListForTicketDetail","/v1/faultService/getWorkTicketListForTicketDetail"},
	{"WebAppService","getWorkTicketRunningCount","/v1/faultService/getWorkTicketRunningCount"},
	{"WebAppService","getWorkTicketUserList","/v1/faultService/getWorkTicketUserList"},
	{"WebAppService","getinverterType","/v1/devService/getinverterType"},
	{"WebAppService","getreportPermissionByUser","/v1/reportService/getPsIdByUserId"},
	{"WebAppService","handleValue","/v1/reportService/handleValue"},
	{"WebAppService","modifyDeviceInfo","/v1/devService/modifyDeviceInfo"},
	{"WebAppService","operaStoreSpareParts","/v1/otherService/operaStoreSpareParts"},
	{"WebAppService","operateBillTransferToUser","/v1/faultService/operateBillTransferToUser"},
	{"WebAppService","queryAllStockInventory","/v1/otherService/queryAllStockInventory"},
	{"WebAppService","queryBatteryBoardsList","/v1/devService/queryBatteryBoardsList"},
	{"WebAppService","queryBatteryBoardsPointsData","/v1/devService/queryBatteryBoardsPointsData"},
	{"WebAppService","queryCodeByType","/v1/commonService/queryCodeByType"},
	{"WebAppService","queryDeviceInfoList","/v1/devService/queryDeviceInfoList"},
	{"WebAppService","queryElectricityCalendarData","/v1/reportService/queryElectricityCalendarData"},
	{"WebAppService","queryFaultCodes","/v1/faultService/queryFaultCodes"},
	{"WebAppService","queryFaultLevelAndType","/v1/faultService/queryFaultLevelAndType"},
	{"WebAppService","queryFaultNames","/v1/faultService/queryFaultNames"},
	{"WebAppService","queryMaterialType","/v1/otherService/queryMaterialType"},
	{"WebAppService","queryNounAndKlgList","/v1/faultService/queryNounAndKlgList"},
	{"WebAppService","queryNounList","/v1/faultService/queryNounList"},
	{"WebAppService","queryOptTickctInfo","/v1/faultService/queryOptTickctInfo"},
	{"WebAppService","queryOrgIdByUser","/v1/userService/queryOrgIdByUser"},
	{"WebAppService","queryPsCountryList","/v1/commonService/queryPsCountryList"},
	{"WebAppService","queryPsProvcnList","/v1/commonService/queryPsProvcnList"},
	{"WebAppService","queryPsTypeByPsId","/v1/powerStationService/queryPsTypeByPsId"},
	{"WebAppService","querySparePartsList","/v1/otherService/querySparePartsList"},
	{"WebAppService","queryStoreList","/v1/otherService/queryStoreList"},
	{"WebAppService","querySysTimezone","/v1/commonService/querySysTimezone"},
	{"WebAppService","queryUnInventorySpareList","/v1/otherService/queryUnInventorySpareList"},
	{"WebAppService","queryUserCurveTemplateData","/v1/devService/queryUserCurveTemplateData"},
	{"WebAppService","renewOperation","/v1/devService/renewOperation"},
	{"WebAppService","saveCustomReport","/v1/reportService/saveCustomReport"},
	{"WebAppService","saveDutyInfo","/v1/otherService/saveDutyInfo"},
	{"WebAppService","saveInventory","/v1/otherService/saveInventory"},
	{"WebAppService","saveMaterial","/v1/devService/saveMaterial"},
	{"WebAppService","saveSecondTypeTicket","/v1/faultService/saveSecondTypeTicket"},
	{"WebAppService","saveSelfReportPoint","/v1/reportService/saveSelfReportPoint"},
	{"WebAppService","saveWorkTicket","/v1/faultService/saveWorkTicket"},
	{"WebAppService","secondTypeTicketFlowImplementStep","/v1/faultService/secondTypeTicketFlowImplementStep"},
	{"WebAppService","secondTypeTicketFlowTransferStep","/v1/faultService/secondTypeTicketFlowTransferStep"},
	{"WebAppService","secondTypeUpdateSign","/v1/faultService/secondTypeUpdateSign"},
	{"WebAppService","selectPowerPageList","/v1/powerStationService/selectPowerPageList"},
	{"WebAppService","showAnalyzefxDetail","/v1/reportService/showAnalyzefxDetail"},
	{"WebAppService","showFxReport","/v1/reportService/showFxReport"},
	{"WebAppService","showMaterNameList","/v1/otherService/showMaterNameList"},
	{"WebAppService","showMaterSubTypeList","/v1/otherService/showMaterSubTypeList"},
	{"WebAppService","showPSView","/v1/powerStationService/showPSView"},
	{"WebAppService","showTjReport","/v1/reportService/showTjReport"},
	{"WebAppService","templateLikesInfo","/v1/faultService/templateLikesInfo"},
	{"WebAppService","updOptTicketInfo","/v1/faultService/updOptTicketInfo"},
	{"WebAppService","updataWorkTicketAfterStartProcess","/v1/faultService/updataWorkTicketAfterStartProcess"},
	{"WebAppService","updateBillTicketForTask","/v1/faultService/updateBillTicketForTask"},
	{"WebAppService","updateDutyInfo","/v1/otherService/updateDutyInfo"},
	{"WebAppService","updateKnowledgeBaseUseNumber","/v1/faultService/updateKnowledgeBaseUseNumber"},
	{"WebAppService","updateMaterial","/v1/otherService/updateMaterial"},
	{"WebAppService","updateSpareParts","/v1/otherService/updateSpareParts"},
	{"WebAppService","updateStopReason","/v1/devService/updateStopReason"},
	{"WebAppService","updateTemplate","/v1/faultService/updateKnowledgeBaseCitationCount"},
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
		EndPoints: api.TypeEndPoints {
			api.GetName(showPSView.EndPoint{}): showPSView.Init(apiRoot),
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
