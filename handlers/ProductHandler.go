package handlers

import (
	"abdulrahemfaqih/go-commerce/models"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup
		// menambahkan 1 goroutine ke waitgroup
		wg.Add(1)
		// memulai goroutine untuk melakukan operasi yang membutuhkan waktu lama
		go func() {
			defer wg.Done()
			db.Find(&products)
		}()
		// menunggu goroutine selsai
		wg.Wait()
		ctx.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"error": "product not found"})
			return
		}
		ctx.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&input)
		ctx.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"error": "product not found"})
			return
		}
		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Model(&product).Updates(input)
		ctx.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"error": "product not found"})
			return
		}
		db.Delete(&product)
		ctx.JSON(200, gin.H{"message": "product deleted"})
	}
}