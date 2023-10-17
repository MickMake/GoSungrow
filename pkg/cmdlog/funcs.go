package cmdlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/anicoll/gosungrow/pkg/only"
)

var (
	logFile       string
	logFileHandle *os.File
)

func LogFileSet(lfn string) error {
	var err error

	for range only.Once {
		if lfn != logFile {
			err = LogFileClose()
		}
		logFile = lfn

		if logFile == "" {
			log.SetOutput(os.Stdout)
			break
		}

		logFileHandle, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			break
		}
		log.SetOutput(logFileHandle)
		// log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	}

	return err
}

func LogFileClose() error {
	var err error
	if logFileHandle != nil {
		err = logFileHandle.Close()
		logFileHandle = nil
		logFile = ""
		return err
	}
	return err
}

func timeStamp() string {
	return time.Now().Local().Format(time.UnixDate) + " : "
}

func Printf(format string, args ...interface{}) {
	// format = timeStamp() + format
	log.Printf(format, args...)
}

func LogPrintDate(format string, args ...interface{}) {
	// log.Printf("%s ", TimeNow())
	log.Printf(format, args...)
	// fmt.Println()
}

func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func LogSprintf(format string, args ...interface{}) string {
	// format = timeStamp() + format
	return fmt.Sprintf(format, args...)
}

func LogSprintfDate(format string, args ...interface{}) string {
	ret := fmt.Sprintf("%s ", TimeNow())
	ret += fmt.Sprintf(format, args...)
	return ret
}

type Log struct {
	*os.File
	logLevel int
	err      error
}

func (l *Log) Open(path ...string) error {
	for range only.Once {
		fp := filepath.Join(path...)

		l.File, l.err = os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if l.err != nil {
			break
		}
	}

	return l.err
}

func (l *Log) LogPrintfDate(format string, args ...interface{}) {
	for range only.Once {
		s := LogSprintfDate(format, args...)

		_, l.err = l.File.Write([]byte(s))
		if l.err != nil {
			break
		}
	}
}

func (l *Log) LogPrintf(format string, args ...interface{}) {
	for range only.Once {
		s := LogSprintf(format, args...)

		_, l.err = l.File.Write([]byte(s))
		if l.err != nil {
			break
		}
	}
}

func (l *Log) LogClose() error {
	for range only.Once {
		l.err = l.File.Close()
		if l.err != nil {
			break
		}
	}
	return l.err
}
