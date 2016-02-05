package models

import (
	"encoding/xml"
	"time"
)

type RSSPost struct {
	XMLName     xml.Name  `xml:"item"`
	PostTitle   string    `xml:"title"`
	PostLink    string    `xml:"link"`
	PostContent string    `xml:"description"`
	PostTime    time.Time `xml:"pubDate"`
}

func NewRSSPost(post *Post) *RSSPost {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", post.PostTime, loc)
	return &RSSPost{
		PostTitle:   post.PostTitle,
		PostLink:    post.PostLink(),
		PostContent: post.PostContent,
		PostTime:    time,
	}
}

type RSSChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	ChannelTitle  string    `xml:"title"`
	ChannelLink   string    `xml:"link"`
	ChannelDesc   string    `xml:"description"`
	ChannelEditor string    `xml:"managingEditor"`
	PubDate       time.Time `xml:"pubDate"`
	RSSPosts      []*RSSPost
}

func NewRSSChannel() *RSSChannel {
	return &RSSChannel{}
}

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *RSSChannel
}

func NewRSSFeed() *RSSFeed {
	return &RSSFeed{
		Version: "2.0",
	}
}
