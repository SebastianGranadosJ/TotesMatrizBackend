package repositories

import (
	"errors"
	"strconv"
	"time"
	"totesbackend/dtos"
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
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts"). // Ahora sí debería funcionar
		Preload("Taxes").
		First(&purchaseOrder, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &purchaseOrder, nil
}

func (r *PurchaseOrderRepository) GetPurchaseOrdersByStateID(stateID string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Where("order_state_id = ?", stateID).
		Find(&purchaseOrders).Error

	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) GetPurchaseOrdersByCustomerID(customerID string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Where("CAST(customer_id AS TEXT) = ?", customerID).
		Find(&purchaseOrders).Error

	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) GetPurchaseOrdersBySellerID(sellerID string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Where("CAST(seller_id AS TEXT) = ?", sellerID).
		Find(&purchaseOrders).Error

	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) GetAllPurchaseOrders() ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Find(&purchaseOrders).Error

	if err != nil {
		return nil, errors.New("error retrieving purchase orders")
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) SearchPurchaseOrdersByID(query string) ([]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		Where("CAST(id AS TEXT) LIKE ?", query+"%").
		Find(&purchaseOrders).Error

	if err != nil {
		return nil, err
	}
	return purchaseOrders, nil
}

func (r *PurchaseOrderRepository) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrder) error {
	var existingPurchaseOrder models.PurchaseOrder

	// Preload completo de todas las relaciones relevantes
	if err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		First(&existingPurchaseOrder, "id = ?", purchaseOrder.ID).Error; err != nil {
		return err
	}

	// Actualización de la orden de compra
	if err := r.DB.Model(&existingPurchaseOrder).Updates(purchaseOrder).Error; err != nil {
		return err
	}

	return nil
}

func (r *PurchaseOrderRepository) CreatePurchaseOrder(dto *dtos.CreatePurchaseOrderDTO, subtotal float64, total float64) (*models.PurchaseOrder, error) {
	purchaseOrder := &models.PurchaseOrder{
		SellerID:      dto.SellerID,
		CustomerID:    dto.CustomerID,
		ResponsibleID: dto.ResponsibleID,
		DateTime:      time.Now(),
		SubTotal:      subtotal,
		Total:         total,
		OrderStateID:  1, // Estado inicial
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Crear PurchaseOrder
	if err := tx.Create(purchaseOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Registrar PurchaseOrderItems
	for _, billingItem := range dto.Items {
		purchaseOrderItem := &models.PurchaseOrderItem{
			PurchaseOrderID: purchaseOrder.ID,
			ItemID:          billingItem.ID,
			Amount:          billingItem.Stock,
		}

		if err := tx.Create(purchaseOrderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Registrar descuentos en la relación many-to-many
	var discounts []models.DiscountType
	if len(dto.Discounts) > 0 {
		if err := tx.Where("id IN ?", dto.Discounts).Find(&discounts).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(purchaseOrder).Association("Discounts").Append(discounts); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Registrar impuestos en la relación many-to-many
	var taxes []models.TaxType
	if len(dto.Taxes) > 0 {
		if err := tx.Where("id IN ?", dto.Taxes).Find(&taxes).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(purchaseOrder).Association("Taxes").Append(taxes); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Cargar datos completos de la orden
	var fullPurchaseOrder models.PurchaseOrder
	if err := r.DB.Preload("Discounts").Preload("Taxes").Preload("Items.Item").First(&fullPurchaseOrder, purchaseOrder.ID).Error; err != nil {
		return nil, err
	}

	return &fullPurchaseOrder, nil
}

func (r *PurchaseOrderRepository) ChangePurchaseOrderState(id string, state string) (*models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder

	// Buscar solo por ID sin preloads inicialmente
	if err := r.DB.First(&purchaseOrder, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Convertir el string del estado a entero
	stateInt, err := strconv.Atoi(state)
	if err != nil {
		return nil, errors.New("invalid state ID: " + err.Error())
	}

	// Actualizar solo el campo 'order_state_id'
	if err := r.DB.Model(&purchaseOrder).Update("order_state_id", stateInt).Error; err != nil {
		return nil, err
	}

	// Recargar la orden completa con sus relaciones
	if err := r.DB.Preload("Seller").
		Preload("Responsible").
		Preload("Customer").
		Preload("OrderState").
		Preload("Items.Item").
		Preload("Discounts").
		Preload("Taxes").
		First(&purchaseOrder, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}
