package AppService

import (
	"GoSungrow/iSolarCloud/AppService/acceptPsSharing"
	"GoSungrow/iSolarCloud/AppService/activateEmail"
	"GoSungrow/iSolarCloud/AppService/addConfig"
	"GoSungrow/iSolarCloud/AppService/addDeviceRepair"
	"GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHousehold"
	"GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHouseholdByPsIdS"
	"GoSungrow/iSolarCloud/AppService/addFault"
	"GoSungrow/iSolarCloud/AppService/addFaultOrder"
	"GoSungrow/iSolarCloud/AppService/addFaultPlan"
	"GoSungrow/iSolarCloud/AppService/addFaultRepairSteps"
	"GoSungrow/iSolarCloud/AppService/addHouseholdEvaluation"
	"GoSungrow/iSolarCloud/AppService/addHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/AppService/addHouseholdOpinionFeedback"
	"GoSungrow/iSolarCloud/AppService/addHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/AppService/addOnDutyInfo"
	"GoSungrow/iSolarCloud/AppService/addOperRule"
	"GoSungrow/iSolarCloud/AppService/addOrDelPsStructure"
	"GoSungrow/iSolarCloud/AppService/addOrderStep"
	"GoSungrow/iSolarCloud/AppService/addPowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/addPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/addReportConfigEmail"
	"GoSungrow/iSolarCloud/AppService/addSysAdvancedParam"
	"GoSungrow/iSolarCloud/AppService/addSysOrgNew"
	"GoSungrow/iSolarCloud/AppService/aliPayAppTest"
	"GoSungrow/iSolarCloud/AppService/auditOperRule"
	"GoSungrow/iSolarCloud/AppService/batchAddStationBySn"
	"GoSungrow/iSolarCloud/AppService/batchImportSN"
	"GoSungrow/iSolarCloud/AppService/batchInsertUserAndOrg"
	"GoSungrow/iSolarCloud/AppService/batchModifyDevicesInfoAndPropertis"
	"GoSungrow/iSolarCloud/AppService/batchProcessPlantReport"
	"GoSungrow/iSolarCloud/AppService/batchUpdateDeviceSim"
	"GoSungrow/iSolarCloud/AppService/batchUpdateUserIsAgreeGdpr"
	"GoSungrow/iSolarCloud/AppService/boundMobilePhone"
	"GoSungrow/iSolarCloud/AppService/boundUserMail"
	"GoSungrow/iSolarCloud/AppService/caculateDeviceInputDiscrete"
	"GoSungrow/iSolarCloud/AppService/calculateDeviceDiscrete"
	"GoSungrow/iSolarCloud/AppService/calculateInitialCompensationData"
	"GoSungrow/iSolarCloud/AppService/cancelDeliverMail"
	"GoSungrow/iSolarCloud/AppService/cancelOrderScan"
	"GoSungrow/iSolarCloud/AppService/cancelParamSetTask"
	"GoSungrow/iSolarCloud/AppService/cancelPsSharing"
	"GoSungrow/iSolarCloud/AppService/cancelRechargeOrder"
	"GoSungrow/iSolarCloud/AppService/changRechargeOrderToCancel"
	"GoSungrow/iSolarCloud/AppService/changeHouseholdUser2Installer"
	"GoSungrow/iSolarCloud/AppService/changeRemoteParam"
	"GoSungrow/iSolarCloud/AppService/checkDealerOrgCode"
	"GoSungrow/iSolarCloud/AppService/checkDevSnIsBelongsToUser"
	"GoSungrow/iSolarCloud/AppService/checkInverterResult"
	"GoSungrow/iSolarCloud/AppService/checkIsCanDoParamSet"
	"GoSungrow/iSolarCloud/AppService/checkIsIvScan"
	"GoSungrow/iSolarCloud/AppService/checkOssObjectExist"
	"GoSungrow/iSolarCloud/AppService/checkServiceIsConnect"
	"GoSungrow/iSolarCloud/AppService/checkTechnicalParameters"
	"GoSungrow/iSolarCloud/AppService/checkUnitStatus"
	"GoSungrow/iSolarCloud/AppService/checkUpRechargeDevicePaying"
	"GoSungrow/iSolarCloud/AppService/checkUserAccountUnique"
	"GoSungrow/iSolarCloud/AppService/checkUserAccountUniqueAll"
	"GoSungrow/iSolarCloud/AppService/checkUserInfoUnique"
	"GoSungrow/iSolarCloud/AppService/checkUserIsExist"
	"GoSungrow/iSolarCloud/AppService/checkUserListIsExist"
	"GoSungrow/iSolarCloud/AppService/checkUserPassword"
	"GoSungrow/iSolarCloud/AppService/cloudDeploymentRecord"
	"GoSungrow/iSolarCloud/AppService/comfirmParamModel"
	"GoSungrow/iSolarCloud/AppService/communicationModuleDetail"
	"GoSungrow/iSolarCloud/AppService/compareValidateCode"
	"GoSungrow/iSolarCloud/AppService/componentInfo2Cloud"
	"GoSungrow/iSolarCloud/AppService/confirmFault"
	"GoSungrow/iSolarCloud/AppService/confirmIvFault"
	"GoSungrow/iSolarCloud/AppService/confirmReportConfig"
	"GoSungrow/iSolarCloud/AppService/createAppkeyInfo"
	"GoSungrow/iSolarCloud/AppService/createRenewInvoice"
	"GoSungrow/iSolarCloud/AppService/dealCommandReply"
	"GoSungrow/iSolarCloud/AppService/dealDeletePsFailPsDelete"
	"GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeSubTasks"
	"GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeTasks"
	"GoSungrow/iSolarCloud/AppService/dealFaultOrder"
	"GoSungrow/iSolarCloud/AppService/dealGroupStringDisableOrEnable"
	"GoSungrow/iSolarCloud/AppService/dealNumberOfServiceCalls2Mysql"
	"GoSungrow/iSolarCloud/AppService/dealParamSettingAfterComplete"
	"GoSungrow/iSolarCloud/AppService/dealPsDataSupplement"
	"GoSungrow/iSolarCloud/AppService/dealPsReportEmailSend"
	"GoSungrow/iSolarCloud/AppService/dealRemoteUpgrade"
	"GoSungrow/iSolarCloud/AppService/dealSnElectrifyCheck"
	"GoSungrow/iSolarCloud/AppService/dealSysDeviceSimFlowInfo"
	"GoSungrow/iSolarCloud/AppService/dealSysDeviceSimInfo"
	"GoSungrow/iSolarCloud/AppService/definiteTimeDealSnExpRemind"
	"GoSungrow/iSolarCloud/AppService/definiteTimeDealSnStatus"
	"GoSungrow/iSolarCloud/AppService/delDeviceRepair"
	"GoSungrow/iSolarCloud/AppService/delOperRule"
	"GoSungrow/iSolarCloud/AppService/delayCallApiResidueTimes"
	"GoSungrow/iSolarCloud/AppService/deleteComponent"
	"GoSungrow/iSolarCloud/AppService/deleteCustomerEmployee"
	"GoSungrow/iSolarCloud/AppService/deleteDeviceAccount"
	"GoSungrow/iSolarCloud/AppService/deleteDeviceSimById"
	"GoSungrow/iSolarCloud/AppService/deleteElectricitySettlementData"
	"GoSungrow/iSolarCloud/AppService/deleteFaultPlan"
	"GoSungrow/iSolarCloud/AppService/deleteFirmwareFiles"
	"GoSungrow/iSolarCloud/AppService/deleteHouseholdEvaluation"
	"GoSungrow/iSolarCloud/AppService/deleteHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/AppService/deleteHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/AppService/deleteInverterSnInChnnl"
	"GoSungrow/iSolarCloud/AppService/deleteModuleLog"
	"GoSungrow/iSolarCloud/AppService/deleteOnDutyInfo"
	"GoSungrow/iSolarCloud/AppService/deleteOperateBillFile"
	"GoSungrow/iSolarCloud/AppService/deleteOssObject"
	"GoSungrow/iSolarCloud/AppService/deletePowerDevicePointById"
	"GoSungrow/iSolarCloud/AppService/deletePowerPicture"
	"GoSungrow/iSolarCloud/AppService/deletePowerRobotInfoBySnAndPsId"
	"GoSungrow/iSolarCloud/AppService/deletePowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/AppService/deleteProductionData"
	"GoSungrow/iSolarCloud/AppService/deletePs"
	"GoSungrow/iSolarCloud/AppService/deleteRechargeOrder"
	"GoSungrow/iSolarCloud/AppService/deleteRegularlyConnectionInfo"
	"GoSungrow/iSolarCloud/AppService/deleteReportConfigEmailAddr"
	"GoSungrow/iSolarCloud/AppService/deleteSysAdvancedParam"
	"GoSungrow/iSolarCloud/AppService/deleteSysOrgNew"
	"GoSungrow/iSolarCloud/AppService/deleteTemplate"
	"GoSungrow/iSolarCloud/AppService/deleteUserInfoAllByUserId"
	"GoSungrow/iSolarCloud/AppService/deviceInputDiscreteDeleteTime"
	"GoSungrow/iSolarCloud/AppService/deviceInputDiscreteGetTime"
	"GoSungrow/iSolarCloud/AppService/deviceInputDiscreteInsertTime"
	"GoSungrow/iSolarCloud/AppService/deviceInputDiscreteUpdateTime"
	"GoSungrow/iSolarCloud/AppService/devicePointsDataFromMySql"
	"GoSungrow/iSolarCloud/AppService/deviceReplace"
	"GoSungrow/iSolarCloud/AppService/editDeviceRepair"
	"GoSungrow/iSolarCloud/AppService/editOperRule"
	"GoSungrow/iSolarCloud/AppService/energyPovertyAlleviation"
	"GoSungrow/iSolarCloud/AppService/energyTrend"
	"GoSungrow/iSolarCloud/AppService/exportParamSettingValPDF"
	"GoSungrow/iSolarCloud/AppService/exportPlantReportPDF"
	"GoSungrow/iSolarCloud/AppService/faultAutoClose"
	"GoSungrow/iSolarCloud/AppService/faultCloseRemindOrderHandler"
	"GoSungrow/iSolarCloud/AppService/findCodeValueList"
	"GoSungrow/iSolarCloud/AppService/findEmgOrgInfo"
	"GoSungrow/iSolarCloud/AppService/findEnvironmentInfo"
	"GoSungrow/iSolarCloud/AppService/findFromHbaseAndRedis"
	"GoSungrow/iSolarCloud/AppService/findInfoByuuid"
	"GoSungrow/iSolarCloud/AppService/findLossAnalysisList"
	"GoSungrow/iSolarCloud/AppService/findOnDutyInfo"
	"GoSungrow/iSolarCloud/AppService/findPsType"
	"GoSungrow/iSolarCloud/AppService/findSingleStationPR"
	"GoSungrow/iSolarCloud/AppService/findUserPassword"
	"GoSungrow/iSolarCloud/AppService/genTLSUserSigByUserAccount"
	"GoSungrow/iSolarCloud/AppService/generateRandomPassword"
	"GoSungrow/iSolarCloud/AppService/getAPIServiceInfo"
	"GoSungrow/iSolarCloud/AppService/getAccessedPermission"
	"GoSungrow/iSolarCloud/AppService/getAllDeviceByPsId"
	"GoSungrow/iSolarCloud/AppService/getAllPowerDeviceSetName"
	"GoSungrow/iSolarCloud/AppService/getAllPowerRobotViewInfoByPsId"
	"GoSungrow/iSolarCloud/AppService/getAllPsIdByOrgIds"
	"GoSungrow/iSolarCloud/AppService/getAllUserRemindCount"
	"GoSungrow/iSolarCloud/AppService/getAndOutletsAndUnit"
	"GoSungrow/iSolarCloud/AppService/getApiCallsForAppkeys"
	"GoSungrow/iSolarCloud/AppService/getAreaInfoCodeByCounty"
	"GoSungrow/iSolarCloud/AppService/getAreaList"
	"GoSungrow/iSolarCloud/AppService/getAutoCreatePowerStation"
	"GoSungrow/iSolarCloud/AppService/getBackReadValue"
	"GoSungrow/iSolarCloud/AppService/getBatchNewestPointData"
	"GoSungrow/iSolarCloud/AppService/getCallApiResidueTimes"
	"GoSungrow/iSolarCloud/AppService/getChangedPsListByTime"
	"GoSungrow/iSolarCloud/AppService/getChnnlListByPsId"
	"GoSungrow/iSolarCloud/AppService/getCloudList"
	"GoSungrow/iSolarCloud/AppService/getCloudServiceMappingConfig"
	"GoSungrow/iSolarCloud/AppService/getCommunicationDeviceConfigInfo"
	"GoSungrow/iSolarCloud/AppService/getCommunicationModuleMonitorData"
	"GoSungrow/iSolarCloud/AppService/getComponentModelFactory"
	"GoSungrow/iSolarCloud/AppService/getConfigList"
	"GoSungrow/iSolarCloud/AppService/getConnectionInfoBySnAndLocalPort"
	"GoSungrow/iSolarCloud/AppService/getCountDown"
	"GoSungrow/iSolarCloud/AppService/getCountryServiceInfo"
	"GoSungrow/iSolarCloud/AppService/getCounty"
	"GoSungrow/iSolarCloud/AppService/getCustomerEmployee"
	"GoSungrow/iSolarCloud/AppService/getCustomerList"
	"GoSungrow/iSolarCloud/AppService/getDataFromHBase"
	"GoSungrow/iSolarCloud/AppService/getDataFromHbaseByRowKey"
	"GoSungrow/iSolarCloud/AppService/getDevInstalledPowerByPsId"
	"GoSungrow/iSolarCloud/AppService/getDevRecord"
	"GoSungrow/iSolarCloud/AppService/getDevRunRecordList"
	"GoSungrow/iSolarCloud/AppService/getDevSimByList"
	"GoSungrow/iSolarCloud/AppService/getDevSimList"
	"GoSungrow/iSolarCloud/AppService/getDeviceAccountById"
	"GoSungrow/iSolarCloud/AppService/getDeviceFaultStatisticsData"
	"GoSungrow/iSolarCloud/AppService/getDeviceInfo"
	"GoSungrow/iSolarCloud/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/AppService/getDeviceModelInfoList"
	"GoSungrow/iSolarCloud/AppService/getDevicePointMinuteDataList"
	"GoSungrow/iSolarCloud/AppService/getDevicePoints"
	"GoSungrow/iSolarCloud/AppService/getDevicePropertys"
	"GoSungrow/iSolarCloud/AppService/getDeviceRepairDetail"
	"GoSungrow/iSolarCloud/AppService/getDeviceTechBranchCount"
	"GoSungrow/iSolarCloud/AppService/getDeviceTypeInfoList"
	"GoSungrow/iSolarCloud/AppService/getDeviceTypeList"
	"GoSungrow/iSolarCloud/AppService/getDstInfo"
	"GoSungrow/iSolarCloud/AppService/getElectricitySettlementData"
	"GoSungrow/iSolarCloud/AppService/getElectricitySettlementDetailData"
	"GoSungrow/iSolarCloud/AppService/getEncryptPublicKey"
	"GoSungrow/iSolarCloud/AppService/getFaultCount"
	"GoSungrow/iSolarCloud/AppService/getFaultDetail"
	"GoSungrow/iSolarCloud/AppService/getFaultMsgByFaultCode"
	"GoSungrow/iSolarCloud/AppService/getFaultMsgListByPageWithYYYYMM"
	"GoSungrow/iSolarCloud/AppService/getFaultMsgListWithYYYYMM"
	"GoSungrow/iSolarCloud/AppService/getFaultPlanList"
	"GoSungrow/iSolarCloud/AppService/getFileOperationRecordOne"
	"GoSungrow/iSolarCloud/AppService/getFormulaFaultAnalyseList"
	"GoSungrow/iSolarCloud/AppService/getGroupStringCheckResult"
	"GoSungrow/iSolarCloud/AppService/getGroupStringCheckRule"
	"GoSungrow/iSolarCloud/AppService/getHisData"
	"GoSungrow/iSolarCloud/AppService/getHistoryInfo"
	"GoSungrow/iSolarCloud/AppService/getHouseholdEvaluation"
	"GoSungrow/iSolarCloud/AppService/getHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/AppService/getHouseholdOpinionFeedback"
	"GoSungrow/iSolarCloud/AppService/getHouseholdPsInstallerByUserId"
	"GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getHouseholdUserInfo"
	"GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderInfo"
	"GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderList"
	"GoSungrow/iSolarCloud/AppService/getI18nConfigByType"
	"GoSungrow/iSolarCloud/AppService/getI18nFileInfo"
	"GoSungrow/iSolarCloud/AppService/getI18nInfoByKey"
	"GoSungrow/iSolarCloud/AppService/getI18nVersion"
	"GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"GoSungrow/iSolarCloud/AppService/getInfoFromAMap"
	"GoSungrow/iSolarCloud/AppService/getInfomationFromRedis"
	"GoSungrow/iSolarCloud/AppService/getInstallInfoList"
	"GoSungrow/iSolarCloud/AppService/getInstallerInfoByDealerOrgCodeOrId"
	"GoSungrow/iSolarCloud/AppService/getInvertDataList"
	"GoSungrow/iSolarCloud/AppService/getInverterDataCount"
	"GoSungrow/iSolarCloud/AppService/getInverterProcess"
	"GoSungrow/iSolarCloud/AppService/getInverterUuidBytotalId"
	"GoSungrow/iSolarCloud/AppService/getIvEchartsData"
	"GoSungrow/iSolarCloud/AppService/getIvEchartsDataById"
	"GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"GoSungrow/iSolarCloud/AppService/getListMiFromHBase"
	"GoSungrow/iSolarCloud/AppService/getMapInfo"
	"GoSungrow/iSolarCloud/AppService/getMapMiFromHBase"
	"GoSungrow/iSolarCloud/AppService/getMenuAndPrivileges"
	"GoSungrow/iSolarCloud/AppService/getMicrogridEStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getModuleLogInfo"
	"GoSungrow/iSolarCloud/AppService/getModuleLogTaskList"
	"GoSungrow/iSolarCloud/AppService/getNationProvJSON"
	"GoSungrow/iSolarCloud/AppService/getNeedOpAsynOpRecordList"
	"GoSungrow/iSolarCloud/AppService/getNoticeInfo"
	"GoSungrow/iSolarCloud/AppService/getNumberOfServiceCalls"
	"GoSungrow/iSolarCloud/AppService/getOSSConfig"
	"GoSungrow/iSolarCloud/AppService/getOperRuleDetail"
	"GoSungrow/iSolarCloud/AppService/getOperateBillFileId"
	"GoSungrow/iSolarCloud/AppService/getOperateTicketForDetail"
	"GoSungrow/iSolarCloud/AppService/getOrCreateNetEaseUserToken"
	"GoSungrow/iSolarCloud/AppService/getOrderDataList"
	"GoSungrow/iSolarCloud/AppService/getOrderDataSql2"
	"GoSungrow/iSolarCloud/AppService/getOrderDatas"
	"GoSungrow/iSolarCloud/AppService/getOrderDetail"
	"GoSungrow/iSolarCloud/AppService/getOrderStatistics"
	"GoSungrow/iSolarCloud/AppService/getOrgIdNameByUserId"
	"GoSungrow/iSolarCloud/AppService/getOrgInfoByDealerOrgCode"
	"GoSungrow/iSolarCloud/AppService/getOrgListByName"
	"GoSungrow/iSolarCloud/AppService/getOrgListByUserId"
	"GoSungrow/iSolarCloud/AppService/getOrgListForUser"
	"GoSungrow/iSolarCloud/AppService/getOssObjectStream"
	"GoSungrow/iSolarCloud/AppService/getOwnerFaultConfigList"
	"GoSungrow/iSolarCloud/AppService/getPListinfoFromMysql"
	"GoSungrow/iSolarCloud/AppService/getParamSetTemplate4NewProtocol"
	"GoSungrow/iSolarCloud/AppService/getParamSetTemplatePointInfo"
	"GoSungrow/iSolarCloud/AppService/getParamterSettingBase"
	"GoSungrow/iSolarCloud/AppService/getPhotoInfo"
	"GoSungrow/iSolarCloud/AppService/getPlanedOrNotPsList"
	"GoSungrow/iSolarCloud/AppService/getPlantReportPDFList"
	"GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTechList"
	"GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTree"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskDetailList"
	"GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskList"
	"GoSungrow/iSolarCloud/AppService/getPowerFormulaFaultAnalyse"
	"GoSungrow/iSolarCloud/AppService/getPowerPictureList"
	"GoSungrow/iSolarCloud/AppService/getPowerRobotInfoByRobotSn"
	"GoSungrow/iSolarCloud/AppService/getPowerRobotSweepAttrByPsId"
	"GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategyList"
	"GoSungrow/iSolarCloud/AppService/getPowerSettingCharges"
	"GoSungrow/iSolarCloud/AppService/getPowerSettingHistoryRecords"
	"GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationPR"
	"GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSql"
	"GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSqlCount"
	"GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/AppService/getPowerTrendDayData"
	"GoSungrow/iSolarCloud/AppService/getPrivateCloudValidityPeriod"
	"GoSungrow/iSolarCloud/AppService/getProvInfoListByNationCode"
	"GoSungrow/iSolarCloud/AppService/getPsAuthKey"
	"GoSungrow/iSolarCloud/AppService/getPsCurveInfo"
	"GoSungrow/iSolarCloud/AppService/getPsDataSupplementTaskList"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailByUserTokens"
	"GoSungrow/iSolarCloud/AppService/getPsDetailForSinglePage"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/AppService/getPsInstallerByPsId"
	"GoSungrow/iSolarCloud/AppService/getPsInstallerOrgInfoByPsId"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/getPsListByName"
	"GoSungrow/iSolarCloud/AppService/getPsListForPsDataByPsId"
	"GoSungrow/iSolarCloud/AppService/getPsListStaticData"
	"GoSungrow/iSolarCloud/AppService/getPsReport"
	"GoSungrow/iSolarCloud/AppService/getPsUser"
	"GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"GoSungrow/iSolarCloud/AppService/getRechargeOrderDetail"
	"GoSungrow/iSolarCloud/AppService/getRechargeOrderItemDeviceList"
	"GoSungrow/iSolarCloud/AppService/getRechargeOrderList"
	"GoSungrow/iSolarCloud/AppService/getRegionalTree"
	"GoSungrow/iSolarCloud/AppService/getRemoteParamSettingList"
	"GoSungrow/iSolarCloud/AppService/getRemoteUpgradeDeviceList"
	"GoSungrow/iSolarCloud/AppService/getRemoteUpgradeScheduleDetails"
	"GoSungrow/iSolarCloud/AppService/getRemoteUpgradeSubTasksList"
	"GoSungrow/iSolarCloud/AppService/getRemoteUpgradeTaskList"
	"GoSungrow/iSolarCloud/AppService/getReportData"
	"GoSungrow/iSolarCloud/AppService/getReportEmailConfigInfo"
	"GoSungrow/iSolarCloud/AppService/getReportExportColumns"
	"GoSungrow/iSolarCloud/AppService/getReportListByUserId"
	"GoSungrow/iSolarCloud/AppService/getRobotDynamicCleaningView"
	"GoSungrow/iSolarCloud/AppService/getRobotNumAndSweepCapacity"
	"GoSungrow/iSolarCloud/AppService/getRuleUnit"
	"GoSungrow/iSolarCloud/AppService/getSendReportConfigCron"
	"GoSungrow/iSolarCloud/AppService/getSerialNum"
	"GoSungrow/iSolarCloud/AppService/getShieldMapConditionList"
	"GoSungrow/iSolarCloud/AppService/getSimIdBySnList"
	"GoSungrow/iSolarCloud/AppService/getSingleIVData"
	"GoSungrow/iSolarCloud/AppService/getSnChangeRecord"
	"GoSungrow/iSolarCloud/AppService/getSnConnectionInfo"
	"GoSungrow/iSolarCloud/AppService/getStationInfoSql"
	"GoSungrow/iSolarCloud/AppService/getSungwsConfigCache"
	"GoSungrow/iSolarCloud/AppService/getSungwsGlobalConfigCache"
	"GoSungrow/iSolarCloud/AppService/getSweepDevParamSetTemplate"
	"GoSungrow/iSolarCloud/AppService/getSweepRobotDevList"
	"GoSungrow/iSolarCloud/AppService/getSysMsg"
	"GoSungrow/iSolarCloud/AppService/getSysOrgNewList"
	"GoSungrow/iSolarCloud/AppService/getSysOrgNewOne"
	"GoSungrow/iSolarCloud/AppService/getSysUserById"
	"GoSungrow/iSolarCloud/AppService/getTableDataSql"
	"GoSungrow/iSolarCloud/AppService/getTableDataSqlCount"
	"GoSungrow/iSolarCloud/AppService/getTemplateByInfoType"
	"GoSungrow/iSolarCloud/AppService/getTemplateList"
	"GoSungrow/iSolarCloud/AppService/getUUIDByUpuuid"
	"GoSungrow/iSolarCloud/AppService/getUpTimePoint"
	"GoSungrow/iSolarCloud/AppService/getUserById"
	"GoSungrow/iSolarCloud/AppService/getUserByInstaller"
	"GoSungrow/iSolarCloud/AppService/getUserDevOnlineOffineCount"
	"GoSungrow/iSolarCloud/AppService/getUserGDPRAttrs"
	"GoSungrow/iSolarCloud/AppService/getUserHavePowerStationCount"
	"GoSungrow/iSolarCloud/AppService/getUserInfoByUserAccounts"
	"GoSungrow/iSolarCloud/AppService/getUserList"
	"GoSungrow/iSolarCloud/AppService/getUserPsOrderList"
	"GoSungrow/iSolarCloud/AppService/getValFromHBase"
	"GoSungrow/iSolarCloud/AppService/getValidateCode"
	"GoSungrow/iSolarCloud/AppService/getValidateCodeAtRegister"
	"GoSungrow/iSolarCloud/AppService/getWeatherInfo"
	"GoSungrow/iSolarCloud/AppService/getWechatPushConfig"
	"GoSungrow/iSolarCloud/AppService/getWorkInfo"
	"GoSungrow/iSolarCloud/AppService/groupStringCheck"
	"GoSungrow/iSolarCloud/AppService/handleDevByCommunicationSN"
	"GoSungrow/iSolarCloud/AppService/householdResetPassBySN"
	"GoSungrow/iSolarCloud/AppService/immediatePayment"
	"GoSungrow/iSolarCloud/AppService/importExcelData"
	"GoSungrow/iSolarCloud/AppService/incomeStatistics"
	"GoSungrow/iSolarCloud/AppService/informPush"
	"GoSungrow/iSolarCloud/AppService/insertEmgOrgInfo"
	"GoSungrow/iSolarCloud/AppService/insightSynDeviceStructure2Cloud"
	"GoSungrow/iSolarCloud/AppService/intoDataToHbase"
	"GoSungrow/iSolarCloud/AppService/ipLocationQuery"
	"GoSungrow/iSolarCloud/AppService/isHave2GSn"
	"GoSungrow/iSolarCloud/AppService/judgeDevIsHasInitSetTemplate"
	"GoSungrow/iSolarCloud/AppService/judgeIsSettingMan"
	"GoSungrow/iSolarCloud/AppService/listOssFiles"
	"GoSungrow/iSolarCloud/AppService/loadAreaInfo"
	"GoSungrow/iSolarCloud/AppService/loadPowerStation"
	"GoSungrow/iSolarCloud/AppService/login"
	"GoSungrow/iSolarCloud/AppService/loginByToken"
	"GoSungrow/iSolarCloud/AppService/logout"
	"GoSungrow/iSolarCloud/AppService/mobilePhoneHasBound"
	"GoSungrow/iSolarCloud/AppService/modifiedDeviceInfo"
	"GoSungrow/iSolarCloud/AppService/modifyEmgOrgStruc"
	"GoSungrow/iSolarCloud/AppService/modifyFaultPlan"
	"GoSungrow/iSolarCloud/AppService/modifyOnDutyInfo"
	"GoSungrow/iSolarCloud/AppService/modifyPassword"
	"GoSungrow/iSolarCloud/AppService/modifyPersonalUnitList"
	"GoSungrow/iSolarCloud/AppService/modifyPsUser"
	"GoSungrow/iSolarCloud/AppService/moduleLogParamSet"
	"GoSungrow/iSolarCloud/AppService/operateOssFile"
	"GoSungrow/iSolarCloud/AppService/operationPowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/AppService/orgPowerReport"
	"GoSungrow/iSolarCloud/AppService/paramSetTryAgain"
	"GoSungrow/iSolarCloud/AppService/paramSetting"
	"GoSungrow/iSolarCloud/AppService/planPower"
	"GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"GoSungrow/iSolarCloud/AppService/powerTrendChartData"
	"GoSungrow/iSolarCloud/AppService/psForcastInfo"
	"GoSungrow/iSolarCloud/AppService/psHourPointsValue"
	"GoSungrow/iSolarCloud/AppService/queryAllPsIdAndName"
	"GoSungrow/iSolarCloud/AppService/queryBatchCreatePsTaskList"
	"GoSungrow/iSolarCloud/AppService/queryBatchSpeedyAddPowerStationResult"
	"GoSungrow/iSolarCloud/AppService/queryCardStatusCTCC"
	"GoSungrow/iSolarCloud/AppService/queryChildAccountList"
	"GoSungrow/iSolarCloud/AppService/queryCompensationRecordData"
	"GoSungrow/iSolarCloud/AppService/queryCompensationRecordList"
	"GoSungrow/iSolarCloud/AppService/queryComponent"
	"GoSungrow/iSolarCloud/AppService/queryComponentTechnicalParam"
	"GoSungrow/iSolarCloud/AppService/queryCountryGridAndRelation"
	"GoSungrow/iSolarCloud/AppService/queryCountryList"
	"GoSungrow/iSolarCloud/AppService/queryCtrlTaskById"
	"GoSungrow/iSolarCloud/AppService/queryDeviceInfo"
	"GoSungrow/iSolarCloud/AppService/queryDeviceInfoForApp"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListByUserId"
	"GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/AppService/queryDeviceModelTechnical"
	"GoSungrow/iSolarCloud/AppService/queryDevicePointDayMonthYearDataList"
	"GoSungrow/iSolarCloud/AppService/queryDevicePointMinuteDataList"
	"GoSungrow/iSolarCloud/AppService/queryDevicePointsDayMonthYearDataList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"GoSungrow/iSolarCloud/AppService/queryDeviceRepairList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceTypeInfoList"
	"GoSungrow/iSolarCloud/AppService/queryEnvironmentList"
	"GoSungrow/iSolarCloud/AppService/queryFaultList"
	"GoSungrow/iSolarCloud/AppService/queryFaultPlanDetail"
	"GoSungrow/iSolarCloud/AppService/queryFaultRepairSteps"
	"GoSungrow/iSolarCloud/AppService/queryFaultTypeAndLevelByCode"
	"GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevice"
	"GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevicePage"
	"GoSungrow/iSolarCloud/AppService/queryFirmwareFilesPage"
	"GoSungrow/iSolarCloud/AppService/queryInfotoAlert"
	"GoSungrow/iSolarCloud/AppService/queryInverterModelList"
	"GoSungrow/iSolarCloud/AppService/queryInverterVersionList"
	"GoSungrow/iSolarCloud/AppService/queryM2MCardInfoCMCC"
	"GoSungrow/iSolarCloud/AppService/queryM2MCardTermInfoCMCC"
	"GoSungrow/iSolarCloud/AppService/queryModelInfoByModelId"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/AppService/queryNoticeList"
	"GoSungrow/iSolarCloud/AppService/queryNumberOfRenewalReminders"
	"GoSungrow/iSolarCloud/AppService/queryOperRules"
	"GoSungrow/iSolarCloud/AppService/queryOrderList"
	"GoSungrow/iSolarCloud/AppService/queryOrderStep"
	"GoSungrow/iSolarCloud/AppService/queryOrgGenerationReport"
	"GoSungrow/iSolarCloud/AppService/queryOrgInfoList"
	"GoSungrow/iSolarCloud/AppService/queryOrgPowerElecPercent"
	"GoSungrow/iSolarCloud/AppService/queryOrgPsCompensationRecordList"
	"GoSungrow/iSolarCloud/AppService/queryParamSettingTask"
	"GoSungrow/iSolarCloud/AppService/queryPersonalUnitList"
	"GoSungrow/iSolarCloud/AppService/queryPointDataTopOne"
	"GoSungrow/iSolarCloud/AppService/queryPowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/queryPsAreaByUserIdAndAreaCode"
	"GoSungrow/iSolarCloud/AppService/queryPsCompensationRecordList"
	"GoSungrow/iSolarCloud/AppService/queryPsDataByDate"
	"GoSungrow/iSolarCloud/AppService/queryPsIdList"
	"GoSungrow/iSolarCloud/AppService/queryPsListByUserIdAndAreaCode"
	"GoSungrow/iSolarCloud/AppService/queryPsNameByPsId"
	"GoSungrow/iSolarCloud/AppService/queryPsPrByDate"
	"GoSungrow/iSolarCloud/AppService/queryPsProfit"
	"GoSungrow/iSolarCloud/AppService/queryPsReportComparativeAnalysisOfPowerGeneration"
	"GoSungrow/iSolarCloud/AppService/queryPsStructureList"
	"GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId"
	"GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId2"
	"GoSungrow/iSolarCloud/AppService/queryRepairRuleList"
	"GoSungrow/iSolarCloud/AppService/queryReportListForManagementPage"
	"GoSungrow/iSolarCloud/AppService/queryReportMsg"
	"GoSungrow/iSolarCloud/AppService/querySharingPs"
	"GoSungrow/iSolarCloud/AppService/querySysAdvancedParam"
	"GoSungrow/iSolarCloud/AppService/queryTimeBySN"
	"GoSungrow/iSolarCloud/AppService/queryTrafficByDateCTCC"
	"GoSungrow/iSolarCloud/AppService/queryTrafficCTCC"
	"GoSungrow/iSolarCloud/AppService/queryUnitList"
	"GoSungrow/iSolarCloud/AppService/queryUnitUuidBytotalId"
	"GoSungrow/iSolarCloud/AppService/queryUserBtnPri"
	"GoSungrow/iSolarCloud/AppService/queryUserByUserIds"
	"GoSungrow/iSolarCloud/AppService/queryUserExtensionAttribute"
	"GoSungrow/iSolarCloud/AppService/queryUserForStep"
	"GoSungrow/iSolarCloud/AppService/queryUserList"
	"GoSungrow/iSolarCloud/AppService/queryUserProcessPri"
	"GoSungrow/iSolarCloud/AppService/queryUserWechatBindRel"
	"GoSungrow/iSolarCloud/AppService/queryUuidByTotalIdAndUuid"
	"GoSungrow/iSolarCloud/AppService/rechargeOrderSetMeal"
	"GoSungrow/iSolarCloud/AppService/renewSendReportConfirmEmail"
	"GoSungrow/iSolarCloud/AppService/reportList"
	"GoSungrow/iSolarCloud/AppService/saveCustomerEmployee"
	"GoSungrow/iSolarCloud/AppService/saveDevSimList"
	"GoSungrow/iSolarCloud/AppService/saveDeviceAccountBatchData"
	"GoSungrow/iSolarCloud/AppService/saveEnviromentIncomeInfos"
	"GoSungrow/iSolarCloud/AppService/saveEnvironmentCurve"
	"GoSungrow/iSolarCloud/AppService/saveFirmwareFile"
	"GoSungrow/iSolarCloud/AppService/saveIncomeSettingInfos"
	"GoSungrow/iSolarCloud/AppService/saveOrUpdateGroupStringCheckRule"
	"GoSungrow/iSolarCloud/AppService/saveParamModel"
	"GoSungrow/iSolarCloud/AppService/savePowerCharges"
	"GoSungrow/iSolarCloud/AppService/savePowerDevicePoint"
	"GoSungrow/iSolarCloud/AppService/savePowerRobotInfo"
	"GoSungrow/iSolarCloud/AppService/savePowerRobotSweepAttr"
	"GoSungrow/iSolarCloud/AppService/savePowerSettingCharges"
	"GoSungrow/iSolarCloud/AppService/savePowerSettingInfo"
	"GoSungrow/iSolarCloud/AppService/saveProductionBatchData"
	"GoSungrow/iSolarCloud/AppService/saveRechargeOrderObj"
	"GoSungrow/iSolarCloud/AppService/saveRechargeOrderOtherInfo"
	"GoSungrow/iSolarCloud/AppService/saveRepair"
	"GoSungrow/iSolarCloud/AppService/saveReportExportColumns"
	"GoSungrow/iSolarCloud/AppService/saveSetParam"
	"GoSungrow/iSolarCloud/AppService/saveSysUserMsg"
	"GoSungrow/iSolarCloud/AppService/saveTemplate"
	"GoSungrow/iSolarCloud/AppService/searchM2MMonthFlowCMCC"
	"GoSungrow/iSolarCloud/AppService/selectSysTranslationNames"
	"GoSungrow/iSolarCloud/AppService/sendPsTimeZoneInstruction"
	"GoSungrow/iSolarCloud/AppService/setUpFormulaFaultAnalyse"
	"GoSungrow/iSolarCloud/AppService/setUserGDPRAttrs"
	"GoSungrow/iSolarCloud/AppService/settingNotice"
	"GoSungrow/iSolarCloud/AppService/shareMyPs"
	"GoSungrow/iSolarCloud/AppService/sharePsBySN"
	"GoSungrow/iSolarCloud/AppService/showInverterByUnit"
	"GoSungrow/iSolarCloud/AppService/showOnlineUsers"
	"GoSungrow/iSolarCloud/AppService/showWarning"
	"GoSungrow/iSolarCloud/AppService/snIsExist"
	"GoSungrow/iSolarCloud/AppService/snsIsExist"
	"GoSungrow/iSolarCloud/AppService/speedyAddPowerStation"
	"GoSungrow/iSolarCloud/AppService/stationDeviceHistoryDataList"
	"GoSungrow/iSolarCloud/AppService/stationUnitsList"
	"GoSungrow/iSolarCloud/AppService/stationsDiscreteData"
	"GoSungrow/iSolarCloud/AppService/stationsIncomeList"
	"GoSungrow/iSolarCloud/AppService/stationsPointReport"
	"GoSungrow/iSolarCloud/AppService/stationsYearPlanReport"
	"GoSungrow/iSolarCloud/AppService/sureAndImportSelettlementData"
	"GoSungrow/iSolarCloud/AppService/sweepDevParamSet"
	"GoSungrow/iSolarCloud/AppService/sweepDevRunControl"
	"GoSungrow/iSolarCloud/AppService/sweepDevStrategyIssue"
	"GoSungrow/iSolarCloud/AppService/sysTimeZoneList"
	"GoSungrow/iSolarCloud/AppService/unLockUser"
	"GoSungrow/iSolarCloud/AppService/unlockChildAccount"
	"GoSungrow/iSolarCloud/AppService/updateCommunicationModuleState"
	"GoSungrow/iSolarCloud/AppService/updateDevInstalledPower"
	"GoSungrow/iSolarCloud/AppService/updateFault"
	"GoSungrow/iSolarCloud/AppService/updateFaultData"
	"GoSungrow/iSolarCloud/AppService/updateFaultMsgByFaultCode"
	"GoSungrow/iSolarCloud/AppService/updateFaultStatus"
	"GoSungrow/iSolarCloud/AppService/updateHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/AppService/updateInverterSn2ModuleSn"
	"GoSungrow/iSolarCloud/AppService/updateOperateTicketAttachmentId"
	"GoSungrow/iSolarCloud/AppService/updateOrderDeviceByCustomerService"
	"GoSungrow/iSolarCloud/AppService/updateOwnerFaultConfig"
	"GoSungrow/iSolarCloud/AppService/updateParamSettingSysMsg"
	"GoSungrow/iSolarCloud/AppService/updatePlatformLevelFaultLevel"
	"GoSungrow/iSolarCloud/AppService/updatePowerDevicePoint"
	"GoSungrow/iSolarCloud/AppService/updatePowerRobotInfo"
	"GoSungrow/iSolarCloud/AppService/updatePowerRobotSweepAttr"
	"GoSungrow/iSolarCloud/AppService/updatePowerStationForHousehold"
	"GoSungrow/iSolarCloud/AppService/updatePowerStationInfo"
	"GoSungrow/iSolarCloud/AppService/updatePowerUserInfo"
	"GoSungrow/iSolarCloud/AppService/updateReportConfigByEmailAddr"
	"GoSungrow/iSolarCloud/AppService/updateShareAttr"
	"GoSungrow/iSolarCloud/AppService/updateSnIsSureFlag"
	"GoSungrow/iSolarCloud/AppService/updateStationPics"
	"GoSungrow/iSolarCloud/AppService/updateSysAdvancedParam"
	"GoSungrow/iSolarCloud/AppService/updateSysOrgNew"
	"GoSungrow/iSolarCloud/AppService/updateTemplate"
	"GoSungrow/iSolarCloud/AppService/updateUinfoNetEaseUser"
	"GoSungrow/iSolarCloud/AppService/updateUserExtensionAttribute"
	"GoSungrow/iSolarCloud/AppService/updateUserLanguage"
	"GoSungrow/iSolarCloud/AppService/updateUserPosition"
	"GoSungrow/iSolarCloud/AppService/updateUserUpOrg"
	"GoSungrow/iSolarCloud/AppService/upgrade"
	"GoSungrow/iSolarCloud/AppService/upgrate"
	"GoSungrow/iSolarCloud/AppService/uploadFileToOss"
	"GoSungrow/iSolarCloud/AppService/userAgreeGdprProtocol"
	"GoSungrow/iSolarCloud/AppService/userInfoUniqueCheck"
	"GoSungrow/iSolarCloud/AppService/userMailHasBound"
	"GoSungrow/iSolarCloud/AppService/userRegister"
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
			api.GetName(login.EndPoint{}):                       login.Init(apiRoot),
			api.GetName(findPsType.EndPoint{}):                  findPsType.Init(apiRoot),
			api.GetName(getPowerDevicePointNames.EndPoint{}):    getPowerDevicePointNames.Init(apiRoot),
			api.GetName(getPowerStatistics.EndPoint{}):          getPowerStatistics.Init(apiRoot),
			api.GetName(getPsDetailWithPsType.EndPoint{}):       getPsDetailWithPsType.Init(apiRoot),
			api.GetName(getPsList.EndPoint{}):                   getPsList.Init(apiRoot),
			api.GetName(getHouseholdStoragePsReport.EndPoint{}): getHouseholdStoragePsReport.Init(apiRoot),
			api.GetName(queryUnitList.EndPoint{}):               queryUnitList.Init(apiRoot),
			api.GetName(queryMutiPointDataList.EndPoint{}):      queryMutiPointDataList.Init(apiRoot),
			api.GetName(communicationModuleDetail.EndPoint{}):   communicationModuleDetail.Init(apiRoot),
			api.GetName(getHistoryInfo.EndPoint{}):              getHistoryInfo.Init(apiRoot), // @TODO - Not working.
			api.GetName(energyTrend.EndPoint{}):                 energyTrend.Init(apiRoot),
			api.GetName(getTemplateList.EndPoint{}):             getTemplateList.Init(apiRoot),
			api.GetName(getDevicePoints.EndPoint{}):             getDevicePoints.Init(apiRoot), // @TODO - Doesn't seem to generate anything.
			api.GetName(getPsDetail.EndPoint{}):                 getPsDetail.Init(apiRoot),
			api.GetName(getPsHealthState.EndPoint{}):            getPsHealthState.Init(apiRoot),
			api.GetName(queryAllPsIdAndName.EndPoint{}):         queryAllPsIdAndName.Init(apiRoot),
			api.GetName(getAllPowerDeviceSetName.EndPoint{}):    getAllPowerDeviceSetName.Init(apiRoot),
			api.GetName(getOrgListByName.EndPoint{}):            getOrgListByName.Init(apiRoot),
			api.GetName(getPsListByName.EndPoint{}):             getPsListByName.Init(apiRoot),
			api.GetName(checkUnitStatus.EndPoint{}):             checkUnitStatus.Init(apiRoot), // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceInfo.EndPoint{}):               getDeviceInfo.Init(apiRoot),   // @TODO - Returns null.
			api.GetName(getDeviceList.EndPoint{}):               getDeviceList.Init(apiRoot),
			api.GetName(powerDevicePointList.EndPoint{}):        powerDevicePointList.Init(apiRoot),  // @TODO - CRITICAL - GetByJson all point_id definitions.
			api.GetName(queryDeviceListForApp.EndPoint{}):       queryDeviceListForApp.Init(apiRoot), // @TODO - CRITICAL - Show more detail info on devices.
			api.GetName(queryDeviceListByUserId.EndPoint{}):     queryDeviceListByUserId.Init(apiRoot),
			api.GetName(getPsListStaticData.EndPoint{}):         getPsListStaticData.Init(apiRoot),
			api.GetName(getPsReport.EndPoint{}):                 getPsReport.Init(apiRoot),
			api.GetName(getPsUser.EndPoint{}):                   getPsUser.Init(apiRoot),
			api.GetName(getPsWeatherList.EndPoint{}):            getPsWeatherList.Init(apiRoot),
			api.GetName(getAreaList.EndPoint{}):                 getAreaList.Init(apiRoot),
			api.GetName(getCloudList.EndPoint{}):                getCloudList.Init(apiRoot),
			api.GetName(getConfigList.EndPoint{}):               getConfigList.Init(apiRoot),          // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceModelInfoList.EndPoint{}):      getDeviceModelInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device model info.
			api.GetName(getDeviceTypeList.EndPoint{}):           getDeviceTypeList.Init(apiRoot),
			api.GetName(getDeviceTypeInfoList.EndPoint{}):       getDeviceTypeInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device types.
			api.GetName(getInvertDataList.EndPoint{}):           getInvertDataList.Init(apiRoot),
			api.GetName(queryUserList.EndPoint{}):               queryUserList.Init(apiRoot),
			api.GetName(getPowerPictureList.EndPoint{}):         getPowerPictureList.Init(apiRoot),
			api.GetName(getUserList.EndPoint{}):                 getUserList.Init(apiRoot),
			api.GetName(getUserPsOrderList.EndPoint{}):          getUserPsOrderList.Init(apiRoot),
			api.GetName(reportList.EndPoint{}):                  reportList.Init(apiRoot),

			api.GetName(getKpiInfo.EndPoint{}): getKpiInfo.Init(apiRoot),
			api.GetName(getMapInfo.EndPoint{}): getMapInfo.Init(apiRoot), // @TODO - No data.

			// Tasks
			api.GetName(getPowerDeviceSetTaskDetailList.EndPoint{}): getPowerDeviceSetTaskDetailList.Init(apiRoot), // @TODO - Needs to be checked.
			api.GetName(getPowerDeviceSetTaskList.EndPoint{}):       getPowerDeviceSetTaskList.Init(apiRoot),
			api.GetName(getPsDataSupplementTaskList.EndPoint{}):     getPsDataSupplementTaskList.Init(apiRoot),
			api.GetName(queryBatchCreatePsTaskList.EndPoint{}):      queryBatchCreatePsTaskList.Init(apiRoot),
			api.GetName(getModuleLogTaskList.EndPoint{}):            getModuleLogTaskList.Init(apiRoot),
			api.GetName(getRemoteUpgradeTaskList.EndPoint{}):        getRemoteUpgradeTaskList.Init(apiRoot),     // @TODO - No data.
			api.GetName(getRemoteUpgradeSubTasksList.EndPoint{}):    getRemoteUpgradeSubTasksList.Init(apiRoot), // @TODO - No data.
			api.GetName(queryParamSettingTask.EndPoint{}):           queryParamSettingTask.Init(apiRoot),
			api.GetName(queryCtrlTaskById.EndPoint{}):               queryCtrlTaskById.Init(apiRoot), // @TODO - Returns 404 error. Disabled.

			// Use the following for all critical data every 5 minutes.
			api.GetName(queryDeviceList.EndPoint{}): queryDeviceList.Init(apiRoot), // @TODO - This gives you ALL info at 5 min intervals.

			// Disabled from here on.

			api.GetName(cancelParamSetTask.EndPoint{}):            cancelParamSetTask.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeSubTasks.EndPoint{}): dealFailRemoteUpgradeSubTasks.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeTasks.EndPoint{}):    dealFailRemoteUpgradeTasks.Init(apiRoot),

			api.GetName(updateTemplate.EndPoint{}):                  updateTemplate.Init(apiRoot),
			api.GetName(getTemplateByInfoType.EndPoint{}):           getTemplateByInfoType.Init(apiRoot),
			api.GetName(deleteTemplate.EndPoint{}):                  deleteTemplate.Init(apiRoot),
			api.GetName(getParamSetTemplate4NewProtocol.EndPoint{}): getParamSetTemplate4NewProtocol.Init(apiRoot),
			api.GetName(getParamSetTemplatePointInfo.EndPoint{}):    getParamSetTemplatePointInfo.Init(apiRoot),
			api.GetName(getSweepDevParamSetTemplate.EndPoint{}):     getSweepDevParamSetTemplate.Init(apiRoot),
			api.GetName(judgeDevIsHasInitSetTemplate.EndPoint{}):    judgeDevIsHasInitSetTemplate.Init(apiRoot),
			api.GetName(saveTemplate.EndPoint{}):                    saveTemplate.Init(apiRoot),

			api.GetName(queryDevicePointMinuteDataList.EndPoint{}):          queryDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(getDevicePointMinuteDataList.EndPoint{}):            getDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(queryDeviceInfo.EndPoint{}):                         queryDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceInfoForApp.EndPoint{}):                   queryDeviceInfoForApp.Init(apiRoot),
			api.GetName(queryDeviceModelTechnical.EndPoint{}):               queryDeviceModelTechnical.Init(apiRoot),
			api.GetName(queryDevicePointDayMonthYearDataList.EndPoint{}):    queryDevicePointDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDevicePointsDayMonthYearDataList.EndPoint{}):   queryDevicePointsDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDeviceRealTimeDataByPsKeys.EndPoint{}):         queryDeviceRealTimeDataByPsKeys.Init(apiRoot),
			api.GetName(addDeviceRepair.EndPoint{}):                         addDeviceRepair.Init(apiRoot),
			api.GetName(addDeviceToStructureForHousehold.EndPoint{}):        addDeviceToStructureForHousehold.Init(apiRoot),
			api.GetName(addDeviceToStructureForHouseholdByPsIdS.EndPoint{}): addDeviceToStructureForHouseholdByPsIdS.Init(apiRoot),
			api.GetName(batchModifyDevicesInfoAndPropertis.EndPoint{}):      batchModifyDevicesInfoAndPropertis.Init(apiRoot),
			api.GetName(batchUpdateDeviceSim.EndPoint{}):                    batchUpdateDeviceSim.Init(apiRoot),
			api.GetName(caculateDeviceInputDiscrete.EndPoint{}):             caculateDeviceInputDiscrete.Init(apiRoot),
			api.GetName(calculateDeviceDiscrete.EndPoint{}):                 calculateDeviceDiscrete.Init(apiRoot),
			api.GetName(checkUpRechargeDevicePaying.EndPoint{}):             checkUpRechargeDevicePaying.Init(apiRoot),
			api.GetName(dealSysDeviceSimFlowInfo.EndPoint{}):                dealSysDeviceSimFlowInfo.Init(apiRoot),
			api.GetName(dealSysDeviceSimInfo.EndPoint{}):                    dealSysDeviceSimInfo.Init(apiRoot),
			api.GetName(delDeviceRepair.EndPoint{}):                         delDeviceRepair.Init(apiRoot),
			api.GetName(deleteDeviceAccount.EndPoint{}):                     deleteDeviceAccount.Init(apiRoot),
			api.GetName(deleteDeviceSimById.EndPoint{}):                     deleteDeviceSimById.Init(apiRoot),
			api.GetName(deletePowerDevicePointById.EndPoint{}):              deletePowerDevicePointById.Init(apiRoot),
			api.GetName(editDeviceRepair.EndPoint{}):                        editDeviceRepair.Init(apiRoot),
			api.GetName(getAllDeviceByPsId.EndPoint{}):                      getAllDeviceByPsId.Init(apiRoot),
			api.GetName(getCommunicationDeviceConfigInfo.EndPoint{}):        getCommunicationDeviceConfigInfo.Init(apiRoot),
			api.GetName(getDeviceAccountById.EndPoint{}):                    getDeviceAccountById.Init(apiRoot),
			api.GetName(getDeviceFaultStatisticsData.EndPoint{}):            getDeviceFaultStatisticsData.Init(apiRoot),
			api.GetName(getDevicePropertys.EndPoint{}):                      getDevicePropertys.Init(apiRoot),
			api.GetName(getDeviceRepairDetail.EndPoint{}):                   getDeviceRepairDetail.Init(apiRoot),
			api.GetName(getDeviceTechBranchCount.EndPoint{}):                getDeviceTechBranchCount.Init(apiRoot),
			api.GetName(getPowerDeviceModelTechList.EndPoint{}):             getPowerDeviceModelTechList.Init(apiRoot),
			api.GetName(getPowerDeviceModelTree.EndPoint{}):                 getPowerDeviceModelTree.Init(apiRoot),
			api.GetName(getPowerDevicePointInfo.EndPoint{}):                 getPowerDevicePointInfo.Init(apiRoot),
			api.GetName(getRechargeOrderItemDeviceList.EndPoint{}):          getRechargeOrderItemDeviceList.Init(apiRoot),
			api.GetName(getRemoteUpgradeDeviceList.EndPoint{}):              getRemoteUpgradeDeviceList.Init(apiRoot),
			api.GetName(insightSynDeviceStructure2Cloud.EndPoint{}):         insightSynDeviceStructure2Cloud.Init(apiRoot),
			api.GetName(modifiedDeviceInfo.EndPoint{}):                      modifiedDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceRepairList.EndPoint{}):                   queryDeviceRepairList.Init(apiRoot),
			api.GetName(queryDeviceTypeInfoList.EndPoint{}):                 queryDeviceTypeInfoList.Init(apiRoot),
			api.GetName(queryFaultTypeByDevice.EndPoint{}):                  queryFaultTypeByDevice.Init(apiRoot),
			api.GetName(queryFaultTypeByDevicePage.EndPoint{}):              queryFaultTypeByDevicePage.Init(apiRoot),
			api.GetName(saveDeviceAccountBatchData.EndPoint{}):              saveDeviceAccountBatchData.Init(apiRoot),
			api.GetName(savePowerDevicePoint.EndPoint{}):                    savePowerDevicePoint.Init(apiRoot),
			api.GetName(stationDeviceHistoryDataList.EndPoint{}):            stationDeviceHistoryDataList.Init(apiRoot),
			api.GetName(updateOrderDeviceByCustomerService.EndPoint{}):      updateOrderDeviceByCustomerService.Init(apiRoot),
			api.GetName(updatePowerDevicePoint.EndPoint{}):                  updatePowerDevicePoint.Init(apiRoot),

			api.GetName(psHourPointsValue.EndPoint{}):                 psHourPointsValue.Init(apiRoot),
			api.GetName(getPsCurveInfo.EndPoint{}):                    getPsCurveInfo.Init(apiRoot),
			api.GetName(exportPlantReportPDF.EndPoint{}):              exportPlantReportPDF.Init(apiRoot),
			api.GetName(getTableDataSql.EndPoint{}):                   getTableDataSql.Init(apiRoot),
			api.GetName(getTableDataSqlCount.EndPoint{}):              getTableDataSqlCount.Init(apiRoot),
			api.GetName(getPListinfoFromMysql.EndPoint{}):             getPListinfoFromMysql.Init(apiRoot),
			api.GetName(getDataFromHBase.EndPoint{}):                  getDataFromHBase.Init(apiRoot),
			api.GetName(getListMiFromHBase.EndPoint{}):                getListMiFromHBase.Init(apiRoot),
			api.GetName(getMapMiFromHBase.EndPoint{}):                 getMapMiFromHBase.Init(apiRoot),
			api.GetName(getValFromHBase.EndPoint{}):                   getValFromHBase.Init(apiRoot),
			api.GetName(getPowerStationTableDataSql.EndPoint{}):       getPowerStationTableDataSql.Init(apiRoot),
			api.GetName(getPowerStationTableDataSqlCount.EndPoint{}):  getPowerStationTableDataSqlCount.Init(apiRoot),
			api.GetName(getDataFromHbaseByRowKey.EndPoint{}):          getDataFromHbaseByRowKey.Init(apiRoot),
			api.GetName(findInfoByuuid.EndPoint{}):                    findInfoByuuid.Init(apiRoot),
			api.GetName(getStationInfoSql.EndPoint{}):                 getStationInfoSql.Init(apiRoot),
			api.GetName(getPowerTrendDayData.EndPoint{}):              getPowerTrendDayData.Init(apiRoot),
			api.GetName(powerTrendChartData.EndPoint{}):               powerTrendChartData.Init(apiRoot),
			api.GetName(getUpTimePoint.EndPoint{}):                    getUpTimePoint.Init(apiRoot),
			api.GetName(getSerialNum.EndPoint{}):                      getSerialNum.Init(apiRoot),
			api.GetName(getPsAuthKey.EndPoint{}):                      getPsAuthKey.Init(apiRoot),
			api.GetName(getApiCallsForAppkeys.EndPoint{}):             getApiCallsForAppkeys.Init(apiRoot),
			api.GetName(exportParamSettingValPDF.EndPoint{}):          exportParamSettingValPDF.Init(apiRoot),
			api.GetName(devicePointsDataFromMySql.EndPoint{}):         devicePointsDataFromMySql.Init(apiRoot),
			api.GetName(acceptPsSharing.EndPoint{}):                   acceptPsSharing.Init(apiRoot),
			api.GetName(activateEmail.EndPoint{}):                     activateEmail.Init(apiRoot),
			api.GetName(addConfig.EndPoint{}):                         addConfig.Init(apiRoot),
			api.GetName(addFault.EndPoint{}):                          addFault.Init(apiRoot),
			api.GetName(addFaultOrder.EndPoint{}):                     addFaultOrder.Init(apiRoot),
			api.GetName(addFaultPlan.EndPoint{}):                      addFaultPlan.Init(apiRoot),
			api.GetName(addFaultRepairSteps.EndPoint{}):               addFaultRepairSteps.Init(apiRoot),
			api.GetName(addHouseholdEvaluation.EndPoint{}):            addHouseholdEvaluation.Init(apiRoot),
			api.GetName(addHouseholdLeaveMessage.EndPoint{}):          addHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(addHouseholdOpinionFeedback.EndPoint{}):       addHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(addHouseholdWorkOrder.EndPoint{}):             addHouseholdWorkOrder.Init(apiRoot),
			api.GetName(addOnDutyInfo.EndPoint{}):                     addOnDutyInfo.Init(apiRoot),
			api.GetName(addOperRule.EndPoint{}):                       addOperRule.Init(apiRoot),
			api.GetName(addOrDelPsStructure.EndPoint{}):               addOrDelPsStructure.Init(apiRoot),
			api.GetName(addOrderStep.EndPoint{}):                      addOrderStep.Init(apiRoot),
			api.GetName(addPowerStationForHousehold.EndPoint{}):       addPowerStationForHousehold.Init(apiRoot),
			api.GetName(addPowerStationInfo.EndPoint{}):               addPowerStationInfo.Init(apiRoot),
			api.GetName(addReportConfigEmail.EndPoint{}):              addReportConfigEmail.Init(apiRoot),
			api.GetName(addSysAdvancedParam.EndPoint{}):               addSysAdvancedParam.Init(apiRoot),
			api.GetName(addSysOrgNew.EndPoint{}):                      addSysOrgNew.Init(apiRoot),
			api.GetName(aliPayAppTest.EndPoint{}):                     aliPayAppTest.Init(apiRoot),
			api.GetName(auditOperRule.EndPoint{}):                     auditOperRule.Init(apiRoot),
			api.GetName(batchAddStationBySn.EndPoint{}):               batchAddStationBySn.Init(apiRoot),
			api.GetName(batchImportSN.EndPoint{}):                     batchImportSN.Init(apiRoot),
			api.GetName(batchInsertUserAndOrg.EndPoint{}):             batchInsertUserAndOrg.Init(apiRoot),
			api.GetName(batchProcessPlantReport.EndPoint{}):           batchProcessPlantReport.Init(apiRoot),
			api.GetName(batchUpdateUserIsAgreeGdpr.EndPoint{}):        batchUpdateUserIsAgreeGdpr.Init(apiRoot),
			api.GetName(boundMobilePhone.EndPoint{}):                  boundMobilePhone.Init(apiRoot),
			api.GetName(boundUserMail.EndPoint{}):                     boundUserMail.Init(apiRoot),
			api.GetName(calculateInitialCompensationData.EndPoint{}):  calculateInitialCompensationData.Init(apiRoot),
			api.GetName(cancelDeliverMail.EndPoint{}):                 cancelDeliverMail.Init(apiRoot),
			api.GetName(cancelOrderScan.EndPoint{}):                   cancelOrderScan.Init(apiRoot),
			api.GetName(cancelPsSharing.EndPoint{}):                   cancelPsSharing.Init(apiRoot),
			api.GetName(cancelRechargeOrder.EndPoint{}):               cancelRechargeOrder.Init(apiRoot),
			api.GetName(changRechargeOrderToCancel.EndPoint{}):        changRechargeOrderToCancel.Init(apiRoot),
			api.GetName(changeHouseholdUser2Installer.EndPoint{}):     changeHouseholdUser2Installer.Init(apiRoot),
			api.GetName(changeRemoteParam.EndPoint{}):                 changeRemoteParam.Init(apiRoot),
			api.GetName(checkDealerOrgCode.EndPoint{}):                checkDealerOrgCode.Init(apiRoot),
			api.GetName(checkDevSnIsBelongsToUser.EndPoint{}):         checkDevSnIsBelongsToUser.Init(apiRoot),
			api.GetName(checkInverterResult.EndPoint{}):               checkInverterResult.Init(apiRoot),
			api.GetName(checkIsCanDoParamSet.EndPoint{}):              checkIsCanDoParamSet.Init(apiRoot),
			api.GetName(checkIsIvScan.EndPoint{}):                     checkIsIvScan.Init(apiRoot),
			api.GetName(checkOssObjectExist.EndPoint{}):               checkOssObjectExist.Init(apiRoot),
			api.GetName(checkServiceIsConnect.EndPoint{}):             checkServiceIsConnect.Init(apiRoot),
			api.GetName(checkTechnicalParameters.EndPoint{}):          checkTechnicalParameters.Init(apiRoot),
			api.GetName(checkUserAccountUnique.EndPoint{}):            checkUserAccountUnique.Init(apiRoot),
			api.GetName(checkUserAccountUniqueAll.EndPoint{}):         checkUserAccountUniqueAll.Init(apiRoot),
			api.GetName(checkUserInfoUnique.EndPoint{}):               checkUserInfoUnique.Init(apiRoot),
			api.GetName(checkUserIsExist.EndPoint{}):                  checkUserIsExist.Init(apiRoot),
			api.GetName(checkUserListIsExist.EndPoint{}):              checkUserListIsExist.Init(apiRoot),
			api.GetName(checkUserPassword.EndPoint{}):                 checkUserPassword.Init(apiRoot),
			api.GetName(cloudDeploymentRecord.EndPoint{}):             cloudDeploymentRecord.Init(apiRoot),
			api.GetName(comfirmParamModel.EndPoint{}):                 comfirmParamModel.Init(apiRoot),
			api.GetName(compareValidateCode.EndPoint{}):               compareValidateCode.Init(apiRoot),
			api.GetName(componentInfo2Cloud.EndPoint{}):               componentInfo2Cloud.Init(apiRoot),
			api.GetName(confirmFault.EndPoint{}):                      confirmFault.Init(apiRoot),
			api.GetName(confirmIvFault.EndPoint{}):                    confirmIvFault.Init(apiRoot),
			api.GetName(confirmReportConfig.EndPoint{}):               confirmReportConfig.Init(apiRoot),
			api.GetName(createAppkeyInfo.EndPoint{}):                  createAppkeyInfo.Init(apiRoot),
			api.GetName(createRenewInvoice.EndPoint{}):                createRenewInvoice.Init(apiRoot),
			api.GetName(dealCommandReply.EndPoint{}):                  dealCommandReply.Init(apiRoot),
			api.GetName(dealDeletePsFailPsDelete.EndPoint{}):          dealDeletePsFailPsDelete.Init(apiRoot),
			api.GetName(dealFaultOrder.EndPoint{}):                    dealFaultOrder.Init(apiRoot),
			api.GetName(dealGroupStringDisableOrEnable.EndPoint{}):    dealGroupStringDisableOrEnable.Init(apiRoot),
			api.GetName(dealNumberOfServiceCalls2Mysql.EndPoint{}):    dealNumberOfServiceCalls2Mysql.Init(apiRoot),
			api.GetName(dealParamSettingAfterComplete.EndPoint{}):     dealParamSettingAfterComplete.Init(apiRoot),
			api.GetName(dealPsDataSupplement.EndPoint{}):              dealPsDataSupplement.Init(apiRoot),
			api.GetName(dealPsReportEmailSend.EndPoint{}):             dealPsReportEmailSend.Init(apiRoot),
			api.GetName(dealRemoteUpgrade.EndPoint{}):                 dealRemoteUpgrade.Init(apiRoot),
			api.GetName(dealSnElectrifyCheck.EndPoint{}):              dealSnElectrifyCheck.Init(apiRoot),
			api.GetName(definiteTimeDealSnExpRemind.EndPoint{}):       definiteTimeDealSnExpRemind.Init(apiRoot),
			api.GetName(definiteTimeDealSnStatus.EndPoint{}):          definiteTimeDealSnStatus.Init(apiRoot),
			api.GetName(delOperRule.EndPoint{}):                       delOperRule.Init(apiRoot),
			api.GetName(delayCallApiResidueTimes.EndPoint{}):          delayCallApiResidueTimes.Init(apiRoot),
			api.GetName(deleteComponent.EndPoint{}):                   deleteComponent.Init(apiRoot),
			api.GetName(deleteCustomerEmployee.EndPoint{}):            deleteCustomerEmployee.Init(apiRoot),
			api.GetName(deleteElectricitySettlementData.EndPoint{}):   deleteElectricitySettlementData.Init(apiRoot),
			api.GetName(deleteFaultPlan.EndPoint{}):                   deleteFaultPlan.Init(apiRoot),
			api.GetName(deleteFirmwareFiles.EndPoint{}):               deleteFirmwareFiles.Init(apiRoot),
			api.GetName(deleteHouseholdEvaluation.EndPoint{}):         deleteHouseholdEvaluation.Init(apiRoot),
			api.GetName(deleteHouseholdLeaveMessage.EndPoint{}):       deleteHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(deleteHouseholdWorkOrder.EndPoint{}):          deleteHouseholdWorkOrder.Init(apiRoot),
			api.GetName(deleteInverterSnInChnnl.EndPoint{}):           deleteInverterSnInChnnl.Init(apiRoot),
			api.GetName(deleteModuleLog.EndPoint{}):                   deleteModuleLog.Init(apiRoot),
			api.GetName(deleteOnDutyInfo.EndPoint{}):                  deleteOnDutyInfo.Init(apiRoot),
			api.GetName(deleteOperateBillFile.EndPoint{}):             deleteOperateBillFile.Init(apiRoot),
			api.GetName(deleteOssObject.EndPoint{}):                   deleteOssObject.Init(apiRoot),
			api.GetName(deletePowerPicture.EndPoint{}):                deletePowerPicture.Init(apiRoot),
			api.GetName(deletePowerRobotInfoBySnAndPsId.EndPoint{}):   deletePowerRobotInfoBySnAndPsId.Init(apiRoot),
			api.GetName(deletePowerRobotSweepStrategy.EndPoint{}):     deletePowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(deleteProductionData.EndPoint{}):              deleteProductionData.Init(apiRoot),
			api.GetName(deletePs.EndPoint{}):                          deletePs.Init(apiRoot),
			api.GetName(deleteRechargeOrder.EndPoint{}):               deleteRechargeOrder.Init(apiRoot),
			api.GetName(deleteRegularlyConnectionInfo.EndPoint{}):     deleteRegularlyConnectionInfo.Init(apiRoot),
			api.GetName(deleteReportConfigEmailAddr.EndPoint{}):       deleteReportConfigEmailAddr.Init(apiRoot),
			api.GetName(deleteSysAdvancedParam.EndPoint{}):            deleteSysAdvancedParam.Init(apiRoot),
			api.GetName(deleteSysOrgNew.EndPoint{}):                   deleteSysOrgNew.Init(apiRoot),
			api.GetName(deleteUserInfoAllByUserId.EndPoint{}):         deleteUserInfoAllByUserId.Init(apiRoot),
			api.GetName(deviceInputDiscreteDeleteTime.EndPoint{}):     deviceInputDiscreteDeleteTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteGetTime.EndPoint{}):        deviceInputDiscreteGetTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteInsertTime.EndPoint{}):     deviceInputDiscreteInsertTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteUpdateTime.EndPoint{}):     deviceInputDiscreteUpdateTime.Init(apiRoot),
			api.GetName(deviceReplace.EndPoint{}):                     deviceReplace.Init(apiRoot),
			api.GetName(editOperRule.EndPoint{}):                      editOperRule.Init(apiRoot),
			api.GetName(energyPovertyAlleviation.EndPoint{}):          energyPovertyAlleviation.Init(apiRoot),
			api.GetName(faultAutoClose.EndPoint{}):                    faultAutoClose.Init(apiRoot),
			api.GetName(faultCloseRemindOrderHandler.EndPoint{}):      faultCloseRemindOrderHandler.Init(apiRoot),
			api.GetName(findCodeValueList.EndPoint{}):                 findCodeValueList.Init(apiRoot),
			api.GetName(findEmgOrgInfo.EndPoint{}):                    findEmgOrgInfo.Init(apiRoot),
			api.GetName(findEnvironmentInfo.EndPoint{}):               findEnvironmentInfo.Init(apiRoot),
			api.GetName(findFromHbaseAndRedis.EndPoint{}):             findFromHbaseAndRedis.Init(apiRoot),
			api.GetName(findLossAnalysisList.EndPoint{}):              findLossAnalysisList.Init(apiRoot),
			api.GetName(findOnDutyInfo.EndPoint{}):                    findOnDutyInfo.Init(apiRoot),
			api.GetName(findSingleStationPR.EndPoint{}):               findSingleStationPR.Init(apiRoot),
			api.GetName(findUserPassword.EndPoint{}):                  findUserPassword.Init(apiRoot),
			api.GetName(genTLSUserSigByUserAccount.EndPoint{}):        genTLSUserSigByUserAccount.Init(apiRoot),
			api.GetName(generateRandomPassword.EndPoint{}):            generateRandomPassword.Init(apiRoot),
			api.GetName(getAPIServiceInfo.EndPoint{}):                 getAPIServiceInfo.Init(apiRoot),
			api.GetName(getAccessedPermission.EndPoint{}):             getAccessedPermission.Init(apiRoot),
			api.GetName(getAllPowerRobotViewInfoByPsId.EndPoint{}):    getAllPowerRobotViewInfoByPsId.Init(apiRoot),
			api.GetName(getAllPsIdByOrgIds.EndPoint{}):                getAllPsIdByOrgIds.Init(apiRoot),
			api.GetName(getAllUserRemindCount.EndPoint{}):             getAllUserRemindCount.Init(apiRoot),
			api.GetName(getAndOutletsAndUnit.EndPoint{}):              getAndOutletsAndUnit.Init(apiRoot),
			api.GetName(getAreaInfoCodeByCounty.EndPoint{}):           getAreaInfoCodeByCounty.Init(apiRoot),
			api.GetName(getAutoCreatePowerStation.EndPoint{}):         getAutoCreatePowerStation.Init(apiRoot),
			api.GetName(getBackReadValue.EndPoint{}):                  getBackReadValue.Init(apiRoot),
			api.GetName(getBatchNewestPointData.EndPoint{}):           getBatchNewestPointData.Init(apiRoot),
			api.GetName(getCallApiResidueTimes.EndPoint{}):            getCallApiResidueTimes.Init(apiRoot),
			api.GetName(getChangedPsListByTime.EndPoint{}):            getChangedPsListByTime.Init(apiRoot),
			api.GetName(getChnnlListByPsId.EndPoint{}):                getChnnlListByPsId.Init(apiRoot),
			api.GetName(getCloudServiceMappingConfig.EndPoint{}):      getCloudServiceMappingConfig.Init(apiRoot),
			api.GetName(getCommunicationModuleMonitorData.EndPoint{}): getCommunicationModuleMonitorData.Init(apiRoot),
			api.GetName(getComponentModelFactory.EndPoint{}):          getComponentModelFactory.Init(apiRoot),

			api.GetName(getConnectionInfoBySnAndLocalPort.EndPoint{}):                 getConnectionInfoBySnAndLocalPort.Init(apiRoot),
			api.GetName(getCountDown.EndPoint{}):                                      getCountDown.Init(apiRoot),
			api.GetName(getCountryServiceInfo.EndPoint{}):                             getCountryServiceInfo.Init(apiRoot),
			api.GetName(getCounty.EndPoint{}):                                         getCounty.Init(apiRoot),
			api.GetName(getCustomerEmployee.EndPoint{}):                               getCustomerEmployee.Init(apiRoot),
			api.GetName(getCustomerList.EndPoint{}):                                   getCustomerList.Init(apiRoot),
			api.GetName(getDevInstalledPowerByPsId.EndPoint{}):                        getDevInstalledPowerByPsId.Init(apiRoot),
			api.GetName(getDevRecord.EndPoint{}):                                      getDevRecord.Init(apiRoot),
			api.GetName(getDevRunRecordList.EndPoint{}):                               getDevRunRecordList.Init(apiRoot),
			api.GetName(getDevSimByList.EndPoint{}):                                   getDevSimByList.Init(apiRoot),
			api.GetName(getDevSimList.EndPoint{}):                                     getDevSimList.Init(apiRoot),
			api.GetName(getDstInfo.EndPoint{}):                                        getDstInfo.Init(apiRoot),
			api.GetName(getElectricitySettlementData.EndPoint{}):                      getElectricitySettlementData.Init(apiRoot),
			api.GetName(getElectricitySettlementDetailData.EndPoint{}):                getElectricitySettlementDetailData.Init(apiRoot),
			api.GetName(getEncryptPublicKey.EndPoint{}):                               getEncryptPublicKey.Init(apiRoot),
			api.GetName(getFaultCount.EndPoint{}):                                     getFaultCount.Init(apiRoot),
			api.GetName(getFaultDetail.EndPoint{}):                                    getFaultDetail.Init(apiRoot),
			api.GetName(getFaultMsgByFaultCode.EndPoint{}):                            getFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(getFaultMsgListByPageWithYYYYMM.EndPoint{}):                   getFaultMsgListByPageWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultMsgListWithYYYYMM.EndPoint{}):                         getFaultMsgListWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultPlanList.EndPoint{}):                                  getFaultPlanList.Init(apiRoot),
			api.GetName(getFileOperationRecordOne.EndPoint{}):                         getFileOperationRecordOne.Init(apiRoot),
			api.GetName(getFormulaFaultAnalyseList.EndPoint{}):                        getFormulaFaultAnalyseList.Init(apiRoot),
			api.GetName(getGroupStringCheckResult.EndPoint{}):                         getGroupStringCheckResult.Init(apiRoot),
			api.GetName(getGroupStringCheckRule.EndPoint{}):                           getGroupStringCheckRule.Init(apiRoot),
			api.GetName(getHisData.EndPoint{}):                                        getHisData.Init(apiRoot),
			api.GetName(getHouseholdEvaluation.EndPoint{}):                            getHouseholdEvaluation.Init(apiRoot),
			api.GetName(getHouseholdLeaveMessage.EndPoint{}):                          getHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(getHouseholdOpinionFeedback.EndPoint{}):                       getHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(getHouseholdPsInstallerByUserId.EndPoint{}):                   getHouseholdPsInstallerByUserId.Init(apiRoot),
			api.GetName(getHouseholdUserInfo.EndPoint{}):                              getHouseholdUserInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderInfo.EndPoint{}):                         getHouseholdWorkOrderInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderList.EndPoint{}):                         getHouseholdWorkOrderList.Init(apiRoot),
			api.GetName(getI18nConfigByType.EndPoint{}):                               getI18nConfigByType.Init(apiRoot),
			api.GetName(getI18nFileInfo.EndPoint{}):                                   getI18nFileInfo.Init(apiRoot),
			api.GetName(getI18nInfoByKey.EndPoint{}):                                  getI18nInfoByKey.Init(apiRoot),
			api.GetName(getI18nVersion.EndPoint{}):                                    getI18nVersion.Init(apiRoot),
			api.GetName(getIncomeSettingInfos.EndPoint{}):                             getIncomeSettingInfos.Init(apiRoot),
			api.GetName(getInfoFromAMap.EndPoint{}):                                   getInfoFromAMap.Init(apiRoot),
			api.GetName(getInfomationFromRedis.EndPoint{}):                            getInfomationFromRedis.Init(apiRoot),
			api.GetName(getInstallInfoList.EndPoint{}):                                getInstallInfoList.Init(apiRoot),
			api.GetName(getInstallerInfoByDealerOrgCodeOrId.EndPoint{}):               getInstallerInfoByDealerOrgCodeOrId.Init(apiRoot),
			api.GetName(getInverterDataCount.EndPoint{}):                              getInverterDataCount.Init(apiRoot),
			api.GetName(getInverterProcess.EndPoint{}):                                getInverterProcess.Init(apiRoot),
			api.GetName(getInverterUuidBytotalId.EndPoint{}):                          getInverterUuidBytotalId.Init(apiRoot),
			api.GetName(getIvEchartsData.EndPoint{}):                                  getIvEchartsData.Init(apiRoot),
			api.GetName(getIvEchartsDataById.EndPoint{}):                              getIvEchartsDataById.Init(apiRoot),
			api.GetName(getMenuAndPrivileges.EndPoint{}):                              getMenuAndPrivileges.Init(apiRoot),
			api.GetName(getMicrogridEStoragePsReport.EndPoint{}):                      getMicrogridEStoragePsReport.Init(apiRoot),
			api.GetName(getModuleLogInfo.EndPoint{}):                                  getModuleLogInfo.Init(apiRoot),
			api.GetName(getNationProvJSON.EndPoint{}):                                 getNationProvJSON.Init(apiRoot),
			api.GetName(getNeedOpAsynOpRecordList.EndPoint{}):                         getNeedOpAsynOpRecordList.Init(apiRoot),
			api.GetName(getNoticeInfo.EndPoint{}):                                     getNoticeInfo.Init(apiRoot),
			api.GetName(getNumberOfServiceCalls.EndPoint{}):                           getNumberOfServiceCalls.Init(apiRoot),
			api.GetName(getOSSConfig.EndPoint{}):                                      getOSSConfig.Init(apiRoot),
			api.GetName(getOperRuleDetail.EndPoint{}):                                 getOperRuleDetail.Init(apiRoot),
			api.GetName(getOperateBillFileId.EndPoint{}):                              getOperateBillFileId.Init(apiRoot),
			api.GetName(getOperateTicketForDetail.EndPoint{}):                         getOperateTicketForDetail.Init(apiRoot),
			api.GetName(getOrCreateNetEaseUserToken.EndPoint{}):                       getOrCreateNetEaseUserToken.Init(apiRoot),
			api.GetName(getOrderDataList.EndPoint{}):                                  getOrderDataList.Init(apiRoot),
			api.GetName(getOrderDataSql2.EndPoint{}):                                  getOrderDataSql2.Init(apiRoot),
			api.GetName(getOrderDatas.EndPoint{}):                                     getOrderDatas.Init(apiRoot),
			api.GetName(getOrderDetail.EndPoint{}):                                    getOrderDetail.Init(apiRoot),
			api.GetName(getOrderStatistics.EndPoint{}):                                getOrderStatistics.Init(apiRoot),
			api.GetName(getOrgIdNameByUserId.EndPoint{}):                              getOrgIdNameByUserId.Init(apiRoot),
			api.GetName(getOrgInfoByDealerOrgCode.EndPoint{}):                         getOrgInfoByDealerOrgCode.Init(apiRoot),
			api.GetName(getOrgListByUserId.EndPoint{}):                                getOrgListByUserId.Init(apiRoot),
			api.GetName(getOrgListForUser.EndPoint{}):                                 getOrgListForUser.Init(apiRoot),
			api.GetName(getOssObjectStream.EndPoint{}):                                getOssObjectStream.Init(apiRoot),
			api.GetName(getOwnerFaultConfigList.EndPoint{}):                           getOwnerFaultConfigList.Init(apiRoot),
			api.GetName(getParamterSettingBase.EndPoint{}):                            getParamterSettingBase.Init(apiRoot),
			api.GetName(getPhotoInfo.EndPoint{}):                                      getPhotoInfo.Init(apiRoot),
			api.GetName(getPlanedOrNotPsList.EndPoint{}):                              getPlanedOrNotPsList.Init(apiRoot),
			api.GetName(getPlantReportPDFList.EndPoint{}):                             getPlantReportPDFList.Init(apiRoot),
			api.GetName(getPowerChargeSettingInfo.EndPoint{}):                         getPowerChargeSettingInfo.Init(apiRoot),
			api.GetName(getPowerFormulaFaultAnalyse.EndPoint{}):                       getPowerFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(getPowerRobotInfoByRobotSn.EndPoint{}):                        getPowerRobotInfoByRobotSn.Init(apiRoot),
			api.GetName(getPowerRobotSweepAttrByPsId.EndPoint{}):                      getPowerRobotSweepAttrByPsId.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategy.EndPoint{}):                        getPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategyList.EndPoint{}):                    getPowerRobotSweepStrategyList.Init(apiRoot),
			api.GetName(getPowerSettingCharges.EndPoint{}):                            getPowerSettingCharges.Init(apiRoot),
			api.GetName(getPowerSettingHistoryRecords.EndPoint{}):                     getPowerSettingHistoryRecords.Init(apiRoot),
			api.GetName(getPowerStationBasicInfo.EndPoint{}):                          getPowerStationBasicInfo.Init(apiRoot),
			api.GetName(getPowerStationData.EndPoint{}):                               getPowerStationData.Init(apiRoot),
			api.GetName(getPowerStationForHousehold.EndPoint{}):                       getPowerStationForHousehold.Init(apiRoot),
			api.GetName(getPowerStationInfo.EndPoint{}):                               getPowerStationInfo.Init(apiRoot),
			api.GetName(getPowerStationPR.EndPoint{}):                                 getPowerStationPR.Init(apiRoot),
			api.GetName(getPrivateCloudValidityPeriod.EndPoint{}):                     getPrivateCloudValidityPeriod.Init(apiRoot),
			api.GetName(getProvInfoListByNationCode.EndPoint{}):                       getProvInfoListByNationCode.Init(apiRoot),
			api.GetName(getPsDetailByUserTokens.EndPoint{}):                           getPsDetailByUserTokens.Init(apiRoot),
			api.GetName(getPsDetailForSinglePage.EndPoint{}):                          getPsDetailForSinglePage.Init(apiRoot),
			api.GetName(getPsInstallerByPsId.EndPoint{}):                              getPsInstallerByPsId.Init(apiRoot),
			api.GetName(getPsInstallerOrgInfoByPsId.EndPoint{}):                       getPsInstallerOrgInfoByPsId.Init(apiRoot),
			api.GetName(getPsListForPsDataByPsId.EndPoint{}):                          getPsListForPsDataByPsId.Init(apiRoot),
			api.GetName(getRechargeOrderDetail.EndPoint{}):                            getRechargeOrderDetail.Init(apiRoot),
			api.GetName(getRechargeOrderList.EndPoint{}):                              getRechargeOrderList.Init(apiRoot),
			api.GetName(getRegionalTree.EndPoint{}):                                   getRegionalTree.Init(apiRoot),
			api.GetName(getRemoteParamSettingList.EndPoint{}):                         getRemoteParamSettingList.Init(apiRoot),
			api.GetName(getRemoteUpgradeScheduleDetails.EndPoint{}):                   getRemoteUpgradeScheduleDetails.Init(apiRoot),
			api.GetName(getReportData.EndPoint{}):                                     getReportData.Init(apiRoot),
			api.GetName(getReportEmailConfigInfo.EndPoint{}):                          getReportEmailConfigInfo.Init(apiRoot),
			api.GetName(getReportExportColumns.EndPoint{}):                            getReportExportColumns.Init(apiRoot),
			api.GetName(getReportListByUserId.EndPoint{}):                             getReportListByUserId.Init(apiRoot),
			api.GetName(getRobotDynamicCleaningView.EndPoint{}):                       getRobotDynamicCleaningView.Init(apiRoot),
			api.GetName(getRobotNumAndSweepCapacity.EndPoint{}):                       getRobotNumAndSweepCapacity.Init(apiRoot),
			api.GetName(getRuleUnit.EndPoint{}):                                       getRuleUnit.Init(apiRoot),
			api.GetName(getSendReportConfigCron.EndPoint{}):                           getSendReportConfigCron.Init(apiRoot),
			api.GetName(getShieldMapConditionList.EndPoint{}):                         getShieldMapConditionList.Init(apiRoot),
			api.GetName(getSimIdBySnList.EndPoint{}):                                  getSimIdBySnList.Init(apiRoot),
			api.GetName(getSingleIVData.EndPoint{}):                                   getSingleIVData.Init(apiRoot),
			api.GetName(getSnChangeRecord.EndPoint{}):                                 getSnChangeRecord.Init(apiRoot),
			api.GetName(getSnConnectionInfo.EndPoint{}):                               getSnConnectionInfo.Init(apiRoot),
			api.GetName(getSungwsConfigCache.EndPoint{}):                              getSungwsConfigCache.Init(apiRoot),
			api.GetName(getSungwsGlobalConfigCache.EndPoint{}):                        getSungwsGlobalConfigCache.Init(apiRoot),
			api.GetName(getSweepRobotDevList.EndPoint{}):                              getSweepRobotDevList.Init(apiRoot),
			api.GetName(getSysMsg.EndPoint{}):                                         getSysMsg.Init(apiRoot),
			api.GetName(getSysOrgNewList.EndPoint{}):                                  getSysOrgNewList.Init(apiRoot),
			api.GetName(getSysOrgNewOne.EndPoint{}):                                   getSysOrgNewOne.Init(apiRoot),
			api.GetName(getSysUserById.EndPoint{}):                                    getSysUserById.Init(apiRoot),
			api.GetName(getUUIDByUpuuid.EndPoint{}):                                   getUUIDByUpuuid.Init(apiRoot),
			api.GetName(getUserById.EndPoint{}):                                       getUserById.Init(apiRoot),
			api.GetName(getUserByInstaller.EndPoint{}):                                getUserByInstaller.Init(apiRoot),
			api.GetName(getUserDevOnlineOffineCount.EndPoint{}):                       getUserDevOnlineOffineCount.Init(apiRoot),
			api.GetName(getUserGDPRAttrs.EndPoint{}):                                  getUserGDPRAttrs.Init(apiRoot),
			api.GetName(getUserHavePowerStationCount.EndPoint{}):                      getUserHavePowerStationCount.Init(apiRoot),
			api.GetName(getUserInfoByUserAccounts.EndPoint{}):                         getUserInfoByUserAccounts.Init(apiRoot),
			api.GetName(getValidateCode.EndPoint{}):                                   getValidateCode.Init(apiRoot),
			api.GetName(getValidateCodeAtRegister.EndPoint{}):                         getValidateCodeAtRegister.Init(apiRoot),
			api.GetName(getWeatherInfo.EndPoint{}):                                    getWeatherInfo.Init(apiRoot),
			api.GetName(getWechatPushConfig.EndPoint{}):                               getWechatPushConfig.Init(apiRoot),
			api.GetName(getWorkInfo.EndPoint{}):                                       getWorkInfo.Init(apiRoot),
			api.GetName(groupStringCheck.EndPoint{}):                                  groupStringCheck.Init(apiRoot),
			api.GetName(handleDevByCommunicationSN.EndPoint{}):                        handleDevByCommunicationSN.Init(apiRoot),
			api.GetName(householdResetPassBySN.EndPoint{}):                            householdResetPassBySN.Init(apiRoot),
			api.GetName(immediatePayment.EndPoint{}):                                  immediatePayment.Init(apiRoot),
			api.GetName(importExcelData.EndPoint{}):                                   importExcelData.Init(apiRoot),
			api.GetName(incomeStatistics.EndPoint{}):                                  incomeStatistics.Init(apiRoot),
			api.GetName(informPush.EndPoint{}):                                        informPush.Init(apiRoot),
			api.GetName(insertEmgOrgInfo.EndPoint{}):                                  insertEmgOrgInfo.Init(apiRoot),
			api.GetName(intoDataToHbase.EndPoint{}):                                   intoDataToHbase.Init(apiRoot),
			api.GetName(ipLocationQuery.EndPoint{}):                                   ipLocationQuery.Init(apiRoot),
			api.GetName(isHave2GSn.EndPoint{}):                                        isHave2GSn.Init(apiRoot),
			api.GetName(judgeIsSettingMan.EndPoint{}):                                 judgeIsSettingMan.Init(apiRoot),
			api.GetName(listOssFiles.EndPoint{}):                                      listOssFiles.Init(apiRoot),
			api.GetName(loadAreaInfo.EndPoint{}):                                      loadAreaInfo.Init(apiRoot),
			api.GetName(loadPowerStation.EndPoint{}):                                  loadPowerStation.Init(apiRoot),
			api.GetName(loginByToken.EndPoint{}):                                      loginByToken.Init(apiRoot),
			api.GetName(logout.EndPoint{}):                                            logout.Init(apiRoot),
			api.GetName(mobilePhoneHasBound.EndPoint{}):                               mobilePhoneHasBound.Init(apiRoot),
			api.GetName(modifyEmgOrgStruc.EndPoint{}):                                 modifyEmgOrgStruc.Init(apiRoot),
			api.GetName(modifyFaultPlan.EndPoint{}):                                   modifyFaultPlan.Init(apiRoot),
			api.GetName(modifyOnDutyInfo.EndPoint{}):                                  modifyOnDutyInfo.Init(apiRoot),
			api.GetName(modifyPassword.EndPoint{}):                                    modifyPassword.Init(apiRoot),
			api.GetName(modifyPersonalUnitList.EndPoint{}):                            modifyPersonalUnitList.Init(apiRoot),
			api.GetName(modifyPsUser.EndPoint{}):                                      modifyPsUser.Init(apiRoot),
			api.GetName(moduleLogParamSet.EndPoint{}):                                 moduleLogParamSet.Init(apiRoot),
			api.GetName(operateOssFile.EndPoint{}):                                    operateOssFile.Init(apiRoot),
			api.GetName(operationPowerRobotSweepStrategy.EndPoint{}):                  operationPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(orgPowerReport.EndPoint{}):                                    orgPowerReport.Init(apiRoot),
			api.GetName(paramSetTryAgain.EndPoint{}):                                  paramSetTryAgain.Init(apiRoot),
			api.GetName(paramSetting.EndPoint{}):                                      paramSetting.Init(apiRoot),
			api.GetName(planPower.EndPoint{}):                                         planPower.Init(apiRoot),
			api.GetName(psForcastInfo.EndPoint{}):                                     psForcastInfo.Init(apiRoot),
			api.GetName(queryBatchSpeedyAddPowerStationResult.EndPoint{}):             queryBatchSpeedyAddPowerStationResult.Init(apiRoot),
			api.GetName(queryCardStatusCTCC.EndPoint{}):                               queryCardStatusCTCC.Init(apiRoot),
			api.GetName(queryChildAccountList.EndPoint{}):                             queryChildAccountList.Init(apiRoot),
			api.GetName(queryCompensationRecordData.EndPoint{}):                       queryCompensationRecordData.Init(apiRoot),
			api.GetName(queryCompensationRecordList.EndPoint{}):                       queryCompensationRecordList.Init(apiRoot),
			api.GetName(queryComponent.EndPoint{}):                                    queryComponent.Init(apiRoot),
			api.GetName(queryComponentTechnicalParam.EndPoint{}):                      queryComponentTechnicalParam.Init(apiRoot),
			api.GetName(queryCountryGridAndRelation.EndPoint{}):                       queryCountryGridAndRelation.Init(apiRoot),
			api.GetName(queryCountryList.EndPoint{}):                                  queryCountryList.Init(apiRoot),
			api.GetName(queryEnvironmentList.EndPoint{}):                              queryEnvironmentList.Init(apiRoot),
			api.GetName(queryFaultList.EndPoint{}):                                    queryFaultList.Init(apiRoot),
			api.GetName(queryFaultPlanDetail.EndPoint{}):                              queryFaultPlanDetail.Init(apiRoot),
			api.GetName(queryFaultRepairSteps.EndPoint{}):                             queryFaultRepairSteps.Init(apiRoot),
			api.GetName(queryFaultTypeAndLevelByCode.EndPoint{}):                      queryFaultTypeAndLevelByCode.Init(apiRoot),
			api.GetName(queryFirmwareFilesPage.EndPoint{}):                            queryFirmwareFilesPage.Init(apiRoot),
			api.GetName(queryInfotoAlert.EndPoint{}):                                  queryInfotoAlert.Init(apiRoot),
			api.GetName(queryInverterModelList.EndPoint{}):                            queryInverterModelList.Init(apiRoot),
			api.GetName(queryInverterVersionList.EndPoint{}):                          queryInverterVersionList.Init(apiRoot),
			api.GetName(queryM2MCardInfoCMCC.EndPoint{}):                              queryM2MCardInfoCMCC.Init(apiRoot),
			api.GetName(queryM2MCardTermInfoCMCC.EndPoint{}):                          queryM2MCardTermInfoCMCC.Init(apiRoot),
			api.GetName(queryModelInfoByModelId.EndPoint{}):                           queryModelInfoByModelId.Init(apiRoot),
			api.GetName(queryNoticeList.EndPoint{}):                                   queryNoticeList.Init(apiRoot),
			api.GetName(queryNumberOfRenewalReminders.EndPoint{}):                     queryNumberOfRenewalReminders.Init(apiRoot),
			api.GetName(queryOperRules.EndPoint{}):                                    queryOperRules.Init(apiRoot),
			api.GetName(queryOrderList.EndPoint{}):                                    queryOrderList.Init(apiRoot),
			api.GetName(queryOrderStep.EndPoint{}):                                    queryOrderStep.Init(apiRoot),
			api.GetName(queryOrgGenerationReport.EndPoint{}):                          queryOrgGenerationReport.Init(apiRoot),
			api.GetName(queryOrgInfoList.EndPoint{}):                                  queryOrgInfoList.Init(apiRoot),
			api.GetName(queryOrgPowerElecPercent.EndPoint{}):                          queryOrgPowerElecPercent.Init(apiRoot),
			api.GetName(queryOrgPsCompensationRecordList.EndPoint{}):                  queryOrgPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPersonalUnitList.EndPoint{}):                             queryPersonalUnitList.Init(apiRoot),
			api.GetName(queryPointDataTopOne.EndPoint{}):                              queryPointDataTopOne.Init(apiRoot),
			api.GetName(queryPowerStationInfo.EndPoint{}):                             queryPowerStationInfo.Init(apiRoot),
			api.GetName(queryPsAreaByUserIdAndAreaCode.EndPoint{}):                    queryPsAreaByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsCompensationRecordList.EndPoint{}):                     queryPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPsDataByDate.EndPoint{}):                                 queryPsDataByDate.Init(apiRoot),
			api.GetName(queryPsIdList.EndPoint{}):                                     queryPsIdList.Init(apiRoot),
			api.GetName(queryPsListByUserIdAndAreaCode.EndPoint{}):                    queryPsListByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsNameByPsId.EndPoint{}):                                 queryPsNameByPsId.Init(apiRoot),
			api.GetName(queryPsPrByDate.EndPoint{}):                                   queryPsPrByDate.Init(apiRoot),
			api.GetName(queryPsProfit.EndPoint{}):                                     queryPsProfit.Init(apiRoot),
			api.GetName(queryPsReportComparativeAnalysisOfPowerGeneration.EndPoint{}): queryPsReportComparativeAnalysisOfPowerGeneration.Init(apiRoot),
			api.GetName(queryPsStructureList.EndPoint{}):                              queryPsStructureList.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId.EndPoint{}):                       queryPuuidsByCommandTotalId.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId2.EndPoint{}):                      queryPuuidsByCommandTotalId2.Init(apiRoot),
			api.GetName(queryRepairRuleList.EndPoint{}):                               queryRepairRuleList.Init(apiRoot),
			api.GetName(queryReportListForManagementPage.EndPoint{}):                  queryReportListForManagementPage.Init(apiRoot),
			api.GetName(queryReportMsg.EndPoint{}):                                    queryReportMsg.Init(apiRoot),
			api.GetName(querySharingPs.EndPoint{}):                                    querySharingPs.Init(apiRoot),
			api.GetName(querySysAdvancedParam.EndPoint{}):                             querySysAdvancedParam.Init(apiRoot),
			api.GetName(queryTimeBySN.EndPoint{}):                                     queryTimeBySN.Init(apiRoot),
			api.GetName(queryTrafficByDateCTCC.EndPoint{}):                            queryTrafficByDateCTCC.Init(apiRoot),
			api.GetName(queryTrafficCTCC.EndPoint{}):                                  queryTrafficCTCC.Init(apiRoot),
			api.GetName(queryUnitUuidBytotalId.EndPoint{}):                            queryUnitUuidBytotalId.Init(apiRoot),
			api.GetName(queryUserBtnPri.EndPoint{}):                                   queryUserBtnPri.Init(apiRoot),
			api.GetName(queryUserByUserIds.EndPoint{}):                                queryUserByUserIds.Init(apiRoot),
			api.GetName(queryUserExtensionAttribute.EndPoint{}):                       queryUserExtensionAttribute.Init(apiRoot),
			api.GetName(queryUserForStep.EndPoint{}):                                  queryUserForStep.Init(apiRoot),
			api.GetName(queryUserProcessPri.EndPoint{}):                               queryUserProcessPri.Init(apiRoot),
			api.GetName(queryUserWechatBindRel.EndPoint{}):                            queryUserWechatBindRel.Init(apiRoot),
			api.GetName(queryUuidByTotalIdAndUuid.EndPoint{}):                         queryUuidByTotalIdAndUuid.Init(apiRoot),
			api.GetName(rechargeOrderSetMeal.EndPoint{}):                              rechargeOrderSetMeal.Init(apiRoot),
			api.GetName(renewSendReportConfirmEmail.EndPoint{}):                       renewSendReportConfirmEmail.Init(apiRoot),
			api.GetName(saveCustomerEmployee.EndPoint{}):                              saveCustomerEmployee.Init(apiRoot),
			api.GetName(saveDevSimList.EndPoint{}):                                    saveDevSimList.Init(apiRoot),
			api.GetName(saveEnviromentIncomeInfos.EndPoint{}):                         saveEnviromentIncomeInfos.Init(apiRoot),
			api.GetName(saveEnvironmentCurve.EndPoint{}):                              saveEnvironmentCurve.Init(apiRoot),
			api.GetName(saveFirmwareFile.EndPoint{}):                                  saveFirmwareFile.Init(apiRoot),
			api.GetName(saveIncomeSettingInfos.EndPoint{}):                            saveIncomeSettingInfos.Init(apiRoot),
			api.GetName(saveOrUpdateGroupStringCheckRule.EndPoint{}):                  saveOrUpdateGroupStringCheckRule.Init(apiRoot),
			api.GetName(saveParamModel.EndPoint{}):                                    saveParamModel.Init(apiRoot),
			api.GetName(savePowerCharges.EndPoint{}):                                  savePowerCharges.Init(apiRoot),
			api.GetName(savePowerRobotInfo.EndPoint{}):                                savePowerRobotInfo.Init(apiRoot),
			api.GetName(savePowerRobotSweepAttr.EndPoint{}):                           savePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(savePowerSettingCharges.EndPoint{}):                           savePowerSettingCharges.Init(apiRoot),
			api.GetName(savePowerSettingInfo.EndPoint{}):                              savePowerSettingInfo.Init(apiRoot),
			api.GetName(saveProductionBatchData.EndPoint{}):                           saveProductionBatchData.Init(apiRoot),
			api.GetName(saveRechargeOrderObj.EndPoint{}):                              saveRechargeOrderObj.Init(apiRoot),
			api.GetName(saveRechargeOrderOtherInfo.EndPoint{}):                        saveRechargeOrderOtherInfo.Init(apiRoot),
			api.GetName(saveRepair.EndPoint{}):                                        saveRepair.Init(apiRoot),
			api.GetName(saveReportExportColumns.EndPoint{}):                           saveReportExportColumns.Init(apiRoot),
			api.GetName(saveSetParam.EndPoint{}):                                      saveSetParam.Init(apiRoot),
			api.GetName(saveSysUserMsg.EndPoint{}):                                    saveSysUserMsg.Init(apiRoot),
			api.GetName(searchM2MMonthFlowCMCC.EndPoint{}):                            searchM2MMonthFlowCMCC.Init(apiRoot),
			api.GetName(selectSysTranslationNames.EndPoint{}):                         selectSysTranslationNames.Init(apiRoot),
			api.GetName(sendPsTimeZoneInstruction.EndPoint{}):                         sendPsTimeZoneInstruction.Init(apiRoot),
			api.GetName(setUpFormulaFaultAnalyse.EndPoint{}):                          setUpFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(setUserGDPRAttrs.EndPoint{}):                                  setUserGDPRAttrs.Init(apiRoot),
			api.GetName(settingNotice.EndPoint{}):                                     settingNotice.Init(apiRoot),
			api.GetName(shareMyPs.EndPoint{}):                                         shareMyPs.Init(apiRoot),
			api.GetName(sharePsBySN.EndPoint{}):                                       sharePsBySN.Init(apiRoot),
			api.GetName(showInverterByUnit.EndPoint{}):                                showInverterByUnit.Init(apiRoot),
			api.GetName(showOnlineUsers.EndPoint{}):                                   showOnlineUsers.Init(apiRoot),
			api.GetName(showWarning.EndPoint{}):                                       showWarning.Init(apiRoot),
			api.GetName(snIsExist.EndPoint{}):                                         snIsExist.Init(apiRoot),
			api.GetName(snsIsExist.EndPoint{}):                                        snsIsExist.Init(apiRoot),
			api.GetName(speedyAddPowerStation.EndPoint{}):                             speedyAddPowerStation.Init(apiRoot),
			api.GetName(stationUnitsList.EndPoint{}):                                  stationUnitsList.Init(apiRoot),
			api.GetName(stationsDiscreteData.EndPoint{}):                              stationsDiscreteData.Init(apiRoot),
			api.GetName(stationsIncomeList.EndPoint{}):                                stationsIncomeList.Init(apiRoot),
			api.GetName(stationsPointReport.EndPoint{}):                               stationsPointReport.Init(apiRoot),
			api.GetName(stationsYearPlanReport.EndPoint{}):                            stationsYearPlanReport.Init(apiRoot),
			api.GetName(sureAndImportSelettlementData.EndPoint{}):                     sureAndImportSelettlementData.Init(apiRoot),
			api.GetName(sweepDevParamSet.EndPoint{}):                                  sweepDevParamSet.Init(apiRoot),
			api.GetName(sweepDevRunControl.EndPoint{}):                                sweepDevRunControl.Init(apiRoot),
			api.GetName(sweepDevStrategyIssue.EndPoint{}):                             sweepDevStrategyIssue.Init(apiRoot),
			api.GetName(sysTimeZoneList.EndPoint{}):                                   sysTimeZoneList.Init(apiRoot),
			api.GetName(unLockUser.EndPoint{}):                                        unLockUser.Init(apiRoot),
			api.GetName(unlockChildAccount.EndPoint{}):                                unlockChildAccount.Init(apiRoot),
			api.GetName(updateCommunicationModuleState.EndPoint{}):                    updateCommunicationModuleState.Init(apiRoot),
			api.GetName(updateDevInstalledPower.EndPoint{}):                           updateDevInstalledPower.Init(apiRoot),
			api.GetName(updateFault.EndPoint{}):                                       updateFault.Init(apiRoot),
			api.GetName(updateFaultData.EndPoint{}):                                   updateFaultData.Init(apiRoot),
			api.GetName(updateFaultMsgByFaultCode.EndPoint{}):                         updateFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(updateFaultStatus.EndPoint{}):                                 updateFaultStatus.Init(apiRoot),
			api.GetName(updateHouseholdWorkOrder.EndPoint{}):                          updateHouseholdWorkOrder.Init(apiRoot),
			api.GetName(updateInverterSn2ModuleSn.EndPoint{}):                         updateInverterSn2ModuleSn.Init(apiRoot),
			api.GetName(updateOperateTicketAttachmentId.EndPoint{}):                   updateOperateTicketAttachmentId.Init(apiRoot),
			api.GetName(updateOwnerFaultConfig.EndPoint{}):                            updateOwnerFaultConfig.Init(apiRoot),
			api.GetName(updateParamSettingSysMsg.EndPoint{}):                          updateParamSettingSysMsg.Init(apiRoot),
			api.GetName(updatePlatformLevelFaultLevel.EndPoint{}):                     updatePlatformLevelFaultLevel.Init(apiRoot),
			api.GetName(updatePowerRobotInfo.EndPoint{}):                              updatePowerRobotInfo.Init(apiRoot),
			api.GetName(updatePowerRobotSweepAttr.EndPoint{}):                         updatePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(updatePowerStationForHousehold.EndPoint{}):                    updatePowerStationForHousehold.Init(apiRoot),
			api.GetName(updatePowerStationInfo.EndPoint{}):                            updatePowerStationInfo.Init(apiRoot),
			api.GetName(updatePowerUserInfo.EndPoint{}):                               updatePowerUserInfo.Init(apiRoot),
			api.GetName(updateReportConfigByEmailAddr.EndPoint{}):                     updateReportConfigByEmailAddr.Init(apiRoot),
			api.GetName(updateShareAttr.EndPoint{}):                                   updateShareAttr.Init(apiRoot),
			api.GetName(updateSnIsSureFlag.EndPoint{}):                                updateSnIsSureFlag.Init(apiRoot),
			api.GetName(updateStationPics.EndPoint{}):                                 updateStationPics.Init(apiRoot),
			api.GetName(updateSysAdvancedParam.EndPoint{}):                            updateSysAdvancedParam.Init(apiRoot),
			api.GetName(updateSysOrgNew.EndPoint{}):                                   updateSysOrgNew.Init(apiRoot),
			api.GetName(updateUinfoNetEaseUser.EndPoint{}):                            updateUinfoNetEaseUser.Init(apiRoot),
			api.GetName(updateUserExtensionAttribute.EndPoint{}):                      updateUserExtensionAttribute.Init(apiRoot),
			api.GetName(updateUserLanguage.EndPoint{}):                                updateUserLanguage.Init(apiRoot),
			api.GetName(updateUserPosition.EndPoint{}):                                updateUserPosition.Init(apiRoot),
			api.GetName(updateUserUpOrg.EndPoint{}):                                   updateUserUpOrg.Init(apiRoot),
			api.GetName(upgrade.EndPoint{}):                                           upgrade.Init(apiRoot),
			api.GetName(upgrate.EndPoint{}):                                           upgrate.Init(apiRoot),
			api.GetName(uploadFileToOss.EndPoint{}):                                   uploadFileToOss.Init(apiRoot),
			api.GetName(userAgreeGdprProtocol.EndPoint{}):                             userAgreeGdprProtocol.Init(apiRoot),
			api.GetName(userInfoUniqueCheck.EndPoint{}):                               userInfoUniqueCheck.Init(apiRoot),
			api.GetName(userMailHasBound.EndPoint{}):                                  userMailHasBound.Init(apiRoot),
			api.GetName(userRegister.EndPoint{}):                                      userRegister.Init(apiRoot),
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
