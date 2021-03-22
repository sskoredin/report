package entity

import "time"

type Customer struct {
	//id in iiko
	Identifier     string     `json:"id"`
	Organization   string     `json:"organization"`
	Name           string     `json:"name"`
	Surname        string     `json:"surname"`
	Phone          string     `json:"phone"`
	Email          string     `json:"email"`
	Anonymized     bool       `json:"anonymized"`
	Sex            int        `json:"sex"`
	FirstOrderDate *time.Time `json:"firstOrderDate"`
	LastVisitDate  *time.Time `json:"lastVisitDate"`
	RegisteredDate *time.Time `json:"registeredDate"`
	WhenCreated    *time.Time `json:"whenCreated"`
}

type CustomerIiko struct {
	//id in iiko
	Identifier     string  `json:"id"`
	Organization   string  `json:"organization"`
	Name           string  `json:"name"`
	Surname        string  `json:"surname"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	Anonymized     bool    `json:"anonymized"`
	Sex            int     `json:"sex"`
	FirstOrderDate *string `json:"firstOrderDate"`
	LastVisitDate  *string `json:"lastVisitDate"`
	RegisteredDate *string `json:"registeredDate"`
	WhenCreated    *string `json:"whenCreated"`
}

func (c CustomerIiko) Convert() Customer {
	return Customer{
		Identifier:     c.Identifier,
		Organization:   c.Organization,
		Name:           c.Name,
		Surname:        c.Surname,
		Phone:          c.Phone,
		Email:          c.Email,
		Anonymized:     c.Anonymized,
		Sex:            c.Sex,
		FirstOrderDate: c.ParseDate(c.FirstOrderDate),
		LastVisitDate:  c.ParseDate(c.LastVisitDate),
		RegisteredDate: c.ParseDate(c.RegisteredDate),
		WhenCreated:    c.ParseDate(c.WhenCreated),
	}
}

func (c CustomerIiko) ParseDate(date *string) *time.Time {
	if date == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", *date)
	if err != nil {
		return nil
	}
	return &t
}
