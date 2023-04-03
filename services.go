package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMock = map[string]float64{
	"BTC": 20.0,
	"ETH": 200.0,
	"Sol": 10.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMock[ticker]
	if !ok {
		return price, fmt.Errorf("The given ticker %s is not supported", ticker)
	}

	return price, nil
}
