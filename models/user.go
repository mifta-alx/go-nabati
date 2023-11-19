package models

import (
	"time"
)

type User struct {
	Id        int64     `gorm:"primaryKey:autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar" json:"username"`
	Password  string    `gorm:"type:varchar" json:"password"`
	Role      string    `gorm:"type:varchar" json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
