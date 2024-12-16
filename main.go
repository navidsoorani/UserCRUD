package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"simpleCrud/controllers"
	"simpleCrud/models"
	"simpleCrud/repository"
)

func main() {
	// Database connection
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local" // Update with your credentials
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Set up Gin and the user controller
	r := gin.Default()
	userRepo := repository.NewGormUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	r.POST("/users", userController.Create)
	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.GetByID)
	r.PUT("/users/:id", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 by default
}
