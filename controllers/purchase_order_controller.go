package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"totesbackend/config"
	"totesbackend/controllers/utilities"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/services"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderController struct {
	Service *services.PurchaseOrderService
	Auth    *utilities.AuthorizationUtil
}

func NewPurchaseOrderController(service *services.PurchaseOrderService, auth *utilities.AuthorizationUtil) *PurchaseOrderController {
	return &PurchaseOrderController{Service: service, Auth: auth}
}

func (poc *PurchaseOrderController) GetPurchaseOrderByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDER_BY_ID

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchaseOrder ID"})
		return
	}

	purchaseOrder, err := poc.Service.GetPurchaseOrderByID(strconv.Itoa(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "purchaseOrder not found"})
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
		Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
		Discounts:     extractDiscountIds(purchaseOrder.Discounts),
		Taxes:         extractTaxIds(purchaseOrder.Taxes),
	}
	fmt.Print(purchaseOrder)

	c.JSON(http.StatusOK, purchaseOrderDTO)
}

func (poc *PurchaseOrderController) GetAllPurchaseOrders(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_PURCHASE_ORDERS

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

	purchaseOrders, err := poc.Service.GetAllPurchaseOrders()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	var purchaseOrderDTOs []dtos.GetPurchaseOrderDTO
	for _, purchaseOrder := range purchaseOrders {
		purchaseOrderDTOs = append(purchaseOrderDTOs, dtos.GetPurchaseOrderDTO{
			ID:            purchaseOrder.ID,
			SellerID:      purchaseOrder.SellerID,
			CustomerID:    purchaseOrder.CustomerID,
			ResponsibleID: purchaseOrder.ResponsibleID,
			DateTime:      purchaseOrder.DateTime,
			SubTotal:      purchaseOrder.SubTotal,
			Total:         purchaseOrder.Total,
			OrderStateID:  purchaseOrder.OrderStateID,
			Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
			Discounts:     extractDiscountIds(purchaseOrder.Discounts),
			Taxes:         extractTaxIds(purchaseOrder.Taxes),
		})
	}

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) SearchPurchaseOrdersByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_PURCHASE_ORDERS_BY_ID

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

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

	var purchaseOrderDTOs []dtos.GetPurchaseOrderDTO
	for _, purchaseOrder := range purchaseOrders {
		purchaseOrderDTOs = append(purchaseOrderDTOs, dtos.GetPurchaseOrderDTO{
			ID:            purchaseOrder.ID,
			SellerID:      purchaseOrder.SellerID,
			CustomerID:    purchaseOrder.CustomerID,
			ResponsibleID: purchaseOrder.ResponsibleID,
			DateTime:      purchaseOrder.DateTime,
			SubTotal:      purchaseOrder.SubTotal,
			Total:         purchaseOrder.Total,
			OrderStateID:  purchaseOrder.OrderStateID,
			Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
			Discounts:     extractDiscountIds(purchaseOrder.Discounts),
			Taxes:         extractTaxIds(purchaseOrder.Taxes),
		})
	}

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersByCustomerID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDERS_BY_CUSTOMER_ID

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

	customerID := c.Param("customerID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No purchase orders found"})
		return
	}

	var purchaseOrderDTOs []dtos.GetPurchaseOrderDTO
	for _, purchaseOrder := range purchaseOrders {
		purchaseOrderDTOs = append(purchaseOrderDTOs, dtos.GetPurchaseOrderDTO{
			ID:            purchaseOrder.ID,
			SellerID:      purchaseOrder.SellerID,
			CustomerID:    purchaseOrder.CustomerID,
			ResponsibleID: purchaseOrder.ResponsibleID,
			DateTime:      purchaseOrder.DateTime,
			SubTotal:      purchaseOrder.SubTotal,
			Total:         purchaseOrder.Total,
			OrderStateID:  purchaseOrder.OrderStateID,
			Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
			Discounts:     extractDiscountIds(purchaseOrder.Discounts),
			Taxes:         extractTaxIds(purchaseOrder.Taxes),
		})
	}

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersBySellerID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDERS_BY_SELLER_ID

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

	sellerID := c.Param("sellerID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersBySellerID(sellerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	var purchaseOrderDTOs []dtos.GetPurchaseOrderDTO
	for _, purchaseOrder := range purchaseOrders {
		purchaseOrderDTOs = append(purchaseOrderDTOs, dtos.GetPurchaseOrderDTO{
			ID:            purchaseOrder.ID,
			SellerID:      purchaseOrder.SellerID,
			CustomerID:    purchaseOrder.CustomerID,
			ResponsibleID: purchaseOrder.ResponsibleID,
			DateTime:      purchaseOrder.DateTime,
			SubTotal:      purchaseOrder.SubTotal,
			Total:         purchaseOrder.Total,
			OrderStateID:  purchaseOrder.OrderStateID,
			Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
			Discounts:     extractDiscountIds(purchaseOrder.Discounts),
			Taxes:         extractTaxIds(purchaseOrder.Taxes),
		})
	}

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) UpdatePurchaseOrderState(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_PURCHASE_ORDER_STATE

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

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

	purchaseOrderDTO := dtos.GetPurchaseOrderDTO{
		ID:            purchaseOrder.ID,
		SellerID:      purchaseOrder.SellerID,
		CustomerID:    purchaseOrder.CustomerID,
		ResponsibleID: purchaseOrder.ResponsibleID,
		DateTime:      purchaseOrder.DateTime,
		SubTotal:      purchaseOrder.SubTotal,
		Total:         purchaseOrder.Total,
		OrderStateID:  purchaseOrder.OrderStateID,
		Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
		Discounts:     extractDiscountIds(purchaseOrder.Discounts),
		Taxes:         extractTaxIds(purchaseOrder.Taxes),
	}

	c.JSON(http.StatusOK, purchaseOrderDTO)
}

/*func (poc *PurchaseOrderController) UpdatePurchaseOrder(c *gin.Context) {
    permissionId := config.PERMISSION_UPDATE_PURCHASE_ORDER

    if !poc.Auth.CheckPermission(c, permissionId) {
        return
    }

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
    purchaseOrder.Items = extractPurchaseOrderBillingItems(dto.Items)
    purchaseOrder.Discounts = dto.Discounts
    purchaseOrder.Taxes = dto.Taxes

    updatedOrder, err := poc.Service.UpdatePurchaseOrder(purchaseOrder)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    var response dtos.GetPurchaseOrderDTO
    response.ID = updatedOrder.ID
    response.SellerID = updatedOrder.SellerID
    response.CustomerID = updatedOrder.CustomerID
    response.ResponsibleID = updatedOrder.ResponsibleID
    response.DateTime = updatedOrder.DateTime
    response.Items = extractBillingItems(updatedOrder.Items)
    response.Discounts = updatedOrder.Discounts
    response.Taxes = updatedOrder.Taxes

    c.JSON(http.StatusOK, response)
}*/

func (poc *PurchaseOrderController) CreatePurchaseOrder(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_PURCHASE_ORDER

	if !poc.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.CreatePurchaseOrderDTO
	//body, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println("JSON recibido:", string(body))

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	purchaseOrder, err := poc.Service.CreatePurchaseOrder(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		Items:         extractPurchaseOrderBillingItems(purchaseOrder.Items),
		Discounts:     extractDiscountIds(purchaseOrder.Discounts),
		Taxes:         extractTaxIds(purchaseOrder.Taxes),
	}

	c.JSON(http.StatusCreated, purchaseOrderDTO)
}

func extractPurchaseOrderBillingItems(items []models.PurchaseOrderItem) []dtos.BillingItemDTO {
	var billingItems []dtos.BillingItemDTO
	for _, item := range items {
		billingItems = append(billingItems, dtos.BillingItemDTO{
			ID:    item.ItemID,
			Stock: item.Amount,
		})
	}
	return billingItems
}
