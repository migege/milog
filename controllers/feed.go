package controllers

import (
	"time"

	"github.com/migege/milog/models"
)

type FeedController struct {
	BaseController
}

func (this *FeedController) RSS() {
	posts, err := models.NewPostModel().Offset("-PostId", 0, 10)
	if err != nil {
		panic(err)
	}
	var rss_posts []*models.RSSPost
	for _, post := range posts {
		rss_post := models.NewRSSPost(post)
		rss_post.PostLink = blogUrl + rss_post.PostLink
		rss_posts = append(rss_posts, rss_post)
	}
	rss := models.NewRSSFeed()
	channel := models.NewRSSChannel()
	channel.PubDate = time.Now()
	channel.ChannelTitle = blogTitle
	channel.ChannelLink = blogUrl
	channel.ChannelDesc = blogDesc
	channel.RSSPosts = rss_posts
	rss.Channel = channel

	this.Data["xml"] = rss
	this.ServeXML()
}
