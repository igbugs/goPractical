package scheduler

import "sync"

var (
	mu        sync.Mutex
	jobParams = []JobParam{}
)

type JobParam struct {
	City string
	Pn   int
	Kd   string
}

func NewScheduler() *JobParam  {
	return &JobParam{}
}

func (j *JobParam) Pop() *JobParam {
	mu.Lock()
	length := len(jobParams)
	if length < 1 {
		mu.Unlock()
		return nil
	}

	job := jobParams[length-1]
	jobParams = jobParams[:length-1]

	mu.Unlock()
	return &job
}

func (j *JobParam) Append(city string, pn int, kd string) {
	mu.Lock()
	jobParams = append(jobParams, JobParam{
		City: city,
		Pn:   pn,
		Kd:   kd,
	})

	mu.Unlock()
}

func (j *JobParam) Count() int {
	return len(jobParams)

}
