package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID      uint
	Username    string `json:"username" form:"username"`
	FullName    string `json:"full-name" form:"full-name"`
	TweetID     uint
	TextComment string `json:"text-comment" form:"text-comment"`
}
