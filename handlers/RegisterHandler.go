package handlers

import (
	"abdulrahemfaqih/go-commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		if input.Name == "" || input.Email == "" || input.Password == "" {
			ctx.JSON(400, gin.H{"message": "All fields are required"})
			return
		}

		// cek apakah user sudah by username
		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err == nil {
			ctx.JSON(400, gin.H{"message": "Username already exists"})
			return
		}

		// cek apakah user sudah by email
		if err := db.Where("email = ?", input.Email).First(&user).Error; err == nil {
			ctx.JSON(400, gin.H{"message": "Email already exists"})
			return
		}

		// generate hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Failed to hash password"})
			return
		}

		// buat user
		newUser := models.User{
			Username: input.Username,
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashedPassword),
		}
		if err := db.Create(&newUser).Error; err != nil {
			ctx.JSON(500, gin.H{"message": "Failed to create user"})
			return
		}

		token, err := CreateToken(newUser.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Failed to create token"})
			return
		}

		ctx.JSON(201, gin.H{"message": "User created successfully", "token": token})
	}
}
