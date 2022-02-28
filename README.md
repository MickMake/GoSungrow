# GoSungrow - iSolarCloud API written in GoLang.

## What is it?

This GoLang package has a complete, (almost complete), implementation of the iSolarCloud API.
There's been no published specs on this, so I've had to figure it all out based on the [Android app](https://play.google.com/store/apps/details?id=com.isolarcloud.manager), using javascript IDEs and various other means.

Note:
- [iSolarCloud](https://isolarcloud.com) has no interest in developing a public API.
- Their "API" implementation is so broken with security and coding issues, I'm surprised it hasn't been exploited yet.

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/iSolarCloudLogin.png?raw=true)

![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/iSolarCloud.png?raw=true)


## What state is it in?

This is currently under development.

So far I have mapped out all the API calls, but now figuring out JSON POST data formats.

It's tricky as their "API" changes regularly.

I've currently mapped out these API EndPoints. Only about 6%, but most of the API endpoints are a repeat and a lot of the data is repeated.
So, in reality 50% of the critical endpoints are implemented.
```
+-------------------+-------------------+--------------------+------------+
|       AREAS       | ENABLED ENDPOINTS | DISABLED ENDPOINTS | COVERAGE % |
+-------------------+-------------------+--------------------+------------+
| AliSmsService     |                 0 |                  1 | 0.0 %      |
| AppService        |                51 |                523 | 9.8 %      |
| MttvScreenService |                 0 |                 30 | 0.0 %      |
| PowerPointService |                 0 |                  1 | 0.0 %      |
| WebAppService     |                 3 |                186 | 1.6 %      |
| WebIscmAppService |                 0 |                184 | 0.0 %      |
| ----------------  | ----------------  | -----------------  | ---------  |
| Total             |                54 |                925 | 5.8 %      |
+-------------------+-------------------+--------------------+------------+
```

## What does it do?

This GoLang package does several things:
1. Provides ready access to all API calls via a simple get/put framework.
2. Update a GitHub repo with SunGrow PV data, (provide full revision history for any changes made to the SunGrow PV).
3. Update a Google sheet with SunGrow PV data.

## What is the roadmap?

I've planned a number of features, but my main goal is to be able to interface with my HASS instance.
1. MQTT client to push to things like [HomeAssistant](https://www.home-assistant.io/).
2. IFTTT support.
3. Graphing a logging support.


## Use case example:
### Fetch PV data from the API.

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

Produce basic real-time stats
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

### Reporting examples.

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

Produce daily report for template 8040 for date 2022/02/24 spit out CSV to STDOUT.
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
![alt text](https://github.com/MickMake/GoSungrow/blob/master/docs/AppService_queryMutiPointDataList-20220224-8042.png?raw=true)

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

