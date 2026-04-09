package blibli

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytecorner/blibli-integration/internal/config"
)

type Client struct {
	config     *config.Config
	httpClient *http.Client
}

// CLIENT INITIALIZATION
func NewClient(cfg *config.Config) *Client {
	return &Client{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CORE HTTP REQUEST HANDLER
func (c *Client) DoRequest(method, endpoint string, body interface{}) ([]byte, error) {
	fullURL := fmt.Sprintf("%s%s%s", c.config.BaseURL, c.config.APIRouter, endpoint)
	
	var reqBody []byte
	var err error
	var md5Hash string

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		hash := md5.Sum(reqBody)
		md5Hash = base64.StdEncoding.EncodeToString(hash[:])
	}

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().UnixMilli()
	contentType := "application/json"
	
	signature := c.generateSignature(method, md5Hash, contentType, timestamp, c.config.APIRouter+endpoint)

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("API-Seller-Key", c.config.SellerKey)
	req.Header.Set("Authorization", fmt.Sprintf("Signature %s:%s", c.config.ClientID, signature))
	req.Header.Set("BusinessPartnerCode", c.config.StoreID)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("api request failed with status %d: %s", res.StatusCode, string(resBody))
	}

	return resBody, nil
}

// SECURITY & SIGNATURE GENERATION
func (c *Client) generateSignature(method, contentMD5, contentType string, timestamp int64, endpoint string) string {
	signatureString := fmt.Sprintf("%s\n%s\n%s\n%d\n%s", 
		method, 
		contentMD5, 
		contentType, 
		timestamp, 
		endpoint,
	)

	mac := hmac.New(sha256.New, []byte(c.config.ClientSecret))
	mac.Write([]byte(signatureString))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}