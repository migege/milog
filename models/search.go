package models

import (
	"errors"
	"sort"

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

type WeightedPost struct {
	Post   *Post
	Weight int
}

type WeightedPosts []WeightedPost

func (s WeightedPosts) Len() int {
	return len(s)
}

func (s WeightedPosts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByWeight struct {
	WeightedPosts
}

func (s ByWeight) Less(i, j int) bool {
	return s.WeightedPosts[i].Weight > s.WeightedPosts[j].Weight
}

func (this *SearchModel) Search(term string) ([]*Post, error) {
	sc := SphinxClient()
	defer sc.Close()

	res, err := sc.Query(term, "", "")
	if err != nil {
		return nil, err
	}

	var post_id []int
	post_weight := make(map[int]int)
	for _, match := range res.Matches {
		post_id = append(post_id, int(match.DocId))
		post_weight[int(match.DocId)] = match.Weight
	}

	if len(post_id) < 1 {
		return nil, errors.New("no search results.")
	}

	posts, err := NewPostModel().IdIn(post_id)
	if err == nil {
		var weighted WeightedPosts
		for _, post := range posts {
			weighted = append(weighted, WeightedPost{post, post_weight[post.PostId]})
		}
		sort.Stable(ByWeight{weighted})

		var sorted []*Post
		for _, post := range weighted {
			sorted = append(sorted, post.Post)
		}
		return sorted, err
	} else {
		return posts, err
	}
}
