package main

import (
	"fmt"
	"time"
)

//var sendHis = make(map[int][]string)

func main() {
	//a := make([]int, 0)
	////a := []int{}
	//for i := 0; i <= 100; i++ {
	//	a = append(a, i)
	//	fmt.Println(len(a), cap(a))
	//}

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(r.Int())
	//}

	//var msg = make(map[string]interface{}, 16)
	//
	//m := []byte(`{"ip":"192.168.137.108","data":"ooooooooooooooooooooooo"}`)
	//err := json.Unmarshal(m, &msg)
	//if err != nil {
	//	logging.Error("unmarshal failed, err:%v", err)
	//}
	//
	//fmt.Println(msg["ip"])
	//
	//
	////var a uint16 = 11111
	////fmt.Printf("uint16 length:%v", len([]byte(a)))
	//
	//buf := make([]byte, 6)
	//fmt.Printf("buf length: %v", len(buf))

	//strList := []string{"qq", "ww", "ee"}
	//for i := 0; i <= 10; i++ {
	//	for _, j := range strList {
	//		//sendHis = map[int][]string{
	//		//	i: append(sendHis[i], j),
	//		//}
	//		sendHis[i] = append(sendHis[i], j)
	//	}
	//}
	//
	//fmt.Println(len(sendHis))
	//for k, v := range sendHis {
	//	fmt.Printf("k: %d, v: %s", k, v)
	//}

	ticker1 := time.NewTicker(10 * time.Second)
	ticker2 := time.NewTicker(time.Duration(10*3/2) * time.Second)

	for {
		for i := 0; i <= 10; i++ {
			select {
			case <-ticker1.C:
				fmt.Println("ticker1:", time.Now().Unix())
			case <-ticker2.C:
				fmt.Println("ticker2:", time.Now().Unix())
			}

			fmt.Printf("out: %v\n", time.Now().Unix())
		}
	}
}
