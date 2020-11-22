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
}

func (pair *Pair) reverse() {
	pair.BaseAssetID, pair.QuoteAssetID = pair.QuoteAssetID, pair.BaseAssetID
	pair.BaseAmount, pair.QuoteAmount = pair.QuoteAmount, pair.BaseAmount
}

// ReadPairs list all pairs
func ListPairs(token string) ([]*Pair, error) {
	const uri = "/api/pairs"
	resp, err := Request(context.Background()).SetHeader("Authorization", "Bearer "+token).Get(uri)
	if err != nil {
		return nil, err
	}

	var body struct {
		Pairs []*Pair `json:"pairs,omitempty"`
	}

	if err := UnmarshalResponse(resp, &body); err != nil {
		return nil, err
	}

	return body.Pairs, nil
}
