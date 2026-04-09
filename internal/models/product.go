package models

// API RESPONSE STRUCTURES
type BlibliResponse struct {
	Success      bool        `json:"success"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	RequestId    string      `json:"requestId"`
	Data         interface{} `json:"data,omitempty"`
}

type ProductListResponse struct {
	Content          []ProductItem `json:"content"`
	TotalElements    int           `json:"totalElements"`
	TotalPages       int           `json:"totalPages"`
	Number           int           `json:"number"`
	NumberOfElements int           `json:"numberOfElements"`
}

// PRODUCT ENTITY
type ProductItem struct {
	ProductSku   string  `json:"productSku"`
	ProductName  string  `json:"productName"`
	ItemSku      string  `json:"itemSku"`
	Price        float64 `json:"price"`
	SalePrice    float64 `json:"salePrice"`
	Stock        int     `json:"stock"`
	ProductStatus string `json:"productStatus"`
}