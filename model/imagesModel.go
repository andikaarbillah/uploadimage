package model

import (
	"mime/multipart"
	"time"
)

type Images struct {
	Id         string `json:"id" gorm:"primaryKey;type:varchar(55);index"`
	Path       string `json:"path" gorm:"type:varchar(55);not null"`
	Created_at time.Time	`json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time	`json:"updated_at" gorm:"autoCreateTime"`
}

type ImageRequest struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type ImageResponse struct {
	Id         string `json:"id" gorm:"primaryKey;type:varchar(55);index"`
	Path       string `json:"path" gorm:"type:varchar(55);not null"`
	Created_at time.Time	`json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time	`json:"updated_at" gorm:"autoCreateTime"`
}
