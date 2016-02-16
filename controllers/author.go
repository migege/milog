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

func (this *AuthorController) ById() {
	author_id_str := this.Ctx.Input.Param(":id")
	author_id, _ := strconv.Atoi(author_id_str)
	author, err := models.NewAuthorModel().ById(author_id)
	if err != nil {
		this.Abort("404")
	}
	this.Redirect(fmt.Sprintf("/author/%s", author.AuthorName), 301)
}

func (this *AuthorController) ByName() {
	this.TplName = "author.tpl"

	author_name := this.Ctx.Input.Param(":name")

	if post_count, err := models.NewPostModel().Count("Author__AuthorName", author_name); err == nil {
		paginator := pagination.SetPaginator(this.Ctx, postsPerPage, post_count)
		if author, err := models.NewAuthorModel().ByName(author_name); err == nil {
			this.Data["Author"] = author
			posts, err := models.NewPostModel().ByAuthorId(author.AuthorId, "-PostId", paginator.Offset(), postsPerPage)
			if err != nil {
				panic(err)
			}
			this.Data["Posts"] = posts
			this.Data["PageTitle"] = fmt.Sprintf("%s - Author - %s", author.DisplayName, blogTitle)
		} else {
			this.Abort("404")
		}
	} else {
		panic(err)
	}
}

func (this *AuthorController) Signup() {
	this.TplName = "signup.tpl"
	this.Data["PageTitle"] = fmt.Sprintf("Sign Up - %s", blogTitle)
}
