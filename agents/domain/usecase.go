package domain

type UseCase interface {
	GetAssigned(agentID uint64) (*AgentTaskResponse, error)
	UpdateTask(req *TaskUpdateRequest) (*AgentTask, error)
}
