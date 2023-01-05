package main

import (
	"CrowFundingV2/src/auth"
	"CrowFundingV2/src/handlers"
	"CrowFundingV2/src/modules/user"
	"fmt"
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
	authService := auth.NewJWTService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.GkFDeRlaOVpcx8Ee86ANEEKyxNEqe8lydZPgArTklIvhlchJ10ytOYL43vsTRQy5")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID!")
		fmt.Println("INVALID!")
		fmt.Println("INVALID!")
	}

	userHandler := handlers.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
