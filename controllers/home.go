package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	post_count, err := models.NewPostModel().Count("", "")
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)

	posts, err := models.NewPostModel().Offset("-PostId", paginator.Offset(), postsPerPage)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts

	latest_comments, err := models.NewCommentModel().Latest(10)
	if err == nil {
		this.Data["LatestComments"] = latest_comments
	}

	links, err := models.NewLinkModel().AllLinks()
	if err == nil {
		this.Data["Links"] = links
	}

	this.Data["PageTitle"] = blogTitle
	this.TplName = "home.tpl"
}
