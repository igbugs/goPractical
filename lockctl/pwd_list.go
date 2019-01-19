package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"logging"
)

type PwdLsReq struct {
	LockNo string `json:"lock_no"`
	PwdNo  int    `json:"pwd_no"`
}

type PwdLsResp struct {
	Data []*PwdLsRespData `json:"data"`
	RequestStatus
}

type PwdLsRespData struct {
	LockNo         string `json:"lock_no"`
	PwdNo          int    `json:"pwd_no"`
	Status         string `json:"status"`
	PwdText        string `json:"pwd_text"`
	ValidTimeStart int64  `json:"valid_time_start"`
	ValidTimeEnd   int64  `json:"valid_time_end"`
	PwdUserName    string `json:"pwd_user_name"`
	PwdUserMobile  string `json:"pwd_user_mobile"`
	PwdUserIdcard  string `json:"pwd_user_idcard"`
}

func PwdList(ctx *cli.Context, token string, pr *PwdLsReq) (rlt *PwdLsResp, err error) {
	logging.Debug("PwdList input parameter(pr): %#v", pr)
	body, err := Post(ctx, token, "/pwd/list", pr)
	rlt = &PwdLsResp{}
	err = json.Unmarshal(body, rlt)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return nil, err
	}

	return rlt, nil
}
