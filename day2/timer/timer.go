package timer

import (
	"fmt"
	"time"
)

func testTimer() {
	// NewTimer 只执行一次，之后从 time.C 中取值时会进行阻塞
	timer := time.NewTimer(time.Second)

	for v := range timer.C {
		fmt.Printf("time: %v\n", v)
		timer.Reset(time.Second)
	}

}

func testTicker() {
	// NewTicker 会一直进行在时间间隔之后执行
	ticker := time.NewTicker(time.Second)

	for v := range ticker.C {
		fmt.Printf("time: %v\n", v)
	}

}

func timestampToTime(timestamp int64) {
	t := time.Unix(timestamp, 0)

	fmt.Printf("convert timestamp to time: %v\n", t)
}

func test() {
	start := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		_ = i
	}
	end := time.Now().UnixNano()
	fmt.Printf("Cost time is: %d\n", (end-start)/1000)
}

func main() {
	now := time.Now()
	fmt.Printf("current time is %v\n", now)
	fmt.Printf("%d/%d/%d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("timestamp: %d nanoSec: %d\n", now.Unix(), now.UnixNano())

	testTimer()
	testTicker()
	//time.Sleep(time.Minute)

	timestampToTime(now.Unix())

	str := now.Format("2006/01/02 15:04:05")
	fmt.Printf("str: %v us\n", str)

	test()
}
