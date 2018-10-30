package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"log_agent/common/conf"
	"time"
	"logging"
	"context"
	"encoding/json"
	"fmt"
)

type EtcdClient struct {
	client *clientv3.Client
	address []string
	watchKey string
	dataChan chan []*conf.MsgLogConf
}

var (
	etcdClient *EtcdClient
)

func Init(address []string, watchKey string) (err error) {
	etcdClient = &EtcdClient{
		address: address,
		watchKey: watchKey,
		dataChan: make(chan []*conf.MsgLogConf),
	}

	etcdClient.client, err = clientv3.New(clientv3.Config {
		Endpoints: address,
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		logging.Error("create etcd client failed, address:%v err:%v", err, address)
		return
	}

	go etcdClient.watch()
	return
}

func (ec *EtcdClient) watch() {
	for {
		resultCh := ec.client.Watch(context.Background(), ec.watchKey)
		logging.Debug("watch return, resultCh:%v", resultCh)

		for v := range resultCh {
			logging.Debug("watch from resultCh, vaule:%v", v)
			if v.Err() != nil {
				logging.Error("watch key:%s failed, err:%v", ec.watchKey, v.Err())
				continue
			}

			for _, event := range v.Events {
				logging.Debug("eventType:%v key:%s value:%s", event.Type, event.Kv.Key, string(event.Kv.Value))
				var conf []*conf.MsgLogConf
				if event.Type == clientv3.EventTypeDelete {
					// 此时的conf 是空值
					ec.dataChan <- conf
					continue
				}

				err := json.Unmarshal(event.Kv.Value, &conf)
				if err != nil {
					logging.Error("unmarshal etcd value failed, key:%v value:%v, err:%v", ec.watchKey, string(event.Kv.Value), err)
					continue
				}

				ec.dataChan <- conf
			}
		}
	}
}

func Watch() <-chan []*conf.MsgLogConf {
	return etcdClient.dataChan
}

func GetConfig(key string) (conf []*conf.MsgLogConf, err error) {
	resp, err := etcdClient.client.Get(context.Background(), key)
	if err != nil {
		logging.Error("get key:%v from etcd failed, err:%v", key, err)
		return
	}

	if len(resp.Kvs) == 0 {
		logging.Error("get key:%v from etcd failed, len(resp.kvs)=0", key)
		return
	}

	keyVals := resp.Kvs[0]
	logging.Debug("get key:%s from etcd success, key:%s value:%s", key, keyVals.Key, keyVals.Value)

	err = json.Unmarshal(keyVals.Value, &conf)
	if err != nil {
		logging.Error("unmarshal from keyVals.Value failed, err:%v, data:%v", err, string(keyVals.Value))
		return
	}

	logging.Debug("get config from etcd success, conf: %#v", conf)
	return
}

func GetSystemInfoConfig(key string) (conf []*conf.MsgSystemConf, err error) {
	resp, err := etcdClient.client.Get(context.Background(), key)
	if err != nil {
		logging.Error("get key: %v from etcd failed, err: %v", key, err)
		return
	}

	if len(resp.Kvs) == 0 {
		logging.Error("get key: %v from etcd failed, len(resp.Kvs)=0", key)
		err = fmt.Errorf("not found value of %s", key)
		return
	}

	keyVals := resp.Kvs[0]
	logging.Debug("get key:%v from etcd success, key:%v value:%v", key, keyVals.Key, keyVals.Value)

	err = json.Unmarshal(keyVals.Value, &conf)
	if err != nil {
		logging.Error("unmarshal from keyVals.Value failed, err:%v, data:%v", err, string(keyVals.Value))
		return
	}

	logging.Debug("get config from etcd success, conf: %#v", conf)
	return

}

