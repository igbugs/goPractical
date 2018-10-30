package ip

import (
	"net"
	"logging"
)

func GetLocalIP() (localIP string, err error) {
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

		logging.Debug("get local ip: %#v\n", ip.IP.String())
		localIP = ip.IP.String()
		return
	}
	return
}