package migrations

import (
	"abdulrahemfaqih/go-commerce/models"
	"log"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
	)

	// add gorm db function
	db.Model(&models.Product{}).AddForeignKey("category_id", "product_categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Transaction{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.TransactionItem{}).AddForeignKey("transaction_id", "transactions(id)", "CASCADE", "CASCADE")
	db.Model(&models.TransactionItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
}

func FreshMigrate(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.DropTableIfExists(
		&models.User{},
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
	)
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
	Migrate(db)
	log.Println("Database migrated successfully")
}
