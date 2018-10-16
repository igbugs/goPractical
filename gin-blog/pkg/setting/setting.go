package setting

import (
	"time"
	"log"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func LoadBase()  {
	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
}

func LoadApp()  {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'app': %v", err)
	}
	PageSize = sec.Key("page_size").MustInt(10)
	JwtSecret = sec.Key("jwt_secret").MustString("!@#$%^&*)(*&")
}

func LoadServer()  {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get sectioin 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")

	HttpPort = sec.Key("http_port").MustInt(8000)

	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(60)) * time.Second
}

func init()  {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}