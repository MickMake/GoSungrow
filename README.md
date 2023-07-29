# GoSungrow - iSolarCloud API written in GoLang.

Note: The next major release, (v3.0.x), is here! Check out the [Latest releases here](https://github.com/MickMake/GoSungrow/releases).

[![Go Reference](https://pkg.go.dev/badge/github.com/MickMake/GoSungrow.svg)](https://pkg.go.dev/github.com/MickMake/GoSungrow)

Find this useful? [You can support development of this app](https://paypal.me/MickMake)

![image](https://github.com/MickMake/GoSungrow/assets/17118367/c015b207-9aed-4aab-b521-57408bba85f5)


## What is it?

This GoLang package has a complete implementation of the iSolarCloud API.
There's been no published specs on this, so I've had to figure it all out based on the [Android app](https://play.google.com/store/apps/details?id=com.isolarcloud.manager), using javascript IDEs and various other means.

Note:
- [iSolarCloud](https://isolarcloud.com) has no interest in developing a public API.
- Their "API" implementation is so broken with security and coding issues, I'm surprised it hasn't been exploited yet.
- iSolarCloud reached out to me, (based off this GitHub page), to see what can be done about these security issues. So that's a very good thing.
- As of 1st August 2022 I discovered a number of security holes. At which point I notified AusCert, the vendor and a number of other security companies - hopefully they will get patched because my tool makes it very easy.

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/iSolarCloudLogin.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/iSolarCloud.png?raw=true)

I'm currently using it in my [HomeAssistant](https://www.home-assistant.io/) instance.

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/SunGrowOnHASSIO1.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/SunGrowOnHASSIO2.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/SunGrowOnHASSIO3.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/SunGrowOnHASSIO4.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/SunGrowOnHASSIO5.png?raw=true)


![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/Grafana1.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/Grafana2.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/Grafana3.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/Grafana4.png?raw=true)


## What state is it in?

It was originally intended for my needs, (seeing all data in [HomeAssistant](https://www.home-assistant.io/)), but there seems to be a big interest in this tool.
So I spent some time working on the v3.0.0 release, adding features.

I have now mapped out all the API calls. All the read-only endpoints are mapped out and fully tested. The write-only calls haven't been tested fully.
It's tricky as their "API" changes regularly; however I've accommodated for quick changes in the v3.0.0 release.

Most endpoints contain repeated data. The main endpoints that house most of the data I've provided easy commands to access.
Of course, all endpoints are accessible. So go for your life.

```
+-------------------+-------------------+--------------------+------------+
|       AREAS       | ENABLED ENDPOINTS | DISABLED ENDPOINTS | COVERAGE % |
+-------------------+-------------------+--------------------+------------+
| AliSmsService     |                 1 |                  0 |  100 %     |
| AppService        |               574 |                  0 |  100 %     |
| MttvScreenService |                30 |                  0 |  100 %     |
| NullArea          |                 1 |                  0 |  100 %     |
| PowerPointService |                 1 |                  0 |  100 %     |
| WebAppService     |               190 |                  0 |  100 %     |
| WebIscmAppService |               184 |                  0 |  100 %     |
| ----------------  | ----------------  | -----------------  | ---------  |
| Total             |               981 |                  0 |  100 %     |
+-------------------+-------------------+--------------------+------------+
```


## What does it do?

This GoLang package does several things:
1. Provides ready access to all API calls via a simple get/put framework.
2. MQTT client to push to [HomeAssistant](https://www.home-assistant.io/).
3. Graphing any data points, over daily, monthly and yearly with 5 minute to 1 hour granularity.
4. Output data to various formats - tables, JSON, CSV, raw, Graphing, XML, XLSX, Markdown, SQL, HTML and GoLang structures.


## What is the roadmap?

I've implemented most of the features I've wanted to, except for... 
1. IFTTT support.

The most recent version has changed the code-base substantially, making it a lot more robust to changes in the API JSON schema.


## Using GoSungrow:

### Config and login.

Add your username and password to the config. (See [the website](https://portalau.isolarcloud.com/))
Once done, it's a case of set and forget. GoSungrow will handle the re-authentication for you.
```
% ./bin/GoSungrow config write --user=USERNAME --password=PASSWORD
Using config file '/Users/mick/.GoSungrow/config.json'
```

Login to SunGrow website.
```
% ./bin/GoSungrow api login
Email:	your@email.address
Create Date:	Tue Nov 16 23:30:12 CST 2021
Login Last Date:	2022-03-10 17:14:49
Login Last IP:
Login State:	1
User Account:	mickmake
User Id:	276937
User Name:	MickMake
Is Online:	0
Token:	424242_42424242424242424242424242424242
Token File:	/Users/mick/.GoSungrow/AppService_login.json
```


### High level reporting examples.
For more examples see the EXAMPLES.md and examples.txt files.
[EXAMPLES.md](https://github.com/MickMake/GoSungrow/blob/master/EXAMPLES.md)
[examples.txt](https://github.com/MickMake/GoSungrow/blob/master/examples.txt)


Show all devices on your iSolarCloud account.
```
% ./bin/GoSungrow show ps list
┏━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┓
┃ Ps Key           ┃ Ps Id   ┃ Device Type ┃ Device Code ┃ Channel Id ┃ Serial #    ┃ Factory Name ┃ Device Model ┃
┣━━━━━━━━━━━━━━━━━━╇━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━┫
┃ 1129147_14_1_1   │ 1129147 │ 14          │ 1           │ 1          │ B2192301528 │ SUNGROW      │ SH10RT       ┃
┃ 1129147_22_247_1 │ 1129147 │ 22          │ 247         │ 1          │ B2192301528 │ SUNGROW      │ WiNet-S      ┃
┃ 1129147_43_2_1   │ 1129147 │ 43          │ 2           │ 1          │ B2192301528 │ SUNGROW      │ SBR096       ┃
┃ 1171348_14_1_2   │ 1171348 │ 14          │ 1           │ 2          │ B2281302388 │ SUNGROW      │ SH10RT-V112  ┃
┃ 1171348_22_247_2 │ 1171348 │ 22          │ 247         │ 2          │ B2281302388 │ SUNGROW      │ WiNet-S      ┃
┃ 1171348_43_2_2   │ 1171348 │ 43          │ 2           │ 2          │ B2281302388 │ SUNGROW      │ SBR096       ┃
┗━━━━━━━━━━━━━━━━━━┷━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━┛
```


Show the device tree on your iSolarCloud account.
```
% ./bin/GoSungrow show ps tree
+	PsId:1129147	PsName:MickMake	PsKey:1129147_11_0_0	DeviceName:MickMake	Uuid:844763
+--	PsId:1129147	PsName:MickMake	PsKey:1129147_14_1_1	DeviceName:SH10RT	Uuid:844775
+----	PsId:1129147	PsName:MickMake	PsKey:1129147_43_2_1	DeviceName:Battery_001_002	Uuid:1155386
+--	PsId:1129147	PsName:MickMake	PsKey:1129147_22_247_1	DeviceName:WiNet-S	Uuid:844774
+	PsId:1171348	PsName:MickMake42	PsKey:1171348_11_0_0	DeviceName:MickMake42	Uuid:1179860
+--	PsId:1171348	PsName:MickMake42	PsKey:1171348_22_247_2	DeviceName:Communication Module 02_247	Uuid:1179877
+--	PsId:1171348	PsName:MickMake42	PsKey:1171348_14_1_2	DeviceName:Energy Storage System 02_01	Uuid:1179878
+----	PsId:1171348	PsName:MickMake42	PsKey:1171348_43_2_2	DeviceName:Battery 02_02	Uuid:1179879
```


List all known data points for all PS on your account.
```
% ./bin/GoSungrow show ps points 
# Available points:
┏━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Id       ┃ Name                                                     ┃ Unit   ┃ Unit Type ┃ Ps Id   ┃ Device Type ┃ Device Name                 ┃
┣━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━╇━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ p83001   │ Inverter AC Power Normalization                          │ kW/kWp │ 43        │ 1129147 │ 11          │ MickMake                    ┃
┃ p83002   │ Inverter AC Power                                        │ kW     │ 3         │ 1129147 │ 11          │ MickMake                    ┃
┃ p83004   │ Inverter Total Yield                                     │ kWh    │ 7         │ 1129147 │ 11          │ MickMake                    ┃
┃ p83005   │ Daily Equivalent Hours of Meter                          │ h      │ 15        │ 1129147 │ 11          │ MickMake                    ┃
┃ p83006   │ Meter Daily Yield                                        │ kWh    │ 7         │ 1129147 │ 11          │ MickMake                    ┃
┃ p83007   │ Meter PR                                                 │ %      │ 10        │ 1129147 │ 11          │ MickMake                    ┃
┃ p83008   │ Daily Equivalent Hours of Inverter                       │ h      │ 15        │ 1129147 │ 11          │ MickMake                    ┃
┃ p83009   │ Daily Yield by Inverter                                  │ kWh    │ 7         │ 1129147 │ 11          │ MickMake                    ┃
┃ p83010   │ Inverter PR                                              │ %      │ 10        │ 1129147 │ 11          │ MickMake                    ┃
┃ p83011   │ Meter E-daily Consumption                                │ kWh    │ 7         │ 1129147 │ 11          │ MickMake                    ┃

...

┃ p58630   │ Min. Cell Voltage of Module 5                            │ mV     │ 31        │ 1171348 │ 43          │ Battery 02_02               ┃
┃ p58631   │ Min. Cell Voltage of Module 6                            │ mV     │ 31        │ 1171348 │ 43          │ Battery 02_02               ┃
┃ p58632   │ Min. Cell Voltage of Module 7                            │ mV     │ 31        │ 1171348 │ 43          │ Battery 02_02               ┃
┃ p58633   │ Min. Cell Voltage of Module 8                            │ mV     │ 31        │ 1171348 │ 43          │ Battery 02_02               ┃
┃ p58635   │ DC Contactor State                                       │        │ 999       │ 1171348 │ 43          │ Battery 02_02               ┃
┃ p58636   │ Fault Module ID                                          │        │ 999       │ 1171348 │ 43          │ Battery 02_02               ┃
┗━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━┷━━━━━━━━━━━┷━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

Produce data table of device_type 22 on ps_id 1171348 between 20221001 and 20221002 at 60 minute increments.
```
% ./bin/GoSungrow show ps data 1171348 22 20221001 20221002 60

# DataTable AppService.queryMutiPointDataList.ResultData.Data - PsKeys:1171348_22_247_2,1171348_22_247_2,1171348_22_247_2,1171348_22_247_2,1171348_22_247_2,1171348_22_247_2,1171348_22_247_2,1171348_22_247_2 Points:p23001,p23014,p23019,p23020,p23021,p23022,p23023,p23024 PsId:1171348 StartTimeStamp:20221001000000 EndTimeStamp:20221002000000 MinuteInterval:60 
┏━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Timestamp              ┃ Ps Key              ┃ 1171348_22_247_2.p23014    ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 2022-10-01 00:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 01:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 02:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 03:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 04:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 05:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 06:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 07:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 08:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 09:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 10:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 11:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 12:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 13:00:00    ┃ 1171348_22_247_2    ┃ --                         ┃
┃ 2022-10-01 14:00:00    ┃ 1171348_22_247_2    ┃ -68                        ┃
┃ 2022-10-01 15:00:00    ┃ 1171348_22_247_2    ┃ -82                        ┃
┃ 2022-10-01 16:00:00    ┃ 1171348_22_247_2    ┃ -81                        ┃
┃ 2022-10-01 17:00:00    ┃ 1171348_22_247_2    ┃ -81                        ┃
┃ 2022-10-01 18:00:00    ┃ 1171348_22_247_2    ┃ -86                        ┃
┃ 2022-10-01 19:00:00    ┃ 1171348_22_247_2    ┃ -82                        ┃
┃ 2022-10-01 20:00:00    ┃ 1171348_22_247_2    ┃ -90                        ┃
┃ 2022-10-01 21:00:00    ┃ 1171348_22_247_2    ┃ -89                        ┃
┃ 2022-10-01 22:00:00    ┃ 1171348_22_247_2    ┃ -88                        ┃
┃ 2022-10-01 23:00:00    ┃ 1171348_22_247_2    ┃ -86                        ┃
┃ 2022-10-02 00:00:00    ┃ 1171348_22_247_2    ┃ -83                        ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```


Do the same, but with a graph!
```
% ./bin/GoSungrow show ps graph 1171348 22 20221001 20221002 60
Found ps_keys: 1129147_14_1_1,1129147_22_247_1,1129147_43_2_1,1171348_14_1_2,1171348_22_247_2,1171348_43_2_2
Finding points to graph...
Table Headers: Timestamp, Ps Key, 1171348_22_247_2.p23014
Table rows: 25
Found 1 points.
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1171348-1171348_22_247_2.p23014.png'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1171348-1171348_22_247_2.p23014.png?raw=true)


Get all defined report templates.
```
% ./bin/GoSungrow show template list

# DataTable AppService.getTemplateList.ResultData.PageList - DataTable AppService.getTemplateList.ResultData.PageList
┏━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Template Id    ┃ Template Name    ┃ Update Time            ┃
┣━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 7981           ┃ Power            ┃ 2022-02-09 10:03:40    ┃
┃ 8031           ┃ kWh              ┃ 2022-02-15 07:57:36    ┃
┃ 8035           ┃ Hours            ┃ 2022-02-15 08:55:56    ┃
┃ 8033           ┃ kW               ┃ 2022-02-15 09:01:19    ┃
┃ 8037           ┃ MW               ┃ 2022-02-15 09:03:22    ┃
┃ 8038           ┃ MWh              ┃ 2022-02-15 09:09:22    ┃
┃ 8034           ┃ Percent          ┃ 2022-02-15 09:30:41    ┃
┃ 8040           ┃ A                ┃ 2022-02-15 09:30:56    ┃
┃ 8039           ┃ v                ┃ 2022-02-15 09:31:10    ┃
┃ 8036           ┃ C                ┃ 2022-02-15 09:31:35    ┃
┃ 8041           ┃ extras           ┃ 2022-02-15 09:40:04    ┃
┃ 8042           ┃ Critical         ┃ 2022-02-15 13:00:28    ┃
┃ 8092           ┃ ALL1             ┃ 2022-03-09 17:18:21    ┃
┃ 8093           ┃ ALL2             ┃ 2022-03-09 17:20:42    ┃
┃ 8094           ┃ ALL3             ┃ 2022-03-09 17:36:20    ┃
┃ 8095           ┃ ALL4             ┃ 2022-03-09 17:37:56    ┃
┃ 8652           ┃ NewAll1          ┃ 2022-10-04 12:27:00    ┃
┃ 8653           ┃ NewAll2          ┃ 2022-10-04 12:28:58    ┃
┃ 8654           ┃ NewAll03         ┃ 2022-10-04 12:31:39    ┃
┃ 8655           ┃ NewAll04         ┃ 2022-10-04 12:34:51    ┃
┃ 8656           ┃ NewAll05         ┃ 2022-10-04 12:37:34    ┃
┃ 8657           ┃ NewAll06         ┃ 2022-10-04 12:38:27    ┃
┗━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━┛
```


Show all data points used in a report template.
```
% ./bin/GoSungrow show template points 8040

# DataTable WebAppService.queryUserCurveTemplateData.8040.ResultData.PointsData.Devices.[1129147_14_1_1].Points - TemplateId:8040 
┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━━┓
┃ Point Id    ┃ Point Name                        ┃ Ps Id      ┃ Ps Key            ┃ Color      ┃ Detail Id    ┃ Ps Name     ┃ Statistics    ┃ Style    ┃ Unit    ┃ Data List    ┃
┣━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━━╇━━━━━━━━━━━━━━┫
┃ p13002      ┃ MPPT1 Current                     ┃ 1129147    ┃ 1129147_14_1_1    ┃ #FFFF00    ┃ 123808       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13008      ┃ Phase A Current                   ┃ 1129147    ┃ 1129147_14_1_1    ┃ #FF0000    ┃ 123814       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13009      ┃ Phase B Current                   ┃ 1129147    ┃ 1129147_14_1_1    ┃ #00FF00    ┃ 123813       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13010      ┃ Phase C Current                   ┃ 1129147    ┃ 1129147_14_1_1    ┃ #0000FF    ┃ 123812       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13106      ┃ MPPT2 Current                     ┃ 1129147    ┃ 1129147_14_1_1    ┃ #70DB93    ┃ 123807       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13139      ┃ Battery Current                   ┃ 1129147    ┃ 1129147_14_1_1    ┃ #CD7F32    ┃ 123806       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13162      ┃ Max. Charging Current (BMS)       ┃ 1129147    ┃ 1129147_14_1_1    ┃ #C0C0C0    ┃ 123805       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p13163      ┃ Max. Discharging Current (BMS)    ┃ 1129147    ┃ 1129147_14_1_1    ┃ #9F9F9F    ┃ 123804       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p18062      ┃ Phase A Backup Current            ┃ 1129147    ┃ 1129147_14_1_1    ┃ #FF00FF    ┃ 123811       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p18063      ┃ Phase B Backup Current            ┃ 1129147    ┃ 1129147_14_1_1    ┃ #00FFFF    ┃ 123810       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┃ p18064      ┃ Phase C Backup Current            ┃ 1129147    ┃ 1129147_14_1_1    ┃ #000000    ┃ 123809       ┃ MickMake    ┃ 5             ┃ 1        ┃ A       ┃              ┃
┗━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━┷━━━━━━━━━━━━━━┷━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━┷━━━━━━━━━━┷━━━━━━━━━┷━━━━━━━━━━━━━━┛
```


Produce daily report for template 8040 for date 2022/02/24 display on STDOUT.
```
% ./bin/GoSungrow show template data 8040 20220204 20220205 120

# DataTable AppService.queryMutiPointDataList.ResultData.Data - MinuteInterval:120 PsKeys:1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1 Points:p13008,p13010,p13162,p18062,p13009,p18064,p13106,p13139,p13163,p18063,p13002 PsId:1129147 StartTimeStamp:20220204000000 EndTimeStamp:20220205000000 
┏━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Timestamp              ┃ Ps Key            ┃ 1129147_14_1_1.p13002    ┃ 1129147_14_1_1.p13008    ┃ 1129147_14_1_1.p13009    ┃ 1129147_14_1_1.p13010    ┃ 1129147_14_1_1.p13106    ┃ 1129147_14_1_1.p13139    ┃ 1129147_14_1_1.p13162    ┃ 1129147_14_1_1.p13163    ┃ 1129147_14_1_1.p18062    ┃ 1129147_14_1_1.p18063    ┃ 1129147_14_1_1.p18064    ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 2022-02-04 00:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 02:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 04:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 06:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 1.38                     ┃ 1.38                     ┃ 1.29                     ┃ 0                        ┃ 4.6                      ┃ 30                       ┃ 28                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 08:00:00    ┃ 1129147_14_1_1    ┃ 2.39                     ┃ 3.31                     ┃ 3.31                     ┃ 3.31                     ┃ 6.1                      ┃ 0                        ┃ 30                       ┃ 29                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 10:00:00    ┃ 1129147_14_1_1    ┃ 6.35                     ┃ 1.75                     ┃ 1.75                     ┃ 1.75                     ┃ 15.4                     ┃ 24.5                     ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 12:00:00    ┃ 1129147_14_1_1    ┃ 6.99                     ┃ 9.84                     ┃ 9.94                     ┃ 9.84                     ┃ 19.5                     ┃ 0                        ┃ 0                        ┃ 30                       ┃ 0                        ┃ 0.2                      ┃ 0                        ┃
┃ 2022-02-04 14:00:00    ┃ 1129147_14_1_1    ┃ 0.64                     ┃ 0.83                     ┃ 0.83                     ┃ 0.83                     ┃ 1.4                      ┃ 0                        ┃ 0                        ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 16:00:00    ┃ 1129147_14_1_1    ┃ 0.37                     ┃ 3.13                     ┃ 3.13                     ┃ 3.13                     ┃ 0.8                      ┃ 10.1                     ┃ 20                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 18:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 1.01                     ┃ 0.92                     ┃ 1.01                     ┃ 0                        ┃ 3.8                      ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 20:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 1.38                     ┃ 1.47                     ┃ 1.47                     ┃ 0                        ┃ 5.5                      ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-04 22:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┃ 2022-02-05 00:00:00    ┃ 1129147_14_1_1    ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 0                        ┃ 30                       ┃ 30                       ┃ 0                        ┃ 0.1                      ┃ 0                        ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```


And now graph it!
```
% ./bin/GoSungrow show template graph 8040 20220204 20220205 120
Finding points to graph...
Table Headers: Timestamp, Ps Key, 1129147_14_1_1.p13002, 1129147_14_1_1.p13008, 1129147_14_1_1.p13009, 1129147_14_1_1.p13010, 1129147_14_1_1.p13106, 1129147_14_1_1.p13139, 1129147_14_1_1.p13162, 1129147_14_1_1.p13163, 1129147_14_1_1.p18062, 1129147_14_1_1.p18063, 1129147_14_1_1.p18064
Table rows: 13
Found 11 points.
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13002.png'
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13008.png'
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13009.png'
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13010.png'
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13106.png'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1171348-1171348_22_247_2.p23014.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13002.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13008.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13139.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13162.png?raw=true)


List all possible devices
```
% ./bin/GoSungrow show device list
# Available points:
┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Device Type ┃ Name                                           ┃
┣━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 1           │ Inverter                                       ┃
┃ 10          │ String                                         ┃
┃ 11          │ Plant                                          ┃
┃ 12          │ Circuit Protection                             ┃
┃ 13          │ Splitting Device                               ┃
┃ 14          │ Energy Storage System                          ┃
┃ 15          │ Sampling Device                                ┃
┃ 16          │ EMU                                            ┃
┃ 17          │ Unit                                           ┃
┃ 18          │ Temperature and Humidity Sensor                ┃
┃ 19          │ Intelligent Power Distribution Cabinet         ┃
┃ 20          │ Display Device                                 ┃
┃ 21          │ AC Power Distributed Cabinet                   ┃
┃ 22          │ Communication Module                           ┃
┃ 23          │ System-BMS                                     ┃
┃ 24          │ RackBMS                                        ┃
┃ 25          │ DC-DC                                          ┃
┃ 26          │ Energy Management System                       ┃
┃ 28          │ Wind Energy Converter                          ┃
┃ 29          │ SVG                                            ┃
┃ 3           │ Grid-connection Point                          ┃
┃ 30          │ PT Cabinet                                     ┃
┃ 31          │ Bus Protection                                 ┃
┃ 32          │ Cleaning Robot                                 ┃
┃ 33          │ Direct Current Cabinet                         ┃
┃ 34          │ Public Measurement and Control                 ┃
┃ 35          │ Anti-islanding Protection Device               ┃
┃ 36          │ Frequency and Voltage Emergency Control Device ┃
┃ 37          │ PCS                                            ┃
┃ 38          │ Cell BMS                                       ┃
┃ 39          │ Power Quality                                  ┃
┃ 4           │ Combiner Box                                   ┃
┃ 40          │ Shuttle                                        ┃
┃ 41          │ Optimizer                                      ┃
┃ 42          │ Tracking axis communication box                ┃
┃ 43          │ Battery                                        ┃
┃ 44          │ Battery Cluster Management Unit                ┃
┃ 45          │ Local Controller                               ┃
┃ 46          │ Networking Devices                             ┃
┃ 47          │ Energy Storage Unit                            ┃
┃ 48          │ DC Container                                   ┃
┃ 5           │ Meteo Station                                  ┃
┃ 50          │ IO Module                                      ┃
┃ 51          │ Charger                                        ┃
┃ 52          │ Battery System Controller                      ┃
┃ 6           │ Transformer                                    ┃
┃ 7           │ Meter                                          ┃
┃ 8           │ UPS                                            ┃
┃ 9           │ Data Logger                                    ┃
┃ 99          │ Others                                         ┃
┗━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

Get mains power frequency variation graph from template id 8041 on date 2022/02/28
```
% ./bin/GoSungrow show point data 20220301 20220302 120 1129147_14_1_1.p13007

# DataTable AppService.queryMutiPointDataList.ResultData.Data - PsId:1129147 StartTimeStamp:20220301000000 EndTimeStamp:20220302000000 MinuteInterval:120 PsKeys:1129147_14_1_1 Points:p13007 
┏━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ Timestamp              ┃ Ps Key            ┃ 1129147_14_1_1.p13007    ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 2022-03-01 00:00:00    ┃ 1129147_14_1_1    ┃ 49.969997                ┃
┃ 2022-03-01 02:00:00    ┃ 1129147_14_1_1    ┃ 49.98                    ┃
┃ 2022-03-01 04:00:00    ┃ 1129147_14_1_1    ┃ 50.01                    ┃
┃ 2022-03-01 06:00:00    ┃ 1129147_14_1_1    ┃ 49.98                    ┃
┃ 2022-03-01 08:00:00    ┃ 1129147_14_1_1    ┃ 49.98                    ┃
┃ 2022-03-01 10:00:00    ┃ 1129147_14_1_1    ┃ 50.01                    ┃
┃ 2022-03-01 12:00:00    ┃ 1129147_14_1_1    ┃ 50                       ┃
┃ 2022-03-01 14:00:00    ┃ 1129147_14_1_1    ┃ 50.02                    ┃
┃ 2022-03-01 16:00:00    ┃ 1129147_14_1_1    ┃ 49.96                    ┃
┃ 2022-03-01 18:00:00    ┃ 1129147_14_1_1    ┃ 50.01                    ┃
┃ 2022-03-01 20:00:00    ┃ 1129147_14_1_1    ┃ 50                       ┃
┃ 2022-03-01 22:00:00    ┃ 1129147_14_1_1    ┃ 49.969997                ┃
┃ 2022-03-02 00:00:00    ┃ 1129147_14_1_1    ┃ 50.01                    ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━┛

% ./bin/GoSungrow show point graph 20220301 20220302 120 1129147_14_1_1.p13007
Finding points to graph...
Table Headers: Timestamp, Ps Key, 1129147_14_1_1.p13007
Table rows: 13
Found 1 points.
Creating graph file 'AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13007.png'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService.queryMutiPointDataList.ResultData.Data-1129147-1129147_14_1_1.p13007.png?raw=true)


### Using the API instead.
Want to get your hands dirty?

Get basic inverter information for inverter id 1129147
```
% ./bin/GoSungrow api get findPsType '{"ps_id":"1129147"}'
```

```
% ./bin/GoSungrow api get getPsDetailWithPsType '{"ps_id":"1129147"}'
```

Get basic power stats for inverter
```
% ./bin/GoSungrow api get getPowerStatistics '{"ps_id":"1129147"}'
```

Get point_id to point names for different device types
```
% ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"1"}'
```

```
% ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"2"}'
```

```
% ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"7"}'
```

Get all inverters
```
% ./bin/GoSungrow api get getPsList
```

```
% ./bin/GoSungrow api get WebAppService.showPSView '{"ps_id":"1129147"}'
```

Produce basic storage report
```
% ./bin/GoSungrow api get queryMutiPointDataList '{"ps_key":"1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_11_0_0","points":"p13150,p13126,p13142,p13143,p13019,p13141,p13121,p13003,p13149,p83106","minute_interval":"5","start_time_stamp":"20220215000000","end_time_stamp":"20220215235900", "ps_id":"1129147"}'
```

Get the household storage report
```
% ./bin/GoSungrow api get getHouseholdStoragePsReport '{"date_id":"2022","date_type":"4","ps_id":"1129147"}'
```


### Config file.
Show current config.

	% GoSungrow config read

Write current config.

	% GoSungrow config write

Change diff command used in compares.

	% GoSungrow --diff-cmd='sdiff' config write

Change iSolarCloud API token.

	% GoSungrow --cf-token='this is a token string' config write


## Flags available for all commands:
```
% ./bin/GoSungrow help flags
+-----------------+------------+-------------------------+--------------------------------+------------------------------------+
|      FLAG       | SHORT FLAG |       ENVIRONMENT       |          DESCRIPTION           |        VALUE (* = DEFAULT)         |
+-----------------+------------+-------------------------+--------------------------------+------------------------------------+
| --config        |            | GOSUNGROW_CONFIG        | GoSungrow: config file.        | /Users/mick/.GoSungrow/config.json |
| --debug         |            | GOSUNGROW_DEBUG         | GoSungrow: Debug mode.         | false *                            |
| --quiet         |            | GOSUNGROW_QUIET         | GoSungrow: Silence all         | false *                            |
|                 |            |                         | messages.                      |                                    |
| --timeout       |            | GOSUNGROW_TIMEOUT       | Web timeout.                   | 0s                                 |
| --user          | -u         | GOSUNGROW_USER          | SunGrow: api username.         | ------------------                 |
| --password      | -p         | GOSUNGROW_PASSWORD      | SunGrow: api password.         | ---------------------------        |
| --appkey        |            | GOSUNGROW_APPKEY        | SunGrow: api application key.  | 93D72E60331ABDCDC7B39ADC2D1F32B3   |
|                 |            |                         |                                | *                                  |
| --host          |            | GOSUNGROW_HOST          | SunGrow: Provider API URL.     | https://augateway.isolarcloud.com  |
|                 |            |                         |                                | *                                  |
| --token-expiry  |            | GOSUNGROW_TOKEN_EXPIRY  | SunGrow: last login.           | 2022-12-08T16:58:19                |
| --save          | -s         | GOSUNGROW_SAVE          | Save output as a file.         | false *                            |
| --mqtt-user     |            | GOSUNGROW_MQTT_USER     | HASSIO: mqtt username.         | ------------------                 |
| --mqtt-password |            | GOSUNGROW_MQTT_PASSWORD | HASSIO: mqtt password.         | --------------                     |
| --mqtt-host     |            | GOSUNGROW_MQTT_HOST     | HASSIO: mqtt host.             | localhost                          |
| --mqtt-port     |            | GOSUNGROW_MQTT_PORT     | HASSIO: mqtt port.             |                               1883 |
+-----------------+------------+-------------------------+--------------------------------+------------------------------------+
```
