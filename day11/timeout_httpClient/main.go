package main

import (
	"net/http"
	"context"
	"time"
	"fmt"
		"sync"
		"io/ioutil"
)

type RespData struct {
	resp *http.Response
	err error
}

func doCall(ctx context.Context) {
	tr := http.Transport{}
	client := http.Client{
		Transport: &tr,
	}

	respChan := make(chan *RespData, 1)
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		fmt.Printf("new request faild, err: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.Do resp: %v, err: %v\n", resp, err)
		respData := &RespData{
			resp: resp,
			err: err,
		}

		respChan <- respData
		wg.Done()
	}()

	select {
	case <- ctx.Done():
		// 超时的时候，此分支会有空的结构体 传入，标识任务结束
		tr.CancelRequest(req)	// 取消请求
		fmt.Printf("call api timeout\n")
	case result := <- respChan:
		fmt.Printf("call server api succ\n")
		if result.err != nil {
			fmt.Printf("call api failed, err: %v\n", err)
			return
		}

		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp: %v\n", string(data))
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond * 10000)
	defer cancel()
	doCall(ctx)
}
