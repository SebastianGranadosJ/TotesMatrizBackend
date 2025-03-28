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

type InvoiceController struct {
	Service *services.InvoiceService
	Auth    *utilities.AuthorizationUtil
}

func NewInvoiceController(service *services.InvoiceService, auth *utilities.AuthorizationUtil) *InvoiceController {
	return &InvoiceController{Service: service, Auth: auth}
}

func (ic *InvoiceController) GetAllInvoices(c *gin.Context) {
	permissionId := config.PERMISSION_GET_ALL_INVOICES

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	invoices, err := ic.Service.GetAllInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve invoices"})
		return
	}

	var invoiceDTOs []dtos.GetInvoiceDTO
	for _, invoice := range invoices {
		invoiceDTOs = append(invoiceDTOs, dtos.GetInvoiceDTO{
			ID:             invoice.ID,
			EnterpriseData: invoice.EnterpriseData,
			DateTime:       invoice.DateTime,
			CustomerID:     invoice.CustomerID,
			Subtotal:       invoice.Subtotal,
			Total:          invoice.Total,
			Items:          extractInvoiceBillingItems(invoice.Items),
			Discounts:      extractDiscountIds(invoice.Discounts),
			Taxes:          extractTaxIds(invoice.Taxes),
		})
	}

	c.JSON(http.StatusOK, invoiceDTOs)
}

func (ic *InvoiceController) GetInvoiceByID(c *gin.Context) {
	permissionId := config.PERMISSION_GET_INVOICE_BY_ID

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	invoice, err := ic.Service.GetInvoiceByID(strconv.Itoa(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	invoiceDTO := dtos.GetInvoiceDTO{
		ID:             invoice.ID,
		EnterpriseData: invoice.EnterpriseData,
		DateTime:       invoice.DateTime,
		CustomerID:     invoice.CustomerID,
		Subtotal:       invoice.Subtotal,
		Total:          invoice.Total,
		Items:          extractInvoiceBillingItems(invoice.Items),
		Discounts:      extractDiscountIds(invoice.Discounts),
		Taxes:          extractTaxIds(invoice.Taxes),
	}
	fmt.Print(invoice)

	c.JSON(http.StatusOK, invoiceDTO)
}

func (ic *InvoiceController) SearchInvoiceByID(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_INVOICE_BY_ID

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("id")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	invoices, err := ic.Service.SearchInvoiceByID(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching invoices"})
		return
	}

	var invoiceDTOs []dtos.GetInvoiceDTO
	for _, invoice := range invoices {
		invoiceDTOs = append(invoiceDTOs, dtos.GetInvoiceDTO{
			ID:             invoice.ID,
			EnterpriseData: invoice.EnterpriseData,
			DateTime:       invoice.DateTime,
			CustomerID:     invoice.CustomerID,
			Subtotal:       invoice.Subtotal,
			Total:          invoice.Total,
			Items:          extractInvoiceBillingItems(invoice.Items),
			Discounts:      extractDiscountIds(invoice.Discounts),
			Taxes:          extractTaxIds(invoice.Taxes),
		})
	}

	c.JSON(http.StatusOK, invoiceDTOs)
}

func (ic *InvoiceController) SearchInvoiceByCustomerPersonalId(c *gin.Context) {
	permissionId := config.PERMISSION_SEARCH_INVOICE_BY_CUSTOMER_PERSONAL_ID

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	query := c.Query("personal_id")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'personal_id' is required"})
		return
	}

	invoices, err := ic.Service.SearchInvoiceByCustomerPersonalId(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching invoices by customer personal ID"})
		return
	}

	var invoiceDTOs []dtos.GetInvoiceDTO
	for _, invoice := range invoices {
		invoiceDTOs = append(invoiceDTOs, dtos.GetInvoiceDTO{
			ID:             invoice.ID,
			EnterpriseData: invoice.EnterpriseData,
			DateTime:       invoice.DateTime,
			CustomerID:     invoice.CustomerID,
			Subtotal:       invoice.Subtotal,
			Total:          invoice.Total,
			Items:          extractInvoiceBillingItems(invoice.Items),
			Discounts:      extractDiscountIds(invoice.Discounts),
			Taxes:          extractTaxIds(invoice.Taxes),
		})
	}

	c.JSON(http.StatusOK, invoiceDTOs)
}

func (ic *InvoiceController) CreateInvoice(c *gin.Context) {
	permissionId := config.PERMISSION_CREATE_INVOICE

	if !ic.Auth.CheckPermission(c, permissionId) {
		return
	}

	var dto dtos.CreateInvoiceDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	invoice, err := ic.Service.CreateInvoice(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	invoiceDTO := dtos.GetInvoiceDTO{
		ID:             invoice.ID,
		EnterpriseData: invoice.EnterpriseData,
		DateTime:       invoice.DateTime,
		CustomerID:     invoice.CustomerID,
		Subtotal:       invoice.Subtotal,
		Total:          invoice.Total,
		Items:          extractInvoiceBillingItems(invoice.Items),
		Discounts:      extractDiscountIds(invoice.Discounts),
		Taxes:          extractTaxIds(invoice.Taxes),
	}

	// Responder con la factura creada
	c.JSON(http.StatusCreated, invoiceDTO)
}

func extractInvoiceBillingItems(items []models.InvoiceItem) []dtos.BillingItemDTO {
	var billingItems []dtos.BillingItemDTO
	for _, item := range items {
		billingItems = append(billingItems, dtos.BillingItemDTO{
			ID:    item.ItemID,
			Stock: item.Amount,
		})
	}
	return billingItems
}

func extractDiscountIds(discounts []models.DiscountType) []int {
	var ids []int
	for _, discount := range discounts {
		ids = append(ids, discount.ID)
	}
	return ids
}

func extractTaxIds(taxes []models.TaxType) []int {
	var ids []int
	for _, tax := range taxes {
		ids = append(ids, tax.ID)
	}
	return ids
}
