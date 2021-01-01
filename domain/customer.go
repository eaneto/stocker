package domain

import (
	"time"

	"github.com/google/uuid"
)

const CUSTOMER = "customer"

type Customer struct {
	Code uuid.UUID `json:"code"`
	Name string    `json:"name"`
}

type CustomerEntity struct {
	ID        uint
	Code      uuid.UUID
	Name      string
	CreatedAt time.Time
}
