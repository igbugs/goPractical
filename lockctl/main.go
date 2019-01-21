package main

import (
	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
	"logging"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

const (
	SEND   = 1001
	DELETE = 1002
	CHECK  = 1003
)

type sendHisDel struct {
	lock    *sync.RWMutex
	DelList map[int64][]*OperationHis
}

var (
	sid, _ = uuid.NewV4()

	client = &http.Client{
		Timeout: 10 * time.Second,
	}
	opHisChan      = make(chan *CheckPwdStatus, 100)
	sendStatusChan = make(chan *OperationHis, 100)
	signal         = make(chan struct{}, 1)

	statusMsg = map[string]string{
		"01": "启用中",
		"03": "删除中",
		"11": "已启用",
		"13": "已删除",
		"21": "启用失败",
		"23": "删除失败",
		"":   "NotFound PwdNo",
	}

	sendHis = sendHisDel{
		lock:    new(sync.RWMutex),
		DelList: make(map[int64][]*OperationHis),
	}
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.5"

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
	logging.Debug("request host: %s", ctx.String("host"))
	token, err := loginToken(ctx, &Account{
		Acc:    ctx.String("username"),
		Passwd: ctx.String("password"),
	})
	logging.Debug("username: %s, passwd: %s", ctx.String("username"), ctx.String("password"))
	logging.Debug("get token: %s", token)
	if err != nil {
		logging.Error("get token failed, err: %v", err)
		return
	}

	if ctx.String("id-card-file") == "" ||
		ctx.String("lock-file") == "" ||
		ctx.String("outfile") == "" {
		logging.Debug("id-card-file, lock-file and outfile don't empty")
	}

	var outputFile = ctx.String("outfile")
	var cardList = ReadFile(ctx.String("id-card-file"))
	var lockList = ReadFile(ctx.String("lock-file"))

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
		wg := &sync.WaitGroup{}
		for op := range sendStatusChan {
			wg.Add(1)
			go func(op *OperationHis) {
				defer wg.Done()
				body := &PwdLsReq{
					LockNo: op.LockNo,
					PwdNo:  op.PwdNo,
				}

				count := 0
				time.Sleep(5 * time.Second)
				for {
					if count > 60 {
						logging.Debug("call PwdList func more then 60 times, spend time 300s, not query result, op: %#v", op)
						break
					}
					ret, err := PwdList(ctx, token, body)
					if err != nil {
						logging.Error("call PwdList func err: %v", err)
					}
					if ret.RltCode == "HH0000" {
						for _, dataList := range ret.Data {
							opHisChan <- &CheckPwdStatus{
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
						}
						break
					} else {
						logging.Error("PwdList response err: %#v", ret)
					}
					count = count + 1
				}
			}(op)
		}
		wg.Wait()
	}()

	wg := &sync.WaitGroup{}
	ticker1 := time.NewTicker(time.Duration(ctx.Int("interval")) * time.Second)
	ts := time.Now().UnixNano() / 1e6

	for {
		for _, cardNo := range cardList {
			select {
			case <-ticker1.C:
				// 按指定的时间间隔,定时的下发密码
				// 取得一个身份证cardNo 一次同时发送给所有的门锁
				for _, lockNo := range lockList {
					wg.Add(1)
					logging.Debug("send password timestamp: %v", ts)
					go func(lockNo string) {
						defer wg.Done()
						body := &CardAddReq{
							LockNo:         lockNo,
							CardType:       2,
							CardNo:         cardNo,
							ValidTimeStart: ts,
							ValidTimeEnd:   ts + int64(ctx.Int("pwd-valid-time")),
							PwdUserMobile:  ctx.String("phone"),
							PwdUserName:    "test-send-pass-xyb",
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
							sendHis.lock.Lock()
							{
								// 由于此处操作的为一个全局的map, 所有goroutine 存在竞争,需要加锁处理
								sendHis.DelList[op.OpTime] = append(sendHis.DelList[op.OpTime], op)
							}
							sendHis.lock.Unlock()
						} else {
							// 如果调用失败这直接写入文件, 记录下发操作
							opHisChan <- &CheckPwdStatus{
								OpHis: op,
							}
						}
					}(lockNo)
				}
				// 等待所有的goroutine 执行完成
				wg.Wait()

				// 此时可以发送 信号, 进行map 的长度检测, 决定是否执行密码的删除操作
				logging.Debug("send signal to delete worker")
				signal <- struct{}{}
			case <-signal:
				logging.Debug("sendHis length: %d, save pass number length: %d", len(sendHis.DelList), ctx.Int("save-pwd-number"))
				if len(sendHis.DelList) > ctx.Int("save-pwd-number") {
					var keys []int
					for k := range sendHis.DelList {
						keys = append(keys, int(k))
					}
					sort.Ints(keys)
					logging.Debug("the oldest opHistory timestamp: %v", keys[0])
					logging.Debug("the oldest opHistory: %#v", sendHis.DelList[int64(keys[0])])
					logging.Debug("the all opHistory timestamp: %#v", keys)

					for _, del := range sendHis.DelList[int64(keys[0])] {
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
					delete(sendHis.DelList, int64(keys[0]))
					logging.Debug("delete after the sendHis list: %v", sendHis)
				}
			}

			// 每次下发身份证,时间向前推移, 避免产生同一身份证,密码有效期重叠的问题
			ts = ts + int64(ctx.Int("pwd-valid-time")/len(cardList))
		}
	}
}
