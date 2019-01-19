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

type CheckPwdStatus struct {
	OpHis              *OperationHis
	Check              int    `json:"check"`
	PassCheckStatus    string `json:"pass_check_status"`
	PassCheckStatusMsg string `json:"pass_check_status_msg"`

	PwdUserName    string `json:"pwd_user_name"`
	PwdUserIdcard  string `json:"pwd_user_idcard"`
	PwdUserMobile  string `json:"pwd_user_mobile"`
	ValidTimeStart int64  `json:"valid_time_start"`
	ValidTimeEnd   int64  `json:"valid_time_end"`
}

type OperationHis struct {
	LockNo  string `json:"lock_no"`
	PwdText string `json:"pwd_text"`
	PwdNo   int    `json:"pwd_no"`
	OpType  int    `json:"op_type"`
	Result  string `json:"result"`
	RltMsg  string `json:"rlt_msg"`
	OpTime  int64  `json:"op_time"`
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

func WriteFile(path string, op chan *CheckPwdStatus) (err error) {
	var header []string
	_, err = os.Stat(path)
	notExist := os.IsNotExist(err)
	if notExist {
		header = []string{"门锁编号", "身份证号", "加密后身份证号", "操作类型","密码编号",
			"操作结果", "返回信息", "操作时间戳", "是否检查操作状态", "下发密码状态码", "密码检测信息",
		"PwdUserName", "mobile", "生效时间", "过期时间"}
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
		err := w.Write([]string{c.OpHis.LockNo, c.PwdUserIdcard, c.OpHis.PwdText, strconv.Itoa(c.OpHis.OpType),
			strconv.Itoa(c.OpHis.PwdNo), c.OpHis.Result, c.OpHis.RltMsg, strconv.FormatInt(c.OpHis.OpTime, 10),
			strconv.Itoa(c.Check), c.PassCheckStatus, c.PassCheckStatusMsg, c.PwdUserName, c.PwdUserMobile,
			strconv.FormatInt(c.ValidTimeStart, 10), strconv.FormatInt(c.ValidTimeEnd, 10)})
		if err != nil {
			logging.Error("write file failed, err: %v", err)
			return err
		}
		w.Flush()
	}

	return
}
