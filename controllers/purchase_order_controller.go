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
	Log     *utilities.LogUtil
}

func NewPurchaseOrderController(service *services.PurchaseOrderService, auth *utilities.AuthorizationUtil, log *utilities.LogUtil) *PurchaseOrderController {
	return &PurchaseOrderController{Service: service, Auth: auth, Log: log}
}

func (poc *PurchaseOrderController) GetPurchaseOrderByID(c *gin.Context) {

	permissionId := config.PERMISSION_GET_PURCHASE_ORDER_BY_ID

	if err := poc.Log.RegisterLog(c, "Attempting to retrieve Purchase Order by ID"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for GetPurchaseOrderByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	id := c.Param("id")

	purchaseOrder, err := poc.Service.GetPurchaseOrderByID(id)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Purchase Order not found with ID: "+id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Order not found"})
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

	_ = poc.Log.RegisterLog(c, "Successfully retrieved Purchase Order with ID: "+id)

	c.JSON(http.StatusOK, purchaseOrderDTO)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersByStateID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDERS_BY_STATE_ID

	if err := poc.Log.RegisterLog(c, "Attempting to retrieve Purchase Orders by State ID"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for GetPurchaseOrdersByStateID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	stateID := c.Param("stateID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersByStateID(stateID)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Purchase Orders not found for State ID: "+stateID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		_ = poc.Log.RegisterLog(c, "No Purchase Orders found for State ID: "+stateID)
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

	_ = poc.Log.RegisterLog(c, "Successfully retrieved Purchase Orders with State ID: "+stateID)

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) GetAllPurchaseOrders(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_PURCHASE_ORDERS

	if err := poc.Log.RegisterLog(c, "Attempting to retrieve all Purchase Orders"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for GetAllPurchaseOrders")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	purchaseOrders, err := poc.Service.GetAllPurchaseOrders()
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Error retrieving all Purchase Orders")
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

	_ = poc.Log.RegisterLog(c, "Successfully retrieved all Purchase Orders")

	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) SearchPurchaseOrdersByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_PURCHASE_ORDERS_BY_ID

	if err := poc.Log.RegisterLog(c, "Attempting to search Purchase Orders by ID"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for SearchPurchaseOrdersByID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	id := c.Query("id")
	if id == "" {
		_ = poc.Log.RegisterLog(c, "Missing 'id' query parameter in SearchPurchaseOrdersByID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	purchaseOrders, err := poc.Service.SearchPurchaseOrdersByID(id)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Error retrieving Purchase Orders with ID: "+id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		_ = poc.Log.RegisterLog(c, "No Purchase Orders found with ID: "+id)
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

	_ = poc.Log.RegisterLog(c, "Successfully found Purchase Orders with ID containing: "+id)
	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersByCustomerID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDERS_BY_CUSTOMER_ID

	if err := poc.Log.RegisterLog(c, "Attempting to retrieve Purchase Orders by Customer ID"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for GetPurchaseOrdersByCustomerID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	customerID := c.Param("customerID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersByCustomerID(customerID)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Error retrieving Purchase Orders for Customer ID: "+customerID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase Orders not found"})
		return
	}

	if len(purchaseOrders) == 0 {
		_ = poc.Log.RegisterLog(c, "No Purchase Orders found for Customer ID: "+customerID)
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

	_ = poc.Log.RegisterLog(c, "Successfully retrieved Purchase Orders for Customer ID: "+customerID)
	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) GetPurchaseOrdersBySellerID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_PURCHASE_ORDERS_BY_SELLER_ID

	if err := poc.Log.RegisterLog(c, "Attempting to retrieve Purchase Orders by Seller ID"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for GetPurchaseOrdersBySellerID")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	sellerID := c.Param("sellerID")

	purchaseOrders, err := poc.Service.GetPurchaseOrdersBySellerID(sellerID)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Error retrieving Purchase Orders for Seller ID: "+sellerID)
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

	_ = poc.Log.RegisterLog(c, "Successfully retrieved Purchase Orders for Seller ID: "+sellerID)
	c.JSON(http.StatusOK, purchaseOrderDTOs)
}

func (poc *PurchaseOrderController) ChangePurchaseOrderState(c *gin.Context) {
	permissionId := config.PERMISSION_UPDATE_PURCHASE_ORDER_STATE

	if err := poc.Log.RegisterLog(c, "Attempting to update Purchase Order state"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})

		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for UpdatePurchaseOrderState")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})

		return
	}

	id := c.Param("id")

	var request struct {
		OrderStateID int `json:"order_state_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = poc.Log.RegisterLog(c, "Error binding JSON for UpdatePurchaseOrderState: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	orderStateIDStr := strconv.Itoa(request.OrderStateID)
	purchaseOrder, err := poc.Service.ChangePurchaseOrderState(id, orderStateIDStr)

	if err != nil {
		_ = poc.Log.RegisterLog(c, err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}
	fmt.Println(purchaseOrder.OrderStateID)
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

	_ = poc.Log.RegisterLog(c, "Successfully updated Purchase Order state with ID: "+id)
	c.JSON(http.StatusOK, purchaseOrderDTO)
}

func (poc *PurchaseOrderController) CreatePurchaseOrder(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_PURCHASE_ORDER

	if err := poc.Log.RegisterLog(c, "Attempting to create a new Purchase Order"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering log"})
		return
	}

	if !poc.Auth.CheckPermission(c, permissionId) {
		_ = poc.Log.RegisterLog(c, "Permission denied for CreatePurchaseOrder")
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	var dto dtos.CreatePurchaseOrderDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		_ = poc.Log.RegisterLog(c, "Invalid request data for CreatePurchaseOrder: "+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	purchaseOrder, err := poc.Service.CreatePurchaseOrder(&dto)
	if err != nil {
		_ = poc.Log.RegisterLog(c, "Error creating Purchase Order: "+err.Error())
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

	_ = poc.Log.RegisterLog(c, "Successfully created Purchase Order with ID: "+strconv.Itoa(purchaseOrder.ID))
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
