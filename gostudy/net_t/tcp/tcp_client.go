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

func main() {
	var buf [512]byte
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s host:port", os.Args[0])
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp_chat", service)
	checkError(err)

	conn, err := net.DialTCP("tcp_chat", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello Server!"))
	checkError(err)

	n, err = conn.Read(buf[0:])
	checkError(err)

	fmt.Println("Reply from server", rAddr.String(), string(buf[0:n]))
	//conn.Close()

	os.Exit(0)
}
