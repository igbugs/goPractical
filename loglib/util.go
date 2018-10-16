package loglib

import (
	"os"
	"fmt"
	"time"
	"runtime"
	"path/filepath"
	)

type LogBase struct {
	level int
	module string
}


func (l *LogBase) writeLog(file *os.File, ld *LogData) (err error) {
	_, err = fmt.Fprintf(file, "[%s] [%s] [%s] [%s:%s:%d] [%s]\n",
		ld.timeStr, ld.levelStr, ld.module,
		ld.fileName, ld.funcName, ld.lineNo,
		ld.data)
	return
}

func (l *LogBase) formatLog(level int, module string, format string, args ...interface{}) *LogData {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000")

	levelStr := getLevel(level)
	fileName, funcName, lineNo := getFuncInfo(5)

	fileName = filepath.Base(fileName)
	data := fmt.Sprintf(format, args...)

	return &LogData {
		timeStr:  timeStr,
		levelStr: levelStr,
		module:   module,
		fileName: fileName,
		lineNo:   lineNo,
		funcName: funcName,
		data:     data,
	}
}

func (l *LogBase) outLog(level int, out *os.File, format string, args ...interface{}) {
	if l.level > level {
		return
	}

	logData := l.formatLog(level, l.module, format, args...)
	l.writeLog(out, logData)
}

func getFuncInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}

	fileName = file
	lineNo = line
	return
}

