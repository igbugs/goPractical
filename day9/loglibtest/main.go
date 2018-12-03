package main

import (
	"flag"
	"fmt"
	"logging"
	"time"
)

//func logic(logger loglib.Logger) {
//	logger.LogDebug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
//	logger.LogTrace("dskskdskfksdf")
//	logger.LogInfo("dskskdskfksdf")
//	logger.LogWarn("dskskdskfksdf")
//	logger.LogError("dskskdskfksdf")
//}

func logic() {
	for {
		logging.Debug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
		logging.Trace("dskskdskfksdf")
		logging.Info("dskskdskfksdf")
		logging.Warn("dskskdskfksdf")
		logging.Error("dskskdskfksdf")
		logging.Fatal("dskskdskfksdf")
		time.Sleep(time.Second * 5)
	}
}

func main() {

	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "file", "please input logger type")
	flag.Parse()

	var logType int
	if logTypeStr == "file" {
		logType = logging.LogTypeFile
	} else {
		logType = logging.LogTypeConsole
	}

	var outpath = "C:/GoProject/Go3Project/src/day9/loglibtest/loglib.log"
	//logger := loglib.NewLogger(logType, loglib.LogLevelDebug, outpath, "LOGLIB_EXAMPLE")
	//err := logger.Init()

	err := logging.Init(logType, logging.LogLevelDebug, outpath, "LOGLIB_EXAMPLE")
	if err != nil {
		fmt.Printf("logger init failed\n")
		return
	}
	logic()
}
