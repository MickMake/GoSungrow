// Package WebIscmAppService - API endpoints pulled from the sqlite database, (./assets/interface.db), contained within the Android app com.isolarcloud.manager
package WebIscmAppService

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/addPowerDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/addPowerPointManage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/addSubTypeDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/batchAddDevicesPropertis"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/batchDelDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/batchSavePowerDeviceTechnical"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/checkDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/contactMessageOpera"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/delDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteDeviceFactory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteDeviceType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteOneNotice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteOrgNodeInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePicture"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePointInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceChannl"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceParameterPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceSubType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerDeviceTechnical"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deletePowerStore"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteProcessDefinition"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deleteUserNode"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/deployProcess"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/editProcessManageAction"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/findImageInputStreamString"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getAllDevTypeList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getAllNodeByType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getAuthKey"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getAuthKeyList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getCodeByType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getContactMessage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getCountryNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDefinitionIdByKey"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeploymentList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceFactoryListByIds"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDevicePro"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceSubType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceTechnical"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDeviceTypeInfoById"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getDutyUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getFatherPrivileges"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getGroupManSettings"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getGroupManSettingsMembers"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMaterialByListId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMaterialByType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMaterialList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMaxDeviceIdByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getModelPoints"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getMoneyUnitList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getNamecnNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getNationList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOperationRecord"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgAndChildBasicInfoOptions"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgAndStateAndCode"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgForPs"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgListForUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgNodeInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgStationList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgStationListByPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgUserList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgUserMapData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgZtree"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgZtree4User"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgZtreeAsync"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getOrgZtreeForUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPictureList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPointInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPointInfoPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceChannl"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceFactoryListCount"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceModelTechList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerDeviceTypeList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerPlanList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerStation"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerStationInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerStationList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPowerStore"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getProvcnNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getPsTreeMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getRoleByUserIds"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getRootOrgInfoByUserId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSettingUserMapData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getStateNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSubTypeDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSysHomeList2"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSysMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSysOrgPro"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSysUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSystemOrgInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSystemRoleInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getSystemRoleList2"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getTownValueNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getUserMenuLs"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getUserOrgPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getVillageList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getVillageListNew"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getZtreeAsyncSysMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getZtreeChildMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getZtreeMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/getZtreeSysMenu2"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/goToDevicePropertyPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/isCanAddUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/isHasIrradiationData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/isHasPlan"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/loadDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modelPointsPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifyDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifyPowerDeviceChannl"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifySysOrg"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifySystemMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifySystemOrgNode"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifySystemRole"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/modifySystemUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/publishNotice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/queryDeviceListForBackSys"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/queryDutyType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/queryReportDataById"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/resetPasW"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveAuthKey"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveDeviceFactory"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveDeviceType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveIrradiationData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveModelPoints"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveNewNotice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveOrUpdateReport"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveOrgNode"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveOrgUsers"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePicture"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePointManage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerDeviceChannl"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerDeviceParameterPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerDeviceSubType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerDeviceTechnical"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerPlan"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerStationByPowerStore"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePowerStore"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/savePsOrg"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveRelDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveRoleAssign"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveSysMenu"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveSysOrg"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveSysRole"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveSysUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveUserNode"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/saveUserRole"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/searchIrradiationData"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/searchTechnicalNums"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/selectDeviceTypeByPsId"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceTechnicals"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/selectPowerDeviceType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/setupUserRole4AddUser"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/startWorkFlow"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateDevice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateDeviceType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateFaultLevel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateNotice"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updatePointInfo"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceParameterPage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceSubType"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updatePowerDeviceTechnical"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateProcessManage"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateSysOrgPro"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateSysRoleValidFlag"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateUserValidFlag"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/updateValidFlag"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/viewDeviceModel"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/viewDeviceParameter"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/workFlowImplementStep"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/workFlowIsStart"
	"github.com/anicoll/gosungrow/iSolarCloud/WebIscmAppService/workFlowTransferStep"
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
			api.GetName(queryDeviceListForBackSys.EndPoint{}):      queryDeviceListForBackSys.Init(apiRoot),
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
