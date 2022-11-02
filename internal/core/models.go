package core

import (
	"time"
)

type Model struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Account struct {
	Model
	Description string
	Active      bool
	Reference   string
}

func NewAccount(id string, description string, active bool, reference string) *Account {
	timestamp := time.Now()
	return &Account{
		Model: Model{
			ID:        id,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
		Description: description,
		Active:      active,
		Reference:   reference,
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
	Model
	Email     string
	Address   Address
	Forename  string
	Surname   string
	DOB       time.Time
	Reference string
}

func NewCustomer(id string, email string, address Address, forename string, surname string, dob time.Time, reference string) *Customer {
	timestamp := time.Now()
	return &Customer{
		Model: Model{
			ID:        id,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
		Email:     email,
		Address:   address,
		Forename:  forename,
		Surname:   surname,
		DOB:       dob,
		Reference: reference,
	}
}

type Card struct {
	Model
	Reference string
}

func NewCard(id string, reference string) *Card {
	timestamp := time.Now()
	return &Card{
		Model: Model{
			ID:        id,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		},
		Reference: reference,
	}
}
