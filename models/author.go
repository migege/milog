package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	TABLE_NAME_AUTHOR = "t_authors"
)

var (
	ErrWrongUserOrPass = errors.New("error: wrong username or password")
	ErrTimedOut        = errors.New("error: timed out")
)

type Author struct {
	AuthorId             int    `orm:"auto"`
	AuthorName           string `orm:"unique;size(128)"`
	AuthorPassword       string `orm:"size(64)"`
	AuthorMail           string `orm:"unique;size(128)"`
	AuthorUrl            string `orm:"size(128)"`
	AuthorRegisteredTime string `orm:"auto_now_add;type(datetime)"`
	AuthorStatus         int
	DisplayName          string  `orm:"unique;size(128)"`
	Posts                []*Post `orm:"reverse(many)"`
}

func (this *Author) TableName() string {
	return TABLE_NAME_AUTHOR
}

type AuthorModel struct {
}

func NewAuthorModel() *AuthorModel {
	return new(AuthorModel)
}

func (this *AuthorModel) ById(id int) *Author {
	o := ORM()
	author := &Author{AuthorId: id}
	err := o.Read(author)
	if err != nil {
		panic(err)
	}
	return author
}

func (this *AuthorModel) ByName(name string) *Author {
	author := &Author{AuthorName: name}
	o := ORM()
	err := o.QueryTable(TABLE_NAME_AUTHOR).Filter("AuthorName", name).One(author)
	if err != nil {
		panic(err)
	}
	return author
}

// the password is a signature actually, never use plain password
func (this *AuthorModel) Validate(ts, name, password string) error {
	now := time.Now().UnixNano()
	tsn, err := strconv.ParseInt(ts, 10, 64)
	if err != nil || now-tsn > 2*60*1000*1000*1000 {
		return ErrTimedOut
	}
	o := ORM()
	author := &Author{}
	o.QueryTable(TABLE_NAME_AUTHOR).Filter("AuthorName", name).One(author)
	// user exists
	if author.AuthorName == name {
		h := hmac.New(sha256.New, []byte(name))
		fmt.Fprintf(h, "%s%s", ts, author.AuthorPassword)
		if fmt.Sprintf("%02x", h.Sum(nil)) == password {
			return nil
		}
	}
	return ErrWrongUserOrPass
}
