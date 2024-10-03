package entity

type Song struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(200);unique;not null" json:"name"`
	Group      string `gorm:"type:varchar(200);not null" json:"group"`
	Text       string `gorm:"type:text" json:"text"`
	Listenings int    `gorm:"type:int" json:"listenings"`
}
