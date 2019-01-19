package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"logging"
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
	Data PwdDeleteRespData `json:"data"`
	RequestStatus
}

func PwdDelete(ctx *cli.Context, token string, pr *PwdDeleteReq) (rlt *PwdDeleteResp, err error) {
	logging.Debug("PwdDelete input parameter(pr): %#v", pr)
	body, err := Post(ctx, token, "/pwd/delete", pr)
	rlt = &PwdDeleteResp{}
	err = json.Unmarshal(body, rlt)
	if err != nil {
		logging.Error("unmarshal err: %v", err)
		return nil, err
	}

	return rlt, nil
}
