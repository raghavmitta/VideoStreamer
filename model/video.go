package model

import "time"

type Video struct {
	VideoId          string    `json:"VideoId"`
	PublishedAt      time.Time `json:"PublishedAt"`
	Title            string    `json:"Title"`
	Description      string    `json:"Description "`
	DefaultThumbnail string    `json:"DefaultThumbnail"`
	MediumThumbnail  string    `json:"MediumThumbnail"`
	HighThumbnail    string    `json:"HighThumbnail"`
}

func NewVideo(videoId string, publishedAt time.Time, title string, description string, defaultThumbnail string, mediumThumbnail string, highThumbnail string) *Video {
	return &Video{VideoId: videoId, PublishedAt: publishedAt, Title: title, Description: description, DefaultThumbnail: defaultThumbnail, MediumThumbnail: mediumThumbnail, HighThumbnail: highThumbnail}
}
