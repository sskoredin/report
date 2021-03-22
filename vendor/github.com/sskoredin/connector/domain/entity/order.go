package entity

import "time"

type OrderQuery struct {
	Organization string        `json:"organization"`
	Customer     OrderCustomer `json:"customer"`
	Order        Order         `json:"order"`
}

type OrderCustomer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Order struct {
	ID            string       `json:"id"`
	Date          string       `json:"date"`
	Phone         string       `json:"phone"`
	IsSelfService string       `json:"isSelfService"`
	Items         []OrderItem  `json:"items"`
	Address       OrderAddress `json:"address"`
	PaymentItems  PaymentItem  `json:"paymentItems"`
}
type PaymentItem struct {
	Sum                   string        `json:"sum"`
	PaymentType           []interface{} `json:"paymentType"`
	IsProcessedExternally bool          `json:"isProcessedExternally"`
}

type OrderItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Code   string `json:"code"`
	Sum    int    `json:"sum"`
}

type OrderAddress struct {
	City      string `json:"city"`
	Street    string `json:"street"`
	Home      string `json:"home"`
	Housing   string `json:"housing"`
	Apartment string `json:"apartment"`
	Comment   string `json:"comment"`
}

type OrderFile struct {
	ID          uint
	FileName    string
	CreatedTime time.Time
	Size        uint64
	IsProcessed bool
	CreatedAt   time.Time
}
