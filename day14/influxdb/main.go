package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/lunny/log"
	"time"
	"fmt"
	)

const (
	MyDB          = "sys_info"
	username      = "admin"
	password      = ""
	MyMeasurement = "cpu_usage"
)

func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.247.131:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func QueryDB(cli client.Client, cmd string) (result []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}

	if resp, err := cli.Query(q); err == nil {
		if resp.Error() != nil {
			return result, resp.Error()
		}
		result = resp.Results
	} else {
		return result, err
	}
	return result, nil
}

func writePoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   20.2,
		"system": 30.3,
		"user":   34.4,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn := connInflux()
	fmt.Println(conn)

	writePoints(conn)

	// 获取10条记录并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", MyMeasurement, 10)
	result, err := QueryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range result[0].Series[0].Values {
		for j, value := range row {
			//t, err := time.Parse(time.RFC3339, row[j].(string))
			//if err != nil {
			//	log.Fatal(err)
			//}
			////fmt.Println(reflect.TypeOf(row[1]))
			//valu := row[j].(json.Number)
			//log.Printf("[%2d] %s: %s\n", i, t.Format(time.Stamp), valu)
			log.Printf("i:%d j:%d value:%#v\n", i, j, value)
		}
	}
}
