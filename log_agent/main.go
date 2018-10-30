package main

import (
	"fmt"
	"log_agent/common/conf"
	"log_agent/common/ip"
	"log_agent/common/logs"
	"log_agent/etcd"
	"log_agent/kafka"
	"log_agent/tailf"
	"logging"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

func run() (err error) {
	wg.Add(1)
	go tailf.Run(&wg)
	wg.Wait()
	return
}

func main() {
	localIP, err := ip.GetLocalIP()
	if err != nil {
		logging.Error("get local ip failed, err:%v", err)
		panic(fmt.Sprintf("get local ip failed, err:%v", err))
	}

	err = conf.Init()
	if err != nil {
		logging.Fatal("init config failed, err: %v", err)
		panic("init config failed")
	}

	err = logs.Init()
	if err != nil {
		logging.Fatal("init logs failed, err: %v", err)
		panic("init logs failed")
	}

	kafkaAddr := strings.Split(conf.AppKafkaSetting.Address, ",")
	err = kafka.Init(kafkaAddr, conf.AppKafkaSetting.QueueSize)
	if err != nil {
		logging.Fatal("init kafka client failed, err: %v", err)
		panic("init kafka client failed")
	}
	logging.Debug("init kafka success, address:%v", kafkaAddr)

	logging.Debug("init kafka client success")

	etcdKey := fmt.Sprintf(conf.AppEtcdSetting.EtcdKey, localIP)
	logging.Debug("etcd key is %v", etcdKey)

	etcdAddr := strings.Split(conf.AppEtcdSetting.Address, ",")
	err = etcd.Init(etcdAddr, etcdKey)
	if err != nil {
		panic(fmt.Sprintf("init etcd client failed, err:%v", err))
	}
	logging.Debug("init etcd success, address:%v", etcdAddr)

	collectLogConf, err := etcd.GetConfig(etcdKey)
	logging.Debug("etcd conf:%#v", collectLogConf)

	watchCHan := etcd.Watch()
	err = tailf.Init(collectLogConf, watchCHan)
	if err != nil {
		panic(fmt.Sprintf("init tailf client failed, err:%v", err))
	}
	logging.Debug("init tailf client success")

	err = run()
	if err != nil {
		logging.Error("main.run failed, err:%v", err)
		return
	}
	logging.Debug("main.run finished")
}
