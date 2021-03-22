package entity

import "time"

type CustomerBalance struct {
	CustomerID        string
	WalletID          string
	Balance           float64
	OrganizationID    string
	firstOrderDate    *time.Time
	IsBlocked         bool
	IsDeleted         bool
	IikoCardOrdersSum float64
	ReportDate        time.Time
	NeedSetCategory   bool
	CategoryID        string
	CorrectionSum     float64
}
