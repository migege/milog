package controllers

import (
	"fmt"
	"github.com/migege/milog/models"
)

type OptionController struct {
	AdminController
}

func (this *OptionController) setOptionsData(names *[]string) {
	options_map := models.NewOptionModel().Names(names)
	var options []*models.Option
	for _, option := range options_map {
		options = append(options, option)
	}
	this.Data["Options"] = options
}

func (this *OptionController) Basic() {
	this.TplName = "option.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("Basic Options - Admin - %s", blogTitle)

	names := []string{
		models.OPTION_BLOG_TITLE,
		models.OPTION_BLOG_DESC,
		models.OPTION_BLOG_URL,
		models.OPTION_COMMENT_DEFAULT_STATUS,
	}
	this.setOptionsData(&names)
}

func (this *OptionController) DoEdit() {
	err := models.NewOptionModel().Save(this.Input())
	if err != nil {
		panic(err)
	}

	this.Redirect("/admin/options", 302)
}
