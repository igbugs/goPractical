package main

import (
	"go_crawler/lagou_jobs/spider"
	"go_crawler/lagou_jobs/pipeline"
	"sync"
	"logging"
)

var (
	kds = []string{
		"golang",
	}
	citys = []string{
		"北京",
		"上海",
		"广州",
		"深圳",
		"杭州",
		"成都",
	}

	initResults = []spider.InitResult{}
	loopResults = []spider.LoopResult{}
	jobPipeline = pipeline.NewJobPipeline()

	wg sync.WaitGroup
)

func main()  {
	for _, kd := range kds {
		for _, city := range citys {
			wg.Add(1)
			go func(city string, kd string) {
				defer wg.Done()
				initResult, err := spider.InitJobs(city, 1, kd)
				if err != nil {
					logging.Error("spider.InitJobs failed, err:%v", err)
				}

				initResults = append(initResults, initResult...)
				loopResults = append(loopResults, spider.LoopJobs())
			}(city, kd)
		}
	}

	wg.Wait()

	jobPipeline.Push()

	logging.Debug("Init Results: %v", initResults)
	logging.Debug("Loop Results: %v", loopResults)
}
