package main

import (
	"net/http"
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"os"
	"os/signal"
	"golang.org/x/net/context"
	"time"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/gredis"
)

func main() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
