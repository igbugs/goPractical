package session

import (
	"fmt"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"web_chat/common/setting"
)

var (
	GlobalSess *session.Manager
)

func init()  {
	config := &session.ManagerConfig{
		CookieName: "sid",
		Gclifetime: 3600,
		Maxlifetime: 3600,
		CookieLifeTime: 3600,
		ProviderConfig: fmt.Sprintf(setting.RedisSetting.Host, ",100,"),
		EnableSetCookie: true,
	}

	var err error
	GlobalSess, err = session.NewManager("redis", config)
	if err != nil {
		panic(fmt.Sprintf("init session failed, err:%v", err))
		return
	}
	go GlobalSess.GC()
}