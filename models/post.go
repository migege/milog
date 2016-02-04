package models

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	TABLE_NAME_POST = "t_posts"
)

type Post struct {
	PostId           int    `orm:"auto"`
	PostTitle        string `orm:"size(128)"`
	PostContent      string
	PostContentMd    string
	PostTime         string `orm:"auto_now_add;type(datetime)"`
	PostModifiedTime string `orm:"auto_now;type(datetime)"`
	CommentStatus    int
	CommentCount     int
	Author           *Author `orm:"rel(fk);on_delete(do_nothing)"`
	Tags             []*Tag  `orm:"rel(m2m);rel_through(github.com/migege/milog/models.TagRelationship)"`
}

func NewPost() *Post {
	now := time.Now().Format("2006-01-02 15:04:05")
	return &Post{
		PostTime:         now,
		PostModifiedTime: now,
		CommentStatus:    1,
	}
}

func (this *Post) TableName() string {
	return TABLE_NAME_POST
}

type PostModel struct {
}

func NewPostModel() *PostModel {
	return new(PostModel)
}

func (this *PostModel) All(orderby string) []*Post {
	o := ORM()
	var posts []*Post
	_, err := o.QueryTable(TABLE_NAME_POST).OrderBy(orderby).RelatedSel().All(&posts)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		_, err = o.LoadRelated(post, "Tags")
	}
	return posts
}

func (this *PostModel) ById(id int) *Post {
	o := ORM()
	post := &Post{PostId: id}
	err := o.QueryTable(TABLE_NAME_POST).Filter("PostId", id).RelatedSel().One(post)
	if err != nil {
		panic(err)
	}
	_, err = o.LoadRelated(post, "Tags")
	if err != nil {
		panic(err)
	}
	return post
}

func (this *PostModel) ByAuthorId(author_id int, orderby string) []*Post {
	o := ORM()
	var posts []*Post
	o.QueryTable(TABLE_NAME_POST).Filter("Author__AuthorId", author_id).OrderBy(orderby).RelatedSel().All(&posts)
	return posts
}

func (this *PostModel) ByTagName(tag_name, orderby string) []*Post {
	o := ORM()
	var posts []*Post
	_, err := o.QueryTable(TABLE_NAME_POST).Filter("Tags__Tag__TagName", tag_name).RelatedSel().OrderBy(orderby).All(&posts)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		_, err = o.LoadRelated(post, "Tags")
	}
	return posts
}

func (this *PostModel) PostNew(post *Post) (int64, error) {
	o := ORM()
	o.Begin()
	post_id, err := o.Insert(post)
	if err != nil {
		// error inserting post
		o.Rollback()
	} else {
		if len(post.Tags) > 0 {
			m2m := o.QueryM2M(post, "Tags")
			NewTagModel().AddTags(post.Tags)
			_, err = m2m.Add(post.Tags)
			if err != nil {
				// error adding tags
				o.Rollback()
			} else {
				// ok
				o.Commit()
			}
		} else {
			// no tags to add
			o.Commit()
		}
	}
	return post_id, err
}

func (this *PostModel) PostEdit(post *Post) error {
	o := ORM()
	o.Begin()
	_, err := o.QueryTable(TABLE_NAME_POST).Filter("PostId", post.PostId).Update(
		orm.Params{
			"CommentStatus":    post.CommentStatus,
			"PostTitle":        post.PostTitle,
			"PostContent":      post.PostContent,
			"PostContentMd":    post.PostContentMd,
			"PostModifiedTime": post.PostModifiedTime,
		})
	if err != nil {
		// error updating post
		o.Rollback()
	} else {
		if err != nil {
			// error inserting tags
			o.Rollback()
		} else {
			m2m := o.QueryM2M(post, "Tags")
			_, err = m2m.Clear()
			if err != nil {
				// error clearing tags
				o.Rollback()
			} else {
				if len(post.Tags) > 0 {
					NewTagModel().AddTags(post.Tags)
					_, err = m2m.Add(post.Tags)
					if err != nil {
						// error adding tags
						o.Rollback()
					} else {
						// ok
						o.Commit()
					}
				} else {
					// no tags to add
					o.Commit()
				}
			}
		}
	}

	return err
}
