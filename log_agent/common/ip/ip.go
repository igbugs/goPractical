package ip

import (
	"logging"
	"net"
)

var localIP string

func GetLocalIP() (localIP string, err error) {
	if len(localIP) > 0 {
		return localIP, nil
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logging.Error("get local ip failed, err: %v", err)
	}

	for _, addr := range addrs {
		ip, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ip.IP.IsLoopback() {
			continue
		}
		if !ip.IP.IsGlobalUnicast() {
			continue
		}

		logging.Debug("get local ip: %#v", ip.IP.String())
		localIP = ip.IP.String()
		return
	}
	return
}
