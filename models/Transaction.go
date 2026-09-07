package models


type Transaction struct {

	ID     uint              `gorm:"primary_key" json:"id"`
	UserID uint              `json:"user_id"`
	User   *User              `gorm:"foreignkey:UserID" json:"user,omitempty"`
	Amount  float64           `json:"amount"`
	Items  []TransactionItem `gorm:"foreignkey:TransactionID"`
}
