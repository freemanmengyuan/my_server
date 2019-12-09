package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

