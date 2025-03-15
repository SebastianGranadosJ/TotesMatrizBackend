package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PurchaseOrderController struct {
	Service *services.PurchaseOrderService
}

func NewPurchaseOrderController(service *services.PurchaseOrderService) *PurchaseOrderController {
	return &PurchaseOrderController{Service: service}
}

func (poc *PurchaseOrderController) GetPurchaseOrderByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	query := c.Param("id")

	purchaseOrder, err := poc.Service.GetPurchaseOrderByID(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	purchaseOrderDTO := dtos.GetPurchaseOrderDTO{
		ID:            purchaseOrder.ID,
		SellerID:      purchaseOrder.SellerID,
		CustomerID:    purchaseOrder.CustomerID,
		ResponsibleID: purchaseOrder.ResponsibleID,
		DateTime:      purchaseOrder.DateTime,
		SubTotal:      purchaseOrder.SubTotal,
		Total:         purchaseOrder.Total,
		OrderStateID:  purchaseOrder.OrderStateID,
	}

	c.JSON(http.StatusOK, purchaseOrderDTO)
}

func (poc *PurchaseOrderController) GetAllPurchaseOrders(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	purchaseOrders, err := poc.Service.GetAllPurchaseOrders()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	var purchaseOrdersDTO []dtos.GetPurchaseOrderDTO
	for _, order := range purchaseOrders {
		orderDTO := dtos.GetPurchaseOrderDTO{
			ID:            order.ID,
			SellerID:      order.SellerID,
			CustomerID:    order.CustomerID,
			ResponsibleID: order.ResponsibleID,
			DateTime:      order.DateTime,
			SubTotal:      order.SubTotal,
			Total:         order.Total,
			OrderStateID:  order.OrderStateID,
		}
		purchaseOrdersDTO = append(purchaseOrdersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, purchaseOrdersDTO)
}

func (poc *PurchaseOrderController) SearchPurchaseOrdersByID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	purchaseOrders, err := poc.Service.SearchPurchaseOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No purchase orders found"})
		return
	}

	var purchaseOrdersDTO []dtos.GetPurchaseOrderDTO
	for _, order := range purchaseOrders {
		orderDTO := dtos.GetPurchaseOrderDTO{
			ID:            order.ID,
			SellerID:      order.SellerID,
			CustomerID:    order.CustomerID,
			ResponsibleID: order.ResponsibleID,
			DateTime:      order.DateTime,
			SubTotal:      order.SubTotal,
			Total:         order.Total,
			OrderStateID:  order.OrderStateID,
		}
		purchaseOrdersDTO = append(purchaseOrdersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, purchaseOrdersDTO)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersByCustomerID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	customerID := c.Param("customer_id")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No purchase orders found"})
		return
	}

	var purchaseOrdersDTO []dtos.GetPurchaseOrderDTO
	for _, order := range purchaseOrders {
		orderDTO := dtos.GetPurchaseOrderDTO{
			ID:            order.ID,
			SellerID:      order.SellerID,
			CustomerID:    order.CustomerID,
			ResponsibleID: order.ResponsibleID,
			DateTime:      order.DateTime,
			SubTotal:      order.SubTotal,
			Total:         order.Total,
			OrderStateID:  order.OrderStateID,
		}
		purchaseOrdersDTO = append(purchaseOrdersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, purchaseOrdersDTO)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersBySellerID(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	sellerID := c.Param("sellerID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersBySellerID(sellerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	var purchaseOrdersDTO []dtos.GetPurchaseOrderDTO
	for _, order := range purchaseOrders {
		orderDTO := dtos.GetPurchaseOrderDTO{
			ID:            order.ID,
			SellerID:      order.SellerID,
			CustomerID:    order.CustomerID,
			ResponsibleID: order.ResponsibleID,
			DateTime:      order.DateTime,
			SubTotal:      order.SubTotal,
			Total:         order.Total,
			OrderStateID:  order.OrderStateID,
		}
		purchaseOrdersDTO = append(purchaseOrdersDTO, orderDTO)
	}

	c.JSON(http.StatusOK, purchaseOrdersDTO)
}

func (poc *PurchaseOrderController) UpdatePurchaseOrderState(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	var request struct {
		OrderStateID int `json:"order_state_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchaseOrder, err := poc.Service.UpdatePurchaseOrderState(id, request.OrderStateID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
		return
	}

	orderDTO := dtos.GetPurchaseOrderDTO{
		ID:            purchaseOrder.ID,
		SellerID:      purchaseOrder.SellerID,
		CustomerID:    purchaseOrder.CustomerID,
		ResponsibleID: purchaseOrder.ResponsibleID,
		DateTime:      purchaseOrder.DateTime,
		SubTotal:      purchaseOrder.SubTotal,
		Total:         purchaseOrder.Total,
		OrderStateID:  purchaseOrder.OrderStateID,
	}

	c.JSON(http.StatusOK, orderDTO)
}

func (poc *PurchaseOrderController) UpdatePurchaseOrder(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	id := c.Param("id")

	var dto dtos.UpdatePurchaseOrderDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchaseOrder, err := poc.Service.GetPurchaseOrderByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	purchaseOrder.SellerID = dto.SellerID
	purchaseOrder.CustomerID = dto.CustomerID
	purchaseOrder.ResponsibleID = dto.ResponsibleID
	purchaseOrder.DateTime = dto.DateTime
	purchaseOrder.SubTotal = dto.SubTotal
	purchaseOrder.Total = dto.Total
	purchaseOrder.OrderStateID = dto.OrderStateID

	err = poc.Service.UpdatePurchaseOrder(purchaseOrder)
	var dtoPurchaseOrder dtos.GetPurchaseOrderDTO

	dtoPurchaseOrder.ID = purchaseOrder.ID
	dtoPurchaseOrder.SellerID = purchaseOrder.SellerID
	dtoPurchaseOrder.CustomerID = purchaseOrder.CustomerID
	dtoPurchaseOrder.ResponsibleID = purchaseOrder.ResponsibleID
	dtoPurchaseOrder.DateTime = purchaseOrder.DateTime
	dtoPurchaseOrder.SubTotal = purchaseOrder.SubTotal
	dtoPurchaseOrder.Total = purchaseOrder.Total
	dtoPurchaseOrder.OrderStateID = purchaseOrder.OrderStateID

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, dtoPurchaseOrder)
}

func (poc *PurchaseOrderController) CreatePurchaseOrder(c *gin.Context) {
	username := c.GetHeader("Username")
	fmt.Println("Request made by user:", username)

	var dto dtos.CreatePurchaseOrderDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder := models.PurchaseOrder{
		SellerID:      dto.SellerID,
		CustomerID:    dto.CustomerID,
		ResponsibleID: dto.ResponsibleID,
		DateTime:      dto.DateTime,
		SubTotal:      dto.SubTotal,
		Total:         dto.Total,
		OrderStateID:  dto.OrderStateID,
	}

	createdOrder, err := poc.Service.CreatePurchaseOrder(&newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase order"})
		return
	}

	orderDTO := dtos.GetPurchaseOrderDTO{
		ID:            createdOrder.ID,
		SellerID:      createdOrder.SellerID,
		CustomerID:    createdOrder.CustomerID,
		ResponsibleID: createdOrder.ResponsibleID,
		DateTime:      createdOrder.DateTime,
		SubTotal:      createdOrder.SubTotal,
		Total:         createdOrder.Total,
		OrderStateID:  createdOrder.OrderStateID,
	}

	c.JSON(http.StatusCreated, orderDTO)
}
