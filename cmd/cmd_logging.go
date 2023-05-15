package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/cmdLog"
	"strings"
)


type Logging struct {
	logLevel int
	logger   cmdLog.Log
}

// -------------------------------------------------------------------------------- //

const (
	LogLevelDebug   = 0
	LogLevelInfo    = iota
	LogLevelWarning = iota
	LogLevelError   = iota

	LogLevelDebugStr   = "debug"
	LogLevelInfoStr    = "info"
	LogLevelWarningStr = "warning"
	LogLevelErrorStr   = "error"
)

func NewLogging(level string) Logging {
	ret := Logging{}
	ret.SetLogLevel(level)
	ret.logger = cmdLog.Log{}

	return ret
}

func (c *Logging) SetLogLevel(level string) {
	switch strings.ToLower(level) {
		case LogLevelDebugStr:
			c.logLevel = LogLevelDebug
		case LogLevelInfoStr:
			c.logLevel = LogLevelInfo
		case LogLevelWarningStr:
			c.logLevel = LogLevelWarning
		case LogLevelErrorStr:
			c.logLevel = LogLevelError
		default:
			cmdLog.LogPrintDate("Unknown log level, setting to default.")
			c.logLevel = LogLevelInfo
	}
}

func (c *Logging) GetLogLevel() string {
	var ret string
	switch c.logLevel {
		case LogLevelDebug:
			ret = LogLevelDebugStr
		case LogLevelInfo:
			ret = LogLevelInfoStr
		case LogLevelWarning:
			ret = LogLevelWarningStr
		case LogLevelError:
			ret = LogLevelErrorStr
		default:
			ret = LogLevelInfoStr
	}
	return ret
}

func (c *Logging) Debug(format string, args ...interface{}) {
	if LogLevelDebug >= c.logLevel {
		cmdLog.LogPrintDate("DEBUG: " + format, args...)
	}
}

func (c *Logging) Info(format string, args ...interface{}) {
	if LogLevelInfo >= c.logLevel {
		cmdLog.LogPrintDate("INFO: " + format, args...)
	}
}

func (c *Logging) Warning(format string, args ...interface{}) {
	if LogLevelWarning >= c.logLevel {
		cmdLog.LogPrintDate("WARNING: " + format, args...)
	}
}

func (c *Logging) Error(format string, args ...interface{}) {
	if LogLevelError >= c.logLevel {
		cmdLog.LogPrintDate("ERROR: " + format, args...)
	}
}

func (c *Logging) PlainDebug(format string, args ...interface{}) {
	if LogLevelDebug >= c.logLevel {
		fmt.Print(cmdLog.LogSprintf(format, args...))
	}
}

func (c *Logging) PlainInfo(format string, args ...interface{}) {
	if LogLevelInfo >= c.logLevel {
		fmt.Print(cmdLog.LogSprintf(format, args...))
	}
}

func (c *Logging) PlainWarning(format string, args ...interface{}) {
	if LogLevelWarning >= c.logLevel {
		fmt.Print(cmdLog.LogSprintf(format, args...))
	}
}

func (c *Logging) PlainError(format string, args ...interface{}) {
	if LogLevelError >= c.logLevel {
		fmt.Print(cmdLog.LogSprintf(format, args...))
	}
}
