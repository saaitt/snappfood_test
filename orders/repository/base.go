package repository

func (r *repository) Create(item any) (any, error) {
	result := r.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (r *repository) FindByID(item any, id uint64) (any, error) {
	result := r.db.
		Where("id = ?", id).
		First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (r *repository) Update(item any) error {
	return r.db.Save(item).Error
}
