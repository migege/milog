package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/migege/milog/models"
	"strconv"
	"time"
)

var (
	blogTitle string
	blogDesc  string
	blogUrl   string
)

type BaseController struct {
	beego.Controller
	loggedUser *models.Author
}

func init() {
}

func (this *BaseController) Prepare() {
	ts := strconv.FormatInt(time.Now().UnixNano(), 10)
	this.Data["TimeStamp"] = ts

	options := models.NewOptionModel().Names(&[]string{
		models.OPTION_BLOG_TITLE,
		models.OPTION_BLOG_DESC,
		models.OPTION_BLOG_URL,
	})
	blogTitle = options[models.OPTION_BLOG_TITLE].OptionValue
	blogDesc = options[models.OPTION_BLOG_DESC].OptionValue
	blogUrl = options[models.OPTION_BLOG_URL].OptionValue

	this.Data["BlogTitle"] = blogTitle
	this.Data["BlogDesc"] = blogDesc
	this.Data["BlogUrl"] = blogUrl

	logged_username := this.GetSession(SESS_NAME)
	if logged_username == nil {
		logged_username, res := this.GetSecureCookie(COOKIE_SECURE_KEY_USER, COOKIE_NAME_LOGGED_USER)
		if res == true {
			this.SetSession(SESS_NAME, logged_username)
			logged_user := models.NewAuthorModel().ByName(fmt.Sprintf("%s", logged_username))
			this.Data["LoggedUser"] = logged_user
			this.loggedUser = logged_user
		}
	} else {
		logged_user := models.NewAuthorModel().ByName(fmt.Sprintf("%s", logged_username))
		this.Data["LoggedUser"] = logged_user
		this.loggedUser = logged_user
	}
}

func (this *BaseController) CheckLogged() {
	if this.loggedUser == nil {
		this.Redirect("/login", 302)
		this.StopRun()
	}
}
