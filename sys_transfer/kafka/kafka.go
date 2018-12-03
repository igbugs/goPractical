package kafka

import (
	"github.com/Shopify/sarama"
	"logging"
	"sys_transfer/influxdb"
)

var consumer sarama.Consumer

func Init(addr string, topic string) (err error) {
	consumer, err = sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		logging.Error("new kafka consumer failed, err:%v", err)
		return
	}

	logging.Debug("connect to kafka succ")

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		logging.Error("get partitions failed, err:%v", err)
		return
	}

	logging.Debug("get partitions succ, partition:%#v", partitions)

	for _, p := range partitions {
		pc, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			logging.Error("consumer partition failed, err:%v", err)
			continue
		}

		go func() {
			messageChan := pc.Messages()
			for m := range messageChan {
				logging.Debug("recv from kafka, consumer msg body %v, text:%v", m, string(m.Value))

				influxdb.AppendMsg(string(m.Value))
				logging.Debug("append msg to influxdb channel succ, msg:%v", string(m.Value))
			}
		}()
	}
	return
}
