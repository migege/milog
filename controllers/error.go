package controllers

import (
	_ "fmt"
	"github.com/migege/milog/models"
)

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	posts := models.NewPostModel().All("-PostId")
	this.Data["Posts"] = posts
	this.Data["PageTitle"] = blogTitle
	this.Data["Content"] = "404 出错啦，不如看看其它内容吧。"
	this.TplName = "404.tpl"
}
