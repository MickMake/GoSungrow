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
			api.GetName(powerDevicePointList.EndPoint{}):        powerDevicePointList.Init(apiRoot),  // @TODO - CRITICAL - Get all point_id definitions.
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

			// Use the following for all critical data every 5 minutes.
			api.GetName(queryDeviceList.EndPoint{}): queryDeviceList.Init(apiRoot), // @TODO - This gives you ALL info at 5 min intervals.

			// Disabled from here on.
			api.GetName(getPsDataSupplementTaskList.EndPoint{}):     getPsDataSupplementTaskList.Init(apiRoot),
			api.GetName(getPowerDeviceSetTaskDetailList.EndPoint{}): getPowerDeviceSetTaskDetailList.Init(apiRoot),
			api.GetName(getPowerDeviceSetTaskList.EndPoint{}):       getPowerDeviceSetTaskList.Init(apiRoot),
			api.GetName(queryBatchCreatePsTaskList.EndPoint{}):      queryBatchCreatePsTaskList.Init(apiRoot),
			api.GetName(getModuleLogTaskList.EndPoint{}):            getModuleLogTaskList.Init(apiRoot),
			api.GetName(queryParamSettingTask.EndPoint{}):           queryParamSettingTask.Init(apiRoot),
			api.GetName(queryCtrlTaskById.EndPoint{}):               queryCtrlTaskById.Init(apiRoot),
			api.GetName(cancelParamSetTask.EndPoint{}):              cancelParamSetTask.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeSubTasks.EndPoint{}):   dealFailRemoteUpgradeSubTasks.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeTasks.EndPoint{}):      dealFailRemoteUpgradeTasks.Init(apiRoot),
			api.GetName(getRemoteUpgradeSubTasksList.EndPoint{}):    getRemoteUpgradeSubTasksList.Init(apiRoot),
			api.GetName(getRemoteUpgradeTaskList.EndPoint{}):        getRemoteUpgradeTaskList.Init(apiRoot),

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
