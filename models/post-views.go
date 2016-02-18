package models

import "fmt"

const (
	TABLE_NAME_POST_VIEWS = "t_post_views"
)

type PostViews struct {
	Post     *Post  `orm:"pk;rel(fk)"`
	ViewedBy string `orm:"size(16)"`
	Views    int    `orm:"default(0)"`
}

func (this *PostViews) TableName() string {
	return TABLE_NAME_POST_VIEWS
}

func (this *PostViews) TableUnique() [][]string {
	return [][]string{
		[]string{"PostId", "ViewedBy"},
	}
}

func (this *PostModel) ViewedBy(post_id int, viewed_by string) error {
	o := ORM()
	_, err := o.Raw(fmt.Sprintf("insert into %s (post_id,viewed_by,views) values (?,?,1) on duplicate key update views=views+1", TABLE_NAME_POST_VIEWS), post_id, viewed_by).Exec()
	return err
}

func (this *PostModel) MostPopular(top int) ([]*PostViews, error) {
	o := ORM()
	var views []*PostViews
	_, err := o.QueryTable(new(PostViews)).Filter("ViewedBy", "human").OrderBy("-Views").Limit(top, 0).RelatedSel().All(&views)
	return views, err
}
