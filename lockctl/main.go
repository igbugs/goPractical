package main

import (
	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
	"logging"
	"net/http"
	"os"
	"sort"
	"time"
)

const (
	SEND   = 1001
	DELETE = 1002
	CHECK  = 1003
)

var (
	logType, level int
	sid, _         = uuid.NewV4()

	client = &http.Client{
		Timeout: 30 * time.Second,
	}
	opHisChan      = make(chan *CheckPwdStatus, 100)
	sendStatusChan = make(chan *OperationHis, 100)

	statusMsg = map[string]string{
		"01": "启用中",
		"03": "删除中",
		"11": "已启用",
		"13": "已删除",
		"21": "启用失败",
		"23": "删除失败",
		"":   "NotFound PwdNo",
	}

	sendHis = make(map[int64][]*OperationHis)
)

var logTypeMap = map[string]int{
	"console": logging.LogTypeConsole,
	"file":    logging.LogTypeFile,
	"net":     logging.LogTypeNet,
}

var levelMap = map[string]int{
	"debug": logging.LogLevelDebug,
	"trace": logging.LogLevelTrace,
	"info":  logging.LogLevelInfo,
	"warn":  logging.LogLevelWarn,
	"error": logging.LogLevelError,
	"fatal": logging.LogLevelFatal,
}

func main() {
	app := cli.NewApp()
	app.Version = "0.0.8"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host, H",
			Value: "bak.ops.huohetech.com",
			Usage: "ops server address",
		},
		cli.StringFlag{
			Name:  "username, u",
			Value: "gj_1291_4209",
			Usage: "ops account",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "851a5bb6bd12b668",
			Usage: "ops account password",
		},
		cli.StringFlag{
			Name:  "phone, P",
			Value: "13121651514",
			Usage: "customer phone number",
		},
		cli.StringFlag{
			Name:  "id-card-file",
			Value: "",
			Usage: "id card list from `FILE`",
		},
		cli.StringFlag{
			Name:  "lock-file",
			Value: "",
			Usage: "lock list from `FILE`",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 10,
			Usage: "send password interval",
		},
		cli.IntFlag{
			Name:  "save-pwd-number, spn",
			Value: 20,
			Usage: "save password number",
		},
		cli.IntFlag{
			Name:  "pwd-valid-time, pvt",
			Value: 300000,
			Usage: "password validtime(ms)",
		},
		cli.StringFlag{
			Name:  "outfile, o",
			Value: "opHistory.csv",
			Usage: "history record to file",
		},
		cli.StringFlag{
			Name:  "level, l",
			Value: "info",
			Usage: "record log level",
		},
		cli.StringFlag{
			Name:  "log-type, lt",
			Value: "console",
			Usage: "record log type",
		},
	}
	app.Action = action

	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		logging.Fatal("app.Run Fatal: %v", err)
	}
}

func action(ctx *cli.Context) {
	level = levelMap[ctx.String("level")]
	logType = logTypeMap[ctx.String("log-type")]
	err := logging.Init(logType, level, "./lockctl.log", "lockctl")
	if err != nil {
		return
	}
	logging.Debug("init logging success")

	logging.Info("request host: %s", ctx.String("host"))
	token, err := loginToken(ctx, &Account{
		Acc:    ctx.String("username"),
		Passwd: ctx.String("password"),
	})
	logging.Info("username: %s, passwd: %s", ctx.String("username"), ctx.String("password"))
	logging.Info("get token: %s", token)
	if err != nil {
		logging.Error("get token failed, err: %v", err)
		return
	}

	if ctx.String("id-card-file") == "" ||
		ctx.String("lock-file") == "" {
		logging.Error("id-card-file, lock-file don't empty")
		os.Exit(1)
	}

	var outputFile = ctx.String("outfile")
	var cardList = ReadFile(ctx.String("id-card-file"))
	var lockList = ReadFile(ctx.String("lock-file"))

	logging.Debug("card list: %v", cardList)
	logging.Debug("lock list: %v", lockList)

	// 此goroutine 用于文件的写入
	go func() {
		err = WriteFile(outputFile, opHisChan)
		if err != nil {
			logging.Error("write file failed, err: %v", err)
		}
	}()

	// 此goroutine 用于下发或删除等操作的检测, 下发的状态
	go func() {
		logging.Debug("check the send passwd call pwd/list ")
		for op := range sendStatusChan {
			logging.Info("check opHis channel data: %v", op)
			body := &PwdLsReq{
				LockNo: op.LockNo,
				PwdNo:  op.PwdNo,
			}
			flag := false

			count := 0
			for {
				time.Sleep(50 * time.Microsecond)
			LABEL:
				//if count > 60 {
				//	logging.Info("call PwdList func more then 60 times, not query result, op: %#v", op)
				//	break
				//}
				ret, err := PwdList(ctx, token, body)
				if err != nil {
					logging.Error("call PwdList func err: %v", err)
				}
				logging.Debug("PwdList return data: %#v", ret)
				logging.Debug("PwdList return ret.data length: %v", len(ret.Data))
				if len(ret.Data) == 1 {
					logging.Debug("PwdList return ret.data: %v", *ret.Data[0])
				}
				if ret.RltCode == "HH0000" && len(ret.Data) != 0 {
					dataList := ret.Data[0]
					status := &CheckPwdStatus{
						OpHis:              op,
						Check:              1,
						PassCheckStatus:    dataList.Status,
						PassCheckStatusMsg: statusMsg[dataList.Status],
						PwdUserName:        dataList.PwdUserName,
						PwdUserMobile:      dataList.PwdUserMobile,
						PwdUserIdcard:      dataList.PwdUserIdcard,
						ValidTimeStart:     dataList.ValidTimeStart,
						ValidTimeEnd:       dataList.ValidTimeEnd,
					}
					switch dataList.Status {
					case "01":
						if op.OpType == SEND {
							count = count + 1
							goto LABEL
						}
					case "03":
						if op.OpType == DELETE {
							if !flag {
								opHisChan <- status
								flag = true
							}
							count = count + 1
							goto LABEL
						}
					case "11":
						if op.OpType == SEND {
							opHisChan <- status
						}
					case "13":
						if op.OpType == DELETE {
							opHisChan <- status
						}
					case "21":
						if op.OpType == SEND {
							logging.Info("the pwdNo status(21), 启用失败, %#v", op)
							opHisChan <- status
							_, err := PwdDelete(ctx, token, &PwdDeleteReq{
								LockNo: op.LockNo,
								PwdNo:  op.PwdNo,
							})
							if err != nil {
								logging.Error("PwdDelete call failed, err: %v", err)
							}
						}
					case "23":
						if op.OpType == DELETE {
							logging.Info("the pwdNo status(23), 删除失败, %#v", op)
							opHisChan <- status
							_, err := PwdDelete(ctx, token, &PwdDeleteReq{
								LockNo: op.LockNo,
								PwdNo:  op.PwdNo,
							})
							if err != nil {
								logging.Error("PwdDelete call failed, err: %v", err)
							}
						}
					default:
						logging.Info("not found the pwdNo: %v", op)
					}
					break
				} else if len(ret.Data) == 0 {
					logging.Info("PwdList return data length equal 0")
					break
				} else {
					logging.Error("PwdList call err: %#v", err)
				}
				count = count + 1
			}
		}
	}()

	ts := time.Now().UnixNano() / 1e6

	for {
		for _, cardNo := range cardList {
			time.Sleep(time.Duration(ctx.Int("interval")) * time.Second)
			// 按指定的时间间隔,定时的下发密码
			// 取得一个身份证cardNo 一次同时发送给所有的门锁
			for _, lockNo := range lockList {
				logging.Debug("send password timestamp: %v", ts)
				body := &CardAddReq{
					LockNo:         lockNo,
					CardType:       2,
					CardNo:         cardNo,
					ValidTimeStart: ts,
					ValidTimeEnd:   ts + int64(ctx.Int("pwd-valid-time")),
					PwdUserMobile:  ctx.String("phone"),
					PwdUserName:    "test-send-pass",
					Description:    "",
					Extra:          "",
				}
				ret, err := CardAdd(ctx, token, body)
				if err != nil {
					logging.Error("send password failed, err: %v", err)
				}
				logging.Debug("response result: %#v", ret)

				var op = &OperationHis{
					LockNo:  ret.Data.LockNo,
					PwdText: cardNo,
					PwdNo:   ret.Data.PwdNo,
					OpType:  SEND,
					Result:  ret.RltCode,
					RltMsg:  ret.RltMsg,
					OpTime:  ts,
				}
				logging.Debug("send passwd operationHis: %#v", op)
				if ret.RltCode == "HH0000" {
					// 如果调用成功, 则发送到 sendStatusChan 等待进行检测是否启用
					logging.Debug("SEND operation send sendStatusChan check, %#v", op)
					sendStatusChan <- op

					logging.Debug("set the sendHis Map: %v", sendHis)
					sendHis[op.OpTime] = append(sendHis[op.OpTime], op)
				} else {
					// 如果调用失败这直接写入文件, 记录下发操作
					opHisChan <- &CheckPwdStatus{
						OpHis: op,
					}
				}
			}

			// 进行需要删除的密码长度的检测
			time.Sleep(2 * time.Second)
			logging.Debug("sendHis length: %d, save pass number length: %d", len(sendHis), ctx.Int("save-pwd-number"))
			if len(sendHis) > ctx.Int("save-pwd-number") {
				var keys []int
				for k := range sendHis {
					keys = append(keys, int(k))
				}
				sort.Ints(keys)
				logging.Debug("the oldest opHistory timestamp: %v", keys[0])
				logging.Debug("the oldest opHistory: %#v", sendHis[int64(keys[0])])
				logging.Debug("the all opHistory timestamp: %#v", keys)

				for _, del := range sendHis[int64(keys[0])] {
					pr := &PwdDeleteReq{
						LockNo: del.LockNo,
						PwdNo:  del.PwdNo,
						Extra:  "",
					}
					ret, err := PwdDelete(ctx, token, pr)
					if err != nil {
						logging.Error("PwdDelete call failed, err: %v", err)
					}

					var op = &OperationHis{
						LockNo: ret.Data.LockNo,
						PwdNo:  ret.Data.PwdNo,
						OpType: DELETE,
						Result: ret.RltCode,
						RltMsg: ret.RltMsg,
						OpTime: time.Now().UnixNano() / 1e6,
					}

					logging.Debug("PwdDelete response data: %#v", ret)
					if ret.RltCode == "HH0000" {
						// 如果调用成功, 则发送到 sendStatusChan 等待进行检测是否删除
						logging.Debug("DELETE operation send sendStatusChan check, %#v", op)
						sendStatusChan <- op
					} else {
						// 如果调用失败这直接写入文件, 记录删除的操作
						opHisChan <- &CheckPwdStatus{
							OpHis: op,
						}
					}
				}

				// 循环完 以时间戳为key 的所有的门锁的下发记录slice 调用删除接口后, 删除map 中的key
				delete(sendHis, int64(keys[0]))
				logging.Debug("delete after the sendHis list: %v", sendHis)
			}
			// 每次下发身份证,时间向前推移, 避免产生同一身份证,密码有效期重叠的问题
			ts = ts + int64(ctx.Int("pwd-valid-time")/len(cardList)+1000)
		}
	}
}
