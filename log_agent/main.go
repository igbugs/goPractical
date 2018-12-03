package main

import (
	"fmt"
	"log_agent/collect_sys_info"
	"log_agent/common/config"
	"log_agent/common/ip"
	"log_agent/common/logs"
	"log_agent/etcd"
	"log_agent/kafka"
	"log_agent/tailf"
	"logging"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func runCollectLog() (err error) {
	wg.Add(1)
	go tailf.Run(&wg)
	return
}


func runSysInfo(sysInfoConf *config.MsgSystemConf) (err error) {
	wg.Add(1)
	go collect_sys_info.Run(&wg, sysInfoConf.Interval, sysInfoConf.Topic)
	return
}

func main() {
	localIP, err := ip.GetLocalIP()
	if err != nil {
		logging.Error("get local ip failed, err:%v", err)
		panic(fmt.Sprintf("get local ip failed, err:%v", err))
	}

	err = config.Init()
	if err != nil {
		logging.Fatal("init config failed, err: %v", err)
		panic("init config failed")
	}

	err = logs.Init()
	if err != nil {
		logging.Fatal("init logs failed, err: %v", err)
		panic("init logs failed")
	}

	kafkaAddr := strings.Split(config.AppKafkaSetting.Address, ",")
	err = kafka.Init(kafkaAddr, config.AppKafkaSetting.QueueSize)
	if err != nil {
		logging.Fatal("init kafka client failed, err: %v", err)
		panic("init kafka client failed")
	}
	logging.Debug("init kafka success, address:%v", kafkaAddr)

	logging.Debug("init kafka client success")

	etcdKey := fmt.Sprintf(config.AppEtcdSetting.EtcdKey, localIP)
	logging.Debug("etcd key is %v", etcdKey)

	etcdAddr := strings.Split(config.AppEtcdSetting.Address, ",")
	err = etcd.Init(etcdAddr, etcdKey)
	if err != nil {
		panic(fmt.Sprintf("init etcd client failed, err:%v", err))
	}
	logging.Debug("init etcd success, address:%v", etcdAddr)

	collectLogConf, err := etcd.GetConfig(etcdKey)
	logging.Debug("etcd config:%#v", collectLogConf)

	watchCHan := etcd.Watch()
	err = tailf.Init(collectLogConf, watchCHan)
	if err != nil {
		panic(fmt.Sprintf("init tailf client failed, err:%v", err))
	}
	logging.Debug("init tailf client success")

	systemInfoKey := fmt.Sprintf(config.AppEtcdSetting.SystemInfoKey, localIP)
	systemInfoConf, err := etcd.GetSystemInfoConfig(systemInfoKey)
	if err != nil {
		systemInfoConf = &config.MsgSystemConf{
			Topic:    "collect_system_info",
			Interval: 5 * time.Second,
		}
		logging.Error("get collect system info config from etcd failed, use default config: %#v, err: %v",
			systemInfoConf, err)
	}

	err = runCollectLog()
	if err != nil {
		logging.Error("main.runCollectLog failed, err:%v", err)
		return
	}
	logging.Debug("main.runCollectLog finished")

	err = runSysInfo(systemInfoConf)
	if err != nil {
		logging.Error("main.runSysInfo failed, err:%v", err)
		return
	}
	logging.Debug("main.runSysInfo finished")

	wg.Wait()
}
