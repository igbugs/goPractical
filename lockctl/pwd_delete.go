package main

import (
	"bytes"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
	"io/ioutil"
	"logging"
	"net/http"
	"time"
)

type PwdDeleteReq struct {
	LockNo string `json:"lock_no"`
	PwdNo  int    `json:"pwd_no"`
	Extra  string `json:"extra"`
}

type PwdDeleteRespData struct {
	LockNo     string `json:"lock_no"`
	PwdNo      int    `json:"pwd_no"`
	BusinessId string `json:"business_id"`
}

type PwdDeleteResp struct {
	Data    PwdDeleteRespData `json:"data"`
	RltCode string            `json:"rlt_code"`
	RltMsg  string            `json:"rlt_msg"`
}

func PwdDelete(ctx *cli.Context, token string, pr *PwdDeleteReq) (opHis *OperationHis, err error) {
	logging.Debug("PwdDelete input parameter(pr): %#v", pr)
	data, _ := json.Marshal(pr)
	req, err := http.NewRequest("POST",
		"http://"+ctx.String("host")+"/pwd/card/add",
		bytes.NewReader(data))
	if err != nil {
		logging.Error("NewRequest err: %v", err)
		return nil, err
	}
	sid, _ := uuid.NewV4()
	req.Header.Set("version", "1.1")
	req.Header.Set("s_id", sid.String())
	req.Header.Set("access_token", token)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		logging.Error("client.Do err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error("read resp.Body err: %v", err)
		return nil, err
	}
	//fmt.Printf(string(body))

	var pwdResp PwdDeleteResp
	err = json.Unmarshal(body, &pwdResp)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return nil, err
	}

	return &OperationHis{
		LockNo:     pr.LockNo,
		CardNo:     "",
		PwdNo:      pwdResp.Data.PwdNo,
		Type:       DELETE,
		Result:     pwdResp.RltCode,
		TimeStamp:  time.Now().UnixNano() / 1e6,
		ReturnBody: string(body),
	}, nil
}
