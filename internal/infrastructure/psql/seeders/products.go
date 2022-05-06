package seeders

import "github.com/harlesbayu/kuncie/internal/domain"

func CreateProduct() (products []*domain.Products, promo []*domain.Promo, discount *domain.Discount, freeProduct *domain.FreeProduct, discountPrice *domain.DiscountPrice) {
	googleHome := &domain.Products{
		ID:                "53432b76-9ee6-4507-86f1-53af94b8bec4",
		SKU:               "120P90",
		Name:              "Google Home",
		Price:             49.99,
		InventoryQuantity: 10,
	}
	mackbookPro := &domain.Products{
		ID:                "3ca71c95-8fe2-4b18-8fde-54212c474ab7",
		SKU:               "43N23P",
		Name:              "MacBook Pro",
		Price:             5399.99,
		InventoryQuantity: 5,
	}
	alexaSpeaker := &domain.Products{
		ID:                "d21157cf-627e-490c-b094-cd743a4eb1a8",
		SKU:               "A304SD",
		Name:              "Alexa Speaker",
		Price:             109.50,
		InventoryQuantity: 10,
	}
	raspberry := &domain.Products{
		ID:                "9e340345-3ff5-4c5b-bdfe-424fadd0ac25",
		SKU:               "234234",
		Name:              "Raspberry Pi B",
		Price:             30.00,
		InventoryQuantity: 2,
	}
	products = append(products, googleHome)
	products = append(products, mackbookPro)
	products = append(products, alexaSpeaker)
	products = append(products, raspberry)

	promoMacbook := &domain.Promo{
		ID:        "b15bbbf9-4213-4451-93ca-a81fba03f6b0",
		ProductId: "3ca71c95-8fe2-4b18-8fde-54212c474ab7",
		Type:      "free_product",
	}

	promoGoogleHome := &domain.Promo{
		ID:        "9e09dbe9-fa6a-4e0b-a0a7-443eb76b57d1",
		ProductId: "53432b76-9ee6-4507-86f1-53af94b8bec4",
		Type:      "discount_price",
	}

	promoAlexaSpeaker := &domain.Promo{
		ID:        "34d6a758-d095-4272-885a-0177e1b8a20c",
		ProductId: "d21157cf-627e-490c-b094-cd743a4eb1a8",
		Type:      "discount",
	}

	promo = append(promo, promoMacbook)
	promo = append(promo, promoGoogleHome)
	promo = append(promo, promoAlexaSpeaker)

	discount = &domain.Discount{
		ID:           "2cab18e2-48ef-4f3b-8a50-657aa1a0eecb",
		PromoId:      "34d6a758-d095-4272-885a-0177e1b8a20c",
		TotalProduct: 3,
		Discount:     10,
	}

	freeProduct = &domain.FreeProduct{
		ID:        "dc1c2a1d-e995-41fe-a6e1-df99dfef6440",
		PromoId:   "b15bbbf9-4213-4451-93ca-a81fba03f6b0",
		ProductId: "9e340345-3ff5-4c5b-bdfe-424fadd0ac25",
	}

	discountPrice = &domain.DiscountPrice{
		ID:           "0b174a6c-2847-4060-a53a-a83ee2393ed3",
		PromoId:      "9e09dbe9-fa6a-4e0b-a0a7-443eb76b57d1",
		TotalProduct: 3,
		TotalBuy:     2,
	}
	return
}
