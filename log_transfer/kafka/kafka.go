package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log_transfer/es"
	"logging"
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
				var msg = make(map[string]interface{}, 16)

				err = json.Unmarshal(m.Value, &msg)
				if err != nil {
					logging.Error("unmarshal failed, err:%v", err)
					continue
				}

				es.AppendMsg(msg)
				logging.Debug("append msg to es succ, msg:%v", msg)
			}
		}()
	}
	return
}
