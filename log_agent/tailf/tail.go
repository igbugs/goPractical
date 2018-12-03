package tailf

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	"log_agent/common/config"
	"log_agent/common/ip"
	"log_agent/kafka"
	"logging"
)

type TailTask struct {
	Path       string
	ModuleName string
	Topic      string
	tailx      *tail.Tail
	ctx        context.Context
	cancel     context.CancelFunc
}

var localIP string

func init() {
	var err error
	localIP, err = ip.GetLocalIP()
	if err != nil {
		logging.Error("get local ip failed, err:%v", err)
		panic(fmt.Sprintf("get local ip failed, err:%v", err))
	}
}

func NewTailTask(path, module, topic string) (tailTask *TailTask, err error) {
	tailTask = &TailTask{}
	err = tailTask.Init(path, module, topic)
	return
}

func (t *TailTask) Init(path, module, topic string) (err error) {
	t.Path = path
	t.ModuleName = module
	t.Topic = topic

	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.tailx, err = tail.TailFile(path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logging.Error("init tail file err: %v", err)
		return
	}
	return
}

func (t *TailTask) Key() string {
	key := fmt.Sprintf("%s_%s_%s", t.Path, t.ModuleName, t.Topic)
	return key
}

func (t *TailTask) Stop() {
	t.cancel()
}

func (t *TailTask) Run() {
	for {
		select {
		case <-t.ctx.Done():
			logging.Warn("task path:%s module:%s topic:%s is exit", t.Path, t.ModuleName, t.Topic)
			return
		case line, ok := <-t.tailx.Lines:
			if !ok {
				logging.Warn("get message from tailf failed")
				continue
			}

			if len(line.Text) == 0 {
				continue
			}

			logging.Debug("line: %s", line.Text)
			data := &config.MsgData{
				IP:   localIP,
				Data: line.Text,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				continue
			}

			msg := &kafka.Message{
				Data:  string(jsonData),
				Topic: t.Topic,
			}

			err = kafka.SendLog(msg)
			if err != nil {
				logging.Warn("send log failed, err: %v", err)
				continue
			}
			logging.Debug("send to kafka success")
		}
	}
}
