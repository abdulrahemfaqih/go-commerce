package models


type Transaction struct {

	ID     uint              `gorm:"primary_key" json:"id"`
	UserID uint              `json:"user_id"`
	Amout  float64           `json:"amount"`
	Items  []TransactionItem `gorm:"foreignkey:TransactionID"`
	Price  float64           `json:"price"`
}
