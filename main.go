package main

import (
	"fmt"
	log "gin-blog/pkg/logging"
	"gin-blog/routers"
	"syscall"

	"github.com/fvbock/endless"
	"gin-blog/pkg/setting"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Errorf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Errorf("Server err: %v", err)
	}
}