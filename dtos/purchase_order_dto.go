package dtos

import "time"

type GetPurchaseOrderDTO struct {
	ID            int       `json:"id"`
	SellerID      int       `json:"seller_id"`
	CustomerID    int       `json:"customer_id"`
	ResponsibleID int       `json:"responsible_id"`
	DateTime      time.Time `json:"date_time"`
	SubTotal      float64   `json:"sub_total"`
	Total         float64   `json:"total"`
	OrderStateID  int       `json:"order_state_id"`
}

type UpdatePurchaseOrderDTO struct {
	SellerID      int       `json:"seller_id"`
	CustomerID    int       `json:"customer_id"`
	ResponsibleID int       `json:"responsible_id"`
	DateTime      time.Time `json:"date_time"`
	SubTotal      float64   `json:"sub_total"`
	Total         float64   `json:"total"`
	OrderStateID  int       `json:"order_state_id"`
}

type CreatePurchaseOrderDTO struct {
	SellerID      int       `json:"seller_id"`
	CustomerID    int       `json:"customer_id"`
	ResponsibleID int       `json:"responsible_id"`
	DateTime      time.Time `json:"date_time"`
	SubTotal      float64   `json:"sub_total"`
	Total         float64   `json:"total"`
	OrderStateID  int       `json:"order_state_id"`
}
