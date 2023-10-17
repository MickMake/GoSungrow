// Package WebAppService - API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebAppService

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/addMaterial"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/addOptTicketInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/addSpareParts"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/associateQueryFaultNames"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/auditPsDeviceCheck"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/calcOutputRankByDay"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/changeReadStatus"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/checkMaterialName"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/confirmFault"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/copeOperateTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/copySecondTypeTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/copyWorkTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/delOptTicketInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteDuty"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteDutyMid"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteMaterial"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteOrSharedSelfReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteSecondTypeTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteSparePartsById"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deleteWorkTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/deviceFactoryList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/dispartDataPageList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/executeTask"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findCurrentTask"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findDeviceMessageByPskey"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findFactoryMessage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findImgResources"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findMateiralSubType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findMaterialById"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findMyDealedCurrentTask"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findMyDealedImgResources"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findSeriesInverterData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/findWebRole"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getAllPsFaultCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getAllPsFaultCountByUserId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getAllPsList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getAllStore"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getBaseDeviceInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getBoxData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getCBoxTree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getCheckDevTypeList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getCheckUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getChildOrg"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getCo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getCodeTreeMap"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDST"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDataCounts"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDataInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevName"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevTypeList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDeviceDataList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDeviceFactory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDeviceInfoForCheck"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDeviceTreeChild"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDeviceUuid"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDutyInfoById"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getDutyOrgZtree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getElecEffectList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getEnvironmentInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultName"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultOrder"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultOrderByOrderId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultOrderList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getFaultOrderStepList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getHTRoleList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getHistoryComments"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInverteTableListCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInverterDiscreteDistributioList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInverterDiscreteList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInverterFactoryList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getInverterInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getLoadCurveList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getMqttConfigInfoByAppkey"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getMultiPowers"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOndutyQuery"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOperateTicketUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOptTicketsAttachments"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOrderCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOrderData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOrderDataSize"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOrgPsEquipmentList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getOrgPsPowerGenerationSummaryReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getParentUidChain"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerKwhkwpList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerPrList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerPredictionInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerTrendDayData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerTrendMonthData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerTrendYearData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPowerValue"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsBlock"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsBlockData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsBlockTree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsBoxListCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsCBoxDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsContact"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsDataVal"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsDeviceCheckList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsDeviceFaultList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsFaultList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsIdByUserId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsIdState"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsListForWorkTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsPictureMessage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsTicketSizeAndClockNum"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsTree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsTreeChild"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPsValue"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getPscSeriseData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getReportInfoByReportId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getReportListByType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getReportPsTree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getRoleList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSafeEffectList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSecondTypeTicketList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSecondTypeTicketListForTicketDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSelfReportPoint"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSparePartsDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getStatementList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getStoreByStationId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getSysUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getTableList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getWeather"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getWorkTicketList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getWorkTicketListForTicketDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getWorkTicketRunningCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getWorkTicketUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getinverterType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/getreportPermissionByUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/handleValue"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/modifyDeviceInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/operaStoreSpareParts"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/operateBillTransferToUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryAllStockInventory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryBatteryBoardsList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryBatteryBoardsPointsData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryCodeByType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryDeviceInfoList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryElectricityCalendarData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryFaultCodes"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryFaultLevelAndType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryFaultNames"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryMaterialType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryNounAndKlgList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryNounList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryOptTickctInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryOrgIdByUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryPsCountryList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryPsProvcnList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryPsTypeByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/querySparePartsList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryStoreList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/querySysTimezone"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryUnInventorySpareList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/renewOperation"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveCustomReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveDutyInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveInventory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveMaterial"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveSecondTypeTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveSelfReportPoint"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/saveWorkTicket"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/secondTypeTicketFlowImplementStep"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/secondTypeTicketFlowTransferStep"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/secondTypeUpdateSign"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/selectPowerPageList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showAnalyzefxDetail"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showFxReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showMaterNameList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showMaterSubTypeList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showPSView"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/showTjReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/templateLikesInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updOptTicketInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updataWorkTicketAfterStartProcess"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateBillTicketForTask"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateDutyInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateKnowledgeBaseUseNumber"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateMaterial"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateSpareParts"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateStopReason"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/updateTemplate"
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
			api.GetName(showPSView.EndPoint{}):                 showPSView.Init(apiRoot),
			api.GetName(queryUserCurveTemplateData.EndPoint{}): queryUserCurveTemplateData.Init(apiRoot),
			api.GetName(getDeviceUuid.EndPoint{}):              getDeviceUuid.Init(apiRoot), // /v1/devService/getDeviceUuid}

			// Discovered from Chrome Dev Tools
			api.GetName(getMqttConfigInfoByAppkey.EndPoint{}): getMqttConfigInfoByAppkey.Init(apiRoot), // /v1/devService/getDeviceUuid}

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
