package model

import "time"

type ProductTag struct {
	ID      uint    `gorm:"primaryKey;autoIncrement"`
	Color 	uint
	Name 	string 	`gorm:"uniqueIndex"`

	CreateUID *int
	CreateResUser ResUser `gorm:"foreignKey:CreateUID;"`
	CreateAt  time.Time
	WriteUID *int
	WriteResUser  ResUser `gorm:"foreignKey:WriteUID;"`
	WriteAt   time.Time
}
