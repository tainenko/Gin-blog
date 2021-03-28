package main

import (
	"fmt"
	"github.com/gin-blog/models"
	"github.com/gin-blog/pkg/logging"
	"github.com/gin-blog/pkg/setting"
	"github.com/gin-blog/pkg/util"
	"github.com/gin-blog/routers"
	"log"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
	router := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
