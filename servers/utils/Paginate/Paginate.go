package Paginate

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageUp, _ := strconv.Atoi(page)
		if pageUp == 0 {
			pageUp = 1
		}
		pageSizeInt, _ := strconv.Atoi(pageSize)
		switch {
		case pageSizeInt <= 0:
			pageSizeInt = 10
		}
		offset := (pageUp - 1) * pageSizeInt
		return db.Offset(offset).Limit(pageSizeInt)
	}
}
