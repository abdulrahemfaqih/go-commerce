package models


type Transaction struct {

	ID     uint              `gorm:"primary_key" json:"id"`
	UserID uint              `json:"user_id"`
	User   User              `gorm:"foreignkey:UserID"`
	Amout  float64           `json:"amount"`
	Items  []TransactionItem `gorm:"foreignkey:TransactionID"`
	Price  float64           `json:"price"`
}
