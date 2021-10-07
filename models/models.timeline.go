package models

import "gorm.io/gorm"

type Timeline struct {
	gorm.Model
	UserID    uint
	Username  string `json:"username" form:"username"`
	FullName  string `json:"full-name" form:"full-name"`
	TweetID   uint
	CommentID uint
}

type InputTimeline struct {
	UserID uint `json:"user-id" form:"user-id"` //the current user accout
}
