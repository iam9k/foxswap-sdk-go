package foxswap

import (
	"context"

	"github.com/shopspring/decimal"
)

type Pair struct {
	BaseAssetID  string          `json:"base_asset_id,omitempty"`
	BaseAmount   decimal.Decimal `json:"base_amount,omitempty"`
	QuoteAssetID string          `json:"quote_asset_id,omitempty"`
	QuoteAmount  decimal.Decimal `json:"quote_amount,omitempty"`
	FeePercent   decimal.Decimal `json:"fee_percent,omitempty"`
	RouteID      int64           `json:"route_id,omitempty"`
	Liquidity    decimal.Decimal `json:"liquidity,omitempty"`
	Share        decimal.Decimal `json:"share,omitempty"`
	SwapMethod   string          `json:"swap_method,omitempty"`
	Version      int64           `json:"version,omitempty"`
}

func (pair *Pair) reverse() {
	pair.BaseAssetID, pair.QuoteAssetID = pair.QuoteAssetID, pair.BaseAssetID
	pair.BaseAmount, pair.QuoteAmount = pair.QuoteAmount, pair.BaseAmount
}

// ReadPairs list all pairs
func ReadPairs(token string) (pairs []*Pair, ts int64, err error) {
	const uri = "/api/pairs"
	resp, err := Request(context.Background()).SetHeader("Authorization", "Bearer "+token).Get(uri)
	if err != nil {
		return nil, ts, err
	}

	var body struct {
		Pairs                 []*Pair         `json:"pairs,omitempty"`
		Fee_24h               decimal.Decimal `json:"fee_24h,omitempty"`
		Pair_count            int64           `json:"pair_count,omitempty"`
		Transaction_count_24h decimal.Decimal `json:"transaction_count_24h,omitempty"`
		Ts                    int64           `json:"ts,omitempty"`
		Volume_24h            decimal.Decimal `json:"volume_24h,omitempty"`
	}

	if err = UnmarshalResponse(resp, &body); err != nil {
		return nil, ts, err
	}

	ts = body.Ts

	return body.Pairs, ts, nil
}
