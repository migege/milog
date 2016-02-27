package controllers

import (
	"fmt"
	"strconv"

	"github.com/migege/milog/models"
	"github.com/mssola/user_agent"
)

type PostController struct {
	BaseController
}

func (this *PostController) setPost(post *models.Post) {
	this.TplName = "post.tpl"
	this.Data["Author"] = post.Author
	this.Data["Post"] = post
	this.Data["PageTitle"] = fmt.Sprintf("%s - %s", post.PostTitle, blogTitle)

	this.LoadSidebar([]string{"LatestComments", "MostPopular", "TagCloud"})

	// comments
	comments := models.NewCommentModel().ByPostId(post.PostId, "-CommentId")
	this.Data["Comments"] = comments

	comment_author, res := this.Ctx.GetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR)
	if res == true {
		this.Data["CommentAuthor"] = comment_author
	}
	comment_author_mail, res := this.Ctx.GetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR_MAIL)
	if res == true {
		this.Data["CommentAuthorMail"] = comment_author_mail
	}
	comment_author_url, res := this.Ctx.GetSecureCookie(COOKIE_SECURE_KEY_COMMENT, COOKIE_NAME_COMMENT_AUTHOR_URL)
	if res == true {
		this.Data["CommentAuthorUrl"] = comment_author_url
	}
}

func (this *PostController) BySlug() {
	post_slug := this.Ctx.Input.Param(":slug")
	post_model := models.NewPostModel()
	post, err := post_model.BySlug(post_slug)
	if err != nil {
		this.Abort("404")
	}
	this.setPost(post)

	ua := user_agent.New(this.Ctx.Request.UserAgent())
	if ua.Bot() {
		post_model.ViewedBy(post.PostId, "bot")
	} else {
		post_model.ViewedBy(post.PostId, "human")
	}
}

func (this *PostController) ById() {
	id_str := this.Ctx.Input.Param(":id")
	post_id, _ := strconv.Atoi(id_str)
	post, err := models.NewPostModel().ById(post_id)
	if err != nil {
		this.Abort("404")
	}
	this.Redirect(post.PostLink(), 301)
}
