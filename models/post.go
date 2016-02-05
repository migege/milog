package models

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	TABLE_NAME_POST = "t_posts"
)

type Post struct {
	PostId           int    `orm:"auto"`
	PostTitle        string `orm:"size(128)"`
	PostSlug         string `orm:"size(128)"`
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

func (this *PostModel) Count(filter string, v interface{}) (int64, error) {
	if filter != "" {
		return ORM().QueryTable(TABLE_NAME_POST).Filter(filter, v).Count()
	} else {
		return ORM().QueryTable(TABLE_NAME_POST).Count()
	}
}

func (this *PostModel) Offset(orderby string, offset, limit int) ([]*Post, error) {
	o := ORM()
	var posts []*Post
	_, err := o.QueryTable(TABLE_NAME_POST).OrderBy(orderby).Limit(limit, offset).RelatedSel().All(&posts)
	if err != nil {
		return posts, err
	}
	for _, post := range posts {
		_, _ = o.LoadRelated(post, "Tags")
	}
	return posts, err
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

func (this *PostModel) ById(id int) (*Post, error) {
	o := ORM()
	post := &Post{PostId: id}
	err := o.QueryTable(TABLE_NAME_POST).Filter("PostId", id).RelatedSel().One(post)
	if err != nil {
		return post, err
	}
	_, err = o.LoadRelated(post, "Tags")
	if err != nil {
		return post, err
	}
	return post, err
}

func (this *PostModel) BySlug(post_slug string) (*Post, error) {
	o := ORM()
	post := &Post{PostSlug: post_slug}
	err := o.QueryTable(TABLE_NAME_POST).Filter("PostSlug", post_slug).RelatedSel().One(post)
	if err != nil {
		return post, err
	}
	_, err = o.LoadRelated(post, "Tags")
	if err != nil {
		return post, err
	}
	return post, err
}

func (this *PostModel) ByAuthorId(author_id int, orderby string, offset, limit int) ([]*Post, error) {
	o := ORM()
	var posts []*Post
	_, err := o.QueryTable(TABLE_NAME_POST).Filter("Author__AuthorId", author_id).OrderBy(orderby).Limit(limit, offset).RelatedSel().All(&posts)
	return posts, err
}

func (this *PostModel) ByTagName(tag_name, orderby string, offset, limit int) ([]*Post, error) {
	o := ORM()
	var posts []*Post
	_, err := o.QueryTable(TABLE_NAME_POST).Filter("Tags__Tag__TagName", tag_name).OrderBy(orderby).Limit(limit, offset).RelatedSel().All(&posts)
	if err != nil {
		return posts, err
	}
	for _, post := range posts {
		o.LoadRelated(post, "Tags")
	}
	return posts, err
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
			"PostSlug":         post.PostSlug,
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
