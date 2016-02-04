package main

import (
	"github.com/astaxie/beego"
	_ "github.com/migege/milog/routers"
)

func init() {
}

func main() {
	beego.Run()
}
