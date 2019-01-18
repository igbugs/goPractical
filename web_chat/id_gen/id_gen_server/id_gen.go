package main

import (
	"github.com/gin-gonic/gin/json"
	"github.com/sony/sonyflake"
	"logging"
	"net/http"
)

const (
	UNKNOW       = -1
	SUCCESS      = 0
	ERR_INTERNAL = 1001
)

var (
	snowflake *sonyflake.Sonyflake
)

var MsgFlags = map[int]string{
	UNKNOW:       "failed",
	SUCCESS:      "success",
	ERR_INTERNAL: "err_internal",
}

type RespData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	GenId   uint64 `json:"gen_id"`
}

func getMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[UNKNOW]
}

func respMsg(id uint64, code int, w http.ResponseWriter) {
	var resp = RespData{
		Code:    code,
		Message: getMsg(code),
		GenId:   id,
	}
	data, _ := json.Marshal(&resp)
	_, err := w.Write(data)
	if err != nil {
		logging.Error("respMsg write failed, err:%v", err)
		return
	}
}

func initSonyFlake() (err error) {
	settings := sonyflake.Settings{}
	settings.MachineID = func() (u uint16, e error) {
		return 1, nil
	}
	snowflake = sonyflake.NewSonyflake(settings)
	return
}

func idGen(w http.ResponseWriter, r *http.Request) {
	var err error
	var id uint64
	defer func() {
		logging.Info("id gen err info:%v, id:%v, caller ip:%v, url:%v",
			err, id,r.RemoteAddr, r.RequestURI)
	}()

	id, err = snowflake.NextID()
	if err != nil {
		respMsg(0, ERR_INTERNAL, w)
		return
	}
	respMsg(id, SUCCESS, w)
}

func main() {
	err := initSonyFlake()
	if err != nil {
		logging.Error("init snow flake failed, err:%v", err)
		return
	}

	http.HandleFunc("/id/gen", idGen)
	logging.Info("id gen server starting...")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		logging.Error("id gen server err:%v", err)
		return
	}
}
