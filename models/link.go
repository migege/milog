package models

const (
	TABLE_NAME_LINK = "t_links"
)

type Link struct {
	LinkId   int    `orm:"auto"`
	LinkText string `orm:"size(128)"`
	LinkUrl  string `orm:"size(128)"`
	LinkDesc string `orm:"size(128)"`
}

func NewLink() *Link {
	return &Link{}
}

func (this *Link) TableName() string {
	return TABLE_NAME_LINK
}

type LinkModel struct {
}

func NewLinkModel() *LinkModel {
	return &LinkModel{}
}

func (this *LinkModel) AllLinks() ([]*Link, error) {
	o := ORM()
	var links []*Link
	_, err := o.QueryTable(TABLE_NAME_LINK).All(&links)
	return links, err
}
