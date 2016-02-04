package controllers

import (
	"fmt"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) Prepare() {
	this.BaseController.Prepare()
	this.CheckLogged()
}

func (this *AdminController) Get() {
	this.TplName = "admin.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("%s - Admin - %s", this.loggedUser.DisplayName, blogTitle)
}
