package unify

import "time"

const (
	flagConfigFile = "config"
	flagDebug      = "debug"
	flagQuiet      = "quiet"
	flagTimeout    = "timeout"

	defaultConfig  = ""
	defaultTimeout = time.Second * 30
	defaultDebug   = false
	defaultQuiet   = false
)
