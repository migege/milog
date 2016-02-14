package models

import (
	_ "fmt"
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

func (this *TagModel) AddTags(tags []*Tag) {
	o := ORM()
	for _, tag := range tags {
		_, _, _ = o.ReadOrCreate(tag, "TagSlug")
	}
}

func (this *TagModel) AllTags() ([]*Tag, error) {
	o := ORM()
	var tags []*Tag
	_, err := o.QueryTable(TABLE_NAME_TAG).All(&tags)
	return tags, err
}
