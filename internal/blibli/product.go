package blibli

import (
	"encoding/json"
	"fmt"

	"github.com/bytecorner/blibli-integration/internal/models"
)

// HIT LIST PRODUCT ENDPOINT
func (c *Client) GetProductList(page, size int) (*models.ProductListResponse, error) {
	endpoint := fmt.Sprintf("/products?page=%d&size=%d", page, size)
	
	responseBytes, err := c.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product list: %v", err)
	}

	var blibliRes models.BlibliResponse
	if err := json.Unmarshal(responseBytes, &blibliRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal outer response: %v", err)
	}

	if !blibliRes.Success {
		return nil, fmt.Errorf("api returned error: %s - %s", blibliRes.ErrorCode, blibliRes.ErrorMessage)
	}

	dataBytes, err := json.Marshal(blibliRes.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to process response data: %v", err)
	}

	var productList models.ProductListResponse
	if err := json.Unmarshal(dataBytes, &productList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal product list: %v", err)
	}

	return &productList, nil
}