package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log_test"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")
	client, err := sarama.NewSyncProducer([]string{"192.168.247.133:9092"}, config)
	//client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		log.Println(err)
		return
	}
	//defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		log.Println(err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
