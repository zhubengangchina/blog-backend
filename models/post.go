package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"type:varchar(200);not null"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint
	User     User
	Comments []Comment `gorm:"foreignKey:PostID"`
}
