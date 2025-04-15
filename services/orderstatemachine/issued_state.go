package orderstatemachine

import (
	"errors"
	"strconv"
	"totesbackend/models"
)

type IssuedState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewIssuedState(context *OrderStateMachine) *IssuedState {
	return &IssuedState{
		context: context,
		state: &models.OrderStateType{
			ID:          1,
			Description: "IssuedState",
		},
	}
}

func (s *IssuedState) ChangeState(stateID string) error {

	switch stateID {
	case "2":
		return s.changeIssuedToInTransit(stateID)
	case "3":
		return s.changeIssuedToCancelled(stateID)
	default:
		return errors.New("no se puede cambiar de Issued a este estado: " + stateID)
	}

}

func (s *IssuedState) GetId() int {
	return s.state.ID
}

func (s *IssuedState) GetDescription() string {
	return s.state.Description
}

func (s *IssuedState) changeIssuedToInTransit(stateID string) error {

	for _, item := range s.context.PurchaseOrder.Items {
		itemIDStr := strconv.Itoa(item.ItemID)
		hasStock, err := s.context.ItemRepo.HasEnoughStock(itemIDStr, item.Amount)
		if err != nil {
			return errors.New("error checking stock for item with ID: " + itemIDStr + " - " + err.Error())
		}
		if !hasStock {
			return errors.New("insufficient stock for item with ID: " + itemIDStr)
		}
	}

	for _, item := range s.context.PurchaseOrder.Items {
		itemIDStr := strconv.Itoa(item.ItemID)
		if err := s.context.ItemRepo.SubtractItemsFromInventory(itemIDStr, item.Amount); err != nil {
			return errors.New("error subtracting stock for item with ID: " + itemIDStr + " - " + err.Error())
		}
	}

	orderIDStr := strconv.Itoa(s.context.PurchaseOrder.ID)
	newPurchaseOrder, err := s.context.PurchaseOrderRepo.ChangePurchaseOrderState(orderIDStr, stateID)
	if err != nil {
		return errors.New("error changing state of purchase order with ID: " + orderIDStr + " - " + err.Error())
	}
	s.context.PurchaseOrder = newPurchaseOrder
	s.context.CurrentState = NewInTransitState(s.context)

	return nil
}

func (s *IssuedState) changeIssuedToCancelled(stateID string) error {
	orderIDStr := strconv.Itoa(s.context.PurchaseOrder.ID)
	newPurchaseOrder, err := s.context.PurchaseOrderRepo.ChangePurchaseOrderState(orderIDStr, stateID)
	if err != nil {
		return errors.New("error cancelling purchase order with ID: " + orderIDStr + " - " + err.Error())
	}
	s.context.PurchaseOrder = newPurchaseOrder
	s.context.CurrentState = NewCancelledState(s.context)

	return nil
}
