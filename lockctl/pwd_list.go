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
	//RltCode string           `json:"rlt_code"`
	//RltMsg  string           `json:"rlt_msg"`
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

//func PwdList(ctx *cli.Context, token string, pr *PwdLsReq) (rlt *PwdLsResp, err error) {
//	logging.Debug("PwdList input parameter(pr): %#v", pr)
//	data, _ := json.Marshal(pr)
//	req, err := http.NewRequest("POST",
//		"http://"+ctx.String("host")+"/pwd/list",
//		bytes.NewReader(data))
//	if err != nil {
//		logging.Error("NewRequest err: %v", err)
//		return nil, err
//	}
//	req.Header.Set("version", "1.1")
//	req.Header.Set("s_id", sid.String())
//	req.Header.Set("access_token", token)
//	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		logging.Error("client.Do err: %v", err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logging.Error("read resp.Body err: %v", err)
//		return nil, err
//	}
//	logging.Debug("get PwdList response body: %#v", string(body))
//
//	rlt = &PwdLsResp{}
//	err = json.Unmarshal(body, rlt)
//	if err != nil {
//		logging.Error("unmarshal err: %v", err)
//		return nil, err
//	}
//
//	return rlt, nil
//}

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
