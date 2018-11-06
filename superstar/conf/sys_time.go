package conf

import "time"

const (
	SysTimeForm      string = "2006-01-02 15:04:05"
	SysTimeFromShort string = "2006-01-02"
)

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")
