package controllers

import (
	_ "fmt"
	"github.com/migege/milog/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	posts := models.NewPostModel().All("-PostId")
	this.Data["Posts"] = posts
	this.Data["PageTitle"] = blogTitle
	this.TplName = "home.tpl"
}
