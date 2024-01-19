package usecase

import (
	"encoding/json"
	"github.com/saaitt/snappfood_test/agents/domain"
	od "github.com/saaitt/snappfood_test/orders/domain"
)

func (uc *useCase) GetAssigned(agentID uint64) (*domain.AgentTaskResponse, error) {
	var task domain.AgentTask
	if err := uc.repo.FindActiveTask(&task, agentID); err != nil {
		return nil, err
	}

	var message []byte
	select {
	case message = <-uc.consumer.MsgCh:
	default:
		return nil, domain.ErrorNoDelayReportIsOnQueue
	}
	var order od.Order
	if err := json.Unmarshal(message, &order); err != nil {
		return nil, err
	}
	var resp domain.AgentTaskResponse
	var delayReports []od.OrderDelayReport
	if err := uc.repo.FindAllByFilter(delayReports, "order_id", order.ID); err != nil {
		return nil, err
	}
	resp.Order = order
	task.Processed = false
	task.OrderID = order.ID
	task.DelayReportID = delayReports[0].ID
	if err := uc.repo.Create(&task); err != nil {
		return nil, err
	}
	resp.Task = task
	return &resp, nil
}
