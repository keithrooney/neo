package core

import "time"

type ExternalReference struct {
	ID    string
	Model string
}

func NewExternalReference(id string, model string) *ExternalReference {
	return &ExternalReference{
		ID:    id,
		Model: model,
	}
}

type Account struct {
	ID          string
	Description string
	Active      bool
	Reference   ExternalReference
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAccount(id string, description string, active bool, reference ExternalReference) *Account {
	timestamp := time.Now()
	return &Account{
		ID:          id,
		Description: description,
		Active:      active,
		Reference:   reference,
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
	}
}

type County string

type Address struct {
	Line1  string
	Line2  string
	County County
}

func NewAddress(line1 string, line2 string, county County) *Address {
	return &Address{
		Line1:  line1,
		Line2:  line2,
		County: county,
	}
}

type Customer struct {
	ID        string
	Email     string
	Address   Address
	Forename  string
	Surname   string
	DOB       time.Time
	Reference ExternalReference
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(id string, email string, address Address, forename string, surname string, dob time.Time, reference ExternalReference) *Customer {
	timestamp := time.Now()
	return &Customer{
		ID:        id,
		Email:     email,
		Address:   address,
		Forename:  forename,
		Surname:   surname,
		DOB:       dob,
		Reference: reference,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
}

type Card struct {
	ID        string
	Reference ExternalReference
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCard(id string, reference ExternalReference) *Card {
	timestamp := time.Now()
	return &Card{
		ID:        id,
		Reference: reference,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
}
