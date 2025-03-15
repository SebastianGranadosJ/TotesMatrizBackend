package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type PurchaseOrderRepository struct {
	DB *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{DB: db}
}

func (r *PurchaseOrderRepository) GetPurchaseOrderByID(id string) (*models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").First(&purchaseOrder, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *PurchaseOrderRepository) GetPurchaseOrdersByCustomerID(customerID string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		Where("customer_id = ?", customerID).Find(&purchaseOrders).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) GetPurchaseOrdersBySellerID(sellerID string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		Where("seller_id = ?", sellerID).Find(&purchaseOrders).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) GetAllPurchaseOrders() ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").Find(&purchaseOrders).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) SearchPurchaseOrdersByID(query string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").Find(&purchaseOrders).Error
	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) UpdatePurchaseOrderState(id string, state int) (*models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	if err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		First(&purchaseOrder, "id = ?", id).Error; err != nil {
		return nil, err
	}

	purchaseOrder.OrderState.ID = state

	if err := r.DB.Save(&purchaseOrder).Error; err != nil {
		return nil, err
	}
	return &purchaseOrder, nil

}

func (r *PurchaseOrderRepository) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrder) error {
	var existingPurchaseOrder models.PurchaseOrder
	if err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		First(&existingPurchaseOrder, "id = ?", purchaseOrder.ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *PurchaseOrderRepository) CreatePurchaseOrder(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	if err := r.DB.Preload("Seller").Preload("Responsible").Preload("Customer").Preload("OrderState").
		Create(purchaseOrder).Error; err != nil {
		return nil, err
	}
	return purchaseOrder, nil
}
