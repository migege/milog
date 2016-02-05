package models

import (
	"net/url"
	"strconv"

	"github.com/astaxie/beego/orm"
)

const (
	TABLE_NAME_OPTION = "t_options"
)

const (
	OPTION_BLOG_TITLE             = "blog_title"
	OPTION_BLOG_DESC              = "blog_desc"
	OPTION_BLOG_URL               = "blog_url"
	OPTION_COMMENT_DEFAULT_STATUS = "comment_default_status"
	OPTION_POSTS_PER_PAGE         = "posts_per_page"
)

type Option struct {
	OptionId    int    `orm:"auto"`
	OptionName  string `orm:"unique;size(128)"`
	OptionValue string
	OptionDesc  string `orm:"size(128)"`
}

func (this *Option) TableName() string {
	return TABLE_NAME_OPTION
}

func (this *Option) GetInt() (int, error) {
	return strconv.Atoi(this.OptionValue)
}

type OptionModel struct {
}

func NewOptionModel() *OptionModel {
	return new(OptionModel)
}

func (this *OptionModel) Name(name string) *Option {
	o := ORM()
	option := Option{OptionName: name}
	o.QueryTable(TABLE_NAME_OPTION).Filter("OptionName", name).One(&option)
	return &option
}

func (this *OptionModel) Names(names *[]string) map[string]*Option {
	o := ORM()
	var options []*Option
	_, err := o.QueryTable(TABLE_NAME_OPTION).Filter("OptionName__in", names).All(&options)
	if err != nil {
		panic(err)
	}

	m := make(map[string]*Option, len(*names))
	for _, option := range options {
		m[option.OptionName] = option
	}
	for _, name := range *names {
		if _, ok := m[name]; !ok {
			m[name] = &Option{}
		}
	}
	return m
}

func (this *OptionModel) Save(data url.Values) error {
	o := ORM()
	o.Begin()
	for k := range data {
		if v := data.Get(k); v != "" {
			_, err := o.QueryTable(TABLE_NAME_OPTION).Filter("OptionName", k).Update(orm.Params{
				"OptionValue": v,
			})
			if err != nil {
				o.Rollback()
				return err
			}
		}
	}
	o.Commit()
	return nil
}
