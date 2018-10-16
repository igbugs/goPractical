package main

import(
	"fmt"
	"sync"
)


// 通过建立任务的 channel，将不同的类型的任务放入到管道内，执行任务的时候，调用任务的 process 函数，进行任务的处理
func initProgram(threadNum, chanSize int, waitGroup *sync.WaitGroup) (imageChan chan *Task, err error) {

	if chanSize <= 0 || threadNum <= 0 {
		err = fmt.Errorf("invalid parameter")
		return
	}

	imageChan = make(chan *Task, chanSize)
	for i := 0; i <threadNum; i++ {
		waitGroup.Add(1)
		go procImage(imageChan, waitGroup)
	}

	return
}

// 监听发放任务的管道imageChan，获取任务进行处理
func procImage(imageChan chan *Task, wg* sync.WaitGroup) {
	for task := range imageChan {
		err := task.Process()
		if err != nil {
			fmt.Printf("process task:%#v failed, err:%v\n", task, err)
			continue
		}

		fmt.Printf("process task:%#v succ\n", task)
	}
	wg.Done()
}