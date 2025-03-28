package dtos

import "time"

type GetPurchaseOrderDTO struct {
	ID            int              `json:"id"`
	DateTime      time.Time        `json:"date_time"`
	SellerID      int              `json:"seller_id"`
	CustomerID    int              `json:"customer_id"`
	ResponsibleID int              `json:"responsible_id"`
	SubTotal      float64          `json:"sub_total"`
	Total         float64          `json:"total"`
	OrderStateID  int              `json:"order_state_id"`
	Items         []BillingItemDTO `json:"items"`
	Discounts     []int            `json:"discounts"`
	Taxes         []int            `json:"taxes"`
}

type CreatePurchaseOrderDTO struct {
	SellerID      int              `json:"seller_id"`
	CustomerID    int              `json:"customer_id"`
	ResponsibleID int              `json:"responsible_id"`
	Items         []BillingItemDTO `json:"items"`
	Discounts     []int            `json:"discounts"`
	Taxes         []int            `json:"taxes"`
}

type UpdatePurchaseOrderDTO struct {
	SellerID      int              `json:"seller_id"`
	CustomerID    int              `json:"customer_id"`
	ResponsibleID int              `json:"responsible_id"`
	DateTime      time.Time        `json:"date_time"`
	Items         []BillingItemDTO `json:"items"`
	Discounts     []int            `json:"discounts"`
	Taxes         []int            `json:"taxes"`
}
