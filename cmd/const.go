package cmd

import "time"


//goland:noinspection SpellCheckingInspection
const (
	DefaultBinaryName = "GoSungrow"
	EnvPrefix         = "SunGrow"
	defaultConfigFile = "config.json"
	defaultTokenFile  = "AuthToken.json"

	flagConfigFile = "config"
	flagDebug      = "debug"
	flagQuiet      = "quiet"

	flagApiUrl        = "host"
	flagApiTimeout    = "timeout"
	flagApiUsername   = "user"
	flagApiPassword   = "password"
	flagApiAppKey     = "appkey"
	flagApiLastLogin  = "token-expiry"
	flagApiOutputType = "out"

	flagMqttUsername   = "mqtt-user"
	flagMqttPassword   = "mqtt-password"
	flagMqttHost   = "mqtt-host"
	flagMqttPort   = "mqtt-port"

	flagGoogleSheet       = "google-sheet"
	flagGoogleSheetUpdate = "update"

	flagGitUsername = "git-username"
	flagGitPassword = "git-password"
	flagGitKeyFile  = "git-sshkey"
	flagGitToken    = "git-token"
	flagGitRepo     = "git-repo"
	flagGitRepoDir  = "git-dir"
	flagGitDiffCmd  = "diff-cmd"

	defaultHost      = "https://augateway.isolarcloud.com"
	defaultUsername  = "harry@potter.net"
	defaultPassword  = "hogwarts"
	defaultApiAppKey = "93D72E60331ABDCDC7B39ADC2D1F32B3"

	defaultTimeout = time.Duration(time.Second * 30)
)
