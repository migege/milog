package models

import (
	"fmt"
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
	PostStatus       int    `orm:"default(0)"`
	CommentStatus    int
	CommentCount     int
	Author           *Author      `orm:"rel(fk);on_delete(do_nothing)"`
	Tags             []*Tag       `orm:"rel(m2m);rel_through(github.com/migege/milog/models.TagRelationship)"`
	PostViews        []*PostViews `orm:"reverse(many)"`
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

func (this *Post) PostLink() string {
	return fmt.Sprintf("/post/%s", this.PostSlug)
}

func (this *Post) Update(cols ...string) error {
	o := ORM()
	this.PostModifiedTime = time.Now().Format("2006-01-02 15:04:05")
	cols = append(cols, "PostModifiedTime")
	_, err := o.Update(this, cols...)
	return err
}

type PostModel struct {
}

func NewPostModel() *PostModel {
	return new(PostModel)
}

func (this *PostModel) parseArgs(args ...interface{}) orm.Params {
	params := orm.Params{
		"ignore_post_status": false,
		"load_tags":          true,
		"load_views":         false,
	}
	for i, arg := range args {
		switch i {
		case 0:
			if v, ok := arg.(bool); ok {
				params["ignore_post_status"] = v
			}
		case 1:
			if v, ok := arg.(bool); ok {
				params["load_tags"] = v
			}
		case 2:
			if v, ok := arg.(bool); ok {
				params["load_views"] = v
			}
		}
	}
	return params
}

func (this *PostModel) Count(filter string, v interface{}, args ...interface{}) (int64, error) {
	params := this.parseArgs(args...)
	qs := ORM().QueryTable(new(Post))
	switch {
	case filter != "":
		qs = qs.Filter(filter, v)
	case params["ignore_post_status"] == false:
		qs = qs.Filter("PostStatus", 0)
	}
	return qs.Count()
}

func (this *PostModel) Offset(filter string, v interface{}, orderby string, offset, limit int, args ...interface{}) ([]*Post, error) {
	params := this.parseArgs(args...)
	o := ORM()
	qs := o.QueryTable(new(Post))
	switch {
	case filter != "":
		qs = qs.Filter(filter, v)
	case params["ignore_post_status"] == false:
		qs = qs.Filter("PostStatus", 0)
	}
	qs = qs.OrderBy(orderby).Limit(limit, offset).RelatedSel()
	var posts []*Post
	_, err := qs.All(&posts)
	if err != nil {
		return posts, err
	}
	for _, post := range posts {
		if params["load_tags"] == true {
			o.LoadRelated(post, "Tags")
		}
		if params["load_views"] == true {
			o.LoadRelated(post, "PostViews")
		}
	}
	return posts, err
}

func (this *PostModel) ById(id int, args ...interface{}) (*Post, error) {
	posts, err := this.Offset("PostId", id, "-PostId", 0, -1, args...)
	if err != nil || len(posts) != 1 {
		return nil, err
	}
	return posts[0], err
}

func (this *PostModel) BySlug(post_slug string, args ...interface{}) (*Post, error) {
	posts, err := this.Offset("PostSlug", post_slug, "-PostId", 0, -1, args...)
	if err != nil || len(posts) != 1 {
		return nil, err
	}
	return posts[0], err
}

func (this *PostModel) ByAuthorId(author_id int, orderby string, offset, limit int, args ...interface{}) ([]*Post, error) {
	return this.Offset("Author__AuthorId", author_id, orderby, offset, limit, args...)
}

func (this *PostModel) ByTagName(tag_name, orderby string, offset, limit int, args ...interface{}) ([]*Post, error) {
	return this.Offset("Tags__Tag__TagName", tag_name, orderby, offset, limit, args...)
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
