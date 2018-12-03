package conf

import (
	"github.com/go-ini/ini"
	"logging"
)

//type AppConf struct {
//}

type KafkaConf struct {
	Addr  string `ini:"addr"`
	Topic string `ini:"topic"`
}

type ESConf struct {
	Addr      string `ini:"addr"`
	Index     string `ini:"index"`
	ThreadNum int    `ini:"thread_num"`
	QueueSize int    `ini:"queue_size"`
}

type LogConf struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType  string `ini:"log_type"`
	Module   string `ini:"module"`
}

var (
	cfg          *ini.File
	LogSetting   = &LogConf{}
	KafkaSetting = &KafkaConf{}
	ESSetting    = &ESConf{}
)

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.Fatal("cfg.MapTo Setting %v , err: %v", v, err)
	}
}

func Init() (err error) {
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'config/config.ini'; %v", err)
		return
	}

	mapTo("kafka", KafkaSetting)
	mapTo("logs", LogSetting)
	mapTo("es", ESSetting)

	logging.Debug("load config success")
	return
}
