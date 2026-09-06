package handlers

import (
	"abdulrahemfaqih/go-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		if input.Email == "" || input.Password == "" {
			ctx.JSON(400, gin.H{"message": "All fields are required"})
			return
		}

		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			ctx.JSON(401, gin.H{"message": "Invalid email or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			ctx.JSON(401, gin.H{"message": "Invalid email or password"})
			return
		}

		token, err := CreateToken(user.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Failed to create token"})
			return
		}

		ctx.JSON(200, gin.H{"message": "Login successful", "token": token})
	}
}
