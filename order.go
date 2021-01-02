package foxswap

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

const (
	OrderStateTrading  = "Trading"
	OrderStateRejected = "Rejected"
	OrderStateDone     = "Done"
)

type Order struct {
	ID          string          `json:"id,omitempty"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	State       string          `json:"state,omitempty"`
	PayAssetID  string          `json:"pay_asset_id,omitempty"`
	Funds       decimal.Decimal `json:"funds,omitempty"`
	FillAssetID string          `json:"fill_asset_id,omitempty"`
	Amount      decimal.Decimal `json:"amount,omitempty"`
	MinAmount   decimal.Decimal `json:"min_amount,omitempty"`
	PriceImpact decimal.Decimal `json:"price_impact,omitempty"`
	RouteAssets []string        `json:"route_assets,omitempty"`
	// route id
	Routes string `json:"routes,omitempty"`
}

type PreOrderReq struct {
	PayAssetID  string          `json:"pay_asset_id,omitempty"`
	FillAssetID string          `json:"fill_asset_id,omitempty"`
	Funds       decimal.Decimal `json:"funds,omitempty"`
	Amount      decimal.Decimal `json:"amount,omitempty"`
}

func ReadOrder(token, traceId string) (*Order, error) {
	const uri = "/api/orders/{id}"
	resp, err := Request(context.Background()).SetHeader("Authorization", "Bearer "+token).SetPathParams(map[string]string{
		"id": traceId,
	}).Get(uri)
	if err != nil {
		return nil, err
	}

	var order Order
	if _, err := UnmarshalResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
