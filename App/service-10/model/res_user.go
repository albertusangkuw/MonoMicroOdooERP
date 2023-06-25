package model

type ResUser struct {
	UID int  `gorm:"primaryKey"`
	DisplayName string
	GroupID int
}

func (u ResUser) ToArray() []interface{} {
	return []interface{}{u.UID, u.DisplayName}
}
