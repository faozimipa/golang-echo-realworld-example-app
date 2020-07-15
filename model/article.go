package model

import (
	"github.com/jinzhu/gorm"

)

//Article struct
type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	Title       string `gorm:"not null"`
	Description string
	Body        string
	Author      User
	AuthorID    uint
	Comments    []Comment
	Favorites   []User `gorm:"many2many:favorites;"`
	Tags        []Tag  `gorm:"many2many:article_tags;association_autocreate:false"`
}

//Comment struct
type Comment struct {
	gorm.Model
	Article   Article
	ArticleID uint
	User      User
	UserID    uint
	Body      string
}

//Tag struct
type Tag struct {
	gorm.Model
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags;"`
}
