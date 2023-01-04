package main

import (
	"CrowFundingV2/src/handlers"
	"CrowFundingV2/src/modules/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/crowd_funding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
