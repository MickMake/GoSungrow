// Package WebAppService - API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebAppService

import (
	"GoSungrow/iSolarCloud/WebAppService/addMaterial"
	"GoSungrow/iSolarCloud/WebAppService/addOptTicketInfo"
	"GoSungrow/iSolarCloud/WebAppService/addSpareParts"
	"GoSungrow/iSolarCloud/WebAppService/associateQueryFaultNames"
	"GoSungrow/iSolarCloud/WebAppService/auditPsDeviceCheck"
	"GoSungrow/iSolarCloud/WebAppService/calcOutputRankByDay"
	"GoSungrow/iSolarCloud/WebAppService/changeReadStatus"
	"GoSungrow/iSolarCloud/WebAppService/checkMaterialName"
	"GoSungrow/iSolarCloud/WebAppService/confirmFault"
	"GoSungrow/iSolarCloud/WebAppService/copeOperateTicket"
	"GoSungrow/iSolarCloud/WebAppService/copySecondTypeTicket"
	"GoSungrow/iSolarCloud/WebAppService/copyWorkTicket"
	"GoSungrow/iSolarCloud/WebAppService/delOptTicketInfo"
	"GoSungrow/iSolarCloud/WebAppService/deleteDuty"
	"GoSungrow/iSolarCloud/WebAppService/deleteDutyMid"
	"GoSungrow/iSolarCloud/WebAppService/deleteMaterial"
	"GoSungrow/iSolarCloud/WebAppService/deleteOrSharedSelfReport"
	"GoSungrow/iSolarCloud/WebAppService/deleteSecondTypeTicket"
	"GoSungrow/iSolarCloud/WebAppService/deleteSparePartsById"
	"GoSungrow/iSolarCloud/WebAppService/deleteWorkTicket"
	"GoSungrow/iSolarCloud/WebAppService/deviceFactoryList"
	"GoSungrow/iSolarCloud/WebAppService/dispartDataPageList"
	"GoSungrow/iSolarCloud/WebAppService/executeTask"
	"GoSungrow/iSolarCloud/WebAppService/findCurrentTask"
	"GoSungrow/iSolarCloud/WebAppService/findDeviceMessageByPskey"
	"GoSungrow/iSolarCloud/WebAppService/findFactoryMessage"
	"GoSungrow/iSolarCloud/WebAppService/findImgResources"
	"GoSungrow/iSolarCloud/WebAppService/findMateiralSubType"
	"GoSungrow/iSolarCloud/WebAppService/findMaterialById"
	"GoSungrow/iSolarCloud/WebAppService/findMyDealedCurrentTask"
	"GoSungrow/iSolarCloud/WebAppService/findMyDealedImgResources"
	"GoSungrow/iSolarCloud/WebAppService/findSeriesInverterData"
	"GoSungrow/iSolarCloud/WebAppService/findWebRole"
	"GoSungrow/iSolarCloud/WebAppService/getAllPsFaultCount"
	"GoSungrow/iSolarCloud/WebAppService/getAllPsFaultCountByUserId"
	"GoSungrow/iSolarCloud/WebAppService/getAllPsList"
	"GoSungrow/iSolarCloud/WebAppService/getAllStore"
	"GoSungrow/iSolarCloud/WebAppService/getBaseDeviceInfo"
	"GoSungrow/iSolarCloud/WebAppService/getBoxData"
	"GoSungrow/iSolarCloud/WebAppService/getCBoxTree"
	"GoSungrow/iSolarCloud/WebAppService/getCheckDevTypeList"
	"GoSungrow/iSolarCloud/WebAppService/getCheckUserList"
	"GoSungrow/iSolarCloud/WebAppService/getChildOrg"
	"GoSungrow/iSolarCloud/WebAppService/getCo"
	"GoSungrow/iSolarCloud/WebAppService/getCodeTreeMap"
	"GoSungrow/iSolarCloud/WebAppService/getDST"
	"GoSungrow/iSolarCloud/WebAppService/getDataCounts"
	"GoSungrow/iSolarCloud/WebAppService/getDataInfo"
	"GoSungrow/iSolarCloud/WebAppService/getDevList"
	"GoSungrow/iSolarCloud/WebAppService/getDevName"
	"GoSungrow/iSolarCloud/WebAppService/getDevTypeList"
	"GoSungrow/iSolarCloud/WebAppService/getDeviceDataList"
	"GoSungrow/iSolarCloud/WebAppService/getDeviceFactory"
	"GoSungrow/iSolarCloud/WebAppService/getDeviceInfoForCheck"
	"GoSungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"GoSungrow/iSolarCloud/WebAppService/getDeviceTreeChild"
	"GoSungrow/iSolarCloud/WebAppService/getDeviceUuid"
	"GoSungrow/iSolarCloud/WebAppService/getDutyInfoById"
	"GoSungrow/iSolarCloud/WebAppService/getDutyOrgZtree"
	"GoSungrow/iSolarCloud/WebAppService/getElecEffectList"
	"GoSungrow/iSolarCloud/WebAppService/getEnvironmentInfo"
	"GoSungrow/iSolarCloud/WebAppService/getFaultList"
	"GoSungrow/iSolarCloud/WebAppService/getFaultName"
	"GoSungrow/iSolarCloud/WebAppService/getFaultOrder"
	"GoSungrow/iSolarCloud/WebAppService/getFaultOrderByOrderId"
	"GoSungrow/iSolarCloud/WebAppService/getFaultOrderList"
	"GoSungrow/iSolarCloud/WebAppService/getFaultOrderStepList"
	"GoSungrow/iSolarCloud/WebAppService/getHTRoleList"
	"GoSungrow/iSolarCloud/WebAppService/getHistoryComments"
	"GoSungrow/iSolarCloud/WebAppService/getInfo"
	"GoSungrow/iSolarCloud/WebAppService/getInverteTableListCount"
	"GoSungrow/iSolarCloud/WebAppService/getInverterDiscreteDistributioList"
	"GoSungrow/iSolarCloud/WebAppService/getInverterDiscreteList"
	"GoSungrow/iSolarCloud/WebAppService/getInverterFactoryList"
	"GoSungrow/iSolarCloud/WebAppService/getInverterInfo"
	"GoSungrow/iSolarCloud/WebAppService/getLoadCurveList"
	"GoSungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"GoSungrow/iSolarCloud/WebAppService/getMultiPowers"
	"GoSungrow/iSolarCloud/WebAppService/getOndutyQuery"
	"GoSungrow/iSolarCloud/WebAppService/getOperateTicketUserList"
	"GoSungrow/iSolarCloud/WebAppService/getOptTicketsAttachments"
	"GoSungrow/iSolarCloud/WebAppService/getOrderCount"
	"GoSungrow/iSolarCloud/WebAppService/getOrderData"
	"GoSungrow/iSolarCloud/WebAppService/getOrderDataSize"
	"GoSungrow/iSolarCloud/WebAppService/getOrgPsEquipmentList"
	"GoSungrow/iSolarCloud/WebAppService/getOrgPsPowerGenerationSummaryReport"
	"GoSungrow/iSolarCloud/WebAppService/getParentUidChain"
	"GoSungrow/iSolarCloud/WebAppService/getPowerKwhkwpList"
	"GoSungrow/iSolarCloud/WebAppService/getPowerPrList"
	"GoSungrow/iSolarCloud/WebAppService/getPowerPredictionInfo"
	"GoSungrow/iSolarCloud/WebAppService/getPowerTrendDayData"
	"GoSungrow/iSolarCloud/WebAppService/getPowerTrendMonthData"
	"GoSungrow/iSolarCloud/WebAppService/getPowerTrendYearData"
	"GoSungrow/iSolarCloud/WebAppService/getPowerValue"
	"GoSungrow/iSolarCloud/WebAppService/getPsBlock"
	"GoSungrow/iSolarCloud/WebAppService/getPsBlockData"
	"GoSungrow/iSolarCloud/WebAppService/getPsBlockTree"
	"GoSungrow/iSolarCloud/WebAppService/getPsBoxListCount"
	"GoSungrow/iSolarCloud/WebAppService/getPsCBoxDetail"
	"GoSungrow/iSolarCloud/WebAppService/getPsContact"
	"GoSungrow/iSolarCloud/WebAppService/getPsDataVal"
	"GoSungrow/iSolarCloud/WebAppService/getPsDeviceCheckList"
	"GoSungrow/iSolarCloud/WebAppService/getPsDeviceFaultList"
	"GoSungrow/iSolarCloud/WebAppService/getPsFaultList"
	"GoSungrow/iSolarCloud/WebAppService/getPsIdByUserId"
	"GoSungrow/iSolarCloud/WebAppService/getPsIdState"
	"GoSungrow/iSolarCloud/WebAppService/getPsList"
	"GoSungrow/iSolarCloud/WebAppService/getPsListForWorkTicket"
	"GoSungrow/iSolarCloud/WebAppService/getPsPictureMessage"
	"GoSungrow/iSolarCloud/WebAppService/getPsTicketSizeAndClockNum"
	"GoSungrow/iSolarCloud/WebAppService/getPsTree"
	"GoSungrow/iSolarCloud/WebAppService/getPsTreeChild"
	"GoSungrow/iSolarCloud/WebAppService/getPsUserList"
	"GoSungrow/iSolarCloud/WebAppService/getPsValue"
	"GoSungrow/iSolarCloud/WebAppService/getPscSeriseData"
	"GoSungrow/iSolarCloud/WebAppService/getReportInfoByReportId"
	"GoSungrow/iSolarCloud/WebAppService/getReportListByType"
	"GoSungrow/iSolarCloud/WebAppService/getReportPsTree"
	"GoSungrow/iSolarCloud/WebAppService/getRoleList"
	"GoSungrow/iSolarCloud/WebAppService/getSafeEffectList"
	"GoSungrow/iSolarCloud/WebAppService/getSecondTypeTicketList"
	"GoSungrow/iSolarCloud/WebAppService/getSecondTypeTicketListForTicketDetail"
	"GoSungrow/iSolarCloud/WebAppService/getSelfReportPoint"
	"GoSungrow/iSolarCloud/WebAppService/getSparePartsDetail"
	"GoSungrow/iSolarCloud/WebAppService/getStatementList"
	"GoSungrow/iSolarCloud/WebAppService/getStoreByStationId"
	"GoSungrow/iSolarCloud/WebAppService/getSysUserList"
	"GoSungrow/iSolarCloud/WebAppService/getTableList"
	"GoSungrow/iSolarCloud/WebAppService/getUserList"
	"GoSungrow/iSolarCloud/WebAppService/getWeather"
	"GoSungrow/iSolarCloud/WebAppService/getWorkTicketList"
	"GoSungrow/iSolarCloud/WebAppService/getWorkTicketListForTicketDetail"
	"GoSungrow/iSolarCloud/WebAppService/getWorkTicketRunningCount"
	"GoSungrow/iSolarCloud/WebAppService/getWorkTicketUserList"
	"GoSungrow/iSolarCloud/WebAppService/getinverterType"
	"GoSungrow/iSolarCloud/WebAppService/getreportPermissionByUser"
	"GoSungrow/iSolarCloud/WebAppService/handleValue"
	"GoSungrow/iSolarCloud/WebAppService/modifyDeviceInfo"
	"GoSungrow/iSolarCloud/WebAppService/operaStoreSpareParts"
	"GoSungrow/iSolarCloud/WebAppService/operateBillTransferToUser"
	"GoSungrow/iSolarCloud/WebAppService/queryAllStockInventory"
	"GoSungrow/iSolarCloud/WebAppService/queryBatteryBoardsList"
	"GoSungrow/iSolarCloud/WebAppService/queryBatteryBoardsPointsData"
	"GoSungrow/iSolarCloud/WebAppService/queryCodeByType"
	"GoSungrow/iSolarCloud/WebAppService/queryDeviceInfoList"
	"GoSungrow/iSolarCloud/WebAppService/queryElectricityCalendarData"
	"GoSungrow/iSolarCloud/WebAppService/queryFaultCodes"
	"GoSungrow/iSolarCloud/WebAppService/queryFaultLevelAndType"
	"GoSungrow/iSolarCloud/WebAppService/queryFaultNames"
	"GoSungrow/iSolarCloud/WebAppService/queryMaterialType"
	"GoSungrow/iSolarCloud/WebAppService/queryNounAndKlgList"
	"GoSungrow/iSolarCloud/WebAppService/queryNounList"
	"GoSungrow/iSolarCloud/WebAppService/queryOptTickctInfo"
	"GoSungrow/iSolarCloud/WebAppService/queryOrgIdByUser"
	"GoSungrow/iSolarCloud/WebAppService/queryPsCountryList"
	"GoSungrow/iSolarCloud/WebAppService/queryPsProvcnList"
	"GoSungrow/iSolarCloud/WebAppService/queryPsTypeByPsId"
	"GoSungrow/iSolarCloud/WebAppService/querySparePartsList"
	"GoSungrow/iSolarCloud/WebAppService/queryStoreList"
	"GoSungrow/iSolarCloud/WebAppService/querySysTimezone"
	"GoSungrow/iSolarCloud/WebAppService/queryUnInventorySpareList"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/WebAppService/renewOperation"
	"GoSungrow/iSolarCloud/WebAppService/saveCustomReport"
	"GoSungrow/iSolarCloud/WebAppService/saveDutyInfo"
	"GoSungrow/iSolarCloud/WebAppService/saveInventory"
	"GoSungrow/iSolarCloud/WebAppService/saveMaterial"
	"GoSungrow/iSolarCloud/WebAppService/saveSecondTypeTicket"
	"GoSungrow/iSolarCloud/WebAppService/saveSelfReportPoint"
	"GoSungrow/iSolarCloud/WebAppService/saveWorkTicket"
	"GoSungrow/iSolarCloud/WebAppService/secondTypeTicketFlowImplementStep"
	"GoSungrow/iSolarCloud/WebAppService/secondTypeTicketFlowTransferStep"
	"GoSungrow/iSolarCloud/WebAppService/secondTypeUpdateSign"
	"GoSungrow/iSolarCloud/WebAppService/selectPowerPageList"
	"GoSungrow/iSolarCloud/WebAppService/showAnalyzefxDetail"
	"GoSungrow/iSolarCloud/WebAppService/showFxReport"
	"GoSungrow/iSolarCloud/WebAppService/showMaterNameList"
	"GoSungrow/iSolarCloud/WebAppService/showMaterSubTypeList"
	"GoSungrow/iSolarCloud/WebAppService/showPSView"
	"GoSungrow/iSolarCloud/WebAppService/showTjReport"
	"GoSungrow/iSolarCloud/WebAppService/templateLikesInfo"
	"GoSungrow/iSolarCloud/WebAppService/updOptTicketInfo"
	"GoSungrow/iSolarCloud/WebAppService/updataWorkTicketAfterStartProcess"
	"GoSungrow/iSolarCloud/WebAppService/updateBillTicketForTask"
	"GoSungrow/iSolarCloud/WebAppService/updateDutyInfo"
	"GoSungrow/iSolarCloud/WebAppService/updateKnowledgeBaseUseNumber"
	"GoSungrow/iSolarCloud/WebAppService/updateMaterial"
	"GoSungrow/iSolarCloud/WebAppService/updateSpareParts"
	"GoSungrow/iSolarCloud/WebAppService/updateStopReason"
	"GoSungrow/iSolarCloud/WebAppService/updateTemplate"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"fmt"
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
			api.GetName(showPSView.EndPoint{}):                 showPSView.Init(apiRoot),
			api.GetName(queryUserCurveTemplateData.EndPoint{}): queryUserCurveTemplateData.Init(apiRoot),
			api.GetName(getDeviceUuid.EndPoint{}):              getDeviceUuid.Init(apiRoot), // /v1/devService/getDeviceUuid}

			// Discovered from Chrome Dev Tools
			api.GetName(getMqttConfigInfoByAppkey.EndPoint{}):              getMqttConfigInfoByAppkey.Init(apiRoot), // /v1/devService/getDeviceUuid}


			api.GetName(addMaterial.EndPoint{}):                            addMaterial.Init(apiRoot),                            // /v1/otherService/addMaterial}
			api.GetName(addOptTicketInfo.EndPoint{}):                       addOptTicketInfo.Init(apiRoot),                       // /v1/faultService/addOptTicketInfo}
			api.GetName(addSpareParts.EndPoint{}):                          addSpareParts.Init(apiRoot),                          // /v1/otherService/addSpareParts}
			api.GetName(associateQueryFaultNames.EndPoint{}):               associateQueryFaultNames.Init(apiRoot),               // /v1/faultService/associateQueryFaultNames}
			api.GetName(auditPsDeviceCheck.EndPoint{}):                     auditPsDeviceCheck.Init(apiRoot),                     // /v1/devService/auditPsDeviceCheck}
			api.GetName(calcOutputRankByDay.EndPoint{}):                    calcOutputRankByDay.Init(apiRoot),                    // /v1/powerStationService/calcOutputRankByDay}
			api.GetName(changeReadStatus.EndPoint{}):                       changeReadStatus.Init(apiRoot),                       // /v1/reportService/changeReadStatus}
			api.GetName(checkMaterialName.EndPoint{}):                      checkMaterialName.Init(apiRoot),                      // /v1/otherService/checkMaterialName}
			api.GetName(confirmFault.EndPoint{}):                           confirmFault.Init(apiRoot),                           // /v1/faultService/confirmOrCloseFault}
			api.GetName(copeOperateTicket.EndPoint{}):                      copeOperateTicket.Init(apiRoot),                      // /v1/faultService/copeOperateTicket}
			api.GetName(copySecondTypeTicket.EndPoint{}):                   copySecondTypeTicket.Init(apiRoot),                   // /v1/faultService/copySecondTypeTicket}
			api.GetName(copyWorkTicket.EndPoint{}):                         copyWorkTicket.Init(apiRoot),                         // /v1/faultService/copyWorkTicket}
			api.GetName(delOptTicketInfo.EndPoint{}):                       delOptTicketInfo.Init(apiRoot),                       // /v1/faultService/delOptTicketInfo}
			api.GetName(deleteDuty.EndPoint{}):                             deleteDuty.Init(apiRoot),                             // /v1/otherService/deleteDuty}
			api.GetName(deleteDutyMid.EndPoint{}):                          deleteDutyMid.Init(apiRoot),                          // /v1/otherService/deleteDutyMid}
			api.GetName(deleteMaterial.EndPoint{}):                         deleteMaterial.Init(apiRoot),                         // /v1/otherService/deleteMaterial}
			api.GetName(deleteOrSharedSelfReport.EndPoint{}):               deleteOrSharedSelfReport.Init(apiRoot),               // /v1/reportService/deleteOrSharedSelfReport}
			api.GetName(deleteSecondTypeTicket.EndPoint{}):                 deleteSecondTypeTicket.Init(apiRoot),                 // /v1/faultService/deleteSecondTypeTicket}
			api.GetName(deleteSparePartsById.EndPoint{}):                   deleteSparePartsById.Init(apiRoot),                   // /v1/otherService/deleteSparePartsById}
			api.GetName(deleteWorkTicket.EndPoint{}):                       deleteWorkTicket.Init(apiRoot),                       // /v1/faultService/deleteWorkTicket}
			api.GetName(deviceFactoryList.EndPoint{}):                      deviceFactoryList.Init(apiRoot),                      // /v1/devService/deviceFactoryList}
			api.GetName(dispartDataPageList.EndPoint{}):                    dispartDataPageList.Init(apiRoot),                    // /v1/commonService/dispartDataPageList}
			api.GetName(executeTask.EndPoint{}):                            executeTask.Init(apiRoot),                            // /v1/faultService/executeTask}
			api.GetName(findCurrentTask.EndPoint{}):                        findCurrentTask.Init(apiRoot),                        // /v1/faultService/findCurrentTask}
			api.GetName(findDeviceMessageByPskey.EndPoint{}):               findDeviceMessageByPskey.Init(apiRoot),               // /v1/devService/findDeviceMessageByPskey}
			api.GetName(findFactoryMessage.EndPoint{}):                     findFactoryMessage.Init(apiRoot),                     // /v1/devService/findFactoryMessage}
			api.GetName(findImgResources.EndPoint{}):                       findImgResources.Init(apiRoot),                       // /v1/faultService/findImgResources}
			api.GetName(findMateiralSubType.EndPoint{}):                    findMateiralSubType.Init(apiRoot),                    // /v1/otherService/findMateiralSubType}
			api.GetName(findMaterialById.EndPoint{}):                       findMaterialById.Init(apiRoot),                       // /v1/otherService/findMaterialById}
			api.GetName(findMyDealedCurrentTask.EndPoint{}):                findMyDealedCurrentTask.Init(apiRoot),                // /v1/faultService/findMyDealedCurrentTask}
			api.GetName(findMyDealedImgResources.EndPoint{}):               findMyDealedImgResources.Init(apiRoot),               // /v1/faultService/findMyDealedImgResources}
			api.GetName(findSeriesInverterData.EndPoint{}):                 findSeriesInverterData.Init(apiRoot),                 // /v1/devService/findSeriesInverterData}
			api.GetName(findWebRole.EndPoint{}):                            findWebRole.Init(apiRoot),                            // /v1/userService/findWebRole}
			api.GetName(getAllPsFaultCount.EndPoint{}):                     getAllPsFaultCount.Init(apiRoot),                     // /v1/faultService/getAllPsFaultCount}
			api.GetName(getAllPsFaultCountByUserId.EndPoint{}):             getAllPsFaultCountByUserId.Init(apiRoot),             // /v1/faultService/getAllPsFaultCountByUserId}
			api.GetName(getAllPsList.EndPoint{}):                           getAllPsList.Init(apiRoot),                           // /v1/powerStationService/getAllPsList}
			api.GetName(getAllStore.EndPoint{}):                            getAllStore.Init(apiRoot),                            // /v1/otherService/getAllStore}
			api.GetName(getBaseDeviceInfo.EndPoint{}):                      getBaseDeviceInfo.Init(apiRoot),                      // /v1/devService/getBaseDeviceInfo}
			api.GetName(getBoxData.EndPoint{}):                             getBoxData.Init(apiRoot),                             // /v1/devService/getBoxData}
			api.GetName(getCBoxTree.EndPoint{}):                            getCBoxTree.Init(apiRoot),                            // /v1/devService/getCBoxTree}
			api.GetName(getCheckDevTypeList.EndPoint{}):                    getCheckDevTypeList.Init(apiRoot),                    // /v1/devService/getCheckDevTypeList}
			api.GetName(getCheckUserList.EndPoint{}):                       getCheckUserList.Init(apiRoot),                       // /v1/userService/getCheckUserList}
			api.GetName(getChildOrg.EndPoint{}):                            getChildOrg.Init(apiRoot),                            // /v1/otherService/getChildOrg}
			api.GetName(getCo.EndPoint{}):                                  getCo.Init(apiRoot),                                  // /v1/powerStationService/getCo}
			api.GetName(getCodeTreeMap.EndPoint{}):                         getCodeTreeMap.Init(apiRoot),                         // /v1/commonService/getCodeTreeMap}
			api.GetName(getDST.EndPoint{}):                                 getDST.Init(apiRoot),                                 // /v1/userService/getDST}
			api.GetName(getDataCounts.EndPoint{}):                          getDataCounts.Init(apiRoot),                          // /v1/commonService/getDataCounts}
			api.GetName(getDataInfo.EndPoint{}):                            getDataInfo.Init(apiRoot),                            // /v1/powerStationService/getDataInfo}
			api.GetName(getDevList.EndPoint{}):                             getDevList.Init(apiRoot),                             // /v1/devService/getDevList}
			api.GetName(getDevName.EndPoint{}):                             getDevName.Init(apiRoot),                             // /v1/devService/getDevName}
			api.GetName(getDevTypeList.EndPoint{}):                         getDevTypeList.Init(apiRoot),                         // /v1/devService/getDevTypeList}
			api.GetName(getDeviceDataList.EndPoint{}):                      getDeviceDataList.Init(apiRoot),                      // /v1/reportService/getDeviceDataList}
			api.GetName(getDeviceFactory.EndPoint{}):                       getDeviceFactory.Init(apiRoot),                       // /v1/devService/getDeviceFactory}
			api.GetName(getDeviceInfoForCheck.EndPoint{}):                  getDeviceInfoForCheck.Init(apiRoot),                  // /v1/devService/getDeviceInfoForCheck}
			api.GetName(getDevicePointAttrs.EndPoint{}):                    getDevicePointAttrs.Init(apiRoot),                    // /v1/devService/getDevicePointAttrs}
			api.GetName(getDeviceTreeChild.EndPoint{}):                     getDeviceTreeChild.Init(apiRoot),                     // /v1/devService/getDeviceTreeChild}
			api.GetName(getDutyInfoById.EndPoint{}):                        getDutyInfoById.Init(apiRoot),                        // /v1/otherService/getDutyInfoById}
			api.GetName(getDutyOrgZtree.EndPoint{}):                        getDutyOrgZtree.Init(apiRoot),                        // /v1/otherService/getDutyOrgZtree}
			api.GetName(getElecEffectList.EndPoint{}):                      getElecEffectList.Init(apiRoot),                      // /v1/faultService/getElecEffectList}
			api.GetName(getEnvironmentInfo.EndPoint{}):                     getEnvironmentInfo.Init(apiRoot),                     // /v1/devService/getEnvironmentInfo}
			api.GetName(getFaultList.EndPoint{}):                           getFaultList.Init(apiRoot),                           // /v1/faultService/getFaultList}
			api.GetName(getFaultName.EndPoint{}):                           getFaultName.Init(apiRoot),                           // /v1/faultService/getFaultName}
			api.GetName(getFaultOrder.EndPoint{}):                          getFaultOrder.Init(apiRoot),                          // /v1/faultService/getFaultOrder}
			api.GetName(getFaultOrderByOrderId.EndPoint{}):                 getFaultOrderByOrderId.Init(apiRoot),                 // /v1/faultService/getFaultOrderByOrderId}
			api.GetName(getFaultOrderList.EndPoint{}):                      getFaultOrderList.Init(apiRoot),                      // /v1/faultService/getFaultOrderList}
			api.GetName(getFaultOrderStepList.EndPoint{}):                  getFaultOrderStepList.Init(apiRoot),                  // /v1/faultService/getFaultOrderStepList}
			api.GetName(getHTRoleList.EndPoint{}):                          getHTRoleList.Init(apiRoot),                          // /v1/userService/getHTRoleList}
			api.GetName(getHistoryComments.EndPoint{}):                     getHistoryComments.Init(apiRoot),                     // /v1/faultService/getHistoryComments}
			api.GetName(getInfo.EndPoint{}):                                getInfo.Init(apiRoot),                                // /v1/devService/getInfo}
			api.GetName(getInverteTableListCount.EndPoint{}):               getInverteTableListCount.Init(apiRoot),               // /v1/devService/getInverteTableListCount}
			api.GetName(getInverterDiscreteDistributioList.EndPoint{}):     getInverterDiscreteDistributioList.Init(apiRoot),     // /v1/reportService/getInverterDiscreteDistributioList}
			api.GetName(getInverterDiscreteList.EndPoint{}):                getInverterDiscreteList.Init(apiRoot),                // /v1/reportService/getInverterDiscreteList}
			api.GetName(getInverterFactoryList.EndPoint{}):                 getInverterFactoryList.Init(apiRoot),                 // /v1/reportService/getInverterFactoryList}
			api.GetName(getInverterInfo.EndPoint{}):                        getInverterInfo.Init(apiRoot),                        // /v1/devService/getInverterInfo}
			api.GetName(getLoadCurveList.EndPoint{}):                       getLoadCurveList.Init(apiRoot),                       // /v1/reportService/getLoadCurveList}
			api.GetName(getMultiPowers.EndPoint{}):                         getMultiPowers.Init(apiRoot),                         // /v1/devService/getMultiPowers}
			api.GetName(getOndutyQuery.EndPoint{}):                         getOndutyQuery.Init(apiRoot),                         // /v1/otherService/getOndutyQuery}
			api.GetName(getOperateTicketUserList.EndPoint{}):               getOperateTicketUserList.Init(apiRoot),               // /v1/faultService/getOperateTicketUserList}
			api.GetName(getOptTicketsAttachments.EndPoint{}):               getOptTicketsAttachments.Init(apiRoot),               // /v1/faultService/getOptTicketsAttachments}
			api.GetName(getOrderCount.EndPoint{}):                          getOrderCount.Init(apiRoot),                          // /v1/faultService/getOrderCount}
			api.GetName(getOrderData.EndPoint{}):                           getOrderData.Init(apiRoot),                           // /v1/faultService/getOrderData}
			api.GetName(getOrderDataSize.EndPoint{}):                       getOrderDataSize.Init(apiRoot),                       // /v1/faultService/getOrderDataSize}
			api.GetName(getOrgPsEquipmentList.EndPoint{}):                  getOrgPsEquipmentList.Init(apiRoot),                  // /v1/faultService/getOrgPsEquipmentList}
			api.GetName(getOrgPsPowerGenerationSummaryReport.EndPoint{}):   getOrgPsPowerGenerationSummaryReport.Init(apiRoot),   // /v1/reportService/getOrgPsPowerGenerationSummaryReport}
			api.GetName(getParentUidChain.EndPoint{}):                      getParentUidChain.Init(apiRoot),                      // /v1/devService/getParentUidChain}
			api.GetName(getPowerKwhkwpList.EndPoint{}):                     getPowerKwhkwpList.Init(apiRoot),                     // /v1/powerStationService/getPowerKwhkwpList}
			api.GetName(getPowerPrList.EndPoint{}):                         getPowerPrList.Init(apiRoot),                         // /v1/powerStationService/getPowerPrList}
			api.GetName(getPowerPredictionInfo.EndPoint{}):                 getPowerPredictionInfo.Init(apiRoot),                 // /v1/reportService/getPowerPredictionInfo}
			api.GetName(getPowerTrendDayData.EndPoint{}):                   getPowerTrendDayData.Init(apiRoot),                   // /v1/powerStationService/getPowerTrendDayData}
			api.GetName(getPowerTrendMonthData.EndPoint{}):                 getPowerTrendMonthData.Init(apiRoot),                 // /v1/powerStationService/getPowerTrendMonthData}
			api.GetName(getPowerTrendYearData.EndPoint{}):                  getPowerTrendYearData.Init(apiRoot),                  // /v1/powerStationService/getPowerTrendYearData}
			api.GetName(getPowerValue.EndPoint{}):                          getPowerValue.Init(apiRoot),                          // /v1/reportService/getPowerValue}
			api.GetName(getPsBlock.EndPoint{}):                             getPsBlock.Init(apiRoot),                             // /v1/devService/getPsBlock}
			api.GetName(getPsBlockData.EndPoint{}):                         getPsBlockData.Init(apiRoot),                         // /v1/devService/getPsBlockData}
			api.GetName(getPsBlockTree.EndPoint{}):                         getPsBlockTree.Init(apiRoot),                         // /v1/devService/getPsBlockTree}
			api.GetName(getPsBoxListCount.EndPoint{}):                      getPsBoxListCount.Init(apiRoot),                      // /v1/devService/getPsBoxListCount}
			api.GetName(getPsCBoxDetail.EndPoint{}):                        getPsCBoxDetail.Init(apiRoot),                        // /v1/devService/getPsCBoxDetail}
			api.GetName(getPsContact.EndPoint{}):                           getPsContact.Init(apiRoot),                           // /v1/powerStationService/getPsContact}
			api.GetName(getPsDataVal.EndPoint{}):                           getPsDataVal.Init(apiRoot),                           // /v1/devService/getPsDataVal}
			api.GetName(getPsDeviceCheckList.EndPoint{}):                   getPsDeviceCheckList.Init(apiRoot),                   // /v1/devService/getPsDeviceCheckList}
			api.GetName(getPsDeviceFaultList.EndPoint{}):                   getPsDeviceFaultList.Init(apiRoot),                   // /v1/faultService/getPsDeviceFaultList}
			api.GetName(getPsFaultList.EndPoint{}):                         getPsFaultList.Init(apiRoot),                         // /v1/faultService/getPsFaultList}
			api.GetName(getPsIdByUserId.EndPoint{}):                        getPsIdByUserId.Init(apiRoot),                        // /v1/powerStationService/getPsIdByUserId}
			api.GetName(getPsIdState.EndPoint{}):                           getPsIdState.Init(apiRoot),                           // /v1/powerStationService/getPsIdState}
			api.GetName(getPsList.EndPoint{}):                              getPsList.Init(apiRoot),                              // /v1/powerStationService/getPsListForWeb}
			api.GetName(getPsListForWorkTicket.EndPoint{}):                 getPsListForWorkTicket.Init(apiRoot),                 // /v1/faultService/getPsListForWorkTicket}
			api.GetName(getPsPictureMessage.EndPoint{}):                    getPsPictureMessage.Init(apiRoot),                    // /v1/powerStationService/getPsPictureMessage}
			api.GetName(getPsTicketSizeAndClockNum.EndPoint{}):             getPsTicketSizeAndClockNum.Init(apiRoot),             // /v1/faultService/getPsTicketSizeAndClockNum}
			api.GetName(getPsTree.EndPoint{}):                              getPsTree.Init(apiRoot),                              // /v1/devService/getPsTree}
			api.GetName(getPsTreeChild.EndPoint{}):                         getPsTreeChild.Init(apiRoot),                         // /v1/devService/getPsTreeChild}
			api.GetName(getPsUserList.EndPoint{}):                          getPsUserList.Init(apiRoot),                          // /v1/userService/getPsUserList}
			api.GetName(getPsValue.EndPoint{}):                             getPsValue.Init(apiRoot),                             // /v1/powerStationService/getPsValue}
			api.GetName(getPscSeriseData.EndPoint{}):                       getPscSeriseData.Init(apiRoot),                       // /v1/devService/getPscSeriseData}
			api.GetName(getReportInfoByReportId.EndPoint{}):                getReportInfoByReportId.Init(apiRoot),                // /v1/reportService/getReportInfoByReportId}
			api.GetName(getReportListByType.EndPoint{}):                    getReportListByType.Init(apiRoot),                    // /v1/reportService/getReportListByType}
			api.GetName(getReportPsTree.EndPoint{}):                        getReportPsTree.Init(apiRoot),                        // /v1/reportService/getReportPsTree}
			api.GetName(getRoleList.EndPoint{}):                            getRoleList.Init(apiRoot),                            // /v1/userService/getRoleList}
			api.GetName(getSafeEffectList.EndPoint{}):                      getSafeEffectList.Init(apiRoot),                      // /v1/faultService/getSafeEffectList}
			api.GetName(getSecondTypeTicketList.EndPoint{}):                getSecondTypeTicketList.Init(apiRoot),                // /v1/faultService/getSecondTypeTicketList}
			api.GetName(getSecondTypeTicketListForTicketDetail.EndPoint{}): getSecondTypeTicketListForTicketDetail.Init(apiRoot), // /v1/faultService/getSecondTypeTicketListForTicketDetail}
			api.GetName(getSelfReportPoint.EndPoint{}):                     getSelfReportPoint.Init(apiRoot),                     // /v1/reportService/getSelfReportPoint}
			api.GetName(getSparePartsDetail.EndPoint{}):                    getSparePartsDetail.Init(apiRoot),                    // /v1/otherService/getSparePartsDetail}
			api.GetName(getStatementList.EndPoint{}):                       getStatementList.Init(apiRoot),                       // /v1/reportService/getStatementList}
			api.GetName(getStoreByStationId.EndPoint{}):                    getStoreByStationId.Init(apiRoot),                    // /v1/otherService/getStoreByStationId}
			api.GetName(getSysUserList.EndPoint{}):                         getSysUserList.Init(apiRoot),                         // /v1/userService/getSysUserList}
			api.GetName(getTableList.EndPoint{}):                           getTableList.Init(apiRoot),                           // /v1/devService/getTableList}
			api.GetName(getUserList.EndPoint{}):                            getUserList.Init(apiRoot),                            // /v1/userService/getUserList}
			api.GetName(getWeather.EndPoint{}):                             getWeather.Init(apiRoot),                             // /v1/powerStationService/getWeather}
			api.GetName(getWorkTicketList.EndPoint{}):                      getWorkTicketList.Init(apiRoot),                      // /v1/faultService/getWorkTicketList}
			api.GetName(getWorkTicketListForTicketDetail.EndPoint{}):       getWorkTicketListForTicketDetail.Init(apiRoot),       // /v1/faultService/getWorkTicketListForTicketDetail}
			api.GetName(getWorkTicketRunningCount.EndPoint{}):              getWorkTicketRunningCount.Init(apiRoot),              // /v1/faultService/getWorkTicketRunningCount}
			api.GetName(getWorkTicketUserList.EndPoint{}):                  getWorkTicketUserList.Init(apiRoot),                  // /v1/faultService/getWorkTicketUserList}
			api.GetName(getinverterType.EndPoint{}):                        getinverterType.Init(apiRoot),                        // /v1/devService/getinverterType}
			api.GetName(getreportPermissionByUser.EndPoint{}):              getreportPermissionByUser.Init(apiRoot),              // /v1/reportService/getPsIdByUserId}
			api.GetName(handleValue.EndPoint{}):                            handleValue.Init(apiRoot),                            // /v1/reportService/handleValue}
			api.GetName(modifyDeviceInfo.EndPoint{}):                       modifyDeviceInfo.Init(apiRoot),                       // /v1/devService/modifyDeviceInfo}
			api.GetName(operaStoreSpareParts.EndPoint{}):                   operaStoreSpareParts.Init(apiRoot),                   // /v1/otherService/operaStoreSpareParts}
			api.GetName(operateBillTransferToUser.EndPoint{}):              operateBillTransferToUser.Init(apiRoot),              // /v1/faultService/operateBillTransferToUser}
			api.GetName(queryAllStockInventory.EndPoint{}):                 queryAllStockInventory.Init(apiRoot),                 // /v1/otherService/queryAllStockInventory}
			api.GetName(queryBatteryBoardsList.EndPoint{}):                 queryBatteryBoardsList.Init(apiRoot),                 // /v1/devService/queryBatteryBoardsList}
			api.GetName(queryBatteryBoardsPointsData.EndPoint{}):           queryBatteryBoardsPointsData.Init(apiRoot),           // /v1/devService/queryBatteryBoardsPointsData}
			api.GetName(queryCodeByType.EndPoint{}):                        queryCodeByType.Init(apiRoot),                        // /v1/commonService/queryCodeByType}
			api.GetName(queryDeviceInfoList.EndPoint{}):                    queryDeviceInfoList.Init(apiRoot),                    // /v1/devService/queryDeviceInfoList}
			api.GetName(queryElectricityCalendarData.EndPoint{}):           queryElectricityCalendarData.Init(apiRoot),           // /v1/reportService/queryElectricityCalendarData}
			api.GetName(queryFaultCodes.EndPoint{}):                        queryFaultCodes.Init(apiRoot),                        // /v1/faultService/queryFaultCodes}
			api.GetName(queryFaultLevelAndType.EndPoint{}):                 queryFaultLevelAndType.Init(apiRoot),                 // /v1/faultService/queryFaultLevelAndType}
			api.GetName(queryFaultNames.EndPoint{}):                        queryFaultNames.Init(apiRoot),                        // /v1/faultService/queryFaultNames}
			api.GetName(queryMaterialType.EndPoint{}):                      queryMaterialType.Init(apiRoot),                      // /v1/otherService/queryMaterialType}
			api.GetName(queryNounAndKlgList.EndPoint{}):                    queryNounAndKlgList.Init(apiRoot),                    // /v1/faultService/queryNounAndKlgList}
			api.GetName(queryNounList.EndPoint{}):                          queryNounList.Init(apiRoot),                          // /v1/faultService/queryNounList}
			api.GetName(queryOptTickctInfo.EndPoint{}):                     queryOptTickctInfo.Init(apiRoot),                     // /v1/faultService/queryOptTickctInfo}
			api.GetName(queryOrgIdByUser.EndPoint{}):                       queryOrgIdByUser.Init(apiRoot),                       // /v1/userService/queryOrgIdByUser}
			api.GetName(queryPsCountryList.EndPoint{}):                     queryPsCountryList.Init(apiRoot),                     // /v1/commonService/queryPsCountryList}
			api.GetName(queryPsProvcnList.EndPoint{}):                      queryPsProvcnList.Init(apiRoot),                      // /v1/commonService/queryPsProvcnList}
			api.GetName(queryPsTypeByPsId.EndPoint{}):                      queryPsTypeByPsId.Init(apiRoot),                      // /v1/powerStationService/queryPsTypeByPsId}
			api.GetName(querySparePartsList.EndPoint{}):                    querySparePartsList.Init(apiRoot),                    // /v1/otherService/querySparePartsList}
			api.GetName(queryStoreList.EndPoint{}):                         queryStoreList.Init(apiRoot),                         // /v1/otherService/queryStoreList}
			api.GetName(querySysTimezone.EndPoint{}):                       querySysTimezone.Init(apiRoot),                       // /v1/commonService/querySysTimezone}
			api.GetName(queryUnInventorySpareList.EndPoint{}):              queryUnInventorySpareList.Init(apiRoot),              // /v1/otherService/queryUnInventorySpareList}
			api.GetName(renewOperation.EndPoint{}):                         renewOperation.Init(apiRoot),                         // /v1/devService/renewOperation}
			api.GetName(saveCustomReport.EndPoint{}):                       saveCustomReport.Init(apiRoot),                       // /v1/reportService/saveCustomReport}
			api.GetName(saveDutyInfo.EndPoint{}):                           saveDutyInfo.Init(apiRoot),                           // /v1/otherService/saveDutyInfo}
			api.GetName(saveInventory.EndPoint{}):                          saveInventory.Init(apiRoot),                          // /v1/otherService/saveInventory}
			api.GetName(saveMaterial.EndPoint{}):                           saveMaterial.Init(apiRoot),                           // /v1/devService/saveMaterial}
			api.GetName(saveSecondTypeTicket.EndPoint{}):                   saveSecondTypeTicket.Init(apiRoot),                   // /v1/faultService/saveSecondTypeTicket}
			api.GetName(saveSelfReportPoint.EndPoint{}):                    saveSelfReportPoint.Init(apiRoot),                    // /v1/reportService/saveSelfReportPoint}
			api.GetName(saveWorkTicket.EndPoint{}):                         saveWorkTicket.Init(apiRoot),                         // /v1/faultService/saveWorkTicket}
			api.GetName(secondTypeTicketFlowImplementStep.EndPoint{}):      secondTypeTicketFlowImplementStep.Init(apiRoot),      // /v1/faultService/secondTypeTicketFlowImplementStep}
			api.GetName(secondTypeTicketFlowTransferStep.EndPoint{}):       secondTypeTicketFlowTransferStep.Init(apiRoot),       // /v1/faultService/secondTypeTicketFlowTransferStep}
			api.GetName(secondTypeUpdateSign.EndPoint{}):                   secondTypeUpdateSign.Init(apiRoot),                   // /v1/faultService/secondTypeUpdateSign}
			api.GetName(selectPowerPageList.EndPoint{}):                    selectPowerPageList.Init(apiRoot),                    // /v1/powerStationService/selectPowerPageList}
			api.GetName(showAnalyzefxDetail.EndPoint{}):                    showAnalyzefxDetail.Init(apiRoot),                    // /v1/reportService/showAnalyzefxDetail}
			api.GetName(showFxReport.EndPoint{}):                           showFxReport.Init(apiRoot),                           // /v1/reportService/showFxReport}
			api.GetName(showMaterNameList.EndPoint{}):                      showMaterNameList.Init(apiRoot),                      // /v1/otherService/showMaterNameList}
			api.GetName(showMaterSubTypeList.EndPoint{}):                   showMaterSubTypeList.Init(apiRoot),                   // /v1/otherService/showMaterSubTypeList}
			api.GetName(showTjReport.EndPoint{}):                           showTjReport.Init(apiRoot),                           // /v1/reportService/showTjReport}
			api.GetName(templateLikesInfo.EndPoint{}):                      templateLikesInfo.Init(apiRoot),                      // /v1/faultService/templateLikesInfo}
			api.GetName(updOptTicketInfo.EndPoint{}):                       updOptTicketInfo.Init(apiRoot),                       // /v1/faultService/updOptTicketInfo}
			api.GetName(updataWorkTicketAfterStartProcess.EndPoint{}):      updataWorkTicketAfterStartProcess.Init(apiRoot),      // /v1/faultService/updataWorkTicketAfterStartProcess}
			api.GetName(updateBillTicketForTask.EndPoint{}):                updateBillTicketForTask.Init(apiRoot),                // /v1/faultService/updateBillTicketForTask}
			api.GetName(updateDutyInfo.EndPoint{}):                         updateDutyInfo.Init(apiRoot),                         // /v1/otherService/updateDutyInfo}
			api.GetName(updateKnowledgeBaseUseNumber.EndPoint{}):           updateKnowledgeBaseUseNumber.Init(apiRoot),           // /v1/faultService/updateKnowledgeBaseUseNumber}
			api.GetName(updateMaterial.EndPoint{}):                         updateMaterial.Init(apiRoot),                         // /v1/otherService/updateMaterial}
			api.GetName(updateSpareParts.EndPoint{}):                       updateSpareParts.Init(apiRoot),                       // /v1/otherService/updateSpareParts}
			api.GetName(updateStopReason.EndPoint{}):                       updateStopReason.Init(apiRoot),                       // /v1/devService/updateStopReason}
			api.GetName(updateTemplate.EndPoint{}):                         updateTemplate.Init(apiRoot),                         // /v1/faultService/updateKnowledgeBaseCitationCount}
		},
	}

	return area
}

// ****************************************
// Methods not scoped by api.EndPoint interface type

//goland:noinspection GoUnusedExportedFunction
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

//goland:noinspection GoUnusedParameter
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

//goland:noinspection GoUnusedParameter
func (a Area) Call(name api.EndPointName) output.Json {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) GetRequest(name api.EndPointName) output.Json {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) GetResponse(name api.EndPointName) output.Json {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) GetData(name api.EndPointName) output.Json {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

//goland:noinspection GoUnusedParameter
func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
