package model

import (
	"time"
)

type PosCategory struct {
	ID      uint    `gorm:"primaryKey;autoIncrement"`
	ParentID uint
	ChildID map[uint]uint `gorm:"-"`

	Name  string
	Sequence int


	ImageID *int
	Image PosCategoryImage `gorm:"foreignKey:ImageID;"`

	CreateUID *int
	CreateResUser ResUser `gorm:"foreignKey:CreateUID;"`
	CreateAt  time.Time
	WriteUID *int
	WriteResUser  ResUser `gorm:"foreignKey:WriteUID;"`
	WriteAt   time.Time
}

func (p *PosCategory)DisplayName() string{
	return p.Name
}




