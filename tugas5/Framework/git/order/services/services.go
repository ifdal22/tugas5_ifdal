package services

import (
	"context"
	"fmt"
	"time"

	cm "tugas5/Framework/git/order/common"
)

type PaymentServices interface {
	StatusHandler(context.Context, cm.StatusRequest) cm.StatusResponse
}

type PaymentService struct{}

type ServiceMiddleware func(PaymentServices) PaymentServices

func utc() string {
	return time.Now().Format("2006-01-02 15:04:05 +0700")
}

func panicRecovery() {
	if r := recover(); r != nil {
		fmt.Printf("Recovering from panic: %v \n", r)
	}
}
