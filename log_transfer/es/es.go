package es

import (
	"github.com/olivere/elastic"
	"logging"
	"context"
)

type ESClient struct {
	client    *elastic.Client
	index     string
	threadNum int
	queueSize int
	queue     chan interface{}
}

var esClient = &ESClient{}

func Init(addr string, index string, threadNum, queueSize int) (err error) {
	logging.Debug("init es addr:%s index:%s thread_num:%d queue_size:%d",
		addr, index, threadNum, queueSize)
	client, err := elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		logging.Error("init es client failed, err:%v", err)
		return
	}

	esClient.client = client
	esClient.index = index
	esClient.threadNum = threadNum
	esClient.queueSize = queueSize
	esClient.queue = make(chan interface{}, queueSize)

	for i := 0; i < threadNum; i++ {
		go insertES()
	}
	return
}

func AppendMsg(msg interface{}) {
	logging.Debug("append msg to es queue, msg:%#v", msg)
	esClient.queue <- msg
}

func insertES() {
	for data := range esClient.queue {
		_, err := esClient.client.Index().
			Index(esClient.index).
			Type(esClient.index).
			BodyJson(data).
			Do(context.Background())
		if err != nil {
			logging.Error("do insert es failed, err;%v", err)
			continue
		}
	}
}
