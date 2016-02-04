package controllers

import (
	_ "fmt"
)

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["Content"] = "404 Not Found"
	this.TplName = "404.tpl"
}
