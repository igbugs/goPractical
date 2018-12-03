package tailf

import (
	config "log_agent/common/config"
	"logging"
	"sync"
)

type TailTaskMgr struct {
	tailTaskMap    map[string]*TailTask
	collectLogList []*config.MsgLogConf
	etcdChan       <-chan []*config.MsgLogConf
}

var (
	tailTaskMgr *TailTaskMgr
)

func Init(collectLogList []*config.MsgLogConf, etcdCh <-chan []*config.MsgLogConf) (err error) {
	tailTaskMgr = &TailTaskMgr{
		collectLogList: collectLogList,
		etcdChan:       etcdCh,
		tailTaskMap:    make(map[string]*TailTask),
	}

	for _, conf := range collectLogList {
		if tailTaskMgr.exist(conf) {
			logging.Debug("init tail task failed, config:%#v, duplicate config", conf)
			continue
		}

		tailTask, err := NewTailTask(conf.Path, conf.ModuleName, conf.Topic)
		if err != nil {
			logging.Error("init tail task failed, config:%#v, err:%#v", conf, err)
			continue
		}

		go tailTask.Run()
		tailTaskMgr.tailTaskMap[tailTask.Key()] = tailTask
	}
	return
}

func (t *TailTaskMgr) exist(conf *config.MsgLogConf) bool {
	for _, tailTask := range t.tailTaskMap {
		if tailTask.Path == conf.Path &&
			tailTask.ModuleName == conf.ModuleName &&
			tailTask.Topic == conf.Topic {
			return true
		}
	}
	return false
}

func (t *TailTaskMgr) listTask() {
	for key, task := range t.tailTaskMap {
		logging.Debug("list tailTask: ===<key:%s task:%v>=== is running", key, task)
	}
}

func (t *TailTaskMgr) run() {
	for {
		t.listTask()
		tmpCollectLogList := <-t.etcdChan
		logging.Debug("the etcd config have changed, new collectLogList:%#v", tmpCollectLogList)

		// 判断是否有新增的日志收集配置
		for _, conf := range tmpCollectLogList {
			// 如果对应的日志收集配置，已经存在。那么不需要做任何事情
			if t.exist(conf) {
				logging.Debug("the tail log task of config:%#v already running", conf)
				continue
			}

			logging.Debug("new tail log task of config:%#v is running", conf)
			// 不存在则实例化新的日志收集任务配置
			tailTask, err := NewTailTask(conf.Path, conf.ModuleName, conf.Topic)
			if err != nil {
				logging.Error("init tail task failed, config:%#v", conf, err)
				continue
			}

			go tailTask.Run()
			tailTaskMgr.tailTaskMap[tailTask.Key()] = tailTask
		}

		// 从已经运行的配置里面，判断 是否存在于 tailTaskMap 里，没有则说明此任务已经删除，可以停止继续收集
		for key, task := range t.tailTaskMap {
			found := false
			for _, conf := range tmpCollectLogList {
				if task.Path == conf.Path &&
					task.ModuleName == conf.ModuleName &&
					task.Topic == conf.Topic {
					found = true
					break
				}
			}

			if found == false {
				// 旧有的collectLog 配置，不在新的 从etcd里取出的配置
				task.Stop()
				delete(t.tailTaskMap, key)
			}
		}
	}
}

func Run(wg *sync.WaitGroup) {
	tailTaskMgr.run()
	wg.Done()
}
