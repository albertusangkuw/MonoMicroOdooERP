package model

import (
	"errors"
	"service-10/utils"
)

const maxWidth = 128
const maxHeight = 128

type PosCategoryImage struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	ImageBase64   string `gorm:"type:text"`
	AttachmentId  int
	ImageSizeByte int
}

func (p *PosCategoryImage) FitImage() error {
	p.ImageBase64 = utils.ResizeImage(p.ImageBase64, maxWidth, maxHeight)
	if p.ImageBase64 == "" {
		return errors.New("unable to Fit Image")
	}
	return nil
}
