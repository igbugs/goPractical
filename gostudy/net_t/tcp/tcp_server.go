package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}

		rAddr := conn.RemoteAddr()
		fmt.Println("Recevie from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client!"))
		if err2 != nil {
			return
		}
	}
}

func main() {
	service := ":5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp_chat", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp_chat", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		//conn.Close()		// handleClient 执行完成后，关闭连接
	}
}
