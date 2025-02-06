package models

import (
	"time"
)

type Tea struct {
	Id        int       `json:"id" gorm:"primarykey"`
	TeaName   string    `json:"tea_name" gorm:"type:varchar(50)"`
	Type      string    `json:"type" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
