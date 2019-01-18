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
	QUERY  = 1003
)

var (
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
	opHisChan = make(chan *OperationHis, 100)

	sendHis = make(map[int64][]*OperationHis)
	sid, _  = uuid.NewV4()
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

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
			Value: "record.csv",
			Usage: "history record to file",
		},
	}
	app.Action = action

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

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

	go func() {
		err = WriteFile(outputFile, opHisChan)
		if err != nil {
			logging.Error("write file failed, err: %v", err)
		}
	}()
	ticker := time.NewTicker(time.Duration(ctx.Int("interval")) * time.Second)
	ts := time.Now().UnixNano() / 1e6

	for {
		for _, cardNo := range cardList {
			select {
			case <-ticker.C:
				for _, lockNo := range lockList {
					logging.Debug("send password timestamp: %v", ts)
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
					opHisChan <- ret

					sendHis[ret.TimeStamp] = append(sendHis[ret.TimeStamp], ret)

					logging.Debug("set the sendHis: %v", sendHis)

				}
			}
			ts = ts + 3600*1000

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

				for _, op := range sendHis[int64(keys[0])] {
					pr := &PwdDeleteReq{
						LockNo: op.LockNo,
						PwdNo:  op.PwdNo,
						Extra:  "",
					}
					ret, err := PwdDelete(ctx, token, pr)
					if err != nil {
						logging.Error("PwdDelete call failed, err: %v", err)
					}

					logging.Debug("PwdDelete response data: %#v", ret)
					opHisChan <- ret

					delete(sendHis, int64(keys[0]))
					logging.Debug("delete after the sendHis: %v", sendHis)
				}
			}
		}
	}
}
