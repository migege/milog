package controllers

import (
	"fmt"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type TagController struct {
	BaseController
}

func (this *TagController) ByName() {
	this.TplName = "home.tpl"
	tag_name := this.Ctx.Input.Param(":tag")

	post_count, err := models.NewPostModel().Count("Tags__Tag__TagSlug", tag_name)
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)

	posts, err := models.NewPostModel().ByTagName(tag_name, "-PostId", paginator.Offset(), postsPerPage, false, true, true)
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

	if latest_comments, err := models.NewCommentModel().Latest(10); err == nil {
		this.Data["LatestComments"] = latest_comments
	}

	if posts, err := models.NewPostModel().MostPopular(10); err == nil {
		this.Data["MostPopular"] = posts
	}

	if tag, err := models.NewTagModel().BySlug(tag_name); err == nil {
		this.Data["Content"] = fmt.Sprintf("Tag - %s", tag.TagName)
		this.Data["PageTitle"] = fmt.Sprintf("%s - Tag - %s", tag.TagName, blogTitle)
	}
}
