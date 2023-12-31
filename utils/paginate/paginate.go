package paginate

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page   int
	Limit  int
	Title  string
	Rating float64
}

func Paginate(page, length int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * length
		return db.Offset(offset).Limit(length)
	}
}

func GetParams(c echo.Context) Pagination {
	params := Pagination{
		Page:   cast.ToInt(c.QueryParam("page")),
		Limit:  cast.ToInt(c.QueryParam("limit")),
		Title:  c.QueryParam("title"),
		Rating: cast.ToFloat64(c.QueryParam("rating")),
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.Limit > 30 {
		params.Limit = 30
	}

	return params
}
