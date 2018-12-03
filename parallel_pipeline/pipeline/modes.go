package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		//a := []int{}
		a := make([]int, 0)
		for v := range in {
			a = append(a, v)
		}
		// read 一个 chunkSize 后的时间
		fmt.Println("Read done: ", time.Now().Sub(startTime))

		sort.Ints(a)
		fmt.Println("InMemSort done: ", time.Now().Sub(startTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func ReaderSource(r io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		byteRead := 0
		for {
			n, err := r.Read(buffer)
			byteRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			// 当 存在错误的时候，break;
			// 或者 chunkSize 不等于 -1 (即，chunkSize 有大小限制)，
			// 读取的byteRead 字节数大于 chunkSize 的大小，则break
			if err != nil || (chunkSize != -1 && byteRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(w io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		w.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	go func() {
		for i := 0; i < count; i++ {
			out <- r.Int()
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		// 判断两者之一有数据的话，进入循环
		for ok1 || ok2 {
			// 判断 in2 是否有数据，没有的话，或者 v1 <= v2 的话 v1 的值输入out channel
			if !ok2 || (ok1 && v1 <= v2) {
				// 去除v1 到out 的merge channel 中
				out <- v1
				// 再次去除下一个的v1 ,重新赋值
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done: ", time.Now().Sub(startTime))
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	// 不断的递归调用，直至inputs 的长度为 1
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
