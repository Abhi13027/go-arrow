// quotes.go
package arrow

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// QuoteRequest represents a request for quote data for a specific instrument.
//
// This struct is used to specify which instrument's quote data should be retrieved,
// identified by the exchange and trading symbol combination.
type QuoteRequest struct {
	Exchange string `json:"exchange"` // Name of the exchange where the instrument is traded (e.g., NSE, BSE, INDEX, NFO).
	Symbol   string `json:"symbol"`   // Trading symbol of the instrument (e.g., RELIANCE, NIFTY, BANKNIFTY).
	Mode     string `json:"mode"`     // Mode parameter (e.g., "full", "ltp", "quote" for data).
}

// QuoteLTP represents the Last Traded Price data for a trading instrument.
//
// This struct contains essential price information including the last traded price,
// closing price, and the unique token identifier for the instrument.
type QuoteLTP struct {
	Token int `json:"token"` // Unique token identifier for the trading instrument.
	Ltp   int `json:"ltp"`   // Last traded price of the instrument.
	Close int `json:"close"` // Closing price of the instrument from the previous trading session.
}

// QuoteLTPResponse represents the API response structure for LTP quotes.
//
// This struct encapsulates the response from the quotes LTP API endpoint,
// containing both the quote data and the API response status.
type QuoteLTPResponse struct {
	Data   []QuoteLTP `json:"data"`   // Array of QuoteLTP objects representing quote data for requested instruments.
	Status string     `json:"status"` // API response status indicating success or failure.
}

// GetQuotesLTP retrieves the Last Traded Price data for the specified instruments.
//
// This method sends a POST request to the "/info/quotes/ltp" endpoint to fetch
// real-time LTP data including the last traded price and closing price for
// the requested instruments.
//
// The method accepts multiple instruments in a single request, allowing efficient
// batch retrieval of quote data across different exchanges and symbols.
//
// Parameters:
//   - requests: A slice of QuoteRequest structs specifying the exchange and symbol
//     combinations for which LTP data should be retrieved.
//
// Returns:
//   - A slice of QuoteLTP structs containing the token, LTP, and closing price for each instrument.
//   - An error if the API request fails, authentication is invalid, or response parsing fails.
//
// Example usage:
//
//	requests := []arrow.QuoteRequest{
//	    {Exchange: "INDEX", Symbol: "NIFTY"},
//	    {Exchange: "INDEX", Symbol: "BANKNIFTY"},
//	    {Exchange: "NSE", Symbol: "RELIANCE"},
//	}
//
//	quotes, err := client.GetQuotesLTP(requests)
//	if err != nil {
//	    log.Printf("Failed to get quotes: %v", err)
//	    return
//	}
//
//	for _, quote := range quotes {
//	    fmt.Printf("Token: %d, LTP: %d, Close: %d\n",
//	               quote.Token, quote.Ltp, quote.Close)
//	}
func (c *Client) GetQuotes(requests []QuoteRequest, mode string) ([]QuoteLTP, error) {
	endpoint := fmt.Sprintf("info/quotes/%s", mode)

	// Marshal the request payload to JSON.
	payload, err := json.Marshal(requests)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal quote requests")
		return nil, err
	}

	// Send a POST request to the API to fetch LTP data.
	resp, err := c.request(endpoint, "POST", payload)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch quotes LTP")
		return nil, err
	}

	var result QuoteLTPResponse
	// Parse the JSON response into the QuoteLTPResponse struct.
	if err := json.Unmarshal(resp, &result); err != nil {
		log.Error().Err(err).Msg("Failed to parse quotes LTP response")
		return nil, err
	}

	// Check if the API response status indicates success.
	if result.Status != "success" {
		return nil, fmt.Errorf("quotes LTP retrieval failed with status: %s", result.Status)
	}

	log.Info().Int("count", len(result.Data)).Msg("Quotes LTP retrieved successfully")
	return result.Data, nil
}
