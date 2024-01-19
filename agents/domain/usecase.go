package domain

type UseCase interface {
	GetAssigned() (*AgentHistory, error)
}
