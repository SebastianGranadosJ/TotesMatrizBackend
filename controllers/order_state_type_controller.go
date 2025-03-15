package controllers

import (
	"fmt"
	"net/http"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type OrderStateTypeController struct {
	Service *services.OrderStateTypeService
}

func NewOrderStateTypeController(service *services.OrderStateTypeService) *OrderStateTypeController {
	return &OrderStateTypeController{Service: service}
}

func (ostc *OrderStateTypeController) GetOrderStateTypeByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	orderStateType, err := ostc.Service.GetOrderStateTypeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order State Type not found"})
		return
	}

	c.JSON(http.StatusOK, orderStateType)
}

func (ostc *OrderStateTypeController) GetOrderStateTypes(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	orderStateTypes, err := ostc.Service.GetAllOrderStateTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Order State Types"})
		return
	}
	c.JSON(http.StatusOK, orderStateTypes)

}
