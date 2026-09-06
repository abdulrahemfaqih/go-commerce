package models


type ProductCategory struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
