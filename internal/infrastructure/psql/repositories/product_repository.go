package repositories

import (
	"github.com/harlesbayu/kuncie/internal/domain"
)

func (r *DBRepository) GetListProduct(searchBy, searchValue, sortBy, sortType string, page, perPage int64) (products []domain.Products, count int64, err error) {
	tx := r.db.Model(domain.Products{}).
		Select("products.id, products.sku, products.name, products.price, products.inventory_quantity")

	if err = tx.Count(&count).Error; err != nil {
		return
	}

	_generatePaginationQueries(tx, page, perPage)
	if err = tx.Find(&products).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) GetProductAndPromo(productId string) (products domain.PromoAndProduct, err error) {
	tx := r.db.Table("products as p").
		Select(""+
			"p.id as product_id, "+
			"p.sku, "+
			"p.name as product_name, "+
			"p.price, "+
			"p.inventory_quantity as quantity, "+
			"p2.id as promo_id,"+
			"p2.type as promo_type").
		Joins("INNER JOIN promos as p2 on p.id = p2.product_id").
		Where("p.id = ?", productId)

	if err = tx.Find(&products).Error; err != nil {
		return
	}

	return
}
