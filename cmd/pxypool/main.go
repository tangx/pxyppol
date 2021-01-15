package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/browser"
	"github.com/tangx/pxypool/pkg/checker"
	"github.com/tangx/pxypool/pkg/crawler/xiladaili"
	"github.com/tangx/pxypool/pkg/keeper"
)

func main() {
	r := gin.Default()
	r.GET("/ping")
	r.GET("list", browser.ListHandler)
	r.GET("/random", browser.RandomHandler)
	_ = r.Run()
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	go checker.Initial()
	go keeper.Initial()
	// go kuaidaili.Initial()
	go xiladaili.Initial()
}
