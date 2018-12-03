package logs

import (
	"fmt"
	"log_transfer/common/config"
	"logging"
	"path/filepath"
)

func Init() (err error) {
	var logType, level int

	var logTypeMap = map[string]int{
		"console": logging.LogTypeConsole,
		"file":    logging.LogTypeFile,
		"net":     logging.LogTypeNet,
	}
	var levelMap = map[string]int{
		"debug": logging.LogLevelDebug,
		"trace": logging.LogLevelTrace,
		"info":  logging.LogLevelInfo,
		"warn":  logging.LogLevelWarn,
		"error": logging.LogLevelError,
		"fatal": logging.LogLevelFatal,
	}

	if v, ok := logTypeMap[conf.LogSetting.LogType]; ok {
		logType = v
	} else {
		logType = logTypeMap["console"]
	}

	if v, ok := levelMap[conf.LogSetting.LogLevel]; ok {
		level = v
	} else {
		level = levelMap["debug"]
	}

	filePath := conf.LogSetting.Filename
	path, filename := filepath.Split(filePath)

	err = logging.MustDir(filename, path)
	if err != nil {
		fmt.Printf("make dir failed, err: %v\n", err)
	}

	err = logging.Init(logType, level, filePath, conf.LogSetting.Module)
	if err != nil {
		return
	}

	logging.Debug("init logging success")
	return
}
