package repositories

import (
	"fmt"
	"github.com/harlesbayu/kuncie/internal/shared/constants"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(dbConn *gorm.DB) *DBRepository {
	return &DBRepository{
		db: dbConn,
	}
}

func (r *DBRepository) DB() *gorm.DB {
	return r.db
}

func _generatePaginationQueries(tx *gorm.DB, page, perPage int64) {
	if page == 1 {
		page = 0
	}
	if perPage == 0 {
		perPage = constants.StorageMaximumListCount
	}
	offset := page*perPage - perPage
	tx.Limit(int(perPage)).Offset(int(offset))
}

func _getRealDBFieldName(fieldName string) (query string) {
	switch fieldName {
	case "name":
		query = "name"
	}

	return
}

func _generateFilterAndSortQueries(tx *gorm.DB, sortBy, sortType, searchBy, searchValue string) {
	if sortBy != "" {
		if sortBy == "name" {
			if sortType == "asc" {
				tx.Order(fmt.Sprintf("lower(%s) asc", _getRealDBFieldName(sortBy)))
			} else {
				tx.Order(fmt.Sprintf("lower(%s) desc", _getRealDBFieldName(sortBy)))
			}
		} else if _getRealDBFieldName(sortBy) != "" {
			if sortType == "asc" {
				tx.Order(fmt.Sprintf("%s asc", _getRealDBFieldName(sortBy)))
			} else if sortType == "desc" {
				tx.Order(fmt.Sprintf("%s desc", _getRealDBFieldName(sortBy)))
			}
		}
	} else {
		tx.Order("created_at desc")
	}

	if _getRealDBFieldName(searchBy) != "" && searchValue != "" {
		tx.Where(fmt.Sprintf("%s = ?", _getRealDBFieldName(searchBy)), searchValue)
	}
}
