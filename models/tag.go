package models

import (
	"fmt"
	"strings"
)

const (
	TABLE_NAME_TAG              = "t_tags"
	TABLE_NAME_TAG_RELATIONSHIP = "t_tag_relationships"
)

type Tag struct {
	TagId   int     `orm:"auto"`
	TagName string  `orm:"size(128)"`
	TagSlug string  `orm:"size(128)"`
	Posts   []*Post `orm:"reverse(many)"`
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) TableName() string {
	return TABLE_NAME_TAG
}

func (this *Tag) formatSlug() {
	this.TagSlug = strings.Replace(strings.ToLower(this.TagName), " ", "-", -1)
}

type TagRelationship struct {
	Id   int   `orm:"auto"`
	Post *Post `orm:"rel(fk)"`
	Tag  *Tag  `orm:"rel(fk)"`
}

func (this *TagRelationship) TableName() string {
	return TABLE_NAME_TAG_RELATIONSHIP
}

func (this *TagRelationship) TableUnique() [][]string {
	return [][]string{
		[]string{"PostId", "TagId"},
	}
}

type TagModel struct {
}

func NewTagModel() *TagModel {
	return &TagModel{}
}

func (this *TagModel) BySlug(tag_slug string) (*Tag, error) {
	o := ORM()
	tag := NewTag()
	err := o.QueryTable(NewTag()).Filter("TagSlug", tag_slug).One(tag)
	return tag, err
}

func (this *TagModel) AddTags(tags []*Tag) {
	o := ORM()
	for _, tag := range tags {
		if tag == nil {
			continue
		}
		tag.formatSlug()
		o.ReadOrCreate(tag, "TagSlug")
	}
}

func (this *TagModel) AllTags() ([]*Tag, error) {
	o := ORM()
	var tags []*Tag
	_, err := o.QueryTable(TABLE_NAME_TAG).All(&tags)
	return tags, err
}

type TagCount struct {
	TagName string
	TagSlug string
	Counts  int
}

func (this *TagModel) MostPopular(limit int) ([]TagCount, error) {
	o := ORM()
	var counts []TagCount
	_, err := o.Raw(fmt.Sprintf("select tag_slug,tag_name,counts from ( select t1.tag_id,t2.tag_slug,t2.tag_name,count(distinct t1.post_id) counts from t_tag_relationships t1 join t_tags t2 on t1.tag_id=t2.tag_id join t_posts t3 on t1.post_id=t3.post_id where t3.post_status=0 group by tag_id order by counts desc limit %d) t where 1=1 order by tag_name", limit)).QueryRows(&counts)
	return counts, err
}
