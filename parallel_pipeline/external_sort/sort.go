package main

import (
	"bufio"
	"fmt"
	"os"
	"parallel_pipeline/pipeline"
	"strconv"
)

func main() {
	//p := createPipeline("small.in", 512, 4)
	//writeToFile(p, "small.out")
	//printFile("small.out")

	//p := createNetPipeline("small.in", 512, 4)
	//writeToFile(p, "small.out")
	//printFile("small.out")

	//p := createPipeline("large.in", 800000000, 80)
	//writeToFile(p, "large.out")
	//printFile("large.out")

	p := createNetPipeline("large.in", 800000000, 8)
	writeToFile(p, "large.out")
	printFile("large.out")
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	// 开始进行计时
	pipeline.Init()

	//sortResults := []<-chan int{}
	sortResults := make([]<-chan int, 0)
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		// whence: 0 从头开始读取文件
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func createNetPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	// 开始进行计时
	pipeline.Init()

	sortAddr := make([]string, 0)
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		// whence: 0 从头开始读取文件
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetWorkerSink(addr, pipeline.InMemSort(source))
		// 将每个 chunk 读取后，启动的tcp server 的地址收集起来
		sortAddr = append(sortAddr, addr)
	}

	sortResults := make([]<-chan int, 0)
	for _, addr := range sortAddr {
		// 连接每个的tcp server 将 远程读取的数据，发送到 out channel 后输出到 sortResults的channel 的切片 合集中去
		sortResults = append(sortResults, pipeline.NetWorkerSource(addr))
	}
	// 对 sortResults channel 进行两两merge， 汇总到 一个 统一的out channel 中
	return pipeline.MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	//for v := range p {
	//	fmt.Println(v)
	//}

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
