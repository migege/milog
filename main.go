package main

import (
	"github.com/astaxie/beego"
	_ "milog/routers"
)

func init() {
}

func main() {
	beego.Run()
}
