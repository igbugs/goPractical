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
	"time"
	"log_agent/collect_sys_info"
)

var (
	wg sync.WaitGroup
)

func run(sysInfoConf *conf.MsgSystemConf) (err error) {
	wg.Add(2)
	go collect_sys_info.Run(&wg, sysInfoConf.Interval, sysInfoConf.Topic)
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

	systemInfoKey := fmt.Sprintf(conf.AppEtcdSetting.SystemInfoKey, localIP)
	systemInfoConf, err := etcd.GetSystemInfoConfig(systemInfoKey)
	if err != nil {
		systemInfoConf = &conf.MsgSystemConf{
			Topic: "collect_system_info",
			Interval: 5 * time.Second,
		}
		logging.Error("get collect system info config from etcd failed, use default conf: %#v, err: %v",
			systemInfoConf, err)
	}

	err = run(systemInfoConf)
	if err != nil {
		logging.Error("main.run failed, err:%v", err)
		return
	}
	logging.Debug("main.run finished")
}
