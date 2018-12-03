package main

import (
	"configlib"
	"fmt"
)

type ServerConf struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
}

type DbConf struct {
	User     string  `ini:"user"`
	Password string  `ini:"password"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Database string  `ini:"database"`
	Rate     float32 `ini:"rate"`
}

type Config struct {
	Server ServerConf `ini:"server"`
	CartDb DbConf     `ini:"cartdb"`
}

func main() {
	var conf Config
	filename := "./example.ini"
	err := configlib.UnMarshalFile(filename, &conf)
	if err != nil {
		fmt.Printf("unmarshal file failed, err: %v\n", err)
		return
	}

	fmt.Printf("config: %#v\n", conf)

	err = configlib.MarshalFile("C:/GoProject/Go3Project/test.ini", conf)
	if err != nil {
		fmt.Printf("marshal file failed, err: %v\n", err)
		return
	}
}
