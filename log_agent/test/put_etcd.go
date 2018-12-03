package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log_agent/common/config"
	"log_agent/common/ip"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err: ", err)
	}

	fmt.Println("connect success")
	defer cli.Close()

	logconf := &config.MsgLogConf{
		Path:       "/tmp/a.log",
		ModuleName: "nginx_log",
		Topic:      "nginx_log",
	}
	var logconfs []*config.MsgLogConf
	logconfs = append(logconfs, logconf)
	jsonData, err := json.Marshal(logconfs)
	fmt.Println(string(jsonData))
	if err != nil {
		fmt.Println("json marshal failed, err: ", err)
	}

	localIP, errno := ip.GetLocalIP()
	if errno != nil {
		fmt.Printf("get local ip failed, err:%v", errno)
	}

	key := fmt.Sprintf("/logagent/%s/config", localIP)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cli.Put(ctx, key, string(jsonData))
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Get(ctx, key)
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	cancel()

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
