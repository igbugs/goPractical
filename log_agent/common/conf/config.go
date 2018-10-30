package conf

import (
		"github.com/go-ini/ini"
	"logging"
)

// app 从配置文件读取配置
//type AppConfig struct {
//	Logs AppLogConf		`ini:"logs"`
//	Kafka AppKafkaConf	`ini:"kafka"`
//	Etcd AppEtcdConf	`ini:"etcd"`
//}

type AppLogConf struct {
	LogLevel string		`ini:"log_level"`
	Filename string		`ini:"filename"`
	LogType string		`ini:"log_type"`
	Module string		`ini:"module"`
}

type AppKafkaConf struct {
	Address string		`ini:"address"`
	QueueSize int		`ini:"queue_size"`
}

type AppEtcdConf struct {
	Address string		`ini:"address"`
	EtcdKey string		`ini:"etcd_key"`
	SystemInfoKey string	`ini:"system_info_key"`
}

var (
	cfg *ini.File
	//AppSet = AppConfig{}
	AppLogSetting = &AppLogConf{}
	AppKafkaSetting = &AppKafkaConf{}
	AppEtcdSetting = &AppEtcdConf{}

)

func mapTo(section string, v interface{})  {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.Fatal("cfg.MapTo Setting %v , err: %v", v, err)
	}
}

func Init() (err error) {
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'conf/conf.ini'; %v", err)
		return
}

	mapTo("kafka", AppKafkaSetting)
	mapTo("logs", AppLogSetting)
	mapTo("etcd", AppEtcdSetting)

	logging.Debug("load config success")
	return
}

