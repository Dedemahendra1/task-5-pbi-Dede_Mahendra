package models

type Photo struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Title    string `gorm:"type:varchar(300)" json:"title"`
	Caption  string `gorm:"type:varchar(300)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(300)" json:"photo_url"`
}
