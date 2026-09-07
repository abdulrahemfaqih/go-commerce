package main

import (
	"abdulrahemfaqih/go-commerce/config"
	"abdulrahemfaqih/go-commerce/handlers"
	"abdulrahemfaqih/go-commerce/middlewares"
	migrations "abdulrahemfaqih/go-commerce/migrations"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fresh := flag.Bool("fresh", false, "Drop all tables and re-migrate")
	flag.Parse()

	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer db.Close()

	if *fresh {
		migrations.FreshMigrate(db)
		log.Println("Database reset selesai")
		return
	}

	migrations.Migrate(db)

	router := gin.Default()

	// routes crud product categories
	router.GET("/product-categories", handlers.ListProductCategories(db))
	router.GET("/product-categories/:id", handlers.GetProductCategory(db))
	router.POST("/product-categories", middlewares.AuthMiddleware(), handlers.CreateProductCategory(db))
	router.PUT("/product-categories/:id", middlewares.AuthMiddleware(), handlers.UpdateProductCategory(db))
	router.DELETE("/product-categories/:id", middlewares.AuthMiddleware(), handlers.DeleteProductCategory(db))

	// routes crud products
	router.GET("/products", handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", middlewares.AuthMiddleware(), handlers.CreateProduct(db))
	router.PUT("/products/:id", middlewares.AuthMiddleware(), handlers.UpdateProduct(db))
	router.DELETE("/products/:id", middlewares.AuthMiddleware(), handlers.DeleteProduct(db))

	router.POST("/transactions", middlewares.AuthMiddleware(), handlers.CreateTransaction(db))
	router.GET("/transactions/:id", middlewares.AuthMiddleware(), handlers.GetTransactionWithItems(db))

	// routes login dan register
	router.POST("/register", handlers.Register(db))
	router.POST("/login", handlers.Login(db))
	router.POST("/logout", middlewares.AuthMiddleware(), handlers.Logout())

	router.Run(":8080")
}
