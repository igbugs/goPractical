package setting

import (
	"github.com/go-ini/ini"
	"logging"
	"time"
)

type App struct {
	UserPasswordSalt string
	IdGenUrl         string
	RunMode          string
	HttpPort         int
	ReadTimeout      int
	WriteTimeout     int
}

var AppSetting = &App{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	DbName   string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		logging.Error("Fail to parse 'config/config.ini'; err:%v", err)
	}

	mapTo("app", AppSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.Error("cfg.MapTo Setting %v , err:%v", v, err)
	}
}
