package models



type TransactionItem struct {

	ID            uint    `gorm:"primary_key" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	Transaction   Transaction `gorm:"foreignkey:TransactionID"`
	ProductID     uint    `json:"product_id"`
	Product       Product `gorm:"foreignkey:ProductID"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
}
