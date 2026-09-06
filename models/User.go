package models



type User struct {

	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
