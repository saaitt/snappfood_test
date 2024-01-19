package repository

func (r *repository) FindActiveTask(item any, agentId uint64) error {
	result := r.db.
		Where("agent_id", agentId).
		Where("processed = NULL").
		First(&item)
	return result.Error
}
