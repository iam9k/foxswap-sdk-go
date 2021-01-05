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
	PayAmount   decimal.Decimal `json:"pay_amount,omitempty"`
	FillAssetID string          `json:"fill_asset_id,omitempty"`
	FillAmount  decimal.Decimal `json:"fill_amount,omitempty"`
	MinAmount   decimal.Decimal `json:"min_amount,omitempty"`
	RouteAssets []string        `json:"route_assets,omitempty"`
	// route id
	Routes string `json:"routes,omitempty"`

	// deprecated, Use PayAmount instead
	Funds decimal.Decimal `json:"funds,omitempty"`
	// deprecated, Use FillAmount instead
	Amount decimal.Decimal `json:"amount,omitempty"`
}

type PreOrderReq struct {
	PayAssetID  string `json:"pay_asset_id,omitempty"`
	FillAssetID string `json:"fill_asset_id,omitempty"`
	// pay amount 和 fill amount 二选一
	PayAmount  decimal.Decimal `json:"pay_amount,omitempty"`
	FillAmount decimal.Decimal `json:"fill_amount,omitempty"`
	// deprecated
	Funds decimal.Decimal `json:"funds,omitempty"`
	// deprecated
	Amount    decimal.Decimal `json:"amount,omitempty"`
	MinAmount decimal.Decimal `json:"min_amount,omitempty"`
}

func (req *PreOrderReq) fixDeprecated() {
	if req.Funds.IsPositive() {
		req.PayAmount = req.Funds
	}

	if req.Amount.IsPositive() {
		req.FillAmount = req.Amount
	}
}

func PreOrderWithPairs(pairs []*Pair, req *PreOrderReq) (*Order, error) {
	var (
		order *Order
		err   error
	)

	req.fixDeprecated()
	if req.PayAmount.IsPositive() {
		order, err = Route(pairs, req.PayAssetID, req.FillAssetID, req.PayAmount)
	} else {
		order, err = ReverseRoute(pairs, req.PayAssetID, req.FillAssetID, req.FillAmount)
	}

	if err != nil {
		return nil, err
	}

	order.fixDeprecated()
	return order, nil
}

// ReadOrder return order detail by id
// WithToken required
func ReadOrder(ctx context.Context, id string) (*Order, error) {
	const uri = "/api/orders/{id}"
	resp, err := Request(ctx).SetPathParams(map[string]string{
		"id": id,
	}).Get(uri)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := UnmarshalResponse(resp, &order); err != nil {
		return nil, err
	}

	order.fixDeprecated()
	return &order, nil
}

func (order *Order) fixDeprecated() {
	if order.PayAmount.IsPositive() {
		order.Funds = order.PayAmount
		order.Amount = order.FillAmount
	} else {
		order.PayAmount = order.Funds
		order.FillAmount = order.Amount
	}
}
