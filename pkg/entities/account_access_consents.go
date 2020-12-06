package entities

import "time"

// AccountAccessConsents entity
type AccountAccessConsents struct {
	Data  Data  `json:"Data"`
	Risk  Risk  `json:"Risk"`
	Links Links `json:"Links,omitempty"`
	Meta  Meta  `json:"Meta,omitempty"`
}

// Data entity
type Data struct {
	ConsentID               string    `json:"ConsentId,omitempty"`
	Status                  string    `json:"Status,omitempty"`
	StatusUpdateDateTime    time.Time `json:"StatusUpdateDateTime,omitempty"`
	CreationDateTime        time.Time `json:"CreationDateTime,omitempty"`
	Permissions             []string  `json:"Permissions"`
	ExpirationDateTime      time.Time `json:"ExpirationDateTime"`
	TransactionFromDateTime time.Time `json:"TransactionFromDateTime"`
	TransactionToDateTime   time.Time `json:"TransactionToDateTime"`
}

// Risk entity
type Risk struct {
}

// Links entity
type Links struct {
	Self string `json:"Self"`
}

// Meta entity
type Meta struct {
	TotalPages int `json:"TotalPages"`
}
