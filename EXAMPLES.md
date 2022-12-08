# Examples of all known working commands

## Authentication

    GoSungrow api login


## Top level

```
Available Commands:
  api          Api	- Low-level interface to the SunGrow api.
  data         Data	- Mid-level access to the Sungrow api.
  info         Info	- General iSolarCloud functions.
  show         Show	- High-level Sungrow commands.
  mqtt         Connect to a HASSIO broker.
  version      Version	- Self-manage this executable.
  selfupdate   Version	- Update version of executable.
  daemon       Daemon	- Daemonize commands.
  cron         Cron	- Run a command via schedule.
  config       Config	- Create, update or show config file.
  shell        Shell	- Run as an interactive shell.
  help         Help	- Extended help
  help         Help about any command
  completion   Generate the autocompletion script for the specified shell
```


### api commands

Low-level interface to the SunGrow api.

```
Available Commands:
  ls           Api	- List SunGrow api endpoints/areas
  login        Api	- Login to the SunGrow api.
  get          Api	- Get endpoint details from the SunGrow api.
  raw          Api	- Raw response from the SunGrow api.
  save         Api	- Save the response from the SunGrow api.
  struct       Api	- Show response as Go structure (debug)
  put          Api	- Put details onto the SunGrow api.
```


### data commands

```
Available Commands:
  list         Data	- Get data from the Sungrow api (list)
  table        Data	- Get data from the Sungrow api (table)
  raw          Data	- Get data from the Sungrow api (raw)
  json         Data	- Get data from the Sungrow api (json)
  csv          Data	- Get data from the Sungrow api (csv)
  graph        Data	- Get data from the Sungrow api (graph)
  xml          Data	- Get data from the Sungrow api (xml)
  xlsx         Data	- Get data from the Sungrow api (XLSX)
  md           Data	- Get data from the Sungrow api (MarkDown)
  struct       Data	- Show response as Go structure (debug)
```


### show commands

```
Available Commands:
  ps           PsId	- Ps related Sungrow commands.
  device       Device	- Device related Sungrow commands.
  template     Template	- Template related Sungrow commands.
  point        Point	- Point related Sungrow commands.
```


#### show ps commands

```
Available Commands:
  list         PsId	- Show all devices on account.
  tree         PsId	- Show the PS tree.
  points       PsId	- List points used for a given ps_id.
  data         PsId	- Generate points table for a given ps_id.
  graph        PsId	- Generate graphs of points for a given ps_id.
```


#### show device commands

```
Available Commands:
  list         Device	- List all device types.
  points       Device	- List points used for a given device_type.
  data         Device	- Generate points table for a given device_type.
  graph        Device	- Generate graphs of points for a given device_type.
  models       Device	- Get ALL Sungrow models (large list).
```


#### show template commands

```
Available Commands:
  list         Template	- Get all defined templates.
  points       Template	- List points used for a given template_id.
  data         Template	- Generate points table for a given template_id.
  graph        Template	- Generate graphs of points for a given template_id.
```


#### show point commands

```
Available Commands:
  ps              Point	- List data points used by a given ps_id.
  ps-data         Point	- Generate points table for a given ps_id.
  ps-graph        Point	- Generate graphs of points for a given ps_id.
  device          Point	- List data points used by a device.
  device-data     Point	- Generate points table for a given device.
  device-graph    Point	- Generate graphs of points for a given device.
  template        Point	- List data points used by a report template.
  template-data   Point	- Generate points table for a given report template.
  template-graph  Point	- Generate graphs of points for a given report template.
  data            Point	- Get data points.
  graph           Point	- Graph data points.
  scan            Point	- Scan full list of points.
```


## High-level data commands

### Templates

```
GoSungrow show template list
┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┓
┃ Template Id ┃ Template Name ┃ Update On           ┃
┣━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━┫
┃ 8042        │ Critical      │ 2022-02-15 13:00:28 ┃
┃ 8041        │ extras        │ 2022-02-15 09:40:04 ┃
┃ 8036        │ C             │ 2022-02-15 09:31:35 ┃
┃ 8039        │ v             │ 2022-02-15 09:31:10 ┃
┃ 8040        │ A             │ 2022-02-15 09:30:56 ┃
┃ 8034        │ Percent       │ 2022-02-15 09:30:41 ┃
┃ 8038        │ MWh           │ 2022-02-15 09:09:22 ┃
┃ 8037        │ MW            │ 2022-02-15 09:03:22 ┃
┃ 8033        │ kW            │ 2022-02-15 09:01:19 ┃
┃ 8035        │ Hours         │ 2022-02-15 08:55:56 ┃
┃ 8031        │ kWh           │ 2022-02-15 07:57:36 ┃
┃ 7981        │ Power         │ 2022-02-09 10:03:40 ┃
┗━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━┛
```


```
GoSungrow show template points 7981
┏━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━┓
┃ PointStruct Id        ┃ Description                           ┃ Unit ┃
┣━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━┫
┃ 1129147_11_0_0.p83032 │ Meter AC Power                        │ kW   ┃
┃ 1129147_11_0_0.p83119 │ Daily Feed-in Energy (PV)             │ kWh  ┃
┃ 1129147_11_0_0.p83549 │ Grid active power                     │ kW   ┃
┃ 1129147_11_0_0.p83002 │ Inverter AC Power                     │ kW   ┃
┃ 1129147_11_0_0.p83011 │ Meter E-daily Consumption             │ kWh  ┃
┃ 1129147_11_0_0.p83022 │ Daily Yield of Plant                  │ kWh  ┃
┃ 1129147_11_0_0.p83006 │ Meter Daily Yield                     │ kWh  ┃
┃ 1129147_11_0_0.p83033 │ Plant Power                           │ kW   ┃
┃ 1129147_11_0_0.p83072 │ Daily Feed-in Energy                  │ kWh  ┃
┃ 1129147_11_0_0.p83097 │ Daily Load Energy Consumption from PV │ kWh  ┃
┃ 1129147_11_0_0.p83102 │ Daily Purchased Energy                │ kWh  ┃
┃ 1129147_14_1_1.p13028 │ Daily Battery Charging Energy         │ kWh  ┃
┃ 1129147_14_1_1.p13112 │ Daily PV Yield                        │ kWh  ┃
┃ 1129147_14_1_1.p13147 │ Daily Purchased Energy                │ kWh  ┃
┃ 1129147_14_1_1.p13173 │ Daily Feed-in Energy (PV)             │ kWh  ┃
┃ 1129147_14_1_1.p13174 │ Daily Battery Charging Energy from PV │ kWh  ┃
┃ 1129147_14_1_1.p13199 │ Daily Load Energy Consumption         │ kWh  ┃
┃ 1129147_14_1_1.p13116 │ Daily Load Energy Consumption from PV │ kWh  ┃
┃ 1129147_14_1_1.p13122 │ Daily Feed-in Energy                  │ kWh  ┃
┗━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━┛
```

```
GoSungrow show template data 8042 20220212
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━┓
┃ Date/Time           ┃ Point Id              ┃ Point Name                 ┃ Value       ┃ Units ┃
┣━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━┫
┃ 2022-02-12 00:00:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 396         │ kW    ┃
┃ 2022-02-12 00:05:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 480         │ kW    ┃
┃ 2022-02-12 00:10:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 487         │ kW    ┃
┃ 2022-02-12 00:15:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 497         │ kW    ┃
┃ 2022-02-12 00:20:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 524         │ kW    ┃
┃ 2022-02-12 00:25:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 541         │ kW    ┃
┃ 2022-02-12 00:30:00 │ 1129147_11_0_0.p83106 │ Load Power                 │ 554         │ kW    ┃

...

┃ 2022-02-12 23:55:00 │ 1129147_14_1_1.p13019 │ Internal Air Temperature   │ 19.4        │ ℃     ┃
┃ 2022-02-13 00:00:00 │ 1129147_14_1_1.p13019 │ Internal Air Temperature   │ 22.3        │ ℃     ┃
┗━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━┛
```

    GoSungrow show template graph 8042 20220212 p13019
    GoSungrow show template graph 8042 20220212 p83106


## Hacking


## Base API commands 

```
Usage:
GoSungrow api [command]

Examples:
GoSungrow api login
GoSungrow api get <endpoint> <args>
GoSungrow api raw <endpoint> <args>
GoSungrow api save <endpoint> <args>
GoSungrow api put <endpoint> <args>

Available Commands:
ls           List iSolarCloud api endpoints/areas
login        Login to iSolarCloud
get          Get details from iSolarCloud
raw          Raw details from iSolarCloud
save         Save details from iSolarCloud as JSON
put          Put details onto iSolarCloud
```

## Get device details

    GoSungrow data get getPsList
    GoSungrow data get getPsListByName
    GoSungrow data get getPsDataSupplementTaskList
    GoSungrow data get getPsUser
    GoSungrow data get queryPsIdList
    GoSungrow data get getInstallInfoList
    GoSungrow data get getPsListStaticData
    GoSungrow data get queryAllPsIdAndName
    GoSungrow data get getPsDetail PsId:1171348
    GoSungrow data get getPsDetailWithPsType PsId:1171348
    GoSungrow data get getPsHealthState PsId:1171348
    GoSungrow data get getPsWeatherList PsId:1171348
    GoSungrow data get queryDeviceList PsId:1171348
    GoSungrow data get queryDeviceListByUserId PsId:1171348
    GoSungrow data get queryDeviceListForApp PsId:1171348
    GoSungrow data get findPsType PsId:1171348
    GoSungrow data get getDeviceList PsId:1171348
    GoSungrow data get getIncomeSettingInfos PsId:1171348
    GoSungrow data get getPowerChargeSettingInfo PsId:1171348
    GoSungrow data get getPowerStationInfo PsId:1171348
    GoSungrow data get WebAppService.getDeviceUuid PsKey:1171348
    GoSungrow data get getPowerStationForHousehold PsId:1171348


## Get device data

    GoSungrow data get energyTrend
    GoSungrow data get getKpiInfo
    GoSungrow data get queryPsProfit DateId:20221001 PsId:1171348
    GoSungrow data get getPowerStationData DateId:20221001 PsId:1171348
    GoSungrow data get getHouseholdStoragePsReport DateId:20221001 PsId:1171348
    GoSungrow data get queryPsProfit DateId:202210 PsId:1171348
    GoSungrow data get getPowerStationData DateId:202210 PsId:1171348
    GoSungrow data get getHouseholdStoragePsReport DateId:202210 PsId:1171348
    GoSungrow data get queryPsProfit DateId:2022 PsId:1171348
    GoSungrow data get getPowerStationData DateId:2022 PsId:1171348
    GoSungrow data get getHouseholdStoragePsReport DateId:2022 PsId:1171348
    GoSungrow data get getPowerTrendDayData BeginTime:20221004
    GoSungrow data get getPowerStatistics PsId:1171348
    GoSungrow data get WebAppService.showPSView PsId:1171348


## Reports

    GoSungrow data get getPsReport DateId:20220201 PsId:1171348
    GoSungrow data get getPsReport DateId:20220201 PsId:1171348
    GoSungrow data get getPsReport DateId:20220201 PsId:1171348
    GoSungrow data get reportList PsId:1171348 ReportType:1
    GoSungrow data get reportList PsId:1171348 ReportType:2
    GoSungrow data get reportList PsId:1171348 ReportType:3
    GoSungrow data get reportList PsId:1171348 ReportType:4


## Templates

    GoSungrow data get template 20220202
    GoSungrow data get template 8042 20220212
    GoSungrow data get getTemplateList
    GoSungrow data get WebAppService.queryUserCurveTemplateData TemplateId:8042 StartTime:20220223000000 EndTime:20220223235900


## User/installer/support info

    GoSungrow data get getUserList
    GoSungrow data get getUserPsOrderList
    GoSungrow data get getPhotoInfo
    GoSungrow data get getOrgListByName
    GoSungrow data get getOrgListByUserId
    GoSungrow data get getOrgListForUser
    GoSungrow data get queryUserList
    GoSungrow data get getInstallerInfoByDealerOrgCodeOrId dealer_org_code:AUSCEKK7
    GoSungrow data get getInstallerInfoByDealerOrgCodeOrId OrgId:362245
    GoSungrow data get getInstallerInfoByDealerOrgCodeOrId OrgId:80384
    GoSungrow data get getInstallerInfoByDealerOrgCodeOrId OrgId:80393
    GoSungrow data get getInstallerInfoByDealerOrgCodeOrId OrgId:300977


## Meta-data

	WebIscmAppService.modelPointsPage DeviceModelId:714 DeviceType:14 - Use this to fetch all points for a particular ps_key.
	AppService.getPowerDevicePointNames DeviceType:14 - Point names for device.
	WebIscmAppService.getModelPoints DeviceModelId:714 - Point names for device.
	WebIscmAppService.queryDeviceListForBackSys - Basic relationships between devices.
	AppService.queryInverterModelList
	WebIscmAppService.getDeviceModel
	WebIscmAppService.getPointInfo
	WebIscmAppService.getPointInfoPage
	WebIscmAppService.getPowerDeviceTypeList
	WebIscmAppService.getPsTreeMenu
	WebIscmAppService.getUserMenuLs UserId:276937
	WebIscmAppService.getSysMenu MenuId:613
	WebIscmAppService.getDeviceTechnical
	WebIscmAppService.getDeviceType
	WebIscmAppService.getDeviceTypeInfoById CodeType:14
	WebIscmAppService.getCodeByType
	WebIscmAppService.getDeviceFactoryListByIds
	AppService.getOwnerFaultConfigList
	AppService.getSungwsGlobalConfigCache
	AppService.getRemoteParamSettingList CurPage:1 Size:100 DeviceType:14
	WebAppService.getSelfReportPoint
	WebAppService.getDevicePointAttrs Uuid:1179878 DeviceType2:14 PsId:1171348

    GoSungrow data get getDeviceTypeInfoList
    GoSungrow data get getDeviceTypeList
    GoSungrow data get getInvertDataList
    GoSungrow data get getModuleLogTaskList
    GoSungrow data get powerDevicePointList
    GoSungrow data get getDeviceModelInfoList
    GoSungrow data get queryUnitList
    GoSungrow data get getPowerSettingCharges
    GoSungrow data get getPowerDeviceModelTechList DeviceType:14
    GoSungrow data get getPowerDevicePointInfo Id:1


## Task commands

    GoSungrow data get queryBatchCreatePsTaskList
    GoSungrow data get getPowerDeviceSetTaskDetailList QueryType:2 TaskId:1 Uuid:844763
    GoSungrow data get getPowerDeviceSetTaskDetailList QueryType:7 TaskId:1 Uuid:844763
    GoSungrow data get getPowerDeviceSetTaskList Size:0 CurPage:1
    GoSungrow data get getRemoteUpgradeSubTasksList QueryType:2 TaskId:1577700
    GoSungrow data get getRemoteUpgradeTaskList PsIdList:1171348


## Misc commands

    GoSungrow data get getDevicePoints PointId:13003
    GoSungrow data get getPowerPictureList


## Hacking

    GoSungrow data get checkUnitStatus
    GoSungrow data get getAllPowerDeviceSetName
    GoSungrow data get getAreaList
    GoSungrow data get getCloudList
    GoSungrow data get getConfigList
    GoSungrow data get getDeviceInfo
    GoSungrow data get getOSSConfig
    GoSungrow data get getOssObjectStream
    GoSungrow data get getEncryptPublicKey
    GoSungrow data get getFaultCount
    GoSungrow data get getFaultDetail FaultCode:34
    GoSungrow data get getFaultMsgByFaultCode FaultCode:703
    GoSungrow data get getFaultMsgListWithYYYYMM
