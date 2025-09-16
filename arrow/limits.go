package arrow

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// Limits represents the trading limits and margin details for a user.
type Limits struct {
	Data struct {
		Utilized           float64 `json:"utilized"`
		Allocated          float64 `json:"allocated"`
		NonCashOpen        int     `json:"nonCashOpen"`
		CashOpen           float64 `json:"cashOpen"`
		NonCashCurrent     int     `json:"nonCashCurrent"`
		CashCurrent        float64 `json:"cashCurrent"`
		CashUsed           int     `json:"cashUsed"`
		SpanMargin         int     `json:"spanMargin"`
		ExposureMargin     int     `json:"exposureMargin"`
		OtherMargin        int     `json:"otherMargin"`
		IntradayCashMargin float64 `json:"intradayCashMargin"`
		TotalMargin        float64 `json:"totalMargin"`
		RealizedPnl        int     `json:"realizedPnl"`
		UnrealizedPnl      int     `json:"unrealizedPnl"`
	} `json:"data"`
	Status string `json:"status"`
}

// GetLimits fetches the trading limits and margin details for the authenticated user.
//
// This function sends a GET request to the "/user/limits" endpoint to retrieve available margins,
// blocked funds, collateral, pending orders, and other financial details.
//
// Returns:
//   - A pointer to a Limits struct containing the trading limits if successful.
//   - An error if the request fails or the response cannot be parsed.
func (c *Client) GetLimits() (*Limits, error) {
	endpoint := "/user/limits"

	resp, err := c.request(endpoint, "GET", nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch trading limits")
		return nil, err
	}

	var result Limits
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Error().Err(err).Msg("Failed to parse trading limits response")
		return nil, err
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("failed to retrieve trading limits")
	}

	log.Info().Msg("Trading limits retrieved successfully")
	return &result, nil
}
