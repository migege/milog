package main

import (
	"github.com/astaxie/beego"
	"github.com/migege/milog/controllers"
	_ "github.com/migege/milog/plugins"
	_ "github.com/migege/milog/routers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.AddFuncMap("add", Add)
}

func main() {
	beego.Run()
}
