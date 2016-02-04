package main

import (
	"github.com/astaxie/beego"
	"github.com/migege/milog/controllers"
	_ "github.com/migege/milog/routers"
)

func init() {
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
