package controllers

import (
	"fmt"
	"milog/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.TplName = "login.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("Log In - %s", blogTitle)
	this.Data["Refer"] = this.Ctx.Input.Refer()
}

func (this *LoginController) DoLogin() {
	ret := struct {
		Code    int
		Message string
	}{0, "success"}
	username := this.GetString("log")
	password := this.GetString("pwd")
	ts := this.GetString("ts")
	err := models.NewAuthorModel().Validate(ts, username, password)
	if err == nil {
		this.SetSession(SESS_NAME, username)
	} else {
		ret.Code = 1
		ret.Message = err.Error()
	}
	this.Data["json"] = &ret
	this.ServeJSON()
}

func (this *LoginController) Logout() {
	this.DelSession(SESS_NAME)
	ret := struct {
		Code    int
		Message string
	}{0, "success"}
	this.Data["json"] = &ret
	this.ServeJSON()
}
