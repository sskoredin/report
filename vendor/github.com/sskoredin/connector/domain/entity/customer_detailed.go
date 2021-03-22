package entity

import "time"

type CustomerDetailed struct {
	Anonymized                    bool            `json:"anonymized"`
	Birthday                      interface{}     `json:"birthday"`
	Cards                         []interface{}   `json:"cards"`
	Categories                    []interface{}   `json:"categories"`
	Comment                       interface{}     `json:"comment"`
	ConsentStatus                 int             `json:"consentStatus"`
	CultureName                   string          `json:"cultureName"`
	Email                         interface{}     `json:"email"`
	ID                            string          `json:"id"`
	IikoCardOrdersSum             float64         `json:"iikoCardOrdersSum"`
	IsBlocked                     bool            `json:"isBlocked"`
	IsDeleted                     bool            `json:"isDeleted"`
	MiddleName                    interface{}     `json:"middleName"`
	Name                          string          `json:"name"`
	PersonalDataConsentFrom       interface{}     `json:"personalDataConsentFrom"`
	PersonalDataConsentTo         interface{}     `json:"personalDataConsentTo"`
	PersonalDataProcessingFrom    interface{}     `json:"personalDataProcessingFrom"`
	PersonalDataProcessingTo      interface{}     `json:"personalDataProcessingTo"`
	Phone                         string          `json:"phone"`
	Rank                          interface{}     `json:"rank"`
	ReferrerID                    interface{}     `json:"referrerId"`
	Sex                           int             `json:"sex"`
	ShouldReceiveLoyaltyInfo      interface{}     `json:"shouldReceiveLoyaltyInfo"`
	ShouldReceiveOrderStatusInfo  interface{}     `json:"shouldReceiveOrderStatusInfo"`
	ShouldReceivePromoActionsInfo bool            `json:"shouldReceivePromoActionsInfo"`
	Surname                       interface{}     `json:"surname"`
	UserData                      interface{}     `json:"userData"`
	WalletBalances                []WalletBalance `json:"walletBalances"`
	FirstOrderDate                *time.Time      `json:"firstOrderDate"`
}

type WalletBalance struct {
	Balance float64 `json:"balance"`
	Wallet  struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		ProgramType string `json:"programType"`
		Type        string `json:"type"`
	} `json:"wallet"`
}

func (wb WalletBalance) Convert(customer CustomerDetailed, organizationID string, reportDate time.Time) CustomerBalance {
	return CustomerBalance{
		CustomerID:     customer.ID,
		WalletID:       wb.Wallet.ID,
		Balance:        wb.Balance,
		OrganizationID: organizationID,
		firstOrderDate: customer.FirstOrderDate,
		IsBlocked:      customer.IsBlocked,
		IsDeleted:      customer.IsDeleted,
		ReportDate:     reportDate,
	}
}

func (cd CustomerDetailed) Convert(organizationID string, reportDate time.Time) []CustomerBalance {
	balances := make([]CustomerBalance, len(cd.WalletBalances))
	for i, wb := range cd.WalletBalances {
		balances[i] = wb.Convert(cd, organizationID, reportDate)
	}

	return balances
}
