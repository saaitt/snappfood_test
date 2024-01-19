package usecase

import (
	"github.com/saaitt/snappfood_test/agents/domain"
)

func (uc *useCase) UpdateTask(req *domain.TaskUpdateRequest) (*domain.AgentTask, error) {
	var t domain.AgentTask
	if err := uc.repo.FindAllByFilter(&t, "id", req.ID); err != nil {
		return nil, err
	}

	t.Processed = req.Processed
	return &t, uc.repo.Update(&t)
}
