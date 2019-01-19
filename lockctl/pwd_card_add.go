package main

import (
	"bytes"
	"encoding/json"
	"github.com/urfave/cli"
	"io/ioutil"
	"logging"
	"net/http"
)

type CardAddReq struct {
	LockNo         string `json:"lock_no"`
	CardType       int    `json:"card_type"`
	CardNo        string `json:"card_no"`
	ValidTimeStart int64  `json:"valid_time_start"`
	ValidTimeEnd   int64  `json:"valid_time_end"`
	PwdUserMobile  string `json:"pwd_user_mobile"`
	PwdUserName    string `json:"pwd_user_name"`
	Description    string `json:"description"`
	Extra          string `json:"extra"`
}

type CardAddRespData struct {
	LockNo     string `json:"lock_no"`
	PwdNo      int    `json:"pwd_no"`
	PwdText    string `json:"pwd_text"`
	BusinessId string `json:"business_id"`
}

type CardAddResp struct {
	Data    CardAddRespData `json:"data"`
	RltCode string          `json:"rlt_code"`
	RltMsg  string          `json:"rlt_msg"`
}

func CardAdd(ctx *cli.Context, token string, pr *CardAddReq) (rlt *CardAddResp, err error) {
	logging.Debug("CardAdd input parameter(pr): %#v", pr)
	data, _ := json.Marshal(pr)
	req, err := http.NewRequest("POST",
		"http://"+ctx.String("host")+"/pwd/card/add",
		bytes.NewReader(data))
	if err != nil {
		logging.Error("NewRequest err: %v", err)
		return nil, err
	}

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

	rlt = &CardAddResp{}
	err = json.Unmarshal(body, rlt)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return nil, err
	}

	return rlt, nil
}
