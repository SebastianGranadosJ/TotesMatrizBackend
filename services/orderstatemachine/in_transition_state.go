package orderstatemachine

import (
	"errors"
	"strconv"
	"totesbackend/models"
)

type InTransitState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewInTransitState(context *OrderStateMachine) *InTransitState {
	return &InTransitState{
		context: context,
		state: &models.OrderStateType{
			ID:          2,
			Description: "InTransitState",
		},
	}
}

func (s *InTransitState) ChangeState(stateID string) error {
	switch stateID {
	case "3":
		return s.changeInTransitToCancelled(stateID)
	case "4":
		return s.changeInTransitToAccepted(stateID)
	default:
		return errors.New("cannot transition from InTransit to this state: " + stateID)
	}
}

func (s *InTransitState) GetId() int {
	return s.state.ID
}

func (s *InTransitState) GetDescription() string {
	return s.state.Description
}

func (s *InTransitState) changeInTransitToCancelled(stateID string) error {
	for _, item := range s.context.PurchaseOrder.Items {
		itemIDStr := strconv.Itoa(item.ItemID)
		if err := s.context.ItemRepo.ReturnItemsToInventory(itemIDStr, item.Amount); err != nil {
			return errors.New("failed to return stock for item with ID: " + itemIDStr + " - " + err.Error())
		}
	}

	orderIDStr := strconv.Itoa(s.context.PurchaseOrder.ID)
	newPurchaseOrder, err := s.context.PurchaseOrderRepo.ChangePurchaseOrderState(orderIDStr, stateID)
	if err != nil {
		return errors.New("failed to change purchase order state with ID: " + orderIDStr + " - " + err.Error())
	}
	s.context.PurchaseOrder = newPurchaseOrder
	s.context.CurrentState = NewCancelledState(s.context)
	return nil
}

func (s *InTransitState) changeInTransitToAccepted(stateID string) error {
	// TODO: Call invoice repo here

	orderIDStr := strconv.Itoa(s.context.PurchaseOrder.ID)
	newPurchaseOrder, err := s.context.PurchaseOrderRepo.ChangePurchaseOrderState(orderIDStr, stateID)
	if err != nil {
		return errors.New("failed to accept purchase order with ID: " + orderIDStr + " - " + err.Error())
	}

	s.context.PurchaseOrder = newPurchaseOrder
	s.context.CurrentState = NewApprovedState(s.context)
	return nil
}
