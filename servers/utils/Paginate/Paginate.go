package Paginate

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
