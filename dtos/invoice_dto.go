package dtos

import (
	"time"
)

type GetInvoiceDTO struct {
	ID             int              `json:"id"`
	EnterpriseData string           `json:"enterprise_data"`
	DateTime       time.Time        `json:"date_time"`
	CustomerID     int              `json:"customer_id"`
	Total          float64          `json:"total"`
	Subtotal       float64          `json:"subtotal"`
	Items          []BillingItemDTO `json:"items"`
	Discounts      []int            `json:"discounts"`
	Taxes          []int            `json:"taxes"`
}

type CreateInvoiceDTO struct {
	EnterpriseData string           `json:"enterprise_data"`
	CustomerID     int              `json:"customer_id"`
	Items          []BillingItemDTO `json:"items"`
	Discounts      []int            `json:"discounts"`
	Taxes          []int            `json:"taxes"`
}
