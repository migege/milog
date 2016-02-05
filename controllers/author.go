package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/migege/milog/models"
)

type AuthorController struct {
	BaseController
}

func (this *AuthorController) Get() {
	this.TplName = "author.tpl"

	author_id_str := this.Ctx.Input.Param(":id")
	author_id, _ := strconv.Atoi(author_id_str)

	post_count, err := models.NewPostModel().Count("Author__AuthorId", author_id)
	if err != nil {
		panic(err)
	}
	paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)

	posts, err := models.NewPostModel().ByAuthorId(author_id, "-PostId", paginator.Offset(), postsPerPage)
	if err != nil {
		panic(err)
	}
	this.Data["Posts"] = posts

	author := &models.Author{}
	if len(posts) > 0 {
		author = posts[0].Author
		this.Data["Author"] = author
	} else {
		author = models.NewAuthorModel().ById(author_id)
	}

	this.Data["PageTitle"] = fmt.Sprintf("%s - Author - %s", author.DisplayName, blogTitle)
}

func (this *AuthorController) Signup() {
	this.TplName = "signup.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("Sign Up - %s", blogTitle)
}
