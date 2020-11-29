package middleware

import (
	"time"

	"context"

	cm "tugas5/Framework/git/order/common"
	"tugas5/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) StatusHandler(ctx context.Context, request cm.StatusRequest) cm.StatusResponse {

	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("StatusHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("StatusHandler begins")

	return mw.PaymentServices.StatusHandler(ctx, request)

}
