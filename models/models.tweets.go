package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	UserID    uint
	Username  string `json:"username" form:"username"`
	FullName  string `json:"full-name" form:"full-name"`
	TweetData string `json:"tweet-data" form:"tweet-data"`
	Favorite  int    `json:"favorite" form:"favorite"`
	Retweet   int    `json:"retweet" form:"retweet"`
	CommentID uint
}
