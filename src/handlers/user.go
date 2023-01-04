package handlers

import (
	"CrowFundingV2/src/helper"
	"CrowFundingV2/src/modules/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (handler *userHandler) RegisterUser(context *gin.Context) {
	var userInput user.RegisterUserInput
	err := context.ShouldBindJSON(&userInput)

	//** catch validation dto/input
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(http.StatusUnprocessableEntity, "Register Account Failed.", "error", errorMessage)
		context.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	register, err := handler.userService.Register(userInput)

	//** catch validation save to db
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Register Account Failed.", "error", nil)
		context.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := user.FormatUser(register, "token-token")
	response := helper.APIResponse(http.StatusOK, "Account successfully registered.", "success", formatter)

	context.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	//** catch validation dto/input
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Login Failed.", "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	login, err := h.userService.Login(input)

	//** catch validation get from db
	if err != nil {
		errorMessage := gin.H{"errors": []string{err.Error()}}
		response := helper.APIResponse(http.StatusBadRequest, "Login Failed.", "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := user.FormatUser(login, "token-token")
	response := helper.APIResponse(http.StatusOK, "Successfully Login.", "success", formatter)

	c.JSON(http.StatusOK, response)
}
