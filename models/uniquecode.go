package models

type UniqueCode struct {
	Id     int64  `gorm:"primaryKey:autoIncrement" json:"id"`
	Code   string `gorm:"type:varchar" json:"code"`
	Status int64  `gorm:"default:0" json:"status"`
}
