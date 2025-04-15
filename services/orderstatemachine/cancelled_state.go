package orderstatemachine

import (
	"errors"
	"totesbackend/models"
)

type CancelledState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewCancelledState(context *OrderStateMachine) *CancelledState {
	return &CancelledState{
		context: context,
		state: &models.OrderStateType{
			ID:          3,
			Description: "CancelledState",
		},
	}
}

func (s *CancelledState) ChangeState(stateID string) error {
	return errors.New("cannot change state: cancelled orders cannot transition to another state")
}

func (s *CancelledState) GetId() int {
	return s.state.ID
}

func (s *CancelledState) GetDescription() string {
	return s.state.Description
}
