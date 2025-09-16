package arrow

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type MarginRequest struct {
	Exchange         string `json:"exchange"`
	Quantity         string `json:"quantity"`
	Price            string `json:"price"`
	Product          string `json:"product"`
	TransactionType  string `json:"transactionType"`
	Order            string `json:"order"`
	IncludePositions bool   `json:"includePositions"`
	Symbol           string `json:"symbol"`
}

type MarginResponse struct {
	Data struct {
		RequiredMargin       float64 `json:"requiredMargin"`
		MinimumCashRequired  float64 `json:"minimumCashRequired"`
		MarginUsedAfterTrade float64 `json:"marginUsedAfterTrade"`
		Charge               struct {
			Brokerage      float64 `json:"brokerage"`
			ExchangeTxnFee float64 `json:"exchangeTxnFee"`
			Gst            struct {
				Cgst  float64 `json:"cgst"`
				Igst  float64 `json:"igst"`
				Sgst  float64 `json:"sgst"`
				Total float64 `json:"total"`
			} `json:"gst"`
			Ipft           float64 `json:"ipft"`
			SebiCharges    float64 `json:"sebiCharges"`
			StampDuty      float64 `json:"stampDuty"`
			Total          float64 `json:"total"`
			TransactionTax float64 `json:"transactionTax"`
		} `json:"charge"`
	} `json:"data"`
	Status string `json:"status"`
}

func (c *Client) GetMargin(order MarginRequest) (*MarginResponse, error) {
	endpoint := "/margin/order"

	// Convert order details into JSON payload.
	payload, err := json.Marshal(order)
	if err != nil {
		log.Error().Err(err).Msg("Failed to serialize margin request")
		return nil, err
	}

	// Send the request to the API.
	resp, err := c.request(endpoint, "POST", []byte(payload))
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch margin")
		return nil, err
	}

	// Parse the JSON response into the OrderMargin struct.
	var result MarginResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Error().Err(err).Msg("Failed to parse margin response")
		return nil, err
	}

	return &result, nil
}
