package controllers

import (
	"fmt"

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

	posts, err := models.NewPostModel().Offset("", nil, "-PostId", paginator.Offset(), postsPerPage, false, true, true)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts

	this.SetPostViews(posts)
	this.LoadSidebar([]string{"LatestComments", "MostPopular", "Links", "TagCloud"})

	this.Data["PageTitle"] = fmt.Sprintf("%s %s", blogTitle, blogDesc)
	this.TplName = "home.tpl"
}
