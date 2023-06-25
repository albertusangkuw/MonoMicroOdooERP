package dto

import "time"

type MeetingTypeDTO struct {
	ID      uint    `json:"id"`
	Name  string `json:"name"`
	Color  uint `json:"color"`

	DisplayName string `json:"display_name"`
	CreateUID []interface{} `json:"create_uid"`
	CreateDate      string      `json:"create_date"`
	WriteUID  []interface{} `json:"write_uid"`
	WriteDate      string      `json:"write_date"`
	LastUpdate string `json:"__last_update"`
}

func TimeToString(t time.Time) string{
	return t.Format("2006-01-02 15:04:05")
}