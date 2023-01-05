package main

import (
	"CrowFundingV2/src/auth"
	"CrowFundingV2/src/handlers"
	"CrowFundingV2/src/helper"
	"CrowFundingV2/src/modules/campaign"
	"CrowFundingV2/src/modules/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/crowd_funding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewJWTService()

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService, authService)

	campaignRepo := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepo)
	campaignHandler := handlers.NewCampaignHandler(campaignService, authService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", authMiddleware(authService, userService), campaignHandler.GetListCampaign)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer token
		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(payload["user_id"].(float64))

		userData, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "Unauthorized", "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userData)
	}
}
