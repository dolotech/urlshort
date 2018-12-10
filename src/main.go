package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"logic"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9999", "Configuration port")
	flag.Parse()
	glog.Errorln("port is", port)

	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/urlshort", logic.Urlshort)

	//glog.Info(logic.Short("http://testlogin.al2il.com/k/jumpShare/?Game=niuniu_roomId=14129&platform=nationalHall"))

	e.Start(":"+port)
}
