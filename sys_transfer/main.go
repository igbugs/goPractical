package main

import (
	"logging"
	"sys_transfer/common/config"
	"sys_transfer/common/logs"
	"sys_transfer/influxdb"
	"sys_transfer/kafka"
)

func main() {
	err := conf.Init()
	if err != nil {
		logging.Fatal("init config failed, err: %v", err)
		panic("init config failed")
	}

	err = logs.Init()
	if err != nil {
		logging.Fatal("init logs failed, err: %v", err)
		panic("init logs failed")
	}

	err = influxdb.Init(conf.InfluxSetting.Addr, conf.InfluxSetting.QueueSize, conf.InfluxSetting.ThreadNum)
	if err != nil {
		logging.Error("init influxdb failed, err:%v", err)
		return
	}

	err = kafka.Init(conf.KafkaSetting.Addr, conf.KafkaSetting.Topic)
	if err != nil {
		logging.Error("init kafka failed, err:%v", err)
		return
	}
	select {}
}
