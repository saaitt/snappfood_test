package repository

import (
	"github.com/saaitt/snappfood_test/orders/domain"
	"gorm.io/gorm"
)

func (r *repository) FindIfODRNotProcessed(orderId uint64) (bool, error) {
	var odr domain.OrderDelayReport
	result := r.db.
		Where("order_id = ?", orderId).
		Where("processed = false").
		First(&odr)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, result.Error
	}
	return false, nil
}
