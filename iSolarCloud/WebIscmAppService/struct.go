// Package WebIscmAppService - API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebIscmAppService

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/addPowerDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/addPowerPointManage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/addSubTypeDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/batchAddDevicesPropertis"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/batchDelDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/batchSavePowerDeviceTechnical"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/checkDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/contactMessageOpera"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/delDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteDeviceFactory"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteDeviceType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteOneNotice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteOrgNodeInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePicture"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePointInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceChannl"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceParameterPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceSubType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceTechnical"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deletePowerStore"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteProcessDefinition"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteReport"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deleteUserNode"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/deployProcess"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/editProcessManageAction"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/findImageInputStreamString"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getAllDevTypeList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getAllNodeByType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getAuthKey"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getAuthKeyList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getCodeByType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getContactMessage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getCountryNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDefinitionIdByKey"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeploymentList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceFactoryListByIds"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDevicePro"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceSubType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceTechnical"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDeviceTypeInfoById"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getDutyUserList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getFatherPrivileges"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getGroupManSettings"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getGroupManSettingsMembers"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getMaterialByListId"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getMaterialByType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getMaterialList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getMaxDeviceIdByPsId"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getModelPoints"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getMoneyUnitList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getNamecnNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getNationList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOperationRecord"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgAndChildBasicInfoOptions"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgAndStateAndCode"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgForPs"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgListForUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgNodeInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgStationList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgStationListByPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgUserList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgUserMapData"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtree"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtree4User"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtreeAsync"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getOrgZtreeForUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPictureList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPointInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPointInfoPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceChannl"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactory"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactoryListCount"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelTechList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerDeviceTypeList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerPlanList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerStation"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerStationInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerStationList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPowerStore"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getProvcnNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getPsTreeMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getRoleByUserIds"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getRootOrgInfoByUserId"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSettingUserMapData"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getStateNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSubTypeDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSysHomeList2"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSysMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSysOrgPro"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSysUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSystemOrgInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSystemRoleInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getSystemRoleList2"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getTownValueNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getUserMenuLs"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getUserOrgPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getVillageList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getVillageListNew"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getZtreeAsyncSysMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getZtreeChildMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getZtreeMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu2"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/goToDevicePropertyPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/isCanAddUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/isHasIrradiationData"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/isHasPlan"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/loadDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modelPointsPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifyDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifyPowerDeviceChannl"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifySysOrg"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifySystemMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifySystemOrgNode"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifySystemRole"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/modifySystemUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/publishNotice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/queryDeviceListForBackSys"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/queryDutyType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/queryReportDataById"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/resetPasW"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveAuthKey"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveDeviceFactory"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveDeviceType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveIrradiationData"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveModelPoints"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveNewNotice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveOrUpdateReport"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveOrgNode"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveOrgUsers"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePicture"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePointManage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceChannl"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceParameterPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceSubType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerDeviceTechnical"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerPlan"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerStationByPowerStore"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePowerStore"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/savePsOrg"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveRelDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveRoleAssign"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveSysMenu"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveSysOrg"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveSysRole"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveSysUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveUserNode"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/saveUserRole"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/searchIrradiationData"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/searchTechnicalNums"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/selectDeviceTypeByPsId"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceTechnicals"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/setupUserRole4AddUser"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/startWorkFlow"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateDevice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateDeviceType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateFaultLevel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateNotice"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updatePointInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceParameterPage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceSubType"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceTechnical"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateProcessManage"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateSysOrgPro"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateSysRoleValidFlag"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateUserValidFlag"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/updateValidFlag"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/viewDeviceModel"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/viewDeviceParameter"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/workFlowImplementStep"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/workFlowIsStart"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebIscmAppService/workFlowTransferStep"
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
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
			api.GetName(loadDevice.EndPoint{}):                loadDevice.Init(apiRoot),
			api.GetName(modelPointsPage.EndPoint{}):           modelPointsPage.Init(apiRoot),
			api.GetName(modifyDevice.EndPoint{}):              modifyDevice.Init(apiRoot),
			api.GetName(modifyPowerDeviceChannl.EndPoint{}):   modifyPowerDeviceChannl.Init(apiRoot),
			api.GetName(modifySysOrg.EndPoint{}):              modifySysOrg.Init(apiRoot),
			api.GetName(modifySystemMenu.EndPoint{}):          modifySystemMenu.Init(apiRoot),
			api.GetName(modifySystemOrgNode.EndPoint{}):       modifySystemOrgNode.Init(apiRoot),
			api.GetName(modifySystemRole.EndPoint{}):          modifySystemRole.Init(apiRoot),
			api.GetName(modifySystemUser.EndPoint{}):          modifySystemUser.Init(apiRoot),
			api.GetName(publishNotice.EndPoint{}):             publishNotice.Init(apiRoot),
			api.GetName(queryDeviceListForBackSys.EndPoint{}): queryDeviceListForBackSys.Init(apiRoot),
			api.GetName(queryDutyType.EndPoint{}):             queryDutyType.Init(apiRoot),
			api.GetName(queryReportDataById.EndPoint{}):       queryReportDataById.Init(apiRoot),
			api.GetName(resetPasW.EndPoint{}):                 resetPasW.Init(apiRoot),
			api.GetName(saveAuthKey.EndPoint{}):               saveAuthKey.Init(apiRoot),
			api.GetName(saveDevice.EndPoint{}):                saveDevice.Init(apiRoot),
			api.GetName(saveDeviceFactory.EndPoint{}):         saveDeviceFactory.Init(apiRoot),
			api.GetName(saveDeviceType.EndPoint{}):            saveDeviceType.Init(apiRoot),
			api.GetName(saveIrradiationData.EndPoint{}):       saveIrradiationData.Init(apiRoot),
			api.GetName(saveModelPoints.EndPoint{}):           saveModelPoints.Init(apiRoot),
			api.GetName(saveNewNotice.EndPoint{}):             saveNewNotice.Init(apiRoot),
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
