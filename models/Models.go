package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `joson:"isbn"`
}

type Author struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `gorm:"default:galeone" json:"middle_name"`
}
