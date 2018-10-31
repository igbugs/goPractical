package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"log_agent/common/conf"
	"encoding/json"
)

func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.247.133:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		fmt.Println("connect fialed, err: ", err)
	}

	fmt.Println("connect success")
	defer cli.Close()

	logconf := &conf.MsgLogConf{
		Path: "C:/GoProject/a.log",
		ModuleName: "test_log",
		Topic: "test_log",
	}
	var logconfs []*conf.MsgLogConf
	logconfs = append(logconfs, logconf)
	jsonData, err := json.Marshal(logconfs)
	if err != nil {
		fmt.Println("json marshal failed, err: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	_, err = cli.Put(ctx, "/logagent/172.16.30.251/conf", string(jsonData))
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(),  3 * time.Second)
	resp, err := cli.Get(ctx, "/logagent/172.16.30.251/conf/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	cancel()

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
