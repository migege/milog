package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

const (
	TABLE_NAME_COMMENT = "t_comments"
)

const (
	COMMENT_STATUS_NEW      = 0
	COMMENT_STATUS_APPROVED = 1
)

type Comment struct {
	CommentId         int `orm:"auto"`
	CommentAuthor     string
	CommentAuthorMail string `orm:"size(128)"`
	CommentAuthorUrl  string `orm:"size(128)"`
	CommentAuthorIp   string `orm:"size(128)"`
	CommentTime       string `orm:"auto_now_add;type(datetime)"`
	CommentContent    string
	CommentStatus     int `orm:"column(comment_approved)"`
	CommentAgent      string
	CommentType       int
	CommentParentId   int
	Post              *Post `orm:"rel(one);on_delete(do_nothing)"`
}

func NewComment() *Comment {
	defaultCommentStatusStr := NewOptionModel().Name(OPTION_COMMENT_DEFAULT_STATUS).OptionValue
	defaultCommentStatus, err := strconv.Atoi(defaultCommentStatusStr)
	if err != nil {
		defaultCommentStatus = COMMENT_STATUS_NEW
	}
	comment := &Comment{
		CommentTime:       time.Now().Format("2006-01-02 15:04:05"),
		CommentAuthor:     "anonymous",
		CommentAuthorMail: "nobody@nowhere.com",
		CommentAuthorUrl:  "http://",
		CommentStatus:     defaultCommentStatus,
	}
	return comment
}

func (this *Comment) TableName() string {
	return TABLE_NAME_COMMENT
}

type CommentModel struct {
}

func NewCommentModel() *CommentModel {
	return new(CommentModel)
}

func (this *CommentModel) ById(id int) *Comment {
	o := ORM()
	comment := Comment{CommentId: id}
	err := o.Read(&comment)
	if err != nil {
		panic(err)
	}
	return &comment
}

func (this *CommentModel) ByPostId(id int, orderby string) []*Comment {
	o := ORM()
	var comments []*Comment
	o.QueryTable(TABLE_NAME_COMMENT).Filter("Post__PostId", id).Filter("CommentStatus", COMMENT_STATUS_APPROVED).OrderBy(orderby).All(&comments)
	return comments
}

func (this *CommentModel) Latest(limit int) ([]*Comment, error) {
	o := ORM()
	var comments []*Comment
	_, err := o.QueryTable(TABLE_NAME_COMMENT).Filter("CommentStatus", COMMENT_STATUS_APPROVED).OrderBy("-CommentTime").Limit(limit).RelatedSel().All(&comments)
	return comments, err
}

func (this *CommentModel) AddComment(comment *Comment) (int64, error) {
	o := ORM()
	o.Begin()
	comment_id, err := o.Insert(comment)
	if err == nil {
		// adding comment count of the post
		_, err := o.QueryTable(TABLE_NAME_POST).Filter("PostId", comment.Post.PostId).Update(orm.Params{
			"CommentCount": orm.ColValue(orm.ColAdd, 1),
		})
		if err != nil {
			// error adding count
			o.Rollback()
		} else {
			o.Commit()
		}
	} else {
		// error adding comment
		o.Rollback()
	}
	return comment_id, err
}
