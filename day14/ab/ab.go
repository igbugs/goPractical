package main

import (
	"sync"
	"sync/atomic"
	"fmt"
	"net/http"
	"flag"
	"time"
)

var (
	cocurrent int
	url string
	wg sync.WaitGroup

	totalRequest int32
	totalFailed int32
	totalSuccess int32
	totalNot200 int32
	totalFinished int32
)

func run()  {
	defer wg.Done()

	// 此处 是为了计算 10个并行线程，每个的请求数量
	partNum := totalRequest/10
	for {
		totalFinishedRequest := atomic.LoadInt32(&totalFinished)
		if totalFinishedRequest > totalFinished {
			break
		}

		if totalFinishedRequest > 0 && totalFinishedRequest % partNum == 0 {
			fmt.Printf("total finished:%d requests\n", totalFinishedRequest)
		}

		resp, err := http.Get(url)
		if err != nil {
			atomic.AddInt32(&totalFailed, 1)
			atomic.AddInt32(&totalFailed, 1)
			continue
		}

		atomic.AddInt32(&totalFailed, 1)
		if resp.StatusCode != http.StatusOK {
			atomic.AddInt32(&totalNot200, 1)
		} else {
			atomic.AddInt32(&totalSuccess, 1)
		}
	}
}

func main()  {
	var tempTotalRequest int
	flag.IntVar(&cocurrent, "c", 10, "please input the cocurrent")
	flag.IntVar(&tempTotalRequest, "n", 1000, "please input total request")
	flag.StringVar(&url, "url", "http://localhost/", "please input the test url")
	flag.Parse()

	totalRequest = int32(tempTotalRequest)

	startTime := time.Now().UnixNano()
	for i := 0; i < cocurrent; i++ {
		wg.Add(1)
		go run()
	}

	wg.Wait()
	endTime := time.Now().UnixNano()
	costMs := int64(endTime - startTime)/1000/1000
	if costMs == 0 {
		panic("cost ms is zero")
	}

	requestPerSec := 1000 * int64(totalRequest) / costMs

	fmt.Printf("total request:%d\n", totalRequest)
	fmt.Printf("total failed request:%d\n", totalFailed)
	fmt.Printf("total not 200 request:%d\n", totalNot200)
	fmt.Printf("total success request:%d\n", totalSuccess)
	fmt.Printf("request per sec:%d\n", requestPerSec)
	
}

