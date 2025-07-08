package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  uint
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	PostID uint
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
