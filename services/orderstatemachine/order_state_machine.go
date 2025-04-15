package orderstatemachine

import (
	"fmt"
	"totesbackend/models"
	"totesbackend/repositories"

	"gorm.io/gorm"
)

type OrderStateMachine struct {
	DB                *gorm.DB
	CurrentState      OrderState
	PurchaseOrder     *models.PurchaseOrder
	ItemRepo          *repositories.ItemRepository
	PurchaseOrderRepo *repositories.PurchaseOrderRepository
}

// NewStateMachine construye la máquina y setea el estado actual según el estado de la orden
func NewStateMachine(po *models.PurchaseOrder, ItemRepo *repositories.ItemRepository, PurchaseOrderRepo *repositories.PurchaseOrderRepository) (*OrderStateMachine, error) {
	sm := &OrderStateMachine{
		PurchaseOrder:     po,
		ItemRepo:          ItemRepo,
		PurchaseOrderRepo: PurchaseOrderRepo,
	}

	// Determinar estado inicial en base al OrderStateID de la orden
	switch po.OrderStateID {
	case 1:
		sm.CurrentState = NewIssuedState(sm)
	case 2:
		sm.CurrentState = NewInTransitState(sm)
	case 3:
		sm.CurrentState = NewCancelledState(sm)
	case 4:
		sm.CurrentState = NewApprovedState(sm)
	default:
		return nil, fmt.Errorf("unknown state: %d", po.OrderStateID)
	}

	return sm, nil
}

func (sm *OrderStateMachine) GetCurrentState() OrderState {
	return sm.CurrentState
}

func (sm *OrderStateMachine) ChangeState(stateID string) error {
	return sm.CurrentState.ChangeState(stateID)
}
