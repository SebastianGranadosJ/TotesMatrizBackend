package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type OrderStateTypeRepository struct {
	DB *gorm.DB
}

func NewOrderStateTypeRepository(db *gorm.DB) *OrderStateTypeRepository {
	return &OrderStateTypeRepository{DB: db}
}

func (r *OrderStateTypeRepository) GetOrderStateTypeByID(id string) (*models.OrderStateType, error) {
	var OrderStateType models.OrderStateType
	err := r.DB.First(&OrderStateType, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &OrderStateType, nil
}

func (r *OrderStateTypeRepository) GetAllOrderStateTypes() ([]models.OrderStateType, error) {
	var OrderStateTypes []models.OrderStateType
	err := r.DB.Find(&OrderStateTypes).Error
	if err != nil {
		return nil, err
	}
	return OrderStateTypes, nil
}
