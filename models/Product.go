package models

type Product struct {
	ID       uint    `gorm:"primary_key" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	Price    float64 `gorm:"not null" json:"price"`
	CategoryID uint    `gorm:"not null" json:"category_id"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID"`
}
