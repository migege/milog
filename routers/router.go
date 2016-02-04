package routers

import (
	"github.com/astaxie/beego"
	"github.com/migege/milog/controllers"
)

func init() {
	// basic pages
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{})
	beego.Router("/author/:id([0-9]+)", &controllers.AuthorController{})

	// comments
	beego.Router("/comments-add", &controllers.CommentController{}, "post:DoAddComment")

	// login/logout
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	// signup
	beego.Router("/signup", &controllers.AuthorController{}, "get:Signup")

	// admin pages
	beego.Router("/admin", &controllers.AdminController{})

	// options
	beego.Router("/admin/options", &controllers.OptionController{}, "get:Basic")
	beego.Router("/admin/option-edit", &controllers.OptionController{}, "post:DoEdit")

	// posts
	beego.Router("/admin/post-new", &controllers.PostController{}, "get:PostNew")
	beego.Router("/admin/post-new", &controllers.PostController{}, "post:DoPostNew")
	beego.Router("/admin/post-edit/:id([0-9]+)", &controllers.PostController{}, "get:PostEdit")
	beego.Router("/admin/post-edit", &controllers.PostController{}, "post:DoPostEdit")

	// tags
	beego.Router("/tag/:tag([\\w]+)", &controllers.TagController{}, "get:ByName")
}
