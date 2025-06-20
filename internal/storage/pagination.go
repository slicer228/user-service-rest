package storage

import "gorm.io/gorm"

const (
	itemsPerPageMin = 10
	itemsPerPageMax = 100
)

func paginate(s *gorm.DB, page, items int) *gorm.DB {
	if page <= 0 {
		page = 1
	}

	if items <= 0 {
		items = itemsPerPageMin
	}
	if items > itemsPerPageMax {
		items = itemsPerPageMax
	}

	return s.Offset((page - 1) * itemsPerPageMin).Limit(items)
}
