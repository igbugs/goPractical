package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"sync"
	)

var wg sync.WaitGroup

func main()  {
	consumer, err := sarama.NewConsumer([]string{"192.168.20.200:9092"}, nil)
	if err != nil {
		fmt.Println("failed to start consumer: %v", err)
		return
	}

	fmt.Println("connect success")
	partitions, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Printf("get partition failed, err: %v\n", err)
		return
	}

	for _, p := range partitions {
		pc, err := consumer.ConsumePartition("nginx_log", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		//defer func(p *sarama.PartitionConsumer) {
		//	(*p).AsyncClose()
		//}(&pc)

		wg.Add(1)
		go func() {
			for m := range pc.Messages() {
				//fmt.Printf("message: %v, test:%v\n", m, string(m.Value))
				fmt.Printf("test:%v\n", string(m.Value))
			}
			wg.Done()
		}()
	}

	wg.Wait()

}
