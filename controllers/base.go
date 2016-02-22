package controllers

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/migege/milog/models"
)

var (
	blogTitle    string
	blogDesc     string
	blogUrl      string
	postsPerPage int
)

type BaseController struct {
	beego.Controller
	loggedUser *models.Author
	startTime  time.Time
}

func init() {
}

func (this *BaseController) Prepare() {
	this.startTime = time.Now()
	ts := strconv.FormatInt(time.Now().UnixNano(), 10)
	this.Data["TimeStamp"] = ts

	options := models.NewOptionModel().Names(&[]string{
		models.OPTION_BLOG_TITLE,
		models.OPTION_BLOG_DESC,
		models.OPTION_BLOG_URL,
		models.OPTION_POSTS_PER_PAGE,
	})
	blogTitle = options[models.OPTION_BLOG_TITLE].OptionValue
	blogDesc = options[models.OPTION_BLOG_DESC].OptionValue
	blogUrl = options[models.OPTION_BLOG_URL].OptionValue
	if postsPerPage, _ = options[models.OPTION_POSTS_PER_PAGE].GetInt(); postsPerPage < 1 {
		postsPerPage = 10
	}

	this.Data["BlogTitle"] = blogTitle
	this.Data["BlogDesc"] = blogDesc
	this.Data["BlogUrl"] = blogUrl

	logged_username := this.GetSession(SESS_NAME)
	if logged_username == nil {
		logged_username, res := this.GetSecureCookie(COOKIE_SECURE_KEY_USER, COOKIE_NAME_LOGGED_USER)
		if res == true {
			if logged_user, err := models.NewAuthorModel().ByName(fmt.Sprintf("%s", logged_username)); err == nil {
				this.SetSession(SESS_NAME, logged_username)
				this.Data["LoggedUser"] = logged_user
				this.loggedUser = logged_user
			}
		}
	} else {
		if logged_user, err := models.NewAuthorModel().ByName(fmt.Sprintf("%s", logged_username)); err == nil {
			this.Data["LoggedUser"] = logged_user
			this.loggedUser = logged_user
		}
	}

	this.Data["GoVersion"] = runtime.Version()
}

func (this *BaseController) Render() error {
	this.Data["Duration"] = time.Since(this.startTime).String()
	return this.Controller.Render()
}

func (this *BaseController) CheckLogged() {
	if this.loggedUser == nil {
		this.Redirect("/login", 302)
		this.StopRun()
	}
}

func (this *BaseController) LoadSidebar(widgets []string) {
	for _, widget := range widgets {
		if widget == "LatestComments" {
			if latest_comments, err := models.NewCommentModel().Latest(10); err == nil {
				this.Data["LatestComments"] = latest_comments
			}
		} else if widget == "MostPopular" {
			if posts, err := models.NewPostModel().MostPopular(10); err == nil {
				this.Data["MostPopular"] = posts
			}
		} else if widget == "Links" {
			if links, err := models.NewLinkModel().AllLinks(); err == nil {
				this.Data["Links"] = links
			}
		} else {
		}
	}
}
