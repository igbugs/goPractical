package spider

import (
	"time"
	"go_crawler/lagou_jobs/scheduler"
	"go_crawler/lagou_jobs/pipeline"
	"go_crawler/lagou_jobs/downloader"
	"logging"
	"go_crawler/lagou_jobs/pkg/page"
	"errors"
	"fmt"
	"go_crawler/lagou_jobs/pkg/convert"
	)

var (
	delayTime = time.Tick(time.Millisecond * 500)

	jobScheduler = scheduler.NewScheduler()
	jobPipeline  = pipeline.NewJobPipeline()
)

type InitResult struct {
	City       string
	Kd         string
	TotalPage  int
	TotalCount int
}

type LoopResult struct {
	Success int
	Error   int
	Empty  int
	Errors  []string
}

func InitJobs(city string, pn int, kd string) ([]InitResult, error) {
	var (
		jobs       []downloader.Result
		totalPage  int
		totalCount int
		results    []InitResult
		err        error
	)

	jobs, totalPage, totalCount, err = GetJobs(city, pn, kd)
	if err != nil {
		logging.Error("GetJobs func return failed, err:%v", err)
		return nil, err
	}

	results = append(results, InitResult{
		City: city,
		Kd: kd,
		TotalPage: totalPage,
		TotalCount: totalCount,
	})

	for i := 2; i <= totalPage; i++ {
		jobScheduler.Append(city, i, kd)
	}

	jobPipeline.Append(convert.ToPipelineJobs(jobs))

	return results, nil
}

func GetJobs(city string, pn int, kd string) ([]downloader.Result, int, int, error) {
	totalPage := 0
	jobService := downloader.NewJobService(city)
	result, err := jobService.GetJobs(pn, kd)
	if err != nil {
		logging.Error("jobService.GetJobs return failed, err:%v", err)
		return nil, 0, 0, err
	}

	logging.Debug("GetJobs Code:%d, GetJobs City: %s, Pn: %d, Kd: %s", result.Code, city, pn, kd)

	if result.Code == 0 && result.Success == true {
		content := result.Content
		if content.PositionResult.TotalCount > 0 && content.PageSize > 0 {
			totalPage = page.CalculateTotalPage(float64(content.PositionResult.TotalCount),
				float64(content.PageSize))
		}
	} else {
		return nil, 0, 0, errors.New(fmt.Sprintf("GetJobs City: %s, Pn: %d, Kd: %s, Result: %v",
			city, pn, kd, result))
	}

	return result.Content.PositionResult.Result, totalPage, result.Content.PositionResult.TotalCount, nil
}

func LoopJobs() LoopResult  {
	var (
		result LoopResult
		output = jobScheduler.Count()
		params = make(chan []downloader.Result)
	)

	for i := 0; i < output; i++ {
		<-delayTime
		go func() {
			if jobParam := jobScheduler.Pop(); jobParam != nil {
				jobs, _, _, err := GetJobs(jobParam.City, jobParam.Pn, jobParam.Kd)
				if err != nil {
					result.Error++
					result.Errors = append(result.Errors, err.Error())
				} else {
					params <- jobs
				}
			} else {
				result.Empty++
			}
		}()
	}

	LOOP:
		for {
			select {
			case p := <- params:
				result.Success++
				jobPipeline.Append(convert.ToPipelineJobs(p))
			default:
				if result.Success + result.Error + result.Empty >= output {
					logging.Debug("LoopJobs finished, Break...")
					break LOOP
				}
			}
		}
		return result
}