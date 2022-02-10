# GoSungro - Over The Wire SUNGRO to Gitlab syncing tool.

This tool does several things:
1. Pull a Gitlab repo that holds SUNGRO data.
2. Fetch all data available from the SUNGRO.
3. Save this data as a JSON file.
4. Commit changes to the Gitlab repo.
5. Push changes to remote.

It is intended to provide full revision history for any changes made to the SUNGRO.

## Use case example:
### Record changes made to user data on SUNGRO. (Will clone if not existing.)

	% GoSungro sync 'Updating just user data' users

### Record changes made to all SUNGRO manually.

	% GoSungro sync 'Updated everything'

### Record changes made to the SUNGRO with default commit message.

	% GoSungro sync default

### Record changes made to the SUNGRO every 30 minutes.

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
  -d, --domain string         SUNGRO: Default domain. (default "mickmake.com")
      --git-dir string        Git: Local repo directory.
      --git-password string   Git: Repo password.
      --git-repo string       Git: Repo url for updates.
      --git-sshkey string     Git: Repo SSH keyfile.
      --git-token string      Git: Repo token string.
      --git-username string   Git: Repo username.
      --google-sheet string   Google: Sheet URL for updates.
      --host string           SUNGRO: Provider API URL. (default "https://augateway.isolarcloud.com")
  -i, --id string             SUNGRO: API client id.
  -p, --password string       SUNGRO: Extension password.
  -q, --quiet                 GoSungro: Silence all messages.
  -s, --secret string         SUNGRO: API client secret.
      --timeout duration      SUNGRO: API timeout. (default 30s)
  -u, --user string           SUNGRO: Extension username.
```

## Using environment variables instad of flags.
```
+----------------+------------+------------------+--------------------------------+-----------------------------------------------+
|      FLAG      | SHORT FLAG |   ENVIRONMENT    |          DESCRIPTION           |                    DEFAULT                    |
+----------------+------------+------------------+--------------------------------+-----------------------------------------------+
| --user         | -u         | SUNGRO_USER         | SUNGRO: Extension username.       |                                               |
| --password     | -p         | SUNGRO_PASSWORD     | SUNGRO: Extension password.       |                                               |
| --id           | -i         | SUNGRO_ID           | SUNGRO: API client id.            |                                               |
| --secret       | -s         | SUNGRO_SECRET       | SUNGRO: API client secret.        |                                               |
| --domain       | -d         | SUNGRO_DOMAIN       | SUNGRO: Default domain.           | mickmake.com                                  |
| --host         |            | SUNGRO_HOST         | SUNGRO: Provider API URL.         | https://augateway.isolarcloud.com/v1/userService/login
| --timeout      |            | SUNGRO_TIMEOUT      | SUNGRO: API timeout.              | 30s                                           |
| --google-sheet |            | SUNGRO_GOOGLE_SHEET | Google: Sheet URL for updates. |                                               |
| --git-repo     |            | SUNGRO_GIT_REPO     | Git: Repo url for updates.     |                                               |
| --git-dir      |            | SUNGRO_GIT_DIR      | Git: Local repo directory.     |                                               |
| --git-username |            | SUNGRO_GIT_USERNAME | Git: Repo username.            |                                               |
| --git-password |            | SUNGRO_GIT_PASSWORD | Git: Repo password.            |                                               |
| --git-sshkey   |            | SUNGRO_GIT_SSHKEY   | Git: Repo SSH keyfile.         |                                               |
| --git-token    |            | SUNGRO_GIT_TOKEN    | Git: Repo token string.        |                                               |
| --diff-cmd     |            | SUNGRO_DIFF_CMD     | Git: Command for diffs.        | tkdiff                                        |
| --config       |            | SUNGRO_CONFIG       | GoSungro: config file.            | $HOME/.GoSungro/config.json                      |
| --debug        |            | SUNGRO_DEBUG        | GoSungro: Debug mode.             | false                                         |
| --quiet        | -q         | SUNGRO_QUIET        | GoSungro: Silence all messages.   | false                                         |
+----------------+------------+------------------+--------------------------------+-----------------------------------------------+
```

