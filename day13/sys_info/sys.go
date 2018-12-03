package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func cpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err: %v\n", err)
		return
	}

	for _, ci := range cpuInfos {
		fmt.Printf("cpu: %#v\n", ci)
	}

	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent: %v\n", percent)
	}
}

func memInfo() {
	mInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get memory info failed, err: %v\n", err)
		return
	}
	fmt.Printf("mem info: %v\n", mInfo)
}

func hostInfo() {
	hInfo, err := host.Info()
	if err != nil {
		fmt.Printf("get host info failed, err: %v\n", err)
		return
	}
	fmt.Printf("host info: %v\n", hInfo)
}

func diskInfo() {
	dInfo, _ := disk.Usage("/")
	fmt.Printf("disk: %#v total: %#v free: %#v used: %#v\n", dInfo, dInfo.Total,
		dInfo.Free, dInfo.Used)

	part, _ := disk.Partitions(true)
	for _, v := range part {
		fmt.Printf("part: %#v", v)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func cpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("load: %v\n", info)
}

func netInfo() {
	info, _ := net.IOCounters(true)
	for i, v := range info {
		fmt.Printf("net: %d, %#v, recv: %v, send: %v\n", i, v, v.BytesRecv, v.BytesSent)
	}
}

func main() {
	memInfo()
	//cpuInfo()
	hostInfo()
	diskInfo()
	cpuLoad()
	netInfo()
}
