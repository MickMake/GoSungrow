// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebAppService

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/sungro/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/sungro/WebAppService/showPSView"
	"fmt"
)

var _ api.Area = (*Area)(nil)

// var endpoints = [][]string{
// 	{"WebAppService", "addMaterial", "/v1/otherService/addMaterial"},
// 	{"WebAppService", "addOptTicketInfo", "/v1/faultService/addOptTicketInfo"},
// 	{"WebAppService", "addSpareParts", "/v1/otherService/addSpareParts"},
// 	{"WebAppService", "associateQueryFaultNames", "/v1/faultService/associateQueryFaultNames"},
// 	{"WebAppService", "auditPsDeviceCheck", "/v1/devService/auditPsDeviceCheck"},
// 	{"WebAppService", "calcOutputRankByDay", "/v1/powerStationService/calcOutputRankByDay"},
// 	{"WebAppService", "changeReadStatus", "/v1/reportService/changeReadStatus"},
// 	{"WebAppService", "checkMaterialName", "/v1/otherService/checkMaterialName"},
// 	{"WebAppService", "confirmFault", "/v1/faultService/confirmOrCloseFault"},
// 	{"WebAppService", "copeOperateTicket", "/v1/faultService/copeOperateTicket"},
// 	{"WebAppService", "copySecondTypeTicket", "/v1/faultService/copySecondTypeTicket"},
// 	{"WebAppService", "copyWorkTicket", "/v1/faultService/copyWorkTicket"},
// 	{"WebAppService", "delOptTicketInfo", "/v1/faultService/delOptTicketInfo"},
// 	{"WebAppService", "deleteDuty", "/v1/otherService/deleteDuty"},
// 	{"WebAppService", "deleteDutyMid", "/v1/otherService/deleteDutyMid"},
// 	{"WebAppService", "deleteMaterial", "/v1/otherService/deleteMaterial"},
// 	{"WebAppService", "deleteOrSharedSelfReport", "/v1/reportService/deleteOrSharedSelfReport"},
// 	{"WebAppService", "deleteSecondTypeTicket", "/v1/faultService/deleteSecondTypeTicket"},
// 	{"WebAppService", "deleteSparePartsById", "/v1/otherService/deleteSparePartsById"},
// 	{"WebAppService", "deleteWorkTicket", "/v1/faultService/deleteWorkTicket"},
// 	{"WebAppService", "deviceFactoryList", "/v1/devService/deviceFactoryList"},
// 	{"WebAppService", "dispartDataPageList", "/v1/commonService/dispartDataPageList"},
// 	{"WebAppService", "executeTask", "/v1/faultService/executeTask"},
// 	{"WebAppService", "findCurrentTask", "/v1/faultService/findCurrentTask"},
// 	{"WebAppService", "findDeviceMessageByPskey", "/v1/devService/findDeviceMessageByPskey"},
// 	{"WebAppService", "findFactoryMessage", "/v1/devService/findFactoryMessage"},
// 	{"WebAppService", "findImgResources", "/v1/faultService/findImgResources"},
// 	{"WebAppService", "findMateiralSubType", "/v1/otherService/findMateiralSubType"},
// 	{"WebAppService", "findMaterialById", "/v1/otherService/findMaterialById"},
// 	{"WebAppService", "findMyDealedCurrentTask", "/v1/faultService/findMyDealedCurrentTask"},
// 	{"WebAppService", "findMyDealedImgResources", "/v1/faultService/findMyDealedImgResources"},
// 	{"WebAppService", "findSeriesInverterData", "/v1/devService/findSeriesInverterData"},
// 	{"WebAppService", "findWebRole", "/v1/userService/findWebRole"},
// 	{"WebAppService", "getAllPsFaultCount", "/v1/faultService/getAllPsFaultCount"},
// 	{"WebAppService", "getAllPsFaultCountByUserId", "/v1/faultService/getAllPsFaultCountByUserId"},
// 	{"WebAppService", "getAllPsList", "/v1/powerStationService/getAllPsList"},
// 	{"WebAppService", "getAllStore", "/v1/otherService/getAllStore"},
// 	{"WebAppService", "getBaseDeviceInfo", "/v1/devService/getBaseDeviceInfo"},
// 	{"WebAppService", "getBoxData", "/v1/devService/getBoxData"},
// 	{"WebAppService", "getCBoxTree", "/v1/devService/getCBoxTree"},
// 	{"WebAppService", "getCheckDevTypeList", "/v1/devService/getCheckDevTypeList"},
// 	{"WebAppService", "getCheckUserList", "/v1/userService/getCheckUserList"},
// 	{"WebAppService", "getChildOrg", "/v1/otherService/getChildOrg"},
// 	{"WebAppService", "getCo", "/v1/powerStationService/getCo"},
// 	{"WebAppService", "getCodeTreeMap", "/v1/commonService/getCodeTreeMap"},
// 	{"WebAppService", "getDST", "/v1/userService/getDST"},
// 	{"WebAppService", "getDataCounts", "/v1/commonService/getDataCounts"},
// 	{"WebAppService", "getDataInfo", "/v1/powerStationService/getDataInfo"},
// 	{"WebAppService", "getDevList", "/v1/devService/getDevList"},
// 	{"WebAppService", "getDevName", "/v1/devService/getDevName"},
// 	{"WebAppService", "getDevTypeList", "/v1/devService/getDevTypeList"},
// 	{"WebAppService", "getDeviceDataList", "/v1/reportService/getDeviceDataList"},
// 	{"WebAppService", "getDeviceFactory", "/v1/devService/getDeviceFactory"},
// 	{"WebAppService", "getDeviceInfoForCheck", "/v1/devService/getDeviceInfoForCheck"},
// 	{"WebAppService", "getDevicePointAttrs", "/v1/devService/getDevicePointAttrs"},
// 	{"WebAppService", "getDeviceTreeChild", "/v1/devService/getDeviceTreeChild"},
// 	{"WebAppService", "getDeviceUuid", "/v1/devService/getDeviceUuid"},
// 	{"WebAppService", "getDutyInfoById", "/v1/otherService/getDutyInfoById"},
// 	{"WebAppService", "getDutyOrgZtree", "/v1/otherService/getDutyOrgZtree"},
// 	{"WebAppService", "getElecEffectList", "/v1/faultService/getElecEffectList"},
// 	{"WebAppService", "getEnvironmentInfo", "/v1/devService/getEnvironmentInfo"},
// 	{"WebAppService", "getFaultList", "/v1/faultService/getFaultList"},
// 	{"WebAppService", "getFaultName", "/v1/faultService/getFaultName"},
// 	{"WebAppService", "getFaultOrder", "/v1/faultService/getFaultOrder"},
// 	{"WebAppService", "getFaultOrderByOrderId", "/v1/faultService/getFaultOrderByOrderId"},
// 	{"WebAppService", "getFaultOrderList", "/v1/faultService/getFaultOrderList"},
// 	{"WebAppService", "getFaultOrderStepList", "/v1/faultService/getFaultOrderStepList"},
// 	{"WebAppService", "getHTRoleList", "/v1/userService/getHTRoleList"},
// 	{"WebAppService", "getHistoryComments", "/v1/faultService/getHistoryComments"},
// 	{"WebAppService", "getInfo", "/v1/devService/getInfo"},
// 	{"WebAppService", "getInverteTableListCount", "/v1/devService/getInverteTableListCount"},
// 	{"WebAppService", "getInverterDiscreteDistributioList", "/v1/reportService/getInverterDiscreteDistributioList"},
// 	{"WebAppService", "getInverterDiscreteList", "/v1/reportService/getInverterDiscreteList"},
// 	{"WebAppService", "getInverterFactoryList", "/v1/reportService/getInverterFactoryList"},
// 	{"WebAppService", "getInverterInfo", "/v1/devService/getInverterInfo"},
// 	{"WebAppService", "getLoadCurveList", "/v1/reportService/getLoadCurveList"},
// 	{"WebAppService", "getMultiPowers", "/v1/devService/getMultiPowers"},
// 	{"WebAppService", "getOndutyQuery", "/v1/otherService/getOndutyQuery"},
// 	{"WebAppService", "getOperateTicketUserList", "/v1/faultService/getOperateTicketUserList"},
// 	{"WebAppService", "getOptTicketsAttachments", "/v1/faultService/getOptTicketsAttachments"},
// 	{"WebAppService", "getOrderCount", "/v1/faultService/getOrderCount"},
// 	{"WebAppService", "getOrderData", "/v1/faultService/getOrderData"},
// 	{"WebAppService", "getOrderDataSize", "/v1/faultService/getOrderDataSize"},
// 	{"WebAppService", "getOrgPsEquipmentList", "/v1/faultService/getOrgPsEquipmentList"},
// 	{"WebAppService", "getOrgPsPowerGenerationSummaryReport", "/v1/reportService/getOrgPsPowerGenerationSummaryReport"},
// 	{"WebAppService", "getParentUidChain", "/v1/devService/getParentUidChain"},
// 	{"WebAppService", "getPowerKwhkwpList", "/v1/powerStationService/getPowerKwhkwpList"},
// 	{"WebAppService", "getPowerPrList", "/v1/powerStationService/getPowerPrList"},
// 	{"WebAppService", "getPowerPredictionInfo", "/v1/reportService/getPowerPredictionInfo"},
// 	{"WebAppService", "getPowerTrendDayData", "/v1/powerStationService/getPowerTrendDayData"},
// 	{"WebAppService", "getPowerTrendMonthData", "/v1/powerStationService/getPowerTrendMonthData"},
// 	{"WebAppService", "getPowerTrendYearData", "/v1/powerStationService/getPowerTrendYearData"},
// 	{"WebAppService", "getPowerValue", "/v1/reportService/getPowerValue"},
// 	{"WebAppService", "getPsBlock", "/v1/devService/getPsBlock"},
// 	{"WebAppService", "getPsBlockData", "/v1/devService/getPsBlockData"},
// 	{"WebAppService", "getPsBlockTree", "/v1/devService/getPsBlockTree"},
// 	{"WebAppService", "getPsBoxListCount", "/v1/devService/getPsBoxListCount"},
// 	{"WebAppService", "getPsCBoxDetail", "/v1/devService/getPsCBoxDetail"},
// 	{"WebAppService", "getPsContact", "/v1/powerStationService/getPsContact"},
// 	{"WebAppService", "getPsDataVal", "/v1/devService/getPsDataVal"},
// 	{"WebAppService", "getPsDeviceCheckList", "/v1/devService/getPsDeviceCheckList"},
// 	{"WebAppService", "getPsDeviceFaultList", "/v1/faultService/getPsDeviceFaultList"},
// 	{"WebAppService", "getPsFaultList", "/v1/faultService/getPsFaultList"},
// 	{"WebAppService", "getPsIdByUserId", "/v1/powerStationService/getPsIdByUserId"},
// 	{"WebAppService", "getPsIdState", "/v1/powerStationService/getPsIdState"},
// 	{"WebAppService", "getPsList", "/v1/powerStationService/getPsListForWeb"},
// 	{"WebAppService", "getPsListForWorkTicket", "/v1/faultService/getPsListForWorkTicket"},
// 	{"WebAppService", "getPsPictureMessage", "/v1/powerStationService/getPsPictureMessage"},
// 	{"WebAppService", "getPsTicketSizeAndClockNum", "/v1/faultService/getPsTicketSizeAndClockNum"},
// 	{"WebAppService", "getPsTree", "/v1/devService/getPsTree"},
// 	{"WebAppService", "getPsTreeChild", "/v1/devService/getPsTreeChild"},
// 	{"WebAppService", "getPsUserList", "/v1/userService/getPsUserList"},
// 	{"WebAppService", "getPsValue", "/v1/powerStationService/getPsValue"},
// 	{"WebAppService", "getPscSeriseData", "/v1/devService/getPscSeriseData"},
// 	{"WebAppService", "getReportInfoByReportId", "/v1/reportService/getReportInfoByReportId"},
// 	{"WebAppService", "getReportListByType", "/v1/reportService/getReportListByType"},
// 	{"WebAppService", "getReportPsTree", "/v1/reportService/getReportPsTree"},
// 	{"WebAppService", "getRoleList", "/v1/userService/getRoleList"},
// 	{"WebAppService", "getSafeEffectList", "/v1/faultService/getSafeEffectList"},
// 	{"WebAppService", "getSecondTypeTicketList", "/v1/faultService/getSecondTypeTicketList"},
// 	{"WebAppService", "getSecondTypeTicketListForTicketDetail", "/v1/faultService/getSecondTypeTicketListForTicketDetail"},
// 	{"WebAppService", "getSelfReportPoint", "/v1/reportService/getSelfReportPoint"},
// 	{"WebAppService", "getSparePartsDetail", "/v1/otherService/getSparePartsDetail"},
// 	{"WebAppService", "getStatementList", "/v1/reportService/getStatementList"},
// 	{"WebAppService", "getStoreByStationId", "/v1/otherService/getStoreByStationId"},
// 	{"WebAppService", "getSysUserList", "/v1/userService/getSysUserList"},
// 	{"WebAppService", "getTableList", "/v1/devService/getTableList"},
// 	{"WebAppService", "getUserList", "/v1/userService/getUserList"},
// 	{"WebAppService", "getWeather", "/v1/powerStationService/getWeather"},
// 	{"WebAppService", "getWorkTicketList", "/v1/faultService/getWorkTicketList"},
// 	{"WebAppService", "getWorkTicketListForTicketDetail", "/v1/faultService/getWorkTicketListForTicketDetail"},
// 	{"WebAppService", "getWorkTicketRunningCount", "/v1/faultService/getWorkTicketRunningCount"},
// 	{"WebAppService", "getWorkTicketUserList", "/v1/faultService/getWorkTicketUserList"},
// 	{"WebAppService", "getinverterType", "/v1/devService/getinverterType"},
// 	{"WebAppService", "getreportPermissionByUser", "/v1/reportService/getPsIdByUserId"},
// 	{"WebAppService", "handleValue", "/v1/reportService/handleValue"},
// 	{"WebAppService", "modifyDeviceInfo", "/v1/devService/modifyDeviceInfo"},
// 	{"WebAppService", "operaStoreSpareParts", "/v1/otherService/operaStoreSpareParts"},
// 	{"WebAppService", "operateBillTransferToUser", "/v1/faultService/operateBillTransferToUser"},
// 	{"WebAppService", "queryAllStockInventory", "/v1/otherService/queryAllStockInventory"},
// 	{"WebAppService", "queryBatteryBoardsList", "/v1/devService/queryBatteryBoardsList"},
// 	{"WebAppService", "queryBatteryBoardsPointsData", "/v1/devService/queryBatteryBoardsPointsData"},
// 	{"WebAppService", "queryCodeByType", "/v1/commonService/queryCodeByType"},
// 	{"WebAppService", "queryDeviceInfoList", "/v1/devService/queryDeviceInfoList"},
// 	{"WebAppService", "queryElectricityCalendarData", "/v1/reportService/queryElectricityCalendarData"},
// 	{"WebAppService", "queryFaultCodes", "/v1/faultService/queryFaultCodes"},
// 	{"WebAppService", "queryFaultLevelAndType", "/v1/faultService/queryFaultLevelAndType"},
// 	{"WebAppService", "queryFaultNames", "/v1/faultService/queryFaultNames"},
// 	{"WebAppService", "queryMaterialType", "/v1/otherService/queryMaterialType"},
// 	{"WebAppService", "queryNounAndKlgList", "/v1/faultService/queryNounAndKlgList"},
// 	{"WebAppService", "queryNounList", "/v1/faultService/queryNounList"},
// 	{"WebAppService", "queryOptTickctInfo", "/v1/faultService/queryOptTickctInfo"},
// 	{"WebAppService", "queryOrgIdByUser", "/v1/userService/queryOrgIdByUser"},
// 	{"WebAppService", "queryPsCountryList", "/v1/commonService/queryPsCountryList"},
// 	{"WebAppService", "queryPsProvcnList", "/v1/commonService/queryPsProvcnList"},
// 	{"WebAppService", "queryPsTypeByPsId", "/v1/powerStationService/queryPsTypeByPsId"},
// 	{"WebAppService", "querySparePartsList", "/v1/otherService/querySparePartsList"},
// 	{"WebAppService", "queryStoreList", "/v1/otherService/queryStoreList"},
// 	{"WebAppService", "querySysTimezone", "/v1/commonService/querySysTimezone"},
// 	{"WebAppService", "queryUnInventorySpareList", "/v1/otherService/queryUnInventorySpareList"},
// 	{"WebAppService", "queryUserCurveTemplateData", "/v1/devService/queryUserCurveTemplateData"},
// 	{"WebAppService", "renewOperation", "/v1/devService/renewOperation"},
// 	{"WebAppService", "saveCustomReport", "/v1/reportService/saveCustomReport"},
// 	{"WebAppService", "saveDutyInfo", "/v1/otherService/saveDutyInfo"},
// 	{"WebAppService", "saveInventory", "/v1/otherService/saveInventory"},
// 	{"WebAppService", "saveMaterial", "/v1/devService/saveMaterial"},
// 	{"WebAppService", "saveSecondTypeTicket", "/v1/faultService/saveSecondTypeTicket"},
// 	{"WebAppService", "saveSelfReportPoint", "/v1/reportService/saveSelfReportPoint"},
// 	{"WebAppService", "saveWorkTicket", "/v1/faultService/saveWorkTicket"},
// 	{"WebAppService", "secondTypeTicketFlowImplementStep", "/v1/faultService/secondTypeTicketFlowImplementStep"},
// 	{"WebAppService", "secondTypeTicketFlowTransferStep", "/v1/faultService/secondTypeTicketFlowTransferStep"},
// 	{"WebAppService", "secondTypeUpdateSign", "/v1/faultService/secondTypeUpdateSign"},
// 	{"WebAppService", "selectPowerPageList", "/v1/powerStationService/selectPowerPageList"},
// 	{"WebAppService", "showAnalyzefxDetail", "/v1/reportService/showAnalyzefxDetail"},
// 	{"WebAppService", "showFxReport", "/v1/reportService/showFxReport"},
// 	{"WebAppService", "showMaterNameList", "/v1/otherService/showMaterNameList"},
// 	{"WebAppService", "showMaterSubTypeList", "/v1/otherService/showMaterSubTypeList"},
// 	{"WebAppService", "showPSView", "/v1/powerStationService/showPSView"},
// 	{"WebAppService", "showTjReport", "/v1/reportService/showTjReport"},
// 	{"WebAppService", "templateLikesInfo", "/v1/faultService/templateLikesInfo"},
// 	{"WebAppService", "updOptTicketInfo", "/v1/faultService/updOptTicketInfo"},
// 	{"WebAppService", "updataWorkTicketAfterStartProcess", "/v1/faultService/updataWorkTicketAfterStartProcess"},
// 	{"WebAppService", "updateBillTicketForTask", "/v1/faultService/updateBillTicketForTask"},
// 	{"WebAppService", "updateDutyInfo", "/v1/otherService/updateDutyInfo"},
// 	{"WebAppService", "updateKnowledgeBaseUseNumber", "/v1/faultService/updateKnowledgeBaseUseNumber"},
// 	{"WebAppService", "updateMaterial", "/v1/otherService/updateMaterial"},
// 	{"WebAppService", "updateSpareParts", "/v1/otherService/updateSpareParts"},
// 	{"WebAppService", "updateStopReason", "/v1/devService/updateStopReason"},
// 	{"WebAppService", "updateTemplate", "/v1/faultService/updateKnowledgeBaseCitationCount"},
// }

type Area api.AreaStruct

func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot *api.Web) Area {
	area := Area{
		ApiRoot: apiRoot,
		Name:    api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints{
			// api.GetName(addMaterial.EndPoint{}): addMaterial.Init(apiRoot),	// /v1/otherService/addMaterial}
			// api.GetName(addOptTicketInfo.EndPoint{}): addOptTicketInfo.Init(apiRoot),	// /v1/faultService/addOptTicketInfo}
			// api.GetName(addSpareParts.EndPoint{}): addSpareParts.Init(apiRoot),	// /v1/otherService/addSpareParts}
			// api.GetName(associateQueryFaultNames.EndPoint{}): associateQueryFaultNames.Init(apiRoot),	// /v1/faultService/associateQueryFaultNames}
			// api.GetName(auditPsDeviceCheck.EndPoint{}): auditPsDeviceCheck.Init(apiRoot),	// /v1/devService/auditPsDeviceCheck}
			// api.GetName(calcOutputRankByDay.EndPoint{}): calcOutputRankByDay.Init(apiRoot),	// /v1/powerStationService/calcOutputRankByDay}
			// api.GetName(changeReadStatus.EndPoint{}): changeReadStatus.Init(apiRoot),	// /v1/reportService/changeReadStatus}
			// api.GetName(checkMaterialName.EndPoint{}): checkMaterialName.Init(apiRoot),	// /v1/otherService/checkMaterialName}
			// api.GetName(confirmFault.EndPoint{}): confirmFault.Init(apiRoot),	// /v1/faultService/confirmOrCloseFault}
			// api.GetName(copeOperateTicket.EndPoint{}): copeOperateTicket.Init(apiRoot),	// /v1/faultService/copeOperateTicket}
			// api.GetName(copySecondTypeTicket.EndPoint{}): copySecondTypeTicket.Init(apiRoot),	// /v1/faultService/copySecondTypeTicket}
			// api.GetName(copyWorkTicket.EndPoint{}): copyWorkTicket.Init(apiRoot),	// /v1/faultService/copyWorkTicket}
			// api.GetName(delOptTicketInfo.EndPoint{}): delOptTicketInfo.Init(apiRoot),	// /v1/faultService/delOptTicketInfo}
			// api.GetName(deleteDuty.EndPoint{}): deleteDuty.Init(apiRoot),	// /v1/otherService/deleteDuty}
			// api.GetName(deleteDutyMid.EndPoint{}): deleteDutyMid.Init(apiRoot),	// /v1/otherService/deleteDutyMid}
			// api.GetName(deleteMaterial.EndPoint{}): deleteMaterial.Init(apiRoot),	// /v1/otherService/deleteMaterial}
			// api.GetName(deleteOrSharedSelfReport.EndPoint{}): deleteOrSharedSelfReport.Init(apiRoot),	// /v1/reportService/deleteOrSharedSelfReport}
			// api.GetName(deleteSecondTypeTicket.EndPoint{}): deleteSecondTypeTicket.Init(apiRoot),	// /v1/faultService/deleteSecondTypeTicket}
			// api.GetName(deleteSparePartsById.EndPoint{}): deleteSparePartsById.Init(apiRoot),	// /v1/otherService/deleteSparePartsById}
			// api.GetName(deleteWorkTicket.EndPoint{}): deleteWorkTicket.Init(apiRoot),	// /v1/faultService/deleteWorkTicket}
			// api.GetName(deviceFactoryList.EndPoint{}): deviceFactoryList.Init(apiRoot),	// /v1/devService/deviceFactoryList}
			// api.GetName(dispartDataPageList.EndPoint{}): dispartDataPageList.Init(apiRoot),	// /v1/commonService/dispartDataPageList}
			// api.GetName(executeTask.EndPoint{}): executeTask.Init(apiRoot),	// /v1/faultService/executeTask}
			// api.GetName(findCurrentTask.EndPoint{}): findCurrentTask.Init(apiRoot),	// /v1/faultService/findCurrentTask}
			// api.GetName(findDeviceMessageByPskey.EndPoint{}): findDeviceMessageByPskey.Init(apiRoot),	// /v1/devService/findDeviceMessageByPskey}
			// api.GetName(findFactoryMessage.EndPoint{}): findFactoryMessage.Init(apiRoot),	// /v1/devService/findFactoryMessage}
			// api.GetName(findImgResources.EndPoint{}): findImgResources.Init(apiRoot),	// /v1/faultService/findImgResources}
			// api.GetName(findMateiralSubType.EndPoint{}): findMateiralSubType.Init(apiRoot),	// /v1/otherService/findMateiralSubType}
			// api.GetName(findMaterialById.EndPoint{}): findMaterialById.Init(apiRoot),	// /v1/otherService/findMaterialById}
			// api.GetName(findMyDealedCurrentTask.EndPoint{}): findMyDealedCurrentTask.Init(apiRoot),	// /v1/faultService/findMyDealedCurrentTask}
			// api.GetName(findMyDealedImgResources.EndPoint{}): findMyDealedImgResources.Init(apiRoot),	// /v1/faultService/findMyDealedImgResources}
			// api.GetName(findSeriesInverterData.EndPoint{}): findSeriesInverterData.Init(apiRoot),	// /v1/devService/findSeriesInverterData}
			// api.GetName(findWebRole.EndPoint{}): findWebRole.Init(apiRoot),	// /v1/userService/findWebRole}
			// api.GetName(getAllPsFaultCount.EndPoint{}): getAllPsFaultCount.Init(apiRoot),	// /v1/faultService/getAllPsFaultCount}
			// api.GetName(getAllPsFaultCountByUserId.EndPoint{}): getAllPsFaultCountByUserId.Init(apiRoot),	// /v1/faultService/getAllPsFaultCountByUserId}
			// api.GetName(getAllPsList.EndPoint{}): getAllPsList.Init(apiRoot),	// /v1/powerStationService/getAllPsList}
			// api.GetName(getAllStore.EndPoint{}): getAllStore.Init(apiRoot),	// /v1/otherService/getAllStore}
			// api.GetName(getBaseDeviceInfo.EndPoint{}): getBaseDeviceInfo.Init(apiRoot),	// /v1/devService/getBaseDeviceInfo}
			// api.GetName(getBoxData.EndPoint{}): getBoxData.Init(apiRoot),	// /v1/devService/getBoxData}
			// api.GetName(getCBoxTree.EndPoint{}): getCBoxTree.Init(apiRoot),	// /v1/devService/getCBoxTree}
			// api.GetName(getCheckDevTypeList.EndPoint{}): getCheckDevTypeList.Init(apiRoot),	// /v1/devService/getCheckDevTypeList}
			// api.GetName(getCheckUserList.EndPoint{}): getCheckUserList.Init(apiRoot),	// /v1/userService/getCheckUserList}
			// api.GetName(getChildOrg.EndPoint{}): getChildOrg.Init(apiRoot),	// /v1/otherService/getChildOrg}
			// api.GetName(getCo.EndPoint{}): getCo.Init(apiRoot),	// /v1/powerStationService/getCo}
			// api.GetName(getCodeTreeMap.EndPoint{}): getCodeTreeMap.Init(apiRoot),	// /v1/commonService/getCodeTreeMap}
			// api.GetName(getDST.EndPoint{}): getDST.Init(apiRoot),	// /v1/userService/getDST}
			// api.GetName(getDataCounts.EndPoint{}): getDataCounts.Init(apiRoot),	// /v1/commonService/getDataCounts}
			// api.GetName(getDataInfo.EndPoint{}): getDataInfo.Init(apiRoot),	// /v1/powerStationService/getDataInfo}
			// api.GetName(getDevList.EndPoint{}): getDevList.Init(apiRoot),	// /v1/devService/getDevList}
			// api.GetName(getDevName.EndPoint{}): getDevName.Init(apiRoot),	// /v1/devService/getDevName}
			// api.GetName(getDevTypeList.EndPoint{}): getDevTypeList.Init(apiRoot),	// /v1/devService/getDevTypeList}
			// api.GetName(getDeviceDataList.EndPoint{}): getDeviceDataList.Init(apiRoot),	// /v1/reportService/getDeviceDataList}
			// api.GetName(getDeviceFactory.EndPoint{}): getDeviceFactory.Init(apiRoot),	// /v1/devService/getDeviceFactory}
			// api.GetName(getDeviceInfoForCheck.EndPoint{}): getDeviceInfoForCheck.Init(apiRoot),	// /v1/devService/getDeviceInfoForCheck}
			// api.GetName(getDevicePointAttrs.EndPoint{}): getDevicePointAttrs.Init(apiRoot),	// /v1/devService/getDevicePointAttrs}
			// api.GetName(getDeviceTreeChild.EndPoint{}): getDeviceTreeChild.Init(apiRoot),	// /v1/devService/getDeviceTreeChild}
			// api.GetName(getDeviceUuid.EndPoint{}): getDeviceUuid.Init(apiRoot),	// /v1/devService/getDeviceUuid}
			// api.GetName(getDutyInfoById.EndPoint{}): getDutyInfoById.Init(apiRoot),	// /v1/otherService/getDutyInfoById}
			// api.GetName(getDutyOrgZtree.EndPoint{}): getDutyOrgZtree.Init(apiRoot),	// /v1/otherService/getDutyOrgZtree}
			// api.GetName(getElecEffectList.EndPoint{}): getElecEffectList.Init(apiRoot),	// /v1/faultService/getElecEffectList}
			// api.GetName(getEnvironmentInfo.EndPoint{}): getEnvironmentInfo.Init(apiRoot),	// /v1/devService/getEnvironmentInfo}
			// api.GetName(getFaultList.EndPoint{}): getFaultList.Init(apiRoot),	// /v1/faultService/getFaultList}
			// api.GetName(getFaultName.EndPoint{}): getFaultName.Init(apiRoot),	// /v1/faultService/getFaultName}
			// api.GetName(getFaultOrder.EndPoint{}): getFaultOrder.Init(apiRoot),	// /v1/faultService/getFaultOrder}
			// api.GetName(getFaultOrderByOrderId.EndPoint{}): getFaultOrderByOrderId.Init(apiRoot),	// /v1/faultService/getFaultOrderByOrderId}
			// api.GetName(getFaultOrderList.EndPoint{}): getFaultOrderList.Init(apiRoot),	// /v1/faultService/getFaultOrderList}
			// api.GetName(getFaultOrderStepList.EndPoint{}): getFaultOrderStepList.Init(apiRoot),	// /v1/faultService/getFaultOrderStepList}
			// api.GetName(getHTRoleList.EndPoint{}): getHTRoleList.Init(apiRoot),	// /v1/userService/getHTRoleList}
			// api.GetName(getHistoryComments.EndPoint{}): getHistoryComments.Init(apiRoot),	// /v1/faultService/getHistoryComments}
			// api.GetName(getInfo.EndPoint{}): getInfo.Init(apiRoot),	// /v1/devService/getInfo}
			// api.GetName(getInverteTableListCount.EndPoint{}): getInverteTableListCount.Init(apiRoot),	// /v1/devService/getInverteTableListCount}
			// api.GetName(getInverterDiscreteDistributioList.EndPoint{}): getInverterDiscreteDistributioList.Init(apiRoot),	// /v1/reportService/getInverterDiscreteDistributioList}
			// api.GetName(getInverterDiscreteList.EndPoint{}): getInverterDiscreteList.Init(apiRoot),	// /v1/reportService/getInverterDiscreteList}
			// api.GetName(getInverterFactoryList.EndPoint{}): getInverterFactoryList.Init(apiRoot),	// /v1/reportService/getInverterFactoryList}
			// api.GetName(getInverterInfo.EndPoint{}): getInverterInfo.Init(apiRoot),	// /v1/devService/getInverterInfo}
			// api.GetName(getLoadCurveList.EndPoint{}): getLoadCurveList.Init(apiRoot),	// /v1/reportService/getLoadCurveList}
			// api.GetName(getMultiPowers.EndPoint{}): getMultiPowers.Init(apiRoot),	// /v1/devService/getMultiPowers}
			// api.GetName(getOndutyQuery.EndPoint{}): getOndutyQuery.Init(apiRoot),	// /v1/otherService/getOndutyQuery}
			// api.GetName(getOperateTicketUserList.EndPoint{}): getOperateTicketUserList.Init(apiRoot),	// /v1/faultService/getOperateTicketUserList}
			// api.GetName(getOptTicketsAttachments.EndPoint{}): getOptTicketsAttachments.Init(apiRoot),	// /v1/faultService/getOptTicketsAttachments}
			// api.GetName(getOrderCount.EndPoint{}): getOrderCount.Init(apiRoot),	// /v1/faultService/getOrderCount}
			// api.GetName(getOrderData.EndPoint{}): getOrderData.Init(apiRoot),	// /v1/faultService/getOrderData}
			// api.GetName(getOrderDataSize.EndPoint{}): getOrderDataSize.Init(apiRoot),	// /v1/faultService/getOrderDataSize}
			// api.GetName(getOrgPsEquipmentList.EndPoint{}): getOrgPsEquipmentList.Init(apiRoot),	// /v1/faultService/getOrgPsEquipmentList}
			// api.GetName(getOrgPsPowerGenerationSummaryReport.EndPoint{}): getOrgPsPowerGenerationSummaryReport.Init(apiRoot),	// /v1/reportService/getOrgPsPowerGenerationSummaryReport}
			// api.GetName(getParentUidChain.EndPoint{}): getParentUidChain.Init(apiRoot),	// /v1/devService/getParentUidChain}
			// api.GetName(getPowerKwhkwpList.EndPoint{}): getPowerKwhkwpList.Init(apiRoot),	// /v1/powerStationService/getPowerKwhkwpList}
			// api.GetName(getPowerPrList.EndPoint{}): getPowerPrList.Init(apiRoot),	// /v1/powerStationService/getPowerPrList}
			// api.GetName(getPowerPredictionInfo.EndPoint{}): getPowerPredictionInfo.Init(apiRoot),	// /v1/reportService/getPowerPredictionInfo}
			// api.GetName(getPowerTrendDayData.EndPoint{}): getPowerTrendDayData.Init(apiRoot),	// /v1/powerStationService/getPowerTrendDayData}
			// api.GetName(getPowerTrendMonthData.EndPoint{}): getPowerTrendMonthData.Init(apiRoot),	// /v1/powerStationService/getPowerTrendMonthData}
			// api.GetName(getPowerTrendYearData.EndPoint{}): getPowerTrendYearData.Init(apiRoot),	// /v1/powerStationService/getPowerTrendYearData}
			// api.GetName(getPowerValue.EndPoint{}): getPowerValue.Init(apiRoot),	// /v1/reportService/getPowerValue}
			// api.GetName(getPsBlock.EndPoint{}): getPsBlock.Init(apiRoot),	// /v1/devService/getPsBlock}
			// api.GetName(getPsBlockData.EndPoint{}): getPsBlockData.Init(apiRoot),	// /v1/devService/getPsBlockData}
			// api.GetName(getPsBlockTree.EndPoint{}): getPsBlockTree.Init(apiRoot),	// /v1/devService/getPsBlockTree}
			// api.GetName(getPsBoxListCount.EndPoint{}): getPsBoxListCount.Init(apiRoot),	// /v1/devService/getPsBoxListCount}
			// api.GetName(getPsCBoxDetail.EndPoint{}): getPsCBoxDetail.Init(apiRoot),	// /v1/devService/getPsCBoxDetail}
			// api.GetName(getPsContact.EndPoint{}): getPsContact.Init(apiRoot),	// /v1/powerStationService/getPsContact}
			// api.GetName(getPsDataVal.EndPoint{}): getPsDataVal.Init(apiRoot),	// /v1/devService/getPsDataVal}
			// api.GetName(getPsDeviceCheckList.EndPoint{}): getPsDeviceCheckList.Init(apiRoot),	// /v1/devService/getPsDeviceCheckList}
			// api.GetName(getPsDeviceFaultList.EndPoint{}): getPsDeviceFaultList.Init(apiRoot),	// /v1/faultService/getPsDeviceFaultList}
			// api.GetName(getPsFaultList.EndPoint{}): getPsFaultList.Init(apiRoot),	// /v1/faultService/getPsFaultList}
			// api.GetName(getPsIdByUserId.EndPoint{}): getPsIdByUserId.Init(apiRoot),	// /v1/powerStationService/getPsIdByUserId}
			// api.GetName(getPsIdState.EndPoint{}): getPsIdState.Init(apiRoot),	// /v1/powerStationService/getPsIdState}
			// api.GetName(getPsList.EndPoint{}): getPsList.Init(apiRoot),	// /v1/powerStationService/getPsListForWeb}
			// api.GetName(getPsListForWorkTicket.EndPoint{}): getPsListForWorkTicket.Init(apiRoot),	// /v1/faultService/getPsListForWorkTicket}
			// api.GetName(getPsPictureMessage.EndPoint{}): getPsPictureMessage.Init(apiRoot),	// /v1/powerStationService/getPsPictureMessage}
			// api.GetName(getPsTicketSizeAndClockNum.EndPoint{}): getPsTicketSizeAndClockNum.Init(apiRoot),	// /v1/faultService/getPsTicketSizeAndClockNum}
			// api.GetName(getPsTree.EndPoint{}): getPsTree.Init(apiRoot),	// /v1/devService/getPsTree}
			// api.GetName(getPsTreeChild.EndPoint{}): getPsTreeChild.Init(apiRoot),	// /v1/devService/getPsTreeChild}
			// api.GetName(getPsUserList.EndPoint{}): getPsUserList.Init(apiRoot),	// /v1/userService/getPsUserList}
			// api.GetName(getPsValue.EndPoint{}): getPsValue.Init(apiRoot),	// /v1/powerStationService/getPsValue}
			// api.GetName(getPscSeriseData.EndPoint{}): getPscSeriseData.Init(apiRoot),	// /v1/devService/getPscSeriseData}
			// api.GetName(getReportInfoByReportId.EndPoint{}): getReportInfoByReportId.Init(apiRoot),	// /v1/reportService/getReportInfoByReportId}
			// api.GetName(getReportListByType.EndPoint{}): getReportListByType.Init(apiRoot),	// /v1/reportService/getReportListByType}
			// api.GetName(getReportPsTree.EndPoint{}): getReportPsTree.Init(apiRoot),	// /v1/reportService/getReportPsTree}
			// api.GetName(getRoleList.EndPoint{}): getRoleList.Init(apiRoot),	// /v1/userService/getRoleList}
			// api.GetName(getSafeEffectList.EndPoint{}): getSafeEffectList.Init(apiRoot),	// /v1/faultService/getSafeEffectList}
			// api.GetName(getSecondTypeTicketList.EndPoint{}): getSecondTypeTicketList.Init(apiRoot),	// /v1/faultService/getSecondTypeTicketList}
			// api.GetName(getSecondTypeTicketListForTicketDetail.EndPoint{}): getSecondTypeTicketListForTicketDetail.Init(apiRoot),	// /v1/faultService/getSecondTypeTicketListForTicketDetail}
			// api.GetName(getSelfReportPoint.EndPoint{}): getSelfReportPoint.Init(apiRoot),	// /v1/reportService/getSelfReportPoint}
			// api.GetName(getSparePartsDetail.EndPoint{}): getSparePartsDetail.Init(apiRoot),	// /v1/otherService/getSparePartsDetail}
			// api.GetName(getStatementList.EndPoint{}): getStatementList.Init(apiRoot),	// /v1/reportService/getStatementList}
			// api.GetName(getStoreByStationId.EndPoint{}): getStoreByStationId.Init(apiRoot),	// /v1/otherService/getStoreByStationId}
			// api.GetName(getSysUserList.EndPoint{}): getSysUserList.Init(apiRoot),	// /v1/userService/getSysUserList}
			// api.GetName(getTableList.EndPoint{}): getTableList.Init(apiRoot),	// /v1/devService/getTableList}
			// api.GetName(getUserList.EndPoint{}): getUserList.Init(apiRoot),	// /v1/userService/getUserList}
			// api.GetName(getWeather.EndPoint{}): getWeather.Init(apiRoot),	// /v1/powerStationService/getWeather}
			// api.GetName(getWorkTicketList.EndPoint{}): getWorkTicketList.Init(apiRoot),	// /v1/faultService/getWorkTicketList}
			// api.GetName(getWorkTicketListForTicketDetail.EndPoint{}): getWorkTicketListForTicketDetail.Init(apiRoot),	// /v1/faultService/getWorkTicketListForTicketDetail}
			// api.GetName(getWorkTicketRunningCount.EndPoint{}): getWorkTicketRunningCount.Init(apiRoot),	// /v1/faultService/getWorkTicketRunningCount}
			// api.GetName(getWorkTicketUserList.EndPoint{}): getWorkTicketUserList.Init(apiRoot),	// /v1/faultService/getWorkTicketUserList}
			// api.GetName(getinverterType.EndPoint{}): getinverterType.Init(apiRoot),	// /v1/devService/getinverterType}
			// api.GetName(getreportPermissionByUser.EndPoint{}): getreportPermissionByUser.Init(apiRoot),	// /v1/reportService/getPsIdByUserId}
			// api.GetName(handleValue.EndPoint{}): handleValue.Init(apiRoot),	// /v1/reportService/handleValue}
			// api.GetName(modifyDeviceInfo.EndPoint{}): modifyDeviceInfo.Init(apiRoot),	// /v1/devService/modifyDeviceInfo}
			// api.GetName(operaStoreSpareParts.EndPoint{}): operaStoreSpareParts.Init(apiRoot),	// /v1/otherService/operaStoreSpareParts}
			// api.GetName(operateBillTransferToUser.EndPoint{}): operateBillTransferToUser.Init(apiRoot),	// /v1/faultService/operateBillTransferToUser}
			// api.GetName(queryAllStockInventory.EndPoint{}): queryAllStockInventory.Init(apiRoot),	// /v1/otherService/queryAllStockInventory}
			// api.GetName(queryBatteryBoardsList.EndPoint{}): queryBatteryBoardsList.Init(apiRoot),	// /v1/devService/queryBatteryBoardsList}
			// api.GetName(queryBatteryBoardsPointsData.EndPoint{}): queryBatteryBoardsPointsData.Init(apiRoot),	// /v1/devService/queryBatteryBoardsPointsData}
			// api.GetName(queryCodeByType.EndPoint{}): queryCodeByType.Init(apiRoot),	// /v1/commonService/queryCodeByType}
			// api.GetName(queryDeviceInfoList.EndPoint{}): queryDeviceInfoList.Init(apiRoot),	// /v1/devService/queryDeviceInfoList}
			// api.GetName(queryElectricityCalendarData.EndPoint{}): queryElectricityCalendarData.Init(apiRoot),	// /v1/reportService/queryElectricityCalendarData}
			// api.GetName(queryFaultCodes.EndPoint{}): queryFaultCodes.Init(apiRoot),	// /v1/faultService/queryFaultCodes}
			// api.GetName(queryFaultLevelAndType.EndPoint{}): queryFaultLevelAndType.Init(apiRoot),	// /v1/faultService/queryFaultLevelAndType}
			// api.GetName(queryFaultNames.EndPoint{}): queryFaultNames.Init(apiRoot),	// /v1/faultService/queryFaultNames}
			// api.GetName(queryMaterialType.EndPoint{}): queryMaterialType.Init(apiRoot),	// /v1/otherService/queryMaterialType}
			// api.GetName(queryNounAndKlgList.EndPoint{}): queryNounAndKlgList.Init(apiRoot),	// /v1/faultService/queryNounAndKlgList}
			// api.GetName(queryNounList.EndPoint{}): queryNounList.Init(apiRoot),	// /v1/faultService/queryNounList}
			// api.GetName(queryOptTickctInfo.EndPoint{}): queryOptTickctInfo.Init(apiRoot),	// /v1/faultService/queryOptTickctInfo}
			// api.GetName(queryOrgIdByUser.EndPoint{}): queryOrgIdByUser.Init(apiRoot),	// /v1/userService/queryOrgIdByUser}
			// api.GetName(queryPsCountryList.EndPoint{}): queryPsCountryList.Init(apiRoot),	// /v1/commonService/queryPsCountryList}
			// api.GetName(queryPsProvcnList.EndPoint{}): queryPsProvcnList.Init(apiRoot),	// /v1/commonService/queryPsProvcnList}
			// api.GetName(queryPsTypeByPsId.EndPoint{}): queryPsTypeByPsId.Init(apiRoot),	// /v1/powerStationService/queryPsTypeByPsId}
			// api.GetName(querySparePartsList.EndPoint{}): querySparePartsList.Init(apiRoot),	// /v1/otherService/querySparePartsList}
			// api.GetName(queryStoreList.EndPoint{}): queryStoreList.Init(apiRoot),	// /v1/otherService/queryStoreList}
			// api.GetName(querySysTimezone.EndPoint{}): querySysTimezone.Init(apiRoot),	// /v1/commonService/querySysTimezone}
			// api.GetName(queryUnInventorySpareList.EndPoint{}): queryUnInventorySpareList.Init(apiRoot),	// /v1/otherService/queryUnInventorySpareList}
			api.GetName(queryUserCurveTemplateData.EndPoint{}): queryUserCurveTemplateData.Init(apiRoot),
			// api.GetName(renewOperation.EndPoint{}): renewOperation.Init(apiRoot),	// /v1/devService/renewOperation}
			// api.GetName(saveCustomReport.EndPoint{}): saveCustomReport.Init(apiRoot),	// /v1/reportService/saveCustomReport}
			// api.GetName(saveDutyInfo.EndPoint{}): saveDutyInfo.Init(apiRoot),	// /v1/otherService/saveDutyInfo}
			// api.GetName(saveInventory.EndPoint{}): saveInventory.Init(apiRoot),	// /v1/otherService/saveInventory}
			// api.GetName(saveMaterial.EndPoint{}): saveMaterial.Init(apiRoot),	// /v1/devService/saveMaterial}
			// api.GetName(saveSecondTypeTicket.EndPoint{}): saveSecondTypeTicket.Init(apiRoot),	// /v1/faultService/saveSecondTypeTicket}
			// api.GetName(saveSelfReportPoint.EndPoint{}): saveSelfReportPoint.Init(apiRoot),	// /v1/reportService/saveSelfReportPoint}
			// api.GetName(saveWorkTicket.EndPoint{}): saveWorkTicket.Init(apiRoot),	// /v1/faultService/saveWorkTicket}
			// api.GetName(secondTypeTicketFlowImplementStep.EndPoint{}): secondTypeTicketFlowImplementStep.Init(apiRoot),	// /v1/faultService/secondTypeTicketFlowImplementStep}
			// api.GetName(secondTypeTicketFlowTransferStep.EndPoint{}): secondTypeTicketFlowTransferStep.Init(apiRoot),	// /v1/faultService/secondTypeTicketFlowTransferStep}
			// api.GetName(secondTypeUpdateSign.EndPoint{}): secondTypeUpdateSign.Init(apiRoot),	// /v1/faultService/secondTypeUpdateSign}
			// api.GetName(selectPowerPageList.EndPoint{}): selectPowerPageList.Init(apiRoot),	// /v1/powerStationService/selectPowerPageList}
			// api.GetName(showAnalyzefxDetail.EndPoint{}): showAnalyzefxDetail.Init(apiRoot),	// /v1/reportService/showAnalyzefxDetail}
			// api.GetName(showFxReport.EndPoint{}): showFxReport.Init(apiRoot),	// /v1/reportService/showFxReport}
			// api.GetName(showMaterNameList.EndPoint{}): showMaterNameList.Init(apiRoot),	// /v1/otherService/showMaterNameList}
			// api.GetName(showMaterSubTypeList.EndPoint{}): showMaterSubTypeList.Init(apiRoot),	// /v1/otherService/showMaterSubTypeList}
			api.GetName(showPSView.EndPoint{}): showPSView.Init(apiRoot),
			// api.GetName(showTjReport.EndPoint{}): showTjReport.Init(apiRoot),	// /v1/reportService/showTjReport}
			// api.GetName(templateLikesInfo.EndPoint{}): templateLikesInfo.Init(apiRoot),	// /v1/faultService/templateLikesInfo}
			// api.GetName(updOptTicketInfo.EndPoint{}): updOptTicketInfo.Init(apiRoot),	// /v1/faultService/updOptTicketInfo}
			// api.GetName(updataWorkTicketAfterStartProcess.EndPoint{}): updataWorkTicketAfterStartProcess.Init(apiRoot),	// /v1/faultService/updataWorkTicketAfterStartProcess}
			// api.GetName(updateBillTicketForTask.EndPoint{}): updateBillTicketForTask.Init(apiRoot),	// /v1/faultService/updateBillTicketForTask}
			// api.GetName(updateDutyInfo.EndPoint{}): updateDutyInfo.Init(apiRoot),	// /v1/otherService/updateDutyInfo}
			// api.GetName(updateKnowledgeBaseUseNumber.EndPoint{}): updateKnowledgeBaseUseNumber.Init(apiRoot),	// /v1/faultService/updateKnowledgeBaseUseNumber}
			// api.GetName(updateMaterial.EndPoint{}): updateMaterial.Init(apiRoot),	// /v1/otherService/updateMaterial}
			// api.GetName(updateSpareParts.EndPoint{}): updateSpareParts.Init(apiRoot),	// /v1/otherService/updateSpareParts}
			// api.GetName(updateStopReason.EndPoint{}): updateStopReason.Init(apiRoot),	// /v1/devService/updateStopReason}
			// api.GetName(updateTemplate.EndPoint{}): updateTemplate.Init(apiRoot),	// /v1/faultService/updateKnowledgeBaseCitationCount}

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
