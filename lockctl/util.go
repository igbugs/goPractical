package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"logging"
	"os"
	"strconv"
	"strings"
)

type OperationHis struct {
	LockNo     string `json:"lock_no"`
	CardNo     string `json:"card_id"`
	PwdNo      int    `json:"pwd_no"`
	Type       int    `json:"type"`
	Result     string `json:"result"`
	TimeStamp  int64  `json:"time_stamp"`
	ReturnBody string `json:"return_body"`
}

func ReadFile(path string) (result []string) {
	f, err := os.Open(path)
	if err != nil {
		logging.Error("ReadFile %s err: %v", path, err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if !strings.HasPrefix(line, "#") {
			line = strings.Replace(line, "\n", "", -1)
			result = append(result, line)
		}
	}
	return
}

func WriteFile(path string, op chan *OperationHis) (err error) {
	var header []string
	_, err = os.Stat(path)
	notExist := os.IsNotExist(err)
	if notExist {
		header = []string{"门锁编号", "身份证号", "密码编号", "操作类型", "操作结果", "时间", "失败原因"}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Error("openfile failed, path: %s, err: %v", path, err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if notExist {
		w.Write(header)
	}
	for c := range op {
		err := w.Write([]string{c.LockNo, c.CardNo, strconv.Itoa(c.PwdNo), strconv.Itoa(c.Type),
			c.Result, strconv.FormatInt(c.TimeStamp, 10), c.ReturnBody})
		if err != nil {
			logging.Error("wtite file failed, err: %v", err)
			return err
		}
		w.Flush()
	}

	return
}
