package domain

import (
	"gopkg.in/guregu/null.v3"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Products struct {
	ID                string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	SKU               string    `gorm:"type:varchar(50)" json:"sku"`
	Name              string    `gorm:"type:varchar(150);not null;index" json:"name"`
	Price             float64   `gorm:"type:float;not null;" json:"price"`
	InventoryQuantity int64     `gorm:"type:int;not null;" json:"inventoryQuantity"`
	CreatedAt         time.Time `gorm:"default:now();index" json:"-"`
	UpdatedAt         null.Time `gorm:"default:null;index" json:"-"`
	DeletedAt         soft_delete.DeletedAt
}

type ProductRepository interface {
	GetListProduct(searchBy, searchValue, sortBy, sortType string, page, perPage int64) (products []Products, count int64, err error)
	GetProductAndPromo(productId string) (products PromoAndProduct, err error)
}
