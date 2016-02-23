package routers

import (
	"github.com/astaxie/beego"
	"github.com/migege/milog/controllers"
)

func init() {
	// home
	beego.Router("/", &controllers.MainController{})

	// feed
	beego.Router("/rss", &controllers.FeedController{}, "get:RSS")
	beego.Router("/feed", &controllers.FeedController{}, "get:RSS")

	// authors
	beego.Router("/author/:id([0-9]+)", &controllers.AuthorController{}, "get:ById")
	beego.Router("/author/:name", &controllers.AuthorController{}, "get:ByName")

	// comments
	beego.Router("/comments-add", &controllers.CommentController{}, "post:DoAddComment")

	// login/logout
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	// signup
	beego.Router("/signup", &controllers.AuthorController{}, "get:Signup")

	// tags
	beego.Router("/tag/:tag([\\w\\-]+)", &controllers.TagController{}, "get:ByName")

	// search
	beego.Router("/search/:term", &controllers.SearchController{}, "get:Search")

	// posts
	beego.Router("/post/:id([0-9]+)", &controllers.PostController{}, "get:ById")
	beego.Router("/post/:slug", &controllers.PostController{}, "get:BySlug")

	///////////////////////////////////////////////////////////////////
	// admin pages begin
	beego.Router("/admin", &controllers.AdminController{})

	beego.Router("/admin/options", &controllers.OptionController{}, "get:Basic")
	beego.Router("/admin/option-edit", &controllers.OptionController{}, "post:DoEdit")

	beego.Router("/admin/posts", &controllers.AdminController{}, "get:AllPosts")
	beego.Router("/admin/post-new", &controllers.AdminController{}, "get:PostNew")
	beego.Router("/admin/post-new", &controllers.AdminController{}, "post:DoPostNew")
	beego.Router("/admin/post-edit/:id([0-9]+)", &controllers.AdminController{}, "get:PostEdit")
	beego.Router("/admin/post-edit", &controllers.AdminController{}, "post:DoPostEdit")
	beego.Router("/admin/post-delete/:id([0-9]+)", &controllers.AdminController{}, "get:PostDelete")
	beego.Router("/admin/post-restore/:id([0-9]+)", &controllers.AdminController{}, "get:PostRestore")
	// admin pages end
	///////////////////////////////////////////////////////////////////
}
