package logging

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

	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})

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

func Debug(fmt string, args ...interface{}) {
	logger.Debug(fmt, args...)
}

func Trace(fmt string, args ...interface{}) {
	logger.Trace(fmt, args...)
}

func Info(fmt string, args ...interface{}) {
	logger.Info(fmt, args...)
}

func Warn(fmt string, args ...interface{}) {
	logger.Warn(fmt, args...)
}

func Error(fmt string, args ...interface{}) {
	logger.Error(fmt, args...)
}

func Fatal(fmt string, args ...interface{}) {
	logger.Fatal(fmt, args...)
}

func Close() {
	logger.Close()
}

//func SetLevel(level int) {
//	logger.SetLevel(level)
//}

