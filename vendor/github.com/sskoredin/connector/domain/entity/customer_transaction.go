package entity

import (
	"github.com/sskoredin/go_iiko/domain/enum"
	"time"
)

type CustomerTransaction struct {
	CustomerID      string
	TransactionDate time.Time
	TransactionType enum.TransactionType
	Sum             float64
	BalanceBefore   float64
	BalanceAfter    float64
	IsWriteOff      bool
}
