package loglib

const (
	LogTypeConsole = iota
	LogTypeFile
	LogTypeNet
)

type LogData struct {
	timeStr string
	levelStr string
	module	string
	fileName string
	funcName string
	lineNo	int
	data string
}

var logger Logger = newLogger(LogTypeConsole, LogLevelDebug, "", "default")

type Logger interface {
	Init() error

	LogDebug(format string, args ...interface{})
	LogTrace(format string, args ...interface{})
	LogInfo(format string, args ...interface{})
	LogWarn(format string, args ...interface{})
	LogError(format string, args ...interface{})
	LogFatal(format string, args ...interface{})

	Close()
}

func newLogger(logType, level int, filename string, module string) Logger {
	var logger Logger
	switch logType {
	case LogTypeConsole:
		logger = NewConsole(level, module)
	case LogTypeFile:
		logger = NewFile(level, filename, module)
	default:
		logger = NewFile(level, filename, module)
	}
	return logger
}

func Init(logType, level int, filename string, module string) error {
	logger = newLogger(logType, level, filename, module)
	return logger.Init()
}

func LogDebug(fmt string, args ...interface{}) {
	logger.LogDebug(fmt, args...)
}

func LogTrace(fmt string, args ...interface{}) {
	logger.LogTrace(fmt, args...)
}

func LogInfo(fmt string, args ...interface{}) {
	logger.LogInfo(fmt, args...)
}

func LogWarn(fmt string, args ...interface{}) {
	logger.LogWarn(fmt, args...)
}

func LogError(fmt string, args ...interface{}) {
	logger.LogError(fmt, args...)
}

func LogFatal(fmt string, args ...interface{}) {
	logger.LogFatal(fmt, args...)
}

func Close() {
	logger.Close()
}

//func SetLevel(level int) {
//	logger.SetLevel(level)
//}