package main

import (
	"fmt"
	"net/http"

	"github.com/233liyu/wechat-app/db"
	"github.com/233liyu/wechat-app/logger"
	"github.com/233liyu/wechat-app/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	logger.InitLogger()
	log := logger.GetLogger()

	log.Info("init success")

	h := &service.WeChatHandler{}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/api/on_msg", h.HandleMessage)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("http server start failed: %+v", err))
	}
	return
}
