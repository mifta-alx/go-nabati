package models

type Redemption struct {
	Id           int64      `grom:"primaryKey:autoIncrement" json:"id"`
	Name         string     `grom:"type:varchar" json:"name"`
	Email        string     `grom:"type:varchar" json:"email"`
	UniqueCodeId int64      `json:"uniquecode_id"`
	UniqueCode   UniqueCode `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
