package entity

type Payment struct {
	PaymentType           PaymentType
	IsProcessedExternally bool
	IsPreliminary         bool
	IsExternal            bool
	AdditionalData        string
}

type PaymentType struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}
