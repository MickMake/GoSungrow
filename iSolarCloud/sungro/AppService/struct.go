package AppService

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/sungro/AppService/acceptPsSharing"
	"GoSungrow/iSolarCloud/sungro/AppService/activateEmail"
	"GoSungrow/iSolarCloud/sungro/AppService/addConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/addDeviceRepair"
	"GoSungrow/iSolarCloud/sungro/AppService/addDeviceToStructureForHousehold"
	"GoSungrow/iSolarCloud/sungro/AppService/addDeviceToStructureForHouseholdByPsIdS"
	"GoSungrow/iSolarCloud/sungro/AppService/addFault"
	"GoSungrow/iSolarCloud/sungro/AppService/addFaultOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/addFaultPlan"
	"GoSungrow/iSolarCloud/sungro/AppService/addFaultRepairSteps"
	"GoSungrow/iSolarCloud/sungro/AppService/addHouseholdEvaluation"
	"GoSungrow/iSolarCloud/sungro/AppService/addHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/sungro/AppService/addHouseholdOpinionFeedback"
	"GoSungrow/iSolarCloud/sungro/AppService/addHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/addOnDutyInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/addOperRule"
	"GoSungrow/iSolarCloud/sungro/AppService/addOrDelPsStructure"
	"GoSungrow/iSolarCloud/sungro/AppService/addOrderStep"
	"GoSungrow/iSolarCloud/sungro/AppService/addPowerStationForHousehold"
	"GoSungrow/iSolarCloud/sungro/AppService/addPowerStationInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/addReportConfigEmail"
	"GoSungrow/iSolarCloud/sungro/AppService/addSysAdvancedParam"
	"GoSungrow/iSolarCloud/sungro/AppService/addSysOrgNew"
	"GoSungrow/iSolarCloud/sungro/AppService/aliPayAppTest"
	"GoSungrow/iSolarCloud/sungro/AppService/auditOperRule"
	"GoSungrow/iSolarCloud/sungro/AppService/batchAddStationBySn"
	"GoSungrow/iSolarCloud/sungro/AppService/batchImportSN"
	"GoSungrow/iSolarCloud/sungro/AppService/batchInsertUserAndOrg"
	"GoSungrow/iSolarCloud/sungro/AppService/batchModifyDevicesInfoAndPropertis"
	"GoSungrow/iSolarCloud/sungro/AppService/batchProcessPlantReport"
	"GoSungrow/iSolarCloud/sungro/AppService/batchUpdateDeviceSim"
	"GoSungrow/iSolarCloud/sungro/AppService/batchUpdateUserIsAgreeGdpr"
	"GoSungrow/iSolarCloud/sungro/AppService/boundMobilePhone"
	"GoSungrow/iSolarCloud/sungro/AppService/boundUserMail"
	"GoSungrow/iSolarCloud/sungro/AppService/caculateDeviceInputDiscrete"
	"GoSungrow/iSolarCloud/sungro/AppService/calculateDeviceDiscrete"
	"GoSungrow/iSolarCloud/sungro/AppService/calculateInitialCompensationData"
	"GoSungrow/iSolarCloud/sungro/AppService/cancelDeliverMail"
	"GoSungrow/iSolarCloud/sungro/AppService/cancelOrderScan"
	"GoSungrow/iSolarCloud/sungro/AppService/cancelParamSetTask"
	"GoSungrow/iSolarCloud/sungro/AppService/cancelPsSharing"
	"GoSungrow/iSolarCloud/sungro/AppService/cancelRechargeOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/changRechargeOrderToCancel"
	"GoSungrow/iSolarCloud/sungro/AppService/changeHouseholdUser2Installer"
	"GoSungrow/iSolarCloud/sungro/AppService/changeRemoteParam"
	"GoSungrow/iSolarCloud/sungro/AppService/checkDealerOrgCode"
	"GoSungrow/iSolarCloud/sungro/AppService/checkDevSnIsBelongsToUser"
	"GoSungrow/iSolarCloud/sungro/AppService/checkInverterResult"
	"GoSungrow/iSolarCloud/sungro/AppService/checkIsCanDoParamSet"
	"GoSungrow/iSolarCloud/sungro/AppService/checkIsIvScan"
	"GoSungrow/iSolarCloud/sungro/AppService/checkOssObjectExist"
	"GoSungrow/iSolarCloud/sungro/AppService/checkServiceIsConnect"
	"GoSungrow/iSolarCloud/sungro/AppService/checkTechnicalParameters"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUnitStatus"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUpRechargeDevicePaying"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserAccountUnique"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserAccountUniqueAll"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserInfoUnique"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserIsExist"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserListIsExist"
	"GoSungrow/iSolarCloud/sungro/AppService/checkUserPassword"
	"GoSungrow/iSolarCloud/sungro/AppService/cloudDeploymentRecord"
	"GoSungrow/iSolarCloud/sungro/AppService/comfirmParamModel"
	"GoSungrow/iSolarCloud/sungro/AppService/communicationModuleDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/compareValidateCode"
	"GoSungrow/iSolarCloud/sungro/AppService/componentInfo2Cloud"
	"GoSungrow/iSolarCloud/sungro/AppService/confirmFault"
	"GoSungrow/iSolarCloud/sungro/AppService/confirmIvFault"
	"GoSungrow/iSolarCloud/sungro/AppService/confirmReportConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/createAppkeyInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/createRenewInvoice"
	"GoSungrow/iSolarCloud/sungro/AppService/dealCommandReply"
	"GoSungrow/iSolarCloud/sungro/AppService/dealDeletePsFailPsDelete"
	"GoSungrow/iSolarCloud/sungro/AppService/dealFailRemoteUpgradeSubTasks"
	"GoSungrow/iSolarCloud/sungro/AppService/dealFailRemoteUpgradeTasks"
	"GoSungrow/iSolarCloud/sungro/AppService/dealFaultOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/dealGroupStringDisableOrEnable"
	"GoSungrow/iSolarCloud/sungro/AppService/dealNumberOfServiceCalls2Mysql"
	"GoSungrow/iSolarCloud/sungro/AppService/dealParamSettingAfterComplete"
	"GoSungrow/iSolarCloud/sungro/AppService/dealPsDataSupplement"
	"GoSungrow/iSolarCloud/sungro/AppService/dealPsReportEmailSend"
	"GoSungrow/iSolarCloud/sungro/AppService/dealRemoteUpgrade"
	"GoSungrow/iSolarCloud/sungro/AppService/dealSnElectrifyCheck"
	"GoSungrow/iSolarCloud/sungro/AppService/dealSysDeviceSimFlowInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/dealSysDeviceSimInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/definiteTimeDealSnExpRemind"
	"GoSungrow/iSolarCloud/sungro/AppService/definiteTimeDealSnStatus"
	"GoSungrow/iSolarCloud/sungro/AppService/delDeviceRepair"
	"GoSungrow/iSolarCloud/sungro/AppService/delOperRule"
	"GoSungrow/iSolarCloud/sungro/AppService/delayCallApiResidueTimes"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteComponent"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteCustomerEmployee"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteDeviceAccount"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteDeviceSimById"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteElectricitySettlementData"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteFaultPlan"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteFirmwareFiles"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteHouseholdEvaluation"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteInverterSnInChnnl"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteModuleLog"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteOnDutyInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteOperateBillFile"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteOssObject"
	"GoSungrow/iSolarCloud/sungro/AppService/deletePowerDevicePointById"
	"GoSungrow/iSolarCloud/sungro/AppService/deletePowerPicture"
	"GoSungrow/iSolarCloud/sungro/AppService/deletePowerRobotInfoBySnAndPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/deletePowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteProductionData"
	"GoSungrow/iSolarCloud/sungro/AppService/deletePs"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteRechargeOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteRegularlyConnectionInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteReportConfigEmailAddr"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteSysAdvancedParam"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteSysOrgNew"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteTemplate"
	"GoSungrow/iSolarCloud/sungro/AppService/deleteUserInfoAllByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/deviceInputDiscreteDeleteTime"
	"GoSungrow/iSolarCloud/sungro/AppService/deviceInputDiscreteGetTime"
	"GoSungrow/iSolarCloud/sungro/AppService/deviceInputDiscreteInsertTime"
	"GoSungrow/iSolarCloud/sungro/AppService/deviceInputDiscreteUpdateTime"
	"GoSungrow/iSolarCloud/sungro/AppService/devicePointsDataFromMySql"
	"GoSungrow/iSolarCloud/sungro/AppService/deviceReplace"
	"GoSungrow/iSolarCloud/sungro/AppService/editDeviceRepair"
	"GoSungrow/iSolarCloud/sungro/AppService/editOperRule"
	"GoSungrow/iSolarCloud/sungro/AppService/energyPovertyAlleviation"
	"GoSungrow/iSolarCloud/sungro/AppService/energyTrend"
	"GoSungrow/iSolarCloud/sungro/AppService/exportParamSettingValPDF"
	"GoSungrow/iSolarCloud/sungro/AppService/exportPlantReportPDF"
	"GoSungrow/iSolarCloud/sungro/AppService/faultAutoClose"
	"GoSungrow/iSolarCloud/sungro/AppService/faultCloseRemindOrderHandler"
	"GoSungrow/iSolarCloud/sungro/AppService/findCodeValueList"
	"GoSungrow/iSolarCloud/sungro/AppService/findEmgOrgInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/findEnvironmentInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/findFromHbaseAndRedis"
	"GoSungrow/iSolarCloud/sungro/AppService/findInfoByuuid"
	"GoSungrow/iSolarCloud/sungro/AppService/findLossAnalysisList"
	"GoSungrow/iSolarCloud/sungro/AppService/findOnDutyInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/findPsType"
	"GoSungrow/iSolarCloud/sungro/AppService/findSingleStationPR"
	"GoSungrow/iSolarCloud/sungro/AppService/findUserPassword"
	"GoSungrow/iSolarCloud/sungro/AppService/genTLSUserSigByUserAccount"
	"GoSungrow/iSolarCloud/sungro/AppService/generateRandomPassword"
	"GoSungrow/iSolarCloud/sungro/AppService/getAPIServiceInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getAccessedPermission"
	"GoSungrow/iSolarCloud/sungro/AppService/getAllDeviceByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getAllPowerDeviceSetName"
	"GoSungrow/iSolarCloud/sungro/AppService/getAllPowerRobotViewInfoByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getAllPsIdByOrgIds"
	"GoSungrow/iSolarCloud/sungro/AppService/getAllUserRemindCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getAndOutletsAndUnit"
	"GoSungrow/iSolarCloud/sungro/AppService/getApiCallsForAppkeys"
	"GoSungrow/iSolarCloud/sungro/AppService/getAreaInfoCodeByCounty"
	"GoSungrow/iSolarCloud/sungro/AppService/getAreaList"
	"GoSungrow/iSolarCloud/sungro/AppService/getAutoCreatePowerStation"
	"GoSungrow/iSolarCloud/sungro/AppService/getBackReadValue"
	"GoSungrow/iSolarCloud/sungro/AppService/getBatchNewestPointData"
	"GoSungrow/iSolarCloud/sungro/AppService/getCallApiResidueTimes"
	"GoSungrow/iSolarCloud/sungro/AppService/getChangedPsListByTime"
	"GoSungrow/iSolarCloud/sungro/AppService/getChnnlListByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getCloudList"
	"GoSungrow/iSolarCloud/sungro/AppService/getCloudServiceMappingConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/getCommunicationDeviceConfigInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getCommunicationModuleMonitorData"
	"GoSungrow/iSolarCloud/sungro/AppService/getComponentModelFactory"
	"GoSungrow/iSolarCloud/sungro/AppService/getConfigList"
	"GoSungrow/iSolarCloud/sungro/AppService/getConnectionInfoBySnAndLocalPort"
	"GoSungrow/iSolarCloud/sungro/AppService/getCountDown"
	"GoSungrow/iSolarCloud/sungro/AppService/getCountryServiceInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getCounty"
	"GoSungrow/iSolarCloud/sungro/AppService/getCustomerEmployee"
	"GoSungrow/iSolarCloud/sungro/AppService/getCustomerList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDataFromHBase"
	"GoSungrow/iSolarCloud/sungro/AppService/getDataFromHbaseByRowKey"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevInstalledPowerByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevRecord"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevRunRecordList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevSimByList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevSimList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceAccountById"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceFaultStatisticsData"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceModelInfoList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevicePointMinuteDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevicePoints"
	"GoSungrow/iSolarCloud/sungro/AppService/getDevicePropertys"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceRepairDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceTechBranchCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceTypeInfoList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDeviceTypeList"
	"GoSungrow/iSolarCloud/sungro/AppService/getDstInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getElectricitySettlementData"
	"GoSungrow/iSolarCloud/sungro/AppService/getElectricitySettlementDetailData"
	"GoSungrow/iSolarCloud/sungro/AppService/getEncryptPublicKey"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultMsgByFaultCode"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultMsgListByPageWithYYYYMM"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultMsgListWithYYYYMM"
	"GoSungrow/iSolarCloud/sungro/AppService/getFaultPlanList"
	"GoSungrow/iSolarCloud/sungro/AppService/getFileOperationRecordOne"
	"GoSungrow/iSolarCloud/sungro/AppService/getFormulaFaultAnalyseList"
	"GoSungrow/iSolarCloud/sungro/AppService/getGroupStringCheckResult"
	"GoSungrow/iSolarCloud/sungro/AppService/getGroupStringCheckRule"
	"GoSungrow/iSolarCloud/sungro/AppService/getHisData"
	"GoSungrow/iSolarCloud/sungro/AppService/getHistoryInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdEvaluation"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdLeaveMessage"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdOpinionFeedback"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdPsInstallerByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdUserInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdWorkOrderInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getHouseholdWorkOrderList"
	"GoSungrow/iSolarCloud/sungro/AppService/getI18nConfigByType"
	"GoSungrow/iSolarCloud/sungro/AppService/getI18nFileInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getI18nInfoByKey"
	"GoSungrow/iSolarCloud/sungro/AppService/getI18nVersion"
	"GoSungrow/iSolarCloud/sungro/AppService/getIncomeSettingInfos"
	"GoSungrow/iSolarCloud/sungro/AppService/getInfoFromAMap"
	"GoSungrow/iSolarCloud/sungro/AppService/getInfomationFromRedis"
	"GoSungrow/iSolarCloud/sungro/AppService/getInstallInfoList"
	"GoSungrow/iSolarCloud/sungro/AppService/getInstallerInfoByDealerOrgCodeOrId"
	"GoSungrow/iSolarCloud/sungro/AppService/getInvertDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/getInverterDataCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getInverterProcess"
	"GoSungrow/iSolarCloud/sungro/AppService/getInverterUuidBytotalId"
	"GoSungrow/iSolarCloud/sungro/AppService/getIvEchartsData"
	"GoSungrow/iSolarCloud/sungro/AppService/getIvEchartsDataById"
	"GoSungrow/iSolarCloud/sungro/AppService/getKpiInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getListMiFromHBase"
	"GoSungrow/iSolarCloud/sungro/AppService/getMapInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getMapMiFromHBase"
	"GoSungrow/iSolarCloud/sungro/AppService/getMenuAndPrivileges"
	"GoSungrow/iSolarCloud/sungro/AppService/getMicrogridEStoragePsReport"
	"GoSungrow/iSolarCloud/sungro/AppService/getModuleLogInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getModuleLogTaskList"
	"GoSungrow/iSolarCloud/sungro/AppService/getNationProvJSON"
	"GoSungrow/iSolarCloud/sungro/AppService/getNeedOpAsynOpRecordList"
	"GoSungrow/iSolarCloud/sungro/AppService/getNoticeInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getNumberOfServiceCalls"
	"GoSungrow/iSolarCloud/sungro/AppService/getOSSConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/getOperRuleDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getOperateBillFileId"
	"GoSungrow/iSolarCloud/sungro/AppService/getOperateTicketForDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrCreateNetEaseUserToken"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrderDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrderDataSql2"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrderDatas"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrderDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrderStatistics"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrgIdNameByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrgInfoByDealerOrgCode"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrgListByName"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrgListByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/getOrgListForUser"
	"GoSungrow/iSolarCloud/sungro/AppService/getOssObjectStream"
	"GoSungrow/iSolarCloud/sungro/AppService/getOwnerFaultConfigList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPListinfoFromMysql"
	"GoSungrow/iSolarCloud/sungro/AppService/getParamSetTemplate4NewProtocol"
	"GoSungrow/iSolarCloud/sungro/AppService/getParamSetTemplatePointInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getParamterSettingBase"
	"GoSungrow/iSolarCloud/sungro/AppService/getPhotoInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPlanedOrNotPsList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPlantReportPDFList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerChargeSettingInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDeviceModelTechList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDeviceModelTree"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDevicePointInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDeviceSetTaskDetailList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerDeviceSetTaskList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerFormulaFaultAnalyse"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerPictureList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerRobotInfoByRobotSn"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerRobotSweepAttrByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerRobotSweepStrategyList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerSettingCharges"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerSettingHistoryRecords"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationBasicInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationForHousehold"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationPR"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationTableDataSql"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStationTableDataSqlCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerStatistics"
	"GoSungrow/iSolarCloud/sungro/AppService/getPowerTrendDayData"
	"GoSungrow/iSolarCloud/sungro/AppService/getPrivateCloudValidityPeriod"
	"GoSungrow/iSolarCloud/sungro/AppService/getProvInfoListByNationCode"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsAuthKey"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsCurveInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsDataSupplementTaskList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsDetailByUserTokens"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsDetailForSinglePage"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsHealthState"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsInstallerByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsInstallerOrgInfoByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsList"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsListByName"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsListForPsDataByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsListStaticData"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsReport"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsUser"
	"GoSungrow/iSolarCloud/sungro/AppService/getPsWeatherList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRechargeOrderDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/getRechargeOrderItemDeviceList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRechargeOrderList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRegionalTree"
	"GoSungrow/iSolarCloud/sungro/AppService/getRemoteParamSettingList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRemoteUpgradeDeviceList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRemoteUpgradeScheduleDetails"
	"GoSungrow/iSolarCloud/sungro/AppService/getRemoteUpgradeSubTasksList"
	"GoSungrow/iSolarCloud/sungro/AppService/getRemoteUpgradeTaskList"
	"GoSungrow/iSolarCloud/sungro/AppService/getReportData"
	"GoSungrow/iSolarCloud/sungro/AppService/getReportEmailConfigInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getReportExportColumns"
	"GoSungrow/iSolarCloud/sungro/AppService/getReportListByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/getRobotDynamicCleaningView"
	"GoSungrow/iSolarCloud/sungro/AppService/getRobotNumAndSweepCapacity"
	"GoSungrow/iSolarCloud/sungro/AppService/getRuleUnit"
	"GoSungrow/iSolarCloud/sungro/AppService/getSendReportConfigCron"
	"GoSungrow/iSolarCloud/sungro/AppService/getSerialNum"
	"GoSungrow/iSolarCloud/sungro/AppService/getShieldMapConditionList"
	"GoSungrow/iSolarCloud/sungro/AppService/getSimIdBySnList"
	"GoSungrow/iSolarCloud/sungro/AppService/getSingleIVData"
	"GoSungrow/iSolarCloud/sungro/AppService/getSnChangeRecord"
	"GoSungrow/iSolarCloud/sungro/AppService/getSnConnectionInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getStationInfoSql"
	"GoSungrow/iSolarCloud/sungro/AppService/getSungwsConfigCache"
	"GoSungrow/iSolarCloud/sungro/AppService/getSungwsGlobalConfigCache"
	"GoSungrow/iSolarCloud/sungro/AppService/getSweepDevParamSetTemplate"
	"GoSungrow/iSolarCloud/sungro/AppService/getSweepRobotDevList"
	"GoSungrow/iSolarCloud/sungro/AppService/getSysMsg"
	"GoSungrow/iSolarCloud/sungro/AppService/getSysOrgNewList"
	"GoSungrow/iSolarCloud/sungro/AppService/getSysOrgNewOne"
	"GoSungrow/iSolarCloud/sungro/AppService/getSysUserById"
	"GoSungrow/iSolarCloud/sungro/AppService/getTableDataSql"
	"GoSungrow/iSolarCloud/sungro/AppService/getTableDataSqlCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getTemplateByInfoType"
	"GoSungrow/iSolarCloud/sungro/AppService/getTemplateList"
	"GoSungrow/iSolarCloud/sungro/AppService/getUUIDByUpuuid"
	"GoSungrow/iSolarCloud/sungro/AppService/getUpTimePoint"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserById"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserByInstaller"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserDevOnlineOffineCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserGDPRAttrs"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserHavePowerStationCount"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserInfoByUserAccounts"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserList"
	"GoSungrow/iSolarCloud/sungro/AppService/getUserPsOrderList"
	"GoSungrow/iSolarCloud/sungro/AppService/getValFromHBase"
	"GoSungrow/iSolarCloud/sungro/AppService/getValidateCode"
	"GoSungrow/iSolarCloud/sungro/AppService/getValidateCodeAtRegister"
	"GoSungrow/iSolarCloud/sungro/AppService/getWeatherInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/getWechatPushConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/getWorkInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/groupStringCheck"
	"GoSungrow/iSolarCloud/sungro/AppService/handleDevByCommunicationSN"
	"GoSungrow/iSolarCloud/sungro/AppService/householdResetPassBySN"
	"GoSungrow/iSolarCloud/sungro/AppService/immediatePayment"
	"GoSungrow/iSolarCloud/sungro/AppService/importExcelData"
	"GoSungrow/iSolarCloud/sungro/AppService/incomeStatistics"
	"GoSungrow/iSolarCloud/sungro/AppService/informPush"
	"GoSungrow/iSolarCloud/sungro/AppService/insertEmgOrgInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/insightSynDeviceStructure2Cloud"
	"GoSungrow/iSolarCloud/sungro/AppService/intoDataToHbase"
	"GoSungrow/iSolarCloud/sungro/AppService/ipLocationQuery"
	"GoSungrow/iSolarCloud/sungro/AppService/isHave2GSn"
	"GoSungrow/iSolarCloud/sungro/AppService/judgeDevIsHasInitSetTemplate"
	"GoSungrow/iSolarCloud/sungro/AppService/judgeIsSettingMan"
	"GoSungrow/iSolarCloud/sungro/AppService/listOssFiles"
	"GoSungrow/iSolarCloud/sungro/AppService/loadAreaInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/loadPowerStation"
	"GoSungrow/iSolarCloud/sungro/AppService/login"
	"GoSungrow/iSolarCloud/sungro/AppService/loginByToken"
	"GoSungrow/iSolarCloud/sungro/AppService/logout"
	"GoSungrow/iSolarCloud/sungro/AppService/mobilePhoneHasBound"
	"GoSungrow/iSolarCloud/sungro/AppService/modifiedDeviceInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyEmgOrgStruc"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyFaultPlan"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyOnDutyInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyPassword"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyPersonalUnitList"
	"GoSungrow/iSolarCloud/sungro/AppService/modifyPsUser"
	"GoSungrow/iSolarCloud/sungro/AppService/moduleLogParamSet"
	"GoSungrow/iSolarCloud/sungro/AppService/operateOssFile"
	"GoSungrow/iSolarCloud/sungro/AppService/operationPowerRobotSweepStrategy"
	"GoSungrow/iSolarCloud/sungro/AppService/orgPowerReport"
	"GoSungrow/iSolarCloud/sungro/AppService/paramSetTryAgain"
	"GoSungrow/iSolarCloud/sungro/AppService/paramSetting"
	"GoSungrow/iSolarCloud/sungro/AppService/planPower"
	"GoSungrow/iSolarCloud/sungro/AppService/powerDevicePointList"
	"GoSungrow/iSolarCloud/sungro/AppService/powerTrendChartData"
	"GoSungrow/iSolarCloud/sungro/AppService/psForcastInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/psHourPointsValue"
	"GoSungrow/iSolarCloud/sungro/AppService/queryAllPsIdAndName"
	"GoSungrow/iSolarCloud/sungro/AppService/queryBatchCreatePsTaskList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryBatchSpeedyAddPowerStationResult"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCardStatusCTCC"
	"GoSungrow/iSolarCloud/sungro/AppService/queryChildAccountList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCompensationRecordData"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCompensationRecordList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryComponent"
	"GoSungrow/iSolarCloud/sungro/AppService/queryComponentTechnicalParam"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCountryGridAndRelation"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCountryList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryCtrlTaskById"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceInfoForApp"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceListByUserId"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceListForApp"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceModelTechnical"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDevicePointDayMonthYearDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDevicePointMinuteDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDevicePointsDayMonthYearDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceRealTimeDataByPsKeys"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceRepairList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryDeviceTypeInfoList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryEnvironmentList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultPlanDetail"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultRepairSteps"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultTypeAndLevelByCode"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultTypeByDevice"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFaultTypeByDevicePage"
	"GoSungrow/iSolarCloud/sungro/AppService/queryFirmwareFilesPage"
	"GoSungrow/iSolarCloud/sungro/AppService/queryInfotoAlert"
	"GoSungrow/iSolarCloud/sungro/AppService/queryInverterModelList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryInverterVersionList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryM2MCardInfoCMCC"
	"GoSungrow/iSolarCloud/sungro/AppService/queryM2MCardTermInfoCMCC"
	"GoSungrow/iSolarCloud/sungro/AppService/queryModelInfoByModelId"
	"GoSungrow/iSolarCloud/sungro/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryNoticeList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryNumberOfRenewalReminders"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOperRules"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrderList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrderStep"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrgGenerationReport"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrgInfoList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrgPowerElecPercent"
	"GoSungrow/iSolarCloud/sungro/AppService/queryOrgPsCompensationRecordList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryParamSettingTask"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPersonalUnitList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPointDataTopOne"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPowerStationInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsAreaByUserIdAndAreaCode"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsCompensationRecordList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsDataByDate"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsIdList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsListByUserIdAndAreaCode"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsNameByPsId"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsPrByDate"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsProfit"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsReportComparativeAnalysisOfPowerGeneration"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPsStructureList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPuuidsByCommandTotalId"
	"GoSungrow/iSolarCloud/sungro/AppService/queryPuuidsByCommandTotalId2"
	"GoSungrow/iSolarCloud/sungro/AppService/queryRepairRuleList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryReportListForManagementPage"
	"GoSungrow/iSolarCloud/sungro/AppService/queryReportMsg"
	"GoSungrow/iSolarCloud/sungro/AppService/querySharingPs"
	"GoSungrow/iSolarCloud/sungro/AppService/querySysAdvancedParam"
	"GoSungrow/iSolarCloud/sungro/AppService/queryTimeBySN"
	"GoSungrow/iSolarCloud/sungro/AppService/queryTrafficByDateCTCC"
	"GoSungrow/iSolarCloud/sungro/AppService/queryTrafficCTCC"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUnitList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUnitUuidBytotalId"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserBtnPri"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserByUserIds"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserExtensionAttribute"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserForStep"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserList"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserProcessPri"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUserWechatBindRel"
	"GoSungrow/iSolarCloud/sungro/AppService/queryUuidByTotalIdAndUuid"
	"GoSungrow/iSolarCloud/sungro/AppService/rechargeOrderSetMeal"
	"GoSungrow/iSolarCloud/sungro/AppService/renewSendReportConfirmEmail"
	"GoSungrow/iSolarCloud/sungro/AppService/reportList"
	"GoSungrow/iSolarCloud/sungro/AppService/saveCustomerEmployee"
	"GoSungrow/iSolarCloud/sungro/AppService/saveDevSimList"
	"GoSungrow/iSolarCloud/sungro/AppService/saveDeviceAccountBatchData"
	"GoSungrow/iSolarCloud/sungro/AppService/saveEnviromentIncomeInfos"
	"GoSungrow/iSolarCloud/sungro/AppService/saveEnvironmentCurve"
	"GoSungrow/iSolarCloud/sungro/AppService/saveFirmwareFile"
	"GoSungrow/iSolarCloud/sungro/AppService/saveIncomeSettingInfos"
	"GoSungrow/iSolarCloud/sungro/AppService/saveOrUpdateGroupStringCheckRule"
	"GoSungrow/iSolarCloud/sungro/AppService/saveParamModel"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerCharges"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerDevicePoint"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerRobotInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerRobotSweepAttr"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerSettingCharges"
	"GoSungrow/iSolarCloud/sungro/AppService/savePowerSettingInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/saveProductionBatchData"
	"GoSungrow/iSolarCloud/sungro/AppService/saveRechargeOrderObj"
	"GoSungrow/iSolarCloud/sungro/AppService/saveRechargeOrderOtherInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/saveRepair"
	"GoSungrow/iSolarCloud/sungro/AppService/saveReportExportColumns"
	"GoSungrow/iSolarCloud/sungro/AppService/saveSetParam"
	"GoSungrow/iSolarCloud/sungro/AppService/saveSysUserMsg"
	"GoSungrow/iSolarCloud/sungro/AppService/saveTemplate"
	"GoSungrow/iSolarCloud/sungro/AppService/searchM2MMonthFlowCMCC"
	"GoSungrow/iSolarCloud/sungro/AppService/selectSysTranslationNames"
	"GoSungrow/iSolarCloud/sungro/AppService/sendPsTimeZoneInstruction"
	"GoSungrow/iSolarCloud/sungro/AppService/setUpFormulaFaultAnalyse"
	"GoSungrow/iSolarCloud/sungro/AppService/setUserGDPRAttrs"
	"GoSungrow/iSolarCloud/sungro/AppService/settingNotice"
	"GoSungrow/iSolarCloud/sungro/AppService/shareMyPs"
	"GoSungrow/iSolarCloud/sungro/AppService/sharePsBySN"
	"GoSungrow/iSolarCloud/sungro/AppService/showInverterByUnit"
	"GoSungrow/iSolarCloud/sungro/AppService/showOnlineUsers"
	"GoSungrow/iSolarCloud/sungro/AppService/showWarning"
	"GoSungrow/iSolarCloud/sungro/AppService/snIsExist"
	"GoSungrow/iSolarCloud/sungro/AppService/snsIsExist"
	"GoSungrow/iSolarCloud/sungro/AppService/speedyAddPowerStation"
	"GoSungrow/iSolarCloud/sungro/AppService/stationDeviceHistoryDataList"
	"GoSungrow/iSolarCloud/sungro/AppService/stationUnitsList"
	"GoSungrow/iSolarCloud/sungro/AppService/stationsDiscreteData"
	"GoSungrow/iSolarCloud/sungro/AppService/stationsIncomeList"
	"GoSungrow/iSolarCloud/sungro/AppService/stationsPointReport"
	"GoSungrow/iSolarCloud/sungro/AppService/stationsYearPlanReport"
	"GoSungrow/iSolarCloud/sungro/AppService/sureAndImportSelettlementData"
	"GoSungrow/iSolarCloud/sungro/AppService/sweepDevParamSet"
	"GoSungrow/iSolarCloud/sungro/AppService/sweepDevRunControl"
	"GoSungrow/iSolarCloud/sungro/AppService/sweepDevStrategyIssue"
	"GoSungrow/iSolarCloud/sungro/AppService/sysTimeZoneList"
	"GoSungrow/iSolarCloud/sungro/AppService/unLockUser"
	"GoSungrow/iSolarCloud/sungro/AppService/unlockChildAccount"
	"GoSungrow/iSolarCloud/sungro/AppService/updateCommunicationModuleState"
	"GoSungrow/iSolarCloud/sungro/AppService/updateDevInstalledPower"
	"GoSungrow/iSolarCloud/sungro/AppService/updateFault"
	"GoSungrow/iSolarCloud/sungro/AppService/updateFaultData"
	"GoSungrow/iSolarCloud/sungro/AppService/updateFaultMsgByFaultCode"
	"GoSungrow/iSolarCloud/sungro/AppService/updateFaultStatus"
	"GoSungrow/iSolarCloud/sungro/AppService/updateHouseholdWorkOrder"
	"GoSungrow/iSolarCloud/sungro/AppService/updateInverterSn2ModuleSn"
	"GoSungrow/iSolarCloud/sungro/AppService/updateOperateTicketAttachmentId"
	"GoSungrow/iSolarCloud/sungro/AppService/updateOrderDeviceByCustomerService"
	"GoSungrow/iSolarCloud/sungro/AppService/updateOwnerFaultConfig"
	"GoSungrow/iSolarCloud/sungro/AppService/updateParamSettingSysMsg"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePlatformLevelFaultLevel"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerDevicePoint"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerRobotInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerRobotSweepAttr"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerStationForHousehold"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerStationInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/updatePowerUserInfo"
	"GoSungrow/iSolarCloud/sungro/AppService/updateReportConfigByEmailAddr"
	"GoSungrow/iSolarCloud/sungro/AppService/updateShareAttr"
	"GoSungrow/iSolarCloud/sungro/AppService/updateSnIsSureFlag"
	"GoSungrow/iSolarCloud/sungro/AppService/updateStationPics"
	"GoSungrow/iSolarCloud/sungro/AppService/updateSysAdvancedParam"
	"GoSungrow/iSolarCloud/sungro/AppService/updateSysOrgNew"
	"GoSungrow/iSolarCloud/sungro/AppService/updateTemplate"
	"GoSungrow/iSolarCloud/sungro/AppService/updateUinfoNetEaseUser"
	"GoSungrow/iSolarCloud/sungro/AppService/updateUserExtensionAttribute"
	"GoSungrow/iSolarCloud/sungro/AppService/updateUserLanguage"
	"GoSungrow/iSolarCloud/sungro/AppService/updateUserPosition"
	"GoSungrow/iSolarCloud/sungro/AppService/updateUserUpOrg"
	"GoSungrow/iSolarCloud/sungro/AppService/upgrade"
	"GoSungrow/iSolarCloud/sungro/AppService/upgrate"
	"GoSungrow/iSolarCloud/sungro/AppService/uploadFileToOss"
	"GoSungrow/iSolarCloud/sungro/AppService/userAgreeGdprProtocol"
	"GoSungrow/iSolarCloud/sungro/AppService/userInfoUniqueCheck"
	"GoSungrow/iSolarCloud/sungro/AppService/userMailHasBound"
	"GoSungrow/iSolarCloud/sungro/AppService/userRegister"
	"fmt"
)

var _ api.Area = (*Area)(nil)

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
			api.GetName(getHistoryInfo.EndPoint{}):              getHistoryInfo.Init(apiRoot),

			api.GetName(psHourPointsValue.EndPoint{}):                                 psHourPointsValue.Init(apiRoot),
			api.GetName(getPsCurveInfo.EndPoint{}):                                    getPsCurveInfo.Init(apiRoot),
			api.GetName(queryMutiPointDataList.EndPoint{}):                            queryMutiPointDataList.Init(apiRoot),
			api.GetName(getTemplateList.EndPoint{}):                                   getTemplateList.Init(apiRoot),
			api.GetName(getDevicePoints.EndPoint{}):                                   getDevicePoints.Init(apiRoot),
			api.GetName(getAllPowerDeviceSetName.EndPoint{}):                          getAllPowerDeviceSetName.Init(apiRoot),
			api.GetName(queryAllPsIdAndName.EndPoint{}):                               queryAllPsIdAndName.Init(apiRoot),
			api.GetName(queryDevicePointMinuteDataList.EndPoint{}):                    queryDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(getTemplateByInfoType.EndPoint{}):                             getTemplateByInfoType.Init(apiRoot),
			api.GetName(energyTrend.EndPoint{}):                                       energyTrend.Init(apiRoot),
			api.GetName(queryBatchCreatePsTaskList.EndPoint{}):                        queryBatchCreatePsTaskList.Init(apiRoot),
			api.GetName(exportPlantReportPDF.EndPoint{}):                              exportPlantReportPDF.Init(apiRoot),
			api.GetName(getTableDataSql.EndPoint{}):                                   getTableDataSql.Init(apiRoot),
			api.GetName(getTableDataSqlCount.EndPoint{}):                              getTableDataSqlCount.Init(apiRoot),
			api.GetName(getPListinfoFromMysql.EndPoint{}):                             getPListinfoFromMysql.Init(apiRoot),
			api.GetName(getDataFromHBase.EndPoint{}):                                  getDataFromHBase.Init(apiRoot),
			api.GetName(getListMiFromHBase.EndPoint{}):                                getListMiFromHBase.Init(apiRoot),
			api.GetName(getMapMiFromHBase.EndPoint{}):                                 getMapMiFromHBase.Init(apiRoot),
			api.GetName(getValFromHBase.EndPoint{}):                                   getValFromHBase.Init(apiRoot),
			api.GetName(getPowerStationTableDataSql.EndPoint{}):                       getPowerStationTableDataSql.Init(apiRoot),
			api.GetName(getPowerStationTableDataSqlCount.EndPoint{}):                  getPowerStationTableDataSqlCount.Init(apiRoot),
			api.GetName(getDataFromHbaseByRowKey.EndPoint{}):                          getDataFromHbaseByRowKey.Init(apiRoot),
			api.GetName(findInfoByuuid.EndPoint{}):                                    findInfoByuuid.Init(apiRoot),
			api.GetName(getStationInfoSql.EndPoint{}):                                 getStationInfoSql.Init(apiRoot),
			api.GetName(getDevicePointMinuteDataList.EndPoint{}):                      getDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(powerDevicePointList.EndPoint{}):                              powerDevicePointList.Init(apiRoot),
			api.GetName(getPowerTrendDayData.EndPoint{}):                              getPowerTrendDayData.Init(apiRoot),
			api.GetName(powerTrendChartData.EndPoint{}):                               powerTrendChartData.Init(apiRoot),
			api.GetName(getUpTimePoint.EndPoint{}):                                    getUpTimePoint.Init(apiRoot),
			api.GetName(getSerialNum.EndPoint{}):                                      getSerialNum.Init(apiRoot),
			api.GetName(getModuleLogTaskList.EndPoint{}):                              getModuleLogTaskList.Init(apiRoot),
			api.GetName(queryParamSettingTask.EndPoint{}):                             queryParamSettingTask.Init(apiRoot),
			api.GetName(queryCtrlTaskById.EndPoint{}):                                 queryCtrlTaskById.Init(apiRoot),
			api.GetName(queryDeviceInfo.EndPoint{}):                                   queryDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceInfoForApp.EndPoint{}):                             queryDeviceInfoForApp.Init(apiRoot),
			api.GetName(queryDeviceList.EndPoint{}):                                   queryDeviceList.Init(apiRoot),
			api.GetName(queryDeviceListByUserId.EndPoint{}):                           queryDeviceListByUserId.Init(apiRoot),
			api.GetName(queryDeviceListForApp.EndPoint{}):                             queryDeviceListForApp.Init(apiRoot),
			api.GetName(queryDeviceModelTechnical.EndPoint{}):                         queryDeviceModelTechnical.Init(apiRoot),
			api.GetName(getPsDataSupplementTaskList.EndPoint{}):                       getPsDataSupplementTaskList.Init(apiRoot),
			api.GetName(getPsDetail.EndPoint{}):                                       getPsDetail.Init(apiRoot),
			api.GetName(getPsHealthState.EndPoint{}):                                  getPsHealthState.Init(apiRoot),
			api.GetName(getPowerDeviceSetTaskDetailList.EndPoint{}):                   getPowerDeviceSetTaskDetailList.Init(apiRoot),
			api.GetName(getPowerDeviceSetTaskList.EndPoint{}):                         getPowerDeviceSetTaskList.Init(apiRoot),
			api.GetName(getPsAuthKey.EndPoint{}):                                      getPsAuthKey.Init(apiRoot),
			api.GetName(getApiCallsForAppkeys.EndPoint{}):                             getApiCallsForAppkeys.Init(apiRoot),
			api.GetName(exportParamSettingValPDF.EndPoint{}):                          exportParamSettingValPDF.Init(apiRoot),
			api.GetName(checkUnitStatus.EndPoint{}):                                   checkUnitStatus.Init(apiRoot),
			api.GetName(devicePointsDataFromMySql.EndPoint{}):                         devicePointsDataFromMySql.Init(apiRoot),
			api.GetName(queryDevicePointDayMonthYearDataList.EndPoint{}):              queryDevicePointDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDevicePointsDayMonthYearDataList.EndPoint{}):             queryDevicePointsDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDeviceRealTimeDataByPsKeys.EndPoint{}):                   queryDeviceRealTimeDataByPsKeys.Init(apiRoot),
			api.GetName(acceptPsSharing.EndPoint{}):                                   acceptPsSharing.Init(apiRoot),
			api.GetName(activateEmail.EndPoint{}):                                     activateEmail.Init(apiRoot),
			api.GetName(addConfig.EndPoint{}):                                         addConfig.Init(apiRoot),
			api.GetName(addDeviceRepair.EndPoint{}):                                   addDeviceRepair.Init(apiRoot),
			api.GetName(addDeviceToStructureForHousehold.EndPoint{}):                  addDeviceToStructureForHousehold.Init(apiRoot),
			api.GetName(addDeviceToStructureForHouseholdByPsIdS.EndPoint{}):           addDeviceToStructureForHouseholdByPsIdS.Init(apiRoot),
			api.GetName(addFault.EndPoint{}):                                          addFault.Init(apiRoot),
			api.GetName(addFaultOrder.EndPoint{}):                                     addFaultOrder.Init(apiRoot),
			api.GetName(addFaultPlan.EndPoint{}):                                      addFaultPlan.Init(apiRoot),
			api.GetName(addFaultRepairSteps.EndPoint{}):                               addFaultRepairSteps.Init(apiRoot),
			api.GetName(addHouseholdEvaluation.EndPoint{}):                            addHouseholdEvaluation.Init(apiRoot),
			api.GetName(addHouseholdLeaveMessage.EndPoint{}):                          addHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(addHouseholdOpinionFeedback.EndPoint{}):                       addHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(addHouseholdWorkOrder.EndPoint{}):                             addHouseholdWorkOrder.Init(apiRoot),
			api.GetName(addOnDutyInfo.EndPoint{}):                                     addOnDutyInfo.Init(apiRoot),
			api.GetName(addOperRule.EndPoint{}):                                       addOperRule.Init(apiRoot),
			api.GetName(addOrDelPsStructure.EndPoint{}):                               addOrDelPsStructure.Init(apiRoot),
			api.GetName(addOrderStep.EndPoint{}):                                      addOrderStep.Init(apiRoot),
			api.GetName(addPowerStationForHousehold.EndPoint{}):                       addPowerStationForHousehold.Init(apiRoot),
			api.GetName(addPowerStationInfo.EndPoint{}):                               addPowerStationInfo.Init(apiRoot),
			api.GetName(addReportConfigEmail.EndPoint{}):                              addReportConfigEmail.Init(apiRoot),
			api.GetName(addSysAdvancedParam.EndPoint{}):                               addSysAdvancedParam.Init(apiRoot),
			api.GetName(addSysOrgNew.EndPoint{}):                                      addSysOrgNew.Init(apiRoot),
			api.GetName(aliPayAppTest.EndPoint{}):                                     aliPayAppTest.Init(apiRoot),
			api.GetName(auditOperRule.EndPoint{}):                                     auditOperRule.Init(apiRoot),
			api.GetName(batchAddStationBySn.EndPoint{}):                               batchAddStationBySn.Init(apiRoot),
			api.GetName(batchImportSN.EndPoint{}):                                     batchImportSN.Init(apiRoot),
			api.GetName(batchInsertUserAndOrg.EndPoint{}):                             batchInsertUserAndOrg.Init(apiRoot),
			api.GetName(batchModifyDevicesInfoAndPropertis.EndPoint{}):                batchModifyDevicesInfoAndPropertis.Init(apiRoot),
			api.GetName(batchProcessPlantReport.EndPoint{}):                           batchProcessPlantReport.Init(apiRoot),
			api.GetName(batchUpdateDeviceSim.EndPoint{}):                              batchUpdateDeviceSim.Init(apiRoot),
			api.GetName(batchUpdateUserIsAgreeGdpr.EndPoint{}):                        batchUpdateUserIsAgreeGdpr.Init(apiRoot),
			api.GetName(boundMobilePhone.EndPoint{}):                                  boundMobilePhone.Init(apiRoot),
			api.GetName(boundUserMail.EndPoint{}):                                     boundUserMail.Init(apiRoot),
			api.GetName(caculateDeviceInputDiscrete.EndPoint{}):                       caculateDeviceInputDiscrete.Init(apiRoot),
			api.GetName(calculateDeviceDiscrete.EndPoint{}):                           calculateDeviceDiscrete.Init(apiRoot),
			api.GetName(calculateInitialCompensationData.EndPoint{}):                  calculateInitialCompensationData.Init(apiRoot),
			api.GetName(cancelDeliverMail.EndPoint{}):                                 cancelDeliverMail.Init(apiRoot),
			api.GetName(cancelOrderScan.EndPoint{}):                                   cancelOrderScan.Init(apiRoot),
			api.GetName(cancelParamSetTask.EndPoint{}):                                cancelParamSetTask.Init(apiRoot),
			api.GetName(cancelPsSharing.EndPoint{}):                                   cancelPsSharing.Init(apiRoot),
			api.GetName(cancelRechargeOrder.EndPoint{}):                               cancelRechargeOrder.Init(apiRoot),
			api.GetName(changRechargeOrderToCancel.EndPoint{}):                        changRechargeOrderToCancel.Init(apiRoot),
			api.GetName(changeHouseholdUser2Installer.EndPoint{}):                     changeHouseholdUser2Installer.Init(apiRoot),
			api.GetName(changeRemoteParam.EndPoint{}):                                 changeRemoteParam.Init(apiRoot),
			api.GetName(checkDealerOrgCode.EndPoint{}):                                checkDealerOrgCode.Init(apiRoot),
			api.GetName(checkDevSnIsBelongsToUser.EndPoint{}):                         checkDevSnIsBelongsToUser.Init(apiRoot),
			api.GetName(checkInverterResult.EndPoint{}):                               checkInverterResult.Init(apiRoot),
			api.GetName(checkIsCanDoParamSet.EndPoint{}):                              checkIsCanDoParamSet.Init(apiRoot),
			api.GetName(checkIsIvScan.EndPoint{}):                                     checkIsIvScan.Init(apiRoot),
			api.GetName(checkOssObjectExist.EndPoint{}):                               checkOssObjectExist.Init(apiRoot),
			api.GetName(checkServiceIsConnect.EndPoint{}):                             checkServiceIsConnect.Init(apiRoot),
			api.GetName(checkTechnicalParameters.EndPoint{}):                          checkTechnicalParameters.Init(apiRoot),
			api.GetName(checkUpRechargeDevicePaying.EndPoint{}):                       checkUpRechargeDevicePaying.Init(apiRoot),
			api.GetName(checkUserAccountUnique.EndPoint{}):                            checkUserAccountUnique.Init(apiRoot),
			api.GetName(checkUserAccountUniqueAll.EndPoint{}):                         checkUserAccountUniqueAll.Init(apiRoot),
			api.GetName(checkUserInfoUnique.EndPoint{}):                               checkUserInfoUnique.Init(apiRoot),
			api.GetName(checkUserIsExist.EndPoint{}):                                  checkUserIsExist.Init(apiRoot),
			api.GetName(checkUserListIsExist.EndPoint{}):                              checkUserListIsExist.Init(apiRoot),
			api.GetName(checkUserPassword.EndPoint{}):                                 checkUserPassword.Init(apiRoot),
			api.GetName(cloudDeploymentRecord.EndPoint{}):                             cloudDeploymentRecord.Init(apiRoot),
			api.GetName(comfirmParamModel.EndPoint{}):                                 comfirmParamModel.Init(apiRoot),
			api.GetName(compareValidateCode.EndPoint{}):                               compareValidateCode.Init(apiRoot),
			api.GetName(componentInfo2Cloud.EndPoint{}):                               componentInfo2Cloud.Init(apiRoot),
			api.GetName(confirmFault.EndPoint{}):                                      confirmFault.Init(apiRoot),
			api.GetName(confirmIvFault.EndPoint{}):                                    confirmIvFault.Init(apiRoot),
			api.GetName(confirmReportConfig.EndPoint{}):                               confirmReportConfig.Init(apiRoot),
			api.GetName(createAppkeyInfo.EndPoint{}):                                  createAppkeyInfo.Init(apiRoot),
			api.GetName(createRenewInvoice.EndPoint{}):                                createRenewInvoice.Init(apiRoot),
			api.GetName(dealCommandReply.EndPoint{}):                                  dealCommandReply.Init(apiRoot),
			api.GetName(dealDeletePsFailPsDelete.EndPoint{}):                          dealDeletePsFailPsDelete.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeSubTasks.EndPoint{}):                     dealFailRemoteUpgradeSubTasks.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeTasks.EndPoint{}):                        dealFailRemoteUpgradeTasks.Init(apiRoot),
			api.GetName(dealFaultOrder.EndPoint{}):                                    dealFaultOrder.Init(apiRoot),
			api.GetName(dealGroupStringDisableOrEnable.EndPoint{}):                    dealGroupStringDisableOrEnable.Init(apiRoot),
			api.GetName(dealNumberOfServiceCalls2Mysql.EndPoint{}):                    dealNumberOfServiceCalls2Mysql.Init(apiRoot),
			api.GetName(dealParamSettingAfterComplete.EndPoint{}):                     dealParamSettingAfterComplete.Init(apiRoot),
			api.GetName(dealPsDataSupplement.EndPoint{}):                              dealPsDataSupplement.Init(apiRoot),
			api.GetName(dealPsReportEmailSend.EndPoint{}):                             dealPsReportEmailSend.Init(apiRoot),
			api.GetName(dealRemoteUpgrade.EndPoint{}):                                 dealRemoteUpgrade.Init(apiRoot),
			api.GetName(dealSnElectrifyCheck.EndPoint{}):                              dealSnElectrifyCheck.Init(apiRoot),
			api.GetName(dealSysDeviceSimFlowInfo.EndPoint{}):                          dealSysDeviceSimFlowInfo.Init(apiRoot),
			api.GetName(dealSysDeviceSimInfo.EndPoint{}):                              dealSysDeviceSimInfo.Init(apiRoot),
			api.GetName(definiteTimeDealSnExpRemind.EndPoint{}):                       definiteTimeDealSnExpRemind.Init(apiRoot),
			api.GetName(definiteTimeDealSnStatus.EndPoint{}):                          definiteTimeDealSnStatus.Init(apiRoot),
			api.GetName(delDeviceRepair.EndPoint{}):                                   delDeviceRepair.Init(apiRoot),
			api.GetName(delOperRule.EndPoint{}):                                       delOperRule.Init(apiRoot),
			api.GetName(delayCallApiResidueTimes.EndPoint{}):                          delayCallApiResidueTimes.Init(apiRoot),
			api.GetName(deleteComponent.EndPoint{}):                                   deleteComponent.Init(apiRoot),
			api.GetName(deleteCustomerEmployee.EndPoint{}):                            deleteCustomerEmployee.Init(apiRoot),
			api.GetName(deleteDeviceAccount.EndPoint{}):                               deleteDeviceAccount.Init(apiRoot),
			api.GetName(deleteDeviceSimById.EndPoint{}):                               deleteDeviceSimById.Init(apiRoot),
			api.GetName(deleteElectricitySettlementData.EndPoint{}):                   deleteElectricitySettlementData.Init(apiRoot),
			api.GetName(deleteFaultPlan.EndPoint{}):                                   deleteFaultPlan.Init(apiRoot),
			api.GetName(deleteFirmwareFiles.EndPoint{}):                               deleteFirmwareFiles.Init(apiRoot),
			api.GetName(deleteHouseholdEvaluation.EndPoint{}):                         deleteHouseholdEvaluation.Init(apiRoot),
			api.GetName(deleteHouseholdLeaveMessage.EndPoint{}):                       deleteHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(deleteHouseholdWorkOrder.EndPoint{}):                          deleteHouseholdWorkOrder.Init(apiRoot),
			api.GetName(deleteInverterSnInChnnl.EndPoint{}):                           deleteInverterSnInChnnl.Init(apiRoot),
			api.GetName(deleteModuleLog.EndPoint{}):                                   deleteModuleLog.Init(apiRoot),
			api.GetName(deleteOnDutyInfo.EndPoint{}):                                  deleteOnDutyInfo.Init(apiRoot),
			api.GetName(deleteOperateBillFile.EndPoint{}):                             deleteOperateBillFile.Init(apiRoot),
			api.GetName(deleteOssObject.EndPoint{}):                                   deleteOssObject.Init(apiRoot),
			api.GetName(deletePowerDevicePointById.EndPoint{}):                        deletePowerDevicePointById.Init(apiRoot),
			api.GetName(deletePowerPicture.EndPoint{}):                                deletePowerPicture.Init(apiRoot),
			api.GetName(deletePowerRobotInfoBySnAndPsId.EndPoint{}):                   deletePowerRobotInfoBySnAndPsId.Init(apiRoot),
			api.GetName(deletePowerRobotSweepStrategy.EndPoint{}):                     deletePowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(deleteProductionData.EndPoint{}):                              deleteProductionData.Init(apiRoot),
			api.GetName(deletePs.EndPoint{}):                                          deletePs.Init(apiRoot),
			api.GetName(deleteRechargeOrder.EndPoint{}):                               deleteRechargeOrder.Init(apiRoot),
			api.GetName(deleteRegularlyConnectionInfo.EndPoint{}):                     deleteRegularlyConnectionInfo.Init(apiRoot),
			api.GetName(deleteReportConfigEmailAddr.EndPoint{}):                       deleteReportConfigEmailAddr.Init(apiRoot),
			api.GetName(deleteSysAdvancedParam.EndPoint{}):                            deleteSysAdvancedParam.Init(apiRoot),
			api.GetName(deleteSysOrgNew.EndPoint{}):                                   deleteSysOrgNew.Init(apiRoot),
			api.GetName(deleteTemplate.EndPoint{}):                                    deleteTemplate.Init(apiRoot),
			api.GetName(deleteUserInfoAllByUserId.EndPoint{}):                         deleteUserInfoAllByUserId.Init(apiRoot),
			api.GetName(deviceInputDiscreteDeleteTime.EndPoint{}):                     deviceInputDiscreteDeleteTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteGetTime.EndPoint{}):                        deviceInputDiscreteGetTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteInsertTime.EndPoint{}):                     deviceInputDiscreteInsertTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteUpdateTime.EndPoint{}):                     deviceInputDiscreteUpdateTime.Init(apiRoot),
			api.GetName(deviceReplace.EndPoint{}):                                     deviceReplace.Init(apiRoot),
			api.GetName(editDeviceRepair.EndPoint{}):                                  editDeviceRepair.Init(apiRoot),
			api.GetName(editOperRule.EndPoint{}):                                      editOperRule.Init(apiRoot),
			api.GetName(energyPovertyAlleviation.EndPoint{}):                          energyPovertyAlleviation.Init(apiRoot),
			api.GetName(faultAutoClose.EndPoint{}):                                    faultAutoClose.Init(apiRoot),
			api.GetName(faultCloseRemindOrderHandler.EndPoint{}):                      faultCloseRemindOrderHandler.Init(apiRoot),
			api.GetName(findCodeValueList.EndPoint{}):                                 findCodeValueList.Init(apiRoot),
			api.GetName(findEmgOrgInfo.EndPoint{}):                                    findEmgOrgInfo.Init(apiRoot),
			api.GetName(findEnvironmentInfo.EndPoint{}):                               findEnvironmentInfo.Init(apiRoot),
			api.GetName(findFromHbaseAndRedis.EndPoint{}):                             findFromHbaseAndRedis.Init(apiRoot),
			api.GetName(findLossAnalysisList.EndPoint{}):                              findLossAnalysisList.Init(apiRoot),
			api.GetName(findOnDutyInfo.EndPoint{}):                                    findOnDutyInfo.Init(apiRoot),
			api.GetName(findSingleStationPR.EndPoint{}):                               findSingleStationPR.Init(apiRoot),
			api.GetName(findUserPassword.EndPoint{}):                                  findUserPassword.Init(apiRoot),
			api.GetName(genTLSUserSigByUserAccount.EndPoint{}):                        genTLSUserSigByUserAccount.Init(apiRoot),
			api.GetName(generateRandomPassword.EndPoint{}):                            generateRandomPassword.Init(apiRoot),
			api.GetName(getAPIServiceInfo.EndPoint{}):                                 getAPIServiceInfo.Init(apiRoot),
			api.GetName(getAccessedPermission.EndPoint{}):                             getAccessedPermission.Init(apiRoot),
			api.GetName(getAllDeviceByPsId.EndPoint{}):                                getAllDeviceByPsId.Init(apiRoot),
			api.GetName(getAllPowerRobotViewInfoByPsId.EndPoint{}):                    getAllPowerRobotViewInfoByPsId.Init(apiRoot),
			api.GetName(getAllPsIdByOrgIds.EndPoint{}):                                getAllPsIdByOrgIds.Init(apiRoot),
			api.GetName(getAllUserRemindCount.EndPoint{}):                             getAllUserRemindCount.Init(apiRoot),
			api.GetName(getAndOutletsAndUnit.EndPoint{}):                              getAndOutletsAndUnit.Init(apiRoot),
			api.GetName(getAreaInfoCodeByCounty.EndPoint{}):                           getAreaInfoCodeByCounty.Init(apiRoot),
			api.GetName(getAreaList.EndPoint{}):                                       getAreaList.Init(apiRoot),
			api.GetName(getAutoCreatePowerStation.EndPoint{}):                         getAutoCreatePowerStation.Init(apiRoot),
			api.GetName(getBackReadValue.EndPoint{}):                                  getBackReadValue.Init(apiRoot),
			api.GetName(getBatchNewestPointData.EndPoint{}):                           getBatchNewestPointData.Init(apiRoot),
			api.GetName(getCallApiResidueTimes.EndPoint{}):                            getCallApiResidueTimes.Init(apiRoot),
			api.GetName(getChangedPsListByTime.EndPoint{}):                            getChangedPsListByTime.Init(apiRoot),
			api.GetName(getChnnlListByPsId.EndPoint{}):                                getChnnlListByPsId.Init(apiRoot),
			api.GetName(getCloudList.EndPoint{}):                                      getCloudList.Init(apiRoot),
			api.GetName(getCloudServiceMappingConfig.EndPoint{}):                      getCloudServiceMappingConfig.Init(apiRoot),
			api.GetName(getCommunicationDeviceConfigInfo.EndPoint{}):                  getCommunicationDeviceConfigInfo.Init(apiRoot),
			api.GetName(getCommunicationModuleMonitorData.EndPoint{}):                 getCommunicationModuleMonitorData.Init(apiRoot),
			api.GetName(getComponentModelFactory.EndPoint{}):                          getComponentModelFactory.Init(apiRoot),
			api.GetName(getConfigList.EndPoint{}):                                     getConfigList.Init(apiRoot),
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
			api.GetName(getDeviceAccountById.EndPoint{}):                              getDeviceAccountById.Init(apiRoot),
			api.GetName(getDeviceFaultStatisticsData.EndPoint{}):                      getDeviceFaultStatisticsData.Init(apiRoot),
			api.GetName(getDeviceInfo.EndPoint{}):                                     getDeviceInfo.Init(apiRoot),
			api.GetName(getDeviceList.EndPoint{}):                                     getDeviceList.Init(apiRoot),
			api.GetName(getDeviceModelInfoList.EndPoint{}):                            getDeviceModelInfoList.Init(apiRoot),
			api.GetName(getDevicePropertys.EndPoint{}):                                getDevicePropertys.Init(apiRoot),
			api.GetName(getDeviceRepairDetail.EndPoint{}):                             getDeviceRepairDetail.Init(apiRoot),
			api.GetName(getDeviceTechBranchCount.EndPoint{}):                          getDeviceTechBranchCount.Init(apiRoot),
			api.GetName(getDeviceTypeInfoList.EndPoint{}):                             getDeviceTypeInfoList.Init(apiRoot),
			api.GetName(getDeviceTypeList.EndPoint{}):                                 getDeviceTypeList.Init(apiRoot),
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
			api.GetName(getInvertDataList.EndPoint{}):                                 getInvertDataList.Init(apiRoot),
			api.GetName(getInverterDataCount.EndPoint{}):                              getInverterDataCount.Init(apiRoot),
			api.GetName(getInverterProcess.EndPoint{}):                                getInverterProcess.Init(apiRoot),
			api.GetName(getInverterUuidBytotalId.EndPoint{}):                          getInverterUuidBytotalId.Init(apiRoot),
			api.GetName(getIvEchartsData.EndPoint{}):                                  getIvEchartsData.Init(apiRoot),
			api.GetName(getIvEchartsDataById.EndPoint{}):                              getIvEchartsDataById.Init(apiRoot),
			api.GetName(getKpiInfo.EndPoint{}):                                        getKpiInfo.Init(apiRoot),
			api.GetName(getMapInfo.EndPoint{}):                                        getMapInfo.Init(apiRoot),
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
			api.GetName(getOrgListByName.EndPoint{}):                                  getOrgListByName.Init(apiRoot),
			api.GetName(getOrgListByUserId.EndPoint{}):                                getOrgListByUserId.Init(apiRoot),
			api.GetName(getOrgListForUser.EndPoint{}):                                 getOrgListForUser.Init(apiRoot),
			api.GetName(getOssObjectStream.EndPoint{}):                                getOssObjectStream.Init(apiRoot),
			api.GetName(getOwnerFaultConfigList.EndPoint{}):                           getOwnerFaultConfigList.Init(apiRoot),
			api.GetName(getParamSetTemplate4NewProtocol.EndPoint{}):                   getParamSetTemplate4NewProtocol.Init(apiRoot),
			api.GetName(getParamSetTemplatePointInfo.EndPoint{}):                      getParamSetTemplatePointInfo.Init(apiRoot),
			api.GetName(getParamterSettingBase.EndPoint{}):                            getParamterSettingBase.Init(apiRoot),
			api.GetName(getPhotoInfo.EndPoint{}):                                      getPhotoInfo.Init(apiRoot),
			api.GetName(getPlanedOrNotPsList.EndPoint{}):                              getPlanedOrNotPsList.Init(apiRoot),
			api.GetName(getPlantReportPDFList.EndPoint{}):                             getPlantReportPDFList.Init(apiRoot),
			api.GetName(getPowerChargeSettingInfo.EndPoint{}):                         getPowerChargeSettingInfo.Init(apiRoot),
			api.GetName(getPowerDeviceModelTechList.EndPoint{}):                       getPowerDeviceModelTechList.Init(apiRoot),
			api.GetName(getPowerDeviceModelTree.EndPoint{}):                           getPowerDeviceModelTree.Init(apiRoot),
			api.GetName(getPowerDevicePointInfo.EndPoint{}):                           getPowerDevicePointInfo.Init(apiRoot),
			api.GetName(getPowerFormulaFaultAnalyse.EndPoint{}):                       getPowerFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(getPowerPictureList.EndPoint{}):                               getPowerPictureList.Init(apiRoot),
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
			api.GetName(getPsListByName.EndPoint{}):                                   getPsListByName.Init(apiRoot),
			api.GetName(getPsListForPsDataByPsId.EndPoint{}):                          getPsListForPsDataByPsId.Init(apiRoot),
			api.GetName(getPsListStaticData.EndPoint{}):                               getPsListStaticData.Init(apiRoot),
			api.GetName(getPsReport.EndPoint{}):                                       getPsReport.Init(apiRoot),
			api.GetName(getPsUser.EndPoint{}):                                         getPsUser.Init(apiRoot),
			api.GetName(getPsWeatherList.EndPoint{}):                                  getPsWeatherList.Init(apiRoot),
			api.GetName(getRechargeOrderDetail.EndPoint{}):                            getRechargeOrderDetail.Init(apiRoot),
			api.GetName(getRechargeOrderItemDeviceList.EndPoint{}):                    getRechargeOrderItemDeviceList.Init(apiRoot),
			api.GetName(getRechargeOrderList.EndPoint{}):                              getRechargeOrderList.Init(apiRoot),
			api.GetName(getRegionalTree.EndPoint{}):                                   getRegionalTree.Init(apiRoot),
			api.GetName(getRemoteParamSettingList.EndPoint{}):                         getRemoteParamSettingList.Init(apiRoot),
			api.GetName(getRemoteUpgradeDeviceList.EndPoint{}):                        getRemoteUpgradeDeviceList.Init(apiRoot),
			api.GetName(getRemoteUpgradeScheduleDetails.EndPoint{}):                   getRemoteUpgradeScheduleDetails.Init(apiRoot),
			api.GetName(getRemoteUpgradeSubTasksList.EndPoint{}):                      getRemoteUpgradeSubTasksList.Init(apiRoot),
			api.GetName(getRemoteUpgradeTaskList.EndPoint{}):                          getRemoteUpgradeTaskList.Init(apiRoot),
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
			api.GetName(getSweepDevParamSetTemplate.EndPoint{}):                       getSweepDevParamSetTemplate.Init(apiRoot),
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
			api.GetName(getUserList.EndPoint{}):                                       getUserList.Init(apiRoot),
			api.GetName(getUserPsOrderList.EndPoint{}):                                getUserPsOrderList.Init(apiRoot),
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
			api.GetName(insightSynDeviceStructure2Cloud.EndPoint{}):                   insightSynDeviceStructure2Cloud.Init(apiRoot),
			api.GetName(intoDataToHbase.EndPoint{}):                                   intoDataToHbase.Init(apiRoot),
			api.GetName(ipLocationQuery.EndPoint{}):                                   ipLocationQuery.Init(apiRoot),
			api.GetName(isHave2GSn.EndPoint{}):                                        isHave2GSn.Init(apiRoot),
			api.GetName(judgeDevIsHasInitSetTemplate.EndPoint{}):                      judgeDevIsHasInitSetTemplate.Init(apiRoot),
			api.GetName(judgeIsSettingMan.EndPoint{}):                                 judgeIsSettingMan.Init(apiRoot),
			api.GetName(listOssFiles.EndPoint{}):                                      listOssFiles.Init(apiRoot),
			api.GetName(loadAreaInfo.EndPoint{}):                                      loadAreaInfo.Init(apiRoot),
			api.GetName(loadPowerStation.EndPoint{}):                                  loadPowerStation.Init(apiRoot),
			api.GetName(loginByToken.EndPoint{}):                                      loginByToken.Init(apiRoot),
			api.GetName(logout.EndPoint{}):                                            logout.Init(apiRoot),
			api.GetName(mobilePhoneHasBound.EndPoint{}):                               mobilePhoneHasBound.Init(apiRoot),
			api.GetName(modifiedDeviceInfo.EndPoint{}):                                modifiedDeviceInfo.Init(apiRoot),
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
			api.GetName(queryDeviceRepairList.EndPoint{}):                             queryDeviceRepairList.Init(apiRoot),
			api.GetName(queryDeviceTypeInfoList.EndPoint{}):                           queryDeviceTypeInfoList.Init(apiRoot),
			api.GetName(queryEnvironmentList.EndPoint{}):                              queryEnvironmentList.Init(apiRoot),
			api.GetName(queryFaultList.EndPoint{}):                                    queryFaultList.Init(apiRoot),
			api.GetName(queryFaultPlanDetail.EndPoint{}):                              queryFaultPlanDetail.Init(apiRoot),
			api.GetName(queryFaultRepairSteps.EndPoint{}):                             queryFaultRepairSteps.Init(apiRoot),
			api.GetName(queryFaultTypeAndLevelByCode.EndPoint{}):                      queryFaultTypeAndLevelByCode.Init(apiRoot),
			api.GetName(queryFaultTypeByDevice.EndPoint{}):                            queryFaultTypeByDevice.Init(apiRoot),
			api.GetName(queryFaultTypeByDevicePage.EndPoint{}):                        queryFaultTypeByDevicePage.Init(apiRoot),
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
			api.GetName(queryUserList.EndPoint{}):                                     queryUserList.Init(apiRoot),
			api.GetName(queryUserProcessPri.EndPoint{}):                               queryUserProcessPri.Init(apiRoot),
			api.GetName(queryUserWechatBindRel.EndPoint{}):                            queryUserWechatBindRel.Init(apiRoot),
			api.GetName(queryUuidByTotalIdAndUuid.EndPoint{}):                         queryUuidByTotalIdAndUuid.Init(apiRoot),
			api.GetName(rechargeOrderSetMeal.EndPoint{}):                              rechargeOrderSetMeal.Init(apiRoot),
			api.GetName(renewSendReportConfirmEmail.EndPoint{}):                       renewSendReportConfirmEmail.Init(apiRoot),
			api.GetName(reportList.EndPoint{}):                                        reportList.Init(apiRoot),
			api.GetName(saveCustomerEmployee.EndPoint{}):                              saveCustomerEmployee.Init(apiRoot),
			api.GetName(saveDevSimList.EndPoint{}):                                    saveDevSimList.Init(apiRoot),
			api.GetName(saveDeviceAccountBatchData.EndPoint{}):                        saveDeviceAccountBatchData.Init(apiRoot),
			api.GetName(saveEnviromentIncomeInfos.EndPoint{}):                         saveEnviromentIncomeInfos.Init(apiRoot),
			api.GetName(saveEnvironmentCurve.EndPoint{}):                              saveEnvironmentCurve.Init(apiRoot),
			api.GetName(saveFirmwareFile.EndPoint{}):                                  saveFirmwareFile.Init(apiRoot),
			api.GetName(saveIncomeSettingInfos.EndPoint{}):                            saveIncomeSettingInfos.Init(apiRoot),
			api.GetName(saveOrUpdateGroupStringCheckRule.EndPoint{}):                  saveOrUpdateGroupStringCheckRule.Init(apiRoot),
			api.GetName(saveParamModel.EndPoint{}):                                    saveParamModel.Init(apiRoot),
			api.GetName(savePowerCharges.EndPoint{}):                                  savePowerCharges.Init(apiRoot),
			api.GetName(savePowerDevicePoint.EndPoint{}):                              savePowerDevicePoint.Init(apiRoot),
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
			api.GetName(saveTemplate.EndPoint{}):                                      saveTemplate.Init(apiRoot),
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
			api.GetName(stationDeviceHistoryDataList.EndPoint{}):                      stationDeviceHistoryDataList.Init(apiRoot),
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
			api.GetName(updateOrderDeviceByCustomerService.EndPoint{}):                updateOrderDeviceByCustomerService.Init(apiRoot),
			api.GetName(updateOwnerFaultConfig.EndPoint{}):                            updateOwnerFaultConfig.Init(apiRoot),
			api.GetName(updateParamSettingSysMsg.EndPoint{}):                          updateParamSettingSysMsg.Init(apiRoot),
			api.GetName(updatePlatformLevelFaultLevel.EndPoint{}):                     updatePlatformLevelFaultLevel.Init(apiRoot),
			api.GetName(updatePowerDevicePoint.EndPoint{}):                            updatePowerDevicePoint.Init(apiRoot),
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
			api.GetName(updateTemplate.EndPoint{}):                                    updateTemplate.Init(apiRoot),
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

			// "psHourPointsValue":                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/psHourPointsValue"}
			// "getPsCurveInfo":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getPsCurveInfo"}
			// "queryMutiPointDataList":                nullEndPoint.Init(apiRoot), // "/v1/commonService/queryMutiPointDataList"}
			// "getTemplateList":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getTemplateList"}
			// "getDevicePoints":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getDevicePoints"}
			// "getAllPowerDeviceSetName":              nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPowerDeviceSetName"}
			// "queryAllPsIdAndName":                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryAllPsIdAndName"}
			// "queryDevicePointMinuteDataList":        nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointMinuteDataList"}
			// "getTemplateByInfoType":                 nullEndPoint.Init(apiRoot), // "/v1/messageService/getTemplateByInfoType"}
			// "energyTrend":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/energyTrend"}
			// "queryBatchCreatePsTaskList":            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryBatchCreatePsTaskList"}
			// "exportPlantReportPDF":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/exportPlantReportPDF"}
			// "getTableDataSql":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getTableDataSql"}
			// "getTableDataSqlCount":                  nullEndPoint.Init(apiRoot), // "/v1/devService/getTableDataSqlCount"}
			// "getPListinfoFromMysql":                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPListinfoFromMysql"}
			// "getDataFromHBase":                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getDataFromHBase"}
			// "getListMiFromHBase":                    nullEndPoint.Init(apiRoot), // "/v1/commonService/getListMiFromHBase"}
			// "getMapMiFromHBase":                     nullEndPoint.Init(apiRoot), // "/v1/commonService/getMapMiFromHBase"}
			// "getValFromHBase":                       nullEndPoint.Init(apiRoot), // "/v1/commonService/getValFromHBase"}
			// "getPowerStationTableDataSql":           nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerStationTableDataSql"}
			// "getPowerStationTableDataSqlCount":      nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerStationTableDataSqlCount"}
			// "getDataFromHbaseByRowKey":              nullEndPoint.Init(apiRoot), // "/v1/commonService/getDataFromHbaseByRowKey"}
			// "findInfoByuuid":                        nullEndPoint.Init(apiRoot), // "/v1/devService/findInfoByuuid"}
			// "getStationInfoSql":                     nullEndPoint.Init(apiRoot), // "/v1/devService/getStationInfoSql"}
			// "getDevicePointMinuteDataList":          nullEndPoint.Init(apiRoot), // "/v1/commonService/getDevicePointMinuteDataList"}
			// "powerDevicePointList":                  nullEndPoint.Init(apiRoot), // "/v1/reportService/powerDevicePointList"}
			// "getPowerTrendDayData":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerTrendDayData"}
			// "powerTrendChartData":                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/powerTrendChartData"}
			// "getUpTimePoint":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getUpTimePoint"}
			// "getSerialNum":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getSerialNum"}
			// "getModuleLogTaskList":                  nullEndPoint.Init(apiRoot), // "/integrationService/getModuleLogTaskList"}
			// "queryParamSettingTask":                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryParamSettingTask"}
			// "queryCtrlTaskById":                     nullEndPoint.Init(apiRoot), // "/v1/devService/queryCtrlTaskById"}
			// "queryDeviceInfo":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceInfoForApp"}
			// "queryDeviceInfoForApp":                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceInfoForApp"}
			// "queryDeviceList":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceList"}
			// "queryDeviceListByUserId":               nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceListByUserId"}
			// "queryDeviceListForApp":                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceListForApp"}
			// "queryDeviceModelTechnical":             nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceModelTechnical"}
			// "getPsDataSupplementTaskList":           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDataSupplementTaskList"}
			// "getPsDetail":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetail"}
			// "getPsHealthState":                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsHealthState"}
			// "getPowerDeviceSetTaskDetailList":       nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceSetTaskDetailList"}
			// "getPowerDeviceSetTaskList":             nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceSetTaskList"}
			// "getPsAuthKey":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsAuthKey"}
			// "getApiCallsForAppkeys":                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getApiCallsForAppkeys"}
			// "exportParamSettingValPDF":              nullEndPoint.Init(apiRoot), // "/v1/devService/exportParamSettingValPDF"}
			// "checkUnitStatus":                       nullEndPoint.Init(apiRoot), // "/v1/devService/checkUnitStatus"}
			// "devicePointsDataFromMySql":             nullEndPoint.Init(apiRoot), // "/v1/devService/devicePointsDataFromMySql"}
			// "queryDevicePointDayMonthYearDataList":  nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointDayMonthYearDataList"}
			// "queryDevicePointsDayMonthYearDataList": nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointsDayMonthYearDataList"}
			// "queryDeviceRealTimeDataByPsKeys":       nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceRealTimeDataByPsKeys"}
			// "acceptPsSharing":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/acceptPsSharing"}
			// "activateEmail":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/activateEmail"}
			// "addConfig":                                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/addConfig"}
			// "addDeviceRepair":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/addDeviceRepair"}
			// "addDeviceToStructureForHousehold":                  nullEndPoint.Init(apiRoot), // "/devDataHandleService/addDeviceToStructureForHousehold"}
			// "addDeviceToStructureForHouseholdByPsIdS":           nullEndPoint.Init(apiRoot), // "/devDataHandleService/addDeviceToStructureForHouseholdByPsIdS"}
			// "addFault":                                          nullEndPoint.Init(apiRoot), // "/v1/faultService/addFault"}
			// "addFaultOrder":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultOrder"}
			// "addFaultPlan":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultPlan"}
			// "addFaultRepairSteps":                               nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultRepairSteps"}
			// "addHouseholdEvaluation":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdEvaluation"}
			// "addHouseholdLeaveMessage":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdLeaveMessage"}
			// "addHouseholdOpinionFeedback":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdOpinionFeedback"}
			// "addHouseholdWorkOrder":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdWorkOrder"}
			// "addOnDutyInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/otherService/addOnDutyInfo"}
			// "addOperRule":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/addOperRule"}
			// "addOrDelPsStructure":                               nullEndPoint.Init(apiRoot), // "/v1/devService/addOrDelPsStructure"}
			// "addOrderStep":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/addOrderStep"}
			// "addPowerStationForHousehold":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/addPowerStationForHousehold"}
			// "addPowerStationInfo":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/addPowerStationInfo"}
			// "addReportConfigEmail":                              nullEndPoint.Init(apiRoot), // "/v1/reportService/addReportConfigEmail"}
			// "addSysAdvancedParam":                               nullEndPoint.Init(apiRoot), // "/v1/devService/addSysAdvancedParam"}
			// "addSysOrgNew":                                      nullEndPoint.Init(apiRoot), // "/v1/otherService/addSysOrgNew"}
			// "aliPayAppTest":                                     nullEndPoint.Init(apiRoot), // "/onlinepay/aliPayAppTest"}
			// "auditOperRule":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/auditOperRule"}
			// "batchAddStationBySn":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/batchAddStationBySn"}
			// "batchImportSN":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/batchImportSN"}
			// "batchInsertUserAndOrg":                             nullEndPoint.Init(apiRoot), // "/v1/userService/batchInsertUserAndOrg"}
			// "batchModifyDevicesInfoAndPropertis":                nullEndPoint.Init(apiRoot), // "/v1/devService/batchModifyDevicesInfoAndPropertis"}
			// "batchProcessPlantReport":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/batchProcessPlantReport"}
			// "batchUpdateDeviceSim":                              nullEndPoint.Init(apiRoot), // "/v1/devService/batchUpdateDeviceSim"}
			// "batchUpdateUserIsAgreeGdpr":                        nullEndPoint.Init(apiRoot), // "/v1/userService/batchUpdateUserIsAgreeGdpr"}
			// "boundMobilePhone":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/boundMobilePhone"}
			// "boundUserMail":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/boundUserMail"}
			// "caculateDeviceInputDiscrete":                       nullEndPoint.Init(apiRoot), // "/v1/devService/caculateDeviceInputDiscrete"}
			// "calculateDeviceDiscrete":                           nullEndPoint.Init(apiRoot), // "/v1/devService/calculateDeviceDiscrete"}
			// "calculateInitialCompensationData":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/calculateInitialCompensationData"}
			// "cancelDeliverMail":                                 nullEndPoint.Init(apiRoot), // "/v1/messageService/cancelDeliverMail"}
			// "cancelOrderScan":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/cancelOrderScan"}
			// "cancelParamSetTask":                                nullEndPoint.Init(apiRoot), // "/v1/devService/cancelParamSetTask"}
			// "cancelPsSharing":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/cancelPsSharing"}
			// "cancelRechargeOrder":                               nullEndPoint.Init(apiRoot), // "/onlinepay/cancelRechargeOrder"}
			// "changRechargeOrderToCancel":                        nullEndPoint.Init(apiRoot), // "/onlinepay/changRechargeOrderToCancel"}
			// "changeHouseholdUser2Installer":                     nullEndPoint.Init(apiRoot), // "/v1/orgService/changeHouseholdUser2Installer"}
			// "changeRemoteParam":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/changeRemoteParam"}
			// "checkDealerOrgCode":                                nullEndPoint.Init(apiRoot), // "/v1/orgService/checkDealerOrgCode"}
			// "checkDevSnIsBelongsToUser":                         nullEndPoint.Init(apiRoot), // "/v1/userService/checkDevSnIsBelongsToUser"}
			// "checkInverterResult":                               nullEndPoint.Init(apiRoot), // "/v1/devService/checkInverterResult"}
			// "checkIsCanDoParamSet":                              nullEndPoint.Init(apiRoot), // "/v1/devService/checkIsCanDoParamSet"}
			// "checkIsIvScan":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/checkIsIvScan"}
			// "checkOssObjectExist":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/checkOssObjectExist"}
			// "checkServiceIsConnect":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/checkServiceIsConnect"}
			// "checkTechnicalParameters":                          nullEndPoint.Init(apiRoot), // "/v1/devService/checkTechnicalParameters"}
			// "checkUpRechargeDevicePaying":                       nullEndPoint.Init(apiRoot), // "/onlinepay/checkUpRechargeDevicePaying"}
			// "checkUserAccountUnique":                            nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserAccountUnique"}
			// "checkUserAccountUniqueAll":                         nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserAccountUniqueAll"}
			// "checkUserInfoUnique":                               nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserInfoUnique"}
			// "checkUserIsExist":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserIsExist"}
			// "checkUserListIsExist":                              nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserListIsExist"}
			// "checkUserPassword":                                 nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserPassword"}
			// "cloudDeploymentRecord":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/cloudDeploymentRecord"}
			// "comfirmParamModel":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/comfirmParamModel"}
			// "compareValidateCode":                               nullEndPoint.Init(apiRoot), // "/v1/userService/compareValidateCode"}
			// "componentInfo2Cloud":                               nullEndPoint.Init(apiRoot), // "/v1/devService/componentInfo2Cloud"}
			// "confirmFault":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/confirmFault"}
			// "confirmIvFault":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/confirmIvFault"}
			// "confirmReportConfig":                               nullEndPoint.Init(apiRoot), // "/v1/reportService/confirmReportConfig"}
			// "createAppkeyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/createAppkeyInfo"}
			// "createRenewInvoice":                                nullEndPoint.Init(apiRoot), // "/onlinepay/createRenewInvoice"}
			// "dealCommandReply":                                  nullEndPoint.Init(apiRoot), // "/devDataHandleService/dealCommandReply"}
			// "dealDeletePsFailPsDelete":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/dealDeletePsFailPsDelete"}
			// "dealFailRemoteUpgradeSubTasks":                     nullEndPoint.Init(apiRoot), // "/v1/devService/dealFailRemoteUpgradeSubTasks"}
			// "dealFailRemoteUpgradeTasks":                        nullEndPoint.Init(apiRoot), // "/v1/devService/dealFailRemoteUpgradeTasks"}
			// "dealFaultOrder":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/dealFaultOrder"}
			// "dealGroupStringDisableOrEnable":                    nullEndPoint.Init(apiRoot), // "/v1/devService/dealGroupStringDisableOrEnable"}
			// "dealNumberOfServiceCalls2Mysql":                    nullEndPoint.Init(apiRoot), // "/v1/commonService/dealNumberOfServiceCalls2Mysql"}
			// "dealParamSettingAfterComplete":                     nullEndPoint.Init(apiRoot), // "/v1/devService/dealParamSettingAfterComplete"}
			// "dealPsDataSupplement":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/dealPsDataSupplement"}
			// "dealPsReportEmailSend":                             nullEndPoint.Init(apiRoot), // "/v1/reportService/dealPsReportEmailSend"}
			// "dealRemoteUpgrade":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/dealRemoteUpgrade"}
			// "dealSnElectrifyCheck":                              nullEndPoint.Init(apiRoot), // "/v1/devService/dealSnElectrifyCheck"}
			// "dealSysDeviceSimFlowInfo":                          nullEndPoint.Init(apiRoot), // "/v1/devService/dealSysDeviceSimFlowInfo"}
			// "dealSysDeviceSimInfo":                              nullEndPoint.Init(apiRoot), // "/v1/devService/dealSysDeviceSimInfo"}
			// "definiteTimeDealSnExpRemind":                       nullEndPoint.Init(apiRoot), // "/v1/devService/definiteTimeDealSnExpRemind"}
			// "definiteTimeDealSnStatus":                          nullEndPoint.Init(apiRoot), // "/v1/devService/definiteTimeDealSnStatus"}
			// "delDeviceRepair":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/delDeviceRepair"}
			// "delOperRule":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/delOperRule"}
			// "delayCallApiResidueTimes":                          nullEndPoint.Init(apiRoot), // "/v1/commonService/delayCallApiResidueTimes"}
			// "deleteComponent":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/deleteComponent"}
			// "deleteCustomerEmployee":                            nullEndPoint.Init(apiRoot), // "/v1/devService/deleteCustomerEmployee"}
			// "deleteDeviceAccount":                               nullEndPoint.Init(apiRoot), // "/v1/devService/deleteDeviceAccount"}
			// "deleteDeviceSimById":                               nullEndPoint.Init(apiRoot), // "/v1/devService/deleteDeviceSimById"}
			// "deleteElectricitySettlementData":                   nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteElectricitySettlementData"}
			// "deleteFaultPlan":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteFaultPlan"}
			// "deleteFirmwareFiles":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteFirmwareFiles"}
			// "deleteHouseholdEvaluation":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdEvaluation"}
			// "deleteHouseholdLeaveMessage":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdLeaveMessage"}
			// "deleteHouseholdWorkOrder":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdWorkOrder"}
			// "deleteInverterSnInChnnl":                           nullEndPoint.Init(apiRoot), // "/devDataHandleService/deleteInverterSnInChnnl"}
			// "deleteModuleLog":                                   nullEndPoint.Init(apiRoot), // "/integrationService/deleteModuleLog"}
			// "deleteOnDutyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteOnDutyInfo"}
			// "deleteOperateBillFile":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteOperateBillFile"}
			// "deleteOssObject":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteOssObject"}
			// "deletePowerDevicePointById":                        nullEndPoint.Init(apiRoot), // "/v1/reportService/deletePowerDevicePointById"}
			// "deletePowerPicture":                                nullEndPoint.Init(apiRoot), // "/v1/powerStationService/deletePowerPicture"}
			// "deletePowerRobotInfoBySnAndPsId":                   nullEndPoint.Init(apiRoot), // "/v1/devService/deletePowerRobotInfoBySnAndPsId"}
			// "deletePowerRobotSweepStrategy":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deletePowerRobotSweepStrategy"}
			// "deleteProductionData":                              nullEndPoint.Init(apiRoot), // "/v1/devService/deleteProductionData"}
			// "deletePs":                                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/deletePs"}
			// "deleteRechargeOrder":                               nullEndPoint.Init(apiRoot), // "/onlinepay/deleteRechargeOrder"}
			// "deleteRegularlyConnectionInfo":                     nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteRegularlyConnectionInfo"}
			// "deleteReportConfigEmailAddr":                       nullEndPoint.Init(apiRoot), // "/v1/reportService/deleteReportConfigEmailAddr"}
			// "deleteSysAdvancedParam":                            nullEndPoint.Init(apiRoot), // "/v1/devService/deleteSysAdvancedParam"}
			// "deleteSysOrgNew":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteSysOrgNew"}
			// "deleteTemplate":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/deleteTemplate"}
			// "deleteUserInfoAllByUserId":                         nullEndPoint.Init(apiRoot), // "/v1/userService/deleteUserInfoAllByUserId"}
			// "deviceInputDiscreteDeleteTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteDeleteTime"}
			// "deviceInputDiscreteGetTime":                        nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteGetTime"}
			// "deviceInputDiscreteInsertTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteInsertTime"}
			// "deviceInputDiscreteUpdateTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteUpdateTime"}
			// "deviceReplace":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceReplace"}
			// "editDeviceRepair":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/editDeviceRepair"}
			// "editOperRule":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/editOperRule"}
			// "energyPovertyAlleviation":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/energyPovertyAlleviation"}
			// "faultAutoClose":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/faultAutoClose"}
			// "faultCloseRemindOrderHandler":                      nullEndPoint.Init(apiRoot), // "/v1/faultService/faultCloseRemindOrderHandler"}
			// "findCodeValueList":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/findCodeValueList"}
			// "findEmgOrgInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/findEmgOrgInfo"}
			// "findEnvironmentInfo":                               nullEndPoint.Init(apiRoot), // "/v1/devService/findEnvironmentInfo"}
			// "findFromHbaseAndRedis":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/findFromHbaseAndRedis"}
			// "findLossAnalysisList":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/findLossAnalysisList"}
			// "findOnDutyInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/findOnDutyInfo"}
			// "findSingleStationPR":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/findSingleStationPR"}
			// "findUserPassword":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/findUserPassword"}
			// "genTLSUserSigByUserAccount":                        nullEndPoint.Init(apiRoot), // "/v1/userService/genTLSUserSigByUserAccount"}
			// "generateRandomPassword":                            nullEndPoint.Init(apiRoot), // "/v1/userService/generateRandomPassword"}
			// "getAPIServiceInfo":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getAPIServiceInfo"}
			// "getAccessedPermission":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/getAccessedPermission"}
			// "getAllDeviceByPsId":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getAllDeviceByPsId"}
			// "getAllPowerRobotViewInfoByPsId":                    nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPowerRobotViewInfoByPsId"}
			// "getAllPsIdByOrgIds":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPsIdByOrgIds"}
			// "getAllUserRemindCount":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getAllUserRemindCount"}
			// "getAndOutletsAndUnit":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getAndOutletsAndUnit"}
			// "getAreaInfoCodeByCounty":                           nullEndPoint.Init(apiRoot), // "/v1/commonService/getAreaInfoCodeByCounty"}
			// "getAreaList":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getAreaList"}
			// "getAutoCreatePowerStation":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getAutoCreatePowerStation"}
			// "getBackReadValue":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getBackReadValue"}
			// "getBatchNewestPointData":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getBatchNewestPointData"}
			// "getCallApiResidueTimes":                            nullEndPoint.Init(apiRoot), // "/v1/commonService/getCallApiResidueTimes"}
			// "getChangedPsListByTime":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getChangedPsListByTime"}
			// "getChnnlListByPsId":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getChnnlListByPsId"}
			// "getCloudList":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getCloudList"}
			// "getCloudServiceMappingConfig":                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getCloudServiceMappingConfig"}
			// "getCommunicationDeviceConfigInfo":                  nullEndPoint.Init(apiRoot), // "/v1/devService/getCommunicationDeviceConfigInfo"}
			// "getCommunicationModuleMonitorData":                 nullEndPoint.Init(apiRoot), // "/v1/devService/getCommunicationModuleMonitorData"}
			// "getComponentModelFactory":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getComponentModelFactory"}
			// "getConfigList":                                     nullEndPoint.Init(apiRoot), // "/devDataHandleService/getConfigList"}
			// "getConnectionInfoBySnAndLocalPort":                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getConnectionInfoBySnAndLocalPort"}
			// "getCountDown":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/getCountDown"}
			// "getCountryServiceInfo":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/getCountryServiceInfo"}
			// "getCounty":                                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getCounty"}
			// "getCustomerEmployee":                               nullEndPoint.Init(apiRoot), // "/v1/devService/getCustomerEmployee"}
			// "getCustomerList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getCustomerList"}
			// "getDevInstalledPowerByPsId":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getDevInstalledPowerByPsId"}
			// "getDevRecord":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/getDevRecord"}
			// "getDevRunRecordList":                               nullEndPoint.Init(apiRoot), // "/v1/devService/getDevRunRecordList"}
			// "getDevSimByList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getDevSimByList"}
			// "getDevSimList":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDevSimList"}
			// "getDeviceAccountById":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceAccountById"}
			// "getDeviceFaultStatisticsData":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceFaultStatisticsData"}
			// "getDeviceInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceInfo"}
			// "getDeviceList":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceList"}
			// "getDeviceModelInfoList":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceModelInfoList"}
			// "getDevicePropertys":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getDevicePropertys"}
			// "getDeviceRepairDetail":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceRepairDetail"}
			// "getDeviceTechBranchCount":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTechBranchCount"}
			// "getDeviceTypeInfoList":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTypeInfoList"}
			// "getDeviceTypeList":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTypeList"}
			// "getDstInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/userService/getDstInfo"}
			// "getElectricitySettlementData":                      nullEndPoint.Init(apiRoot), // "/v1/otherService/getElectricitySettlementData"}
			// "getElectricitySettlementDetailData":                nullEndPoint.Init(apiRoot), // "/v1/otherService/getElectricitySettlementDetailData"}
			// "getEncryptPublicKey":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/getEncryptPublicKey"}
			// "getFaultCount":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultCount"}
			// "getFaultDetail":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultDetail"}
			// "getFaultMsgByFaultCode":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgByFaultCode"}
			// "getFaultMsgListByPageWithYYYYMM":                   nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgListByPageWithYYYYMM"}
			// "getFaultMsgListWithYYYYMM":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgListWithYYYYMM"}
			// "getFaultPlanList":                                  nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultPlanList"}
			// "getFileOperationRecordOne":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getFileOperationRecordOne"}
			// "getFormulaFaultAnalyseList":                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getFormulaFaultAnalyseList"}
			// "getGroupStringCheckResult":                         nullEndPoint.Init(apiRoot), // "/v1/devService/getGroupStringCheckResult"}
			// "getGroupStringCheckRule":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getGroupStringCheckRule"}
			// "getHisData":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/getHisData"}
			// "getHouseholdEvaluation":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdEvaluation"}
			// "getHouseholdLeaveMessage":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdLeaveMessage"}
			// "getHouseholdOpinionFeedback":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdOpinionFeedback"}
			// "getHouseholdPsInstallerByUserId":                   nullEndPoint.Init(apiRoot), // "/v1/userService/getHouseholdPsInstallerByUserId"}
			// "getHouseholdUserInfo":                              nullEndPoint.Init(apiRoot), // "/v1/userService/getHouseholdUserInfo"}
			// "getHouseholdWorkOrderInfo":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdWorkOrderInfo"}
			// "getHouseholdWorkOrderList":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdWorkOrderList"}
			// "getI18nConfigByType":                               nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nConfigByType"}
			// "getI18nFileInfo":                                   nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nFileInfo"}
			// "getI18nInfoByKey":                                  nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nInfoByKey"}
			// "getI18nVersion":                                    nullEndPoint.Init(apiRoot), // "/integrationService/international/getI18nVersion"}
			// "getIncomeSettingInfos":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getIncomeSettingInfos"}
			// "getInfoFromAMap":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/getInfoFromAMap"}
			// "getInfomationFromRedis":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getInfomationFromRedis"}
			// "getInstallInfoList":                                nullEndPoint.Init(apiRoot), // "/v1/orgService/getInstallInfoList"}
			// "getInstallerInfoByDealerOrgCodeOrId":               nullEndPoint.Init(apiRoot), // "/v1/orgService/getInstallerInfoByDealerOrgCodeOrId"}
			// "getInvertDataList":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getInvertDataList"}
			// "getInverterDataCount":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterDataCount"}
			// "getInverterProcess":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterProcess"}
			// "getInverterUuidBytotalId":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterUuidBytotalId"}
			// "getIvEchartsData":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getIvEchartsData"}
			// "getIvEchartsDataById":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getIvEchartsDataById"}
			// "getKpiInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getKpiInfo"}
			// "getMapInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getMapInfo"}
			// "getMenuAndPrivileges":                              nullEndPoint.Init(apiRoot), // "/v1/userService/getMenuAndPrivileges"}
			// "getMicrogridEStoragePsReport":                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getMicrogridEStoragePsReport"}
			// "getModuleLogInfo":                                  nullEndPoint.Init(apiRoot), // "/integrationService/getModuleLogInfo"}
			// "getNationProvJSON":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getNationProvJSON"}
			// "getNeedOpAsynOpRecordList":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getNeedOpAsynOpRecordList"}
			// "getNoticeInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/otherService/getNoticeInfo"}
			// "getNumberOfServiceCalls":                           nullEndPoint.Init(apiRoot), // "/v1/commonService/getNumberOfServiceCalls"}
			// "getOSSConfig":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getOSSConfig"}
			// "getOperRuleDetail":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperRuleDetail"}
			// "getOperateBillFileId":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperateBillFileId"}
			// "getOperateTicketForDetail":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperateTicketForDetail"}
			// "getOrCreateNetEaseUserToken":                       nullEndPoint.Init(apiRoot), // "/v1/userService/getOrCreateNetEaseUserToken"}
			// "getOrderDataList":                                  nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderDataList"}
			// "getOrderDataSql2":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getOrderDataSql2"}
			// "getOrderDatas":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getOrderDatas"}
			// "getOrderDetail":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderDetail"}
			// "getOrderStatistics":                                nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderStatistics"}
			// "getOrgIdNameByUserId":                              nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgIdNameByUserId"}
			// "getOrgInfoByDealerOrgCode":                         nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgInfoByDealerOrgCode"}
			// "getOrgListByName":                                  nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgListByName"}
			// "getOrgListByUserId":                                nullEndPoint.Init(apiRoot), // "/v1/userService/getOrgListByUserId"}
			// "getOrgListForUser":                                 nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgListForUser"}
			// "getOssObjectStream":                                nullEndPoint.Init(apiRoot), // "/v1/commonService/getOssObjectStream"}
			// "getOwnerFaultConfigList":                           nullEndPoint.Init(apiRoot), // "/v1/faultService/getOwnerFaultConfigList"}
			// "getParamSetTemplate4NewProtocol":                   nullEndPoint.Init(apiRoot), // "/v1/devService/getParamSetTemplate4NewProtocol"}
			// "getParamSetTemplatePointInfo":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getParamSetTemplatePointInfo"}
			// "getParamterSettingBase":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getParamterSettingBase"}
			// "getPhotoInfo":                                      nullEndPoint.Init(apiRoot), // "/v1/otherService/getPhotoInfo"}
			// "getPlanedOrNotPsList":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/getPlanedOrNotPsList"}
			// "getPlantReportPDFList":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPlantReportPDFList"}
			// "getPowerChargeSettingInfo":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerChargeSettingInfo"}
			// "getPowerDeviceModelTechList":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceModelTechList"}
			// "getPowerDeviceModelTree":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceModelTree"}
			// "getPowerDevicePointInfo":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/getPowerDevicePointInfo"}
			// "getPowerFormulaFaultAnalyse":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerFormulaFaultAnalyse"}
			// "getPowerPictureList":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerPictureList"}
			// "getPowerRobotInfoByRobotSn":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotInfoByRobotSn"}
			// "getPowerRobotSweepAttrByPsId":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepAttrByPsId"}
			// "getPowerRobotSweepStrategy":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepStrategy"}
			// "getPowerRobotSweepStrategyList":                    nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepStrategyList"}
			// "getPowerSettingCharges":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerSettingCharges"}
			// "getPowerSettingHistoryRecords":                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerSettingHistoryRecords"}
			// "getPowerStationBasicInfo":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationBasicInfo"}
			// "getPowerStationData":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationData"}
			// "getPowerStationForHousehold":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationForHousehold"}
			// "getPowerStationInfo":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationInfo"}
			// "getPowerStationPR":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationPR"}
			// "getPrivateCloudValidityPeriod":                     nullEndPoint.Init(apiRoot), // "/v1/commonService/getPrivateCloudValidityPeriod"}
			// "getProvInfoListByNationCode":                       nullEndPoint.Init(apiRoot), // "/v1/commonService/getProvInfoListByNationCode"}
			// "getPsDetailByUserTokens":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetailByUserTokens"}
			// "getPsDetailForSinglePage":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetailForSinglePage"}
			// "getPsInstallerByPsId":                              nullEndPoint.Init(apiRoot), // "/v1/orgService/getPsInstallerByPsId"}
			// "getPsInstallerOrgInfoByPsId":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsInstallerOrgInfoByPsId"}
			// "getPsListByName":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListByName"}
			// "getPsListForPsDataByPsId":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListForPsDataByPsId"}
			// "getPsListStaticData":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListStaticData"}
			// "getPsReport":                                       nullEndPoint.Init(apiRoot), // "/v1/reportService/getPsReport"}
			// "getPsUser":                                         nullEndPoint.Init(apiRoot), // "/v1/userService/getPsUser"}
			// "getPsWeatherList":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsWeatherList"}
			// "getRechargeOrderDetail":                            nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderDetail"}
			// "getRechargeOrderItemDeviceList":                    nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderItemDeviceList"}
			// "getRechargeOrderList":                              nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderList"}
			// "getRegionalTree":                                   nullEndPoint.Init(apiRoot), // "/v1/orgService/getRegionalTree"}
			// "getRemoteParamSettingList":                         nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteParamSettingList"}
			// "getRemoteUpgradeDeviceList":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeDeviceList"}
			// "getRemoteUpgradeScheduleDetails":                   nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeScheduleDetails"}
			// "getRemoteUpgradeSubTasksList":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeSubTasksList"}
			// "getRemoteUpgradeTaskList":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeTaskList"}
			// "getReportData":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getReportData"}
			// "getReportEmailConfigInfo":                          nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportEmailConfigInfo"}
			// "getReportExportColumns":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportExportColumns"}
			// "getReportListByUserId":                             nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportListByUserId"}
			// "getRobotDynamicCleaningView":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getRobotDynamicCleaningView"}
			// "getRobotNumAndSweepCapacity":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getRobotNumAndSweepCapacity"}
			// "getRuleUnit":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getRuleUnit"}
			// "getSendReportConfigCron":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/getSendReportConfigCron"}
			// "getShieldMapConditionList":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getShieldMapConditionList"}
			// "getSimIdBySnList":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getSimIdBySnList"}
			// "getSingleIVData":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getSingleIVData"}
			// "getSnChangeRecord":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getSnChangeRecord"}
			// "getSnConnectionInfo":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/getSnConnectionInfo"}
			// "getSungwsConfigCache":                              nullEndPoint.Init(apiRoot), // "/v1/commonService/getSungwsConfigCache"}
			// "getSungwsGlobalConfigCache":                        nullEndPoint.Init(apiRoot), // "/v1/commonService/getSungwsGlobalConfigCache"}
			// "getSweepDevParamSetTemplate":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getSweepDevParamSetTemplate"}
			// "getSweepRobotDevList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getSweepRobotDevList"}
			// "getSysMsg":                                         nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysMsg"}
			// "getSysOrgNewList":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysOrgNewList"}
			// "getSysOrgNewOne":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysOrgNewOne"}
			// "getSysUserById":                                    nullEndPoint.Init(apiRoot), // "/v1/userService/getSysUserById"}
			// "getUUIDByUpuuid":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getUUIDByUpuuid"}
			// "getUserById":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/getUserById"}
			// "getUserByInstaller":                                nullEndPoint.Init(apiRoot), // "/v1/userService/getUserByInstaller"}
			// "getUserDevOnlineOffineCount":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getUserDevOnlineOffineCount"}
			// "getUserGDPRAttrs":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/getUserGDPRAttrs"}
			// "getUserHavePowerStationCount":                      nullEndPoint.Init(apiRoot), // "/v1/userService/getUserHavePowerStationCount"}
			// "getUserInfoByUserAccounts":                         nullEndPoint.Init(apiRoot), // "/v1/userService/getUserInfoByUserAccounts"}
			// "getUserList":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/getUserList"}
			// "getUserPsOrderList":                                nullEndPoint.Init(apiRoot), // "/v1/faultService/getUserPsOrderList"}
			// "getValidateCode":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/getValidateCode"}
			// "getValidateCodeAtRegister":                         nullEndPoint.Init(apiRoot), // "/v1/userService/getValidateCodeAtRegister"}
			// "getWeatherInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getWeatherInfo"}
			// "getWechatPushConfig":                               nullEndPoint.Init(apiRoot), // "/v1/userService/getWechatPushConfig"}
			// "getWorkInfo":                                       nullEndPoint.Init(apiRoot), // "/v1/otherService/getWorkInfo"}
			// "groupStringCheck":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/groupStringCheck"}
			// "handleDevByCommunicationSN":                        nullEndPoint.Init(apiRoot), // "/devDataHandleService/handleDevByCommunicationSN"}
			// "householdResetPassBySN":                            nullEndPoint.Init(apiRoot), // "/v1/userService/householdResetPassBySN"}
			// "immediatePayment":                                  nullEndPoint.Init(apiRoot), // "/onlinepay/immediatePayment"}
			// "importExcelData":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/importExcelData"}
			// "incomeStatistics":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/incomeStatistics"}
			// "informPush":                                        nullEndPoint.Init(apiRoot), // "/v1/messageService/informPush"}
			// "insertEmgOrgInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/insertEmgOrgInfo"}
			// "insightSynDeviceStructure2Cloud":                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/insightSynDeviceStructure2Cloud"}
			// "intoDataToHbase":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/intoDataToHbase"}
			// "ipLocationQuery":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/ipLocationQuery"}
			// "isHave2GSn":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/isHave2GSn"}
			// "judgeDevIsHasInitSetTemplate":                      nullEndPoint.Init(apiRoot), // "/v1/devService/judgeDevIsHasInitSetTemplate"}
			// "judgeIsSettingMan":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/judgeIsSettingMan"}
			// "listOssFiles":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/listOssFiles"}
			// "loadAreaInfo":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/loadAreaInfo"}
			// "loadPowerStation":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/loadPowerStation"}
			// "loginByToken":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/loginByToken"}
			// "logout":                                            nullEndPoint.Init(apiRoot), // "/v1/userService/logout"}
			// "mobilePhoneHasBound":                               nullEndPoint.Init(apiRoot), // "/v1/userService/mobilePhoneHasBound"}
			// "modifiedDeviceInfo":                                nullEndPoint.Init(apiRoot), // "/v1/devService/modifiedDeviceInfo"}
			// "modifyEmgOrgStruc":                                 nullEndPoint.Init(apiRoot), // "/v1/otherService/modifyEmgOrgStruc"}
			// "modifyFaultPlan":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/modifyFaultPlan"}
			// "modifyOnDutyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/modifyOnDutyInfo"}
			// "modifyPassword":                                    nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPassword"}
			// "modifyPersonalUnitList":                            nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPersonalUnitList"}
			// "modifyPsUser":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPsUser"}
			// "moduleLogParamSet":                                 nullEndPoint.Init(apiRoot), // "/integrationService/moduleLogParamSet"}
			// "operateOssFile":                                    nullEndPoint.Init(apiRoot), // "/v1/commonService/operateOssFile"}
			// "operationPowerRobotSweepStrategy":                  nullEndPoint.Init(apiRoot), // "/v1/devService/operationPowerRobotSweepStrategy"}
			// "orgPowerReport":                                    nullEndPoint.Init(apiRoot), // "/v1/orgService/orgPowerReport"}
			// "paramSetTryAgain":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/paramSetTryAgain"}
			// "paramSetting":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/paramSetting"}
			// "planPower":                                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/planPower"}
			// "psForcastInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/psForcastInfo"}
			// "queryBatchSpeedyAddPowerStationResult":             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryBatchSpeedyAddPowerStationResult"}
			// "queryCardStatusCTCC":                               nullEndPoint.Init(apiRoot), // "/v1/devService/queryCardStatusCTCC"}
			// "queryChildAccountList":                             nullEndPoint.Init(apiRoot), // "/v1/userService/queryChildAccountList"}
			// "queryCompensationRecordData":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryCompensationRecordData"}
			// "queryCompensationRecordList":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryCompensationRecordList"}
			// "queryComponent":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/queryComponent"}
			// "queryComponentTechnicalParam":                      nullEndPoint.Init(apiRoot), // "/v1/devService/queryComponentTechnicalParam"}
			// "queryCountryGridAndRelation":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryCountryGridAndRelation"}
			// "queryCountryList":                                  nullEndPoint.Init(apiRoot), // "/v1/commonService/queryCountryList"}
			// "queryDeviceRepairList":                             nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceRepairList"}
			// "queryDeviceTypeInfoList":                           nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceTypeInfoList"}
			// "queryEnvironmentList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryEnvironmentList"}
			// "queryFaultList":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultList"}
			// "queryFaultPlanDetail":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultPlanDetail"}
			// "queryFaultRepairSteps":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultRepairSteps"}
			// "queryFaultTypeAndLevelByCode":                      nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeAndLevelByCode"}
			// "queryFaultTypeByDevice":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeByDevice"}
			// "queryFaultTypeByDevicePage":                        nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeByDevicePage"}
			// "queryFirmwareFilesPage":                            nullEndPoint.Init(apiRoot), // "/v1/commonService/queryFirmwareFilesPage"}
			// "queryInfotoAlert":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/queryInfotoAlert"}
			// "queryInverterModelList":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryInverterModelList"}
			// "queryInverterVersionList":                          nullEndPoint.Init(apiRoot), // "/v1/devService/queryInverterVersionList"}
			// "queryM2MCardInfoCMCC":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryM2MCardInfoCMCC"}
			// "queryM2MCardTermInfoCMCC":                          nullEndPoint.Init(apiRoot), // "/v1/devService/queryM2MCardTermInfoCMCC"}
			// "queryModelInfoByModelId":                           nullEndPoint.Init(apiRoot), // "/v1/devService/queryModelInfoByModelId"}
			// "queryNoticeList":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/queryNoticeList"}
			// "queryNumberOfRenewalReminders":                     nullEndPoint.Init(apiRoot), // "/v1/devService/queryNumberOfRenewalReminders"}
			// "queryOperRules":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOperRules"}
			// "queryOrderList":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOrderList"}
			// "queryOrderStep":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOrderStep"}
			// "queryOrgGenerationReport":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/queryOrgGenerationReport"}
			// "queryOrgInfoList":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/queryOrgInfoList"}
			// "queryOrgPowerElecPercent":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/queryOrgPowerElecPercent"}
			// "queryOrgPsCompensationRecordList":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryOrgPsCompensationRecordList"}
			// "queryPersonalUnitList":                             nullEndPoint.Init(apiRoot), // "/v1/userService/queryPersonalUnitList"}
			// "queryPointDataTopOne":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryPointDataTopOne"}
			// "queryPowerStationInfo":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPowerStationInfo"}
			// "queryPsAreaByUserIdAndAreaCode":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsAreaByUserIdAndAreaCode"}
			// "queryPsCompensationRecordList":                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsCompensationRecordList"}
			// "queryPsDataByDate":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsDataByDate"}
			// "queryPsIdList":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsIdList"}
			// "queryPsListByUserIdAndAreaCode":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsListByUserIdAndAreaCode"}
			// "queryPsNameByPsId":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryPsNameByPsId"}
			// "queryPsPrByDate":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsPrByDate"}
			// "queryPsProfit":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsProfit"}
			// "queryPsReportComparativeAnalysisOfPowerGeneration": nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsReportComparativeAnalysisOfPowerGeneration"}
			// "queryPsStructureList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryPsStructureList"}
			// "queryPuuidsByCommandTotalId":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryPuuidsByCommandTotalId"}
			// "queryPuuidsByCommandTotalId2":                      nullEndPoint.Init(apiRoot), // "/v1/devService/queryPuuidsByCommandTotalId2"}
			// "queryRepairRuleList":                               nullEndPoint.Init(apiRoot), // "/v1/devService/queryRepairRuleList"}
			// "queryReportListForManagementPage":                  nullEndPoint.Init(apiRoot), // "/v1/reportService/queryReportListForManagementPage"}
			// "queryReportMsg":                                    nullEndPoint.Init(apiRoot), // "/v1/reportService/queryReportMsg"}
			// "querySharingPs":                                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/querySharingPs"}
			// "querySysAdvancedParam":                             nullEndPoint.Init(apiRoot), // "/v1/devService/querySysAdvancedParam"}
			// "queryTimeBySN":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/queryTimeBySN"}
			// "queryTrafficByDateCTCC":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryTrafficByDateCTCC"}
			// "queryTrafficCTCC":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/queryTrafficCTCC"}
			// "queryUnitUuidBytotalId":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryUnitUuidBytotalId"}
			// "queryUserBtnPri":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserBtnPri"}
			// "queryUserByUserIds":                                nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserByUserIds"}
			// "queryUserExtensionAttribute":                       nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserExtensionAttribute"}
			// "queryUserForStep":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserForStep"}
			// "queryUserList":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserList"}
			// "queryUserProcessPri":                               nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserProcessPri"}
			// "queryUserWechatBindRel":                            nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserWechatBindRel"}
			// "queryUuidByTotalIdAndUuid":                         nullEndPoint.Init(apiRoot), // "/v1/devService/queryUuidByTotalIdAndUuid"}
			// "rechargeOrderSetMeal":                              nullEndPoint.Init(apiRoot), // "/v1/devService/rechargeOrderSetMeal"}
			// "renewSendReportConfirmEmail":                       nullEndPoint.Init(apiRoot), // "/v1/reportService/renewSendReportConfirmEmail"}
			// "reportList":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/reportList"}
			// "saveCustomerEmployee":                              nullEndPoint.Init(apiRoot), // "/v1/devService/saveCustomerEmployee"}
			// "saveDevSimList":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/saveDevSimList"}
			// "saveDeviceAccountBatchData":                        nullEndPoint.Init(apiRoot), // "/v1/devService/saveDeviceAccountBatchData"}
			// "saveEnviromentIncomeInfos":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/saveEnviromentIncomeInfos"}
			// "saveEnvironmentCurve":                              nullEndPoint.Init(apiRoot), // "/v1/devService/saveEnvironmentCurve"}
			// "saveFirmwareFile":                                  nullEndPoint.Init(apiRoot), // "/v1/commonService/saveFirmwareFile"}
			// "saveIncomeSettingInfos":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/saveIncomeSettingInfos"}
			// "saveOrUpdateGroupStringCheckRule":                  nullEndPoint.Init(apiRoot), // "/v1/devService/saveOrUpdateGroupStringCheckRule"}
			// "saveParamModel":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/saveParamModel"}
			// "savePowerCharges":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerCharges"}
			// "savePowerDevicePoint":                              nullEndPoint.Init(apiRoot), // "/v1/reportService/savePowerDevicePoint"}
			// "savePowerRobotInfo":                                nullEndPoint.Init(apiRoot), // "/v1/devService/savePowerRobotInfo"}
			// "savePowerRobotSweepAttr":                           nullEndPoint.Init(apiRoot), // "/v1/devService/savePowerRobotSweepAttr"}
			// "savePowerSettingCharges":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerSettingCharges"}
			// "savePowerSettingInfo":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerSettingInfo"}
			// "saveProductionBatchData":                           nullEndPoint.Init(apiRoot), // "/v1/devService/saveProductionBatchData"}
			// "saveRechargeOrderObj":                              nullEndPoint.Init(apiRoot), // "/onlinepay/saveRechargeOrderObj"}
			// "saveRechargeOrderOtherInfo":                        nullEndPoint.Init(apiRoot), // "/onlinepay/saveRechargeOrderOtherInfo"}
			// "saveRepair":                                        nullEndPoint.Init(apiRoot), // "/v1/faultService/saveRepair"}
			// "saveReportExportColumns":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/saveReportExportColumns"}
			// "saveSetParam":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/saveSetParam"}
			// "saveSysUserMsg":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/saveSysUserMsg"}
			// "saveTemplate":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/saveTemplate"}
			// "searchM2MMonthFlowCMCC":                            nullEndPoint.Init(apiRoot), // "/v1/devService/searchM2MMonthFlowCMCC"}
			// "selectSysTranslationNames":                         nullEndPoint.Init(apiRoot), // "/v1/reportService/selectSysTranslationNames"}
			// "sendPsTimeZoneInstruction":                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/sendPsTimeZoneInstruction"}
			// "setUpFormulaFaultAnalyse":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/setUpFormulaFaultAnalyse"}
			// "setUserGDPRAttrs":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/setUserGDPRAttrs"}
			// "settingNotice":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/settingNotice"}
			// "shareMyPs":                                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/shareMyPs"}
			// "sharePsBySN":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/sharePsBySN"}
			// "showInverterByUnit":                                nullEndPoint.Init(apiRoot), // "/v1/devService/showInverterByUnit"}
			// "showOnlineUsers":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/showOnlineUsers"}
			// "showWarning":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/showWarning"}
			// "snIsExist":                                         nullEndPoint.Init(apiRoot), // "/v1/devService/snIsExist"}
			// "snsIsExist":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/snsIsExist"}
			// "speedyAddPowerStation":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/speedyAddPowerStation"}
			// "stationDeviceHistoryDataList":                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationDeviceHistoryDataList"}
			// "stationUnitsList":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationUnitsList"}
			// "stationsDiscreteData":                              nullEndPoint.Init(apiRoot), // "/v1/devService/stationsDiscreteData"}
			// "stationsIncomeList":                                nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationsIncomeList"}
			// "stationsPointReport":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationsPointReport"}
			// "stationsYearPlanReport":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/stationsYearPlanReport"}
			// "sureAndImportSelettlementData":                     nullEndPoint.Init(apiRoot), // "/v1/otherService/sureAndImportSelettlementData"}
			// "sweepDevParamSet":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevParamSet"}
			// "sweepDevRunControl":                                nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevRunControl"}
			// "sweepDevStrategyIssue":                             nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevStrategyIssue"}
			// "sysTimeZoneList":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/sysTimeZoneList"}
			// "unLockUser":                                        nullEndPoint.Init(apiRoot), // "/v1/userService/unLockUser"}
			// "unlockChildAccount":                                nullEndPoint.Init(apiRoot), // "/v1/userService/unlockChildAccount"}
			// "updateCommunicationModuleState":                    nullEndPoint.Init(apiRoot), // "/v1/devService/updateCommunicationModuleState"}
			// "updateDevInstalledPower":                           nullEndPoint.Init(apiRoot), // "/v1/devService/updateDevInstalledPower"}
			// "updateFault":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultStatus"}
			// "updateFaultData":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultData"}
			// "updateFaultMsgByFaultCode":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultMsgByFaultCode"}
			// "updateFaultStatus":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultStatus"}
			// "updateHouseholdWorkOrder":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/updateHouseholdWorkOrder"}
			// "updateInverterSn2ModuleSn":                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/updateInverterSn2ModuleSn"}
			// "updateOperateTicketAttachmentId":                   nullEndPoint.Init(apiRoot), // "/v1/faultService/updateOperateTicketAttachmentId"}
			// "updateOrderDeviceByCustomerService":                nullEndPoint.Init(apiRoot), // "/onlinepay/updateOrderDeviceByCustomerService"}
			// "updateOwnerFaultConfig":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/updateOwnerFaultConfig"}
			// "updateParamSettingSysMsg":                          nullEndPoint.Init(apiRoot), // "/v1/devService/updateParamSettingSysMsg"}
			// "updatePlatformLevelFaultLevel":                     nullEndPoint.Init(apiRoot), // "/v1/faultService/updatePlatformLevelFaultLevel"}
			// "updatePowerDevicePoint":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/updatePowerDevicePoint"}
			// "updatePowerRobotInfo":                              nullEndPoint.Init(apiRoot), // "/v1/devService/updatePowerRobotInfo"}
			// "updatePowerRobotSweepAttr":                         nullEndPoint.Init(apiRoot), // "/v1/devService/updatePowerRobotSweepAttr"}
			// "updatePowerStationForHousehold":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updatePowerStationForHousehold"}
			// "updatePowerStationInfo":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updatePowerStationInfo"}
			// "updatePowerUserInfo":                               nullEndPoint.Init(apiRoot), // "/v1/userService/updatePowerUserInfo"}
			// "updateReportConfigByEmailAddr":                     nullEndPoint.Init(apiRoot), // "/v1/reportService/updateReportConfigByEmailAddr"}
			// "updateShareAttr":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updateShareAttr"}
			// "updateSnIsSureFlag":                                nullEndPoint.Init(apiRoot), // "/devDataHandleService/updateSnIsSureFlag"}
			// "updateStationPics":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updateStationPics"}
			// "updateSysAdvancedParam":                            nullEndPoint.Init(apiRoot), // "/v1/devService/updateSysAdvancedParam"}
			// "updateSysOrgNew":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/updateSysOrgNew"}
			// "updateTemplate":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/updateDataCurveTemplate"}
			// "updateUinfoNetEaseUser":                            nullEndPoint.Init(apiRoot), // "/v1/userService/updateUinfoNetEaseUser"}
			// "updateUserExtensionAttribute":                      nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserExtensionAttribute"}
			// "updateUserLanguage":                                nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserLanguage"}
			// "updateUserPosition":                                nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserPosition"}
			// "updateUserUpOrg":                                   nullEndPoint.Init(apiRoot), // "/v1/orgService/updateUserUpOrg"}
			// "upgrade":                                           nullEndPoint.Init(apiRoot), // "/v1/userService/upgrade"}
			// "upgrate":                                           nullEndPoint.Init(apiRoot), // "/v1/userService/upgrade"}
			// "uploadFileToOss":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/uploadFileToOss"}
			// "userAgreeGdprProtocol":                             nullEndPoint.Init(apiRoot), // "/v1/userService/userAgreeGdprProtocol"}
			// "userInfoUniqueCheck":                               nullEndPoint.Init(apiRoot), // "/v1/userService/userInfoUniqueCheck"}
			// "userMailHasBound":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/userMailHasBound"}
			// "userRegister":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/userRegister"}
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
