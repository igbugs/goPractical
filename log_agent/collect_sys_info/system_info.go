package collect_sys_info

import (
	"logging"
	"encoding/json"
	"log_agent/common/ip"
	"log_agent/kafka"
	"sync"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

var sysInfoTopic string

const (
	SystemTypeCpu  = "cpu"
	SystemTypeMem  = "mem"
	SystemTypeDisk = "disk"
	SystemTypeNet  = "net"
	SystemTypeHost = "host"
)

type SystemInfo struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type PartitionStat struct {
	PartStat  disk.PartitionStat `json:"part_stat"`
	PartUsage disk.UsageStat     `json:"part_usage"`
}

type DiskInfo struct {
	Partitions []PartitionStat                `json:"partitions"`
	DiskIO     map[string]disk.IOCountersStat `json:"disk_io"`
}

func collectDisk() {
	var diskInfo DiskInfo
	diskInfo.DiskIO = make(map[string]disk.IOCountersStat)

	part, _ := disk.Partitions(true)
	for _, p := range part {
		var ps PartitionStat
		ps.PartStat = p

		dInfo, _ := disk.Usage(p.Device)
		ps.PartUsage = *dInfo

		diskInfo.Partitions = append(diskInfo.Partitions, ps)
	}

	ioStat, err := disk.IOCounters()
	if err != nil {
		logging.Error("get disk io stat failed, err: %v", err)
	}

	for k, v := range ioStat {
		diskInfo.DiskIO[k] = v
	}
	//diskInfo.DiskIO = ioStat

	sendToKafka(SystemTypeDisk, diskInfo)
}


type CpuInfo struct {
	Percent      float64          `json:"percent"`
	CpuLoad      load.AvgStat     `json:"cpu_load"`
	InfoStat     []cpu.InfoStat  `json:"info_stat"`
	CoreTimeStat []cpu.TimesStat `json:"core_time_stat"`
}

func collectCpu() {
	var cpuInfo CpuInfo

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		logging.Error("collect cpu percent failed, err: %v", err)
		return
	}
	cpuInfo.Percent = percent[0]

	infos, err := cpu.Info()
	if err != nil {
		logging.Error("collect cpu info failed, err: %v", err)
		return
	}

	for _, v := range infos {
		cpuInfo.InfoStat = append(cpuInfo.InfoStat, v)
	}

	cpuLoad, err := load.Avg()
	if err != nil {
		logging.Error("collect cpu load failed, err: %v", err)
		//return
	}
	cpuInfo.CpuLoad = *cpuLoad

	coreTimes, err := cpu.Times(true)
	if err != nil {
		logging.Error("collect coreTimes failed, err: %v", err)
		return
	}
	for _, v := range coreTimes {
		cpuInfo.CoreTimeStat = append(cpuInfo.CoreTimeStat, v)
	}

	logging.Debug("collect cpu succ, cpuinfo: %#v", cpuInfo)
	sendToKafka(SystemTypeCpu, &cpuInfo)
}


type MemInfo struct {
	VMStat *mem.VirtualMemoryStat	`json:"vm_stat"`
	SwapStat *mem.SwapMemoryStat		`json:"swap_stat"`
}

func collectMem()  {
	var memInfo MemInfo

	vmInfo, err := mem.VirtualMemory()
	if err != nil {
		logging.Error("collect memory info failed, err: %v", err)
		return
	}

	memInfo.VMStat = vmInfo

	swapInfo, err := mem.SwapMemory()
	if err != nil {
		logging.Error("collect swap memory info failed, err: %v", err)
		return
	}
	memInfo.SwapStat = swapInfo

	sendToKafka(SystemTypeMem, &memInfo)
}

type NetInfo struct {
	NetInterfaces []net.IOCountersStat	`json:"net_interfaces"`
}

func collectNet()  {
	var netInfo NetInfo
	//netInfo.NetInterfaces = make(map[string]*net.IOCountersStat, 16)

	info, err := net.IOCounters(true)
	if err != nil {
		logging.Error("collect net interfaces data failed, err: %v", err)
		return
	}

	for _, v := range info {
		netInfo.NetInterfaces = append(netInfo.NetInterfaces, v)
	}

	sendToKafka(SystemTypeNet, &netInfo)
}


type HostInfo struct {
	HostInfo *host.InfoStat	`json:"host_info"`
}

func collectHostInfo()  {
	var hostInfo HostInfo

	info, err := host.Info()
	if err != nil {
		logging.Error("collect host info failed, err: %v", err)
		return
	}
	hostInfo.HostInfo = info

	sendToKafka(SystemTypeHost, &hostInfo)
}

func sendToKafka(sysType string, data interface{}) {
	originData, err := json.Marshal(data)
	if err != nil {
		logging.Error("marshal originData failed, err: %v", err)
		return
	}

	localIP, _ := ip.GetLocalIP()
	var sysInfo = SystemInfo{
		IP:   localIP,
		Type: sysType,
		Data: string(originData),
	}

	jsonData, err := json.Marshal(sysInfo)
	if err != nil {
		logging.Error("marshal systemInfo failed, err: %v", err)
		return
	}

	msg := &kafka.Message{
		Topic: sysInfoTopic,
		Data:  string(jsonData),
	}

	err = kafka.SendLog(msg)
	if err != nil {
		logging.Error("send to kafka failed, err: %v", err)
		return
	}
	jsonMsg, _ := json.Marshal(msg)
	logging.Debug("sendLog success, msg: %v", string(jsonMsg))
}

func collect() {
	logging.Debug("start collect system info")
	collectDisk()
	collectCpu()
	collectMem()
	collectNet()
	collectHostInfo()
}

func Run(wg *sync.WaitGroup, interval time.Duration, topic string) {
	sysInfoTopic = topic
	timer := time.NewTicker(interval)
	for {
		select {
		case <-timer.C:
			collect()
		}
	}

	wg.Done()
}
