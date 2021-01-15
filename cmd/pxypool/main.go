package main

import (
	"github.com/tangx/pxypool/pkg/browser"
	"github.com/tangx/pxypool/pkg/checker"
	"github.com/tangx/pxypool/pkg/crawler/kuaidaili"
	"github.com/tangx/pxypool/pkg/crawler/xiladaili"
	"github.com/tangx/pxypool/pkg/keeper"
)

func main() {
	if err := browser.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetReportCaller(true)
	go checker.Initial()
	go keeper.Initial()
	go kuaidaili.Initial()
	go xiladaili.Initial()
}
