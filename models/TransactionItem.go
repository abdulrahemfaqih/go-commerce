package models



type TransactionItem struct {

	ID            uint    `gorm:"primary_key" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
}
