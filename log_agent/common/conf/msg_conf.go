package conf

import "time"

// 设置 收集的消息内容配置
type MsgLogConf struct {
	Path string			`json:"path"`
	ModuleName string	`json:"module_name"`
	Topic string		`json:"topic"`
}

type MsgSystemConf struct {
	Interval time.Duration	`json:"interval"`
	Topic string			`json:"topic"`
}

type MsgData struct {
	IP string		`json:"ip"`
	Data string		`json:"data"`
}