// API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebIscmAppService

import (
	"GoSungrow/iSolarCloud/WebIscmAppService/addPowerDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/addPowerPointManage"
	"GoSungrow/iSolarCloud/WebIscmAppService/addSubTypeDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/batchAddDevicesPropertis"
	"GoSungrow/iSolarCloud/WebIscmAppService/batchDelDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/batchSavePowerDeviceTechnical"
	"GoSungrow/iSolarCloud/WebIscmAppService/checkDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/contactMessageOpera"
	"GoSungrow/iSolarCloud/WebIscmAppService/delDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteDeviceFactory"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteDeviceType"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteOneNotice"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteOrgNodeInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePicture"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePointInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceChannl"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceParameterPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceSubType"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceTechnical"
	"GoSungrow/iSolarCloud/WebIscmAppService/deletePowerStore"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteProcessDefinition"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteReport"
	"GoSungrow/iSolarCloud/WebIscmAppService/deleteUserNode"
	"GoSungrow/iSolarCloud/WebIscmAppService/deployProcess"
	"GoSungrow/iSolarCloud/WebIscmAppService/editProcessManageAction"
	"GoSungrow/iSolarCloud/WebIscmAppService/findImageInputStreamString"
	"GoSungrow/iSolarCloud/WebIscmAppService/getAllDevTypeList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getAllNodeByType"
	"GoSungrow/iSolarCloud/WebIscmAppService/getAuthKey"
	"GoSungrow/iSolarCloud/WebIscmAppService/getAuthKeyList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getCodeByType"
	"GoSungrow/iSolarCloud/WebIscmAppService/getContactMessage"
	"GoSungrow/iSolarCloud/WebIscmAppService/getCountryNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDefinitionIdByKey"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeploymentList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceFactoryListByIds"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDevicePro"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceSubType"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceTechnical"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceType"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDeviceTypeInfoById"
	"GoSungrow/iSolarCloud/WebIscmAppService/getDutyUserList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getFatherPrivileges"
	"GoSungrow/iSolarCloud/WebIscmAppService/getGroupManSettings"
	"GoSungrow/iSolarCloud/WebIscmAppService/getGroupManSettingsMembers"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMaterialByListId"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMaterialByType"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMaterialList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMaxDeviceIdByPsId"
	"GoSungrow/iSolarCloud/WebIscmAppService/getModelPoints"
	"GoSungrow/iSolarCloud/WebIscmAppService/getMoneyUnitList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getNamecnNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getNationList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOperationRecord"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgAndChildBasicInfoOptions"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgAndStateAndCode"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgForPs"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgListForUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgNodeInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgStationList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgStationListByPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgUserList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgUserMapData"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtree"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtree4User"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtreeAsync"
	"GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtreeForUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPictureList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPointInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPointInfoPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceChannl"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactory"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactoryListCount"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelTechList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceTypeList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerPlanList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerStation"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerStationInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerStationList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPowerStore"
	"GoSungrow/iSolarCloud/WebIscmAppService/getProvcnNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getPsTreeMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getRoleByUserIds"
	"GoSungrow/iSolarCloud/WebIscmAppService/getRootOrgInfoByUserId"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSettingUserMapData"
	"GoSungrow/iSolarCloud/WebIscmAppService/getStateNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSubTypeDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSysHomeList2"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSysMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSysOrgPro"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSysUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSystemOrgInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSystemRoleInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/getSystemRoleList2"
	"GoSungrow/iSolarCloud/WebIscmAppService/getTownValueNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getUserMenuLs"
	"GoSungrow/iSolarCloud/WebIscmAppService/getUserOrgPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/getVillageList"
	"GoSungrow/iSolarCloud/WebIscmAppService/getVillageListNew"
	"GoSungrow/iSolarCloud/WebIscmAppService/getZtreeAsyncSysMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getZtreeChildMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getZtreeMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu2"
	"GoSungrow/iSolarCloud/WebIscmAppService/goToDevicePropertyPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/isCanAddUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/isHasIrradiationData"
	"GoSungrow/iSolarCloud/WebIscmAppService/isHasPlan"
	"GoSungrow/iSolarCloud/WebIscmAppService/loadDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/modelPointsPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifyDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifyPowerDeviceChannl"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifySysOrg"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifySystemMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifySystemOrgNode"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifySystemRole"
	"GoSungrow/iSolarCloud/WebIscmAppService/modifySystemUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/publishNotice"
	"GoSungrow/iSolarCloud/WebIscmAppService/queryDeviceList"
	"GoSungrow/iSolarCloud/WebIscmAppService/queryDutyType"
	"GoSungrow/iSolarCloud/WebIscmAppService/queryReportDataById"
	"GoSungrow/iSolarCloud/WebIscmAppService/resetPasW"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveAuthKey"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveDeviceFactory"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveDeviceType"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveIrradiationData"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveModelPoints"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveNewNotice"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveOrUpdateReport"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveOrgNode"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveOrgUsers"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePicture"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePointManage"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceChannl"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceParameterPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceSubType"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceTechnical"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerPlan"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerStationByPowerStore"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePowerStore"
	"GoSungrow/iSolarCloud/WebIscmAppService/savePsOrg"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveRelDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveRoleAssign"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveSysMenu"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveSysOrg"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveSysRole"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveSysUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveUserNode"
	"GoSungrow/iSolarCloud/WebIscmAppService/saveUserRole"
	"GoSungrow/iSolarCloud/WebIscmAppService/searchIrradiationData"
	"GoSungrow/iSolarCloud/WebIscmAppService/searchTechnicalNums"
	"GoSungrow/iSolarCloud/WebIscmAppService/selectDeviceTypeByPsId"
	"GoSungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceTechnicals"
	"GoSungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceType"
	"GoSungrow/iSolarCloud/WebIscmAppService/setupUserRole4AddUser"
	"GoSungrow/iSolarCloud/WebIscmAppService/startWorkFlow"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateDevice"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateDeviceType"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateFaultLevel"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateNotice"
	"GoSungrow/iSolarCloud/WebIscmAppService/updatePointInfo"
	"GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceParameterPage"
	"GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceSubType"
	"GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceTechnical"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateProcessManage"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateSysOrgPro"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateSysRoleValidFlag"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateUserValidFlag"
	"GoSungrow/iSolarCloud/WebIscmAppService/updateValidFlag"
	"GoSungrow/iSolarCloud/WebIscmAppService/viewDeviceModel"
	"GoSungrow/iSolarCloud/WebIscmAppService/viewDeviceParameter"
	"GoSungrow/iSolarCloud/WebIscmAppService/workFlowImplementStep"
	"GoSungrow/iSolarCloud/WebIscmAppService/workFlowIsStart"
	"GoSungrow/iSolarCloud/WebIscmAppService/workFlowTransferStep"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
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
			api.GetName(addPowerDeviceModel.EndPoint{}):            addPowerDeviceModel.Init(apiRoot),
			api.GetName(addPowerPointManage.EndPoint{}):            addPowerPointManage.Init(apiRoot),
			api.GetName(addSubTypeDevice.EndPoint{}):               addSubTypeDevice.Init(apiRoot),
			api.GetName(batchAddDevicesPropertis.EndPoint{}):       batchAddDevicesPropertis.Init(apiRoot),
			api.GetName(batchDelDevice.EndPoint{}):                 batchDelDevice.Init(apiRoot),
			api.GetName(batchSavePowerDeviceTechnical.EndPoint{}):  batchSavePowerDeviceTechnical.Init(apiRoot),
			api.GetName(checkDeviceModel.EndPoint{}):               checkDeviceModel.Init(apiRoot),
			api.GetName(contactMessageOpera.EndPoint{}):            contactMessageOpera.Init(apiRoot),
			api.GetName(delDevice.EndPoint{}):                      delDevice.Init(apiRoot),
			api.GetName(deleteDeviceFactory.EndPoint{}):            deleteDeviceFactory.Init(apiRoot),
			api.GetName(deleteDeviceType.EndPoint{}):               deleteDeviceType.Init(apiRoot),
			api.GetName(deleteMenu.EndPoint{}):                     deleteMenu.Init(apiRoot),
			api.GetName(deleteOneNotice.EndPoint{}):                deleteOneNotice.Init(apiRoot),
			api.GetName(deleteOrgNodeInfo.EndPoint{}):              deleteOrgNodeInfo.Init(apiRoot),
			api.GetName(deletePicture.EndPoint{}):                  deletePicture.Init(apiRoot),
			api.GetName(deletePointInfo.EndPoint{}):                deletePointInfo.Init(apiRoot),
			api.GetName(deletePowerDeviceChannl.EndPoint{}):        deletePowerDeviceChannl.Init(apiRoot),
			api.GetName(deletePowerDeviceModel.EndPoint{}):         deletePowerDeviceModel.Init(apiRoot),
			api.GetName(deletePowerDeviceParameterPage.EndPoint{}): deletePowerDeviceParameterPage.Init(apiRoot),
			api.GetName(deletePowerDeviceSubType.EndPoint{}):       deletePowerDeviceSubType.Init(apiRoot),
			api.GetName(deletePowerDeviceTechnical.EndPoint{}):     deletePowerDeviceTechnical.Init(apiRoot),
			api.GetName(deletePowerStore.EndPoint{}):               deletePowerStore.Init(apiRoot),
			api.GetName(deleteProcessDefinition.EndPoint{}):        deleteProcessDefinition.Init(apiRoot),
			api.GetName(deleteReport.EndPoint{}):                   deleteReport.Init(apiRoot),
			api.GetName(deleteUserNode.EndPoint{}):                 deleteUserNode.Init(apiRoot),
			api.GetName(deployProcess.EndPoint{}):                  deployProcess.Init(apiRoot),
			api.GetName(editProcessManageAction.EndPoint{}):        editProcessManageAction.Init(apiRoot),
			api.GetName(findImageInputStreamString.EndPoint{}):     findImageInputStreamString.Init(apiRoot),
			api.GetName(getAllDevTypeList.EndPoint{}):              getAllDevTypeList.Init(apiRoot),
			api.GetName(getAllNodeByType.EndPoint{}):               getAllNodeByType.Init(apiRoot),
			api.GetName(getAuthKey.EndPoint{}):                     getAuthKey.Init(apiRoot),
			api.GetName(getAuthKeyList.EndPoint{}):                 getAuthKeyList.Init(apiRoot),
			api.GetName(getCodeByType.EndPoint{}):                  getCodeByType.Init(apiRoot),
			api.GetName(getContactMessage.EndPoint{}):              getContactMessage.Init(apiRoot),
			api.GetName(getCountryNew.EndPoint{}):                  getCountryNew.Init(apiRoot),
			api.GetName(getDefinitionIdByKey.EndPoint{}):           getDefinitionIdByKey.Init(apiRoot),
			api.GetName(getDeploymentList.EndPoint{}):              getDeploymentList.Init(apiRoot),
			api.GetName(getDeviceFactoryListByIds.EndPoint{}):      getDeviceFactoryListByIds.Init(apiRoot),
			api.GetName(getDeviceModel.EndPoint{}):                 getDeviceModel.Init(apiRoot),
			api.GetName(getDevicePro.EndPoint{}):                   getDevicePro.Init(apiRoot),
			api.GetName(getDeviceSubType.EndPoint{}):               getDeviceSubType.Init(apiRoot),
			api.GetName(getDeviceTechnical.EndPoint{}):             getDeviceTechnical.Init(apiRoot),
			api.GetName(getDeviceType.EndPoint{}):                  getDeviceType.Init(apiRoot),
			api.GetName(getDeviceTypeInfoById.EndPoint{}):          getDeviceTypeInfoById.Init(apiRoot),
			api.GetName(getDutyUserList.EndPoint{}):                getDutyUserList.Init(apiRoot),
			api.GetName(getFatherPrivileges.EndPoint{}):            getFatherPrivileges.Init(apiRoot),
			api.GetName(getGroupManSettings.EndPoint{}):            getGroupManSettings.Init(apiRoot),
			api.GetName(getGroupManSettingsMembers.EndPoint{}):     getGroupManSettingsMembers.Init(apiRoot),
			api.GetName(getMaterialByListId.EndPoint{}):            getMaterialByListId.Init(apiRoot),
			api.GetName(getMaterialByType.EndPoint{}):              getMaterialByType.Init(apiRoot),
			api.GetName(getMaterialList.EndPoint{}):                getMaterialList.Init(apiRoot),
			api.GetName(getMaxDeviceIdByPsId.EndPoint{}):           getMaxDeviceIdByPsId.Init(apiRoot),
			api.GetName(getModelPoints.EndPoint{}):                 getModelPoints.Init(apiRoot),
			api.GetName(getMoneyUnitList.EndPoint{}):               getMoneyUnitList.Init(apiRoot),
			api.GetName(getNamecnNew.EndPoint{}):                   getNamecnNew.Init(apiRoot),
			api.GetName(getNationList.EndPoint{}):                  getNationList.Init(apiRoot),
			api.GetName(getOperationRecord.EndPoint{}):             getOperationRecord.Init(apiRoot),
			api.GetName(getOrgAndChildBasicInfoOptions.EndPoint{}): getOrgAndChildBasicInfoOptions.Init(apiRoot),
			api.GetName(getOrgAndStateAndCode.EndPoint{}):          getOrgAndStateAndCode.Init(apiRoot),
			api.GetName(getOrgForPs.EndPoint{}):                    getOrgForPs.Init(apiRoot),
			api.GetName(getOrgList.EndPoint{}):                     getOrgList.Init(apiRoot),
			api.GetName(getOrgListForUser.EndPoint{}):              getOrgListForUser.Init(apiRoot),
			api.GetName(getOrgNodeInfo.EndPoint{}):                 getOrgNodeInfo.Init(apiRoot),
			api.GetName(getOrgStationList.EndPoint{}):              getOrgStationList.Init(apiRoot),
			api.GetName(getOrgStationListByPage.EndPoint{}):        getOrgStationListByPage.Init(apiRoot),
			api.GetName(getOrgUserList.EndPoint{}):                 getOrgUserList.Init(apiRoot),
			api.GetName(getOrgUserMapData.EndPoint{}):              getOrgUserMapData.Init(apiRoot),
			api.GetName(getOrgZtree.EndPoint{}):                    getOrgZtree.Init(apiRoot),
			api.GetName(getOrgZtree4User.EndPoint{}):               getOrgZtree4User.Init(apiRoot),
			api.GetName(getOrgZtreeAsync.EndPoint{}):               getOrgZtreeAsync.Init(apiRoot),
			api.GetName(getOrgZtreeForUser.EndPoint{}):             getOrgZtreeForUser.Init(apiRoot),
			api.GetName(getPictureList.EndPoint{}):                 getPictureList.Init(apiRoot),
			api.GetName(getPointInfo.EndPoint{}):                   getPointInfo.Init(apiRoot),
			api.GetName(getPointInfoPage.EndPoint{}):               getPointInfoPage.Init(apiRoot),
			api.GetName(getPowerDevice.EndPoint{}):                 getPowerDevice.Init(apiRoot),
			api.GetName(getPowerDeviceChannl.EndPoint{}):           getPowerDeviceChannl.Init(apiRoot),
			api.GetName(getPowerDeviceFactory.EndPoint{}):          getPowerDeviceFactory.Init(apiRoot),
			api.GetName(getPowerDeviceFactoryListCount.EndPoint{}): getPowerDeviceFactoryListCount.Init(apiRoot),
			api.GetName(getPowerDeviceInfo.EndPoint{}):             getPowerDeviceInfo.Init(apiRoot),
			api.GetName(getPowerDeviceModelList.EndPoint{}):        getPowerDeviceModelList.Init(apiRoot),
			api.GetName(getPowerDeviceModelTechList.EndPoint{}):    getPowerDeviceModelTechList.Init(apiRoot),
			api.GetName(getPowerDeviceTypeList.EndPoint{}):         getPowerDeviceTypeList.Init(apiRoot),
			api.GetName(getPowerPlanList.EndPoint{}):               getPowerPlanList.Init(apiRoot),
			api.GetName(getPowerStation.EndPoint{}):                getPowerStation.Init(apiRoot),
			api.GetName(getPowerStationInfo.EndPoint{}):            getPowerStationInfo.Init(apiRoot),
			api.GetName(getPowerStationList.EndPoint{}):            getPowerStationList.Init(apiRoot),
			api.GetName(getPowerStore.EndPoint{}):                  getPowerStore.Init(apiRoot),
			api.GetName(getProvcnNew.EndPoint{}):                   getProvcnNew.Init(apiRoot),
			api.GetName(getPsTreeMenu.EndPoint{}):                  getPsTreeMenu.Init(apiRoot),
			api.GetName(getRoleByUserIds.EndPoint{}):               getRoleByUserIds.Init(apiRoot),
			api.GetName(getRootOrgInfoByUserId.EndPoint{}):         getRootOrgInfoByUserId.Init(apiRoot),
			api.GetName(getSettingUserMapData.EndPoint{}):          getSettingUserMapData.Init(apiRoot),
			api.GetName(getStateNew.EndPoint{}):                    getStateNew.Init(apiRoot),
			api.GetName(getSubTypeDevice.EndPoint{}):               getSubTypeDevice.Init(apiRoot),
			api.GetName(getSysHomeList2.EndPoint{}):                getSysHomeList2.Init(apiRoot),
			api.GetName(getSysMenu.EndPoint{}):                     getSysMenu.Init(apiRoot),
			api.GetName(getSysOrgPro.EndPoint{}):                   getSysOrgPro.Init(apiRoot),
			api.GetName(getSysUser.EndPoint{}):                     getSysUser.Init(apiRoot),
			api.GetName(getSystemOrgInfo.EndPoint{}):               getSystemOrgInfo.Init(apiRoot),
			api.GetName(getSystemRoleInfo.EndPoint{}):              getSystemRoleInfo.Init(apiRoot),
			api.GetName(getSystemRoleList2.EndPoint{}):             getSystemRoleList2.Init(apiRoot),
			api.GetName(getTownValueNew.EndPoint{}):                getTownValueNew.Init(apiRoot),
			api.GetName(getUserMenuLs.EndPoint{}):                  getUserMenuLs.Init(apiRoot),
			api.GetName(getUserOrgPage.EndPoint{}):                 getUserOrgPage.Init(apiRoot),
			api.GetName(getVillageList.EndPoint{}):                 getVillageList.Init(apiRoot),
			api.GetName(getVillageListNew.EndPoint{}):              getVillageListNew.Init(apiRoot),
			api.GetName(getZtreeAsyncSysMenu.EndPoint{}):           getZtreeAsyncSysMenu.Init(apiRoot),
			api.GetName(getZtreeChildMenu.EndPoint{}):              getZtreeChildMenu.Init(apiRoot),
			api.GetName(getZtreeMenu.EndPoint{}):                   getZtreeMenu.Init(apiRoot),
			api.GetName(getZtreeSysMenu.EndPoint{}):                getZtreeSysMenu.Init(apiRoot),
			api.GetName(getZtreeSysMenu2.EndPoint{}):               getZtreeSysMenu2.Init(apiRoot),
			api.GetName(goToDevicePropertyPage.EndPoint{}):         goToDevicePropertyPage.Init(apiRoot),
			api.GetName(isCanAddUser.EndPoint{}):                   isCanAddUser.Init(apiRoot),
			api.GetName(isHasIrradiationData.EndPoint{}):           isHasIrradiationData.Init(apiRoot),
			api.GetName(isHasPlan.EndPoint{}):                      isHasPlan.Init(apiRoot),
			api.GetName(loadDevice.EndPoint{}):                     loadDevice.Init(apiRoot),
			api.GetName(modelPointsPage.EndPoint{}):                modelPointsPage.Init(apiRoot),
			api.GetName(modifyDevice.EndPoint{}):                   modifyDevice.Init(apiRoot),
			api.GetName(modifyPowerDeviceChannl.EndPoint{}):        modifyPowerDeviceChannl.Init(apiRoot),
			api.GetName(modifySysOrg.EndPoint{}):                   modifySysOrg.Init(apiRoot),
			api.GetName(modifySystemMenu.EndPoint{}):               modifySystemMenu.Init(apiRoot),
			api.GetName(modifySystemOrgNode.EndPoint{}):            modifySystemOrgNode.Init(apiRoot),
			api.GetName(modifySystemRole.EndPoint{}):               modifySystemRole.Init(apiRoot),
			api.GetName(modifySystemUser.EndPoint{}):               modifySystemUser.Init(apiRoot),
			api.GetName(publishNotice.EndPoint{}):                  publishNotice.Init(apiRoot),
			api.GetName(queryDeviceList.EndPoint{}):                queryDeviceList.Init(apiRoot),
			api.GetName(queryDutyType.EndPoint{}):                  queryDutyType.Init(apiRoot),
			api.GetName(queryReportDataById.EndPoint{}):            queryReportDataById.Init(apiRoot),
			api.GetName(resetPasW.EndPoint{}):                      resetPasW.Init(apiRoot),
			api.GetName(saveAuthKey.EndPoint{}):                    saveAuthKey.Init(apiRoot),
			api.GetName(saveDevice.EndPoint{}):                     saveDevice.Init(apiRoot),
			api.GetName(saveDeviceFactory.EndPoint{}):              saveDeviceFactory.Init(apiRoot),
			api.GetName(saveDeviceType.EndPoint{}):                 saveDeviceType.Init(apiRoot),
			api.GetName(saveIrradiationData.EndPoint{}):            saveIrradiationData.Init(apiRoot),
			api.GetName(saveModelPoints.EndPoint{}):                saveModelPoints.Init(apiRoot),
			api.GetName(saveNewNotice.EndPoint{}):                  saveNewNotice.Init(apiRoot),
			api.GetName(saveOrUpdateReport.EndPoint{}):             saveOrUpdateReport.Init(apiRoot),
			api.GetName(saveOrgNode.EndPoint{}):                    saveOrgNode.Init(apiRoot),
			api.GetName(saveOrgUsers.EndPoint{}):                   saveOrgUsers.Init(apiRoot),
			api.GetName(savePicture.EndPoint{}):                    savePicture.Init(apiRoot),
			api.GetName(savePointManage.EndPoint{}):                savePointManage.Init(apiRoot),
			api.GetName(savePowerDeviceChannl.EndPoint{}):          savePowerDeviceChannl.Init(apiRoot),
			api.GetName(savePowerDeviceModel.EndPoint{}):           savePowerDeviceModel.Init(apiRoot),
			api.GetName(savePowerDeviceParameterPage.EndPoint{}):   savePowerDeviceParameterPage.Init(apiRoot),
			api.GetName(savePowerDeviceSubType.EndPoint{}):         savePowerDeviceSubType.Init(apiRoot),
			api.GetName(savePowerDeviceTechnical.EndPoint{}):       savePowerDeviceTechnical.Init(apiRoot),
			api.GetName(savePowerPlan.EndPoint{}):                  savePowerPlan.Init(apiRoot),
			api.GetName(savePowerStationByPowerStore.EndPoint{}):   savePowerStationByPowerStore.Init(apiRoot),
			api.GetName(savePowerStore.EndPoint{}):                 savePowerStore.Init(apiRoot),
			api.GetName(savePsOrg.EndPoint{}):                      savePsOrg.Init(apiRoot),
			api.GetName(saveRelDevice.EndPoint{}):                  saveRelDevice.Init(apiRoot),
			api.GetName(saveRoleAssign.EndPoint{}):                 saveRoleAssign.Init(apiRoot),
			api.GetName(saveSysMenu.EndPoint{}):                    saveSysMenu.Init(apiRoot),
			api.GetName(saveSysOrg.EndPoint{}):                     saveSysOrg.Init(apiRoot),
			api.GetName(saveSysRole.EndPoint{}):                    saveSysRole.Init(apiRoot),
			api.GetName(saveSysUser.EndPoint{}):                    saveSysUser.Init(apiRoot),
			api.GetName(saveUserNode.EndPoint{}):                   saveUserNode.Init(apiRoot),
			api.GetName(saveUserRole.EndPoint{}):                   saveUserRole.Init(apiRoot),
			api.GetName(searchIrradiationData.EndPoint{}):          searchIrradiationData.Init(apiRoot),
			api.GetName(searchTechnicalNums.EndPoint{}):            searchTechnicalNums.Init(apiRoot),
			api.GetName(selectDeviceTypeByPsId.EndPoint{}):         selectDeviceTypeByPsId.Init(apiRoot),
			api.GetName(selectPowerDeviceTechnicals.EndPoint{}):    selectPowerDeviceTechnicals.Init(apiRoot),
			api.GetName(selectPowerDeviceType.EndPoint{}):          selectPowerDeviceType.Init(apiRoot),
			api.GetName(setupUserRole4AddUser.EndPoint{}):          setupUserRole4AddUser.Init(apiRoot),
			api.GetName(startWorkFlow.EndPoint{}):                  startWorkFlow.Init(apiRoot),
			api.GetName(updateDevice.EndPoint{}):                   updateDevice.Init(apiRoot),
			api.GetName(updateDeviceType.EndPoint{}):               updateDeviceType.Init(apiRoot),
			api.GetName(updateFaultLevel.EndPoint{}):               updateFaultLevel.Init(apiRoot),
			api.GetName(updateNotice.EndPoint{}):                   updateNotice.Init(apiRoot),
			api.GetName(updatePointInfo.EndPoint{}):                updatePointInfo.Init(apiRoot),
			api.GetName(updatePowerDeviceModel.EndPoint{}):         updatePowerDeviceModel.Init(apiRoot),
			api.GetName(updatePowerDeviceParameterPage.EndPoint{}): updatePowerDeviceParameterPage.Init(apiRoot),
			api.GetName(updatePowerDeviceSubType.EndPoint{}):       updatePowerDeviceSubType.Init(apiRoot),
			api.GetName(updatePowerDeviceTechnical.EndPoint{}):     updatePowerDeviceTechnical.Init(apiRoot),
			api.GetName(updateProcessManage.EndPoint{}):            updateProcessManage.Init(apiRoot),
			api.GetName(updateSysOrgPro.EndPoint{}):                updateSysOrgPro.Init(apiRoot),
			api.GetName(updateSysRoleValidFlag.EndPoint{}):         updateSysRoleValidFlag.Init(apiRoot),
			api.GetName(updateUserValidFlag.EndPoint{}):            updateUserValidFlag.Init(apiRoot),
			api.GetName(updateValidFlag.EndPoint{}):                updateValidFlag.Init(apiRoot),
			api.GetName(viewDeviceModel.EndPoint{}):                viewDeviceModel.Init(apiRoot),
			api.GetName(viewDeviceParameter.EndPoint{}):            viewDeviceParameter.Init(apiRoot),
			api.GetName(workFlowImplementStep.EndPoint{}):          workFlowImplementStep.Init(apiRoot),
			api.GetName(workFlowIsStart.EndPoint{}):                workFlowIsStart.Init(apiRoot),
			api.GetName(workFlowTransferStep.EndPoint{}):           workFlowTransferStep.Init(apiRoot),
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
