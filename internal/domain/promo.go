package domain

import (
	"gopkg.in/guregu/null.v3"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Promo struct {
	ID          string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar(150);index" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	ProductId   string    `gorm:"type:uuid;default:null" json:"productId"`
	Products    Products  `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Type        string    `sql:"promo_type"`
	CreatedAt   time.Time `gorm:"default:now();index" json:"-"`
	UpdatedAt   null.Time `gorm:"default:null;index" json:"-"`
	DeletedAt   soft_delete.DeletedAt
}

type Discount struct {
	ID           string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	PromoId      string `gorm:"type:uuid;default:null" json:"promoId"`
	Promo        Promo  `gorm:"foreignKey:PromoId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalProduct int64  `gorm:"type:int;not null" json:"totalProduct"`
	Discount     int64  `gorm:"type:int;not null" json:"discount"`
	DeletedAt    soft_delete.DeletedAt
}

type FreeProduct struct {
	ID        string   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	PromoId   string   `gorm:"type:uuid;default:null" json:"promoId"`
	Promo     Promo    `gorm:"foreignKey:PromoId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductId string   `gorm:"type:uuid;default:null" json:"productId"`
	Products  Products `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt soft_delete.DeletedAt
}

type DiscountPrice struct {
	ID           string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	PromoId      string `gorm:"type:uuid;default:null" json:"promoId"`
	Promo        Promo  `gorm:"foreignKey:PromoId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalProduct int64  `gorm:"type:int;not null" json:"totalProduct"`
	TotalBuy     int64  `gorm:"type:int;not null" json:"totalBuy"`
	DeletedAt    soft_delete.DeletedAt
}

type PromoRepository interface {
	GetDiscountPrice(promoId string) (discountPrice DiscountPrice, err error)
	GetFreeProduct(promoId string) (freeProduct PromoFreeProduct, err error)
	GetDiscount(promoId string) (discount Discount, err error)
}
