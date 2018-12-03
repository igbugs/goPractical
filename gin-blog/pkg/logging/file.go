package logging

import (
	"fmt"
	"gin-blog/pkg/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt)
}

//func getLogFileFullPath() string {
//	prefixPath := getLogFilePath()
//	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
//
//	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
//}
//
//func openLogFile(filePath string) *os.File {
//	_, err := os.Stat(filePath)
//	switch {
//	case os.IsNotExist(err):
//		mkDir()
//	case os.IsPermission(err):
//		log.Fatalf("Permission: %v", err)
//	}
//
//	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE| os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatalf("Fail to OpenFile: %v, %s", err)
//	}
//
//	return file
//}
//
//func mkDir()  {
//	dir, _ := os.Getwd()
//	path := dir + "/" + getLogFilePath()
//	err := os.MkdirAll(path, os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}
