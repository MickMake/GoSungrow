package AppService

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService/getPowerDevicePointNames"
	"GoSungro/iSolarCloud/sungro/AppService/login"
	"GoSungro/iSolarCloud/sungro/AppService/nullEndPoint"
	"fmt"
)

var _ api.Area = (*Area)(nil)

type Area api.AreaStruct


func init() {
	name := api.GetArea(Area{})
	fmt.Printf("Name: %s\n", name)
}

func Init() Area {
	fmt.Println("Init()")
	name := api.GetArea(Area{})
	fmt.Printf("Name: %s\n", name)

	area := Area {
		Name:      api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints {
			"acceptPsSharing":                                   nullEndPoint.Init(), // "/v1/powerStationService/acceptPsSharing"}
			"activateEmail":                                     nullEndPoint.Init(), // "/v1/userService/activateEmail"}
			"addConfig":                                         nullEndPoint.Init(), // "/devDataHandleService/addConfig"}
			"addDeviceRepair":                                   nullEndPoint.Init(), // "/v1/devService/addDeviceRepair"}
			"addDeviceToStructureForHousehold":                  nullEndPoint.Init(), // "/devDataHandleService/addDeviceToStructureForHousehold"}
			"addDeviceToStructureForHouseholdByPsIdS":           nullEndPoint.Init(), // "/devDataHandleService/addDeviceToStructureForHouseholdByPsIdS"}
			"addFault":                                          nullEndPoint.Init(), // "/v1/faultService/addFault"}
			"addFaultOrder":                                     nullEndPoint.Init(), // "/v1/faultService/addFaultOrder"}
			"addFaultPlan":                                      nullEndPoint.Init(), // "/v1/faultService/addFaultPlan"}
			"addFaultRepairSteps":                               nullEndPoint.Init(), // "/v1/faultService/addFaultRepairSteps"}
			"addHouseholdEvaluation":                            nullEndPoint.Init(), // "/v1/faultService/addHouseholdEvaluation"}
			"addHouseholdLeaveMessage":                          nullEndPoint.Init(), // "/v1/faultService/addHouseholdLeaveMessage"}
			"addHouseholdOpinionFeedback":                       nullEndPoint.Init(), // "/v1/faultService/addHouseholdOpinionFeedback"}
			"addHouseholdWorkOrder":                             nullEndPoint.Init(), // "/v1/faultService/addHouseholdWorkOrder"}
			"addOnDutyInfo":                                     nullEndPoint.Init(), // "/v1/otherService/addOnDutyInfo"}
			"addOperRule":                                       nullEndPoint.Init(), // "/v1/faultService/addOperRule"}
			"addOrDelPsStructure":                               nullEndPoint.Init(), // "/v1/devService/addOrDelPsStructure"}
			"addOrderStep":                                      nullEndPoint.Init(), // "/v1/faultService/addOrderStep"}
			"addPowerStationForHousehold":                       nullEndPoint.Init(), // "/v1/powerStationService/addPowerStationForHousehold"}
			"addPowerStationInfo":                               nullEndPoint.Init(), // "/v1/powerStationService/addPowerStationInfo"}
			"addReportConfigEmail":                              nullEndPoint.Init(), // "/v1/reportService/addReportConfigEmail"}
			"addSysAdvancedParam":                               nullEndPoint.Init(), // "/v1/devService/addSysAdvancedParam"}
			"addSysOrgNew":                                      nullEndPoint.Init(), // "/v1/otherService/addSysOrgNew"}
			"aliPayAppTest":                                     nullEndPoint.Init(), // "/onlinepay/aliPayAppTest"}
			"auditOperRule":                                     nullEndPoint.Init(), // "/v1/faultService/auditOperRule"}
			"batchAddStationBySn":                               nullEndPoint.Init(), // "/v1/powerStationService/batchAddStationBySn"}
			"batchImportSN":                                     nullEndPoint.Init(), // "/v1/devService/batchImportSN"}
			"batchInsertUserAndOrg":                             nullEndPoint.Init(), // "/v1/userService/batchInsertUserAndOrg"}
			"batchModifyDevicesInfoAndPropertis":                nullEndPoint.Init(), // "/v1/devService/batchModifyDevicesInfoAndPropertis"}
			"batchProcessPlantReport":                           nullEndPoint.Init(), // "/v1/powerStationService/batchProcessPlantReport"}
			"batchUpdateDeviceSim":                              nullEndPoint.Init(), // "/v1/devService/batchUpdateDeviceSim"}
			"batchUpdateUserIsAgreeGdpr":                        nullEndPoint.Init(), // "/v1/userService/batchUpdateUserIsAgreeGdpr"}
			"boundMobilePhone":                                  nullEndPoint.Init(), // "/v1/userService/boundMobilePhone"}
			"boundUserMail":                                     nullEndPoint.Init(), // "/v1/userService/boundUserMail"}
			"caculateDeviceInputDiscrete":                       nullEndPoint.Init(), // "/v1/devService/caculateDeviceInputDiscrete"}
			"calculateDeviceDiscrete":                           nullEndPoint.Init(), // "/v1/devService/calculateDeviceDiscrete"}
			"calculateInitialCompensationData":                  nullEndPoint.Init(), // "/v1/powerStationService/calculateInitialCompensationData"}
			"cancelDeliverMail":                                 nullEndPoint.Init(), // "/v1/messageService/cancelDeliverMail"}
			"cancelOrderScan":                                   nullEndPoint.Init(), // "/v1/devService/cancelOrderScan"}
			"cancelParamSetTask":                                nullEndPoint.Init(), // "/v1/devService/cancelParamSetTask"}
			"cancelPsSharing":                                   nullEndPoint.Init(), // "/v1/powerStationService/cancelPsSharing"}
			"cancelRechargeOrder":                               nullEndPoint.Init(), // "/onlinepay/cancelRechargeOrder"}
			"changRechargeOrderToCancel":                        nullEndPoint.Init(), // "/onlinepay/changRechargeOrderToCancel"}
			"changeHouseholdUser2Installer":                     nullEndPoint.Init(), // "/v1/orgService/changeHouseholdUser2Installer"}
			"changeRemoteParam":                                 nullEndPoint.Init(), // "/v1/devService/changeRemoteParam"}
			"checkDealerOrgCode":                                nullEndPoint.Init(), // "/v1/orgService/checkDealerOrgCode"}
			"checkDevSnIsBelongsToUser":                         nullEndPoint.Init(), // "/v1/userService/checkDevSnIsBelongsToUser"}
			"checkInverterResult":                               nullEndPoint.Init(), // "/v1/devService/checkInverterResult"}
			"checkIsCanDoParamSet":                              nullEndPoint.Init(), // "/v1/devService/checkIsCanDoParamSet"}
			"checkIsIvScan":                                     nullEndPoint.Init(), // "/v1/devService/checkIsIvScan"}
			"checkOssObjectExist":                               nullEndPoint.Init(), // "/v1/commonService/checkOssObjectExist"}
			"checkServiceIsConnect":                             nullEndPoint.Init(), // "/v1/commonService/checkServiceIsConnect"}
			"checkTechnicalParameters":                          nullEndPoint.Init(), // "/v1/devService/checkTechnicalParameters"}
			"checkUnitStatus":                                   nullEndPoint.Init(), // "/v1/devService/checkUnitStatus"}
			"checkUpRechargeDevicePaying":                       nullEndPoint.Init(), // "/onlinepay/checkUpRechargeDevicePaying"}
			"checkUserAccountUnique":                            nullEndPoint.Init(), // "/v1/userService/checkUserAccountUnique"}
			"checkUserAccountUniqueAll":                         nullEndPoint.Init(), // "/v1/userService/checkUserAccountUniqueAll"}
			"checkUserInfoUnique":                               nullEndPoint.Init(), // "/v1/userService/checkUserInfoUnique"}
			"checkUserIsExist":                                  nullEndPoint.Init(), // "/v1/userService/checkUserIsExist"}
			"checkUserListIsExist":                              nullEndPoint.Init(), // "/v1/userService/checkUserListIsExist"}
			"checkUserPassword":                                 nullEndPoint.Init(), // "/v1/userService/checkUserPassword"}
			"cloudDeploymentRecord":                             nullEndPoint.Init(), // "/v1/commonService/cloudDeploymentRecord"}
			"comfirmParamModel":                                 nullEndPoint.Init(), // "/v1/devService/comfirmParamModel"}
			"communicationModuleDetail":                         nullEndPoint.Init(), // "/v1/devService/communicationModuleDetail"}
			"compareValidateCode":                               nullEndPoint.Init(), // "/v1/userService/compareValidateCode"}
			"componentInfo2Cloud":                               nullEndPoint.Init(), // "/v1/devService/componentInfo2Cloud"}
			"confirmFault":                                      nullEndPoint.Init(), // "/v1/faultService/confirmFault"}
			"confirmIvFault":                                    nullEndPoint.Init(), // "/v1/devService/confirmIvFault"}
			"confirmReportConfig":                               nullEndPoint.Init(), // "/v1/reportService/confirmReportConfig"}
			"createAppkeyInfo":                                  nullEndPoint.Init(), // "/v1/userService/createAppkeyInfo"}
			"createRenewInvoice":                                nullEndPoint.Init(), // "/onlinepay/createRenewInvoice"}
			"dealCommandReply":                                  nullEndPoint.Init(), // "/devDataHandleService/dealCommandReply"}
			"dealDeletePsFailPsDelete":                          nullEndPoint.Init(), // "/v1/powerStationService/dealDeletePsFailPsDelete"}
			"dealFailRemoteUpgradeSubTasks":                     nullEndPoint.Init(), // "/v1/devService/dealFailRemoteUpgradeSubTasks"}
			"dealFailRemoteUpgradeTasks":                        nullEndPoint.Init(), // "/v1/devService/dealFailRemoteUpgradeTasks"}
			"dealFaultOrder":                                    nullEndPoint.Init(), // "/v1/faultService/dealFaultOrder"}
			"dealGroupStringDisableOrEnable":                    nullEndPoint.Init(), // "/v1/devService/dealGroupStringDisableOrEnable"}
			"dealNumberOfServiceCalls2Mysql":                    nullEndPoint.Init(), // "/v1/commonService/dealNumberOfServiceCalls2Mysql"}
			"dealParamSettingAfterComplete":                     nullEndPoint.Init(), // "/v1/devService/dealParamSettingAfterComplete"}
			"dealPsDataSupplement":                              nullEndPoint.Init(), // "/v1/powerStationService/dealPsDataSupplement"}
			"dealPsReportEmailSend":                             nullEndPoint.Init(), // "/v1/reportService/dealPsReportEmailSend"}
			"dealRemoteUpgrade":                                 nullEndPoint.Init(), // "/v1/devService/dealRemoteUpgrade"}
			"dealSnElectrifyCheck":                              nullEndPoint.Init(), // "/v1/devService/dealSnElectrifyCheck"}
			"dealSysDeviceSimFlowInfo":                          nullEndPoint.Init(), // "/v1/devService/dealSysDeviceSimFlowInfo"}
			"dealSysDeviceSimInfo":                              nullEndPoint.Init(), // "/v1/devService/dealSysDeviceSimInfo"}
			"definiteTimeDealSnExpRemind":                       nullEndPoint.Init(), // "/v1/devService/definiteTimeDealSnExpRemind"}
			"definiteTimeDealSnStatus":                          nullEndPoint.Init(), // "/v1/devService/definiteTimeDealSnStatus"}
			"delDeviceRepair":                                   nullEndPoint.Init(), // "/v1/devService/delDeviceRepair"}
			"delOperRule":                                       nullEndPoint.Init(), // "/v1/faultService/delOperRule"}
			"delayCallApiResidueTimes":                          nullEndPoint.Init(), // "/v1/commonService/delayCallApiResidueTimes"}
			"deleteComponent":                                   nullEndPoint.Init(), // "/v1/devService/deleteComponent"}
			"deleteCustomerEmployee":                            nullEndPoint.Init(), // "/v1/devService/deleteCustomerEmployee"}
			"deleteDeviceAccount":                               nullEndPoint.Init(), // "/v1/devService/deleteDeviceAccount"}
			"deleteDeviceSimById":                               nullEndPoint.Init(), // "/v1/devService/deleteDeviceSimById"}
			"deleteElectricitySettlementData":                   nullEndPoint.Init(), // "/v1/otherService/deleteElectricitySettlementData"}
			"deleteFaultPlan":                                   nullEndPoint.Init(), // "/v1/faultService/deleteFaultPlan"}
			"deleteFirmwareFiles":                               nullEndPoint.Init(), // "/v1/commonService/deleteFirmwareFiles"}
			"deleteHouseholdEvaluation":                         nullEndPoint.Init(), // "/v1/faultService/deleteHouseholdEvaluation"}
			"deleteHouseholdLeaveMessage":                       nullEndPoint.Init(), // "/v1/faultService/deleteHouseholdLeaveMessage"}
			"deleteHouseholdWorkOrder":                          nullEndPoint.Init(), // "/v1/faultService/deleteHouseholdWorkOrder"}
			"deleteInverterSnInChnnl":                           nullEndPoint.Init(), // "/devDataHandleService/deleteInverterSnInChnnl"}
			"deleteModuleLog":                                   nullEndPoint.Init(), // "/integrationService/deleteModuleLog"}
			"deleteOnDutyInfo":                                  nullEndPoint.Init(), // "/v1/otherService/deleteOnDutyInfo"}
			"deleteOperateBillFile":                             nullEndPoint.Init(), // "/v1/faultService/deleteOperateBillFile"}
			"deleteOssObject":                                   nullEndPoint.Init(), // "/v1/commonService/deleteOssObject"}
			"deletePowerDevicePointById":                        nullEndPoint.Init(), // "/v1/reportService/deletePowerDevicePointById"}
			"deletePowerPicture":                                nullEndPoint.Init(), // "/v1/powerStationService/deletePowerPicture"}
			"deletePowerRobotInfoBySnAndPsId":                   nullEndPoint.Init(), // "/v1/devService/deletePowerRobotInfoBySnAndPsId"}
			"deletePowerRobotSweepStrategy":                     nullEndPoint.Init(), // "/v1/devService/deletePowerRobotSweepStrategy"}
			"deleteProductionData":                              nullEndPoint.Init(), // "/v1/devService/deleteProductionData"}
			"deletePs":                                          nullEndPoint.Init(), // "/v1/powerStationService/deletePs"}
			"deleteRechargeOrder":                               nullEndPoint.Init(), // "/onlinepay/deleteRechargeOrder"}
			"deleteRegularlyConnectionInfo":                     nullEndPoint.Init(), // "/v1/commonService/deleteRegularlyConnectionInfo"}
			"deleteReportConfigEmailAddr":                       nullEndPoint.Init(), // "/v1/reportService/deleteReportConfigEmailAddr"}
			"deleteSysAdvancedParam":                            nullEndPoint.Init(), // "/v1/devService/deleteSysAdvancedParam"}
			"deleteSysOrgNew":                                   nullEndPoint.Init(), // "/v1/otherService/deleteSysOrgNew"}
			"deleteTemplate":                                    nullEndPoint.Init(), // "/v1/devService/deleteTemplate"}
			"deleteUserInfoAllByUserId":                         nullEndPoint.Init(), // "/v1/userService/deleteUserInfoAllByUserId"}
			"deviceInputDiscreteDeleteTime":                     nullEndPoint.Init(), // "/v1/devService/deviceInputDiscreteDeleteTime"}
			"deviceInputDiscreteGetTime":                        nullEndPoint.Init(), // "/v1/devService/deviceInputDiscreteGetTime"}
			"deviceInputDiscreteInsertTime":                     nullEndPoint.Init(), // "/v1/devService/deviceInputDiscreteInsertTime"}
			"deviceInputDiscreteUpdateTime":                     nullEndPoint.Init(), // "/v1/devService/deviceInputDiscreteUpdateTime"}
			"devicePointsDataFromMySql":                         nullEndPoint.Init(), // "/v1/devService/devicePointsDataFromMySql"}
			"deviceReplace":                                     nullEndPoint.Init(), // "/v1/devService/deviceReplace"}
			"editDeviceRepair":                                  nullEndPoint.Init(), // "/v1/devService/editDeviceRepair"}
			"editOperRule":                                      nullEndPoint.Init(), // "/v1/faultService/editOperRule"}
			"energyPovertyAlleviation":                          nullEndPoint.Init(), // "/v1/orgService/energyPovertyAlleviation"}
			"energyTrend":                                       nullEndPoint.Init(), // "/v1/powerStationService/energyTrend"}
			"exportParamSettingValPDF":                          nullEndPoint.Init(), // "/v1/devService/exportParamSettingValPDF"}
			"exportPlantReportPDF":                              nullEndPoint.Init(), // "/v1/powerStationService/exportPlantReportPDF"}
			"faultAutoClose":                                    nullEndPoint.Init(), // "/v1/faultService/faultAutoClose"}
			"faultCloseRemindOrderHandler":                      nullEndPoint.Init(), // "/v1/faultService/faultCloseRemindOrderHandler"}
			"findCodeValueList":                                 nullEndPoint.Init(), // "/v1/commonService/findCodeValueList"}
			"findEmgOrgInfo":                                    nullEndPoint.Init(), // "/v1/otherService/findEmgOrgInfo"}
			"findEnvironmentInfo":                               nullEndPoint.Init(), // "/v1/devService/findEnvironmentInfo"}
			"findFromHbaseAndRedis":                             nullEndPoint.Init(), // "/v1/commonService/findFromHbaseAndRedis"}
			"findInfoByuuid":                                    nullEndPoint.Init(), // "/v1/devService/findInfoByuuid"}
			"findLossAnalysisList":                              nullEndPoint.Init(), // "/v1/powerStationService/findLossAnalysisList"}
			"findOnDutyInfo":                                    nullEndPoint.Init(), // "/v1/otherService/findOnDutyInfo"}
			"findPsType":                                        nullEndPoint.Init(), // "/v1/powerStationService/findPsType"}
			"findSingleStationPR":                               nullEndPoint.Init(), // "/v1/powerStationService/findSingleStationPR"}
			"findUserPassword":                                  nullEndPoint.Init(), // "/v1/devService/findUserPassword"}
			"genTLSUserSigByUserAccount":                        nullEndPoint.Init(), // "/v1/userService/genTLSUserSigByUserAccount"}
			"generateRandomPassword":                            nullEndPoint.Init(), // "/v1/userService/generateRandomPassword"}
			"getAPIServiceInfo":                                 nullEndPoint.Init(), // "/v1/commonService/getAPIServiceInfo"}
			"getAccessedPermission":                             nullEndPoint.Init(), // "/v1/commonService/getAccessedPermission"}
			"getAllDeviceByPsId":                                nullEndPoint.Init(), // "/v1/devService/getAllDeviceByPsId"}
			"getAllPowerDeviceSetName":                          nullEndPoint.Init(), // "/v1/devService/getAllPowerDeviceSetName"}
			"getAllPowerRobotViewInfoByPsId":                    nullEndPoint.Init(), // "/v1/devService/getAllPowerRobotViewInfoByPsId"}
			"getAllPsIdByOrgIds":                                nullEndPoint.Init(), // "/v1/devService/getAllPsIdByOrgIds"}
			"getAllUserRemindCount":                             nullEndPoint.Init(), // "/v1/devService/getAllUserRemindCount"}
			"getAndOutletsAndUnit":                              nullEndPoint.Init(), // "/v1/devService/getAndOutletsAndUnit"}
			"getApiCallsForAppkeys":                             nullEndPoint.Init(), // "/v1/commonService/getApiCallsForAppkeys"}
			"getAreaInfoCodeByCounty":                           nullEndPoint.Init(), // "/v1/commonService/getAreaInfoCodeByCounty"}
			"getAreaList":                                       nullEndPoint.Init(), // "/v1/powerStationService/getAreaList"}
			"getAutoCreatePowerStation":                         nullEndPoint.Init(), // "/v1/powerStationService/getAutoCreatePowerStation"}
			"getBackReadValue":                                  nullEndPoint.Init(), // "/v1/devService/getBackReadValue"}
			"getBatchNewestPointData":                           nullEndPoint.Init(), // "/v1/devService/getBatchNewestPointData"}
			"getCallApiResidueTimes":                            nullEndPoint.Init(), // "/v1/commonService/getCallApiResidueTimes"}
			"getChangedPsListByTime":                            nullEndPoint.Init(), // "/v1/powerStationService/getChangedPsListByTime"}
			"getChnnlListByPsId":                                nullEndPoint.Init(), // "/v1/devService/getChnnlListByPsId"}
			"getCloudList":                                      nullEndPoint.Init(), // "/v1/commonService/getCloudList"}
			"getCloudServiceMappingConfig":                      nullEndPoint.Init(), // "/v1/commonService/getCloudServiceMappingConfig"}
			"getCommunicationDeviceConfigInfo":                  nullEndPoint.Init(), // "/v1/devService/getCommunicationDeviceConfigInfo"}
			"getCommunicationModuleMonitorData":                 nullEndPoint.Init(), // "/v1/devService/getCommunicationModuleMonitorData"}
			"getComponentModelFactory":                          nullEndPoint.Init(), // "/v1/devService/getComponentModelFactory"}
			"getConfigList":                                     nullEndPoint.Init(), // "/devDataHandleService/getConfigList"}
			"getConnectionInfoBySnAndLocalPort":                 nullEndPoint.Init(), // "/v1/commonService/getConnectionInfoBySnAndLocalPort"}
			"getCountDown":                                      nullEndPoint.Init(), // "/v1/devService/getCountDown"}
			"getCountryServiceInfo":                             nullEndPoint.Init(), // "/v1/commonService/getCountryServiceInfo"}
			"getCounty":                                         nullEndPoint.Init(), // "/v1/commonService/getCounty"}
			"getCustomerEmployee":                               nullEndPoint.Init(), // "/v1/devService/getCustomerEmployee"}
			"getCustomerList":                                   nullEndPoint.Init(), // "/v1/devService/getCustomerList"}
			"getDataFromHBase":                                  nullEndPoint.Init(), // "/v1/commonService/getDataFromHBase"}
			"getDataFromHbaseByRowKey":                          nullEndPoint.Init(), // "/v1/commonService/getDataFromHbaseByRowKey"}
			"getDevInstalledPowerByPsId":                        nullEndPoint.Init(), // "/v1/devService/getDevInstalledPowerByPsId"}
			"getDevRecord":                                      nullEndPoint.Init(), // "/v1/devService/getDevRecord"}
			"getDevRunRecordList":                               nullEndPoint.Init(), // "/v1/devService/getDevRunRecordList"}
			"getDevSimByList":                                   nullEndPoint.Init(), // "/v1/devService/getDevSimByList"}
			"getDevSimList":                                     nullEndPoint.Init(), // "/v1/devService/getDevSimList"}
			"getDeviceAccountById":                              nullEndPoint.Init(), // "/v1/devService/getDeviceAccountById"}
			"getDeviceFaultStatisticsData":                      nullEndPoint.Init(), // "/v1/devService/getDeviceFaultStatisticsData"}
			"getDeviceInfo":                                     nullEndPoint.Init(), // "/v1/devService/getDeviceInfo"}
			"getDeviceList":                                     nullEndPoint.Init(), // "/v1/devService/getDeviceList"}
			"getDeviceModelInfoList":                            nullEndPoint.Init(), // "/v1/devService/getDeviceModelInfoList"}
			"getDevicePointMinuteDataList":                      nullEndPoint.Init(), // "/v1/commonService/getDevicePointMinuteDataList"}
			"getDevicePoints":                                   nullEndPoint.Init(), // "/v1/devService/getDevicePoints"}
			"getDevicePropertys":                                nullEndPoint.Init(), // "/v1/devService/getDevicePropertys"}
			"getDeviceRepairDetail":                             nullEndPoint.Init(), // "/v1/devService/getDeviceRepairDetail"}
			"getDeviceTechBranchCount":                          nullEndPoint.Init(), // "/v1/devService/getDeviceTechBranchCount"}
			"getDeviceTypeInfoList":                             nullEndPoint.Init(), // "/v1/devService/getDeviceTypeInfoList"}
			"getDeviceTypeList":                                 nullEndPoint.Init(), // "/v1/devService/getDeviceTypeList"}
			"getDstInfo":                                        nullEndPoint.Init(), // "/v1/userService/getDstInfo"}
			"getElectricitySettlementData":                      nullEndPoint.Init(), // "/v1/otherService/getElectricitySettlementData"}
			"getElectricitySettlementDetailData":                nullEndPoint.Init(), // "/v1/otherService/getElectricitySettlementDetailData"}
			"getEncryptPublicKey":                               nullEndPoint.Init(), // "/v1/commonService/getEncryptPublicKey"}
			"getFaultCount":                                     nullEndPoint.Init(), // "/v1/faultService/getFaultCount"}
			"getFaultDetail":                                    nullEndPoint.Init(), // "/v1/faultService/getFaultDetail"}
			"getFaultMsgByFaultCode":                            nullEndPoint.Init(), // "/v1/faultService/getFaultMsgByFaultCode"}
			"getFaultMsgListByPageWithYYYYMM":                   nullEndPoint.Init(), // "/v1/faultService/getFaultMsgListByPageWithYYYYMM"}
			"getFaultMsgListWithYYYYMM":                         nullEndPoint.Init(), // "/v1/faultService/getFaultMsgListWithYYYYMM"}
			"getFaultPlanList":                                  nullEndPoint.Init(), // "/v1/faultService/getFaultPlanList"}
			"getFileOperationRecordOne":                         nullEndPoint.Init(), // "/v1/commonService/getFileOperationRecordOne"}
			"getFormulaFaultAnalyseList":                        nullEndPoint.Init(), // "/v1/powerStationService/getFormulaFaultAnalyseList"}
			"getGroupStringCheckResult":                         nullEndPoint.Init(), // "/v1/devService/getGroupStringCheckResult"}
			"getGroupStringCheckRule":                           nullEndPoint.Init(), // "/v1/devService/getGroupStringCheckRule"}
			"getHisData":                                        nullEndPoint.Init(), // "/v1/devService/getHisData"}
			"getHistoryInfo":                                    nullEndPoint.Init(), // "/v1/powerStationService/getHistoryInfo"}
			"getHouseholdEvaluation":                            nullEndPoint.Init(), // "/v1/faultService/getHouseholdEvaluation"}
			"getHouseholdLeaveMessage":                          nullEndPoint.Init(), // "/v1/faultService/getHouseholdLeaveMessage"}
			"getHouseholdOpinionFeedback":                       nullEndPoint.Init(), // "/v1/faultService/getHouseholdOpinionFeedback"}
			"getHouseholdPsInstallerByUserId":                   nullEndPoint.Init(), // "/v1/userService/getHouseholdPsInstallerByUserId"}
			"getHouseholdStoragePsReport":                       nullEndPoint.Init(), // "/v1/powerStationService/getHouseholdStoragePsReport"}
			"getHouseholdUserInfo":                              nullEndPoint.Init(), // "/v1/userService/getHouseholdUserInfo"}
			"getHouseholdWorkOrderInfo":                         nullEndPoint.Init(), // "/v1/faultService/getHouseholdWorkOrderInfo"}
			"getHouseholdWorkOrderList":                         nullEndPoint.Init(), // "/v1/faultService/getHouseholdWorkOrderList"}
			"getI18nConfigByType":                               nullEndPoint.Init(), // "/integrationService/i18nfile/getI18nConfigByType"}
			"getI18nFileInfo":                                   nullEndPoint.Init(), // "/integrationService/i18nfile/getI18nFileInfo"}
			"getI18nInfoByKey":                                  nullEndPoint.Init(), // "/integrationService/i18nfile/getI18nInfoByKey"}
			"getI18nVersion":                                    nullEndPoint.Init(), // "/integrationService/international/getI18nVersion"}
			"getIncomeSettingInfos":                             nullEndPoint.Init(), // "/v1/powerStationService/getIncomeSettingInfos"}
			"getInfoFromAMap":                                   nullEndPoint.Init(), // "/v1/commonService/getInfoFromAMap"}
			"getInfomationFromRedis":                            nullEndPoint.Init(), // "/v1/devService/getInfomationFromRedis"}
			"getInstallInfoList":                                nullEndPoint.Init(), // "/v1/orgService/getInstallInfoList"}
			"getInstallerInfoByDealerOrgCodeOrId":               nullEndPoint.Init(), // "/v1/orgService/getInstallerInfoByDealerOrgCodeOrId"}
			"getInvertDataList":                                 nullEndPoint.Init(), // "/v1/devService/getInvertDataList"}
			"getInverterDataCount":                              nullEndPoint.Init(), // "/v1/devService/getInverterDataCount"}
			"getInverterProcess":                                nullEndPoint.Init(), // "/v1/devService/getInverterProcess"}
			"getInverterUuidBytotalId":                          nullEndPoint.Init(), // "/v1/devService/getInverterUuidBytotalId"}
			"getIvEchartsData":                                  nullEndPoint.Init(), // "/v1/devService/getIvEchartsData"}
			"getIvEchartsDataById":                              nullEndPoint.Init(), // "/v1/devService/getIvEchartsDataById"}
			"getKpiInfo":                                        nullEndPoint.Init(), // "/v1/powerStationService/getKpiInfo"}
			"getListMiFromHBase":                                nullEndPoint.Init(), // "/v1/commonService/getListMiFromHBase"}
			"getMapInfo":                                        nullEndPoint.Init(), // "/v1/powerStationService/getMapInfo"}
			"getMapMiFromHBase":                                 nullEndPoint.Init(), // "/v1/commonService/getMapMiFromHBase"}
			"getMenuAndPrivileges":                              nullEndPoint.Init(), // "/v1/userService/getMenuAndPrivileges"}
			"getMicrogridEStoragePsReport":                      nullEndPoint.Init(), // "/v1/powerStationService/getMicrogridEStoragePsReport"}
			"getModuleLogInfo":                                  nullEndPoint.Init(), // "/integrationService/getModuleLogInfo"}
			"getModuleLogTaskList":                              nullEndPoint.Init(), // "/integrationService/getModuleLogTaskList"}
			"getNationProvJSON":                                 nullEndPoint.Init(), // "/v1/commonService/getNationProvJSON"}
			"getNeedOpAsynOpRecordList":                         nullEndPoint.Init(), // "/v1/commonService/getNeedOpAsynOpRecordList"}
			"getNoticeInfo":                                     nullEndPoint.Init(), // "/v1/otherService/getNoticeInfo"}
			"getNumberOfServiceCalls":                           nullEndPoint.Init(), // "/v1/commonService/getNumberOfServiceCalls"}
			"getOSSConfig":                                      nullEndPoint.Init(), // "/v1/commonService/getOSSConfig"}
			"getOperRuleDetail":                                 nullEndPoint.Init(), // "/v1/faultService/getOperRuleDetail"}
			"getOperateBillFileId":                              nullEndPoint.Init(), // "/v1/faultService/getOperateBillFileId"}
			"getOperateTicketForDetail":                         nullEndPoint.Init(), // "/v1/faultService/getOperateTicketForDetail"}
			"getOrCreateNetEaseUserToken":                       nullEndPoint.Init(), // "/v1/userService/getOrCreateNetEaseUserToken"}
			"getOrderDataList":                                  nullEndPoint.Init(), // "/v1/faultService/getOrderDataList"}
			"getOrderDataSql2":                                  nullEndPoint.Init(), // "/v1/devService/getOrderDataSql2"}
			"getOrderDatas":                                     nullEndPoint.Init(), // "/v1/devService/getOrderDatas"}
			"getOrderDetail":                                    nullEndPoint.Init(), // "/v1/faultService/getOrderDetail"}
			"getOrderStatistics":                                nullEndPoint.Init(), // "/v1/faultService/getOrderStatistics"}
			"getOrgIdNameByUserId":                              nullEndPoint.Init(), // "/v1/orgService/getOrgIdNameByUserId"}
			"getOrgInfoByDealerOrgCode":                         nullEndPoint.Init(), // "/v1/orgService/getOrgInfoByDealerOrgCode"}
			"getOrgListByName":                                  nullEndPoint.Init(), // "/v1/orgService/getOrgListByName"}
			"getOrgListByUserId":                                nullEndPoint.Init(), // "/v1/userService/getOrgListByUserId"}
			"getOrgListForUser":                                 nullEndPoint.Init(), // "/v1/orgService/getOrgListForUser"}
			"getOssObjectStream":                                nullEndPoint.Init(), // "/v1/commonService/getOssObjectStream"}
			"getOwnerFaultConfigList":                           nullEndPoint.Init(), // "/v1/faultService/getOwnerFaultConfigList"}
			"getPListinfoFromMysql":                             nullEndPoint.Init(), // "/v1/powerStationService/getPListinfoFromMysql"}
			"getParamSetTemplate4NewProtocol":                   nullEndPoint.Init(), // "/v1/devService/getParamSetTemplate4NewProtocol"}
			"getParamSetTemplatePointInfo":                      nullEndPoint.Init(), // "/v1/devService/getParamSetTemplatePointInfo"}
			"getParamterSettingBase":                            nullEndPoint.Init(), // "/v1/devService/getParamterSettingBase"}
			"getPhotoInfo":                                      nullEndPoint.Init(), // "/v1/otherService/getPhotoInfo"}
			"getPlanedOrNotPsList":                              nullEndPoint.Init(), // "/v1/faultService/getPlanedOrNotPsList"}
			"getPlantReportPDFList":                             nullEndPoint.Init(), // "/v1/powerStationService/getPlantReportPDFList"}
			"getPowerChargeSettingInfo":                         nullEndPoint.Init(), // "/v1/powerStationService/getPowerChargeSettingInfo"}
			"getPowerDeviceModelTechList":                       nullEndPoint.Init(), // "/v1/devService/getPowerDeviceModelTechList"}
			"getPowerDeviceModelTree":                           nullEndPoint.Init(), // "/v1/devService/getPowerDeviceModelTree"}
			"getPowerDevicePointInfo":                           nullEndPoint.Init(), // "/v1/reportService/getPowerDevicePointInfo"}
			api.GetName(getPowerDevicePointNames.EndPoint{}): getPowerDevicePointNames.Init(),
			"getPowerDeviceSetTaskDetailList":                   nullEndPoint.Init(), // "/v1/devService/getPowerDeviceSetTaskDetailList"}
			"getPowerDeviceSetTaskList":                         nullEndPoint.Init(), // "/v1/devService/getPowerDeviceSetTaskList"}
			"getPowerFormulaFaultAnalyse":                       nullEndPoint.Init(), // "/v1/powerStationService/getPowerFormulaFaultAnalyse"}
			"getPowerPictureList":                               nullEndPoint.Init(), // "/v1/powerStationService/getPowerPictureList"}
			"getPowerRobotInfoByRobotSn":                        nullEndPoint.Init(), // "/v1/devService/getPowerRobotInfoByRobotSn"}
			"getPowerRobotSweepAttrByPsId":                      nullEndPoint.Init(), // "/v1/devService/getPowerRobotSweepAttrByPsId"}
			"getPowerRobotSweepStrategy":                        nullEndPoint.Init(), // "/v1/devService/getPowerRobotSweepStrategy"}
			"getPowerRobotSweepStrategyList":                    nullEndPoint.Init(), // "/v1/devService/getPowerRobotSweepStrategyList"}
			"getPowerSettingCharges":                            nullEndPoint.Init(), // "/v1/powerStationService/getPowerSettingCharges"}
			"getPowerSettingHistoryRecords":                     nullEndPoint.Init(), // "/v1/powerStationService/getPowerSettingHistoryRecords"}
			"getPowerStationBasicInfo":                          nullEndPoint.Init(), // "/v1/powerStationService/getPowerStationBasicInfo"}
			"getPowerStationData":                               nullEndPoint.Init(), // "/v1/powerStationService/getPowerStationData"}
			"getPowerStationForHousehold":                       nullEndPoint.Init(), // "/v1/powerStationService/getPowerStationForHousehold"}
			"getPowerStationInfo":                               nullEndPoint.Init(), // "/v1/powerStationService/getPowerStationInfo"}
			"getPowerStationPR":                                 nullEndPoint.Init(), // "/v1/powerStationService/getPowerStationPR"}
			"getPowerStationTableDataSql":                       nullEndPoint.Init(), // "/v1/devService/getPowerStationTableDataSql"}
			"getPowerStationTableDataSqlCount":                  nullEndPoint.Init(), // "/v1/devService/getPowerStationTableDataSqlCount"}
			"getPowerStatistics":                                nullEndPoint.Init(), // "/v1/powerStationService/getPowerStatistics"}
			"getPowerTrendDayData":                              nullEndPoint.Init(), // "/v1/powerStationService/getPowerTrendDayData"}
			"getPrivateCloudValidityPeriod":                     nullEndPoint.Init(), // "/v1/commonService/getPrivateCloudValidityPeriod"}
			"getProvInfoListByNationCode":                       nullEndPoint.Init(), // "/v1/commonService/getProvInfoListByNationCode"}
			"getPsAuthKey":                                      nullEndPoint.Init(), // "/v1/powerStationService/getPsAuthKey"}
			"getPsCurveInfo":                                    nullEndPoint.Init(), // "/v1/devService/getPsCurveInfo"}
			"getPsDataSupplementTaskList":                       nullEndPoint.Init(), // "/v1/powerStationService/getPsDataSupplementTaskList"}
			"getPsDetail":                                       nullEndPoint.Init(), // "/v1/powerStationService/getPsDetail"}
			"getPsDetailByUserTokens":                           nullEndPoint.Init(), // "/v1/powerStationService/getPsDetailByUserTokens"}
			"getPsDetailForSinglePage":                          nullEndPoint.Init(), // "/v1/powerStationService/getPsDetailForSinglePage"}
			"getPsDetailWithPsType":                             nullEndPoint.Init(), // "/v1/powerStationService/getPsDetailWithPsType"}
			"getPsHealthState":                                  nullEndPoint.Init(), // "/v1/powerStationService/getPsHealthState"}
			"getPsInstallerByPsId":                              nullEndPoint.Init(), // "/v1/orgService/getPsInstallerByPsId"}
			"getPsInstallerOrgInfoByPsId":                       nullEndPoint.Init(), // "/v1/powerStationService/getPsInstallerOrgInfoByPsId"}
			"getPsList":                                         nullEndPoint.Init(), // "/v1/powerStationService/getPsList"}
			"getPsListByName":                                   nullEndPoint.Init(), // "/v1/powerStationService/getPsListByName"}
			"getPsListForPsDataByPsId":                          nullEndPoint.Init(), // "/v1/powerStationService/getPsListForPsDataByPsId"}
			"getPsListStaticData":                               nullEndPoint.Init(), // "/v1/powerStationService/getPsListStaticData"}
			"getPsReport":                                       nullEndPoint.Init(), // "/v1/reportService/getPsReport"}
			"getPsUser":                                         nullEndPoint.Init(), // "/v1/userService/getPsUser"}
			"getPsWeatherList":                                  nullEndPoint.Init(), // "/v1/powerStationService/getPsWeatherList"}
			"getRechargeOrderDetail":                            nullEndPoint.Init(), // "/onlinepay/getRechargeOrderDetail"}
			"getRechargeOrderItemDeviceList":                    nullEndPoint.Init(), // "/onlinepay/getRechargeOrderItemDeviceList"}
			"getRechargeOrderList":                              nullEndPoint.Init(), // "/onlinepay/getRechargeOrderList"}
			"getRegionalTree":                                   nullEndPoint.Init(), // "/v1/orgService/getRegionalTree"}
			"getRemoteParamSettingList":                         nullEndPoint.Init(), // "/v1/devService/getRemoteParamSettingList"}
			"getRemoteUpgradeDeviceList":                        nullEndPoint.Init(), // "/v1/devService/getRemoteUpgradeDeviceList"}
			"getRemoteUpgradeScheduleDetails":                   nullEndPoint.Init(), // "/v1/devService/getRemoteUpgradeScheduleDetails"}
			"getRemoteUpgradeSubTasksList":                      nullEndPoint.Init(), // "/v1/devService/getRemoteUpgradeSubTasksList"}
			"getRemoteUpgradeTaskList":                          nullEndPoint.Init(), // "/v1/devService/getRemoteUpgradeTaskList"}
			"getReportData":                                     nullEndPoint.Init(), // "/v1/powerStationService/getReportData"}
			"getReportEmailConfigInfo":                          nullEndPoint.Init(), // "/v1/reportService/getReportEmailConfigInfo"}
			"getReportExportColumns":                            nullEndPoint.Init(), // "/v1/reportService/getReportExportColumns"}
			"getReportListByUserId":                             nullEndPoint.Init(), // "/v1/reportService/getReportListByUserId"}
			"getRobotDynamicCleaningView":                       nullEndPoint.Init(), // "/v1/devService/getRobotDynamicCleaningView"}
			"getRobotNumAndSweepCapacity":                       nullEndPoint.Init(), // "/v1/devService/getRobotNumAndSweepCapacity"}
			"getRuleUnit":                                       nullEndPoint.Init(), // "/v1/powerStationService/getRuleUnit"}
			"getSendReportConfigCron":                           nullEndPoint.Init(), // "/v1/reportService/getSendReportConfigCron"}
			"getSerialNum":                                      nullEndPoint.Init(), // "/v1/devService/getSerialNum"}
			"getShieldMapConditionList":                         nullEndPoint.Init(), // "/v1/commonService/getShieldMapConditionList"}
			"getSimIdBySnList":                                  nullEndPoint.Init(), // "/v1/devService/getSimIdBySnList"}
			"getSingleIVData":                                   nullEndPoint.Init(), // "/v1/devService/getSingleIVData"}
			"getSnChangeRecord":                                 nullEndPoint.Init(), // "/v1/devService/getSnChangeRecord"}
			"getSnConnectionInfo":                               nullEndPoint.Init(), // "/v1/commonService/getSnConnectionInfo"}
			"getStationInfoSql":                                 nullEndPoint.Init(), // "/v1/devService/getStationInfoSql"}
			"getSungwsConfigCache":                              nullEndPoint.Init(), // "/v1/commonService/getSungwsConfigCache"}
			"getSungwsGlobalConfigCache":                        nullEndPoint.Init(), // "/v1/commonService/getSungwsGlobalConfigCache"}
			"getSweepDevParamSetTemplate":                       nullEndPoint.Init(), // "/v1/devService/getSweepDevParamSetTemplate"}
			"getSweepRobotDevList":                              nullEndPoint.Init(), // "/v1/devService/getSweepRobotDevList"}
			"getSysMsg":                                         nullEndPoint.Init(), // "/v1/otherService/getSysMsg"}
			"getSysOrgNewList":                                  nullEndPoint.Init(), // "/v1/otherService/getSysOrgNewList"}
			"getSysOrgNewOne":                                   nullEndPoint.Init(), // "/v1/otherService/getSysOrgNewOne"}
			"getSysUserById":                                    nullEndPoint.Init(), // "/v1/userService/getSysUserById"}
			"getTableDataSql":                                   nullEndPoint.Init(), // "/v1/devService/getTableDataSql"}
			"getTableDataSqlCount":                              nullEndPoint.Init(), // "/v1/devService/getTableDataSqlCount"}
			"getTemplateByInfoType":                             nullEndPoint.Init(), // "/v1/messageService/getTemplateByInfoType"}
			"getTemplateList":                                   nullEndPoint.Init(), // "/v1/devService/getTemplateList"}
			"getUUIDByUpuuid":                                   nullEndPoint.Init(), // "/v1/devService/getUUIDByUpuuid"}
			"getUpTimePoint":                                    nullEndPoint.Init(), // "/v1/devService/getUpTimePoint"}
			"getUserById":                                       nullEndPoint.Init(), // "/v1/userService/getUserById"}
			"getUserByInstaller":                                nullEndPoint.Init(), // "/v1/userService/getUserByInstaller"}
			"getUserDevOnlineOffineCount":                       nullEndPoint.Init(), // "/v1/devService/getUserDevOnlineOffineCount"}
			"getUserGDPRAttrs":                                  nullEndPoint.Init(), // "/v1/userService/getUserGDPRAttrs"}
			"getUserHavePowerStationCount":                      nullEndPoint.Init(), // "/v1/userService/getUserHavePowerStationCount"}
			"getUserInfoByUserAccounts":                         nullEndPoint.Init(), // "/v1/userService/getUserInfoByUserAccounts"}
			"getUserList":                                       nullEndPoint.Init(), // "/v1/userService/getUserList"}
			"getUserPsOrderList":                                nullEndPoint.Init(), // "/v1/faultService/getUserPsOrderList"}
			"getValFromHBase":                                   nullEndPoint.Init(), // "/v1/commonService/getValFromHBase"}
			"getValidateCode":                                   nullEndPoint.Init(), // "/v1/userService/getValidateCode"}
			"getValidateCodeAtRegister":                         nullEndPoint.Init(), // "/v1/userService/getValidateCodeAtRegister"}
			"getWeatherInfo":                                    nullEndPoint.Init(), // "/v1/powerStationService/getWeatherInfo"}
			"getWechatPushConfig":                               nullEndPoint.Init(), // "/v1/userService/getWechatPushConfig"}
			"getWorkInfo":                                       nullEndPoint.Init(), // "/v1/otherService/getWorkInfo"}
			"groupStringCheck":                                  nullEndPoint.Init(), // "/v1/devService/groupStringCheck"}
			"handleDevByCommunicationSN":                        nullEndPoint.Init(), // "/devDataHandleService/handleDevByCommunicationSN"}
			"householdResetPassBySN":                            nullEndPoint.Init(), // "/v1/userService/householdResetPassBySN"}
			"immediatePayment":                                  nullEndPoint.Init(), // "/onlinepay/immediatePayment"}
			"importExcelData":                                   nullEndPoint.Init(), // "/v1/devService/importExcelData"}
			"incomeStatistics":                                  nullEndPoint.Init(), // "/v1/powerStationService/incomeStatistics"}
			"informPush":                                        nullEndPoint.Init(), // "/v1/messageService/informPush"}
			"insertEmgOrgInfo":                                  nullEndPoint.Init(), // "/v1/otherService/insertEmgOrgInfo"}
			"insightSynDeviceStructure2Cloud":                   nullEndPoint.Init(), // "/v1/powerStationService/insightSynDeviceStructure2Cloud"}
			"intoDataToHbase":                                   nullEndPoint.Init(), // "/v1/commonService/intoDataToHbase"}
			"ipLocationQuery":                                   nullEndPoint.Init(), // "/v1/commonService/ipLocationQuery"}
			"isHave2GSn":                                        nullEndPoint.Init(), // "/v1/devService/isHave2GSn"}
			"judgeDevIsHasInitSetTemplate":                      nullEndPoint.Init(), // "/v1/devService/judgeDevIsHasInitSetTemplate"}
			"judgeIsSettingMan":                                 nullEndPoint.Init(), // "/v1/faultService/judgeIsSettingMan"}
			"listOssFiles":                                      nullEndPoint.Init(), // "/v1/commonService/listOssFiles"}
			"loadAreaInfo":                                      nullEndPoint.Init(), // "/v1/commonService/loadAreaInfo"}
			"loadPowerStation":                                  nullEndPoint.Init(), // "/v1/powerStationService/loadPowerStation"}
			api.GetName(login.EndPoint{}): login.Init(),
			"loginByToken":                                      nullEndPoint.Init(), // "/v1/userService/loginByToken"}
			"logout":                                            nullEndPoint.Init(), // "/v1/userService/logout"}
			"mobilePhoneHasBound":                               nullEndPoint.Init(), // "/v1/userService/mobilePhoneHasBound"}
			"modifiedDeviceInfo":                                nullEndPoint.Init(), // "/v1/devService/modifiedDeviceInfo"}
			"modifyEmgOrgStruc":                                 nullEndPoint.Init(), // "/v1/otherService/modifyEmgOrgStruc"}
			"modifyFaultPlan":                                   nullEndPoint.Init(), // "/v1/faultService/modifyFaultPlan"}
			"modifyOnDutyInfo":                                  nullEndPoint.Init(), // "/v1/otherService/modifyOnDutyInfo"}
			"modifyPassword":                                    nullEndPoint.Init(), // "/v1/userService/modifyPassword"}
			"modifyPersonalUnitList":                            nullEndPoint.Init(), // "/v1/userService/modifyPersonalUnitList"}
			"modifyPsUser":                                      nullEndPoint.Init(), // "/v1/userService/modifyPsUser"}
			"moduleLogParamSet":                                 nullEndPoint.Init(), // "/integrationService/moduleLogParamSet"}
			"operateOssFile":                                    nullEndPoint.Init(), // "/v1/commonService/operateOssFile"}
			"operationPowerRobotSweepStrategy":                  nullEndPoint.Init(), // "/v1/devService/operationPowerRobotSweepStrategy"}
			"orgPowerReport":                                    nullEndPoint.Init(), // "/v1/orgService/orgPowerReport"}
			"paramSetTryAgain":                                  nullEndPoint.Init(), // "/v1/devService/paramSetTryAgain"}
			"paramSetting":                                      nullEndPoint.Init(), // "/v1/devService/paramSetting"}
			"planPower":                                         nullEndPoint.Init(), // "/v1/powerStationService/planPower"}
			"powerDevicePointList":                              nullEndPoint.Init(), // "/v1/reportService/powerDevicePointList"}
			"powerTrendChartData":                               nullEndPoint.Init(), // "/v1/powerStationService/powerTrendChartData"}
			"psForcastInfo":                                     nullEndPoint.Init(), // "/v1/powerStationService/psForcastInfo"}
			"psHourPointsValue":                                 nullEndPoint.Init(), // "/v1/powerStationService/psHourPointsValue"}
			"queryAllPsIdAndName":                               nullEndPoint.Init(), // "/v1/powerStationService/queryAllPsIdAndName"}
			"queryBatchCreatePsTaskList":                        nullEndPoint.Init(), // "/v1/powerStationService/queryBatchCreatePsTaskList"}
			"queryBatchSpeedyAddPowerStationResult":             nullEndPoint.Init(), // "/v1/powerStationService/queryBatchSpeedyAddPowerStationResult"}
			"queryCardStatusCTCC":                               nullEndPoint.Init(), // "/v1/devService/queryCardStatusCTCC"}
			"queryChildAccountList":                             nullEndPoint.Init(), // "/v1/userService/queryChildAccountList"}
			"queryCompensationRecordData":                       nullEndPoint.Init(), // "/v1/powerStationService/queryCompensationRecordData"}
			"queryCompensationRecordList":                       nullEndPoint.Init(), // "/v1/powerStationService/queryCompensationRecordList"}
			"queryComponent":                                    nullEndPoint.Init(), // "/v1/devService/queryComponent"}
			"queryComponentTechnicalParam":                      nullEndPoint.Init(), // "/v1/devService/queryComponentTechnicalParam"}
			"queryCountryGridAndRelation":                       nullEndPoint.Init(), // "/v1/devService/queryCountryGridAndRelation"}
			"queryCountryList":                                  nullEndPoint.Init(), // "/v1/commonService/queryCountryList"}
			"queryCtrlTaskById":                                 nullEndPoint.Init(), // "/v1/devService/queryCtrlTaskById"}
			"queryDeviceInfo":                                   nullEndPoint.Init(), // "/v1/devService/queryDeviceInfoForApp"}
			"queryDeviceInfoForApp":                             nullEndPoint.Init(), // "/v1/devService/queryDeviceInfoForApp"}
			"queryDeviceList":                                   nullEndPoint.Init(), // "/v1/devService/queryDeviceList"}
			"queryDeviceListByUserId":                           nullEndPoint.Init(), // "/v1/devService/queryDeviceListByUserId"}
			"queryDeviceListForApp":                             nullEndPoint.Init(), // "/v1/devService/queryDeviceListForApp"}
			"queryDeviceModelTechnical":                         nullEndPoint.Init(), // "/v1/devService/queryDeviceModelTechnical"}
			"queryDevicePointDayMonthYearDataList":              nullEndPoint.Init(), // "/v1/commonService/queryDevicePointDayMonthYearDataList"}
			"queryDevicePointMinuteDataList":                    nullEndPoint.Init(), // "/v1/commonService/queryDevicePointMinuteDataList"}
			"queryDevicePointsDayMonthYearDataList":             nullEndPoint.Init(), // "/v1/commonService/queryDevicePointsDayMonthYearDataList"}
			"queryDeviceRealTimeDataByPsKeys":                   nullEndPoint.Init(), // "/v1/devService/queryDeviceRealTimeDataByPsKeys"}
			"queryDeviceRepairList":                             nullEndPoint.Init(), // "/v1/devService/queryDeviceRepairList"}
			"queryDeviceTypeInfoList":                           nullEndPoint.Init(), // "/v1/devService/queryDeviceTypeInfoList"}
			"queryEnvironmentList":                              nullEndPoint.Init(), // "/v1/devService/queryEnvironmentList"}
			"queryFaultList":                                    nullEndPoint.Init(), // "/v1/faultService/queryFaultList"}
			"queryFaultPlanDetail":                              nullEndPoint.Init(), // "/v1/faultService/queryFaultPlanDetail"}
			"queryFaultRepairSteps":                             nullEndPoint.Init(), // "/v1/faultService/queryFaultRepairSteps"}
			"queryFaultTypeAndLevelByCode":                      nullEndPoint.Init(), // "/v1/faultService/queryFaultTypeAndLevelByCode"}
			"queryFaultTypeByDevice":                            nullEndPoint.Init(), // "/v1/faultService/queryFaultTypeByDevice"}
			"queryFaultTypeByDevicePage":                        nullEndPoint.Init(), // "/v1/faultService/queryFaultTypeByDevicePage"}
			"queryFirmwareFilesPage":                            nullEndPoint.Init(), // "/v1/commonService/queryFirmwareFilesPage"}
			"queryInfotoAlert":                                  nullEndPoint.Init(), // "/v1/devService/queryInfotoAlert"}
			"queryInverterModelList":                            nullEndPoint.Init(), // "/v1/devService/queryInverterModelList"}
			"queryInverterVersionList":                          nullEndPoint.Init(), // "/v1/devService/queryInverterVersionList"}
			"queryM2MCardInfoCMCC":                              nullEndPoint.Init(), // "/v1/devService/queryM2MCardInfoCMCC"}
			"queryM2MCardTermInfoCMCC":                          nullEndPoint.Init(), // "/v1/devService/queryM2MCardTermInfoCMCC"}
			"queryModelInfoByModelId":                           nullEndPoint.Init(), // "/v1/devService/queryModelInfoByModelId"}
			"queryMutiPointDataList":                            nullEndPoint.Init(), // "/v1/commonService/queryMutiPointDataList"}
			"queryNoticeList":                                   nullEndPoint.Init(), // "/v1/otherService/queryNoticeList"}
			"queryNumberOfRenewalReminders":                     nullEndPoint.Init(), // "/v1/devService/queryNumberOfRenewalReminders"}
			"queryOperRules":                                    nullEndPoint.Init(), // "/v1/faultService/queryOperRules"}
			"queryOrderList":                                    nullEndPoint.Init(), // "/v1/faultService/queryOrderList"}
			"queryOrderStep":                                    nullEndPoint.Init(), // "/v1/faultService/queryOrderStep"}
			"queryOrgGenerationReport":                          nullEndPoint.Init(), // "/v1/orgService/queryOrgGenerationReport"}
			"queryOrgInfoList":                                  nullEndPoint.Init(), // "/v1/userService/queryOrgInfoList"}
			"queryOrgPowerElecPercent":                          nullEndPoint.Init(), // "/v1/orgService/queryOrgPowerElecPercent"}
			"queryOrgPsCompensationRecordList":                  nullEndPoint.Init(), // "/v1/powerStationService/queryOrgPsCompensationRecordList"}
			"queryParamSettingTask":                             nullEndPoint.Init(), // "/v1/devService/queryParamSettingTask"}
			"queryPersonalUnitList":                             nullEndPoint.Init(), // "/v1/userService/queryPersonalUnitList"}
			"queryPointDataTopOne":                              nullEndPoint.Init(), // "/v1/devService/queryPointDataTopOne"}
			"queryPowerStationInfo":                             nullEndPoint.Init(), // "/v1/powerStationService/queryPowerStationInfo"}
			"queryPsAreaByUserIdAndAreaCode":                    nullEndPoint.Init(), // "/v1/powerStationService/queryPsAreaByUserIdAndAreaCode"}
			"queryPsCompensationRecordList":                     nullEndPoint.Init(), // "/v1/powerStationService/queryPsCompensationRecordList"}
			"queryPsDataByDate":                                 nullEndPoint.Init(), // "/v1/powerStationService/queryPsDataByDate"}
			"queryPsIdList":                                     nullEndPoint.Init(), // "/v1/powerStationService/queryPsIdList"}
			"queryPsListByUserIdAndAreaCode":                    nullEndPoint.Init(), // "/v1/powerStationService/queryPsListByUserIdAndAreaCode"}
			"queryPsNameByPsId":                                 nullEndPoint.Init(), // "/v1/devService/queryPsNameByPsId"}
			"queryPsPrByDate":                                   nullEndPoint.Init(), // "/v1/powerStationService/queryPsPrByDate"}
			"queryPsProfit":                                     nullEndPoint.Init(), // "/v1/powerStationService/queryPsProfit"}
			"queryPsReportComparativeAnalysisOfPowerGeneration": nullEndPoint.Init(), // "/v1/powerStationService/queryPsReportComparativeAnalysisOfPowerGeneration"}
			"queryPsStructureList":                              nullEndPoint.Init(), // "/v1/devService/queryPsStructureList"}
			"queryPuuidsByCommandTotalId":                       nullEndPoint.Init(), // "/v1/devService/queryPuuidsByCommandTotalId"}
			"queryPuuidsByCommandTotalId2":                      nullEndPoint.Init(), // "/v1/devService/queryPuuidsByCommandTotalId2"}
			"queryRepairRuleList":                               nullEndPoint.Init(), // "/v1/devService/queryRepairRuleList"}
			"queryReportListForManagementPage":                  nullEndPoint.Init(), // "/v1/reportService/queryReportListForManagementPage"}
			"queryReportMsg":                                    nullEndPoint.Init(), // "/v1/reportService/queryReportMsg"}
			"querySharingPs":                                    nullEndPoint.Init(), // "/v1/powerStationService/querySharingPs"}
			"querySysAdvancedParam":                             nullEndPoint.Init(), // "/v1/devService/querySysAdvancedParam"}
			"queryTimeBySN":                                     nullEndPoint.Init(), // "/v1/devService/queryTimeBySN"}
			"queryTrafficByDateCTCC":                            nullEndPoint.Init(), // "/v1/devService/queryTrafficByDateCTCC"}
			"queryTrafficCTCC":                                  nullEndPoint.Init(), // "/v1/devService/queryTrafficCTCC"}
			"queryUnitList":                                     nullEndPoint.Init(), // "/v1/userService/queryUnitList"}
			"queryUnitUuidBytotalId":                            nullEndPoint.Init(), // "/v1/devService/queryUnitUuidBytotalId"}
			"queryUserBtnPri":                                   nullEndPoint.Init(), // "/v1/userService/queryUserBtnPri"}
			"queryUserByUserIds":                                nullEndPoint.Init(), // "/v1/userService/queryUserByUserIds"}
			"queryUserExtensionAttribute":                       nullEndPoint.Init(), // "/v1/userService/queryUserExtensionAttribute"}
			"queryUserForStep":                                  nullEndPoint.Init(), // "/v1/userService/queryUserForStep"}
			"queryUserList":                                     nullEndPoint.Init(), // "/v1/userService/queryUserList"}
			"queryUserProcessPri":                               nullEndPoint.Init(), // "/v1/userService/queryUserProcessPri"}
			"queryUserWechatBindRel":                            nullEndPoint.Init(), // "/v1/userService/queryUserWechatBindRel"}
			"queryUuidByTotalIdAndUuid":                         nullEndPoint.Init(), // "/v1/devService/queryUuidByTotalIdAndUuid"}
			"rechargeOrderSetMeal":                              nullEndPoint.Init(), // "/v1/devService/rechargeOrderSetMeal"}
			"renewSendReportConfirmEmail":                       nullEndPoint.Init(), // "/v1/reportService/renewSendReportConfirmEmail"}
			"reportList":                                        nullEndPoint.Init(), // "/v1/powerStationService/reportList"}
			"saveCustomerEmployee":                              nullEndPoint.Init(), // "/v1/devService/saveCustomerEmployee"}
			"saveDevSimList":                                    nullEndPoint.Init(), // "/v1/devService/saveDevSimList"}
			"saveDeviceAccountBatchData":                        nullEndPoint.Init(), // "/v1/devService/saveDeviceAccountBatchData"}
			"saveEnviromentIncomeInfos":                         nullEndPoint.Init(), // "/v1/powerStationService/saveEnviromentIncomeInfos"}
			"saveEnvironmentCurve":                              nullEndPoint.Init(), // "/v1/devService/saveEnvironmentCurve"}
			"saveFirmwareFile":                                  nullEndPoint.Init(), // "/v1/commonService/saveFirmwareFile"}
			"saveIncomeSettingInfos":                            nullEndPoint.Init(), // "/v1/powerStationService/saveIncomeSettingInfos"}
			"saveOrUpdateGroupStringCheckRule":                  nullEndPoint.Init(), // "/v1/devService/saveOrUpdateGroupStringCheckRule"}
			"saveParamModel":                                    nullEndPoint.Init(), // "/v1/devService/saveParamModel"}
			"savePowerCharges":                                  nullEndPoint.Init(), // "/v1/powerStationService/savePowerCharges"}
			"savePowerDevicePoint":                              nullEndPoint.Init(), // "/v1/reportService/savePowerDevicePoint"}
			"savePowerRobotInfo":                                nullEndPoint.Init(), // "/v1/devService/savePowerRobotInfo"}
			"savePowerRobotSweepAttr":                           nullEndPoint.Init(), // "/v1/devService/savePowerRobotSweepAttr"}
			"savePowerSettingCharges":                           nullEndPoint.Init(), // "/v1/powerStationService/savePowerSettingCharges"}
			"savePowerSettingInfo":                              nullEndPoint.Init(), // "/v1/powerStationService/savePowerSettingInfo"}
			"saveProductionBatchData":                           nullEndPoint.Init(), // "/v1/devService/saveProductionBatchData"}
			"saveRechargeOrderObj":                              nullEndPoint.Init(), // "/onlinepay/saveRechargeOrderObj"}
			"saveRechargeOrderOtherInfo":                        nullEndPoint.Init(), // "/onlinepay/saveRechargeOrderOtherInfo"}
			"saveRepair":                                        nullEndPoint.Init(), // "/v1/faultService/saveRepair"}
			"saveReportExportColumns":                           nullEndPoint.Init(), // "/v1/reportService/saveReportExportColumns"}
			"saveSetParam":                                      nullEndPoint.Init(), // "/v1/devService/saveSetParam"}
			"saveSysUserMsg":                                    nullEndPoint.Init(), // "/v1/otherService/saveSysUserMsg"}
			"saveTemplate":                                      nullEndPoint.Init(), // "/v1/devService/saveTemplate"}
			"searchM2MMonthFlowCMCC":                            nullEndPoint.Init(), // "/v1/devService/searchM2MMonthFlowCMCC"}
			"selectSysTranslationNames":                         nullEndPoint.Init(), // "/v1/reportService/selectSysTranslationNames"}
			"sendPsTimeZoneInstruction":                         nullEndPoint.Init(), // "/devDataHandleService/sendPsTimeZoneInstruction"}
			"setUpFormulaFaultAnalyse":                          nullEndPoint.Init(), // "/v1/powerStationService/setUpFormulaFaultAnalyse"}
			"setUserGDPRAttrs":                                  nullEndPoint.Init(), // "/v1/userService/setUserGDPRAttrs"}
			"settingNotice":                                     nullEndPoint.Init(), // "/v1/userService/settingNotice"}
			"shareMyPs":                                         nullEndPoint.Init(), // "/v1/powerStationService/shareMyPs"}
			"sharePsBySN":                                       nullEndPoint.Init(), // "/v1/powerStationService/sharePsBySN"}
			"showInverterByUnit":                                nullEndPoint.Init(), // "/v1/devService/showInverterByUnit"}
			"showOnlineUsers":                                   nullEndPoint.Init(), // "/v1/userService/showOnlineUsers"}
			"showWarning":                                       nullEndPoint.Init(), // "/v1/userService/showWarning"}
			"snIsExist":                                         nullEndPoint.Init(), // "/v1/devService/snIsExist"}
			"snsIsExist":                                        nullEndPoint.Init(), // "/v1/devService/snsIsExist"}
			"speedyAddPowerStation":                             nullEndPoint.Init(), // "/v1/powerStationService/speedyAddPowerStation"}
			"stationDeviceHistoryDataList":                      nullEndPoint.Init(), // "/v1/powerStationService/stationDeviceHistoryDataList"}
			"stationUnitsList":                                  nullEndPoint.Init(), // "/v1/powerStationService/stationUnitsList"}
			"stationsDiscreteData":                              nullEndPoint.Init(), // "/v1/devService/stationsDiscreteData"}
			"stationsIncomeList":                                nullEndPoint.Init(), // "/v1/powerStationService/stationsIncomeList"}
			"stationsPointReport":                               nullEndPoint.Init(), // "/v1/powerStationService/stationsPointReport"}
			"stationsYearPlanReport":                            nullEndPoint.Init(), // "/v1/reportService/stationsYearPlanReport"}
			"sureAndImportSelettlementData":                     nullEndPoint.Init(), // "/v1/otherService/sureAndImportSelettlementData"}
			"sweepDevParamSet":                                  nullEndPoint.Init(), // "/v1/devService/sweepDevParamSet"}
			"sweepDevRunControl":                                nullEndPoint.Init(), // "/v1/devService/sweepDevRunControl"}
			"sweepDevStrategyIssue":                             nullEndPoint.Init(), // "/v1/devService/sweepDevStrategyIssue"}
			"sysTimeZoneList":                                   nullEndPoint.Init(), // "/v1/commonService/sysTimeZoneList"}
			"unLockUser":                                        nullEndPoint.Init(), // "/v1/userService/unLockUser"}
			"unlockChildAccount":                                nullEndPoint.Init(), // "/v1/userService/unlockChildAccount"}
			"updateCommunicationModuleState":                    nullEndPoint.Init(), // "/v1/devService/updateCommunicationModuleState"}
			"updateDevInstalledPower":                           nullEndPoint.Init(), // "/v1/devService/updateDevInstalledPower"}
			"updateFault":                                       nullEndPoint.Init(), // "/v1/faultService/updateFaultStatus"}
			"updateFaultData":                                   nullEndPoint.Init(), // "/v1/faultService/updateFaultData"}
			"updateFaultMsgByFaultCode":                         nullEndPoint.Init(), // "/v1/faultService/updateFaultMsgByFaultCode"}
			"updateFaultStatus":                                 nullEndPoint.Init(), // "/v1/faultService/updateFaultStatus"}
			"updateHouseholdWorkOrder":                          nullEndPoint.Init(), // "/v1/faultService/updateHouseholdWorkOrder"}
			"updateInverterSn2ModuleSn":                         nullEndPoint.Init(), // "/devDataHandleService/updateInverterSn2ModuleSn"}
			"updateOperateTicketAttachmentId":                   nullEndPoint.Init(), // "/v1/faultService/updateOperateTicketAttachmentId"}
			"updateOrderDeviceByCustomerService":                nullEndPoint.Init(), // "/onlinepay/updateOrderDeviceByCustomerService"}
			"updateOwnerFaultConfig":                            nullEndPoint.Init(), // "/v1/faultService/updateOwnerFaultConfig"}
			"updateParamSettingSysMsg":                          nullEndPoint.Init(), // "/v1/devService/updateParamSettingSysMsg"}
			"updatePlatformLevelFaultLevel":                     nullEndPoint.Init(), // "/v1/faultService/updatePlatformLevelFaultLevel"}
			"updatePowerDevicePoint":                            nullEndPoint.Init(), // "/v1/reportService/updatePowerDevicePoint"}
			"updatePowerRobotInfo":                              nullEndPoint.Init(), // "/v1/devService/updatePowerRobotInfo"}
			"updatePowerRobotSweepAttr":                         nullEndPoint.Init(), // "/v1/devService/updatePowerRobotSweepAttr"}
			"updatePowerStationForHousehold":                    nullEndPoint.Init(), // "/v1/powerStationService/updatePowerStationForHousehold"}
			"updatePowerStationInfo":                            nullEndPoint.Init(), // "/v1/powerStationService/updatePowerStationInfo"}
			"updatePowerUserInfo":                               nullEndPoint.Init(), // "/v1/userService/updatePowerUserInfo"}
			"updateReportConfigByEmailAddr":                     nullEndPoint.Init(), // "/v1/reportService/updateReportConfigByEmailAddr"}
			"updateShareAttr":                                   nullEndPoint.Init(), // "/v1/powerStationService/updateShareAttr"}
			"updateSnIsSureFlag":                                nullEndPoint.Init(), // "/devDataHandleService/updateSnIsSureFlag"}
			"updateStationPics":                                 nullEndPoint.Init(), // "/v1/powerStationService/updateStationPics"}
			"updateSysAdvancedParam":                            nullEndPoint.Init(), // "/v1/devService/updateSysAdvancedParam"}
			"updateSysOrgNew":                                   nullEndPoint.Init(), // "/v1/otherService/updateSysOrgNew"}
			"updateTemplate":                                    nullEndPoint.Init(), // "/v1/devService/updateDataCurveTemplate"}
			"updateUinfoNetEaseUser":                            nullEndPoint.Init(), // "/v1/userService/updateUinfoNetEaseUser"}
			"updateUserExtensionAttribute":                      nullEndPoint.Init(), // "/v1/userService/updateUserExtensionAttribute"}
			"updateUserLanguage":                                nullEndPoint.Init(), // "/v1/userService/updateUserLanguage"}
			"updateUserPosition":                                nullEndPoint.Init(), // "/v1/userService/updateUserPosition"}
			"updateUserUpOrg":                                   nullEndPoint.Init(), // "/v1/orgService/updateUserUpOrg"}
			"upgrade":                                           nullEndPoint.Init(), // "/v1/userService/upgrade"}
			"upgrate":                                           nullEndPoint.Init(), // "/v1/userService/upgrade"}
			"uploadFileToOss":                                   nullEndPoint.Init(), // "/v1/commonService/uploadFileToOss"}
			"userAgreeGdprProtocol":                             nullEndPoint.Init(), // "/v1/userService/userAgreeGdprProtocol"}
			"userInfoUniqueCheck":                               nullEndPoint.Init(), // "/v1/userService/userInfoUniqueCheck"}
			"userMailHasBound":                                  nullEndPoint.Init(), // "/v1/userService/userMailHasBound"}
			"userRegister":                                      nullEndPoint.Init(), // "/v1/userService/userRegister"}
		},
	}

	fmt.Printf("area[%s]: %v\n", name, area)

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

func (a Area) Init() api.AreaStruct {
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
