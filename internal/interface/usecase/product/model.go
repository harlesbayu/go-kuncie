package product

type Product struct {
	ID                string  `json:"id""`
	SKU               string  `json:"sku"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	InventoryQuantity int64   `json:"inventoryQuantity"`
}

type ListProductRequest struct {
	SearchBy    string `json:"searchBy" validate:"required_with=SearchValue"`
	SearchValue string `json:"searchValue" validate:"required_with=SearchBy"`
	SortBy      string `json:"sortBy" validate:"required_with=SortType"`
	SortType    string `json:"sortType" validate:"required_with=SortBy,omitempty,oneof=asc desc"`
	Page        int64  `json:"page"`
	PerPage     int64  `json:"perPage"`
}

type ListProductResponse struct {
	Products []Product `json:"products"`
	Page     int64     `json:"page"`
	PerPage  int64     `json:"perPage"`
	Count    int64     `json:"count"`
}

type ScanProductRequest struct {
	ProductId string `json:"productId"`
	Quantity  int64  `json:"quantity"`
}

type ScanProductResponse struct {
	Products []PromoAndProduct `json:"products"`
	Price    float64           `json:"price"`
}

type PromoAndProduct struct {
	ProductName string `json:"productName"`
}
