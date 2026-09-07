package handlers

import (
	"abdulrahemfaqih/go-commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Transaction not found"})
			return
		}

		c.JSON(200, transaction)
	}
}

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Items []struct {
				ProductID uint `json:"product_id"`
				Quantity  int  `json:"quantity"`
			} `json:"items"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		if len(input.Items) == 0 {
			c.JSON(400, gin.H{"message": "Cart cannot be empty"})
			return
		}

		userID := c.MustGet("user_id").(uint) // dari AuthMiddleware

		var totalAmount float64
		var transactionItems []models.TransactionItem

		// Hitung total di Backend secara aman dari harga asli di database
		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(400, gin.H{"message": "Product not found"})
				return
			}

			subtotal := product.Price * float64(item.Quantity)
			totalAmount += subtotal

			// Catat item beserta harga aslinya saat transaksi terjadi
			transactionItems = append(transactionItems, models.TransactionItem{
				ProductID: product.ID,
				Price:     product.Price,
				Quantity:  item.Quantity,
			})
		}

		// Buat transaksi dengan total amount yang dihitung sendiri oleh Backend
		transaction := models.Transaction{
			UserID: userID,
			Amount: totalAmount,
			Items:  transactionItems,
		}

		if err := db.Create(&transaction).Error; err != nil {
			c.JSON(500, gin.H{"message": "Failed to create transaction: " + err.Error()})
			return
		}

		c.JSON(201, transaction)
	}
}
