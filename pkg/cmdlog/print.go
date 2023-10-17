package cmdlog

import (
	"fmt"
	"strings"
)

func New(level string) Log {
	ret := Log{}
	ret.SetLogLevel(level)
	return ret
}

func (l *Log) SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case LogLevelDebugStr:
		l.logLevel = LogLevelDebug
	case LogLevelInfoStr:
		l.logLevel = LogLevelInfo
	case LogLevelWarningStr:
		l.logLevel = LogLevelWarning
	case LogLevelErrorStr:
		l.logLevel = LogLevelError
	default:
		LogPrintDate("Unknown log level, setting to default.")
		l.logLevel = LogLevelInfo
	}
}

func (l *Log) GetLogLevel() string {
	var ret string
	switch l.logLevel {
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

func (l *Log) Debug(format string, args ...interface{}) {
	if l.IsDebug() {
		LogPrintDate("DEBUG: "+format, args...)
	}
}

func (l *Log) Info(format string, args ...interface{}) {
	if l.IsInfo() {
		LogPrintDate("INFO: "+format, args...)
	}
}

func (l *Log) Warning(format string, args ...interface{}) {
	if l.IsWarning() {
		LogPrintDate("WARNING: "+format, args...)
	}
}

func (l *Log) Error(format string, args ...interface{}) {
	if l.IsError() {
		LogPrintDate("ERROR: "+format, args...)
	}
}

func (l *Log) PlainDebug(format string, args ...interface{}) {
	if l.IsDebug() {
		fmt.Print(LogSprintf(format, args...))
	}
}

func (l *Log) PlainInfo(format string, args ...interface{}) {
	if l.IsInfo() {
		fmt.Print(LogSprintf(format, args...))
	}
}

func (l *Log) PlainWarning(format string, args ...interface{}) {
	if l.IsWarning() {
		fmt.Print(LogSprintf(format, args...))
	}
}

func (l *Log) PlainError(format string, args ...interface{}) {
	if l.IsError() {
		fmt.Print(LogSprintf(format, args...))
	}
}
