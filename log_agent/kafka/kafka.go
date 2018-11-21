package kafka

import (
	"github.com/Shopify/sarama"
	"logging"
	"fmt"
)

type Message struct {
	Data  string
	Topic string
}

var (
	client  sarama.SyncProducer
	msgChan chan *Message
)

func Init(addr []string, chanSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addr, config)
	logging.Debug("the address is %v", addr)
	if err != nil {
		logging.Error("new producer failed, err: %v", err)
		return
	}
	// 初始化完成后不能进行 client 的关闭，后续会一直进行使用
	//defer client.Close()

	msgChan = make(chan *Message, chanSize)
	go sendKafka()
	return
}

func SendLog(msg *Message) (err error) {
	if len(msg.Data) == 0 {
		return
	}

	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("msgChan is full")
	}

	return
}

func sendKafka() {
	for msg := range msgChan {
		kafkaMsg := &sarama.ProducerMessage{}
		kafkaMsg.Topic = msg.Topic
		kafkaMsg.Value = sarama.StringEncoder(msg.Data)

		pid, offset, err := client.SendMessage(kafkaMsg)
		if err != nil {
			logging.Error("send message to kafka failed, err: %v", err)
			continue
		}
		logging.Debug("pid: %v offset: %v", pid, offset)
	}
}
