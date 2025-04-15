package services

import (
	"errors"
	"strconv"
	"totesbackend/dtos"
	"totesbackend/models"
	"totesbackend/repositories"
	"totesbackend/services/orderstatemachine"
)

type PurchaseOrderService struct {
	PurchaseOrderRepo *repositories.PurchaseOrderRepository
	ItemRepo          *repositories.ItemRepository
	BillingService    *BillingService
}

func NewPurchaseOrderService(purchaseOrderRepo *repositories.PurchaseOrderRepository,
	itemRepo *repositories.ItemRepository, billingService *BillingService) *PurchaseOrderService {
	return &PurchaseOrderService{
		PurchaseOrderRepo: purchaseOrderRepo,
		ItemRepo:          itemRepo,
		BillingService:    billingService,
	}
}

func (s *PurchaseOrderService) CreatePurchaseOrder(dto *dtos.CreatePurchaseOrderDTO) (*models.PurchaseOrder, error) {
	// Verificar stock de los items
	for _, item := range dto.Items {
		itemID := strconv.Itoa(item.ID)
		hasStock, err := s.ItemRepo.HasEnoughStock(itemID, item.Stock)
		if err != nil {
			return nil, err
		}
		if !hasStock {
			return nil, errors.New("stock insuficiente para el item con ID " + itemID)
		}
	}

	// Calcular subtotal
	subtotal, err := s.BillingService.CalculateSubtotal(dto.Items)
	if err != nil {
		return nil, err
	}

	// Convertir IDs de descuentos e impuestos
	var discountIDs, taxIDs []string
	for _, id := range dto.Discounts {
		discountIDs = append(discountIDs, strconv.Itoa(id))
	}
	for _, id := range dto.Taxes {
		taxIDs = append(taxIDs, strconv.Itoa(id))
	}

	// Calcular total
	total, err := s.BillingService.CalculateTotal(discountIDs, taxIDs, dto.Items)
	if err != nil {
		return nil, err
	}

	// Crear la orden de compra
	purchaseOrder, err := s.PurchaseOrderRepo.CreatePurchaseOrder(dto, subtotal, total)
	if err != nil {
		return nil, err
	}

	return purchaseOrder, nil
}

func (s *PurchaseOrderService) GetPurchaseOrderByID(id string) (*models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.GetPurchaseOrderByID(id)
}

func (s *PurchaseOrderService) GetAllPurchaseOrders() ([]models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.GetAllPurchaseOrders()
}

func (s *PurchaseOrderService) SearchPurchaseOrdersByID(query string) ([]models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.SearchPurchaseOrdersByID(query)
}

func (s *PurchaseOrderService) GetPurchaseOrdersByCustomerID(customerID string) ([]models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.GetPurchaseOrdersByCustomerID(customerID)
}

func (s *PurchaseOrderService) GetPurchaseOrdersBySellerID(sellerID string) ([]models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.GetPurchaseOrdersBySellerID(sellerID)
}

func (s *PurchaseOrderService) ChangePurchaseOrderState(id string, targetStateID string) (*models.PurchaseOrder, error) {
	po, err := s.PurchaseOrderRepo.GetPurchaseOrderByID(id)
	if err != nil {
		return nil, errors.New("orden de compra no encontrada con ID: " + id)
	}

	stateMachine, err := orderstatemachine.NewStateMachine(po, s.ItemRepo, s.PurchaseOrderRepo)
	if err != nil {
		return nil, err
	}

	if err := stateMachine.ChangeState(targetStateID); err != nil {
		return nil, err
	}

	return stateMachine.PurchaseOrder, nil
}

func (s *PurchaseOrderService) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrder) error {
	return s.PurchaseOrderRepo.UpdatePurchaseOrder(purchaseOrder)
}

func (s *PurchaseOrderService) GetPurchaseOrdersByStateID(stateID string) ([]models.PurchaseOrder, error) {
	return s.PurchaseOrderRepo.GetPurchaseOrdersByStateID(stateID)
}
