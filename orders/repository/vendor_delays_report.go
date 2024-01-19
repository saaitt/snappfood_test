package repository

import (
	"fmt"
	"github.com/saaitt/snappfood_test/orders/domain"
	"time"
)

func (r *repository) VendorDelaysReport() ([]domain.VendorDelayReport, error) {

	sql := fmt.Sprintf(`select vendor_id,
									   count(odr.id) delay_counts
								from order_delay_reports odr
								INNER JOIN public.orders o
									on o.id = odr.order_id
								where o.created_at > '%s'
									group by o.vendor_id
									order by delay_counts DESC;
			`, time.Now().UTC().Add(-7*(time.Hour*24)).Format(time.DateOnly))

	var reports []domain.VendorDelayReport
	result := r.db.Raw(sql).Scan(&reports)
	if result.Error != nil {
		return nil, result.Error
	}
	return reports, nil
}
