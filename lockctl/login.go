package main

import (
	"encoding/json"
	"errors"
	"github.com/urfave/cli"
	"logging"
)

type Account struct {
	Acc    string `json:"account"`
	Passwd string `json:"password"`
}

type LoginRespData struct {
	AccessToken   string `json:"access_token"`
	ExpiresSecond int    `json:"expires_second"`
}

type LoginResp struct {
	Data    LoginRespData `json:"data"`
	RequestStatus
	//RltCode string        `json:"rlt_code"`
	//RltMsg  string        `json:"rlt_msg"`
}

//func loginToken(ctx *cli.Context, acc *Account) (string, error) {
//	data, _ := json.Marshal(acc)
//	req, err := http.NewRequest("POST",
//		"http://"+ctx.String("host")+"/login",
//		bytes.NewReader(data))
//	if err != nil {
//		logging.Error("NewRequest err: %v", err)
//		return "", err
//	}
//	req.Header.Set("version", "1.1")
//	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		logging.Error("client.Do err: %v", err)
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logging.Error("read resp.Body err: %v", err)
//		return "", err
//	}
//	logging.Debug("login response body: %#v", string(body))
//
//	var loginResp LoginResp
//	err = json.Unmarshal(body, &loginResp)
//	if err != nil {
//		logging.Error("unmarshal err: %v", err)
//		return "", err
//	}
//
//	if loginResp.RltCode != "HH0000" {
//		logging.Error("resp data: %#v", loginResp)
//		return "", errors.New("login resp code isn't HH0000")
//	}
//	return loginResp.Data.AccessToken, nil
//}

func loginToken(ctx *cli.Context, acc *Account) (string, error) {
	body, err := Post(ctx, "", "/login", acc)

	loginResp := &LoginResp{}
	err = json.Unmarshal(body, loginResp)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return "", err
	}

	if loginResp.RltCode != "HH0000" {
		logging.Error("resp data: %#v", loginResp)
		return "", errors.New("login resp code isn't HH0000")
	}
	return loginResp.Data.AccessToken, nil
}