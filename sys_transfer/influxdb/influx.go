package influxdb

import (
	"encoding/json"
	"github.com/influxdata/influxdb/client/v2"
	"log_agent/collect_sys_info"
	"logging"
	"time"
)

const (
	DB             = "sys_info"
	username       = "admin"
	password       = ""
	CPUMeasurement = "cpu_usage"
)

var (
	influxClient client.Client
	msgChan      chan string
)

func Init(addr string, qs, thread int) (err error) {
	influxClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	if err != nil {
		logging.Error("init influxdb failed, err:%v", err)
		return
	}

	msgChan = make(chan string, qs)
	for i := 0; i < thread; i++ {
		go insertDB()
	}
	return
}

func AppendMsg(data string) {
	msgChan <- data
}

func insertDB() {
	for data := range msgChan {
		var sysInfo = &collect_sys_info.SystemInfo{}
		err := json.Unmarshal([]byte(data), sysInfo)
		if err != nil {
			logging.Error("kafka msg data unmasrshal failed, err:%v", err)
			continue
		}

		parseType(sysInfo)
	}
}

func parseType(sysInfo *collect_sys_info.SystemInfo) {
	switch sysInfo.Type {
	case "cpu":
		parseCPU(sysInfo)
	}
}

func parseCPU(sysInfo *collect_sys_info.SystemInfo) {
	var cpuInfo = &collect_sys_info.CpuInfo{}
	err := json.Unmarshal([]byte(sysInfo.Data), cpuInfo)
	if err != nil {
		logging.Error("sysInfo.Data json unmarshal failed, err:%v", err)
		return
	}

	logging.Debug("influxdb.parseCPU get cpu info succ, cpuInfo:%v", cpuInfo)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB,
		Precision: "s",
	})
	if err != nil {
		logging.Error("new batch points failed, err:%v", err)
		return
	}

	tags := map[string]string{
		"host": sysInfo.IP,
	}
	fields := map[string]interface{}{
		"usage":  cpuInfo.Percent,
		"load1":  cpuInfo.CpuLoad.Load1,
		"load5":  cpuInfo.CpuLoad.Load5,
		"load15": cpuInfo.CpuLoad.Load15,
		//"cputime": cpuInfo.CoreTimeStat,
	}

	pt, err := client.NewPoint(
		CPUMeasurement,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		logging.Error("new points failed, err:%v", err)
		return
	}

	bp.AddPoint(pt)
	if err = influxClient.Write(bp); err != nil {
		logging.Error("insert data  to influxdb failed, err:%v", err)
		return
	}

	logging.Debug("insert data to influxdb succ")
}
