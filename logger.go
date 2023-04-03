package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggerService struct {
	next priceFetcher
}

func NewLoggingService(svc priceFetcher) PriceFetcher {
	return &loggerService{
		next: svc,
	}
}

func (l *loggerService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"took":  time.Since(begin),
			"error": err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now())

	return l.next.FetchPrice(ctx, ticker)
}
