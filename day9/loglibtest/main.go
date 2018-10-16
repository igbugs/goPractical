package main

import (
	"flag"
	"fmt"
	"loglib"
	"time"
)

//func logic(logger loglib.Logger) {
//	logger.LogDebug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
//	logger.LogTrace("dskskdskfksdf")
//	logger.LogInfo("dskskdskfksdf")
//	logger.LogWarn("dskskdskfksdf")
//	logger.LogError("dskskdskfksdf")
//}

func logic()  {
	for {
		loglib.LogDebug("dskskdskfksdf, user_id:%d username:%s", 388338, "username")
		loglib.LogTrace("dskskdskfksdf")
		loglib.LogInfo("dskskdskfksdf")
		loglib.LogWarn("dskskdskfksdf")
		loglib.LogError("dskskdskfksdf")
		loglib.LogFatal("dskskdskfksdf")
		time.Sleep(time.Second * 5)
	}
}

func main() {

	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "file",  "please input logger type")
	flag.Parse()

	var logType int
	if logTypeStr == "file" {
		logType = loglib.LogTypeFile
	} else {
		logType = loglib.LogTypeConsole
	}

	var outpath = "C:/GoProject/Go3Project/src/day9/loglibtest/loglib.log"
	//logger := loglib.NewLogger(logType, loglib.LogLevelDebug, outpath, "LOGLIB_EXAMPLE")
	//err := logger.Init()

	err := loglib.Init(logType, loglib.LogLevelDebug, outpath, "LOGLIB_EXAMPLE")
	if err != nil {
		fmt.Printf("logger init failed\n")
		return
	}
	logic()
}
