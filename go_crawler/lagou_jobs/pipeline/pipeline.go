package pipeline

import (
	"sync"
	"go_crawler/lagou_jobs/model"
	"logging"
)

var (
	mu   sync.Mutex
	jobs []LgJob
)

type LgJob struct {
	City     string
	District string

	CompanyShortName string
	CompanyFullName  string
	CompanyLabelList string
	CompanySize      string
	FinanceStage     string

	IndustryField  string
	IndustryLables string

	PositionName      string
	PositionLables    string
	PositionAdvantage string
	WorkYear          string
	Education         string
	Salary            string

	Longitude   float64
	Latitude    float64
	Linestaion string

	CreateTime int64
	AddTime    int64
}

func NewJobPipeline() *LgJob {
	return &LgJob{}
}

func (j *LgJob) Append(js []LgJob) {
	mu.Lock()
	jobs = append(jobs, js...)
	mu.Unlock()
}

func (j *LgJob) Get() []LgJob {
	return jobs
}

func (j *LgJob) Push() error {
	for _, v := range j.Get() {
		if err := model.DB.Create(v).Error; err != nil {
			logging.Error("insert db failed, err: %v", err)
			return err
		}
	}
	return nil
}
