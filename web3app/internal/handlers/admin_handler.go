package handlers

import (
	"SimpleId/internal/model"
	"SimpleId/internal/model/converter"
	"SimpleId/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userUseCase    usecases.UserUseCase
	sharingUseCase usecases.SharingUseCase
}

func NewAdminHandler(userUseCase usecases.UserUseCase, sharingUseCase usecases.SharingUseCase) *AdminHandler {
	return &AdminHandler{
		userUseCase:    userUseCase,
		sharingUseCase: sharingUseCase,
	}
}

func (handler *AdminHandler) ApproveUser(c *gin.Context) {
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

	results, err := handler.userUseCase.Approve(input)
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
		Message: "New User had been approved successfully",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) ApproveSharing(c *gin.Context) {
	var input model.ApproveSharingRequest
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

	contracAddress, request := converter.ApproveSharingRequestToString(input)

	results, err := handler.sharingUseCase.ApproveSharing(contracAddress, request)
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
		Message: "New User had been approved successfully",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) GetSharingRequest(c *gin.Context) {
	var input model.GetSharingRequest

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

	results, err := handler.sharingUseCase.GetRequest(input)
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
		Message: "",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) ShowAllSharing(c *gin.Context) {
	results, err := handler.sharingUseCase.ShowAll()
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
		Message: "",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) ShowAllSharingByStatus(c *gin.Context) {
	var input model.StatusFilterRequest

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

	results, err := handler.sharingUseCase.ShowAllByStatus(input.Status)
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
		Message: "",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) RequestSharing(c *gin.Context) {
	var input model.SharingRequest

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

	results, err := handler.sharingUseCase.RequestSharing(input)
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
		Message: "",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}

func (handler *AdminHandler) ReceiveSharing(c *gin.Context) {
	var input model.ReceiveSharingRequest

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

	results, err := handler.sharingUseCase.GetSharingInformation(input.Secret, input.ContractAddress)
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
		Message: "New User had been approved successfully",
		Data:    results,
	}
	c.JSON(http.StatusOK, response)
}
