package orderstatemachine

type OrderState interface {
	// ChangeState contiene la lógica para verificar y realizar una transición válida
	ChangeState(stateID string) error

	// GetId retorna el ID del estado (relacionado con OrderStateType.ID)
	GetId() int

	// GetDescription retorna el nombre legible del estado (OrderStateType.Description)
	GetDescription() string
}
