package services

import (
	"totesbackend/models"
	"totesbackend/repositories"
)

type PurchaseOrderService struct {
	Repo *repositories.PurchaseOrderRepository
}

func NewPurchaseOrderService(repo *repositories.PurchaseOrderRepository) *PurchaseOrderService {
	return &PurchaseOrderService{Repo: repo}
}

func (s *PurchaseOrderService) GetPurchaseOrderByID(id string) (*models.PurchaseOrder, error) {
	return s.Repo.GetPurchaseOrderByID(id)
}

func (s *PurchaseOrderService) GetAllPurchaseOrders() ([]models.PurchaseOrder, error) {
	return s.Repo.GetAllPurchaseOrders()
}

func (s *PurchaseOrderService) SearchPurchaseOrdersByID(query string) ([]models.PurchaseOrder, error) {
	return s.Repo.SearchPurchaseOrdersByID(query)
}

func (s *PurchaseOrderService) GetPurchaseOrdersByCustomerID(customerID string) ([]models.PurchaseOrder, error) {
	return s.Repo.GetPurchaseOrdersByCustomerID(customerID)
}

func (s *PurchaseOrderService) GetPurchaseOrdersBySellerID(sellerID string) ([]models.PurchaseOrder, error) {
	return s.Repo.GetPurchaseOrdersBySellerID(sellerID)
}

func (s *PurchaseOrderService) UpdatePurchaseOrderState(id string, state int) (*models.PurchaseOrder, error) {
	return s.Repo.UpdatePurchaseOrderState(id, state)
}

func (s *PurchaseOrderService) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrder) error {
	return s.Repo.UpdatePurchaseOrder(purchaseOrder)
}

func (s *PurchaseOrderService) CreatePurchaseOrder(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	return s.Repo.CreatePurchaseOrder(purchaseOrder)
}
