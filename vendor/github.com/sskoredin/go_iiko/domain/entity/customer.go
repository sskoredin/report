package entity

import "time"

type Customer struct {
	Anonymized     bool        `json:"anonymized"`
	Birthday       interface{} `json:"birthday"`
	Comment        interface{} `json:"comment"`
	Email          interface{} `json:"email"`
	FirstOrderDate string      `json:"firstOrderDate"`
	ID             string      `json:"id"`
	LastVisitDate  string      `json:"lastVisitDate"`
	Name           string      `json:"name"`
	Phone          string      `json:"phone"`
	RegisteredDate interface{} `json:"registeredDate"`
	Sex            int         `json:"sex"`
	Surname        interface{} `json:"surname"`
	WhenCreated    time.Time   `json:"whenCreated"`
}
