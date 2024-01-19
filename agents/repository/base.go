package repository

import "fmt"

func (r *repository) Create(item any) error {
	return r.db.Create(item).Error
}

func (r *repository) FindByID(item any, id uint64) error {
	return r.db.
		Where("id = ?", id).
		First(&item).Error
}

func (r *repository) Update(item any) error {
	return r.db.Save(item).Error
}

func (r *repository) FindAllByFilter(items any, key, value any) error {
	return r.db.
		Where(fmt.Sprintf("%v = ?", key), value).
		Order("created_at DESC").
		Find(&items).Error
}
