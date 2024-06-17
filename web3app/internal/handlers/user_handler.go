package handlers

import (
	"SimpleId/internal/model"
	"SimpleId/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (handler *UserHandler) Register(c *gin.Context) {
	var input model.RegisterUserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	results, err := handler.userUseCase.Register(input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "New User had been registered successfully",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) Login(c *gin.Context) {
	var input model.LoginUserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	results, err := handler.userUseCase.Login(input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Login Success",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) GetUserInformation(c *gin.Context) {
	var input model.UserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	results, err := handler.userUseCase.GetInformation(input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Get User Information Success",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) AddUserInformation(c *gin.Context) {
	var input model.UserInformation
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	results, _, err := handler.userUseCase.AddInformation(input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Adding User Information Success",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) UpdateUserInformation(c *gin.Context) {
	var input model.UserInformation
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	results, _, err := handler.userUseCase.UpdateInformation(input)
	if err != nil {
		response := model.Response{
			Success: false,
			Message: "Something Wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := model.Response{
		Success: true,
		Message: "Updating User Information Success",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}
