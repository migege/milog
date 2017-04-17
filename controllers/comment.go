package controllers

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/migege/milog/models"
	"github.com/migege/milog/plugins"
)

var (
	cpt *captcha.Captcha
)

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

type CommentController struct {
	BaseController
}

func (this *CommentController) DoAddComment() {
	if cpt.VerifyReq(this.Ctx.Request) != true {
		panic(errors.New("error: human test"))
	}

	content := this.GetString("comment")
	if content == "" {
		panic(errors.New("error: empty comment"))
	}

	if err := plugins.Hooks.Callback("PreComment", content); err != nil {
		panic(err)
	}

	post_id := this.Input().Get("post_id")
	int_post_id, _ := strconv.Atoi(post_id)
	if int_post_id == 0 {
		this.StopRun()
	}

	comment_parent_id := this.Input().Get("comment_parent_id")
	int_comment_parent_id, _ := strconv.Atoi(comment_parent_id)

	comment := models.NewComment()
	logged_user := this.loggedUser
	if logged_user != nil {
		comment.CommentAuthor = logged_user.DisplayName
		comment.CommentAuthorMail = logged_user.AuthorMail
		comment.CommentAuthorUrl = logged_user.AuthorUrl
	} else {
		if this.GetString("comment_author") != "" {
			comment.CommentAuthor = this.GetString("comment_author")
		}
		if this.GetString("comment_author_mail") != "" {
			comment.CommentAuthorMail = this.GetString("comment_author_mail")
		}
		if this.GetString("comment_author_url") != "" {
			comment.CommentAuthorUrl = this.GetString("comment_author_url")
		}
	}
	comment.CommentAuthorIp = this.Ctx.Input.IP()
	comment.CommentContent = content
	comment.CommentAgent = this.Ctx.Request.UserAgent()
	comment.CommentParentId = int_comment_parent_id

	post := models.NewPost()
	post.PostId = int_post_id
	comment.Post = post

	comment_id, err := models.NewCommentModel().AddComment(comment)
	if err != nil {
		panic(err)
	}

	if logged_user == nil {
		this.Ctx.SetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR, comment.CommentAuthor)
		this.Ctx.SetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR_MAIL, comment.CommentAuthorMail)
		this.Ctx.SetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR_URL, comment.CommentAuthorUrl)
	}

	ret := struct {
		Code    int
		Message string
		Data    int64
	}{0, "success", comment_id}
	this.Data["json"] = &ret
	this.ServeJSON()
}
