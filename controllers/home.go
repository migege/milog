package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	posts_per_page := 10
	post_count, err := models.NewPostModel().Count()
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, posts_per_page, post_count)

	posts, err := models.NewPostModel().Offset("-PostId", paginator.Offset(), posts_per_page)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts

	latest_comments, err := models.NewCommentModel().Latest(10)
	if err == nil {
		this.Data["LatestComments"] = latest_comments
	}

	this.Data["PageTitle"] = blogTitle
	this.TplName = "home.tpl"
}
