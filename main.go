package main

import (
	"abdulrahemfaqih/go-commerce/config"
	"abdulrahemfaqih/go-commerce/handlers"
	"abdulrahemfaqih/go-commerce/middlewares"
	migrations "abdulrahemfaqih/go-commerce/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer db.Close()

	migrations.Migrate(db)

	// routes crud products
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.GET("/products/:id", middlewares.AuthMiddleware(), handlers.GetProduct(db))
	router.POST("/products", middlewares.AuthMiddleware(), handlers.CreateProduct(db))
	router.PUT("/products/:id", middlewares.AuthMiddleware(), handlers.UpdateProduct(db))
	router.DELETE("/products/:id", middlewares.AuthMiddleware(), handlers.DeleteProduct(db))

	// routes login dan register
	router.POST("/register", handlers.Register(db))
	router.POST("/login", handlers.Login(db))

	router.Run(":8080")
}
