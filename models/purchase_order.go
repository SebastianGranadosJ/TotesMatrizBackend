package models

import (
	"time"
)

type PurchaseOrder struct {
	ID            int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	SellerID      int                 `gorm:"not null" json:"seller_id"`
	Seller        Employee            `gorm:"foreignKey:SellerID;references:ID" json:"seller"`
	CustomerID    int                 `gorm:"not null" json:"customer_id"`
	Customer      Customer            `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	ResponsibleID int                 `json:"responsible_id"`
	Responsible   Employee            `gorm:"foreignKey:ResponsibleID;references:ID" json:"responsible"`
	DateTime      time.Time           `json:"date_time" time_format:"2006-01-02T15:04:05"`
	Items         []PurchaseOrderItem `gorm:"foreignKey:PurchaseOrderID" json:"items"`
	SubTotal      float64             `gorm:"not null" json:"sub_total"`
	OrderStateID  int                 `gorm:"not null" json:"order_state_id"`
	OrderState    OrderStateType      `gorm:"foreignKey:OrderStateID;references:ID" json:"order_state"`
	Discounts     []DiscountType      `gorm:"many2many:purchase_order_discounts;" json:"discounts"`
	Taxes         []TaxType           `gorm:"many2many:purchase_order_taxes;" json:"taxes"`
	Total         float64             `gorm:"not null" json:"total"`
}

type PurchaseOrderItem struct {
	PurchaseOrderID int `gorm:"primaryKey"`
	ItemID          int `gorm:"primaryKey"`
	PurchaseOrder   PurchaseOrder
	Item            Item
	Amount          int `gorm:"not null"`
}
