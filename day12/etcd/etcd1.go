package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
)

func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.20.200:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("content failed, err: ", err)
	}

	fmt.Println("connect success")
	defer cli.Close()
}
