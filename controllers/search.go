package controllers

import (
	"fmt"

	"github.com/migege/milog/models"
)

type SearchController struct {
	BaseController
}

func (this *SearchController) doSearch(term string) {
	this.TplName = "home.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("Search: %s - %s", term, blogTitle)
	this.Data["Content"] = fmt.Sprintf("Search - %s", term)

	if posts, err := models.NewSearchModel().Search(term); err == nil {
		this.Data["Posts"] = posts
		this.SetPostViews(posts)
	}

	this.LoadSidebar([]string{"LatestComments", "MostPopular"})
}

func (this *SearchController) Search() {
	term := this.Ctx.Input.Param(":term")
	this.doSearch(term)
}
