package dto

import (
	"fmt"
	"time"
)

type JWTData struct {
	Uid uint `json:"uid"`
	UserContext  struct {
		Lang string `json:"lang"`
		Tz string `json:"tz"`
	} `json:"user_context"`
	DB string `json:"db"`
	CompanyID int `json:"company_id"`
	PartnerID int `json:"partner_id"`
	GroupID []int `json:"group_id"`
	Exp int64 `json:"exp"`

}

func (c *JWTData) Valid() error {
	now := time.Now().Unix()
	if now > c.Exp {
		return fmt.Errorf("token had been expired")
	}
	return nil
}

