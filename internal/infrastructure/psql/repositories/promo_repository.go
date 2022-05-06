package repositories

import "github.com/harlesbayu/kuncie/internal/domain"

func (r *DBRepository) GetDiscountPrice(promoId string) (discountPrice domain.DiscountPrice, err error) {
	tx := r.db.Model(domain.DiscountPrice{}).
		Select("total_product, total_buy").
		Where("promo_id = ?", promoId)

	if err = tx.Find(&discountPrice).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) GetFreeProduct(promoId string) (freeProduct domain.PromoFreeProduct, err error) {
	tx := r.db.Table("free_products as fp").
		Select("fp.product_id, p.name as product_name").
		Joins("INNER JOIN products as p ON fp.product_id = p.id").
		Where("fp.promo_id = ?", promoId)

	if err = tx.Find(&freeProduct).Error; err != nil {
		return
	}

	return
}

func (r *DBRepository) GetDiscount(promoId string) (discount domain.Discount, err error) {
	tx := r.db.Model(domain.Discount{}).
		Select("total_product, discount").
		Where("promo_id = ?", promoId)

	if err = tx.Find(&discount).Error; err != nil {
		return
	}

	return
}
