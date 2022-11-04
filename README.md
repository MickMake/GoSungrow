# GoSungrow - iSolarCloud API written in GoLang.

Note: I'm currently going through a big refactor. The next major release, (v3.0.0), will:
- support better api changes. This allows a much quicker change in my code when the api changes.
- support multiple devices. Previous versions only allowed query of one device.
- improve cli for queries. Now can fetch endpoints in either "list" or "table" formats, the latter showing point data in a similar fashion to pivot tables in spreadsheets.
- multiple output formats supported - not only csv, png and ascii, but markdown, xml, HTML, sql and plain MQTT.
- lots of bug fixes. :-).

[Latest releases here](https://github.com/MickMake/GoSungrow/releases)


## What is it?

This GoLang package has a complete, (almost complete), implementation of the iSolarCloud API.
There's been no published specs on this, so I've had to figure it all out based on the [Android app](https://play.google.com/store/apps/details?id=com.isolarcloud.manager), using javascript IDEs and various other means.

Note:
- [iSolarCloud](https://isolarcloud.com) has no interest in developing a public API.
- Their "API" implementation is so broken with security and coding issues, I'm surprised it hasn't been exploited yet.
- iSolarCloud reached out to me, (based off this GitHub page), to see what can be done about these security issues. So that's a very good thing.

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

This is currently usable for my needs, (seeing all data in [HomeAssistant](https://www.home-assistant.io/)), but there's quite a few API endpoints that need to be implemented for completeness.

So far I have mapped out all the API calls. It's tricky as their "API" changes regularly.

I've currently mapped out these API EndPoints. All of the critical read only endpoints are mapped out, with none of the write endpoints.
Most endpoints are a repeat and a lot of the data contained is repeated.
So, in reality 50% of all endpoints are implemented.
```
+-------------------+-------------------+--------------------+------------+
|       AREAS       | ENABLED ENDPOINTS | DISABLED ENDPOINTS | COVERAGE % |
+-------------------+-------------------+--------------------+------------+
| AliSmsService     |                 0 |                  1 | 0.0 %      |
| AppService        |                52 |                522 | 10.0 %     |
| MttvScreenService |                 0 |                 30 | 0.0 %      |
| PowerPointService |                 0 |                  1 | 0.0 %      |
| WebAppService     |                 4 |                186 | 2.2 %      |
| WebIscmAppService |                 0 |                184 | 0.0 %      |
| ----------------  | ----------------  | -----------------  | ---------  |
| Total             |                56 |                924 | 6.1 %      |
+-------------------+-------------------+--------------------+------------+
```


## What does it do?

This GoLang package does several things:
1. Provides ready access to all API calls via a simple get/put framework.
2. MQTT client to push to [HomeAssistant](https://www.home-assistant.io/).
3. Graphing any points data, (over daily, monthly and yearly).
4. Update a GitHub repo with SunGrow PV data, (provide full revision history for any changes made to the SunGrow PV).
5. Update a Google sheet with SunGrow PV data.


## What is the roadmap?

I've implemented most of the features I've wanted to, except for... 
1. IFTTT support.


## Using GoSungrow:

### Config and login.

Add your username and password to the config. (See [the website](https://portalau.isolarcloud.com/))
```
$ ./bin/GoSungrow config write --user=USERNAME --password=PASSWORD
Using config file '/Users/mick/.GoSungrow/config.json'
```

Login to SunGrow website.
```
$ ./bin/GoSungrow api login
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

Produce point-in-time statistics
```
$ ./bin/GoSungrow data get stats
+---------------------+-----------+---------------------------+---------+------+
| Date                | Point Id  | Point Name                | Value   | Unit |
+---------------------+-----------+---------------------------+---------+------+
| 2022-02-28 11:05:00 | 1129147.0 | Co2 Reduce                | 3.589   | kg   |
| 2022-02-28 11:05:00 | 1129147.0 | Co2 Reduce Total          | 1047    | kg   |
| 2022-02-28 11:05:00 | 1129147.0 | Curr Power                | 641     | W    |
| 2022-02-28 11:05:00 | 1129147.0 | Daily Irradiation         | --      |      |
| 2022-02-28 11:05:00 | 1129147.0 | Equivalent Hour           | 0.36    | Hour |
| 2022-02-28 11:05:00 | 1129147.0 | Es Discharge Energy       | 4.8     | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Es Energy                 | 1.6     | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Es Power                  | 641     | W    |
| 2022-02-28 11:05:00 | 1129147.0 | Es Total Discharge Energy | 109.9   | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Es Total Energy           | 106.2   | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Installed Power Map       | 10      | kWp  |
| 2022-02-28 11:05:00 | 1129147.0 | Pv Energy                 | 3.6     | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Pv Power                  | 2.291   | kW   |
| 2022-02-28 11:05:00 | 1129147.0 | Radiation                 | --      |      |
| 2022-02-28 11:05:00 | 1129147.0 | Today Energy              | 3.6     | kWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Today Income              | 1.346   | AUD  |
| 2022-02-28 11:05:00 | 1129147.0 | Total Capacity            | 10      | kWp  |
| 2022-02-28 11:05:00 | 1129147.0 | Total Energy              | 1.05    | MWh  |
| 2022-02-28 11:05:00 | 1129147.0 | Total Income              | 302.074 | AUD  |
| 2022-02-28 11:05:00 | 1129147.0 | Use Energy                | 7.6     | kWh  |
+---------------------+-----------+---------------------------+---------+------+
```

Get all defined report templates.
```
./bin/GoSungrow data get templates
+-------------+---------------+---------------------+
| Template Id | Template Name | Update On           |
+-------------+---------------+---------------------+
| 8042        | Critical      | 2022-02-15 13:00:28 |
| 8041        | extras        | 2022-02-15 09:40:04 |
| 8036        | C             | 2022-02-15 09:31:35 |
| 8039        | v             | 2022-02-15 09:31:10 |
| 8040        | A             | 2022-02-15 09:30:56 |
| 8034        | Percent       | 2022-02-15 09:30:41 |
| 8038        | MWh           | 2022-02-15 09:09:22 |
| 8037        | MW            | 2022-02-15 09:03:22 |
| 8033        | kW            | 2022-02-15 09:01:19 |
| 8035        | Hours         | 2022-02-15 08:55:56 |
| 8031        | kWh           | 2022-02-15 07:57:36 |
| 7981        | Power         | 2022-02-09 10:03:40 |
+-------------+---------------+---------------------+
```
Show all data points used in a report template.
```
$ ./bin/GoSungrow data get template-points 8040
+---------+----------------------------+------+
| PointId | Description                | Unit |
+---------+----------------------------+------+
| p83106  | Load Power                 | kW   |
| p13003  | Total DC Power             | kW   |
| p13019  | Internal Air Temperature   | ℃    |
| p13142  | Battery Health (SOH)       | %    |
| p13149  | Purchased Power            | kW   |
| p13150  | Battery Discharging Power  | kW   |
| p13121  | Total Export Active  Power | kW   |
| p13126  | Battery Charging Power     | kW   |
| p13141  | Battery Level (SOC)        | %    |
| p13143  | Battery Temperature        | ℃    |
+---------+----------------------------+------+
```

Produce daily report for template 8040 for date 2022/02/24 display on STDOUT.
```
$ ./bin/GoSungrow data get template 8040 20220224
+---------------------+-----------------------+----------------------------+-------------+-------+
| Date/Time           | Point Id              | Point Name                 | Value       | Units |
+---------------------+-----------------------+----------------------------+-------------+-------+
| 2022-02-24 00:00:00 | 1129147_11_0_0.p83106 | Load Power                 | 818.0       | kW    |
| 2022-02-24 00:05:00 | 1129147_11_0_0.p83106 | Load Power                 | 756.0       | kW    |
| 2022-02-24 00:10:00 | 1129147_11_0_0.p83106 | Load Power                 | 571.0       | kW    |
| 2022-02-24 00:15:00 | 1129147_11_0_0.p83106 | Load Power                 | 557.0       | kW    |
| 2022-02-24 00:20:00 | 1129147_11_0_0.p83106 | Load Power                 | 553.0       | kW    |
| 2022-02-24 00:25:00 | 1129147_11_0_0.p83106 | Load Power                 | 558.0       | kW    |
| 2022-02-24 00:30:00 | 1129147_11_0_0.p83106 | Load Power                 | 562.0       | kW    |

...

+---------------------+-----------------------+----------------------------+-------------+-------+
```

Produce daily report for template 8042 for date 2022/02/24 and save as a JSON file.
```
$ ./bin/GoSungrow data save template 8042 20220224
```

Produce graph of daily report for template 8040 for date 2022/02/24 and point_id P13019.
```
$ ./bin/GoSungrow data get template-points 8040
+---------+----------------------------+------+
| PointId | Description                | Unit |
+---------+----------------------------+------+
| p83106  | Load Power                 | kW   |
| p13003  | Total DC Power             | kW   |
| p13121  | Total Export Active  Power | kW   |
| p13142  | Battery Health (SOH)       | %    |
| p13150  | Battery Discharging Power  | kW   |
| p13019  | Internal Air Temperature   | ℃    |
| p13126  | Battery Charging Power     | kW   |
| p13141  | Battery Level (SOC)        | %    |
| p13143  | Battery Temperature        | ℃    |
| p13149  | Purchased Power            | kW   |
+---------+----------------------------+------+

$ ./bin/GoSungrow data graph template 8042 20220224 '{"search_string":"p13019"}'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220224-8042-InternalAirTemperature.png?raw=true)

Produce graphs of all points from template 8042 for date 2022/02/28 with a different width and height.
```
$ ./bin/GoSungrow data graph template 8042 20220228 '{"width":2400,"height":500}'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-PurchasedPower.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-LoadPower.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-BatteryTemperature.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-InternalAirTemperature.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-BatteryChargingPower.png?raw=true)
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8042-BatteryDischargingPower.png?raw=true)

Produce daily report for point_id p83106 for date 2022/02/24.
```
$ ./bin/GoSungrow data get points 20220224 1129147_11_0_0.p83106
+---------------------+-----------------------+------------+--------+-------+
| Date/Time           | Point Id              | Point Name | Value  | Units |
+---------------------+-----------------------+------------+--------+-------+
| 2022-02-24 00:00:00 | 1129147_11_0_0.p83106 |            | 818.0  |       |
| 2022-02-24 00:05:00 | 1129147_11_0_0.p83106 |            | 756.0  |       |
| 2022-02-24 00:10:00 | 1129147_11_0_0.p83106 |            | 571.0  |       |
| 2022-02-24 00:15:00 | 1129147_11_0_0.p83106 |            | 557.0  |       |
| 2022-02-24 00:20:00 | 1129147_11_0_0.p83106 |            | 553.0  |       |
| 2022-02-24 00:25:00 | 1129147_11_0_0.p83106 |            | 558.0  |       |
| 2022-02-24 00:30:00 | 1129147_11_0_0.p83106 |            | 562.0  |       |
| 2022-02-24 00:35:00 | 1129147_11_0_0.p83106 |            | 561.0  |       |

...

+---------------------+-----------------------+------------+--------+-------+
```

Get all point ids by device_type.
```
+-------------+------------+----------+----------------------------------------+
| Device Type | Point Type | Point Id | Point Name                             |
+-------------+------------+----------+----------------------------------------+
| 1           | 2          | 1        | Daily Yield                            |
| 1           | 1          | 24       | Total Active Power                     |
| 1           | 1          | 2        | Total Yield                            |
| 1           | 1          | 14       | Total DC Power                         |
| 1           | 1          | 18       | Phase A Voltage                        |
| 1           | 1          | 19       | Phase B Voltage                        |
| 1           | 1          | 20       | Phase C Voltage                        |
| 1           | 1          | 21       | Phase A Current                        |
| 1           | 1          | 22       | Phase B Current                        |
| 1           | 1          | 23       | Phase C Current                        |
| 1           | 1          | 25       | Total Reactive Power                   |
| 1           | 1          | 27       | Grid Frequency                         |
| 1           | 1          | 15       | A-B Line Voltage                       |
| 1           | 1          | 16       | B-C Line Voltage                       |
| 1           | 1          | 17       | C-A Line Voltage                       |
| 1           | 1          | 4        | Internal Air Temperature               |
| 1           | 1          | 26       | Total Power Factor                     |
| 1           | 1          | 90       | Negative Voltage to Ground             |
| 1           | 1          | 43       | Total Apparent Power                   |
| 1           | 1          | 33       | Module 1 Temperature                   |
| 1           | 1          | 34       | Module 2 Temperature                   |
| 1           | 1          | 39       | Positive Impedance to Ground           |
| 1           | 1          | 40       | Negative Impedance to Ground           |
| 1           | 1          | 41       | P-limit-actual                         |
| 1           | 1          | 42       | Reactive Power Regulation Actual Value |
| 1           | 2          | 67       | Daily Yield (Theoretical)              |
| 1           | 1          | 5        | MPPT1 Voltage                          |
| 1           | 1          | 6        | MPPT1 Current                          |
| 1           | 1          | 7        | MPPT2 Voltage                          |
| 1           | 1          | 8        | MPPT2 Current                          |
| 1           | 1          | 70       | String 1 Current                       |
| 1           | 1          | 71       | String 2 Current                       |
| 1           | 1          | 72       | String 3 Current                       |
| 1           | 1          | 73       | String 4 Current                       |
| 1           | 1          | 74       | String 5 Current                       |
| 1           | 1          | 75       | String 6 Current                       |
| 1           | 1          | 76       | String 7 Current                       |

...

+-------------+------------+----------+----------------------------------------+
```

Get mains power frequency variation graph from template id 8041 on date 2022/02/28
```
./bin/GoSungrow data graph template 8041 20220228 '{"search_string":"p13007","min_left_axis":49,"max_left_axis":51}'
```
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220228-8041-GridFrequency.png?raw=true)


### Using the API instead.

Get basic inverter information for inverter id 1129147
```
$ ./bin/GoSungrow api get findPsType '{"ps_id":"1129147"}'
```

```
$ ./bin/GoSungrow api get getPsDetailWithPsType '{"ps_id":"1129147"}'
```

Get basic power stats for inverter
```
$ ./bin/GoSungrow api get getPowerStatistics '{"ps_id":"1129147"}'
```

Get point_id to point names for different device types
```
$ ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"1"}'
```

```
$ ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"2"}'
```

```
$ ./bin/GoSungrow api get getPowerDevicePointNames '{"device_type":"7"}'
```

Get all inverters
```
$ ./bin/GoSungrow api get getPsList
```

```
$ ./bin/GoSungrow api get WebAppService.showPSView '{"ps_id":"1129147"}'
```

Produce basic storage report
```
$ ./bin/GoSungrow api get queryMutiPointDataList '{"ps_key":"1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_14_1_1,1129147_11_0_0","points":"p13150,p13126,p13142,p13143,p13019,p13141,p13121,p13003,p13149,p83106","minute_interval":"5","start_time_stamp":"20220215000000","end_time_stamp":"20220215235900", "ps_id":"1129147"}'
```

Get the household storage report
```
$ ./bin/GoSungrow api get getHouseholdStoragePsReport '{"date_id":"2022","date_type":"4","ps_id":"1129147"}'
```

### Git commands
Record statistics data from iSolarCloud to GitHub. (Will clone if not existing.)

	% GoSungrow git sync 'Updating statistics' statistics

Record all changes made to GitHub.

	% GoSungrow git sync 'Update everything'

Record changes with default commit message.

	% GoSungrow git sync default

Record changes made every 30 minutes.

	% GoSungrow cron run ./30 . . . . git sync default

List files in repo, (identical to ls).

```
% GoSungrow git ls -l
 - rw- r-- r--  admin-mickh admin-mickh     51B 10.Jan'22 13:31 README.md
 - rw- rw- r--  admin-mickh admin-mickh  15.60K 10.Jan'22 13:31 contact.json
 - rw- rw- r--  admin-mickh admin-mickh    496B 10.Jan'22 13:31 department.json
 - rw- rw- r--  admin-mickh admin-mickh 132.14K 10.Jan'22 13:31 device.json
 - rw- rw- r--  admin-mickh admin-mickh   1.29K 10.Jan'22 13:31 domain.json
 - rw- rw- r--  admin-mickh admin-mickh 858.86K 10.Jan'22 13:31 model.json
 - rw- rw- r--  admin-mickh admin-mickh  24.90K 10.Jan'22 13:31 presence.json
 - rw- rw- r--  admin-mickh admin-mickh  53.66K 10.Jan'22 13:31 profile.json
 - rw- rw- r--  admin-mickh admin-mickh     88B 10.Jan'22 13:31 site.json
 - rw- rw- r--  admin-mickh admin-mickh  16.70K 10.Jan'22 13:31 user.json
```

Show changes made to a JSON file.

	% GoSungrow git diff devices.json

### Other available Gitlab commands.
Clone repo.

	% GoSungrow git clone

Pull repo.

	% GoSungrow git pull

Add files to repo.

	% GoSungrow git add .

Push repo.

	% GoSungrow git push

Commit changes to repo.

	% GoSungrow git commit 'this is a commit message'

### Config file.
Show current config.

	% GoSungrow config read

Change diff command used in compares.

	% GoSungrow --diff-cmd='sdiff' config write

Change Git repo directory.

	% GoSungrow --git-dir=/some/other/directory config write

Change Git repo url.

	% GoSungrow --git-url=https://github.com/MickMake/iSolarCloudData config write

Change iSolarCloud API token.

	% GoSungrow --cf-token='this is a token string' config write


## Flags available for all commands:
```
      --config string         GoSungrow: config file. (default "$HOME/.GoSungrow/config.json")
      --debug                 GoSungrow: Debug mode.
      --diff-cmd string       Git: Command for diffs. (default "tkdiff")
      --git-dir string        Git: Local repo directory.
      --git-password string   Git: Repo password.
      --git-repo string       Git: Repo url for updates.
      --git-sshkey string     Git: Repo SSH keyfile.
      --git-token string      Git: Repo token string.
      --git-username string   Git: Repo username.
      --google-sheet string   Google: Sheet URL for updates.
      --host string           iSolarCloud: Provider API URL. (default "https://augateway.isolarcloud.com")
  -p, --password string       iSolarCloud: Extension password.
  -q, --quiet                 iSolarCloud: Silence all messages.
      --timeout duration      iSolarCloud: API timeout. (default 30s)
  -u, --user string           iSolarCloud: Extension username.
```

## Using environment variables instad of flags.
```
+----------------+------------+---------------------+--------------------------------+-----------------------------------------------+
|      FLAG      | SHORT FLAG |   ENVIRONMENT       |          DESCRIPTION           |                    DEFAULT                    |
+----------------+------------+---------------------+--------------------------------+-----------------------------------------------+
| --user         | -u         | SUNGROW_USER         | SUNGRO: API username.          |                                               |
| --password     | -p         | SUNGROW_PASSWORD     | SUNGRO: API password.          |                                               |
| --host         |            | SUNGROW_HOST         | SUNGRO: Provider API URL.      | https://augateway.isolarcloud.com             |
| --timeout      |            | SUNGROW_TIMEOUT      | SUNGRO: API timeout.           | 30s                                           |
| --google-sheet |            | SUNGROW_GOOGLE_SHEET | Google: Sheet URL for updates. |                                               |
| --git-repo     |            | SUNGROW_GIT_REPO     | Git: Repo url for updates.     |                                               |
| --git-dir      |            | SUNGROW_GIT_DIR      | Git: Local repo directory.     |                                               |
| --git-username |            | SUNGROW_GIT_USERNAME | Git: Repo username.            |                                               |
| --git-password |            | SUNGROW_GIT_PASSWORD | Git: Repo password.            |                                               |
| --git-sshkey   |            | SUNGROW_GIT_SSHKEY   | Git: Repo SSH keyfile.         |                                               |
| --git-token    |            | SUNGROW_GIT_TOKEN    | Git: Repo token string.        |                                               |
| --diff-cmd     |            | SUNGROW_DIFF_CMD     | Git: Command for diffs.        | tkdiff                                        |
| --config       |            | SUNGROW_CONFIG       | GoSungrow: config file.         | $HOME/.GoSungrow/config.json                   |
| --debug        |            | SUNGROW_DEBUG        | GoSungrow: Debug mode.          | false                                         |
| --quiet        | -q         | SUNGROW_QUIET        | GoSungrow: Silence all messages.| false                                         |
+----------------+------------+------------------+--------------------------------+--------------------------------------------------+
```

