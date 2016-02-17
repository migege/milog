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

	posts, err := models.NewPostModel().Offset("", nil, "-PostId", paginator.Offset(), postsPerPage, false, true, true)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts

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
