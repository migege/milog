package controllers

import (
	"fmt"

	"github.com/migege/milog/models"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Prepare() {
	this.BaseController.Prepare()
	this.CheckLogged()

	post_count, _ := models.NewPostModel().Count("", nil)
	this.Data["PostCount"] = post_count
}

func (this *AdminController) Get() {
	this.TplName = "admin.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("%s - Admin - %s", this.loggedUser.DisplayName, blogTitle)
}
