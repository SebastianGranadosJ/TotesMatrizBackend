package models

import (
	"time"
)

type PurchaseOrder struct {
	ID            int            `gorm:"primaryKey;autoIncrement" json:"id"`
	SellerID      int            `gorm:"not null" json:"seller_id"`
	CustomerID    int            `gorm:"not null" json:"customer_id"`
	ResponsibleID int            `json:"responsible_id"`
	DateTime      time.Time      `gorm:"not null" json:"date_time"`
	SubTotal      float64        `gorm:"not null" json:"sub_total"`
	Total         float64        `gorm:"not null" json:"total"`
	OrderStateID  int            `gorm:"not null" json:"order_state_id"`
	Seller        Employee       `gorm:"foreignKey:SellerID;references:ID" json:"seller"`
	Customer      Customer       `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	Responsible   Employee       `gorm:"foreignKey:ResponsibleID;references:ID" json:"responsible"`
	OrderState    OrderStateType `gorm:"foreignKey:OrderStateID;references:ID" json:"order_state"`
}
