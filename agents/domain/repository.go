package domain

type Repository interface {
	Create(item any) error
	Update(item any) error
	FindByID(item any, id uint64) error
	FindAllByFilter(item any, key, value any) error
	FindActiveTask(item any, agentId uint64) error
}
