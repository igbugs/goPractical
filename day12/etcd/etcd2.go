package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.20.200:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("content failed, err: ", err)
	}

	fmt.Println("connect success")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/config/", "sample_value2w3e")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/config/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
