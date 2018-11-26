package main

import (
	"gin-blog/models"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	select {}
	//下面的作用于 select{} 的作用相似, 目的都是为了阻塞程序的运行
	//t := time.NewTimer(time.Second * 10)
	////for {
	//	select {
	//	case <-t.C:
	//		// 每10s 就重置下计时器，否则只执行一次
	//		//t.Reset(time.Second * 10)
	//		log.Println("haha")
	//	}
	////}
}
