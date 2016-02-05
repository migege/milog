package controllers

import (
	_ "fmt"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type TagController struct {
	BaseController
}

func (this *TagController) ByName() {
	this.TplName = "home.tpl"
	tag_name := this.Ctx.Input.Param(":tag")

	post_count, err := models.NewPostModel().Count("Tags__Tag__TagName", tag_name)
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)

	posts, err := models.NewPostModel().ByTagName(tag_name, "-PostId", paginator.Offset(), postsPerPage)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts
}
