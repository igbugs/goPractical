package logging

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type File struct {
	*LogBase
	filename string
	file     *os.File

	logChan  chan *LogData
	wg       *sync.WaitGroup
	currHour int
}

func NewFile(level int, filename string, module string) Logger {
	logger := &File{
		LogBase: &LogBase{
			level:  level,
			module: module,
		},
		filename: filename,
		logChan:  make(chan *LogData, 10000),
		wg:       &sync.WaitGroup{},
		currHour: time.Now().Hour(),
	}

	logger.wg.Add(1)
	go logger.syncLog()

	return logger
}

func (f *File) outLog(level int, out *os.File, format string, args ...interface{}) {
	if f.level > level {
		return
	}

	logData := f.formatLog(level, f.module, format, args...)
	//f.writeLog(logData)

	select {
	case f.logChan <- logData:
	default:
	}
}

func (f *File) syncLog() {
	for data := range f.logChan {
		f.splitLog()
		f.writeLog(f.file, data)
	}
	f.wg.Done()
}

func (f *File) splitLog() {
	now := time.Now()
	//if now.Hour() == f.currHour {
	//	return
	//}
	//
	//f.currHour = now.Hour()

	if now.Hour() == f.currHour {
		return
	}

	f.file.Sync()
	f.file.Close()

	newFileName := fmt.Sprintf("%s.%04d-%02d-%02d-%02d%s", strings.Trim(f.filename, ".log"),
		now.Year(), now.Month(), now.Day(), now.Hour(), ".log")

	os.Rename(f.filename, newFileName)
	f.currHour = now.Hour()

	f.Init()
}

func (f *File) Init() (err error) {
	f.file, err = os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("openfile failed .....")
		return
	}
	return
}

func (f *File) Debug(format string, args ...interface{}) {
	f.outLog(LogLevelDebug, f.file, format, args...)
}

func (f *File) Trace(format string, args ...interface{}) {
	f.outLog(LogLevelTrace, f.file, format, args...)
}

func (f *File) Info(format string, args ...interface{}) {
	f.outLog(LogLevelInfo, f.file, format, args...)
}

func (f *File) Warn(format string, args ...interface{}) {
	f.outLog(LogLevelWarn, f.file, format, args...)
}

func (f *File) Error(format string, args ...interface{}) {
	f.outLog(LogLevelError, f.file, format, args...)
}

func (f *File) Fatal(format string, args ...interface{}) {
	f.outLog(LogLevelFatal, f.file, format, args...)
}

func (f *File) Close() {
	if f.file != nil {
		f.file.Close()
	}
}
