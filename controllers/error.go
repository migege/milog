package controllers

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	post_count, err := models.NewPostModel().Count("", "")
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)

	posts, err := models.NewPostModel().Offset("", nil, "-PostId", paginator.Offset(), postsPerPage, false, true, true)
	if err != nil {
		panic(err)
	}

	views := make(map[int]int)
	for _, post := range posts {
		for _, view := range post.PostViews {
			if view.ViewedBy == "human" {
				views[post.PostId] = view.Views
				break
			}
		}
	}
	this.Data["Views"] = views

	this.LoadSidebar([]string{"LatestComments", "MostPopular"})

	this.Data["Posts"] = posts
	this.Data["PageTitle"] = blogTitle
	this.Data["Content"] = "404 - 您要找的页面不见啦，不如看看其它内容吧"
	this.TplName = "404.tpl"
}
