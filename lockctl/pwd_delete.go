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
	//RltCode string            `json:"rlt_code"`
	//RltMsg  string            `json:"rlt_msg"`
}

//func PwdDelete(ctx *cli.Context, token string, pr *PwdDeleteReq) (rlt *PwdDeleteResp, err error) {
//	logging.Debug("PwdDelete input parameter(pr): %#v", pr)
//	data, _ := json.Marshal(pr)
//	req, err := http.NewRequest("POST",
//		"http://"+ctx.String("host")+"/pwd/delete",
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
//	//fmt.Printf(string(body))
//
//	rlt = &PwdDeleteResp{}
//	err = json.Unmarshal(body, rlt)
//	if err != nil {
//		logging.Error("unmarshal err: %v", err)
//		return nil, err
//	}
//
//	return rlt, nil
//}

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