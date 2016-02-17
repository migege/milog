package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
)

func init() {
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqldb := beego.AppConfig.String("mysqldb")
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlcharset := beego.AppConfig.String("mysqlcharset")
	conn_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb, mysqlcharset)

	//	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn_str)
	orm.RegisterModel(
		new(Post),
		new(Author),
		new(Comment),
		new(Option),
		new(Tag),
		new(TagRelationship),
		new(Link),
		new(PostViews),
	)
	o = orm.NewOrm()
}

func ORM() orm.Ormer {
	return o
}
