package entity

type Transaction struct {
	PosBalanceBefore            *string `json:"PosBalanceBefore"`
	APIClientLogin              string  `json:"apiClientLogin"`
	BalanceAfter                float64 `json:"balanceAfter"`
	BalanceBefore               float64 `json:"balanceBefore"`
	CardNumbers                 *string `json:"cardNumbers"`
	CertificateBlockReason      *string `json:"certificateBlockReason"`
	CertificateCounteragent     *string `json:"certificateCounteragent"`
	CertificateCounteragentType *string `json:"certificateCounteragentType"`
	CertificateEmitentName      *string `json:"certificateEmitentName"`
	CertificateNominal          *string `json:"certificateNominal"`
	CertificateNumber           *string `json:"certificateNumber"`
	CertificateStatus           string  `json:"certificateStatus"`
	CertificateType             *string `json:"certificateType"`
	Comment                     *string `json:"comment"`
	CouponNumber                *string `json:"couponNumber"`
	CouponSeries                *string `json:"couponSeries"`
	IikoBizUser                 *string `json:"iikoBizUser"`
	MarketingCampaignName       *string `json:"marketingCampaignName"`
	OrderCreateDate             *string `json:"orderCreateDate"`
	OrderNumber                 *string `json:"orderNumber"`
	OrderSum                    int     `json:"orderSum"`
	PhoneNumber                 string  `json:"phoneNumber"`
	ProgramName                 string  `json:"programName"`
	TransactionCreateDate       string  `json:"transactionCreateDate"`
	TransactionSum              float64 `json:"transactionSum"`
	TransactionType             string  `json:"transactionType"`
}
