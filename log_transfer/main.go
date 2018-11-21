package main

import (
	"log_transfer/common/config"
	"logging"
	"log_transfer/common/logs"
	"log_transfer/es"
	"log_transfer/kafka"
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

	err = es.Init(conf.ESSetting.Addr, conf.ESSetting.Index, conf.ESSetting.ThreadNum, conf.ESSetting.QueueSize)
	if err != nil {
		logging.Error("init es failed, err:%v", err)
		return
	}

	err = kafka.Init(conf.KafkaSetting.Addr, conf.KafkaSetting.Topic)
	if err != nil {
		logging.Error("init kafka failed, err:%v", err)
		return
	}
	select {}
}
