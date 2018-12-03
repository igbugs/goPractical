package logging

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
			level:  level,
			module: module,
		},
	}

	return logger
}

func (c *Console) Init() error {
	return nil
}

func (c *Console) Debug(format string, args ...interface{}) {
	c.outLog(LogLevelDebug, os.Stdout, format, args...)
}

func (c *Console) Trace(format string, args ...interface{}) {
	c.outLog(LogLevelTrace, os.Stdout, format, args...)
}

func (c *Console) Info(format string, args ...interface{}) {
	c.outLog(LogLevelInfo, os.Stdout, format, args...)
}

func (c *Console) Warn(format string, args ...interface{}) {
	c.outLog(LogLevelWarn, os.Stdout, format, args...)
}

func (c *Console) Error(format string, args ...interface{}) {
	c.outLog(LogLevelError, os.Stdout, format, args...)
}

func (c *Console) Fatal(format string, args ...interface{}) {
	c.outLog(LogLevelFatal, os.Stdout, format, args...)
}

func (c *Console) Close() {}
