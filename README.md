# GoSungro - iSolarCloud API written in GoLang.

## What is it?

This GoLang package has a complete, (almost complete), implementation of the iSolarCloud API.
There's been no published specs on this, so I've had to figure it all out based on the Android app, using javascript IDEs and various other means.
Note: [iSolarCloud](https://isolarcloud.com) has no interest in developing a public API. In fact their implementation is so broken with security and coding issues, I'm surprised it hasn't been exploited yet.

![alt text](https://github.com/MickMake/GoSungro/blob/master/docs/iSolarCloudLogin.png?raw=true)

![alt text](https://github.com/MickMake/GoSungro/blob/master/docs/iSolarCloud.png?raw=true)


## What state is it in?

This is currently under development. So far I have mapped out all the API calls, but now figuring out JSON POST data formats.
It's tricky as their "API" changes regularly.


## What does it do?

This GoLang package does several things:
1. Update a GitHub repo with SunGro PV data, (provide full revision history for any changes made to the SunGro PV).
2. Update a Google sheet with SunGro PV data.
3. Provides ready access to all API calls via a simple get/put framework.

To be added:
1. MQTT client to push to things like HomeAssistant.


## Use case example:
### Record statistics data from SUNGRO to GitHub. (Will clone if not existing.)

	% GoSungro sync 'Updating statistics' statistics

### Record all changes made to GitHub.

	% GoSungro sync 'Update everything'

### Record changes with default commit message.

	% GoSungro sync default

### Record changes made every 30 minutes.

	% GoSungro cron run ./30 . . . . sync default

### List files in repo, (identical to ls).

```
% GoSungro git ls -l
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

### Show changes made to a JSON file.

	% GoSungro git diff devices.json

### Other available Gitlab commands.
Clone repo.

	% GoSungro git clone

Pull repo.

	% GoSungro git pull

Add files to repo.

	% GoSungro git add .

Push repo.

	% GoSungro git push

Commit changes to repo.

	% GoSungro git commit 'this is a commit message'

### Config file.
Show current config.

	% GoSungro config read

Change diff command used in compares.

	% GoSungro --diff-cmd='sdiff' config write

Change Git repo directory.

	% GoSungro --git-dir=/some/other/directory config write

Change Git repo url.

	% GoSungro --git-url=https://github.com/MickMake/iSolarCloudData config write

Change SUNGRO API token.

	% GoSungro --cf-token='this is a token string' config write


## Flags available for all commands:
```
      --config string         GoSungro: config file. (default "$HOME/.GoSungro/config.json")
      --debug                 GoSungro: Debug mode.
      --diff-cmd string       Git: Command for diffs. (default "tkdiff")
      --git-dir string        Git: Local repo directory.
      --git-password string   Git: Repo password.
      --git-repo string       Git: Repo url for updates.
      --git-sshkey string     Git: Repo SSH keyfile.
      --git-token string      Git: Repo token string.
      --git-username string   Git: Repo username.
      --google-sheet string   Google: Sheet URL for updates.
      --host string           SUNGRO: Provider API URL. (default "https://augateway.isolarcloud.com")
  -p, --password string       SUNGRO: Extension password.
  -q, --quiet                 SUNGRO: Silence all messages.
      --timeout duration      SUNGRO: API timeout. (default 30s)
  -u, --user string           SUNGRO: Extension username.
```

## Using environment variables instad of flags.
```
+----------------+------------+---------------------+--------------------------------+-----------------------------------------------+
|      FLAG      | SHORT FLAG |   ENVIRONMENT       |          DESCRIPTION           |                    DEFAULT                    |
+----------------+------------+---------------------+--------------------------------+-----------------------------------------------+
| --user         | -u         | SUNGRO_USER         | SUNGRO: API username.          |                                               |
| --password     | -p         | SUNGRO_PASSWORD     | SUNGRO: API password.          |                                               |
| --host         |            | SUNGRO_HOST         | SUNGRO: Provider API URL.      | https://augateway.isolarcloud.com             |
| --timeout      |            | SUNGRO_TIMEOUT      | SUNGRO: API timeout.           | 30s                                           |
| --google-sheet |            | SUNGRO_GOOGLE_SHEET | Google: Sheet URL for updates. |                                               |
| --git-repo     |            | SUNGRO_GIT_REPO     | Git: Repo url for updates.     |                                               |
| --git-dir      |            | SUNGRO_GIT_DIR      | Git: Local repo directory.     |                                               |
| --git-username |            | SUNGRO_GIT_USERNAME | Git: Repo username.            |                                               |
| --git-password |            | SUNGRO_GIT_PASSWORD | Git: Repo password.            |                                               |
| --git-sshkey   |            | SUNGRO_GIT_SSHKEY   | Git: Repo SSH keyfile.         |                                               |
| --git-token    |            | SUNGRO_GIT_TOKEN    | Git: Repo token string.        |                                               |
| --diff-cmd     |            | SUNGRO_DIFF_CMD     | Git: Command for diffs.        | tkdiff                                        |
| --config       |            | SUNGRO_CONFIG       | GoSungro: config file.         | $HOME/.GoSungro/config.json                   |
| --debug        |            | SUNGRO_DEBUG        | GoSungro: Debug mode.          | false                                         |
| --quiet        | -q         | SUNGRO_QUIET        | GoSungro: Silence all messages.| false                                         |
+----------------+------------+------------------+--------------------------------+--------------------------------------------------+
```

