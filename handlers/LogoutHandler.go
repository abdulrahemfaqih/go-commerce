package handlers

import (
	"github.com/gin-gonic/gin"
)

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Mengambil user_id yang telah diset oleh AuthMiddleware
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Logout successful",
			"user_id": userID,
		})
	}
}
