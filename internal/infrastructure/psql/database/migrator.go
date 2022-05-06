package database

import (
	"github.com/harlesbayu/kuncie/internal/domain"
	"github.com/harlesbayu/kuncie/internal/infrastructure/psql/seeders"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) (err error) {
	var products = new(domain.Products)
	var promo = new(domain.Promo)
	var discount = new(domain.Discount)
	var freeProduct = new(domain.FreeProduct)
	var DiscountPrice = new(domain.DiscountPrice)

	if err = db.Migrator().DropTable(DiscountPrice); err != nil {
		return
	}
	if err = db.Migrator().DropTable(freeProduct); err != nil {
		return
	}
	if err = db.Migrator().DropTable(discount); err != nil {
		return
	}
	if err = db.Migrator().DropTable(promo); err != nil {
		return
	}
	if err = db.Migrator().DropTable(products); err != nil {
		return
	}

	if err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return
	}

	if err = db.Exec("DROP TYPE IF EXISTS promo_type cascade").Error; err != nil {
		return
	}

	err = db.Exec("CREATE TYPE promo_type AS ENUM('discount', 'free_product', 'discount_price')").Error
	if err != nil {
		return
	}

	if err = db.AutoMigrate(products); err != nil {
		return
	}
	if err = db.AutoMigrate(promo); err != nil {
		return
	}
	if err = db.AutoMigrate(discount); err != nil {
		return
	}
	if err = db.AutoMigrate(freeProduct); err != nil {
		return
	}
	if err = db.AutoMigrate(DiscountPrice); err != nil {
		return
	}

	seedProduct, seedPromo, seedDiscount, seedFreeProduct, seedDiscountPrice := seeders.CreateProduct()

	for _, seedProductVal := range seedProduct {
		db.Create(seedProductVal)
	}
	for _, seedPromoVal := range seedPromo {
		db.Create(seedPromoVal)
	}
	db.Create(seedDiscount)
	db.Create(seedFreeProduct)
	db.Create(seedDiscountPrice)

	return
}
