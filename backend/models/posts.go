package models

import (
	"time"
)

type Posts struct {
	ID       uint64    `gorm:"->;column:id"`
	Type     uint16    `gorm:"->;column:type"`
	ResID    uint64    `gorm:"->;column:res_id"`
	Title    string    `gorm:"->;column:title"`
	Text     string    `gorm:"->;column:text"`
	DateUnix uint32    `gorm:"->;column:date"`
	S_Date   time.Time `gorm:"->;column:s_date"`
	Link     string    `gorm:"->;column:link"`
}

func (Posts) TableName() string {
	return "posts"
}
