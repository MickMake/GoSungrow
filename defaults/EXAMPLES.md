# Examples of all known working commands

## Authentication

    GoSungrow api login


## High-level info commands

```
Usage:
  GoSungrow info [command]

Available Commands:
  get          Info	- Get info from iSolarCloud (table)
  raw          Info	- Get info from iSolarCloud (raw)
  json         Info	- Get info from iSolarCloud (json)
  csv          Info	- Get info from iSolarCloud (json)
  put          Info	- Set info on iSolarCloud
```

```
Usage:
  GoSungrow info get [command]

Available Commands:
  point-names         Info	- Get iSolarCloud point names.
  mqtt                Info	- Get iSolarCloud MQTT service login details.
  search-point-names  Info	- Get iSolarCloud search point names.
  devices             Info	- Get iSolarCloud devices.
  models              Info	- Get ALL iSolarCloud models.
  templates           Info	- Get all defined templates.
  template-points     Info	- List data points used in report template.
  device-points       Info	- List all available device data points.
```


## Get device details

    GoSungrow info get device
    GoSungrow info get search-point-names
    GoSungrow info get models


## Templates

```
GoSungrow info get templates
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


## High-level data commands

```
Usage:
  GoSungrow data [command]

Available Commands:
  get          Data	- Get high-level data from iSolarCloud (table)
  raw          Data	- Get high-level data from iSolarCloud (raw)
  json         Data	- Get high-level data from iSolarCloud (json)
  csv          Data	- Get high-level data from iSolarCloud (json)
  graph        Data	- Get high-level data from iSolarCloud (graph)
```

```
Usage:
  GoSungrow data get [command]

Available Commands:
  stats        Data	- Get current inverter stats, (last 5 minutes).
  template     Data	- Get data from report template.
  points       Data	- Get points data for a specific date.
  real-time    Data	- Get iSolarCloud real-time data.
  psdetails    Data	- Get iSolarCloud ps details.
```


## General

```
GoSungrow data get stats
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┳━━━━━━┓
┃ Date                ┃ Point Id  ┃ Point Name                ┃ Value   ┃ Unit ┃
┣━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━╇━━━━━━┫
┃ 2022-10-08 09:55:00 │ 1129147.0 │ Co2 Reduce                │ 0       │ kg   ┃
┃ 2022-10-08 09:55:00 │ 1129147.0 │ Co2 Reduce Total          │ 5819    │ kg   ┃
┃ 2022-10-08 09:55:00 │ 1129147.0 │ Curr Power                │ 788     │ W    ┃
┃ 2022-10-08 09:55:00 │ 1129147.0 │ Daily Irradiation         │ --      │      ┃
┃ 2022-10-08 09:55:00 │ 1129147.0 │ Equivalent Hour           │ 0.27    │ Hour ┃

...

┃ 2022-10-08 09:55:00 │ 1171348.0 │ Today Income              │ 0.397   │ AUD  ┃
┃ 2022-10-08 09:55:00 │ 1171348.0 │ Total Capacity            │         │      ┃
┃ 2022-10-08 09:55:00 │ 1171348.0 │ Total Energy              │ 176.5   │ kWh  ┃
┃ 2022-10-08 09:55:00 │ 1171348.0 │ Total Income              │ 35.806  │ AUD  ┃
┃ 2022-10-08 09:55:00 │ 1171348.0 │ Use Energy                │ 6.7     │ kWh  ┃
┗━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━┷━━━━━━┛
```



## Templates

```
GoSungrow info get template-points 7981
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
GoSungrow data get template 8042 20220212
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

    GoSungrow data graph template 8042 20220212 '{"search_string":"p13019","min_left_axis":-6000,"max_left_axis":12000}'
    GoSungrow data graph template 8042 20220212 '{"search_string":"p13019"}'
    GoSungrow data graph template 8042 20220212 '{"search_string":"p83106","min_left_axis":-6000,"max_left_axis":12000}'
    GoSungrow data graph template 8042 20220212 '{"search_string":"p83106"}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":"1","value_column":"4","search_column":"3","search_string":"p83106","file_name":"foo.png"}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p13019","min_left_axis":0,"max_left_axis":0}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p13149","min_left_axis":0,"max_left_axis":0}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p83106","file_name":"foo.png"}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p83106","min_left_axis":-6000,"max_left_axis":1000}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p83106","min_left_axis":-6000,"max_left_axis":12000}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p83106","min_left_axis":-6000,"max_left_axis":42000}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":2,"search_string":"p83106","min_left_axis":0,"max_left_axis":0}'
    GoSungrow data graph template 8042 20220212 '{"title":"Testing 1. 2. 3.","time_column":1,"value_column":4,"search_column":3,"search_string":"p83106","file_name":"foo.png"}'
    GoSungrow data raw template '' 20220201
    GoSungrow data save template '' 20220201
    GoSungrow data save template 8042 20220212


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

    GoSungrow api get getPsList
    GoSungrow api get getPsListByName
    GoSungrow api get getPsDataSupplementTaskList
    GoSungrow api get getPsUser
    GoSungrow api get queryPsIdList
    GoSungrow api get getInstallInfoList
    GoSungrow api get getPsListStaticData
    GoSungrow api get queryAllPsIdAndName
    GoSungrow api get getPsDetail '{"ps_id":"1171348"}'
    GoSungrow api get getPsDetailWithPsType '{"ps_id":"1171348"}'
    GoSungrow api get getPsHealthState '{"ps_id":"1171348"}'
    GoSungrow api get getPsWeatherList '{"ps_id":"1171348"}'
    GoSungrow api get queryDeviceList '{"ps_id":"1171348"}'
    GoSungrow api get queryDeviceListByUserId '{"ps_id":"1171348"}'
    GoSungrow api get queryDeviceListForApp '{"ps_id":"1171348"}'
    GoSungrow api get findPsType '{"ps_id":"1171348"}'
    GoSungrow api get getDeviceList '{"ps_id":"1171348"}'
    GoSungrow api get getIncomeSettingInfos '{"ps_id":"1171348"}'
    GoSungrow api get getPowerChargeSettingInfo '{"ps_id":"1171348"}'
    GoSungrow api get getPowerStationInfo '{"ps_id":"1171348"}'
    GoSungrow api get WebAppService.getDeviceUuid '{"ps_key":"1171348"}'
    GoSungrow api get getPowerStationForHousehold '{"ps_id":"1171348"}'


## Get device data

    GoSungrow api get energyTrend
    GoSungrow api get getKpiInfo
    GoSungrow api get queryPsProfit '{"date_id":"20221001","date_type":"1","ps_id":"1171348"}'
    GoSungrow api get getPowerStationData '{"date_id":"20221001","date_type":"1","ps_id":"1171348"}'
    GoSungrow api get getHouseholdStoragePsReport '{"date_id":"20221001","date_type":"1","ps_id":"1171348"}'
    GoSungrow api get queryPsProfit '{"date_id":"202210","date_type":"2","ps_id":"1171348"}'
    GoSungrow api get getPowerStationData '{"date_id":"202210","date_type":"2","ps_id":"1171348"}'
    GoSungrow api get getHouseholdStoragePsReport '{"date_id":"202210","date_type":"2","ps_id":"1171348"}'
    GoSungrow api get queryPsProfit '{"date_id":"2022","date_type":"3","ps_id":"1171348"}'
    GoSungrow api get getPowerStationData '{"date_id":"2022","date_type":"3","ps_id":"1171348"}'
    GoSungrow api get getHouseholdStoragePsReport '{"date_id":"2022","date_type":"3","ps_id":"1171348"}'
    GoSungrow api get getPowerTrendDayData '{"BeginTime":"20221004"}'
    GoSungrow api get getPowerStatistics '{"ps_id":"1171348"}'
    GoSungrow api get WebAppService.showPSView '{"ps_id":"1171348"}'


## Reports

    GoSungrow api get getPsReport '{"report_type":"1","date_id":"20220201","date_type":"1","ps_id":"1171348"}'
    GoSungrow api get getPsReport '{"report_type":"1","date_id":"20220201","date_type":"2","ps_id":"1171348"}'
    GoSungrow api get getPsReport '{"report_type":"1","date_id":"20220201","date_type":"3","ps_id":"1171348"}'
    GoSungrow api get reportList '{"ps_id":"1171348","report_type":"1"}'
    GoSungrow api get reportList '{"ps_id":"1171348","report_type":"2"}'
    GoSungrow api get reportList '{"ps_id":"1171348","report_type":"3"}'
    GoSungrow api get reportList '{"ps_id":"1171348","report_type":"4"}'


## Templates

    GoSungrow api get template 20220202
    GoSungrow api get template 8042 20220212
    GoSungrow api get getTemplateList
    GoSungrow api get WebAppService.queryUserCurveTemplateData '{"template_id":"8042","date_type":"1","start_time":"20220223000000","end_time":"20220223235900"}'


## User/installer/support info

    GoSungrow api get getUserList
    GoSungrow api get getUserPsOrderList
    GoSungrow api get getPhotoInfo
    GoSungrow api get getOrgListByName
    GoSungrow api get getOrgListByUserId
    GoSungrow api get getOrgListForUser
    GoSungrow api get queryUserList
    GoSungrow api get getInstallerInfoByDealerOrgCodeOrId '{"dealer_org_code":"AUSCEKK7"}'
    GoSungrow api get getInstallerInfoByDealerOrgCodeOrId '{"org_id":"362245"}'
    GoSungrow api get getInstallerInfoByDealerOrgCodeOrId '{"org_id":"80384"}'
    GoSungrow api get getInstallerInfoByDealerOrgCodeOrId '{"org_id":"80393"}'
    GoSungrow api get getInstallerInfoByDealerOrgCodeOrId '{"org_id":"300977"}'


## Meta-data

    GoSungrow api get getDeviceTypeInfoList
    GoSungrow api get getDeviceTypeList
    GoSungrow api get getInvertDataList
    GoSungrow api get getModuleLogTaskList
    GoSungrow api get powerDevicePointList
    GoSungrow api get getPowerDevicePointNames '{"device_type":"1"}'
    GoSungrow api get getPowerDevicePointNames '{"device_type":"2"}'
    GoSungrow api get getPowerDevicePointNames '{"device_type":"7"}'
    GoSungrow api get getDeviceModelInfoList
    GoSungrow api get queryUnitList
    GoSungrow api get getPowerSettingCharges
    GoSungrow api get getPowerDevicePointNames '{"device_type":"1"}'
    GoSungrow api get getPowerDeviceModelTechList '{"device_type":"1"}'
    GoSungrow api get getPowerDevicePointInfo '{"id":"1"}'


## Task commands

    GoSungrow api get queryBatchCreatePsTaskList
    GoSungrow api get getPowerDeviceSetTaskDetailList '{"query_type":"2","task_id":"1","uuid":"844763"}'
    GoSungrow api get getPowerDeviceSetTaskDetailList '{"query_type":"7","task_id":"1","uuid":"844763"}'
    GoSungrow api get getPowerDeviceSetTaskDetailList '{"size":0,"curPage":0}'
    GoSungrow api get getPowerDeviceSetTaskList '{"size":0,"curPage":0}'
    GoSungrow api get getPowerDeviceSetTaskList '{"size":0,"curPage":1}'
    GoSungrow api get getRemoteUpgradeSubTasksList '{"query_type":"1","task_id":"1577700"}'
    GoSungrow api get getRemoteUpgradeTaskList '{"ps_id_list":"1171348"}'


## Misc commands

    GoSungrow api get getDevicePoints '{"point_id":"13003"}'
    GoSungrow api get getPowerPictureList


## Hacking

    GoSungrow api get checkUnitStatus
    GoSungrow api get getAllPowerDeviceSetName
    GoSungrow api get getAreaList
    GoSungrow api get getCloudList
    GoSungrow api get getConfigList
    GoSungrow api get getDeviceInfo
    GoSungrow api get getOSSConfig
    GoSungrow api get getOssObjectStream
    GoSungrow api get getEncryptPublicKey
    GoSungrow api get getFaultCount
    GoSungrow api get getFaultDetail '{"fault_code":"34"}'
    GoSungrow api get getFaultMsgByFaultCode '{"fault_code":"703"}'
    GoSungrow api get getFaultMsgListWithYYYYMM
