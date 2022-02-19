package AppService

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService/getPowerDevicePointNames"
	"GoSungro/iSolarCloud/sungro/AppService/getPowerStatistics"
	"GoSungro/iSolarCloud/sungro/AppService/getPsDetailWithPsType"
	"GoSungro/iSolarCloud/sungro/AppService/getPsList"
	"GoSungro/iSolarCloud/sungro/AppService/login"
	"GoSungro/iSolarCloud/sungro/AppService/nullEndPoint"
	"fmt"
)

var _ api.Area = (*Area)(nil)

type Area api.AreaStruct


func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot *api.Web) Area {
	// fmt.Println("Init(apiRoot)")
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)

	area := Area {
		ApiRoot:   apiRoot,
		Name:      api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints {
			"acceptPsSharing":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/acceptPsSharing"}
			"activateEmail":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/activateEmail"}
			"addConfig":                                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/addConfig"}
			"addDeviceRepair":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/addDeviceRepair"}
			"addDeviceToStructureForHousehold":                  nullEndPoint.Init(apiRoot), // "/devDataHandleService/addDeviceToStructureForHousehold"}
			"addDeviceToStructureForHouseholdByPsIdS":           nullEndPoint.Init(apiRoot), // "/devDataHandleService/addDeviceToStructureForHouseholdByPsIdS"}
			"addFault":                                          nullEndPoint.Init(apiRoot), // "/v1/faultService/addFault"}
			"addFaultOrder":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultOrder"}
			"addFaultPlan":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultPlan"}
			"addFaultRepairSteps":                               nullEndPoint.Init(apiRoot), // "/v1/faultService/addFaultRepairSteps"}
			"addHouseholdEvaluation":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdEvaluation"}
			"addHouseholdLeaveMessage":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdLeaveMessage"}
			"addHouseholdOpinionFeedback":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdOpinionFeedback"}
			"addHouseholdWorkOrder":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/addHouseholdWorkOrder"}
			"addOnDutyInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/otherService/addOnDutyInfo"}
			"addOperRule":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/addOperRule"}
			"addOrDelPsStructure":                               nullEndPoint.Init(apiRoot), // "/v1/devService/addOrDelPsStructure"}
			"addOrderStep":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/addOrderStep"}
			"addPowerStationForHousehold":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/addPowerStationForHousehold"}
			"addPowerStationInfo":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/addPowerStationInfo"}
			"addReportConfigEmail":                              nullEndPoint.Init(apiRoot), // "/v1/reportService/addReportConfigEmail"}
			"addSysAdvancedParam":                               nullEndPoint.Init(apiRoot), // "/v1/devService/addSysAdvancedParam"}
			"addSysOrgNew":                                      nullEndPoint.Init(apiRoot), // "/v1/otherService/addSysOrgNew"}
			"aliPayAppTest":                                     nullEndPoint.Init(apiRoot), // "/onlinepay/aliPayAppTest"}
			"auditOperRule":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/auditOperRule"}
			"batchAddStationBySn":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/batchAddStationBySn"}
			"batchImportSN":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/batchImportSN"}
			"batchInsertUserAndOrg":                             nullEndPoint.Init(apiRoot), // "/v1/userService/batchInsertUserAndOrg"}
			"batchModifyDevicesInfoAndPropertis":                nullEndPoint.Init(apiRoot), // "/v1/devService/batchModifyDevicesInfoAndPropertis"}
			"batchProcessPlantReport":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/batchProcessPlantReport"}
			"batchUpdateDeviceSim":                              nullEndPoint.Init(apiRoot), // "/v1/devService/batchUpdateDeviceSim"}
			"batchUpdateUserIsAgreeGdpr":                        nullEndPoint.Init(apiRoot), // "/v1/userService/batchUpdateUserIsAgreeGdpr"}
			"boundMobilePhone":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/boundMobilePhone"}
			"boundUserMail":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/boundUserMail"}
			"caculateDeviceInputDiscrete":                       nullEndPoint.Init(apiRoot), // "/v1/devService/caculateDeviceInputDiscrete"}
			"calculateDeviceDiscrete":                           nullEndPoint.Init(apiRoot), // "/v1/devService/calculateDeviceDiscrete"}
			"calculateInitialCompensationData":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/calculateInitialCompensationData"}
			"cancelDeliverMail":                                 nullEndPoint.Init(apiRoot), // "/v1/messageService/cancelDeliverMail"}
			"cancelOrderScan":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/cancelOrderScan"}
			"cancelParamSetTask":                                nullEndPoint.Init(apiRoot), // "/v1/devService/cancelParamSetTask"}
			"cancelPsSharing":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/cancelPsSharing"}
			"cancelRechargeOrder":                               nullEndPoint.Init(apiRoot), // "/onlinepay/cancelRechargeOrder"}
			"changRechargeOrderToCancel":                        nullEndPoint.Init(apiRoot), // "/onlinepay/changRechargeOrderToCancel"}
			"changeHouseholdUser2Installer":                     nullEndPoint.Init(apiRoot), // "/v1/orgService/changeHouseholdUser2Installer"}
			"changeRemoteParam":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/changeRemoteParam"}
			"checkDealerOrgCode":                                nullEndPoint.Init(apiRoot), // "/v1/orgService/checkDealerOrgCode"}
			"checkDevSnIsBelongsToUser":                         nullEndPoint.Init(apiRoot), // "/v1/userService/checkDevSnIsBelongsToUser"}
			"checkInverterResult":                               nullEndPoint.Init(apiRoot), // "/v1/devService/checkInverterResult"}
			"checkIsCanDoParamSet":                              nullEndPoint.Init(apiRoot), // "/v1/devService/checkIsCanDoParamSet"}
			"checkIsIvScan":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/checkIsIvScan"}
			"checkOssObjectExist":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/checkOssObjectExist"}
			"checkServiceIsConnect":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/checkServiceIsConnect"}
			"checkTechnicalParameters":                          nullEndPoint.Init(apiRoot), // "/v1/devService/checkTechnicalParameters"}
			"checkUnitStatus":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/checkUnitStatus"}
			"checkUpRechargeDevicePaying":                       nullEndPoint.Init(apiRoot), // "/onlinepay/checkUpRechargeDevicePaying"}
			"checkUserAccountUnique":                            nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserAccountUnique"}
			"checkUserAccountUniqueAll":                         nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserAccountUniqueAll"}
			"checkUserInfoUnique":                               nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserInfoUnique"}
			"checkUserIsExist":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserIsExist"}
			"checkUserListIsExist":                              nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserListIsExist"}
			"checkUserPassword":                                 nullEndPoint.Init(apiRoot), // "/v1/userService/checkUserPassword"}
			"cloudDeploymentRecord":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/cloudDeploymentRecord"}
			"comfirmParamModel":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/comfirmParamModel"}
			"communicationModuleDetail":                         nullEndPoint.Init(apiRoot), // "/v1/devService/communicationModuleDetail"}
			"compareValidateCode":                               nullEndPoint.Init(apiRoot), // "/v1/userService/compareValidateCode"}
			"componentInfo2Cloud":                               nullEndPoint.Init(apiRoot), // "/v1/devService/componentInfo2Cloud"}
			"confirmFault":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/confirmFault"}
			"confirmIvFault":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/confirmIvFault"}
			"confirmReportConfig":                               nullEndPoint.Init(apiRoot), // "/v1/reportService/confirmReportConfig"}
			"createAppkeyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/createAppkeyInfo"}
			"createRenewInvoice":                                nullEndPoint.Init(apiRoot), // "/onlinepay/createRenewInvoice"}
			"dealCommandReply":                                  nullEndPoint.Init(apiRoot), // "/devDataHandleService/dealCommandReply"}
			"dealDeletePsFailPsDelete":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/dealDeletePsFailPsDelete"}
			"dealFailRemoteUpgradeSubTasks":                     nullEndPoint.Init(apiRoot), // "/v1/devService/dealFailRemoteUpgradeSubTasks"}
			"dealFailRemoteUpgradeTasks":                        nullEndPoint.Init(apiRoot), // "/v1/devService/dealFailRemoteUpgradeTasks"}
			"dealFaultOrder":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/dealFaultOrder"}
			"dealGroupStringDisableOrEnable":                    nullEndPoint.Init(apiRoot), // "/v1/devService/dealGroupStringDisableOrEnable"}
			"dealNumberOfServiceCalls2Mysql":                    nullEndPoint.Init(apiRoot), // "/v1/commonService/dealNumberOfServiceCalls2Mysql"}
			"dealParamSettingAfterComplete":                     nullEndPoint.Init(apiRoot), // "/v1/devService/dealParamSettingAfterComplete"}
			"dealPsDataSupplement":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/dealPsDataSupplement"}
			"dealPsReportEmailSend":                             nullEndPoint.Init(apiRoot), // "/v1/reportService/dealPsReportEmailSend"}
			"dealRemoteUpgrade":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/dealRemoteUpgrade"}
			"dealSnElectrifyCheck":                              nullEndPoint.Init(apiRoot), // "/v1/devService/dealSnElectrifyCheck"}
			"dealSysDeviceSimFlowInfo":                          nullEndPoint.Init(apiRoot), // "/v1/devService/dealSysDeviceSimFlowInfo"}
			"dealSysDeviceSimInfo":                              nullEndPoint.Init(apiRoot), // "/v1/devService/dealSysDeviceSimInfo"}
			"definiteTimeDealSnExpRemind":                       nullEndPoint.Init(apiRoot), // "/v1/devService/definiteTimeDealSnExpRemind"}
			"definiteTimeDealSnStatus":                          nullEndPoint.Init(apiRoot), // "/v1/devService/definiteTimeDealSnStatus"}
			"delDeviceRepair":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/delDeviceRepair"}
			"delOperRule":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/delOperRule"}
			"delayCallApiResidueTimes":                          nullEndPoint.Init(apiRoot), // "/v1/commonService/delayCallApiResidueTimes"}
			"deleteComponent":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/deleteComponent"}
			"deleteCustomerEmployee":                            nullEndPoint.Init(apiRoot), // "/v1/devService/deleteCustomerEmployee"}
			"deleteDeviceAccount":                               nullEndPoint.Init(apiRoot), // "/v1/devService/deleteDeviceAccount"}
			"deleteDeviceSimById":                               nullEndPoint.Init(apiRoot), // "/v1/devService/deleteDeviceSimById"}
			"deleteElectricitySettlementData":                   nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteElectricitySettlementData"}
			"deleteFaultPlan":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteFaultPlan"}
			"deleteFirmwareFiles":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteFirmwareFiles"}
			"deleteHouseholdEvaluation":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdEvaluation"}
			"deleteHouseholdLeaveMessage":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdLeaveMessage"}
			"deleteHouseholdWorkOrder":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteHouseholdWorkOrder"}
			"deleteInverterSnInChnnl":                           nullEndPoint.Init(apiRoot), // "/devDataHandleService/deleteInverterSnInChnnl"}
			"deleteModuleLog":                                   nullEndPoint.Init(apiRoot), // "/integrationService/deleteModuleLog"}
			"deleteOnDutyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteOnDutyInfo"}
			"deleteOperateBillFile":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/deleteOperateBillFile"}
			"deleteOssObject":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteOssObject"}
			"deletePowerDevicePointById":                        nullEndPoint.Init(apiRoot), // "/v1/reportService/deletePowerDevicePointById"}
			"deletePowerPicture":                                nullEndPoint.Init(apiRoot), // "/v1/powerStationService/deletePowerPicture"}
			"deletePowerRobotInfoBySnAndPsId":                   nullEndPoint.Init(apiRoot), // "/v1/devService/deletePowerRobotInfoBySnAndPsId"}
			"deletePowerRobotSweepStrategy":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deletePowerRobotSweepStrategy"}
			"deleteProductionData":                              nullEndPoint.Init(apiRoot), // "/v1/devService/deleteProductionData"}
			"deletePs":                                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/deletePs"}
			"deleteRechargeOrder":                               nullEndPoint.Init(apiRoot), // "/onlinepay/deleteRechargeOrder"}
			"deleteRegularlyConnectionInfo":                     nullEndPoint.Init(apiRoot), // "/v1/commonService/deleteRegularlyConnectionInfo"}
			"deleteReportConfigEmailAddr":                       nullEndPoint.Init(apiRoot), // "/v1/reportService/deleteReportConfigEmailAddr"}
			"deleteSysAdvancedParam":                            nullEndPoint.Init(apiRoot), // "/v1/devService/deleteSysAdvancedParam"}
			"deleteSysOrgNew":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/deleteSysOrgNew"}
			"deleteTemplate":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/deleteTemplate"}
			"deleteUserInfoAllByUserId":                         nullEndPoint.Init(apiRoot), // "/v1/userService/deleteUserInfoAllByUserId"}
			"deviceInputDiscreteDeleteTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteDeleteTime"}
			"deviceInputDiscreteGetTime":                        nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteGetTime"}
			"deviceInputDiscreteInsertTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteInsertTime"}
			"deviceInputDiscreteUpdateTime":                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceInputDiscreteUpdateTime"}
			"devicePointsDataFromMySql":                         nullEndPoint.Init(apiRoot), // "/v1/devService/devicePointsDataFromMySql"}
			"deviceReplace":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/deviceReplace"}
			"editDeviceRepair":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/editDeviceRepair"}
			"editOperRule":                                      nullEndPoint.Init(apiRoot), // "/v1/faultService/editOperRule"}
			"energyPovertyAlleviation":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/energyPovertyAlleviation"}
			"energyTrend":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/energyTrend"}
			"exportParamSettingValPDF":                          nullEndPoint.Init(apiRoot), // "/v1/devService/exportParamSettingValPDF"}
			"exportPlantReportPDF":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/exportPlantReportPDF"}
			"faultAutoClose":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/faultAutoClose"}
			"faultCloseRemindOrderHandler":                      nullEndPoint.Init(apiRoot), // "/v1/faultService/faultCloseRemindOrderHandler"}
			"findCodeValueList":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/findCodeValueList"}
			"findEmgOrgInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/findEmgOrgInfo"}
			"findEnvironmentInfo":                               nullEndPoint.Init(apiRoot), // "/v1/devService/findEnvironmentInfo"}
			"findFromHbaseAndRedis":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/findFromHbaseAndRedis"}
			"findInfoByuuid":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/findInfoByuuid"}
			"findLossAnalysisList":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/findLossAnalysisList"}
			"findOnDutyInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/findOnDutyInfo"}
			"findPsType":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/findPsType"}
			"findSingleStationPR":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/findSingleStationPR"}
			"findUserPassword":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/findUserPassword"}
			"genTLSUserSigByUserAccount":                        nullEndPoint.Init(apiRoot), // "/v1/userService/genTLSUserSigByUserAccount"}
			"generateRandomPassword":                            nullEndPoint.Init(apiRoot), // "/v1/userService/generateRandomPassword"}
			"getAPIServiceInfo":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getAPIServiceInfo"}
			"getAccessedPermission":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/getAccessedPermission"}
			"getAllDeviceByPsId":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getAllDeviceByPsId"}
			"getAllPowerDeviceSetName":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPowerDeviceSetName"}
			"getAllPowerRobotViewInfoByPsId":                    nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPowerRobotViewInfoByPsId"}
			"getAllPsIdByOrgIds":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getAllPsIdByOrgIds"}
			"getAllUserRemindCount":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getAllUserRemindCount"}
			"getAndOutletsAndUnit":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getAndOutletsAndUnit"}
			"getApiCallsForAppkeys":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/getApiCallsForAppkeys"}
			"getAreaInfoCodeByCounty":                           nullEndPoint.Init(apiRoot), // "/v1/commonService/getAreaInfoCodeByCounty"}
			"getAreaList":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getAreaList"}
			"getAutoCreatePowerStation":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getAutoCreatePowerStation"}
			"getBackReadValue":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getBackReadValue"}
			"getBatchNewestPointData":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getBatchNewestPointData"}
			"getCallApiResidueTimes":                            nullEndPoint.Init(apiRoot), // "/v1/commonService/getCallApiResidueTimes"}
			"getChangedPsListByTime":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getChangedPsListByTime"}
			"getChnnlListByPsId":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getChnnlListByPsId"}
			"getCloudList":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getCloudList"}
			"getCloudServiceMappingConfig":                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getCloudServiceMappingConfig"}
			"getCommunicationDeviceConfigInfo":                  nullEndPoint.Init(apiRoot), // "/v1/devService/getCommunicationDeviceConfigInfo"}
			"getCommunicationModuleMonitorData":                 nullEndPoint.Init(apiRoot), // "/v1/devService/getCommunicationModuleMonitorData"}
			"getComponentModelFactory":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getComponentModelFactory"}
			"getConfigList":                                     nullEndPoint.Init(apiRoot), // "/devDataHandleService/getConfigList"}
			"getConnectionInfoBySnAndLocalPort":                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getConnectionInfoBySnAndLocalPort"}
			"getCountDown":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/getCountDown"}
			"getCountryServiceInfo":                             nullEndPoint.Init(apiRoot), // "/v1/commonService/getCountryServiceInfo"}
			"getCounty":                                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getCounty"}
			"getCustomerEmployee":                               nullEndPoint.Init(apiRoot), // "/v1/devService/getCustomerEmployee"}
			"getCustomerList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getCustomerList"}
			"getDataFromHBase":                                  nullEndPoint.Init(apiRoot), // "/v1/commonService/getDataFromHBase"}
			"getDataFromHbaseByRowKey":                          nullEndPoint.Init(apiRoot), // "/v1/commonService/getDataFromHbaseByRowKey"}
			"getDevInstalledPowerByPsId":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getDevInstalledPowerByPsId"}
			"getDevRecord":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/getDevRecord"}
			"getDevRunRecordList":                               nullEndPoint.Init(apiRoot), // "/v1/devService/getDevRunRecordList"}
			"getDevSimByList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getDevSimByList"}
			"getDevSimList":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDevSimList"}
			"getDeviceAccountById":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceAccountById"}
			"getDeviceFaultStatisticsData":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceFaultStatisticsData"}
			"getDeviceInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceInfo"}
			"getDeviceList":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceList"}
			"getDeviceModelInfoList":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceModelInfoList"}
			"getDevicePointMinuteDataList":                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getDevicePointMinuteDataList"}
			"getDevicePoints":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getDevicePoints"}
			"getDevicePropertys":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getDevicePropertys"}
			"getDeviceRepairDetail":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceRepairDetail"}
			"getDeviceTechBranchCount":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTechBranchCount"}
			"getDeviceTypeInfoList":                             nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTypeInfoList"}
			"getDeviceTypeList":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getDeviceTypeList"}
			"getDstInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/userService/getDstInfo"}
			"getElectricitySettlementData":                      nullEndPoint.Init(apiRoot), // "/v1/otherService/getElectricitySettlementData"}
			"getElectricitySettlementDetailData":                nullEndPoint.Init(apiRoot), // "/v1/otherService/getElectricitySettlementDetailData"}
			"getEncryptPublicKey":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/getEncryptPublicKey"}
			"getFaultCount":                                     nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultCount"}
			"getFaultDetail":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultDetail"}
			"getFaultMsgByFaultCode":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgByFaultCode"}
			"getFaultMsgListByPageWithYYYYMM":                   nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgListByPageWithYYYYMM"}
			"getFaultMsgListWithYYYYMM":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultMsgListWithYYYYMM"}
			"getFaultPlanList":                                  nullEndPoint.Init(apiRoot), // "/v1/faultService/getFaultPlanList"}
			"getFileOperationRecordOne":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getFileOperationRecordOne"}
			"getFormulaFaultAnalyseList":                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getFormulaFaultAnalyseList"}
			"getGroupStringCheckResult":                         nullEndPoint.Init(apiRoot), // "/v1/devService/getGroupStringCheckResult"}
			"getGroupStringCheckRule":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getGroupStringCheckRule"}
			"getHisData":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/getHisData"}
			"getHistoryInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getHistoryInfo"}
			"getHouseholdEvaluation":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdEvaluation"}
			"getHouseholdLeaveMessage":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdLeaveMessage"}
			"getHouseholdOpinionFeedback":                       nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdOpinionFeedback"}
			"getHouseholdPsInstallerByUserId":                   nullEndPoint.Init(apiRoot), // "/v1/userService/getHouseholdPsInstallerByUserId"}
			"getHouseholdStoragePsReport":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getHouseholdStoragePsReport"}
			"getHouseholdUserInfo":                              nullEndPoint.Init(apiRoot), // "/v1/userService/getHouseholdUserInfo"}
			"getHouseholdWorkOrderInfo":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdWorkOrderInfo"}
			"getHouseholdWorkOrderList":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getHouseholdWorkOrderList"}
			"getI18nConfigByType":                               nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nConfigByType"}
			"getI18nFileInfo":                                   nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nFileInfo"}
			"getI18nInfoByKey":                                  nullEndPoint.Init(apiRoot), // "/integrationService/i18nfile/getI18nInfoByKey"}
			"getI18nVersion":                                    nullEndPoint.Init(apiRoot), // "/integrationService/international/getI18nVersion"}
			"getIncomeSettingInfos":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getIncomeSettingInfos"}
			"getInfoFromAMap":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/getInfoFromAMap"}
			"getInfomationFromRedis":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getInfomationFromRedis"}
			"getInstallInfoList":                                nullEndPoint.Init(apiRoot), // "/v1/orgService/getInstallInfoList"}
			"getInstallerInfoByDealerOrgCodeOrId":               nullEndPoint.Init(apiRoot), // "/v1/orgService/getInstallerInfoByDealerOrgCodeOrId"}
			"getInvertDataList":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getInvertDataList"}
			"getInverterDataCount":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterDataCount"}
			"getInverterProcess":                                nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterProcess"}
			"getInverterUuidBytotalId":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getInverterUuidBytotalId"}
			"getIvEchartsData":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getIvEchartsData"}
			"getIvEchartsDataById":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getIvEchartsDataById"}
			"getKpiInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getKpiInfo"}
			"getListMiFromHBase":                                nullEndPoint.Init(apiRoot), // "/v1/commonService/getListMiFromHBase"}
			"getMapInfo":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getMapInfo"}
			"getMapMiFromHBase":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getMapMiFromHBase"}
			"getMenuAndPrivileges":                              nullEndPoint.Init(apiRoot), // "/v1/userService/getMenuAndPrivileges"}
			"getMicrogridEStoragePsReport":                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getMicrogridEStoragePsReport"}
			"getModuleLogInfo":                                  nullEndPoint.Init(apiRoot), // "/integrationService/getModuleLogInfo"}
			"getModuleLogTaskList":                              nullEndPoint.Init(apiRoot), // "/integrationService/getModuleLogTaskList"}
			"getNationProvJSON":                                 nullEndPoint.Init(apiRoot), // "/v1/commonService/getNationProvJSON"}
			"getNeedOpAsynOpRecordList":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getNeedOpAsynOpRecordList"}
			"getNoticeInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/otherService/getNoticeInfo"}
			"getNumberOfServiceCalls":                           nullEndPoint.Init(apiRoot), // "/v1/commonService/getNumberOfServiceCalls"}
			"getOSSConfig":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/getOSSConfig"}
			"getOperRuleDetail":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperRuleDetail"}
			"getOperateBillFileId":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperateBillFileId"}
			"getOperateTicketForDetail":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/getOperateTicketForDetail"}
			"getOrCreateNetEaseUserToken":                       nullEndPoint.Init(apiRoot), // "/v1/userService/getOrCreateNetEaseUserToken"}
			"getOrderDataList":                                  nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderDataList"}
			"getOrderDataSql2":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getOrderDataSql2"}
			"getOrderDatas":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/getOrderDatas"}
			"getOrderDetail":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderDetail"}
			"getOrderStatistics":                                nullEndPoint.Init(apiRoot), // "/v1/faultService/getOrderStatistics"}
			"getOrgIdNameByUserId":                              nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgIdNameByUserId"}
			"getOrgInfoByDealerOrgCode":                         nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgInfoByDealerOrgCode"}
			"getOrgListByName":                                  nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgListByName"}
			"getOrgListByUserId":                                nullEndPoint.Init(apiRoot), // "/v1/userService/getOrgListByUserId"}
			"getOrgListForUser":                                 nullEndPoint.Init(apiRoot), // "/v1/orgService/getOrgListForUser"}
			"getOssObjectStream":                                nullEndPoint.Init(apiRoot), // "/v1/commonService/getOssObjectStream"}
			"getOwnerFaultConfigList":                           nullEndPoint.Init(apiRoot), // "/v1/faultService/getOwnerFaultConfigList"}
			"getPListinfoFromMysql":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPListinfoFromMysql"}
			"getParamSetTemplate4NewProtocol":                   nullEndPoint.Init(apiRoot), // "/v1/devService/getParamSetTemplate4NewProtocol"}
			"getParamSetTemplatePointInfo":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getParamSetTemplatePointInfo"}
			"getParamterSettingBase":                            nullEndPoint.Init(apiRoot), // "/v1/devService/getParamterSettingBase"}
			"getPhotoInfo":                                      nullEndPoint.Init(apiRoot), // "/v1/otherService/getPhotoInfo"}
			"getPlanedOrNotPsList":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/getPlanedOrNotPsList"}
			"getPlantReportPDFList":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPlantReportPDFList"}
			"getPowerChargeSettingInfo":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerChargeSettingInfo"}
			"getPowerDeviceModelTechList":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceModelTechList"}
			"getPowerDeviceModelTree":                           nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceModelTree"}
			"getPowerDevicePointInfo":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/getPowerDevicePointInfo"}
			api.GetName(getPowerDevicePointNames.EndPoint{}): getPowerDevicePointNames.Init(apiRoot),
			"getPowerDeviceSetTaskDetailList":                   nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceSetTaskDetailList"}
			"getPowerDeviceSetTaskList":                         nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerDeviceSetTaskList"}
			"getPowerFormulaFaultAnalyse":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerFormulaFaultAnalyse"}
			"getPowerPictureList":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerPictureList"}
			"getPowerRobotInfoByRobotSn":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotInfoByRobotSn"}
			"getPowerRobotSweepAttrByPsId":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepAttrByPsId"}
			"getPowerRobotSweepStrategy":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepStrategy"}
			"getPowerRobotSweepStrategyList":                    nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerRobotSweepStrategyList"}
			"getPowerSettingCharges":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerSettingCharges"}
			"getPowerSettingHistoryRecords":                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerSettingHistoryRecords"}
			"getPowerStationBasicInfo":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationBasicInfo"}
			"getPowerStationData":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationData"}
			"getPowerStationForHousehold":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationForHousehold"}
			"getPowerStationInfo":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationInfo"}
			"getPowerStationPR":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerStationPR"}
			"getPowerStationTableDataSql":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerStationTableDataSql"}
			"getPowerStationTableDataSqlCount":                  nullEndPoint.Init(apiRoot), // "/v1/devService/getPowerStationTableDataSqlCount"}
			api.GetName(getPowerStatistics.EndPoint{}): getPowerStatistics.Init(apiRoot),
			"getPowerTrendDayData":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPowerTrendDayData"}
			"getPrivateCloudValidityPeriod":                     nullEndPoint.Init(apiRoot), // "/v1/commonService/getPrivateCloudValidityPeriod"}
			"getProvInfoListByNationCode":                       nullEndPoint.Init(apiRoot), // "/v1/commonService/getProvInfoListByNationCode"}
			"getPsAuthKey":                                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsAuthKey"}
			"getPsCurveInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/getPsCurveInfo"}
			"getPsDataSupplementTaskList":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDataSupplementTaskList"}
			"getPsDetail":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetail"}
			"getPsDetailByUserTokens":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetailByUserTokens"}
			"getPsDetailForSinglePage":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsDetailForSinglePage"}
			api.GetName(getPsDetailWithPsType.EndPoint{}): getPsDetailWithPsType.Init(apiRoot),
			"getPsHealthState":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsHealthState"}
			"getPsInstallerByPsId":                              nullEndPoint.Init(apiRoot), // "/v1/orgService/getPsInstallerByPsId"}
			"getPsInstallerOrgInfoByPsId":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsInstallerOrgInfoByPsId"}
			api.GetName(getPsList.EndPoint{}): getPsList.Init(apiRoot),
			"getPsListByName":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListByName"}
			"getPsListForPsDataByPsId":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListForPsDataByPsId"}
			"getPsListStaticData":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsListStaticData"}
			"getPsReport":                                       nullEndPoint.Init(apiRoot), // "/v1/reportService/getPsReport"}
			"getPsUser":                                         nullEndPoint.Init(apiRoot), // "/v1/userService/getPsUser"}
			"getPsWeatherList":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getPsWeatherList"}
			"getRechargeOrderDetail":                            nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderDetail"}
			"getRechargeOrderItemDeviceList":                    nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderItemDeviceList"}
			"getRechargeOrderList":                              nullEndPoint.Init(apiRoot), // "/onlinepay/getRechargeOrderList"}
			"getRegionalTree":                                   nullEndPoint.Init(apiRoot), // "/v1/orgService/getRegionalTree"}
			"getRemoteParamSettingList":                         nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteParamSettingList"}
			"getRemoteUpgradeDeviceList":                        nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeDeviceList"}
			"getRemoteUpgradeScheduleDetails":                   nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeScheduleDetails"}
			"getRemoteUpgradeSubTasksList":                      nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeSubTasksList"}
			"getRemoteUpgradeTaskList":                          nullEndPoint.Init(apiRoot), // "/v1/devService/getRemoteUpgradeTaskList"}
			"getReportData":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getReportData"}
			"getReportEmailConfigInfo":                          nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportEmailConfigInfo"}
			"getReportExportColumns":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportExportColumns"}
			"getReportListByUserId":                             nullEndPoint.Init(apiRoot), // "/v1/reportService/getReportListByUserId"}
			"getRobotDynamicCleaningView":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getRobotDynamicCleaningView"}
			"getRobotNumAndSweepCapacity":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getRobotNumAndSweepCapacity"}
			"getRuleUnit":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getRuleUnit"}
			"getSendReportConfigCron":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/getSendReportConfigCron"}
			"getSerialNum":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/getSerialNum"}
			"getShieldMapConditionList":                         nullEndPoint.Init(apiRoot), // "/v1/commonService/getShieldMapConditionList"}
			"getSimIdBySnList":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/getSimIdBySnList"}
			"getSingleIVData":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getSingleIVData"}
			"getSnChangeRecord":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getSnChangeRecord"}
			"getSnConnectionInfo":                               nullEndPoint.Init(apiRoot), // "/v1/commonService/getSnConnectionInfo"}
			"getStationInfoSql":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/getStationInfoSql"}
			"getSungwsConfigCache":                              nullEndPoint.Init(apiRoot), // "/v1/commonService/getSungwsConfigCache"}
			"getSungwsGlobalConfigCache":                        nullEndPoint.Init(apiRoot), // "/v1/commonService/getSungwsGlobalConfigCache"}
			"getSweepDevParamSetTemplate":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getSweepDevParamSetTemplate"}
			"getSweepRobotDevList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getSweepRobotDevList"}
			"getSysMsg":                                         nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysMsg"}
			"getSysOrgNewList":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysOrgNewList"}
			"getSysOrgNewOne":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/getSysOrgNewOne"}
			"getSysUserById":                                    nullEndPoint.Init(apiRoot), // "/v1/userService/getSysUserById"}
			"getTableDataSql":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getTableDataSql"}
			"getTableDataSqlCount":                              nullEndPoint.Init(apiRoot), // "/v1/devService/getTableDataSqlCount"}
			"getTemplateByInfoType":                             nullEndPoint.Init(apiRoot), // "/v1/messageService/getTemplateByInfoType"}
			"getTemplateList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getTemplateList"}
			"getUUIDByUpuuid":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/getUUIDByUpuuid"}
			"getUpTimePoint":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/getUpTimePoint"}
			"getUserById":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/getUserById"}
			"getUserByInstaller":                                nullEndPoint.Init(apiRoot), // "/v1/userService/getUserByInstaller"}
			"getUserDevOnlineOffineCount":                       nullEndPoint.Init(apiRoot), // "/v1/devService/getUserDevOnlineOffineCount"}
			"getUserGDPRAttrs":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/getUserGDPRAttrs"}
			"getUserHavePowerStationCount":                      nullEndPoint.Init(apiRoot), // "/v1/userService/getUserHavePowerStationCount"}
			"getUserInfoByUserAccounts":                         nullEndPoint.Init(apiRoot), // "/v1/userService/getUserInfoByUserAccounts"}
			"getUserList":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/getUserList"}
			"getUserPsOrderList":                                nullEndPoint.Init(apiRoot), // "/v1/faultService/getUserPsOrderList"}
			"getValFromHBase":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/getValFromHBase"}
			"getValidateCode":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/getValidateCode"}
			"getValidateCodeAtRegister":                         nullEndPoint.Init(apiRoot), // "/v1/userService/getValidateCodeAtRegister"}
			"getWeatherInfo":                                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/getWeatherInfo"}
			"getWechatPushConfig":                               nullEndPoint.Init(apiRoot), // "/v1/userService/getWechatPushConfig"}
			"getWorkInfo":                                       nullEndPoint.Init(apiRoot), // "/v1/otherService/getWorkInfo"}
			"groupStringCheck":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/groupStringCheck"}
			"handleDevByCommunicationSN":                        nullEndPoint.Init(apiRoot), // "/devDataHandleService/handleDevByCommunicationSN"}
			"householdResetPassBySN":                            nullEndPoint.Init(apiRoot), // "/v1/userService/householdResetPassBySN"}
			"immediatePayment":                                  nullEndPoint.Init(apiRoot), // "/onlinepay/immediatePayment"}
			"importExcelData":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/importExcelData"}
			"incomeStatistics":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/incomeStatistics"}
			"informPush":                                        nullEndPoint.Init(apiRoot), // "/v1/messageService/informPush"}
			"insertEmgOrgInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/insertEmgOrgInfo"}
			"insightSynDeviceStructure2Cloud":                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/insightSynDeviceStructure2Cloud"}
			"intoDataToHbase":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/intoDataToHbase"}
			"ipLocationQuery":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/ipLocationQuery"}
			"isHave2GSn":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/isHave2GSn"}
			"judgeDevIsHasInitSetTemplate":                      nullEndPoint.Init(apiRoot), // "/v1/devService/judgeDevIsHasInitSetTemplate"}
			"judgeIsSettingMan":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/judgeIsSettingMan"}
			"listOssFiles":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/listOssFiles"}
			"loadAreaInfo":                                      nullEndPoint.Init(apiRoot), // "/v1/commonService/loadAreaInfo"}
			"loadPowerStation":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/loadPowerStation"}
			api.GetName(login.EndPoint{}): login.Init(apiRoot),
			"loginByToken":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/loginByToken"}
			"logout":                                            nullEndPoint.Init(apiRoot), // "/v1/userService/logout"}
			"mobilePhoneHasBound":                               nullEndPoint.Init(apiRoot), // "/v1/userService/mobilePhoneHasBound"}
			"modifiedDeviceInfo":                                nullEndPoint.Init(apiRoot), // "/v1/devService/modifiedDeviceInfo"}
			"modifyEmgOrgStruc":                                 nullEndPoint.Init(apiRoot), // "/v1/otherService/modifyEmgOrgStruc"}
			"modifyFaultPlan":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/modifyFaultPlan"}
			"modifyOnDutyInfo":                                  nullEndPoint.Init(apiRoot), // "/v1/otherService/modifyOnDutyInfo"}
			"modifyPassword":                                    nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPassword"}
			"modifyPersonalUnitList":                            nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPersonalUnitList"}
			"modifyPsUser":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/modifyPsUser"}
			"moduleLogParamSet":                                 nullEndPoint.Init(apiRoot), // "/integrationService/moduleLogParamSet"}
			"operateOssFile":                                    nullEndPoint.Init(apiRoot), // "/v1/commonService/operateOssFile"}
			"operationPowerRobotSweepStrategy":                  nullEndPoint.Init(apiRoot), // "/v1/devService/operationPowerRobotSweepStrategy"}
			"orgPowerReport":                                    nullEndPoint.Init(apiRoot), // "/v1/orgService/orgPowerReport"}
			"paramSetTryAgain":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/paramSetTryAgain"}
			"paramSetting":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/paramSetting"}
			"planPower":                                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/planPower"}
			"powerDevicePointList":                              nullEndPoint.Init(apiRoot), // "/v1/reportService/powerDevicePointList"}
			"powerTrendChartData":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/powerTrendChartData"}
			"psForcastInfo":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/psForcastInfo"}
			"psHourPointsValue":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/psHourPointsValue"}
			"queryAllPsIdAndName":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryAllPsIdAndName"}
			"queryBatchCreatePsTaskList":                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryBatchCreatePsTaskList"}
			"queryBatchSpeedyAddPowerStationResult":             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryBatchSpeedyAddPowerStationResult"}
			"queryCardStatusCTCC":                               nullEndPoint.Init(apiRoot), // "/v1/devService/queryCardStatusCTCC"}
			"queryChildAccountList":                             nullEndPoint.Init(apiRoot), // "/v1/userService/queryChildAccountList"}
			"queryCompensationRecordData":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryCompensationRecordData"}
			"queryCompensationRecordList":                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryCompensationRecordList"}
			"queryComponent":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/queryComponent"}
			"queryComponentTechnicalParam":                      nullEndPoint.Init(apiRoot), // "/v1/devService/queryComponentTechnicalParam"}
			"queryCountryGridAndRelation":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryCountryGridAndRelation"}
			"queryCountryList":                                  nullEndPoint.Init(apiRoot), // "/v1/commonService/queryCountryList"}
			"queryCtrlTaskById":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryCtrlTaskById"}
			"queryDeviceInfo":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceInfoForApp"}
			"queryDeviceInfoForApp":                             nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceInfoForApp"}
			"queryDeviceList":                                   nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceList"}
			"queryDeviceListByUserId":                           nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceListByUserId"}
			"queryDeviceListForApp":                             nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceListForApp"}
			"queryDeviceModelTechnical":                         nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceModelTechnical"}
			"queryDevicePointDayMonthYearDataList":              nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointDayMonthYearDataList"}
			"queryDevicePointMinuteDataList":                    nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointMinuteDataList"}
			"queryDevicePointsDayMonthYearDataList":             nullEndPoint.Init(apiRoot), // "/v1/commonService/queryDevicePointsDayMonthYearDataList"}
			"queryDeviceRealTimeDataByPsKeys":                   nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceRealTimeDataByPsKeys"}
			"queryDeviceRepairList":                             nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceRepairList"}
			"queryDeviceTypeInfoList":                           nullEndPoint.Init(apiRoot), // "/v1/devService/queryDeviceTypeInfoList"}
			"queryEnvironmentList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryEnvironmentList"}
			"queryFaultList":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultList"}
			"queryFaultPlanDetail":                              nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultPlanDetail"}
			"queryFaultRepairSteps":                             nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultRepairSteps"}
			"queryFaultTypeAndLevelByCode":                      nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeAndLevelByCode"}
			"queryFaultTypeByDevice":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeByDevice"}
			"queryFaultTypeByDevicePage":                        nullEndPoint.Init(apiRoot), // "/v1/faultService/queryFaultTypeByDevicePage"}
			"queryFirmwareFilesPage":                            nullEndPoint.Init(apiRoot), // "/v1/commonService/queryFirmwareFilesPage"}
			"queryInfotoAlert":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/queryInfotoAlert"}
			"queryInverterModelList":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryInverterModelList"}
			"queryInverterVersionList":                          nullEndPoint.Init(apiRoot), // "/v1/devService/queryInverterVersionList"}
			"queryM2MCardInfoCMCC":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryM2MCardInfoCMCC"}
			"queryM2MCardTermInfoCMCC":                          nullEndPoint.Init(apiRoot), // "/v1/devService/queryM2MCardTermInfoCMCC"}
			"queryModelInfoByModelId":                           nullEndPoint.Init(apiRoot), // "/v1/devService/queryModelInfoByModelId"}
			"queryMutiPointDataList":                            nullEndPoint.Init(apiRoot), // "/v1/commonService/queryMutiPointDataList"}
			"queryNoticeList":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/queryNoticeList"}
			"queryNumberOfRenewalReminders":                     nullEndPoint.Init(apiRoot), // "/v1/devService/queryNumberOfRenewalReminders"}
			"queryOperRules":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOperRules"}
			"queryOrderList":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOrderList"}
			"queryOrderStep":                                    nullEndPoint.Init(apiRoot), // "/v1/faultService/queryOrderStep"}
			"queryOrgGenerationReport":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/queryOrgGenerationReport"}
			"queryOrgInfoList":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/queryOrgInfoList"}
			"queryOrgPowerElecPercent":                          nullEndPoint.Init(apiRoot), // "/v1/orgService/queryOrgPowerElecPercent"}
			"queryOrgPsCompensationRecordList":                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryOrgPsCompensationRecordList"}
			"queryParamSettingTask":                             nullEndPoint.Init(apiRoot), // "/v1/devService/queryParamSettingTask"}
			"queryPersonalUnitList":                             nullEndPoint.Init(apiRoot), // "/v1/userService/queryPersonalUnitList"}
			"queryPointDataTopOne":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryPointDataTopOne"}
			"queryPowerStationInfo":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPowerStationInfo"}
			"queryPsAreaByUserIdAndAreaCode":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsAreaByUserIdAndAreaCode"}
			"queryPsCompensationRecordList":                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsCompensationRecordList"}
			"queryPsDataByDate":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsDataByDate"}
			"queryPsIdList":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsIdList"}
			"queryPsListByUserIdAndAreaCode":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsListByUserIdAndAreaCode"}
			"queryPsNameByPsId":                                 nullEndPoint.Init(apiRoot), // "/v1/devService/queryPsNameByPsId"}
			"queryPsPrByDate":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsPrByDate"}
			"queryPsProfit":                                     nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsProfit"}
			"queryPsReportComparativeAnalysisOfPowerGeneration": nullEndPoint.Init(apiRoot), // "/v1/powerStationService/queryPsReportComparativeAnalysisOfPowerGeneration"}
			"queryPsStructureList":                              nullEndPoint.Init(apiRoot), // "/v1/devService/queryPsStructureList"}
			"queryPuuidsByCommandTotalId":                       nullEndPoint.Init(apiRoot), // "/v1/devService/queryPuuidsByCommandTotalId"}
			"queryPuuidsByCommandTotalId2":                      nullEndPoint.Init(apiRoot), // "/v1/devService/queryPuuidsByCommandTotalId2"}
			"queryRepairRuleList":                               nullEndPoint.Init(apiRoot), // "/v1/devService/queryRepairRuleList"}
			"queryReportListForManagementPage":                  nullEndPoint.Init(apiRoot), // "/v1/reportService/queryReportListForManagementPage"}
			"queryReportMsg":                                    nullEndPoint.Init(apiRoot), // "/v1/reportService/queryReportMsg"}
			"querySharingPs":                                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/querySharingPs"}
			"querySysAdvancedParam":                             nullEndPoint.Init(apiRoot), // "/v1/devService/querySysAdvancedParam"}
			"queryTimeBySN":                                     nullEndPoint.Init(apiRoot), // "/v1/devService/queryTimeBySN"}
			"queryTrafficByDateCTCC":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryTrafficByDateCTCC"}
			"queryTrafficCTCC":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/queryTrafficCTCC"}
			"queryUnitList":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/queryUnitList"}
			"queryUnitUuidBytotalId":                            nullEndPoint.Init(apiRoot), // "/v1/devService/queryUnitUuidBytotalId"}
			"queryUserBtnPri":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserBtnPri"}
			"queryUserByUserIds":                                nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserByUserIds"}
			"queryUserExtensionAttribute":                       nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserExtensionAttribute"}
			"queryUserForStep":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserForStep"}
			"queryUserList":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserList"}
			"queryUserProcessPri":                               nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserProcessPri"}
			"queryUserWechatBindRel":                            nullEndPoint.Init(apiRoot), // "/v1/userService/queryUserWechatBindRel"}
			"queryUuidByTotalIdAndUuid":                         nullEndPoint.Init(apiRoot), // "/v1/devService/queryUuidByTotalIdAndUuid"}
			"rechargeOrderSetMeal":                              nullEndPoint.Init(apiRoot), // "/v1/devService/rechargeOrderSetMeal"}
			"renewSendReportConfirmEmail":                       nullEndPoint.Init(apiRoot), // "/v1/reportService/renewSendReportConfirmEmail"}
			"reportList":                                        nullEndPoint.Init(apiRoot), // "/v1/powerStationService/reportList"}
			"saveCustomerEmployee":                              nullEndPoint.Init(apiRoot), // "/v1/devService/saveCustomerEmployee"}
			"saveDevSimList":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/saveDevSimList"}
			"saveDeviceAccountBatchData":                        nullEndPoint.Init(apiRoot), // "/v1/devService/saveDeviceAccountBatchData"}
			"saveEnviromentIncomeInfos":                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/saveEnviromentIncomeInfos"}
			"saveEnvironmentCurve":                              nullEndPoint.Init(apiRoot), // "/v1/devService/saveEnvironmentCurve"}
			"saveFirmwareFile":                                  nullEndPoint.Init(apiRoot), // "/v1/commonService/saveFirmwareFile"}
			"saveIncomeSettingInfos":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/saveIncomeSettingInfos"}
			"saveOrUpdateGroupStringCheckRule":                  nullEndPoint.Init(apiRoot), // "/v1/devService/saveOrUpdateGroupStringCheckRule"}
			"saveParamModel":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/saveParamModel"}
			"savePowerCharges":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerCharges"}
			"savePowerDevicePoint":                              nullEndPoint.Init(apiRoot), // "/v1/reportService/savePowerDevicePoint"}
			"savePowerRobotInfo":                                nullEndPoint.Init(apiRoot), // "/v1/devService/savePowerRobotInfo"}
			"savePowerRobotSweepAttr":                           nullEndPoint.Init(apiRoot), // "/v1/devService/savePowerRobotSweepAttr"}
			"savePowerSettingCharges":                           nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerSettingCharges"}
			"savePowerSettingInfo":                              nullEndPoint.Init(apiRoot), // "/v1/powerStationService/savePowerSettingInfo"}
			"saveProductionBatchData":                           nullEndPoint.Init(apiRoot), // "/v1/devService/saveProductionBatchData"}
			"saveRechargeOrderObj":                              nullEndPoint.Init(apiRoot), // "/onlinepay/saveRechargeOrderObj"}
			"saveRechargeOrderOtherInfo":                        nullEndPoint.Init(apiRoot), // "/onlinepay/saveRechargeOrderOtherInfo"}
			"saveRepair":                                        nullEndPoint.Init(apiRoot), // "/v1/faultService/saveRepair"}
			"saveReportExportColumns":                           nullEndPoint.Init(apiRoot), // "/v1/reportService/saveReportExportColumns"}
			"saveSetParam":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/saveSetParam"}
			"saveSysUserMsg":                                    nullEndPoint.Init(apiRoot), // "/v1/otherService/saveSysUserMsg"}
			"saveTemplate":                                      nullEndPoint.Init(apiRoot), // "/v1/devService/saveTemplate"}
			"searchM2MMonthFlowCMCC":                            nullEndPoint.Init(apiRoot), // "/v1/devService/searchM2MMonthFlowCMCC"}
			"selectSysTranslationNames":                         nullEndPoint.Init(apiRoot), // "/v1/reportService/selectSysTranslationNames"}
			"sendPsTimeZoneInstruction":                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/sendPsTimeZoneInstruction"}
			"setUpFormulaFaultAnalyse":                          nullEndPoint.Init(apiRoot), // "/v1/powerStationService/setUpFormulaFaultAnalyse"}
			"setUserGDPRAttrs":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/setUserGDPRAttrs"}
			"settingNotice":                                     nullEndPoint.Init(apiRoot), // "/v1/userService/settingNotice"}
			"shareMyPs":                                         nullEndPoint.Init(apiRoot), // "/v1/powerStationService/shareMyPs"}
			"sharePsBySN":                                       nullEndPoint.Init(apiRoot), // "/v1/powerStationService/sharePsBySN"}
			"showInverterByUnit":                                nullEndPoint.Init(apiRoot), // "/v1/devService/showInverterByUnit"}
			"showOnlineUsers":                                   nullEndPoint.Init(apiRoot), // "/v1/userService/showOnlineUsers"}
			"showWarning":                                       nullEndPoint.Init(apiRoot), // "/v1/userService/showWarning"}
			"snIsExist":                                         nullEndPoint.Init(apiRoot), // "/v1/devService/snIsExist"}
			"snsIsExist":                                        nullEndPoint.Init(apiRoot), // "/v1/devService/snsIsExist"}
			"speedyAddPowerStation":                             nullEndPoint.Init(apiRoot), // "/v1/powerStationService/speedyAddPowerStation"}
			"stationDeviceHistoryDataList":                      nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationDeviceHistoryDataList"}
			"stationUnitsList":                                  nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationUnitsList"}
			"stationsDiscreteData":                              nullEndPoint.Init(apiRoot), // "/v1/devService/stationsDiscreteData"}
			"stationsIncomeList":                                nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationsIncomeList"}
			"stationsPointReport":                               nullEndPoint.Init(apiRoot), // "/v1/powerStationService/stationsPointReport"}
			"stationsYearPlanReport":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/stationsYearPlanReport"}
			"sureAndImportSelettlementData":                     nullEndPoint.Init(apiRoot), // "/v1/otherService/sureAndImportSelettlementData"}
			"sweepDevParamSet":                                  nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevParamSet"}
			"sweepDevRunControl":                                nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevRunControl"}
			"sweepDevStrategyIssue":                             nullEndPoint.Init(apiRoot), // "/v1/devService/sweepDevStrategyIssue"}
			"sysTimeZoneList":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/sysTimeZoneList"}
			"unLockUser":                                        nullEndPoint.Init(apiRoot), // "/v1/userService/unLockUser"}
			"unlockChildAccount":                                nullEndPoint.Init(apiRoot), // "/v1/userService/unlockChildAccount"}
			"updateCommunicationModuleState":                    nullEndPoint.Init(apiRoot), // "/v1/devService/updateCommunicationModuleState"}
			"updateDevInstalledPower":                           nullEndPoint.Init(apiRoot), // "/v1/devService/updateDevInstalledPower"}
			"updateFault":                                       nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultStatus"}
			"updateFaultData":                                   nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultData"}
			"updateFaultMsgByFaultCode":                         nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultMsgByFaultCode"}
			"updateFaultStatus":                                 nullEndPoint.Init(apiRoot), // "/v1/faultService/updateFaultStatus"}
			"updateHouseholdWorkOrder":                          nullEndPoint.Init(apiRoot), // "/v1/faultService/updateHouseholdWorkOrder"}
			"updateInverterSn2ModuleSn":                         nullEndPoint.Init(apiRoot), // "/devDataHandleService/updateInverterSn2ModuleSn"}
			"updateOperateTicketAttachmentId":                   nullEndPoint.Init(apiRoot), // "/v1/faultService/updateOperateTicketAttachmentId"}
			"updateOrderDeviceByCustomerService":                nullEndPoint.Init(apiRoot), // "/onlinepay/updateOrderDeviceByCustomerService"}
			"updateOwnerFaultConfig":                            nullEndPoint.Init(apiRoot), // "/v1/faultService/updateOwnerFaultConfig"}
			"updateParamSettingSysMsg":                          nullEndPoint.Init(apiRoot), // "/v1/devService/updateParamSettingSysMsg"}
			"updatePlatformLevelFaultLevel":                     nullEndPoint.Init(apiRoot), // "/v1/faultService/updatePlatformLevelFaultLevel"}
			"updatePowerDevicePoint":                            nullEndPoint.Init(apiRoot), // "/v1/reportService/updatePowerDevicePoint"}
			"updatePowerRobotInfo":                              nullEndPoint.Init(apiRoot), // "/v1/devService/updatePowerRobotInfo"}
			"updatePowerRobotSweepAttr":                         nullEndPoint.Init(apiRoot), // "/v1/devService/updatePowerRobotSweepAttr"}
			"updatePowerStationForHousehold":                    nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updatePowerStationForHousehold"}
			"updatePowerStationInfo":                            nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updatePowerStationInfo"}
			"updatePowerUserInfo":                               nullEndPoint.Init(apiRoot), // "/v1/userService/updatePowerUserInfo"}
			"updateReportConfigByEmailAddr":                     nullEndPoint.Init(apiRoot), // "/v1/reportService/updateReportConfigByEmailAddr"}
			"updateShareAttr":                                   nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updateShareAttr"}
			"updateSnIsSureFlag":                                nullEndPoint.Init(apiRoot), // "/devDataHandleService/updateSnIsSureFlag"}
			"updateStationPics":                                 nullEndPoint.Init(apiRoot), // "/v1/powerStationService/updateStationPics"}
			"updateSysAdvancedParam":                            nullEndPoint.Init(apiRoot), // "/v1/devService/updateSysAdvancedParam"}
			"updateSysOrgNew":                                   nullEndPoint.Init(apiRoot), // "/v1/otherService/updateSysOrgNew"}
			"updateTemplate":                                    nullEndPoint.Init(apiRoot), // "/v1/devService/updateDataCurveTemplate"}
			"updateUinfoNetEaseUser":                            nullEndPoint.Init(apiRoot), // "/v1/userService/updateUinfoNetEaseUser"}
			"updateUserExtensionAttribute":                      nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserExtensionAttribute"}
			"updateUserLanguage":                                nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserLanguage"}
			"updateUserPosition":                                nullEndPoint.Init(apiRoot), // "/v1/userService/updateUserPosition"}
			"updateUserUpOrg":                                   nullEndPoint.Init(apiRoot), // "/v1/orgService/updateUserUpOrg"}
			"upgrade":                                           nullEndPoint.Init(apiRoot), // "/v1/userService/upgrade"}
			"upgrate":                                           nullEndPoint.Init(apiRoot), // "/v1/userService/upgrade"}
			"uploadFileToOss":                                   nullEndPoint.Init(apiRoot), // "/v1/commonService/uploadFileToOss"}
			"userAgreeGdprProtocol":                             nullEndPoint.Init(apiRoot), // "/v1/userService/userAgreeGdprProtocol"}
			"userInfoUniqueCheck":                               nullEndPoint.Init(apiRoot), // "/v1/userService/userInfoUniqueCheck"}
			"userMailHasBound":                                  nullEndPoint.Init(apiRoot), // "/v1/userService/userMailHasBound"}
			"userRegister":                                      nullEndPoint.Init(apiRoot), // "/v1/userService/userRegister"}
		},
	}

	// fmt.Printf("area[%s]: %v\n", name, area)

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
