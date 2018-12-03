package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogBase struct {
	level  int
	module string
}

func (l *LogBase) writeLog(file *os.File, ld *LogData) (err error) {
	_, err = fmt.Fprintf(file, "[%s] [%s] [%s] [%s:%s:%d] %s\n",
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

	return &LogData{
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

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if CheckNotExist(src) {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
//	f, err := os.OpenFile(name, flag, perm)
//	if err != nil {
//		return nil, err
//	}
//
//	return f, err
//}

func MustDir(fileName, filePath string) (err error) {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	if CheckPermission(src) {
		return fmt.Errorf("file.CheckPermission Permission denied src: %c", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	//f, err := Open(src + fileName, os.O_APPEND|os.O_CREATE, 0644)
	//if err != nil {
	//	return nil, fmt.Errorf("fail to OpenFile: %v", err)
	//}

	return nil
}
