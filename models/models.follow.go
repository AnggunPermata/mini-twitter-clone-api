package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowedUserID uint   `json:"followed-user-id" form:"followed-user-id"` //FK from user ID
	UserID         uint   `json:"user-id" form:"user-id"`                   //the current user accout
	FullName       string `json:"full-name" form:"full-name"`
	Username       string `json:"username" form:"username"`
	FollowStatus   string `json:"follow-status" form:"follow-status"` //status: follow, unfollow
}
