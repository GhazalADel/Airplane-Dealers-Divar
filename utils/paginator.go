package utils

import (
	"gorm.io/gorm"
)

func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * PAGE_SIZE
		return db.Offset(offset).Limit(PAGE_SIZE)
	}
}
