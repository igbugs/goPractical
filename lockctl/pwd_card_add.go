package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"logging"
)

type CardAddReq struct {
	LockNo         string `json:"lock_no"`
	CardType       int    `json:"card_type"`
	CardNo         string `json:"card_no"`
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
	Data CardAddRespData `json:"data"`
	RequestStatus
}

func CardAdd(ctx *cli.Context, token string, pr *CardAddReq) (rlt *CardAddResp, err error) {
	logging.Debug("CardAdd input parameter(pr): %#v", pr)
	body, err := Post(ctx, token, "/pwd/card/add", pr)
	rlt = &CardAddResp{}
	err = json.Unmarshal(body, rlt)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return nil, err
	}

	return rlt, nil
}
