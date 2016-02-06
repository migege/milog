package models

import (
	"encoding/xml"
	"time"
)

type RSSPost struct {
	XMLName     xml.Name `xml:"item"`
	PostTitle   string   `xml:"title"`
	PostLink    string   `xml:"link"`
	PostContent string   `xml:"description"`
	PostTime    string   `xml:"pubDate"`
	GUID        string   `xml:"guid"`
}

func NewRSSPost(post *Post) *RSSPost {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	post_time, _ := time.ParseInLocation("2006-01-02 15:04:05", post.PostTime, loc)
	return &RSSPost{
		PostTitle:   post.PostTitle,
		PostLink:    post.PostLink(),
		PostContent: post.PostContent,
		PostTime:    post_time.Format("Mon, 02 Jan 2006 15:04:05 MST"),
	}
}

type RSSSelf struct {
	XMLName xml.Name `xml:"atom:link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

func NewRSSSelf(href string) *RSSSelf {
	return &RSSSelf{
		Href: href,
		Rel:  "self",
		Type: "application/rss+xml",
	}
}

type RSSChannel struct {
	XMLName      xml.Name `xml:"channel"`
	ChannelTitle string   `xml:"title"`
	ChannelLink  string   `xml:"link"`
	ChannelDesc  string   `xml:"description"`
	PubDate      string   `xml:"pubDate"`
	Self         *RSSSelf
	RSSPosts     []*RSSPost
}

func NewRSSChannel() *RSSChannel {
	return &RSSChannel{}
}

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	NS      string   `xml:"xmlns:atom,attr"`
	Channel *RSSChannel
}

func NewRSSFeed() *RSSFeed {
	return &RSSFeed{
		Version: "2.0",
		NS:      "http://www.w3.org/2005/Atom",
	}
}
