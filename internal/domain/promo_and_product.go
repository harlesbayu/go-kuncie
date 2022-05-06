package domain

type PromoAndProduct struct {
	ProductId   string  `json:"productId"`
	SKU         string  `json:"sku"`
	ProductName string  `json:"productName"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	PromoId     string  `json:"promoId"`
	PromoType   string  `json:"promoType"`
}

type PromoFreeProduct struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
}
