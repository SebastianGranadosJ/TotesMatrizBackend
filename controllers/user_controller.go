package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"totesbackend/dtos"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	user, err := uc.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userDTO := dtos.GetUserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Token:     user.Token,
		UserType:  user.UserTypeID,
		UserState: user.UserStateTypeID,
	}

	c.JSON(http.StatusOK, userDTO)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	users, err := uc.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:        user.ID,
			Email:     user.Email,
			Password:  user.Password,
			Token:     user.Token,
			UserType:  user.UserTypeID,
			UserState: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}

	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) SearchUsersByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	users, err := uc.Service.SearchUsersByID(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:        user.ID,
			Email:     user.Email,
			Password:  user.Password,
			Token:     user.Token,
			UserType:  user.UserTypeID,
			UserState: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}
	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) SearchUsersByEmail(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	users, err := uc.Service.SearchUsersByEmail(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	var usersDTO []dtos.GetUserDTO
	for _, user := range users {

		userDTO := dtos.GetUserDTO{
			ID:        user.ID,
			Email:     user.Email,
			Password:  user.Password,
			Token:     user.Token,
			UserType:  user.UserTypeID,
			UserState: user.UserStateTypeID,
		}

		usersDTO = append(usersDTO, userDTO)
	}
	c.JSON(http.StatusOK, usersDTO)
}

func (uc *UserController) UpdateUserState(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	var request struct {
		UserState int `json:"user_state"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.Service.UpdateUserState(id, request.UserState)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userDTO := dtos.GetUserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Token:     user.Token,
		UserType:  user.UserTypeID,
		UserState: user.UserStateTypeID,
	}

	c.JSON(http.StatusOK, userDTO)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	var dto dtos.GetUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.Service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user.Email = dto.Email
	user.Password = dto.Password
	user.Token = dto.Token
	user.UserTypeID = dto.UserType
	user.UserStateTypeID = dto.UserState

	err = uc.Service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, dto)
}
