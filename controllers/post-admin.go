package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/migege/milog/models"
)

func (this *AdminController) AllPosts() {
	this.TplName = "admin-posts.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("All Posts - %s - Admin - %s", this.loggedUser.DisplayName, blogTitle)
	posts := models.NewPostModel().All("-PostId", true, true)
	this.Data["Posts"] = posts

	views := make(map[int]int)
	for _, post := range posts {
		for _, view := range post.PostViews {
			if view.ViewedBy == "human" {
				views[post.PostId] = view.Views
				break
			}
		}
	}
	this.Data["Views"] = views
}

func (this *AdminController) PostNew() {
	this.TplName = "admin-post.tpl"
	tags, _ := models.NewTagModel().AllTags()
	this.Data["AllTags"] = tags
	this.Data["PageTitle"] = fmt.Sprintf("New Post - Admin - %s", blogTitle)
}

func (this *AdminController) DoPostNew() {
	post_title := this.GetString("post-title")
	post_slug := this.GetString("post-slug")
	_, err := strconv.Atoi(post_slug)
	if err == nil {
		panic(errors.New("error: post slug should not be a number"))
	}
	post_content := this.GetString("post-content")
	post_content_md := this.GetString("post-content-md")
	if post_title == "" || post_content == "" || post_content_md == "" || post_slug == "" {
		panic(errors.New("error: empty post title, slug or content"))
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
	post.PostSlug = post_slug
	post.PostContent = post_content
	post.PostContentMd = post_content_md
	post.Author = this.loggedUser
	_, err = models.NewPostModel().PostNew(post)
	if err != nil {
		panic(err)
	} else {
		this.Redirect(post.PostLink(), 302)
	}
}

func (this *AdminController) PostEdit() {
	this.TplName = "admin-post.tpl"
	id := this.Ctx.Input.Param(":id")
	post_id, _ := strconv.Atoi(id)

	post, err := models.NewPostModel().ById(post_id, true)
	if err != nil {
		panic(err)
	} else if post.Author.AuthorId != this.loggedUser.AuthorId {
		panic(errors.New("error: can't edit another one's post"))
	}

	tags, _ := models.NewTagModel().AllTags()
	this.Data["Post"] = post
	this.Data["AllTags"] = tags
	this.Data["PageTitle"] = fmt.Sprintf("Editing Post: %s - %s", post.PostTitle, blogTitle)
}

func (this *AdminController) DoPostEdit() {
	id := this.GetString("post-id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	post_title := this.GetString("post-title")
	post_slug := this.GetString("post-slug")
	_, err = strconv.Atoi(post_slug)
	if err == nil {
		panic(errors.New("error: post slug should not be a number"))
	}
	post_content := this.GetString("post-content")
	post_content_md := this.GetString("post-content-md")
	if post_title == "" || post_content == "" || post_content_md == "" || post_slug == "" {
		panic(errors.New("error: empty post title, slug or content"))
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
	post.PostSlug = post_slug
	post.PostContent = post_content
	post.PostContentMd = post_content_md
	post.PostModifiedTime = time.Now().Format("2006-01-02 15:04:05")
	post.Tags = tags

	err = models.NewPostModel().PostEdit(post)
	if err != nil {
		panic(err)
	}
	this.Redirect(post.PostLink(), 302)
}

func (this *AdminController) PostDelete() {
	id := this.Ctx.Input.Param(":id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	post, err := models.NewPostModel().ById(post_id)
	if err != nil {
		panic(err)
	}
	post.PostStatus = -1
	err = post.Update("PostStatus")
	if err != nil {
		panic(err)
	}
	this.Redirect("/admin", 302)
}

func (this *AdminController) PostRestore() {
	id := this.Ctx.Input.Param(":id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	post, err := models.NewPostModel().ById(post_id, true)
	if err != nil {
		panic(err)
	}
	post.PostStatus = 0
	err = post.Update("PostStatus")
	if err != nil {
		panic(err)
	}
	this.Redirect("/admin", 302)
}
