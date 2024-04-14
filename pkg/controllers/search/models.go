package search

type Response struct {
	Message      string `json:"message"`
	MerchantName string `json:"merchant_name,omitempty"`
}
