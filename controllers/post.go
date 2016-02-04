package controllers

import (
	"errors"
	"fmt"
	"milog/models"
	"strconv"
	"strings"
	"time"
)

type PostController struct {
	BaseController
}

func (this *PostController) Get() {
	this.TplName = "post.tpl"
	id_str := this.Ctx.Input.Param(":id")
	post_id, _ := strconv.Atoi(id_str)

	post := models.NewPostModel().ById(post_id)

	this.Data["Author"] = post.Author
	this.Data["Post"] = post
	this.Data["PageTitle"] = fmt.Sprintf("%s - %s", post.PostTitle, blogTitle)

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

func (this *PostController) PostNew() {
	this.CheckLogged()
	this.TplName = "admin-post.tpl"
	tags := models.NewTagModel().AllTags()
	this.Data["AllTags"] = tags
	this.Data["PageTitle"] = fmt.Sprintf("New Post - Admin - %s", blogTitle)
}

func (this *PostController) DoPostNew() {
	this.CheckLogged()
	post_title := this.GetString("post-title")
	post_content := this.GetString("post-content")
	post_content_md := this.GetString("post-content-md")
	if post_title == "" || post_content == "" || post_content_md == "" {
		panic(errors.New("error: empty post title or content"))
	}
	comment_status, err := this.GetInt("comment-status", 1)
	if err != nil {
		panic(err)
	}
	post_tags := this.Input()["post-tags"]
	var tags []*models.Tag
	for _, t := range post_tags {
		t = strings.TrimSpace(t)
		tag := &models.Tag{
			TagName: t,
			TagSlug: strings.ToLower(t),
		}
		tags = append(tags, tag)
	}

	post := models.NewPost()
	post.Tags = tags
	post.CommentStatus = comment_status
	post.PostTitle = post_title
	post.PostContent = post_content
	post.PostContentMd = post_content_md
	post.Author = this.loggedUser
	post_id, err := models.NewPostModel().PostNew(post)
	if err != nil {
		panic(err)
	} else {
		this.Redirect(fmt.Sprintf("/post/%d", post_id), 302)
	}
}

func (this *PostController) PostEdit() {
	this.CheckLogged()
	this.TplName = "admin-post.tpl"
	id := this.Ctx.Input.Param(":id")
	post_id, _ := strconv.Atoi(id)

	post := models.NewPostModel().ById(post_id)
	if post.Author.AuthorId != this.loggedUser.AuthorId {
		panic(errors.New("error: can't edit another one's post"))
	}

	tags := models.NewTagModel().AllTags()
	this.Data["Post"] = post
	this.Data["AllTags"] = tags
	this.Data["PageTitle"] = fmt.Sprintf("Editing Post: %s - %s", post.PostTitle, blogTitle)
}

func (this *PostController) DoPostEdit() {
	this.CheckLogged()
	id := this.GetString("post-id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	post_title := this.GetString("post-title")
	post_content := this.GetString("post-content")
	post_content_md := this.GetString("post-content-md")
	if post_title == "" || post_content == "" || post_content_md == "" {
		panic(errors.New("error: empty post title or content"))
	}
	comment_status, err := this.GetInt("comment-status", 1)
	if err != nil {
		panic(err)
	}
	post_tags := this.Input()["post-tags"]
	var tags []*models.Tag
	for _, t := range post_tags {
		t = strings.TrimSpace(t)
		tag := &models.Tag{
			TagName: t,
			TagSlug: strings.ToLower(t),
		}
		tags = append(tags, tag)
	}

	post := models.NewPost()
	post.PostId = post_id
	post.CommentStatus = comment_status
	post.PostTitle = post_title
	post.PostContent = post_content
	post.PostContentMd = post_content_md
	post.PostModifiedTime = time.Now().Format("2006-01-02 15:04:05")
	post.Tags = tags

	err = models.NewPostModel().PostEdit(post)
	if err != nil {
		panic(err)
	}
	this.Redirect(fmt.Sprintf("/post/%d", post_id), 302)
}
