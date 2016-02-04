package controllers

import (
	_ "fmt"
	"github.com/migege/milog/models"
)

type TagController struct {
	BaseController
}

func (this *TagController) ByName() {
	this.TplName = "home.tpl"
	tag_name := this.Ctx.Input.Param(":tag")
	posts := models.NewPostModel().ByTagName(tag_name, "-PostId")
	this.Data["Posts"] = posts
}
