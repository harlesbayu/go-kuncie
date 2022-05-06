package product

import (
	"github.com/harlesbayu/kuncie/internal/domain"
	"github.com/harlesbayu/kuncie/internal/shared/utils"
)

type service struct {
	productRepo domain.ProductRepository
	promoRepo   domain.PromoRepository
}

func NewService(productRepo domain.ProductRepository, promoRepo domain.PromoRepository) *service {
	return &service{
		productRepo: productRepo,
		promoRepo:   promoRepo,
	}
}

func (s *service) ListProduct(sess *utils.Session, request *ListProductRequest) (list ListProductResponse, err error) {
	products, count, err := s.productRepo.GetListProduct(
		request.SearchBy,
		request.SearchValue,
		request.SortBy,
		request.SortType,
		request.Page,
		request.PerPage)

	response := make([]Product, 0)

	for _, product := range products {
		resp := Product{
			ID:                product.ID,
			SKU:               product.SKU,
			Name:              product.Name,
			Price:             product.Price,
			InventoryQuantity: product.InventoryQuantity,
		}

		response = append(response, resp)
	}

	list = ListProductResponse{
		Products: response,
		Page:     request.Page,
		PerPage:  request.PerPage,
		Count:    count,
	}

	return
}

func (s *service) ScanProduct(sess *utils.Session, request *ScanProductRequest) (res *ScanProductResponse, err error) {
	product, err := s.productRepo.GetProductAndPromo(request.ProductId)
	if err != nil || product.Quantity < 1 {
		return
	}

	var price float64
	products := make([]PromoAndProduct, 0)

	for i := 0; int64(i) < request.Quantity; i++ {
		products = append(products, PromoAndProduct{
			ProductName: product.ProductName,
		})
	}

	price = float64(request.Quantity) * product.Price

	if product.PromoType == "discount_price" {
		promo, errPromo := s.promoRepo.GetDiscountPrice(product.PromoId)
		if errPromo != nil {
			return
		}

		if promo.TotalProduct == request.Quantity {
			price = float64(promo.TotalBuy) * product.Price
		}
	} else if product.PromoType == "free_product" {
		promo, errPromo := s.promoRepo.GetFreeProduct(product.PromoId)
		if errPromo != nil {
			return
		}

		products = append(products, PromoAndProduct{
			ProductName: promo.ProductName,
		})
	} else if product.PromoType == "discount" {
		promo, errPromo := s.promoRepo.GetDiscount(product.PromoId)
		if errPromo != nil {
			return
		}

		if promo.TotalProduct == request.Quantity {
			price = price - (price * float64(promo.Discount) / 100)
		}
	}

	res = &ScanProductResponse{
		Products: products,
		Price:    price,
	}

	return
}
