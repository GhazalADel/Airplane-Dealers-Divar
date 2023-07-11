package utils

import (
	"Airplane-Divar/consts"

	"gorm.io/gorm"
)

func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * consts.PAGE_SIZE
		return db.Offset(offset).Limit(consts.PAGE_SIZE)
	}
}
