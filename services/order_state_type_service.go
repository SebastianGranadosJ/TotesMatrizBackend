package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type OrderStateTypeService struct {
	Repo *repositories.OrderStateTypeRepository
}

func NewOrderStateTypeService(repo *repositories.OrderStateTypeRepository) *OrderStateTypeService {
	return &OrderStateTypeService{Repo: repo}
}

func (s *OrderStateTypeService) GetAllOrderStateTypes() ([]models.OrderStateType, error) {
	return s.Repo.GetAllOrderStateTypes()
}

func (s *OrderStateTypeService) GetOrderStateTypeByID(id string) (*models.OrderStateType, error) {
	return s.Repo.GetOrderStateTypeByID(id)
}
