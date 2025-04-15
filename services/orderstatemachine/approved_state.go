package orderstatemachine

import (
	"errors"
	"totesbackend/models"
)

type ApprovedState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewApprovedState(context *OrderStateMachine) *ApprovedState {
	return &ApprovedState{
		context: context,
		state: &models.OrderStateType{
			ID:          4,
			Description: "ApprovedState",
		},
	}
}

func (s *ApprovedState) ChangeState(stateID string) error {
	return errors.New("cannot change state: approved orders cannot transition to another state")
}

func (s *ApprovedState) GetId() int {
	return s.state.ID
}

func (s *ApprovedState) GetDescription() string {
	return s.state.Description
}
