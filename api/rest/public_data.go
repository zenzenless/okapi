package rest

import (
	"encoding/json"

	"net/http"

	"github.com/james-zhang-bing/okapi"
	requests "github.com/james-zhang-bing/okapi/requests/rest/public"
	responses "github.com/james-zhang-bing/okapi/responses/public_data"
)

// PublicData
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data
type PublicData struct {
	client *ClientRest
}

// NewPublicData returns a pointer to a fresh PublicData
func NewPublicData(c *ClientRest) *PublicData {
	return &PublicData{c}
}

// GetInstruments
// Retrieve a list of instruments with open contracts.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-instruments
func (c *PublicData) GetInstruments(req requests.GetInstruments) (response responses.GetInstruments, err error) {
	p := "/api/v5/public/instruments"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetDeliveryExerciseHistory
// Retrieve the estimated delivery price, which will only have a return value one hour before the delivery/exercise.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-instruments
func (c *PublicData) GetDeliveryExerciseHistory(req requests.GetDeliveryExerciseHistory) (response responses.GetDeliveryExerciseHistory, err error) {
	p := "/api/v5/public/delivery-exercise-history"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetOpenInterest
// Retrieve the total open interest for contracts on OKEx.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-open-interest
func (c *PublicData) GetOpenInterest(req requests.GetOpenInterest) (response responses.GetOpenInterest, err error) {
	p := "/api/v5/public/open-interest"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetLimitPrice
// Retrieve the highest buy limit and lowest sell limit of the instrument.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-limit-price
func (c *PublicData) GetLimitPrice(req requests.GetLimitPrice) (response responses.GetLimitPrice, err error) {
	p := "/api/v5/public/price-limit"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetOptionMarketData
// Retrieve option market data.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-option-market-data
func (c *PublicData) GetOptionMarketData(req requests.GetOptionMarketData) (response responses.GetOptionMarketData, err error) {
	p := "/api/v5/public/opt-summary"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetEstimatedDeliveryExercisePrice
// Retrieve the estimated delivery price which will only have a return value one hour before the delivery/exercise.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-estimated-delivery-Exercise-price
func (c *PublicData) GetEstimatedDeliveryExercisePrice(req requests.GetEstimatedDeliveryExercisePrice) (response responses.GetEstimatedDeliveryExercisePrice, err error) {
	p := "/api/v5/public/estimated-price"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetDiscountRateAndInterestFreeQuota
// Retrieve discount rate level and interest-free quota.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-discount-rate-and-interest-free-quota
func (c *PublicData) GetDiscountRateAndInterestFreeQuota(req requests.GetDiscountRateAndInterestFreeQuota) (response responses.GetDiscountRateAndInterestFreeQuota, err error) {
	p := "/api/v5/public/discount-rate-interest-free-quota"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetSystemTime
// Retrieve API server time.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-system-time
func (c *PublicData) GetSystemTime() (response responses.GetSystemTime, err error) {
	p := "/api/v5/public/time"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetLiquidationOrders
// Retrieve information on liquidation orders in the last 7 days.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-liquidation-orders
func (c *PublicData) GetLiquidationOrders(req requests.GetLiquidationOrders) (response responses.GetLiquidationOrders, err error) {
	p := "/api/v5/public/liquidation-orders"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetMarkPrice
// Retrieve mark price.
//
// We set the mark price based on the SPOT index and at a reasonable basis to prevent individual users from manipulating the market and causing the contract price to fluctuate.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-mark-price
func (c *PublicData) GetMarkPrice(req requests.GetMarkPrice) (response responses.GetMarkPrice, err error) {
	p := "/api/v5/public/mark-price"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetPositionTiers
// Position information，Maximum leverage depends on your borrowings and margin ratio.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-position-tiers
func (c *PublicData) GetPositionTiers(req requests.GetPositionTiers) (response responses.GetPositionTiers, err error) {
	p := "/api/v5/public/position-tiers"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetInterestRateAndLoanQuota
// Get margin interest rate
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-position-tiers
func (c *PublicData) GetInterestRateAndLoanQuota() (response responses.GetInterestRateAndLoanQuota, err error) {
	p := "/api/v5/public/interest-rate-loan-quota"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// GetUnderlying
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-underlying
func (c *PublicData) GetUnderlying(req requests.GetUnderlying) (response responses.GetUnderlying, err error) {
	p := "/api/v5/public/underlying"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// https://www.okx.com/docs-v5/zh/#rest-api-public-data-get-unit-convert
func (c *PublicData) ConvertContractUnit(req requests.ConvertContractUnit) (response responses.ConvertContractUnit, err error) {
	p := "/api/v5/public/convert-contract-coin"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// https://www.okx.com/docs-v5/zh/#rest-api-public-data-get-funding-rate
func (c *PublicData) GetFundingRate(req requests.GetFundingRate) (response responses.GetFundingRate, err error) {
	p := "/api/v5/public/funding-rate"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}

// https://www.okx.com/docs-v5/zh/#rest-api-public-data-get-funding-rate-history
func (c *PublicData) GetFundingRateHistory(req requests.GetFundingRateHistory) (response responses.GetFundingRateHistory, err error) {
	p := "/api/v5/public/funding-rate-history"
	m := okapi.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = Decode(d, &response)
	return
}
