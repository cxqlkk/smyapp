package main

import (
	"go.uber.org/zap"
	"net/http"
	"smy/config"
	"smy/controller"
)
var Logger *zap.Logger

func main(){
	controller.DefaultWebRouter.Use(controller.SessionWare)
	controller.Manual() //add routes after midWare
	fs := http.FileServer(http.Dir(config.UpLoadPath))
	controller.DefaultWebRouter.AddRouter("/files/", http.StripPrefix("/files", fs))//fileServer
	controller.DefaultWebRouter.Run("127.0.0.1:8090") //127.0.0.1 nginx
	//session.DefalutSessionManger.SessionGc()
}
