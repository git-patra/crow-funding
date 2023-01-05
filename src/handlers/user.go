package handlers

import (
	"CrowFundingV2/src/auth"
	"CrowFundingV2/src/helper"
	"CrowFundingV2/src/modules/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
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
	//** catch error save to db
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Register Account Failed.", "error", err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	token, err := handler.authService.GenerateToken(register.ID)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, "Register Account Failed.", "error", err.Error())
		context.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := user.FormatUser(register, token)
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

	token, err := h.authService.GenerateToken(login.ID)
	if err != nil {
		errorMessage := gin.H{"errors": []string{err.Error()}}
		response := helper.APIResponse(http.StatusBadRequest, "Login Failed.", "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	formatter := user.FormatUser(login, token)
	response := helper.APIResponse(http.StatusOK, "Successfully Login.", "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	//** catch validation dto/input
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Email checking failed.", "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	isEmailExist, err := h.userService.IsEmailAvailable(input)
	//** catch validation get from db
	if err != nil {
		errorMessage := gin.H{"errors": []string{err.Error()}}
		response := helper.APIResponse(http.StatusBadRequest, "Email checking failed.", "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	data := gin.H{
		"is_available": isEmailExist,
	}

	metaMessage := "Email h" +
		"as been registered!"

	if isEmailExist {
		metaMessage = "Email is available."
	}

	response := helper.APIResponse(http.StatusBadRequest, metaMessage, "success", data)
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Failed to upload avatar image.", "error", data)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	// get from JWT Token
	userId := 1
	path := fmt.Sprintf("uploads/images/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(http.StatusUnprocessableEntity, "Failed to upload avatar image.", "error", data)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	_, err = handler.userService.SaveAvatar(userId, path)

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(http.StatusUnprocessableEntity, "Avatar successfuly uploaded.", "error", data)
	c.JSON(http.StatusBadRequest, response)
}
