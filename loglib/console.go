package loglib

import "os"

type Console struct {
	*LogBase
}

func NewConsole(level int, module string) Logger {
	//logger := &Console{}
	//logger.LogBase = &LogBase{
	//	level:  level,
	//	module: module,
	//}

	logger := &Console{
		LogBase: &LogBase{
			level: level,
			module: module,
		},
	}

	return logger
}

func (c *Console) Init() error {
	return nil
}

func (c *Console) LogDebug(format string, args ...interface{}) {
	c.outLog(LogLevelDebug, os.Stdout, format, args...)
}

func (c *Console) LogTrace(format string, args ...interface{}) {
	c.outLog(LogLevelTrace, os.Stdout, format, args...)
}

func (c *Console) LogInfo(format string, args ...interface{}) {
	c.outLog(LogLevelInfo, os.Stdout, format, args...)
}

func (c *Console) LogWarn(format string, args ...interface{}) {
	c.outLog(LogLevelWarn, os.Stdout, format, args...)
}

func (c *Console) LogError(format string, args ...interface{}) {
	c.outLog(LogLevelError, os.Stdout, format, args...)
}

func (c *Console) LogFatal(format string, args ...interface{}) {
	c.outLog(LogLevelFatal, os.Stdout, format, args...)
}

func (c *Console) Close() {}