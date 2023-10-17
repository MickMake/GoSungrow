package cmdlog

func (l *Log) IsDebug() bool {
	if LogLevelDebug >= l.logLevel {
		return true
	}
	return false
}

func (l *Log) IsInfo() bool {
	if LogLevelInfo >= l.logLevel {
		return true
	}
	return false
}

func (l *Log) IsWarning() bool {
	if LogLevelWarning >= l.logLevel {
		return true
	}
	return false
}

func (l *Log) IsError() bool {
	if LogLevelError >= l.logLevel {
		return true
	}
	return false
}
