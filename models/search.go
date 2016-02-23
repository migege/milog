package models

import (
	"errors"

	"github.com/yunge/sphinx"
)

func SphinxClient() *sphinx.Client {
	sc := sphinx.NewClient()
	if err := sc.Error(); err != nil {
		panic(err)
	}
	return sc
}

type SearchModel struct {
}

func NewSearchModel() *SearchModel {
	return &SearchModel{}
}

func (this *SearchModel) Search(term string) ([]*Post, error) {
	sc := SphinxClient()
	defer sc.Close()

	res, err := sc.Query(term, "", "")
	if err != nil {
		return nil, err
	}

	var post_id []int
	for _, match := range res.Matches {
		post_id = append(post_id, int(match.DocId))
	}

	if len(post_id) < 1 {
		return nil, errors.New("no search results.")
	}

	return NewPostModel().IdIn(post_id)
}
