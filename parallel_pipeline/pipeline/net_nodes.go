package pipeline

import (
	"bufio"
	"net"
)

func NetWorkerSink(addr string, in <-chan int) {
	// 启动了一个 tcp 的server 接收来源网络的数据请求，把数据发送到远端
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()

		// 写数据到远端的 tcp 连接
		WriteSink(writer, in)
	}()
}

func NetWorkerSource(addr string) <-chan int {
	// 去连接tcp 的server端， 读取数据
	out := make(chan int, 1024)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
