package main

import (
	"fmt"
	"net"
	"logging"
)

func GetLocalIP() (ip string, err error) {
	var localIP string
		if len(localIP) > 0 {
			ip = localIP
			return
		}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logging.Error("get local ip failed, err: %v", err)
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		logging.Debug("get local ip: %#v\n", ipAddr.IP.String())
		localIP = ipAddr.IP.String()
		ip = localIP
		return
	}
	return
}


//var (
//	localIP string
//)

//func GetLocalIP() (ip string, err error) {
//	//if len(localIP) > 0 {
//	//	ip = localIP
//	//	return
//	//}
//
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return
//	}
//
//	for _, addr := range addrs {
//		ipAddr, ok := addr.(*net.IPNet)
//		if !ok {
//			continue
//		}
//
//		if ipAddr.IP.IsLoopback() {
//			continue
//		}
//
//		if !ipAddr.IP.IsGlobalUnicast() {
//			continue
//		}
//
//		logging.Debug("get local ip:%#v\n", ipAddr.IP.String())
//		localIP = ipAddr.IP.String()
//		ip = localIP
//		return
//	}
//	return
//}

func main()  {
	localIP, err := GetLocalIP()
	if err != nil {
		fmt.Printf("getlocalip failed, err:%v\n", err)
	}
	fmt.Printf("localIP: %v\n", localIP)

}