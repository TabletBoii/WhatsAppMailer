package models

import "time"

func (ProjectItems) TableName() string {
	return "project_items"
}

type ProjectItems struct {
	ID         uint32    `gorm:"->;column:id"`
	ProjectID  uint32    `gorm:"->;column:project_id"`
	ItemID     string    `gorm:"->;column:item_id"`
	Type       uint32    `gorm:"->;column:type"`
	ResID      uint32    `gorm:"->;column:res_id"`
	CountryID  uint32    `gorm:"->;column:country_id"`
	RegionID   uint32    `gorm:"->;column:region_id"`
	CategoryID uint32    `gorm:"->;column:category_id"`
	Lang       uint32    `gorm:"->;column:lang"`
	Date       time.Time `gorm:"->;column:date"`
	S_Date     time.Time `gorm:"->;column:s_date"`
	Sex        int32     `gorm:"->;column:sex"`
	StartDate  time.Time `gorm:"->;column:start_date"`
	CreatedAt  time.Time `gorm:"->;column:created_at"`
	UserID     int32     `gorm:"->;column:user_id"`
	Sentiment  int8      `gorm:"->;column:sentiment"`
	DeletedAt  time.Time `gorm:"->;column:deleted_at"`
	DeletedBy  int32     `gorm:"->;column:deleted_by"`
}
