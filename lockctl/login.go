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
}

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